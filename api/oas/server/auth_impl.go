package server

import (
	"bytes"
	"context"
	"encoding/base32"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/entcache"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/api/oas"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/oauthclient"
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
	mfaCachePrefix             = "mfa:"
	tokenCachePrefix           = "token:"
	resetCachePrefix           = "reset:"
	loginFailCachePrefix       = "loginfail:"
	forgetPwdBeginCachePrefix  = "forgetpwdbegin:"
	forgetPwdEmailCachePrefix  = "forgetpwdemail:"
	forgetPwdVerifyCachePrefix = "forgetpwdverify:"

	spmKeyPrefix = "spm:"

	CallBackUrlResetPassword = "/login/reset-password"
	CallBackUrlMFA           = "/login/verify-factor"
	CallBackUrlCaptcha       = "/captcha"

	captchaWidth  = 200
	captchaHeight = 100
)

type Options struct {
	CaptchaCollectNum int           `json:"captchaCollectNum"` // # captcha memory store collect num
	CaptchaExpire     time.Duration `json:"captchaExpire"`     // # captcha expire time
	CaptchaLength     int           `json:"captchaLength"`     // # captcha length
	CaptchaTimes      int           `json:"captchaTimes"`      // # if login fail times, captcha will force show
	CaptchaTTL        time.Duration `json:"captchaTTL"`        // # captcha ttl
	LoginFailTimes    int           `json:"loginFailTimes"`    // # if login fail times, captcha will force show
	LoginFailLockTime time.Duration `json:"loginFailLockTime"` // #  lock time while login upper to max fail times
	StateTokenTTL     time.Duration `json:"stateTokenTTL"`     // # state token ttl
	StateTokenSecret  string        `json:"stateTokenSecret"`  // # state token secret
	SpmTTL            time.Duration `json:"spmTTL"`            // # spm ttl
	JWT               struct {
		SigningMethod   string        `json:"signingMethod"`
		SigningKey      string        `json:"signingKey"`
		TokenTTL        time.Duration `json:"tokenTTL"`
		RefreshTokenTTL time.Duration `json:"refreshTokenTTL"`
	} `json:"jwt"`
}

type OASOptions struct {
	Msgsrv struct {
		BaseUrl string `yaml:"baseUrl"`
	} `yaml:"msgsrv"`
}

// AuthService is the server API for  service.
type AuthService struct {
	Options
	DB    *ent.Client
	Cache cache.Cache
	Cnf   *conf.AppConfiguration

	HttpClient *http.Client
	OASOptions OASOptions

	LogoutHandler func(*gin.Context)

	captchaStore captcha.Store
}

func (s *AuthService) Apply(cnf *conf.AppConfiguration) error {
	s.Options = Options{
		CaptchaCollectNum: 1000,
		CaptchaExpire:     time.Minute * 2,
		CaptchaLength:     6,
		CaptchaTimes:      3,
		LoginFailTimes:    10,
		LoginFailLockTime: time.Hour * 24,
		SpmTTL:            time.Second * 5,
	}
	s.Cnf = cnf
	err := cnf.Sub("auth").Unmarshal(&s.Options)
	if err != nil {
		return err
	}

	s.OASOptions = OASOptions{}
	err = cnf.Sub("oas").Unmarshal(&s.OASOptions)
	if err != nil {
		return err
	}

	// Initialize the captcha
	s.captchaStore = captcha.NewMemoryStore(s.CaptchaCollectNum, s.CaptchaExpire)
	captcha.SetCustomStore(s.captchaStore)
	return nil
}

func (s *AuthService) Captcha(ctx *gin.Context, req *oas.CaptchaRequest) (*oas.Captcha, error) {
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
func (s *AuthService) Login(ctx *gin.Context, req *oas.LoginRequest) (res *oas.LoginResponse, err error) {
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

func (s *AuthService) RefreshToken(ctx *gin.Context, req *oas.RefreshTokenRequest) (*oas.LoginResponse, error) {
	token, err := jwt.ParseWithClaims(req.RefreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		token.Method = jwt.GetSigningMethod(s.Options.JWT.SigningMethod)
		return []byte(s.Options.JWT.SigningKey), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	subject := token.Claims.(*jwt.RegisteredClaims).Subject
	uid, err := strconv.Atoi(subject)
	if err != nil {
		return nil, err
	}

	tid, tstr, err := createToken(strconv.Itoa(uid), s.Options, false)
	if err != nil {
		return nil, err
	}
	err = s.Cache.Set(ctx, tid, uid, cache.WithTTL(s.Options.JWT.TokenTTL))
	if err != nil {
		return nil, err
	}
	return &oas.LoginResponse{
		AccessToken: tstr,
		ExpiresIn:   int(s.Options.JWT.TokenTTL.Seconds()),
	}, nil
}

func (s *AuthService) VerifyFactor(ctx *gin.Context, req *oas.VerifyFactorRequest) (*oas.LoginResponse, error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token, s.Options)
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

func (s *AuthService) Logout(ctx *gin.Context) error {
	s.LogoutHandler(ctx)
	return nil
}

func (s *AuthService) ResetPassword(ctx *gin.Context, req *oas.ResetPasswordRequest) (res *oas.LoginResponse, err error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token, s.Options)
	if err != nil {
		return nil, err
	}

	var uid int
	cacheKey := resetCachePrefix + id
	if err = s.Cache.Get(ctx, cacheKey, &uid); err != nil {
		return nil, err
	}

	pwd := s.DB.UserPassword.Query().Where(userpassword.UserID(uid),
		userpassword.SceneEQ(userpassword.SceneLogin)).OnlyX(ctx)
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

func (s *AuthService) resetPasswordPrepare(ctx *gin.Context, profile *ent.UserLoginProfile) (res *oas.LoginResponse, err error) {
	sid := uuid.New().String()
	res = &oas.LoginResponse{
		CallbackUrl: CallBackUrlResetPassword,
		StateToken:  createStateToken(sid, s.Options),
	}
	err = s.Cache.Set(ctx, resetCachePrefix+sid, profile.UserID, cache.WithTTL(s.Options.StateTokenTTL))
	return
}

func (s *AuthService) mfaPrepare(ctx *gin.Context, profile *ent.UserLoginProfile) (res *oas.LoginResponse, err error) {
	if !profile.MfaEnabled {
		return nil, nil
	}
	if profile.MfaEnabled && profile.MfaStatus != typex.SimpleStatusActive {
		return nil, errors.New("mfa not active")
	}
	sid := uuid.New().String()
	res = &oas.LoginResponse{
		CallbackUrl: CallBackUrlMFA,
		StateToken:  createStateToken(sid, s.Options),
	}
	err = s.Cache.Set(ctx, mfaCachePrefix+sid, profile.ID, cache.WithTTL(s.Cnf.Duration("auth.stateTokenTTL")))
	return
}

func updateLastLogin(ctx *gin.Context, pc *ent.UserLoginProfileClient, uid int) error {
	cip := ctx.ClientIP()
	// no mater what, update last login time and ip
	return pc.Update().Where(userloginprofile.UserID(uid)).
		SetLastLoginIP(cip).SetUpdatedBy(uid).SetPasswordReset(false).
		SetLastLoginAt(time.Now()).Exec(ctx)
}

func (s *AuthService) loginToken(ctx *gin.Context, uid int) (*oas.LoginResponse, error) {
	usr := s.DB.User.GetX(ctx, uid)

	tid, tstr, err := createToken(strconv.Itoa(uid), s.Options, false)
	if err != nil {
		return nil, err
	}

	_, trstr, err := createToken(strconv.Itoa(uid), s.Options, true)
	if err != nil {
		return nil, err
	}

	err = s.Cache.Set(ctx, tid, uid, cache.WithTTL(s.Options.JWT.TokenTTL))
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
		ExpiresIn:    int(s.Options.JWT.TokenTTL.Seconds()),
		RefreshToken: trstr,
		User: &oas.User{
			ID:           usr.ID,
			DisplayName:  usr.DisplayName,
			AvatarFileId: usr.AvatarFileID,
			Domains:      domains,
		},
	}, nil
}

func (s *AuthService) checkPwd(ctx *gin.Context, req *oas.LoginRequest) (*ent.UserPassword, error) {
	pwd, err := s.DB.UserPassword.Query().Where(
		userpassword.HasUserWith(user.HasIdentitiesWith(useridentity.Code(req.Body.Username))),
		userpassword.SceneEQ(userpassword.SceneLogin), userpassword.StatusEQ(typex.SimpleStatusActive),
	).Select(userpassword.FieldUserID, userpassword.FieldSalt, userpassword.FieldPassword).Only(entcache.Skip(ctx))
	if err != nil {
		return nil, err
	}

	given := resource.SaltSecret(req.Body.Password, pwd.Salt)
	if given != pwd.Password {
		return pwd, status.ErrMismatchPWD // return user id
	}
	return pwd, nil
}

func createToken(subject string, opts Options, refresh bool) (tokenID, tokenStr string, err error) {
	tokenID = tokenCachePrefix + subject + ":" + uuid.New().String()
	ttl := opts.JWT.TokenTTL
	if refresh {
		ttl = opts.JWT.RefreshTokenTTL
	}
	claims := jwt.RegisteredClaims{
		Subject:   subject,
		ID:        tokenID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(opts.JWT.SigningMethod), claims)
	tokenStr, err = token.SignedString([]byte(opts.JWT.SigningKey))
	return
}

func createStateToken(id string, opts Options) string {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(opts.StateTokenTTL).Unix(),
		"iat": time.Now().Unix(),
		"jti": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(opts.StateTokenSecret))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func parseStateToken(token string, opts Options) (id string, err error) {
	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(opts.StateTokenSecret), nil
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
func (s *AuthService) MfaQRCode(ctx *gin.Context, userID int, secret string) ([]byte, error) {
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

func (s *AuthService) BindMfaPrepare(ctx *gin.Context) (*oas.Mfa, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	pn, err := s.DB.User.Query().Where(user.ID(uid)).Select(user.FieldPrincipalName).String(ctx)
	if err != nil {
		return nil, err
	}
	sid := uuid.New().String()
	stateToken := createStateToken(sid, s.Options)
	val := map[string]string{
		"uid":    strconv.Itoa(uid),
		"secret": resource.GeneralMFASecret(),
	}
	err = s.Cache.Set(ctx, mfaCachePrefix+sid, val, cache.WithTTL(s.Options.StateTokenTTL))
	if err != nil {
		return nil, err
	}

	secByte, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(val["secret"])
	if err != nil {
		return nil, err
	}
	var tid int
	if tid, err = s.tryGetTenantID(ctx); err != nil {
		return nil, err
	}
	uorg, err := s.DB.Org.Query().Where(org.ID(tid)).Only(ctx)
	issuer := uorg.Domain
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: pn,
		Secret:      secByte,
	})
	if err != nil {
		return nil, err
	}
	return &oas.Mfa{
		QrCodeUri:     key.String(),
		PrincipalName: pn,
		Secret:        val["secret"],
		StateToken:    stateToken,
		StateTokenTTL: s.Options.StateTokenTTL.Seconds(),
	}, nil
}

func (s *AuthService) BindMfa(ctx *gin.Context, req *oas.BindMfaRequest) (bool, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	id, err := parseStateToken(req.Body.StateToken, s.Options)
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

func (s *AuthService) UnBindMfa(ctx *gin.Context, req *oas.UnBindMfaRequest) (bool, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	up, err := s.DB.UserLoginProfile.Query().Where(userloginprofile.UserID(uid)).Only(ctx)
	if !totp.Validate(req.OtpToken, up.MfaSecret) {
		return false, errors.New("invalid code")
	}
	err = s.DB.UserLoginProfile.UpdateOneID(up.ID).ClearMfaEnabled().ClearMfaStatus().ClearMfaSecret().Exec(ctx)
	return err == nil, err
}

func (s *AuthService) GetUserRootOrg(ctx *gin.Context, uid int) (uorg *ent.Org, err error) {
	uorg, err = s.DB.OrgUser.Query().Where(orguser.UserIDEQ(uid)).
		QueryOrg().Unique(false).Where(org.KindEQ(org.KindRoot), org.StatusEQ(typex.SimpleStatusActive)).Order(ent.Desc(org.FieldPath)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return uorg, nil
}

func (s *AuthService) logFailHandler(ctx *gin.Context, uid string, clear bool) (int, error) {
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
	err = s.Cache.Set(ctx, key, count, cache.WithTTL(s.LoginFailLockTime)) // 以账户锁定时间作为过期时间
	return count, err
}

// ForgetPwdBegin 忘记密码验证用户账户，开始修改密码流程
func (s *AuthService) ForgetPwdBegin(ctx *gin.Context, req *oas.ForgetPwdBeginRequest) (*oas.ForgetPwdBeginResponse, error) {
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
		verifies = append(verifies, &oas.ForgetPwdVerify{Kind: "email", Value: resource.MaskEmail(u.Email)})
	}
	// 生成临时token
	sid := uuid.New().String()
	stateToken := createStateToken(sid, s.Options)
	err = s.Cache.Set(ctx, forgetPwdBeginCachePrefix+sid, u.ID, cache.WithTTL(s.Options.StateTokenTTL))
	if err != nil {
		return nil, err
	}
	return &oas.ForgetPwdBeginResponse{
		StateToken:    stateToken,
		StateTokenTTL: s.Options.StateTokenTTL.Seconds(),
		Verifies:      verifies,
	}, nil
}

// ForgetPwdReset 忘记密码设置新密码
func (s *AuthService) ForgetPwdReset(ctx *gin.Context, req *oas.ForgetPwdResetRequest) (bool, error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token, s.Options)
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
	if err != nil {
		return false, err
	}
	return true, nil
}

// ForgetPwdSendEmail 忘记密码 发送邮件验证码
func (s *AuthService) ForgetPwdSendEmail(ctx *gin.Context, req *oas.ForgetPwdSendEmailRequest) (string, error) {
	token := req.StateToken
	id, err := parseStateToken(token, s.Options)
	if err != nil {
		return "", err
	}
	var uid int
	cacheKey := forgetPwdBeginCachePrefix + id
	if err = s.Cache.Get(ctx, cacheKey, &uid); err != nil {
		return "", err
	}
	// 生成验证码
	captchaId := captcha.NewLen(6)
	digits := s.captchaStore.Get(captchaId, false)
	captchaCode := ""
	for _, v := range digits {
		captchaCode = captchaCode + strconv.Itoa(int(v))
	}
	usr, err := s.DB.User.Get(ctx, uid)
	if err != nil {
		return "", err
	}
	if usr.Email == "" {
		return "", fmt.Errorf("email is nil")
	}
	uorg, err := s.GetUserRootOrg(ctx, usr.ID)
	if err != nil {
		return "", err
	}

	params := []postAlertsInput{
		{
			Annotations: map[string]string{
				"to":            usr.Email,
				"displayName":   usr.DisplayName,
				"captchaCode":   captchaCode,
				"captchaExpire": strconv.Itoa(int(s.CaptchaExpire.Minutes())),
			},
			Labels: map[string]string{
				"receiver":  "email",
				"alertname": "SendCaptchaCode",
				"tenant":    strconv.Itoa(uorg.ID),
				"timestamp": strconv.Itoa(int(time.Now().Unix())),
			},
		},
	}
	err = s.postAlerts(ctx, params)
	if err != nil {
		return "", err
	}
	return captchaId, nil
}

// ForgetPwdVerifyEmail 忘记密码 邮件验证身份
func (s *AuthService) ForgetPwdVerifyEmail(ctx *gin.Context, req *oas.ForgetPwdVerifyEmailRequest) (*oas.ForgetPwdBeginResponse, error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token, s.Options)
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
	stateToken := createStateToken(sid, s.Options)
	err = s.Cache.Set(ctx, forgetPwdVerifyCachePrefix+sid, uid, cache.WithTTL(s.Options.StateTokenTTL))
	if err != nil {
		return nil, err
	}
	s.Cache.Del(ctx, cacheKey)
	return &oas.ForgetPwdBeginResponse{
		StateToken:    stateToken,
		StateTokenTTL: s.Options.StateTokenTTL.Seconds(),
	}, nil
}

// ForgetPwdVerifyMfa 忘记密码 mfa验证身份
func (s *AuthService) ForgetPwdVerifyMfa(ctx *gin.Context, req *oas.ForgetPwdVerifyMfaRequest) (*oas.ForgetPwdBeginResponse, error) {
	token := req.Body.StateToken
	id, err := parseStateToken(token, s.Options)
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
	stateToken := createStateToken(sid, s.Options)
	err = s.Cache.Set(ctx, forgetPwdVerifyCachePrefix+sid, profile.UserID, cache.WithTTL(s.Options.StateTokenTTL))
	if err != nil {
		return nil, err
	}
	s.Cache.Del(ctx, cacheKey)
	return &oas.ForgetPwdBeginResponse{
		StateToken:    stateToken,
		StateTokenTTL: s.Options.StateTokenTTL.Seconds(),
	}, nil
}

func (s *AuthService) tryGetTenantID(c *gin.Context) (tid int, err error) {
	if str := c.GetHeader("X-Tenant-ID"); str != "" {
		if tid, err = strconv.Atoi(str); err != nil {
			return 0, err
		}
	}
	return
}

// verifyTenantID 验证登录用户是否加入tid
func (s *AuthService) verifyTenantID(c *gin.Context, tid int) error {
	uid, err := identity.UserIDFromContext(c)
	if err != nil {
		return err
	}
	has, err := s.DB.OrgUser.Query().Where(orguser.UserID(uid), orguser.OrgID(tid)).Exist(c)
	if !has {
		return fmt.Errorf("invaild tenantID")
	}
	if err != nil {
		return err
	}
	return nil
}

// CreateSpm 创建spm key
func (s *AuthService) CreateSpm(ctx *gin.Context) (string, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return "", err
	}
	tid, err := s.tryGetTenantID(ctx)
	if err != nil {
		return "", err
	}

	err = s.verifyTenantID(ctx, tid)
	if err != nil {
		return "", err
	}

	spm := fmt.Sprintf("%d-%d-%d", tid, uid, time.Now().Unix())
	key := resource.SHA256(spm)
	err = s.Cache.Set(ctx, spmKeyPrefix+key, uid, cache.WithTTL(s.Options.SpmTTL))
	if err != nil {
		return "", err
	}
	return key, nil
}

// GetSpmAuth 根据spm 获取登录信息
func (s *AuthService) GetSpmAuth(c *gin.Context, r *oas.GetSpmAuthRequest) (*oas.LoginResponse, error) {
	var uid int
	err := s.Cache.Get(c, spmKeyPrefix+r.Spm, &uid)
	if err != nil {
		return nil, err
	}
	err = s.Cache.Del(c, spmKeyPrefix+r.Spm)
	if err != nil {
		return nil, err
	}
	if uid == 0 {
		return nil, fmt.Errorf("invaild spm")
	}
	return s.loginToken(c, uid)
}

// Token oauth获取accessToken
func (s *AuthService) Token(c *gin.Context, r *oas.TokenRequest) (*oas.TokenResponse, error) {
	oc, err := s.DB.OauthClient.Query().Where(
		oauthclient.GrantTypesEQ(oauthclient.GrantTypes(r.Body.GrantType)),
		oauthclient.ClientID(r.Body.ClientID),
		oauthclient.ClientSecret(r.Body.ClientSecret),
		oauthclient.StatusEQ(typex.SimpleStatusActive),
	).Only(c)
	if err != nil {
		return nil, fmt.Errorf("the clientID or clientSecret is incorrect or the status is not active")
	}

	tid, tstr, err := createToken(strconv.Itoa(oc.UserID), s.Options, false)
	if err != nil {
		return nil, err
	}
	err = s.Cache.Set(c, tid, oc.UserID, cache.WithTTL(s.Options.JWT.TokenTTL))
	if err != nil {
		return nil, err
	}
	// 更新认证时间
	err = s.DB.OauthClient.Update().Where(oauthclient.ID(oc.ID)).
		SetLastAuthAt(time.Now()).SetUpdatedBy(oc.UpdatedBy).Exec(c)
	if err != nil {
		return nil, err
	}
	return &oas.TokenResponse{
		AccessToken: tstr,
		ExpiresIn:   int(s.Options.JWT.TokenTTL.Seconds()),
	}, nil
}

type postAlertsInput struct {
	Annotations  map[string]string `json:"annotations,omitempty"`
	StartsAt     time.Time         `json:"startsAt,omitempty"`
	EndsAt       time.Time         `json:"endsAt,omitempty"`
	Labels       map[string]string `json:"labels"`
	GeneratorURL string            `json:"generatorURL,omitempty"`
}

func (s *AuthService) postAlerts(ctx context.Context, params []postAlertsInput) error {
	val, err := json.Marshal(params)
	if err != nil {
		return err
	}
	body := strings.NewReader(string(val))
	req, err := http.NewRequest("POST", s.OASOptions.Msgsrv.BaseUrl+"/api/v2/alerts", body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := s.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf(resp.Status)
}
