package server

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/dgryski/dgoogauth"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/api/oas"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
	"net/url"
	"rsc.io/qr"
	"strconv"
	"time"
)

var Cnf *conf.AppConfiguration

var (
	mfaCachePrefix   = "mfa:"
	tokenCachePrefix = "token:"
)

// Service is the server API for  service.
type Service struct {
	oas.UnimplementedServer
	DB    *ent.Client
	Redis *redis.Client
}

func (s Service) Auth(c *gin.Context, req *oas.AuthRequest) (res *oas.LoginResponse, err error) {
	pwd, err := s.checkPwd(c, req)
	if err != nil {
		return nil, err
	}
	profile, err := s.DB.UserLoginProfile.Query().Where(userloginprofile.UserID(pwd.UserID)).First(c)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if profile != nil {
		if profile.PasswordReset {
			res = &oas.LoginResponse{
				CallbackUrl: "/login/resetPassword",
			}
			return
		}
		if profile.MfaEnabled {
			res = &oas.LoginResponse{
				CallbackUrl: "/login/verifyFactor",
			}
			return
		}
	}
	return s.loginToken(c, pwd.UserID)
}

func (s Service) loginToken(c *gin.Context, uid int) (*oas.LoginResponse, error) {
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
	s.Redis.SetEX(c, tid, uid, tokenTTL)
	return &oas.LoginResponse{
		CallbackUrl:  "/login/success",
		AccessToken:  tstr,
		ExpiresIn:    int(tokenTTL.Seconds()),
		RefreshToken: trstr,
	}, nil
}

func (s Service) checkPwd(c *gin.Context, req *oas.AuthRequest) (*ent.UserPassword, error) {
	pwd, err := s.DB.UserPassword.Query().Where(
		userpassword.HasUserWith(user.PrincipalName(req.Body.Username)),
		userpassword.SceneEQ(userpassword.SceneLogin), userpassword.StatusEQ(typex.SimpleStatusActive),
	).Select(userpassword.FieldUserID, userpassword.FieldSalt, userpassword.FieldPassword).Only(c)
	if err != nil {
		return nil, err
	}
	sha := sha256.New()
	sha.Write([]byte(req.Body.Password))
	sha.Write([]byte(pwd.Salt))
	given := hex.EncodeToString(sha.Sum(nil))
	if given != pwd.Password {
		return nil, errors.New("password not match")
	}
	return pwd, nil
}

func (s Service) mfaPrepare(c *gin.Context, profile *ent.UserLoginProfile) (res *oas.LoginResponse, err error) {
	if !profile.MfaEnabled {
		return nil, nil
	}
	if profile.MfaEnabled && profile.MfaStatus != typex.SimpleStatusActive {
		return nil, errors.New("mfa not active")
	}
	sid := uuid.New().String()
	res = &oas.LoginResponse{
		CallbackUrl: "/login/verifyFactor",
		StateToken:  createStateToken(sid),
	}
	s.Redis.SetEX(c, mfaCachePrefix+sid, profile.ID, Cnf.Duration("auth.stateTokenTTL"))
	return
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
		"id":  id,
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
	id = tk.Claims.(jwt.MapClaims)["id"].(string)
	return
}

func (s Service) buildOTPC(secret string) *dgoogauth.OTPConfig {
	otpc := &dgoogauth.OTPConfig{
		Secret:      secret,
		WindowSize:  3,
		HotpCounter: 0,
	}
	return otpc
}

// MfaQRCode generate a QR code for MFA, the code is a png image
func (s Service) MfaQRCode(c *gin.Context, username, secret string) ([]byte, error) {
	issuer := c.Request.Host
	l, err := url.Parse("otpauth://totp/" + issuer + ":" + url.PathEscape(username))
	if err != nil {
		return nil, err
	}
	params := url.Values{}
	params.Add("secret", secret)
	params.Add("issuer", issuer)
	l.RawQuery = params.Encode()
	code, err := qr.Encode(l.String(), qr.Q)
	if err != nil {
		return nil, err
	}
	return code.PNG(), nil
}

func (s Service) Logout(c *gin.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) VerifyFactor(c *gin.Context, req *oas.VefityFactorRequest) (*oas.LoginResponse, error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token)
	if err != nil {
		return nil, err
	}
	pid, err := s.Redis.Get(c, mfaCachePrefix+id).Int()
	if err != nil {
		return nil, err
	}
	profile, err := s.DB.UserLoginProfile.Get(c, pid)
	if err != nil {
		return nil, err
	}
	if profile.MfaEnabled {
		otpc := s.buildOTPC(profile.MfaSecret)
		v, err := otpc.Authenticate(req.Body.OtpToken)
		if err != nil || !v {
			return nil, errors.New("invalid code")
		}
	}
	return s.loginToken(c, profile.ID)
}
