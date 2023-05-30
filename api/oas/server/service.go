package server

import (
	"ariga.io/entcache"
	"bytes"
	"context"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/tsingsun/woocoo/pkg/auth"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/api/oas"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
	"github.com/woocoos/knockout/service/resource"
	"github.com/woocoos/knockout/status"
	"image/png"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	Cnf *conf.AppConfiguration
)

var (
	mfaCachePrefix             = "mfa:"
	tokenCachePrefix           = "token:"
	resetCachePrefix           = "reset:"
	loginFailCachePrefix       = "loginfail:"
	forgetPwdBeginCachePrefix  = "forgetpwdbegin:"
	forgetPwdEmailCachePrefix  = "forgetpwdemail:"
	forgetPwdVerifyCachePrefix = "forgetpwdverify:"

	CallBackUrlResetPassword = "/login/reset-password"
	CallBackUrlMFA           = "/login/verify-factor"
	CallBackUrlCaptcha       = "/captcha"

	captchaWidth  = 200
	captchaHeight = 100
)

type Options struct {
	CaptchaCollectNum int             `json:"captchaCollectNum"` // # captcha memory store collect num
	CaptchaExpire     time.Duration   `json:"captchaExpire"`     // # captcha expire time
	CaptchaLength     int             `json:"captchaLength"`     // # captcha length
	CaptchaTimes      int             `json:"captchaTimes"`      // # if login fail times, captcha will force show
	CaptchaTTL        time.Duration   `json:"captchaTTL"`        // # captcha ttl
	LoginFailTimes    int             `json:"loginFailTimes"`    // # if login fail times, captcha will force show
	LoginFailLockTime time.Duration   `json:"loginFailLockTime"` // #  lock time while login upper to max fail times
	StateTokenTTL     time.Duration   `json:"stateTokenTTL"`     // # state token ttl
	StateTokenSecret  string          `json:"stateTokenSecret"`  // # state token secret
	JWT               auth.JWTOptions `json:"jwt"`
}

// Service is the server API for  service.
type Service struct {
	oas.UnimplementedServer
	Options
	DB    *ent.Client
	Cache cache.Cache

	LogoutHandler func(*gin.Context)

	captchaStore captcha.Store
}

func (s *Service) Apply(cnf *conf.AppConfiguration) error {
	s.Options.CaptchaCollectNum = 1000
	s.Options.CaptchaExpire = time.Minute * 2
	s.Options.CaptchaLength = 6
	s.Options.CaptchaTimes = 3
	s.Options.LoginFailTimes = 10
	s.Options.LoginFailLockTime = time.Hour * 24
	err := cnf.Sub("auth").Unmarshal(&s.Options)
	if err != nil {
		return err
	}
	// Initialize the captcha
	s.captchaStore = captcha.NewMemoryStore(s.CaptchaCollectNum, s.CaptchaExpire)
	captcha.SetCustomStore(s.captchaStore)
	return nil
}

func (s *Service) Captcha(ctx *gin.Context, req *oas.CaptchaRequest) (*oas.Captcha, error) {
	captchaId := captcha.NewLen(s.CaptchaLength)
	if req.Body.W == 0 {
		req.Body.W = captchaWidth
	}
	if req.Body.H == 0 {
		req.Body.H = captchaHeight
	}
	var buf bytes.Buffer
	err := captcha.WriteImage(&buf, captchaId, req.Body.W, req.Body.H)
	if err != nil {
		return nil, err
	}
	return &oas.Captcha{
		CaptchaId:    captchaId,
		CaptchaImage: "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes()),
	}, err
}

// Login login
func (s *Service) Login(ctx *gin.Context, req *oas.LoginRequest) (res *oas.LoginResponse, err error) {
	failCount := 0
	s.Cache.Get(ctx, loginFailCachePrefix+req.Body.Username, &failCount)
	if failCount >= s.CaptchaTimes && !captcha.VerifyString(req.Body.CaptchaId, req.Body.Captcha) {
		ctx.Status(http.StatusBadRequest)
		return nil, status.ErrCaptchaNotMatch
	}
	if failCount >= s.LoginFailTimes {
		ctx.Status(http.StatusForbidden)
		return nil, status.ErrLoginFailUpperLimit
	}
	pwd, err := s.checkPwd(ctx, req)
	if err != nil {
		if errors.Is(err, status.ErrMismatchPWD) {
			ctx.Status(http.StatusBadRequest)
			var errL error
			failCount, errL = s.logFailHandler(ctx, req.Body.Username, false)
			if errL != nil {
				return nil, errors.Join(err, errL)
			}
			if failCount >= s.CaptchaTimes {
				return &oas.LoginResponse{CallbackUrl: CallBackUrlCaptcha}, err
			}
			if failCount >= s.LoginFailTimes {
				return nil, status.ErrLoginFailUpperLimit
			}
		}
		return nil, err
	}

	profile, err := s.DB.UserLoginProfile.Query().Where(userloginprofile.UserID(pwd.UserID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	if !profile.CanLogin {
		return nil, errors.New("user not allowed to login")
	}

	if profile.MfaEnabled {
		return s.mfaPrepare(ctx, profile)
	}

	if profile.PasswordReset {
		return s.resetPasswordPrepare(ctx, profile)
	}

	_ = updateLastLogin(ctx, s.DB.UserLoginProfile, profile.UserID)
	return s.loginToken(ctx, pwd.UserID)
}

func (s *Service) VerifyFactor(ctx *gin.Context, req *oas.VerifyFactorRequest) (*oas.LoginResponse, error) {
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

	if profile.PasswordReset {
		return s.resetPasswordPrepare(ctx, profile)
	}

	// no need use transaction
	err = updateLastLogin(ctx, s.DB.UserLoginProfile, profile.UserID)
	return s.loginToken(ctx, profile.UserID)
}

func (s *Service) Logout(ctx *gin.Context) error {
	s.LogoutHandler(ctx)
	return nil
}

func (s *Service) ResetPassword(ctx *gin.Context, req *oas.ResetPasswordRequest) (res *oas.LoginResponse, err error) {
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
	npwd := resource.SaltSecret(req.Body.NewPassword, pwd.Salt)

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

func (s *Service) resetPasswordPrepare(ctx *gin.Context, profile *ent.UserLoginProfile) (res *oas.LoginResponse, err error) {
	sid := uuid.New().String()
	res = &oas.LoginResponse{
		CallbackUrl: CallBackUrlResetPassword,
		StateToken:  createStateToken(sid),
	}
	err = s.Cache.Set(ctx, resetCachePrefix+sid, profile.UserID, Cnf.Duration("auth.stateTokenTTL"))
	return
}

func (s *Service) mfaPrepare(ctx *gin.Context, profile *ent.UserLoginProfile) (res *oas.LoginResponse, err error) {
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
	err = s.Cache.Set(ctx, mfaCachePrefix+sid, profile.ID, Cnf.Duration("auth.stateTokenTTL"))
	return
}

func updateLastLogin(ctx *gin.Context, pc *ent.UserLoginProfileClient, uid int) error {
	cip := ctx.ClientIP()
	// no mater what, update last login time and ip
	return pc.Update().Where(userloginprofile.UserID(uid)).
		SetLastLoginIP(cip).SetUpdatedBy(uid).SetPasswordReset(false).
		SetLastLoginAt(time.Now()).Exec(ctx)
}

func (s *Service) loginToken(ctx *gin.Context, uid int) (*oas.LoginResponse, error) {
	usr := s.DB.User.GetX(ctx, uid)

	//orgr, err := s.GetUserRootOrg(ctx, usr.ID)
	//if err != nil {
	//	return nil, err
	//}

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

	var domains []*oas.Domain
	err = s.DB.Org.Query().Where(
		org.HasOrgUserWith(orguser.UserID(uid)),
		org.StatusEQ(typex.SimpleStatusActive),
		org.DomainNotNil(),
		org.KindEQ(org.KindRoot),
	).Select(org.FieldID, org.FieldName).Scan(ctx, &domains)
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
			Domains:     domains,
		},
	}, nil
}

func (s *Service) checkPwd(ctx *gin.Context, req *oas.LoginRequest) (*ent.UserPassword, error) {
	pwd, err := s.DB.UserPassword.Query().Where(
		userpassword.HasUserWith(user.HasIdentitiesWith(useridentity.Code(req.Body.Username))),
		userpassword.SceneEQ(userpassword.SceneLogin), userpassword.StatusEQ(typex.SimpleStatusActive),
	).Select(userpassword.FieldUserID, userpassword.FieldSalt, userpassword.FieldPassword).Only(entcache.Evict(ctx))
	if err != nil {
		return nil, err
	}

	given := resource.SaltSecret(req.Body.Password, pwd.Salt)
	if given != pwd.Password {
		return pwd, status.ErrMismatchPWD // return user id
	}
	return pwd, nil
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
func (s *Service) MfaQRCode(ctx *gin.Context, userID int, secret string) ([]byte, error) {
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
	if secret == "" {
		secret = profile.MfaSecret
	}
	secByte, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		return nil, err
	}
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: profile.Edges.User.PrincipalName,
		Secret:      secByte,
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

func (s *Service) BindMfaPrepare(ctx *gin.Context) (*oas.Mfa, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	pn, err := s.DB.User.Query().Where(user.ID(uid)).Select(user.FieldPrincipalName).String(ctx)
	if err != nil {
		return nil, err
	}
	sid := uuid.New().String()
	stateToken := createStateToken(sid)
	val := make(map[string]string)
	val["uid"] = strconv.Itoa(uid)
	val["secret"] = resource.GeneralMFASecret()
	err = s.Cache.Set(ctx, mfaCachePrefix+sid, val, Cnf.Duration("auth.stateTokenTTL"))
	if err != nil {
		return nil, err
	}

	secByte, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(val["secret"])
	if err != nil {
		return nil, err
	}
	uorg, err := s.GetUserRootOrg(ctx, uid)
	if err != nil {
		return nil, err
	}
	issuer := uorg.Domain
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: pn,
		Secret:      secByte,
	})
	return &oas.Mfa{
		QrCodeUri:     key.String(),
		PrincipalName: pn,
		Secret:        val["secret"],
		StateToken:    stateToken,
		StateTokenTTL: Cnf.Duration("auth.stateTokenTTL").Seconds(),
	}, nil
}

func (s *Service) BindMfa(ctx *gin.Context, req *oas.BindMfaRequest) (bool, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	id, err := parseStateToken(req.Body.StateToken)
	if err != nil {
		return false, err
	}
	//
	var val map[string]string
	if err = s.Cache.Get(ctx, mfaCachePrefix+id, &val); err != nil {
		return false, err
	}
	if val["uid"] != strconv.Itoa(uid) {
		return false, fmt.Errorf("invalid user")
	}
	if !totp.Validate(req.Body.OtpToken, val["secret"]) {
		return false, errors.New("invalid code")
	}
	err = s.DB.UserLoginProfile.UpdateOneID(uid).SetMfaEnabled(true).SetMfaStatus(typex.SimpleStatusActive).SetMfaSecret(val["secret"]).Exec(ctx)
	return err == nil, err
}

func (s *Service) UnBindMfa(ctx *gin.Context, req *oas.UnBindMfaRequest) (bool, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	up, err := s.DB.UserLoginProfile.Query().Where(userloginprofile.UserID(uid)).Only(ctx)
	if !totp.Validate(req.Body.OtpToken, up.MfaSecret) {
		return false, errors.New("invalid code")
	}
	err = s.DB.UserLoginProfile.UpdateOneID(up.ID).ClearMfaEnabled().ClearMfaStatus().ClearMfaSecret().Exec(ctx)
	return err == nil, err
}

func (s *Service) GetUserRootOrg(ctx *gin.Context, uid int) (uorg *ent.Org, err error) {
	uorg, err = s.DB.OrgUser.Query().Where(orguser.UserIDEQ(uid)).
		QueryOrg().Unique(false).Where(org.KindEQ(org.KindRoot), org.StatusEQ(typex.SimpleStatusActive)).Order(ent.Desc(org.FieldPath)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return uorg, nil
}

func (s *Service) logFailHandler(ctx *gin.Context, uid string, clear bool) (int, error) {
	key := loginFailCachePrefix + uid
	if clear {
		return 0, s.Cache.Del(ctx, key)
	}
	count := 0
	err := s.Cache.Get(ctx, key, &count)
	if err != nil && !s.Cache.IsNotFound(err) {
		return 0, err
	}
	count++
	err = s.Cache.Set(ctx, key, count, s.LoginFailLockTime) // 以账户锁定时间作为过期时间
	return count, err
}

// ForgetPwdBegin 忘记密码验证用户账户，开始修改密码流程
func (s *Service) ForgetPwdBegin(ctx *gin.Context, req *oas.ForgetPwdBeginRequest) (*oas.ForgetPwdBeginResponse, error) {
	// 验证验证码
	if !captcha.VerifyString(req.Body.CaptchaId, req.Body.Captcha) {
		return nil, status.ErrCaptchaNotMatch
	}
	// 查询用户
	u, err := s.DB.User.Query().Where(user.HasIdentitiesWith(useridentity.Code(req.Body.Username))).WithLoginProfile().Only(ctx)
	if err != nil {
		return nil, err
	}
	verifies := make([]*oas.ForgetPwdVerify, 0)
	if u.Edges.LoginProfile.MfaEnabled {
		verifies = append(verifies, &oas.ForgetPwdVerify{Kind: "mfa"})
	}
	if &u.Email != nil {
		verifies = append(verifies, &oas.ForgetPwdVerify{Kind: "email", Value: u.Email})
	}
	// 生成临时token
	sid := uuid.New().String()
	stateToken := createStateToken(sid)
	err = s.Cache.Set(ctx, forgetPwdBeginCachePrefix+sid, u.ID, Cnf.Duration("auth.stateTokenTTL"))
	if err != nil {
		return nil, err
	}
	return &oas.ForgetPwdBeginResponse{
		StateToken:    stateToken,
		StateTokenTTL: Cnf.Duration("auth.stateTokenTTL").Seconds(),
		Verifies:      verifies,
	}, nil
}

// ForgetPwdReset 忘记密码设置新密码
func (s *Service) ForgetPwdReset(ctx *gin.Context, req *oas.ForgetPwdResetRequest) (bool, error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token)
	if err != nil {
		return false, err
	}
	var uid int
	cacheKey := forgetPwdVerifyCachePrefix + id
	if err = s.Cache.Get(ctx, cacheKey, &uid); err != nil {
		return false, err
	}
	//
	pwd := s.DB.UserPassword.Query().Where(userpassword.UserID(uid), userpassword.SceneEQ(userpassword.SceneLogin)).OnlyX(ctx)
	npwd := resource.SaltSecret(req.Body.NewPassword, pwd.Salt)

	err = ecx.WithTx(ctx, func(ctx context.Context) (ecx.Transactor, error) {
		return s.DB.Tx(ctx)
	}, func(itx ecx.Transactor) error {
		tx := itx.(*ent.Tx)
		err = tx.UserPassword.UpdateOneID(pwd.ID).SetUpdatedBy(uid).SetPassword(npwd).Exec(ctx)
		if err != nil {
			return err
		}
		s.Cache.Del(ctx, cacheKey) // lint:ignore
		return nil
	})
	return false, nil
}

// ForgetPwdSendEmail 忘记密码 发送邮件验证码
func (s *Service) ForgetPwdSendEmail(ctx *gin.Context, req *oas.ForgetPwdSendEmailRequest) (string, error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token)
	if err != nil {
		return "", err
	}
	var uid int
	cacheKey := forgetPwdBeginCachePrefix + id
	if err = s.Cache.Get(ctx, cacheKey, &uid); err != nil {
		return "", err
	}

	captchaId := captcha.NewLen(6)
	// TODO 后续集成邮件功能时发送captchaId至邮件
	//digits := s.captchaStore.Get(captchaId, false)
	return captchaId, nil
}

// ForgetPwdVerifyEmail 忘记密码 邮件验证身份
func (s *Service) ForgetPwdVerifyEmail(ctx *gin.Context, req *oas.ForgetPwdVerifyEmailRequest) (*oas.ForgetPwdBeginResponse, error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token)
	if err != nil {
		return nil, err
	}
	var uid int
	cacheKey := forgetPwdBeginCachePrefix + id
	if err = s.Cache.Get(ctx, cacheKey, &uid); err != nil {
		return nil, err
	}
	// 验证验证码
	if !captcha.VerifyString(req.Body.CaptchaId, req.Body.Captcha) {
		return nil, status.ErrCaptchaNotMatch
	}
	sid := uuid.New().String()
	stateToken := createStateToken(sid)
	err = s.Cache.Set(ctx, forgetPwdVerifyCachePrefix+sid, uid, Cnf.Duration("auth.stateTokenTTL"))
	if err != nil {
		return nil, err
	}
	s.Cache.Del(ctx, cacheKey)
	return &oas.ForgetPwdBeginResponse{
		StateToken:    stateToken,
		StateTokenTTL: Cnf.Duration("auth.stateTokenTTL").Seconds(),
	}, nil
}

// ForgetPwdVerifyMfa 忘记密码 mfa验证身份
func (s *Service) ForgetPwdVerifyMfa(ctx *gin.Context, req *oas.ForgetPwdVerifyMfaRequest) (*oas.ForgetPwdBeginResponse, error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token)
	if err != nil {
		return nil, err
	}
	var uid int
	cacheKey := forgetPwdBeginCachePrefix + id
	if err = s.Cache.Get(ctx, cacheKey, &uid); err != nil {
		return nil, err
	}
	profile, err := s.DB.UserLoginProfile.Query().Where(userloginprofile.UserID(uid)).Only(ctx)
	if err != nil {
		return nil, err
	}
	// 验证mfa
	if profile.MfaEnabled {
		if !totp.Validate(req.Body.OtpToken, profile.MfaSecret) {
			return nil, errors.New("invalid code")
		}
	} else {
		return nil, fmt.Errorf("the MFA is disabled")
	}
	// 生成临时token
	sid := uuid.New().String()
	stateToken := createStateToken(sid)
	err = s.Cache.Set(ctx, forgetPwdVerifyCachePrefix+sid, profile.UserID, Cnf.Duration("auth.stateTokenTTL"))
	if err != nil {
		return nil, err
	}
	s.Cache.Del(ctx, cacheKey)
	return &oas.ForgetPwdBeginResponse{
		StateToken:    stateToken,
		StateTokenTTL: Cnf.Duration("auth.stateTokenTTL").Seconds(),
	}, nil
}
