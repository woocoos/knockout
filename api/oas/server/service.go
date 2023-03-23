package server

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/api/oas"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
	"image/png"
	"strconv"
	"strings"
	"time"
)

var Cnf *conf.AppConfiguration

var (
	mfaCachePrefix   = "mfa:"
	tokenCachePrefix = "token:"
	resetCachePrefix = "reset:"

	CallBackUrlResetPassword = "/login/reset-password"
	CallBackUrlMFA           = "/login/verify-factor"
)

// Service is the server API for  service.
type Service struct {
	oas.UnimplementedServer
	DB    *ent.Client
	Redis *redis.Client
	Cache cache.Cache

	LogoutHandler func(*gin.Context)
}

func (s Service) Auth(ctx context.Context, req *oas.AuthRequest) (res *oas.LoginResponse, err error) {
	pwd, err := s.checkPwd(ctx, req)
	if err != nil {
		return nil, err
	}
	profile, err := s.DB.UserLoginProfile.Query().Where(userloginprofile.UserID(pwd.UserID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	if !profile.CanLogin {
		return nil, errors.New("user not allowed to login")
	}
	if profile.PasswordReset {
		return s.resetPasswordPrepare(ctx, profile)
	}
	if profile.MfaEnabled {
		return s.mfaPrepare(ctx, profile)
	}
	_ = updateLastLogin(ctx, s.DB.UserLoginProfile, profile.UserID)
	return s.loginToken(ctx, pwd.UserID)
}

func (s Service) VerifyFactor(ctx context.Context, req *oas.VerifyFactorRequest) (*oas.LoginResponse, error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token)
	if err != nil {
		return nil, err
	}

	var pid int
	if err = s.Cache.Get(ctx, mfaCachePrefix+id, &pid); err != nil {
		return nil, err
	}

	profile, err := s.DB.UserLoginProfile.Get(ctx, pid)
	if err != nil {
		return nil, err
	}
	if profile.MfaEnabled {
		if !totp.Validate(req.Body.OtpToken, profile.MfaSecret) {
			return nil, errors.New("invalid code")
		}
	}
	// no need use transaction
	err = updateLastLogin(ctx, s.DB.UserLoginProfile, profile.UserID)
	return s.loginToken(ctx, profile.UserID)
}

func (s Service) Logout(ctx context.Context) error {
	s.LogoutHandler(ctx.Value(gin.ContextKey).(*gin.Context))
	return nil
}

func (s Service) ResetPassword(ctx context.Context, req *oas.ResetPasswordRequest) (res *oas.LoginResponse, err error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token)
	if err != nil {
		return nil, err
	}

	var uid int
	cacheKey := resetCachePrefix + id
	if err = s.Cache.Get(ctx, cacheKey, &uid); err != nil {
		return nil, err
	}

	pwd := s.DB.UserPassword.Query().Where(userpassword.UserID(uid), userpassword.SceneEQ(userpassword.SceneLogin)).OnlyX(ctx)
	npwd := salt(req.Body.NewPassword, pwd.Salt)

	err = ecx.WithTx(ctx, func(ctx context.Context) (ecx.Transactor, error) {
		return s.DB.Tx(ctx)
	}, func(itx ecx.Transactor) error {
		tx := itx.(*ent.Tx)
		err = tx.UserPassword.UpdateOneID(pwd.ID).SetUpdatedBy(uid).SetPassword(npwd).Exec(ctx)
		if err != nil {
			return err
		}
		res, err = s.loginToken(ctx, uid)
		if err != nil {
			return err
		}
		err = updateLastLogin(ctx, tx.UserLoginProfile, uid)
		if err != nil {
			return err
		}
		s.Cache.Del(ctx, cacheKey) // lint:ignore
		return nil
	})
	return
}

func (s Service) resetPasswordPrepare(ctx context.Context, profile *ent.UserLoginProfile) (res *oas.LoginResponse, err error) {
	sid := uuid.New().String()
	res = &oas.LoginResponse{
		CallbackUrl: CallBackUrlResetPassword,
		StateToken:  createStateToken(sid),
	}
	err = s.Cache.Set(ctx, resetCachePrefix+sid, profile.UserID, Cnf.Duration("auth.stateTokenTTL"))
	return
}

func (s Service) mfaPrepare(ctx context.Context, profile *ent.UserLoginProfile) (res *oas.LoginResponse, err error) {
	if !profile.MfaEnabled {
		return nil, nil
	}
	if profile.MfaEnabled && profile.MfaStatus != typex.SimpleStatusActive {
		return nil, errors.New("mfa not active")
	}
	sid := uuid.New().String()
	res = &oas.LoginResponse{
		CallbackUrl: CallBackUrlMFA,
		StateToken:  createStateToken(sid),
	}
	err = s.Cache.Set(ctx, mfaCachePrefix+sid, profile.UserID, Cnf.Duration("auth.stateTokenTTL"))
	return
}

func updateLastLogin(ctx context.Context, pc *ent.UserLoginProfileClient, uid int) error {
	cip := ""
	if gctx := ctx.Value(gin.ContextKey); gctx != nil {
		cip = gctx.(*gin.Context).ClientIP()
	}
	// no mater what, update last login time and ip
	return pc.Update().Where(userloginprofile.UserID(uid)).
		SetLastLoginIP(cip).SetUpdatedBy(uid).SetPasswordReset(false).
		SetLastLoginAt(time.Now()).Exec(ctx)
}

func (s Service) loginToken(ctx context.Context, uid int) (*oas.LoginResponse, error) {
	usr := s.DB.User.GetX(ctx, uid)

	orgr, err := s.GetUserRootOrg(ctx, usr.ID)
	if err != nil {
		return nil, err
	}

	tokenTTL := Cnf.Duration("auth.jwt.tokenTTL")
	tid, tstr, err := createToken(strconv.Itoa(uid), tokenTTL)
	if err != nil {
		return nil, err
	}

	refreshTokenTTL := Cnf.Duration("auth.jwt.refreshTokenTTL")
	_, trstr, err := createToken(strconv.Itoa(uid), refreshTokenTTL)
	if err != nil {
		return nil, err
	}

	err = s.Cache.Set(ctx, tid, uid, tokenTTL)
	if err != nil {
		return nil, err
	}
	return &oas.LoginResponse{
		AccessToken:  tstr,
		ExpiresIn:    int(tokenTTL.Seconds()),
		RefreshToken: trstr,
		User: &oas.User{
			ID:          usr.ID,
			DisplayName: usr.DisplayName,
			DomainId:    orgr.ID,
			DomainName:  orgr.Name,
		},
	}, nil
}

func (s Service) checkPwd(ctx context.Context, req *oas.AuthRequest) (*ent.UserPassword, error) {
	pwd, err := s.DB.UserPassword.Query().Where(
		userpassword.HasUserWith(user.PrincipalName(req.Body.Username)),
		userpassword.SceneEQ(userpassword.SceneLogin), userpassword.StatusEQ(typex.SimpleStatusActive),
	).Select(userpassword.FieldUserID, userpassword.FieldSalt, userpassword.FieldPassword).Only(ctx)
	if err != nil {
		return nil, err
	}

	given := salt(req.Body.Password, pwd.Salt)
	if given != pwd.Password {
		return nil, errors.New("password not match")
	}
	return pwd, nil
}

func salt(ori, salt string) string {
	sha := sha256.New()
	sha.Write([]byte(ori))
	sha.Write([]byte(salt))
	given := hex.EncodeToString(sha.Sum(nil))
	return given
}

func createToken(subject string, ttl time.Duration) (tokenID, tokenStr string, err error) {
	tokenID = tokenCachePrefix + subject + ":" + uuid.New().String()
	claims := jwt.RegisteredClaims{
		Subject:   subject,
		ID:        tokenID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(Cnf.String("auth.jwt.signingMethod")), claims)
	tokenStr, err = token.SignedString([]byte(Cnf.String("auth.jwt.signingKey")))
	return
}

func createStateToken(id string) string {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(Cnf.Duration("auth.stateTokenTTL")).Unix(),
		"iat": time.Now().Unix(),
		"jti": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Cnf.String("auth.stateTokenSecret")))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func parseStateToken(token string) (id string, err error) {
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(Cnf.String("auth.stateTokenSecret")), nil
	})
	if err != nil {
		return
	}
	if !tk.Valid {
		err = errors.New("invalid token")
		return
	}
	id = tk.Claims.(jwt.MapClaims)["jti"].(string)
	return
}

// MfaQRCode generate a QR code for MFA, the code is a png image
func (s Service) MfaQRCode(ctx context.Context, userID int) ([]byte, error) {
	uorg, err := s.GetUserRootOrg(ctx, userID)
	if err != nil {
		return nil, err
	}

	issuer := uorg.Domain
	profile, err := s.DB.UserLoginProfile.Query().WithUser(func(query *ent.UserQuery) {
		query.Select(user.FieldPrincipalName)
	}).Select(userloginprofile.FieldUserID, userloginprofile.FieldMfaSecret).Where(userloginprofile.UserID(userID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	issuer = strings.ReplaceAll(issuer, ":", "-")
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: profile.Edges.User.PrincipalName,
		Secret:      []byte(profile.MfaSecret),
	})
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		return nil, err
	}
	err = png.Encode(&buf, img)
	return buf.Bytes(), err
}

func (s Service) GetUserRootOrg(ctx context.Context, uid int) (uorg *ent.Org, err error) {
	uorg, err = s.DB.OrgUser.Query().Where(orguser.UserIDEQ(uid)).
		QueryOrg().Unique(false).Where(org.KindEQ(org.KindRoot), org.StatusEQ(typex.SimpleStatusActive)).Order(ent.Desc(org.FieldPath)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return uorg, nil
}
