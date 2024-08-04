package auth

import (
	"bytes"
	"context"
	"encoding/base32"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/gds"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler"
	"github.com/woocoos/entcache"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/api/fs"
	"github.com/woocoos/knockout-go/api/fs/alioss"
	"github.com/woocoos/knockout-go/api/msg"
	"github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/fileidentity"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/oauthclient"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
	"github.com/woocoos/knockout/internal/status"
	"github.com/woocoos/knockout/service/resource"
	"image/png"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	mfaCachePrefix             = "mfa:"
	tokenCachePrefix           = "token:"
	resetCachePrefix           = "reset:"
	loginFailCachePrefix       = "loginfail:"
	forgetPwdBeginCachePrefix  = "forgetpwdbegin:"
	forgetPwdEmailCachePrefix  = "forgetpwdemail:"
	forgetPwdVerifyCachePrefix = "forgetpwdverify:"

	spmKeyPrefix = "spm:"

	callBackUrlResetPassword = "/login/reset-password"
	callBackUrlMFA           = "/login/verify-factor"
	callBackUrlCaptcha       = "/captcha"

	captchaWidth  = 200
	captchaHeight = 100
)

// Options is the configuration of AuthServer in the `auth` section.
type Options struct {
	// the path key of cache config, default `redis`
	CacheDriverName   string        `json:"cacheDriverName"`
	CaptchaCollectNum int           `json:"captchaCollectNum"` // captcha memory store collect num
	CaptchaExpire     time.Duration `json:"captchaExpire"`     // captcha expire time
	CaptchaLength     int           `json:"captchaLength"`     // captcha length
	CaptchaTimes      int           `json:"captchaTimes"`      // if login fail times, captcha will force show
	CaptchaTTL        time.Duration `json:"captchaTTL"`        // captcha ttl
	LoginFailTimes    int           `json:"loginFailTimes"`    // if login fail times, captcha will force show
	LoginFailLockTime time.Duration `json:"loginFailLockTime"` // lock time while login upper to max fail times
	StateTokenTTL     time.Duration `json:"stateTokenTTL"`     // state token ttl
	StateTokenSecret  string        `json:"stateTokenSecret"`  // state token secret
	SpmTTL            time.Duration `json:"spmTTL"`            // spm ttl
	JWT               struct {
		SigningMethod   string        `json:"signingMethod"`
		SigningKey      string        `json:"signingKey"`
		TokenTTL        time.Duration `json:"tokenTTL"`
		RefreshTokenTTL time.Duration `json:"refreshTokenTTL"`
	} `json:"jwt"`
}

// ServerImpl is the server API for service.
type ServerImpl struct {
	Options
	db    *ent.Client
	cache cache.Cache

	kosdk *api.SDK

	LogoutHandler func(*gin.Context)

	captchaStore captcha.Store

	webServer *web.Server
}

func NewServer(app *woocoo.App) *ServerImpl {
	var (
		err error
	)
	cnf := app.AppConfiguration()
	s := &ServerImpl{}
	ents := koapp.BuildEntComponents(cnf)
	if cnf.Development {
		s.db = ent.NewClient(ent.Driver(ents["portal"]), ent.Debug())
	} else {
		s.db = ent.NewClient(ent.Driver(ents["portal"]))
	}
	fs.RegisterS3Provider(fs.KindAliOSS, alioss.BuildProvider)
	if s.kosdk, err = api.NewSDK(cnf.Sub("kosdk")); err != nil {
		panic(err)
	}
	if s.cache, err = cache.GetCache("redis"); err != nil {
		panic(err)
	}
	if err = s.Apply(cnf); err != nil {
		panic(err)
	}

	s.buildWebServer(app)
	app.RegisterServer(s.webServer)
	return s
}

func (s *ServerImpl) buildWebServer(app *woocoo.App) *web.Server {
	s.webServer = web.New(web.WithConfiguration(app.AppConfiguration().Sub("web")),
		web.WithGracefulStop(),
		otelweb.RegisterMiddleware(),
	)
	// default group is '/'
	dr := s.webServer.Router().FindGroup("/").Group
	if mdl, ok := s.webServer.HandlerManager().GetMiddleware("jwt"); ok {
		s.LogoutHandler = mdl.(*handler.JWTMiddleware).Config.LogoutHandler
	}
	RegisterAuthHandlers(dr, s)
	RegisterHandlersManual(dr, s)
	return s.webServer
}

func (s *ServerImpl) Apply(cnf *conf.AppConfiguration) error {
	s.Options = Options{
		CacheDriverName:   "redis",
		CaptchaCollectNum: 1000,
		CaptchaExpire:     time.Minute * 2,
		CaptchaLength:     6,
		CaptchaTimes:      3,
		LoginFailTimes:    10,
		LoginFailLockTime: time.Hour * 24,
		SpmTTL:            time.Second * 5,
	}
	err := cnf.Sub("auth").Unmarshal(&s.Options)
	if err != nil {
		return err
	}

	// Initialize the captcha
	s.captchaStore = captcha.NewMemoryStore(s.CaptchaCollectNum, s.CaptchaExpire)
	captcha.SetCustomStore(s.captchaStore)
	return nil
}

// Start implements woocoo.Server but do noting in start, the web server has registered by NewServer.
func (s *ServerImpl) Start(ctx context.Context) error {
	return nil
}

func (s *ServerImpl) Stop(ctx context.Context) error {
	return s.db.Close()
}

func (s *ServerImpl) Captcha(ctx *gin.Context, req *CaptchaRequest) (*Captcha, error) {
	captchaId := captcha.NewLen(s.CaptchaLength)
	if req.W == nil {
		req.W = gds.Ptr(captchaWidth)
	}
	if req.H == nil {
		req.H = gds.Ptr(captchaHeight)
	}
	var buf bytes.Buffer
	err := captcha.WriteImage(&buf, captchaId, *req.W, *req.H)
	if err != nil {
		return nil, err
	}
	return &Captcha{
		CaptchaId:    captchaId,
		CaptchaImage: "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes()),
	}, err
}

// Login login
func (s *ServerImpl) Login(ctx *gin.Context, req *LoginRequest) (res *LoginResponse, err error) {
	failCount := 0
	s.cache.Get(ctx, loginFailCachePrefix+req.Username, &failCount)
	if failCount >= s.CaptchaTimes && !captcha.VerifyString(req.CaptchaId, req.Captcha) {
		ctx.Status(http.StatusBadRequest)
		s.logFailHandler(ctx, req.Username, false)
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
			failCount, errL = s.logFailHandler(ctx, req.Username, false)
			if errL != nil {
				return nil, errors.Join(err, errL)
			}
			if failCount >= s.CaptchaTimes {
				return &LoginResponse{CallbackUrl: callBackUrlCaptcha}, err
			}
			if failCount >= s.LoginFailTimes {
				return nil, status.ErrLoginFailUpperLimit
			}
		}
		return nil, err
	}

	profile, err := s.db.UserLoginProfile.Query().Where(userloginprofile.UserID(pwd.UserID)).Only(ctx)
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

	_ = updateLastLogin(ctx, s.db.UserLoginProfile, profile.UserID)
	return s.loginToken(ctx, pwd.UserID)
}

func (s *ServerImpl) RefreshToken(ctx *gin.Context, req *RefreshTokenRequest) (*LoginResponse, error) {
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
	err = s.cache.Set(ctx, tid, uid, cache.WithTTL(s.Options.JWT.TokenTTL))
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		AccessToken: tstr,
		ExpiresIn:   int(s.Options.JWT.TokenTTL.Seconds()),
	}, nil
}

func (s *ServerImpl) VerifyFactor(ctx *gin.Context, req *VerifyFactorRequest) (*LoginResponse, error) {
	token := req.StateToken
	id, err := parseStateToken(token, s.Options)
	if err != nil {
		return nil, err
	}

	var pid int
	if err = s.cache.Get(ctx, mfaCachePrefix+id, &pid); err != nil {
		return nil, err
	}

	profile, err := s.db.UserLoginProfile.Get(ctx, pid)
	if err != nil {
		return nil, err
	}
	if profile.MfaEnabled {
		if !totp.Validate(req.OtpToken, profile.MfaSecret) {
			return nil, errors.New("invalid code")
		}
	}

	if profile.PasswordReset {
		return s.resetPasswordPrepare(ctx, profile)
	}

	// no need use transaction
	err = updateLastLogin(ctx, s.db.UserLoginProfile, profile.UserID)
	return s.loginToken(ctx, profile.UserID)
}

func (s *ServerImpl) Logout(ctx *gin.Context) error {
	s.LogoutHandler(ctx)
	return nil
}

func (s *ServerImpl) ResetPassword(ctx *gin.Context, req *ResetPasswordRequest) (res *LoginResponse, err error) {
	token := req.StateToken
	id, err := parseStateToken(token, s.Options)
	if err != nil {
		return nil, err
	}

	var uid int
	cacheKey := resetCachePrefix + id
	if err = s.cache.Get(ctx, cacheKey, &uid); err != nil {
		return nil, err
	}

	pwd := s.db.UserPassword.Query().Where(userpassword.UserID(uid),
		userpassword.SceneEQ(userpassword.SceneLogin)).OnlyX(ctx)
	npwd := resource.SaltSecret(req.NewPassword, pwd.Salt)

	err = clientx.WithTx(ctx, func(ctx context.Context) (clientx.Transactor, error) {
		return s.db.Tx(ctx)
	}, func(itx clientx.Transactor) error {
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
		s.cache.Del(ctx, cacheKey) // lint:ignore
		return nil
	})
	return
}

func (s *ServerImpl) resetPasswordPrepare(ctx *gin.Context, profile *ent.UserLoginProfile) (res *LoginResponse, err error) {
	sid := uuid.New().String()
	res = &LoginResponse{
		CallbackUrl: callBackUrlResetPassword,
		StateToken:  createStateToken(sid, s.Options),
	}
	err = s.cache.Set(ctx, resetCachePrefix+sid, profile.UserID, cache.WithTTL(s.Options.StateTokenTTL))
	return
}

func (s *ServerImpl) mfaPrepare(ctx *gin.Context, profile *ent.UserLoginProfile) (res *LoginResponse, err error) {
	if !profile.MfaEnabled {
		return nil, nil
	}
	if profile.MfaEnabled && profile.MfaStatus != typex.SimpleStatusActive {
		return nil, errors.New("mfa not active")
	}
	sid := uuid.New().String()
	res = &LoginResponse{
		CallbackUrl: callBackUrlMFA,
		StateToken:  createStateToken(sid, s.Options),
	}
	err = s.cache.Set(ctx, mfaCachePrefix+sid, profile.ID, cache.WithTTL(s.StateTokenTTL))
	return
}

func updateLastLogin(ctx *gin.Context, pc *ent.UserLoginProfileClient, uid int) error {
	cip := ctx.ClientIP()
	// no mater what, update last login time and ip
	return pc.Update().Where(userloginprofile.UserID(uid)).
		SetLastLoginIP(cip).SetUpdatedBy(uid).SetPasswordReset(false).
		SetLastLoginAt(time.Now()).Exec(ctx)
}

func (s *ServerImpl) loginToken(ctx *gin.Context, uid int) (*LoginResponse, error) {
	usr := s.db.User.GetX(ctx, uid)

	tid, tstr, err := createToken(strconv.Itoa(uid), s.Options, false)
	if err != nil {
		return nil, err
	}

	_, trstr, err := createToken(strconv.Itoa(uid), s.Options, true)
	if err != nil {
		return nil, err
	}

	err = s.cache.Set(ctx, tid, uid, cache.WithTTL(s.Options.JWT.TokenTTL))
	if err != nil {
		return nil, err
	}

	var domains []*Domain
	err = s.db.Org.Query().Where(
		org.HasOrgUserWith(orguser.UserID(uid)),
		org.StatusEQ(typex.SimpleStatusActive),
		org.DomainNotNil(),
		org.KindEQ(org.KindRoot),
	).Select(org.FieldID, org.FieldName).Scan(ctx, &domains)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		AccessToken:  tstr,
		ExpiresIn:    int(s.Options.JWT.TokenTTL.Seconds()),
		RefreshToken: trstr,
		User: &User{
			ID:           usr.ID,
			DisplayName:  usr.DisplayName,
			AvatarFileId: usr.AvatarFileID,
			Domains:      domains,
		},
	}, nil
}

func (s *ServerImpl) checkPwd(ctx *gin.Context, req *LoginRequest) (*ent.UserPassword, error) {
	pwd, err := s.db.UserPassword.Query().Where(
		userpassword.HasUserWith(user.HasIdentitiesWith(useridentity.Code(req.Username))),
		userpassword.SceneEQ(userpassword.SceneLogin), userpassword.StatusEQ(typex.SimpleStatusActive),
	).Select(userpassword.FieldUserID, userpassword.FieldSalt, userpassword.FieldPassword).Only(entcache.Skip(ctx))
	if err != nil {
		return nil, err
	}

	given := resource.SaltSecret(req.Password, pwd.Salt)
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
func (s *ServerImpl) MfaQRCode(ctx *gin.Context, userID int, secret string) ([]byte, error) {
	uorg, err := s.GetUserRootOrg(ctx, userID)
	if err != nil {
		return nil, err
	}

	issuer := uorg.Domain
	profile, err := s.db.UserLoginProfile.Query().WithUser(func(query *ent.UserQuery) {
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

func (s *ServerImpl) BindMfaPrepare(ctx *gin.Context) (*Mfa, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	pn, err := s.db.User.Query().Where(user.ID(uid)).Select(user.FieldPrincipalName).String(ctx)
	if err != nil {
		return nil, err
	}
	sid := uuid.New().String()
	stateToken := createStateToken(sid, s.Options)
	val := map[string]string{
		"uid":    strconv.Itoa(uid),
		"secret": resource.GeneralMFASecret(),
	}
	err = s.cache.Set(ctx, mfaCachePrefix+sid, val, cache.WithTTL(s.Options.StateTokenTTL))
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
	uorg, err := s.db.Org.Query().Where(org.ID(tid)).Only(ctx)
	issuer := uorg.Domain
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: pn,
		Secret:      secByte,
	})
	if err != nil {
		return nil, err
	}
	return &Mfa{
		QrCodeUri:     key.String(),
		PrincipalName: pn,
		Secret:        val["secret"],
		StateToken:    stateToken,
		StateTokenTTL: s.Options.StateTokenTTL.Seconds(),
	}, nil
}

func (s *ServerImpl) BindMfa(ctx *gin.Context, req *BindMfaRequest) (bool, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	id, err := parseStateToken(req.StateToken, s.Options)
	if err != nil {
		return false, err
	}
	//
	var val map[string]string
	if err = s.cache.Get(ctx, mfaCachePrefix+id, &val); err != nil {
		return false, err
	}
	if val["uid"] != strconv.Itoa(uid) {
		return false, fmt.Errorf("invalid user")
	}
	if !totp.Validate(req.OtpToken, val["secret"]) {
		return false, errors.New("invalid code")
	}
	err = s.db.UserLoginProfile.UpdateOneID(uid).SetMfaEnabled(true).SetMfaStatus(typex.SimpleStatusActive).SetMfaSecret(val["secret"]).Exec(ctx)
	return err == nil, err
}

func (s *ServerImpl) UnBindMfa(ctx *gin.Context, req *UnBindMfaRequest) (bool, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	up, err := s.db.UserLoginProfile.Query().Where(userloginprofile.UserID(uid)).Only(ctx)
	if !totp.Validate(req.OtpToken, up.MfaSecret) {
		return false, errors.New("invalid code")
	}
	err = s.db.UserLoginProfile.UpdateOneID(up.ID).ClearMfaEnabled().ClearMfaStatus().ClearMfaSecret().Exec(ctx)
	return err == nil, err
}

func (s *ServerImpl) GetUserRootOrg(ctx *gin.Context, uid int) (uorg *ent.Org, err error) {
	uorg, err = s.db.OrgUser.Query().Where(orguser.UserIDEQ(uid)).
		QueryOrg().Unique(false).Where(org.KindEQ(org.KindRoot), org.StatusEQ(typex.SimpleStatusActive)).Order(ent.Desc(org.FieldPath)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return uorg, nil
}

func (s *ServerImpl) logFailHandler(ctx *gin.Context, uid string, clear bool) (int, error) {
	key := loginFailCachePrefix + uid
	if clear {
		return 0, s.cache.Del(ctx, key)
	}
	count := 0
	err := s.cache.Get(ctx, key, &count)
	if err != nil && !s.cache.IsNotFound(err) {
		return 0, err
	}
	count++
	err = s.cache.Set(ctx, key, count, cache.WithTTL(s.LoginFailLockTime)) // 以账户锁定时间作为过期时间
	return count, err
}

// ForgetPwdBegin 忘记密码验证用户账户，开始修改密码流程
func (s *ServerImpl) ForgetPwdBegin(ctx *gin.Context, req *ForgetPwdBeginRequest) (*ForgetPwdBeginResponse, error) {
	// 验证验证码
	if !captcha.VerifyString(req.CaptchaId, req.Captcha) {
		return nil, status.ErrCaptchaNotMatch
	}
	// 查询用户
	u, err := s.db.User.Query().Where(user.HasIdentitiesWith(useridentity.Code(req.Username))).WithLoginProfile().Only(ctx)
	if err != nil {
		return nil, err
	}
	verifies := make([]*ForgetPwdVerify, 0)
	if u.Edges.LoginProfile.MfaEnabled {
		verifies = append(verifies, &ForgetPwdVerify{Kind: "mfa"})
	}
	if &u.Email != nil {
		verifies = append(verifies, &ForgetPwdVerify{Kind: "email", Value: resource.MaskEmail(u.Email)})
	}
	// 生成临时token
	sid := uuid.New().String()
	stateToken := createStateToken(sid, s.Options)
	err = s.cache.Set(ctx, forgetPwdBeginCachePrefix+sid, u.ID, cache.WithTTL(s.Options.StateTokenTTL))
	if err != nil {
		return nil, err
	}
	return &ForgetPwdBeginResponse{
		StateToken:    stateToken,
		StateTokenTTL: s.Options.StateTokenTTL.Seconds(),
		Verifies:      verifies,
	}, nil
}

// ForgetPwdReset 忘记密码设置新密码
func (s *ServerImpl) ForgetPwdReset(ctx *gin.Context, req *ForgetPwdResetRequest) (bool, error) {
	token := req.StateToken
	id, err := parseStateToken(token, s.Options)
	if err != nil {
		return false, err
	}
	var uid int
	cacheKey := forgetPwdVerifyCachePrefix + id
	if err = s.cache.Get(ctx, cacheKey, &uid); err != nil {
		return false, err
	}
	//
	pwd := s.db.UserPassword.Query().Where(userpassword.UserID(uid), userpassword.SceneEQ(userpassword.SceneLogin)).OnlyX(ctx)
	npwd := resource.SaltSecret(req.NewPassword, pwd.Salt)

	err = clientx.WithTx(ctx, func(ctx context.Context) (clientx.Transactor, error) {
		return s.db.Tx(ctx)
	}, func(itx clientx.Transactor) error {
		tx := itx.(*ent.Tx)
		err = tx.UserPassword.UpdateOneID(pwd.ID).SetUpdatedBy(uid).SetPassword(npwd).Exec(ctx)
		if err != nil {
			return err
		}
		s.cache.Del(ctx, cacheKey) // lint:ignore
		return nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// ForgetPwdSendEmail 忘记密码 发送邮件验证码
func (s *ServerImpl) ForgetPwdSendEmail(ctx *gin.Context, req *ForgetPwdSendEmailRequest) (string, error) {
	token := req.StateToken
	id, err := parseStateToken(token, s.Options)
	if err != nil {
		return "", err
	}
	var uid int
	cacheKey := forgetPwdBeginCachePrefix + id
	if err = s.cache.Get(ctx, cacheKey, &uid); err != nil {
		return "", err
	}
	// 生成验证码
	captchaId := captcha.NewLen(6)
	digits := s.captchaStore.Get(captchaId, false)
	captchaCode := ""
	for _, v := range digits {
		captchaCode = captchaCode + strconv.Itoa(int(v))
	}
	usr, err := s.db.User.Get(ctx, uid)
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

	params := msg.PostableAlerts{
		{
			Annotations: map[string]string{
				"to":            usr.Email,
				"displayName":   usr.DisplayName,
				"captchaCode":   captchaCode,
				"captchaExpire": strconv.Itoa(int(s.CaptchaExpire.Minutes())),
			},
			Alert: &msg.Alert{
				Labels: map[string]string{
					"receiver":  "email",
					"alertname": "SendCaptchaCode",
					"tenant":    strconv.Itoa(uorg.ID),
					"timestamp": strconv.Itoa(int(time.Now().Unix())),
				},
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
func (s *ServerImpl) ForgetPwdVerifyEmail(ctx *gin.Context, req *ForgetPwdVerifyEmailRequest) (*ForgetPwdBeginResponse, error) {
	token := req.StateToken
	id, err := parseStateToken(token, s.Options)
	if err != nil {
		return nil, err
	}
	var uid int
	cacheKey := forgetPwdBeginCachePrefix + id
	if err = s.cache.Get(ctx, cacheKey, &uid); err != nil {
		return nil, err
	}
	// 验证验证码
	if !captcha.VerifyString(req.CaptchaId, req.Captcha) {
		return nil, status.ErrCaptchaNotMatch
	}
	sid := uuid.New().String()
	stateToken := createStateToken(sid, s.Options)
	err = s.cache.Set(ctx, forgetPwdVerifyCachePrefix+sid, uid, cache.WithTTL(s.Options.StateTokenTTL))
	if err != nil {
		return nil, err
	}
	s.cache.Del(ctx, cacheKey)
	return &ForgetPwdBeginResponse{
		StateToken:    stateToken,
		StateTokenTTL: s.Options.StateTokenTTL.Seconds(),
	}, nil
}

// ForgetPwdVerifyMfa 忘记密码 mfa验证身份
func (s *ServerImpl) ForgetPwdVerifyMfa(ctx *gin.Context, req *ForgetPwdVerifyMfaRequest) (*ForgetPwdBeginResponse, error) {
	token := req.StateToken
	id, err := parseStateToken(token, s.Options)
	if err != nil {
		return nil, err
	}
	var uid int
	cacheKey := forgetPwdBeginCachePrefix + id
	if err = s.cache.Get(ctx, cacheKey, &uid); err != nil {
		return nil, err
	}
	profile, err := s.db.UserLoginProfile.Query().Where(userloginprofile.UserID(uid)).Only(ctx)
	if err != nil {
		return nil, err
	}
	// 验证mfa
	if profile.MfaEnabled {
		if !totp.Validate(req.OtpToken, profile.MfaSecret) {
			return nil, errors.New("invalid code")
		}
	} else {
		return nil, fmt.Errorf("the MFA is disabled")
	}
	// 生成临时token
	sid := uuid.New().String()
	stateToken := createStateToken(sid, s.Options)
	err = s.cache.Set(ctx, forgetPwdVerifyCachePrefix+sid, profile.UserID, cache.WithTTL(s.Options.StateTokenTTL))
	if err != nil {
		return nil, err
	}
	s.cache.Del(ctx, cacheKey)
	return &ForgetPwdBeginResponse{
		StateToken:    stateToken,
		StateTokenTTL: s.Options.StateTokenTTL.Seconds(),
	}, nil
}

func (s *ServerImpl) tryGetTenantID(c *gin.Context) (tid int, err error) {
	if str := c.GetHeader("X-Tenant-ID"); str != "" {
		if tid, err = strconv.Atoi(str); err != nil {
			return 0, err
		}
	}
	return
}

// verifyTenantID 验证登录用户是否加入tid
func (s *ServerImpl) verifyTenantID(c *gin.Context, tid int) error {
	uid, err := identity.UserIDFromContext(c)
	if err != nil {
		return err
	}
	has, err := s.db.OrgUser.Query().Where(orguser.UserID(uid), orguser.OrgID(tid)).Exist(c)
	if !has {
		return fmt.Errorf("invaild tenantID")
	}
	if err != nil {
		return err
	}
	return nil
}

// CreateSpm 创建spm key
func (s *ServerImpl) CreateSpm(ctx *gin.Context) (string, error) {
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
	err = s.cache.Set(ctx, spmKeyPrefix+key, uid, cache.WithTTL(s.Options.SpmTTL))
	if err != nil {
		return "", err
	}
	return key, nil
}

// GetSpmAuth 根据spm 获取登录信息
func (s *ServerImpl) GetSpmAuth(c *gin.Context, r *GetSpmAuthRequest) (*LoginResponse, error) {
	var uid int
	err := s.cache.Get(c, spmKeyPrefix+r.Spm, &uid)
	if err != nil {
		return nil, err
	}
	err = s.cache.Del(c, spmKeyPrefix+r.Spm)
	if err != nil {
		return nil, err
	}
	if uid == 0 {
		return nil, fmt.Errorf("invaild spm")
	}
	return s.loginToken(c, uid)
}

// Token oauth获取accessToken
func (s *ServerImpl) Token(c *gin.Context, r *TokenRequest) (*TokenResponse, error) {
	oc, err := s.db.OauthClient.Query().Where(
		oauthclient.GrantTypesEQ(oauthclient.GrantTypes(r.GrantType)),
		oauthclient.ClientID(r.ClientID),
		oauthclient.ClientSecret(r.ClientSecret),
		oauthclient.StatusEQ(typex.SimpleStatusActive),
	).Only(c)
	if err != nil {
		return nil, fmt.Errorf("the clientID or clientSecret is incorrect or the status is not active")
	}

	tid, tstr, err := createToken(strconv.Itoa(oc.UserID), s.Options, false)
	if err != nil {
		return nil, err
	}
	err = s.cache.Set(c, tid, oc.UserID, cache.WithTTL(s.Options.JWT.TokenTTL))
	if err != nil {
		return nil, err
	}
	// 更新认证时间
	err = s.db.OauthClient.Update().Where(oauthclient.ID(oc.ID)).
		SetLastAuthAt(time.Now()).SetUpdatedBy(oc.UpdatedBy).Exec(c)
	if err != nil {
		return nil, err
	}
	return &TokenResponse{
		AccessToken: tstr,
		ExpiresIn:   int(s.Options.JWT.TokenTTL.Seconds()),
	}, nil
}

func (s *ServerImpl) postAlerts(ctx context.Context, params msg.PostableAlerts) error {
	_, err := s.kosdk.Msg().AlertAPI.PostAlerts(ctx, &msg.PostAlertsRequest{
		PostableAlerts: params,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerImpl) GetSTS(ctx *gin.Context, req *GetSTSRequest) (*GetSTSResponse, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	tid, err := s.tryGetTenantID(ctx)
	if err != nil {
		return nil, err
	}
	var fi *ent.FileIdentity
	if req.Bucket != "" && req.Endpoint != "" {
		// 传参取对应identity
		fi, err = s.db.FileIdentity.Query().Where(
			fileidentity.TenantID(tid),
			fileidentity.HasSourceWith(
				filesource.Endpoint(req.Endpoint),
				filesource.Bucket(req.Bucket),
			),
		).WithSource().Only(ctx)
	} else {
		// 不传参去默认值
		fi, err = s.db.FileIdentity.Query().Where(
			fileidentity.TenantID(tid),
			fileidentity.IsDefault(true),
		).WithSource().Only(ctx)
	}
	if fi == nil {
		return nil, fmt.Errorf("the fileidentity is null")
	}
	if err != nil {
		return nil, err
	}

	provider, err := s.kosdk.Fs().GetProvider(ctx, s.toOSSFileSource(fi))
	if err != nil {
		return nil, err
	}
	usr, err := s.db.User.Get(ctx, uid)
	if err != nil {
		return nil, err
	}
	resp, err := provider.GetSTS(usr.PrincipalName)
	if err != nil {
		return nil, err
	}
	return &GetSTSResponse{
		AccessKeyID:     resp.AccessKeyID,
		SecretAccessKey: resp.SecretAccessKey,
		SessionToken:    resp.SessionToken,
		Expiration:      resp.Expiration,
	}, nil
}

func (s *ServerImpl) GetPreSignUrl(ctx *gin.Context, req *GetPreSignUrlRequest) (*GetPreSignUrlResponse, error) {
	fi, path, err := s.convertUrlToFileSource(ctx, req)
	if err != nil {
		return nil, err
	}
	provider, err := s.kosdk.Fs().GetProvider(ctx, s.toOSSFileSource(fi))
	if err != nil {
		return nil, err
	}
	signUrl, err := provider.GetPreSignedURL(fi.Edges.Source.Bucket, path, time.Hour)
	if err != nil {
		return nil, err
	}
	return &GetPreSignUrlResponse{
		URL: signUrl,
	}, nil
}

// convertUrlToFileSource 将url转换为文件源
func (s *ServerImpl) convertUrlToFileSource(ctx *gin.Context, req *GetPreSignUrlRequest) (*ent.FileIdentity, string, error) {
	tid, err := s.tryGetTenantID(ctx)
	if err != nil {
		return nil, "", err
	}
	var fis []*ent.FileIdentity
	// 取出组织对应的fileidentities
	if req.Bucket != "" && req.Endpoint != "" {
		fis, err = s.db.FileIdentity.Query().Where(
			fileidentity.TenantID(tid),
			fileidentity.HasSourceWith(
				filesource.Bucket(req.Bucket),
				filesource.Endpoint(req.Endpoint),
			),
		).WithSource().All(ctx)
	} else {
		fis, err = s.db.FileIdentity.Query().Where(fileidentity.TenantID(tid), fileidentity.IsDefault(true)).WithSource().All(ctx)
	}
	if err != nil {
		return nil, "", err
	}
	var fi *ent.FileIdentity
	// 根据bucketUrl获取对应的fileIdentity
	for _, f := range fis {
		if strings.HasPrefix(req.URL, f.Edges.Source.BucketURL) {
			fi = f
			break
		}
	}
	if fi == nil {
		return nil, "", fmt.Errorf("the fileidentity is null")
	}

	// 解析url的path
	path := ""
	u, err := url.Parse(req.URL)
	if fi.Edges.Source.Kind == filesource.KindMinio {
		path = strings.TrimLeft(u.Path, "/"+fi.Edges.Source.Bucket)
	} else if fi.Edges.Source.Kind == filesource.KindAliOSS {
		path = u.Path
	} else {
		return nil, "", fmt.Errorf("error path")
	}
	return fi, path, nil
}

func (s *ServerImpl) toOSSFileSource(fi *ent.FileIdentity) *fs.SourceConfig {
	return &fs.SourceConfig{
		TenantID:          fi.TenantID,
		Kind:              fs.Kind(fi.Edges.Source.Kind.String()),
		Bucket:            fi.Edges.Source.Bucket,
		BucketUrl:         fi.Edges.Source.BucketURL,
		Endpoint:          fi.Edges.Source.Endpoint,
		EndpointImmutable: fi.Edges.Source.EndpointImmutable,
		StsEndpoint:       fi.Edges.Source.StsEndpoint,
		AccessKeyID:       fi.AccessKeyID,
		AccessKeySecret:   fi.AccessKeySecret,
		Policy:            fi.Policy,
		Region:            fi.Edges.Source.Region,
		RoleArn:           fi.RoleArn,
		DurationSeconds:   fi.DurationSeconds,
	}
}
