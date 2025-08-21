package auth

import (
	"bytes"
	"context"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/tsingsun/woocoo/pkg/auth"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/gds"
	securityX "github.com/tsingsun/woocoo/pkg/security"
	"github.com/tsingsun/woocoo/pkg/store/redisx"
	"github.com/woocoos/entcache"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/api/fs"
	_ "github.com/woocoos/knockout-go/api/fs/alioss"
	"github.com/woocoos/knockout-go/api/msg"
	"github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/authz"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/fileidentity"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/oauthclient"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/quota"
	"github.com/woocoos/knockout/ent/quotaitem"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/useraddr"
	"github.com/woocoos/knockout/ent/userdevice"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
	"github.com/woocoos/knockout/ent/userpasswordpolicy"
	"github.com/woocoos/knockout/internal/errors"
	"github.com/woocoos/knockout/security"
	quotaService "github.com/woocoos/knockout/service/quota"
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
	verifyDeviceCachePrefix    = "verifyDevice:"
	isLogonCachePrefix         = "isLogon:"

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
	CacheDriverName    string        `json:"cacheDriverName"`
	CaptchaCollectNum  int           `json:"captchaCollectNum"` // captcha memory store collect num
	CaptchaExpire      time.Duration `json:"captchaExpire"`     // captcha expire time
	CaptchaLength      int           `json:"captchaLength"`     // captcha length
	CaptchaTimes       int           `json:"captchaTimes"`      // if login fail times, captcha will force show
	CaptchaTTL         time.Duration `json:"captchaTTL"`        // captcha ttl
	LoginFailTimes     int           `json:"loginFailTimes"`    // if login fail times, captcha will force show
	LoginFailLockTime  time.Duration `json:"loginFailLockTime"` // lock time while login upper to max fail times
	StateTokenTTL      time.Duration `json:"stateTokenTTL"`     // state token ttl
	StateTokenSecret   string        `json:"stateTokenSecret"`  // state token secret
	SpmTTL             time.Duration `json:"spmTTL"`            // spm ttl
	VerifyDeviceParams struct {
		DefaultBound    bool     `json:"defaultBound"`    // The device is bound by default when logging in for the first time
		ExcludeAccounts []string `json:"excludeAccounts"` // exclude accounts
	} `json:"verifyDeviceParams"`
	JWT struct {
		SigningMethod   string        `json:"signingMethod"`
		SigningKey      string        `json:"signingKey"`
		PrivateKey      string        `json:"privateKey"`
		TokenTTL        time.Duration `json:"tokenTTL"`
		RefreshTokenTTL time.Duration `json:"refreshTokenTTL"`
	} `json:"jwt"`
	PwdPolicy OptionsPwdPolicy `json:"pwdPolicy"`
}
type OptionsPwdPolicy struct {
	// 密码最短长度，长度应在6-32位之间
	Length int32 `json:"length"`
	// 必须包含的元素，异或：1-小写字母，2-大写字母，4-数字，8-符号
	IncludeElement int32 `json:"includeElement"`
	// 最少包含的不同字符数，最多8个，0代表不限制
	IncludeChar int32 `json:"includeChar"`
	// 是否允许包含用户名
	AllowIncludeUserName bool `json:"allowIncludeUserName"`
	// 有效天数，最大1095天，0代表不过期
	InvalidDay int32 `json:"invalidDay"`
	// 过期后是否限制登录
	InvalidLoginLimit bool `json:"invalidLoginLimit"`
	// 一小时内密码错误最多尝试次数，最大32次，0代表不限次数
	Retry int32 `json:"retry"`
	// 密码错误多少次出现验证码，最大5次，0代表不出现验证码
	CaptchaTimes int32 `json:"captchaTimes"`
}

// ServerImpl is the server API for service.
type ServerImpl struct {
	Options
	db          *ent.Client
	redisClient *redisx.Client

	cache cache.Cache

	kosdk *api.SDK

	LogoutHandler func(*gin.Context)

	captchaStore captcha.Store
}

func NewServerImpl(cnf *conf.AppConfiguration) *ServerImpl {
	var (
		err error
	)
	s := &ServerImpl{}
	if s.kosdk, err = api.NewSDK(cnf.Sub("kosdk")); err != nil {
		panic(err)
	}
	if s.cache, err = cache.GetCache("redis"); err != nil {
		panic(err)
	}
	if err = s.Apply(cnf); err != nil {
		panic(err)
	}

	return s
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
		VerifyDeviceParams: struct {
			DefaultBound    bool     `json:"defaultBound"`
			ExcludeAccounts []string `json:"excludeAccounts"`
		}{DefaultBound: false, ExcludeAccounts: []string{}},
		PwdPolicy: OptionsPwdPolicy{
			Length:               6,
			IncludeElement:       3,
			IncludeChar:          4,
			AllowIncludeUserName: false,
			InvalidDay:           30,
			InvalidLoginLimit:    false,
			Retry:                5,
			CaptchaTimes:         3,
		},
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
	var failCount int32
	s.cache.Get(ctx, loginFailCachePrefix+req.Username, &failCount)
	// 获取密码策略
	upp, err := s.getPasswordPolicy(ctx)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, err
	}
	// 判断用户是否锁定状态
	ui, err := s.db.UserIdentity.Query().Where(
		useridentity.Code(req.Username),
		useridentity.StatusEQ(typex.SimpleStatusActive),
	).WithUser().Only(ctx)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, errors.Codel(errors.ErrUserNameOrPassword)
	}
	if ui.Edges.User.Status == types.UserStatusLocked {
		ctx.Status(http.StatusBadRequest)
		// 返回账号锁定错误
		return nil, errors.Codel(errors.ErrUserHasLocked)
	}
	// 判断密码是否过期
	has, err := s.db.UserPassword.Query().Where(
		userpassword.HasUserWith(user.HasIdentitiesWith(useridentity.Code(req.Username))),
		userpassword.SceneEQ(userpassword.SceneLogin), userpassword.StatusEQ(typex.SimpleStatusDisabled),
	).Exist(entcache.Skip(ctx))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return nil, err
	}
	if has {
		ctx.Status(http.StatusBadRequest)
		// 密码已过期，请重置密码或联系客服修改密码恢复
		return nil, errors.Codel(errors.ErrPasswordExpired)
	}
	// 错误次数过多需要验证码
	if upp.CaptchaTimes > 0 && failCount >= upp.CaptchaTimes {
		// 若错误次数大于验证码应该出现的次数，则前端展示验证码
		if req.CaptchaId == "" || req.Captcha == "" {
			ctx.Status(http.StatusAccepted)
			return &LoginResponse{CallbackUrl: callBackUrlCaptcha}, nil
		}
		if !captcha.VerifyString(req.CaptchaId, req.Captcha) {
			ctx.Status(http.StatusBadRequest)
			return nil, errors.Codel(errors.ErrCaptchaNotMatch)
		}
	}
	pwd, err := s.checkPwd(ctx, req)
	if err != nil {
		// 处理密码错误
		return s.dealPwdError(ctx, req, ui.UserID, upp, err)
	}

	profile, err := s.db.UserLoginProfile.Query().Where(userloginprofile.UserID(pwd.UserID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	if !profile.CanLogin {
		return nil, errors.Codel(errors.ErrUserCanNotLogin)
	}

	// 清除登录失败次数缓存
	s.logFailHandler(ctx, req.Username, true)
	if profile.MfaEnabled {
		return s.mfaPrepare(ctx, profile)
	}

	if profile.PasswordReset {
		return s.resetPasswordPrepare(ctx, profile)
	}

	_ = s.updateLastLogin(ctx, profile.UserID)
	return s.loginToken(ctx, pwd.UserID)
}

func (s *ServerImpl) dealPwdError(ctx *gin.Context, req *LoginRequest, userID int, upp *ent.UserPasswordPolicy, err error) (*LoginResponse, error) {
	ctx.Status(http.StatusBadRequest)
	var parsedError *gin.Error
	has := errors.As(err, &parsedError)
	if has && err.(*gin.Error).Type == errors.ErrPasswordNotMatch {
		var errL error
		failCount, errL := s.logFailHandler(ctx, req.Username, false)
		if errL != nil {
			return nil, errors.Join(err, errL)
		}
		// 登录失败次数大于设定值，锁定用户
		if upp.Retry > 0 && failCount >= upp.Retry {
			// 锁定用户，更新用户状态
			err = s.db.User.UpdateOneID(userID).SetUpdatedBy(userID).SetStatus(types.UserStatusLocked).Exec(ctx)
			if err != nil {
				return nil, err
			}
			// 账户锁定，清除失败次数缓存
			_, _ = s.logFailHandler(ctx, req.Username, true)
			// 发送邮件给指定用户，指定的用户在邮件模板配置
			usr, addr, err := s.getUserInfo(ctx, userID)
			if err != nil {
				return nil, err
			}
			uorg, err := s.GetUserRootOrg(ctx, userID)
			if err != nil {
				return nil, err
			}
			tid, err := s.getTopOrgIdByPath(uorg.Path)
			if err != nil {
				return nil, err
			}
			params := msg.PostableAlerts{
				{
					Annotations: map[string]string{
						"to":            addr.Email,
						"displayName":   usr.DisplayName,
						"principalName": req.Username,
						"pwdRetry":      strconv.Itoa(int(upp.Retry)),
					},
					Alert: &msg.Alert{
						Labels: map[string]string{
							"receiver":  "email",
							"alertname": "UserLockedNotify",
							"tenant":    strconv.Itoa(tid),
							"timestamp": strconv.Itoa(int(time.Now().Unix())),
						},
					},
				},
			}
			_ = s.postAlerts(ctx, params)
			ctx.Status(http.StatusBadRequest)
			// 返回账号锁定错误
			return nil, errors.Codel(errors.ErrUserHasLocked)
		}
		return nil, errors.Codel(errors.ErrPasswordRetry, upp.Retry-failCount)
	}
	return nil, err
}

func (s *ServerImpl) AppOrgs(ctx *gin.Context, req *AppOrgsRequest) ([]*Domain, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	ros, err := s.db.Org.Query().Where(
		org.HasOrgUserWith(orguser.UserID(uid)),
		org.StatusEQ(typex.SimpleStatusActive),
		org.KindEQ(org.KindRoot),
	).Select(org.FieldID, org.FieldName, org.FieldPath, org.FieldLocalCurrency).Order(ent.Asc(org.FieldID)).All(ctx)
	if err != nil {
		return nil, err
	}
	if ros == nil || len(ros) == 0 {
		return nil, errors.Codel(errors.ErrOrgNotFound)
	}
	var domains []*Domain
	// 查询顶级组织
	for _, o := range ros {
		has, err := s.doCheckPermission(ctx, uid, o.ID, "login", req.AppCode)
		if !has || err != nil {
			continue
		}
		// 根据/截取path的第一项
		code := strings.Split(o.Path, "/")[0]
		// 转换成十进制id
		oID, err := strconv.ParseInt(code, 36, 64)
		if err != nil {
			return nil, err
		}
		to, err := s.db.Org.Query().Where(org.ID(int(oID))).Select(org.FieldID, org.FieldName, org.FieldPath, org.FieldLocalCurrency).Only(ctx)
		if err != nil {
			return nil, err
		}
		domains = append(domains, &Domain{
			ID:             o.ID,
			Name:           o.Name,
			LocalCurrency:  o.LocalCurrency,
			ParentID:       to.ID,
			ParentName:     to.Name,
			ParentCurrency: to.LocalCurrency,
		})
	}
	return domains, nil
}

func (s *ServerImpl) OldLoginForApp(ctx *gin.Context, req *OldLoginForAppRequest) (res *LoginResponse, err error) {
	// 验证密码
	pwd, err := s.checkPwd(ctx, &LoginRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, errors.Codel(errors.ErrUserNameOrPassword)
	}

	profile, err := s.db.UserLoginProfile.Query().Where(userloginprofile.UserID(pwd.UserID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	if !profile.CanLogin {
		return nil, errors.Codel(errors.ErrUserCanNotLogin)
	}

	roIDs, err := s.db.Org.Query().Where(
		org.HasOrgUserWith(orguser.UserID(pwd.UserID)),
		org.StatusEQ(typex.SimpleStatusActive),
		org.KindEQ(org.KindRoot),
	).Select(org.FieldID).Ints(ctx)
	if err != nil {
		return nil, err
	}
	if roIDs == nil || len(roIDs) == 0 {
		return nil, errors.Codel(errors.ErrOrgNotFound)
	}
	appAccess := false
	for _, roID := range roIDs {
		// 验证登录权限
		has, err := s.doCheckPermission(ctx, pwd.UserID, roID, "login", req.AppCode)
		if err != nil {
			return nil, err
		}
		if has {
			appAccess = true
			break
		}
	}
	if !appAccess {
		return nil, errors.Codel(errors.ErrNotLoginPermission)
	}

	if profile.MfaEnabled {
		if !totp.Validate(req.OtpToken, profile.MfaSecret) {
			return nil, errors.Codel(errors.ErrMfaInvalidCode)
		}
	} else {
		return nil, errors.Codel(errors.ErrMfaDisable)
	}

	cip := ctx.ClientIP()
	// no mater what, update last login time and ip
	err = s.db.UserLoginProfile.Update().Where(userloginprofile.UserID(profile.UserID)).
		SetLastLoginIP(cip).SetUpdatedBy(profile.UserID).SetLastLoginAt(time.Now()).Exec(ctx)
	return s.loginToken(ctx, pwd.UserID)
}

func (s *ServerImpl) RefreshToken(ctx *gin.Context, req *RefreshTokenRequest) (*LoginResponse, error) {
	token, err := jwt.ParseWithClaims(req.RefreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		token.Method = jwt.GetSigningMethod(s.Options.JWT.SigningMethod)
		key, err := auth.ParseSigningKeyFromString(s.Options.JWT.SigningKey, s.Options.JWT.SigningMethod, false)
		if err != nil {
			return nil, err
		}
		return key, nil
	})
	if err != nil || !token.Valid {
		ctx.Status(http.StatusUnauthorized)
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
	err = s.cache.Set(ctx, tid, strconv.Itoa(uid), cache.WithTTL(s.Options.JWT.TokenTTL))
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		AccessToken: tstr,
		ExpiresIn:   int(s.Options.JWT.TokenTTL.Seconds()),
	}, nil
}
func (s *ServerImpl) OldFingerprintLogin(ctx *gin.Context, req *OldFingerprintLoginRequest) (*LoginResponse, error) {
	// 验证密码
	pwd, err := s.checkPwd(ctx, &LoginRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, errors.Codel(errors.ErrUserNameOrPassword)
	}

	profile, err := s.db.UserLoginProfile.Query().Where(userloginprofile.UserID(pwd.UserID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	if !profile.CanLogin {
		return nil, errors.Codel(errors.ErrUserCanNotLogin)
	}

	roIDs, err := s.db.Org.Query().Where(
		org.HasOrgUserWith(orguser.UserID(pwd.UserID)),
		org.StatusEQ(typex.SimpleStatusActive),
		org.KindEQ(org.KindRoot),
	).Select(org.FieldID).Ints(ctx)
	if err != nil {
		return nil, err
	}
	if roIDs == nil || len(roIDs) == 0 {
		return nil, errors.Codel(errors.ErrOrgNotFound)
	}
	appAccess := false
	for _, roID := range roIDs {
		// 验证登录权限
		has, err := s.doCheckPermission(ctx, pwd.UserID, roID, "login", req.AppCode)
		if err != nil {
			return nil, err
		}
		if has {
			appAccess = true
			break
		}
	}
	if !appAccess {
		return nil, errors.Codel(errors.ErrNotLoginPermission)
	}

	cip := ctx.ClientIP()
	// no mater what, update last login time and ip
	err = s.db.UserLoginProfile.Update().Where(userloginprofile.UserID(profile.UserID)).
		SetLastLoginIP(cip).SetUpdatedBy(profile.UserID).SetLastLoginAt(time.Now()).Exec(ctx)
	return s.loginToken(ctx, pwd.UserID)
}
func (s *ServerImpl) FingerprintLogin(ctx *gin.Context, req *FingerprintLoginRequest) (*LoginResponse, error) {
	token, err := jwt.ParseWithClaims(req.RefreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		token.Method = jwt.GetSigningMethod(s.Options.JWT.SigningMethod)
		key, err := auth.ParseSigningKeyFromString(s.Options.JWT.SigningKey, s.Options.JWT.SigningMethod, false)
		if err != nil {
			return nil, err
		}
		return key, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	subject := token.Claims.(*jwt.RegisteredClaims).Subject
	uid, err := strconv.Atoi(subject)
	if err != nil {
		return nil, err
	}
	profile, err := s.db.UserLoginProfile.Query().Where(userloginprofile.UserID(uid)).Only(ctx)
	if err != nil {
		return nil, err
	}

	if !profile.CanLogin {
		return nil, errors.Codel(errors.ErrUserCanNotLogin)
	}
	_ = s.updateLastLogin(ctx, profile.UserID)
	return s.loginToken(ctx, uid)
}

func (s *ServerImpl) BindFingerprint(ctx *gin.Context, req *BindFingerprintRequest) (bool, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	pwd, err := s.db.UserPassword.Query().Where(
		userpassword.UserID(uid),
		userpassword.SceneEQ(userpassword.SceneLogin), userpassword.StatusEQ(typex.SimpleStatusActive),
	).Select(userpassword.FieldUserID, userpassword.FieldSalt, userpassword.FieldPassword).Only(entcache.Skip(ctx))
	if err != nil {
		return false, errors.Codel(errors.ErrPasswordNotMatch)
	}
	given := resource.SaltSecret(req.UserPassword, pwd.Salt)
	if given != pwd.Password {
		return false, errors.Codel(errors.ErrPasswordNotMatch)
	}
	return true, nil
}

func (s *ServerImpl) VerifyFactor(ctx *gin.Context, req *VerifyFactorRequest) (*LoginResponse, error) {
	token := req.StateToken
	id, err := parseStateToken(token, s.Options)
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
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
			return nil, errors.Codel(errors.ErrMfaInvalidCode)
		}
	}

	if profile.PasswordReset {
		return s.resetPasswordPrepare(ctx, profile)
	}

	// no need use transaction
	_ = s.updateLastLogin(ctx, profile.UserID)
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
		ctx.Status(http.StatusUnauthorized)
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
	// 判断密码是否旧密码
	if npwd == pwd.Password {
		return nil, errors.Codel(errors.ErrPasswordDuplicate)
	}

	err = clientx.WithTx(ctx, func(ctx context.Context) (clientx.Transactor, error) {
		return s.db.Tx(ctx)
	}, func(itx clientx.Transactor) error {
		tx := itx.(*ent.Tx)
		err = tx.UserPassword.UpdateOneID(pwd.ID).SetUpdatedBy(uid).SetPassword(npwd).Exec(ctx)
		if err != nil {
			return err
		}
		// 修改密码，清除token
		_ = s.clearLoginTokensOfRedis(ctx, uid)
		res, err = s.loginToken(ctx, uid)
		if err != nil {
			return err
		}
		_ = s.updateLastLogin(ctx, uid)
		s.cache.Del(ctx, cacheKey) // lint:ignore
		return nil
	})
	return
}

func (s *ServerImpl) resetPasswordPrepare(ctx *gin.Context, profile *ent.UserLoginProfile) (res *LoginResponse, err error) {
	sid := uuid.New().String()
	ctx.Status(http.StatusAccepted)
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
		return nil, errors.Codel(errors.ErrMfaNotActive)
	}
	sid := uuid.New().String()
	ctx.Status(http.StatusAccepted)
	res = &LoginResponse{
		CallbackUrl: callBackUrlMFA,
		StateToken:  createStateToken(sid, s.Options),
	}
	err = s.cache.Set(ctx, mfaCachePrefix+sid, profile.ID, cache.WithTTL(s.StateTokenTTL))
	return
}

// VerifyDeviceSendEmail 验证登录设备 发送邮件验证码
func (s *ServerImpl) VerifyDeviceSendEmail(ctx *gin.Context, req *VerifyDeviceSendEmailRequest) (string, error) {
	token := req.StateToken
	id, err := parseStateToken(token, s.Options)
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		return "", err
	}
	var uid int
	cacheKey := verifyDeviceCachePrefix + id
	if err = s.cache.Get(ctx, cacheKey, &uid); err != nil {
		return "", err
	}
	// 生成验证码
	captchaId := captcha.NewLen(6)
	digits := s.captchaStore.Get(captchaId, true)
	captchaCode := ""
	for _, v := range digits {
		captchaCode = captchaCode + strconv.Itoa(int(v))
	}
	// 保存到redis
	captchaId = uuid.New().String()
	captchaKey := verifyDeviceCachePrefix + captchaId
	err = s.cache.Set(ctx, captchaKey, captchaCode, cache.WithTTL(s.CaptchaExpire))
	if err != nil {
		return "", err
	}
	usr, err := s.db.User.Get(ctx, uid)
	if err != nil {
		return "", err
	}
	addr, err := usr.QueryAddresses().Where(useraddr.AddrTypeEQ(useraddr.AddrTypeContact), useraddr.EmailEqualFold(req.Email)).Only(ctx)
	if err != nil {
		return "", errors.Codel(errors.ErrEmailVerify)
	}
	uorg, err := s.GetUserRootOrg(ctx, uid)
	if err != nil {
		return "", err
	}
	tid, err := s.getTopOrgIdByPath(uorg.Path)
	if err != nil {
		return "", err
	}
	params := msg.PostableAlerts{
		{
			Annotations: map[string]string{
				"to":            addr.Email,
				"displayName":   usr.DisplayName,
				"captchaCode":   captchaCode,
				"captchaExpire": strconv.Itoa(int(s.CaptchaExpire.Minutes())),
			},
			Alert: &msg.Alert{
				Labels: map[string]string{
					"receiver":  "email",
					"alertname": "SendCaptchaCode",
					"tenant":    strconv.Itoa(tid),
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

func (s *ServerImpl) CheckDevice(ctx *gin.Context, req *CheckDeviceRequest) (*CheckDeviceResponse, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	usr, err := s.db.User.Get(ctx, uid)
	if err != nil {
		return nil, err
	}
	profile, err := usr.QueryLoginProfile().Only(ctx)
	if err != nil {
		return nil, err
	}
	// 设备验证：开启设备验证及传递了deviceId
	if profile.VerifyDevice && req.DeviceInfo.DeviceUid != "" {
		// 判断是否忽略账户
		excludeAccounts := s.VerifyDeviceParams.ExcludeAccounts
		if excludeAccounts != nil && len(excludeAccounts) > 0 {
			identities, err := usr.QueryIdentities().All(ctx)
			if err != nil {
				return nil, err
			}
			for _, iden := range identities {
				// 判断是否排除验证
				for _, excludeAccount := range excludeAccounts {
					if iden.Code == excludeAccount {
						return &CheckDeviceResponse{VerifyDevice: false}, nil
					}
				}
			}
		}
		has, err := s.db.UserDevice.Query().Where(userdevice.UserID(profile.UserID)).Exist(ctx)
		if err != nil {
			return nil, err
		}
		// 如果没有设备绑定，允许默认绑定或非第一次登录则无需验证设备
		key := isLogonCachePrefix + strconv.Itoa(uid)
		var isLogon bool
		_ = s.cache.Get(ctx, key, &isLogon)
		if !has && (s.VerifyDeviceParams.DefaultBound || isLogon) {
			// 未绑定设备，默认直接绑定
			ctx1 := securityX.WithContext(ctx, securityX.NewGenericPrincipalByClaims(jwt.MapClaims{
				"sub": strconv.Itoa(uid),
			}))
			err = s.db.UserDevice.Create().SetInput(ent.CreateUserDeviceInput{
				DeviceName:    &req.DeviceInfo.DeviceName,
				DeviceModel:   &req.DeviceInfo.DeviceModel,
				DeviceUID:     req.DeviceInfo.DeviceUid,
				SystemName:    &req.DeviceInfo.SystemName,
				SystemVersion: &req.DeviceInfo.SystemVersion,
				AppVersion:    &req.DeviceInfo.AppVersion,
				Comments:      &req.DeviceInfo.Comments,
			}).SetStatus(typex.SimpleStatusActive).SetUserID(uid).SetUpdatedBy(profile.UserID).Exec(ctx1)
			if err != nil {
				return nil, err
			}
			return &CheckDeviceResponse{VerifyDevice: false}, nil
		}
		// 查询设备绑定数
		q, err := s.db.Quota.Query().Where(
			quota.TenantIDIsNil(),
			quota.UserID(profile.UserID),
			quota.HasQuotaItemWith(
				quotaitem.Code(string(quotaService.ItemCodeUserDevice)),
			),
		).Only(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, err
		}
		// 验证设备
		has, err = s.db.UserDevice.Query().Where(
			userdevice.UserID(profile.UserID),
			userdevice.DeviceUID(req.DeviceInfo.DeviceUid),
		).Exist(ctx)
		if err != nil {
			return nil, err
		}
		if !has {
			// 判断设备是否超出限制
			if q != nil && q.Used >= q.Limit {
				return nil, errors.Codel(errors.ErrUserDeviceLimit, q.Limit)
			}
			// 进行设备验证
			sid := uuid.New().String()
			verifies := make([]*ForgetPwdVerify, 0)
			if profile.MfaEnabled {
				verifies = append(verifies, &ForgetPwdVerify{Kind: "mfa"})
			}
			addr, err := usr.QueryAddresses().Where(useraddr.AddrTypeEQ(useraddr.AddrTypeContact)).Only(ctx)
			if err != nil {
				return nil, err
			}
			if &addr.Email != nil {
				verifies = append(verifies, &ForgetPwdVerify{Kind: "email", Value: resource.MaskEmail(addr.Email)})
			}
			res := &CheckDeviceResponse{
				VerifyDevice: true,
				StateToken:   createStateToken(sid, s.Options),
				Verifies:     verifies,
			}
			err = s.cache.Set(ctx, verifyDeviceCachePrefix+sid, profile.UserID, cache.WithTTL(s.Options.StateTokenTTL))
			if err != nil {
				return nil, err
			}
			return res, nil
		} else {
			// 正常登录，更新设备信息
			ud, _ := s.db.UserDevice.Query().Where(userdevice.UserID(profile.UserID), userdevice.DeviceUID(req.DeviceInfo.DeviceUid)).Only(ctx)
			if ud != nil {
				_ = s.db.UserDevice.UpdateOne(ud).SetUpdatedBy(profile.UserID).SetDeviceModel(req.DeviceInfo.DeviceModel).SetDeviceName(req.DeviceInfo.DeviceName).
					SetAppVersion(req.DeviceInfo.AppVersion).SetSystemVersion(req.DeviceInfo.SystemVersion).Exec(ctx)
			}
		}
	}
	return &CheckDeviceResponse{VerifyDevice: false}, nil
}

// VerifyDevice 验证登录设备并绑定
func (s *ServerImpl) VerifyDevice(ctx *gin.Context, req *VerifyDeviceRequest) (*LoginResponse, error) {
	token := req.StateToken
	id, err := parseStateToken(token, s.Options)
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		return nil, err
	}
	var uid int
	cacheKey := verifyDeviceCachePrefix + id
	if err = s.cache.Get(ctx, cacheKey, &uid); err != nil {
		return nil, err
	}
	captchaKey := verifyDeviceCachePrefix + req.CaptchaId
	if req.Kind == KindEmail {
		// 验证验证码
		var captchaCode string
		if err = s.cache.Get(ctx, captchaKey, &captchaCode); err != nil {
			return nil, errors.Codel(errors.ErrCaptchaInvalid)
		}
		if captchaCode != req.Captcha {
			return nil, errors.Codel(errors.ErrCaptchaNotMatch)
		}
	} else if req.Kind == KindMfa {
		profile, err := s.db.UserLoginProfile.Query().Where(userloginprofile.UserID(uid)).Only(ctx)
		if err != nil {
			return nil, err
		}
		// 验证mfa
		if profile.MfaEnabled {
			if !totp.Validate(req.OtpToken, profile.MfaSecret) {
				return nil, errors.Codel(errors.ErrMfaInvalidCode)
			}
		} else {
			return nil, errors.Codel(errors.ErrMfaDisable)
		}
	} else {
		return nil, errors.Codel(errors.ErrUnsupportedVerify)
	}
	// 保存设备信息
	client := s.db
	// ctx增加用户信息
	ctx1 := securityX.WithContext(ctx, securityX.NewGenericPrincipalByClaims(jwt.MapClaims{
		"sub": strconv.Itoa(uid),
	}))
	// 如果超出设备限制数，则提醒用户
	uds, err := client.UserDevice.Query().Where(userdevice.UserID(uid)).Count(ctx1)
	if err != nil {
		return nil, err
	}
	q, err := s.db.Quota.Query().Where(
		quota.TenantIDIsNil(),
		quota.UserID(uid),
		quota.HasQuotaItemWith(
			quotaitem.Code(string(quotaService.ItemCodeUserDevice)),
		),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if q != nil && int64(uds) >= q.Limit {
		return nil, errors.Codel(errors.ErrUserDeviceLimit, q.Limit)
	}
	err = client.UserDevice.Create().SetInput(ent.CreateUserDeviceInput{
		DeviceName:    &req.DeviceInfo.DeviceName,
		DeviceModel:   &req.DeviceInfo.DeviceModel,
		DeviceUID:     req.DeviceInfo.DeviceUid,
		SystemName:    &req.DeviceInfo.SystemName,
		SystemVersion: &req.DeviceInfo.SystemVersion,
		AppVersion:    &req.DeviceInfo.AppVersion,
		Comments:      &req.DeviceInfo.Comments,
	}).SetStatus(typex.SimpleStatusActive).SetUserID(uid).Exec(ctx1)
	if err != nil {
		return nil, err
	}
	// no need use transaction
	_ = s.updateLastLogin(ctx, uid)
	loginResp, err := s.loginToken(ctx, uid)
	if err != nil {
		return nil, err
	}
	// 发送新设备登录提醒
	usr, addr, err := s.getUserInfo(ctx, uid)
	if err != nil {
		return nil, err
	}
	loginTime := time.Now().Format("2006-01-02 15:04:05")
	params := msg.PostableAlerts{
		{
			Annotations: map[string]string{
				"to":            addr.Email,
				"displayName":   loginResp.User.DisplayName,
				"principalName": usr.PrincipalName,
				"loginTime":     loginTime,
				"deviceName":    req.DeviceInfo.DeviceName,
			},
			Alert: &msg.Alert{
				Labels: map[string]string{
					"receiver":  "email",
					"alertname": "NewDeviceLogin",
					"tenant":    strconv.Itoa(loginResp.User.Domains[0].ParentID),
					"timestamp": strconv.Itoa(int(time.Now().Unix())),
				},
			},
		},
	}
	err = s.postAlerts(ctx, params)
	if err != nil {
		return nil, err
	}
	_ = s.cache.Del(ctx, cacheKey)
	_ = s.cache.Del(ctx, captchaKey)
	return loginResp, nil
}

func (s *ServerImpl) getUserInfo(ctx *gin.Context, uid int) (*ent.User, *ent.UserAddr, error) {
	usr, err := s.db.User.Get(ctx, uid)
	if err != nil {
		return nil, nil, err
	}
	addr, err := usr.QueryAddresses().Where(useraddr.AddrTypeEQ(useraddr.AddrTypeContact)).Only(ctx)
	if err != nil {
		return nil, nil, err
	}
	return usr, addr, nil
}

func (s *ServerImpl) updateLastLogin(ctx *gin.Context, uid int) error {
	pc := s.db.UserLoginProfile
	// 判断是否有ip及loginAt来确定是否首次登录
	profile, err := pc.Query().Where(userloginprofile.UserID(uid)).Only(ctx)
	if err != nil {
		return err
	}
	key := isLogonCachePrefix + strconv.Itoa(uid)
	if profile.LastLoginIP != "" && !profile.LastLoginAt.IsZero() {
		// 有值，登录过
		err = s.cache.Set(ctx, key, true, cache.WithTTL(time.Minute*1))
	} else {
		err = s.cache.Set(ctx, key, false, cache.WithTTL(time.Minute*1))
	}
	//
	cip := ctx.ClientIP()
	// no mater what, update last login time and ip
	return pc.Update().Where(userloginprofile.UserID(uid)).
		SetLastLoginIP(cip).SetUpdatedBy(uid).SetPasswordReset(false).
		SetLastLoginAt(time.Now()).Exec(ctx)
}

func (s *ServerImpl) loginToken(ctx *gin.Context, uid int) (*LoginResponse, error) {
	usr := s.db.User.GetX(entcache.Skip(ctx), uid)

	tid, tstr, err := createToken(strconv.Itoa(uid), s.Options, false)
	if err != nil {
		return nil, err
	}

	_, trstr, err := createToken(strconv.Itoa(uid), s.Options, true)
	if err != nil {
		return nil, err
	}

	err = s.cache.Set(ctx, tid, strconv.Itoa(uid), cache.WithTTL(s.Options.JWT.TokenTTL))
	if err != nil {
		return nil, err
	}

	ros, err := s.db.Org.Query().Where(
		org.HasOrgUserWith(orguser.UserID(uid)),
		org.StatusEQ(typex.SimpleStatusActive),
		org.KindEQ(org.KindRoot),
	).Select(org.FieldID, org.FieldName, org.FieldPath, org.FieldLocalCurrency).Order(ent.Asc(org.FieldID)).All(ctx)
	if err != nil {
		return nil, err
	}
	if ros == nil || len(ros) == 0 {
		return nil, errors.Codel(errors.ErrOrgNotFound)
	}
	var domains []*Domain
	// 查询顶级组织
	for _, o := range ros {
		// 根据/截取path的第一项
		code := strings.Split(o.Path, "/")[0]
		// 转换成十进制id
		oID, err := strconv.ParseInt(code, 36, 64)
		if err != nil {
			return nil, err
		}
		to, err := s.db.Org.Query().Where(org.ID(int(oID))).Select(org.FieldID, org.FieldName, org.FieldPath, org.FieldLocalCurrency).Only(ctx)
		if err != nil {
			return nil, err
		}
		domains = append(domains, &Domain{
			ID:             o.ID,
			Name:           o.Name,
			LocalCurrency:  o.LocalCurrency,
			ParentID:       to.ID,
			ParentName:     to.Name,
			ParentCurrency: to.LocalCurrency,
		})
	}
	return &LoginResponse{
		AccessToken:  tstr,
		ExpiresIn:    int(s.Options.JWT.TokenTTL.Seconds()),
		RefreshToken: trstr,
		User: &User{
			ID:          usr.ID,
			DisplayName: usr.DisplayName,
			Avatar:      usr.Avatar,
			Domains:     domains,
		},
	}, nil
}

func (s *ServerImpl) checkPwd(ctx *gin.Context, req *LoginRequest) (*ent.UserPassword, error) {
	pwd, err := s.db.UserPassword.Query().Where(
		userpassword.HasUserWith(user.HasIdentitiesWith(useridentity.Code(req.Username))),
		userpassword.SceneEQ(userpassword.SceneLogin), userpassword.StatusEQ(typex.SimpleStatusActive),
	).Select(userpassword.FieldUserID, userpassword.FieldSalt, userpassword.FieldPassword).Only(entcache.Skip(ctx))
	if err != nil {
		return nil, errors.Codel(errors.ErrUserNameOrPassword)
	}

	given := resource.SaltSecret(req.Password, pwd.Salt)
	if given != pwd.Password {
		return nil, errors.Codel(errors.ErrPasswordNotMatch)
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
	key, err := auth.ParseSigningKeyFromString(opts.JWT.PrivateKey, opts.JWT.SigningMethod, true)
	if err != nil {
		return "", "", err
	}
	tokenStr, err = token.SignedString(key)
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
		err = errors.Codel(errors.ErrInvalidToken)
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
	issuer, err := s.ParentDomain(ctx, tid)
	if err != nil {
		return nil, err
	}
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

func (s *ServerImpl) ParentDomain(ctx context.Context, orgID int) (string, error) {
	o, err := s.db.Org.Query().Where(org.ID(orgID)).Only(ctx)
	if err != nil {
		return "", err
	}
	if o.Domain == "" && o.ParentID == 0 {
		return "", nil
	}
	if o.Domain != "" {
		return o.Domain, nil
	}
	return s.ParentDomain(ctx, o.ParentID)
}

func (s *ServerImpl) BindMfa(ctx *gin.Context, req *BindMfaRequest) (bool, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	id, err := parseStateToken(req.StateToken, s.Options)
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		return false, err
	}
	//
	var val map[string]string
	if err = s.cache.Get(ctx, mfaCachePrefix+id, &val); err != nil {
		return false, err
	}
	if val["uid"] != strconv.Itoa(uid) {
		return false, errors.Codel(errors.ErrInvalidUser)
	}
	if !totp.Validate(req.OtpToken, val["secret"]) {
		return false, errors.Codel(errors.ErrMfaInvalidCode)
	}
	err = s.db.UserLoginProfile.Update().Where(userloginprofile.UserID(uid)).SetMfaEnabled(true).SetMfaStatus(typex.SimpleStatusActive).SetMfaSecret(val["secret"]).Exec(ctx)
	return err == nil, err
}

func (s *ServerImpl) UnBindMfa(ctx *gin.Context, req *UnBindMfaRequest) (bool, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	up, err := s.db.UserLoginProfile.Query().Where(userloginprofile.UserID(uid)).Only(ctx)
	if !totp.Validate(req.OtpToken, up.MfaSecret) {
		return false, errors.Codel(errors.ErrMfaInvalidCode)
	}
	err = s.db.UserLoginProfile.UpdateOneID(up.ID).ClearMfaEnabled().ClearMfaStatus().ClearMfaSecret().Exec(ctx)
	return err == nil, err
}

func (s *ServerImpl) GetUserRootOrg(ctx *gin.Context, uid int) (uorg *ent.Org, err error) {
	uorg, err = s.db.OrgUser.Query().Where(orguser.UserIDEQ(uid)).
		QueryOrg().Unique(false).Where(
		org.KindEQ(org.KindRoot),
		org.StatusEQ(typex.SimpleStatusActive),
	).Order(ent.Desc(org.FieldPath)).First(ctx)
	if err != nil {
		return nil, err
	}
	return uorg, nil
}
func (s *ServerImpl) getTopOrgIdByPath(path string) (int, error) {
	code := strings.Split(path, "/")[0]
	oID, err := strconv.ParseInt(code, 36, 64)
	if err != nil {
		return 0, err
	}
	return int(oID), nil
}

func (s *ServerImpl) logFailHandler(ctx *gin.Context, uid string, clear bool) (int32, error) {
	key := loginFailCachePrefix + uid
	if clear {
		return 0, s.cache.Del(ctx, key)
	}
	var count int32 = 0
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
		return nil, errors.Codel(errors.ErrCaptchaNotMatch)
	}
	// 查询用户
	u, err := s.db.User.Query().Where(user.HasIdentitiesWith(useridentity.Code(req.Username))).WithLoginProfile().Only(ctx)
	if err != nil {
		return nil, err
	}
	// 判断用户锁定不能重置密码
	if u.Status == types.UserStatusLocked {
		// 返回账号锁定错误
		return nil, errors.Codel(errors.ErrUserHasLocked)
	}
	verifies := make([]*ForgetPwdVerify, 0)
	if u.Edges.LoginProfile.MfaEnabled {
		verifies = append(verifies, &ForgetPwdVerify{Kind: "mfa"})
	}
	addr, err := u.QueryAddresses().Where(useraddr.AddrTypeEQ(useraddr.AddrTypeContact)).Only(ctx)
	if err != nil {
		return nil, err
	}
	if &addr.Email != nil {
		verifies = append(verifies, &ForgetPwdVerify{Kind: "email", Value: resource.MaskEmail(addr.Email)})
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
		ctx.Status(http.StatusUnauthorized)
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
	// 判断密码是否旧密码
	if npwd == pwd.Password {
		return false, errors.Codel(errors.ErrPasswordDuplicate)
	}

	err = clientx.WithTx(ctx, func(ctx context.Context) (clientx.Transactor, error) {
		return s.db.Tx(ctx)
	}, func(itx clientx.Transactor) error {
		tx := itx.(*ent.Tx)
		// SetStatus用于处理密码过期状态恢复
		err = tx.UserPassword.UpdateOneID(pwd.ID).SetUpdatedBy(uid).SetPassword(npwd).SetStatus(typex.SimpleStatusActive).Exec(ctx)
		if err != nil {
			return err
		}
		// 修改密码，清除token
		_ = s.clearLoginTokensOfRedis(ctx, uid)
		usr, addr, err := s.getUserInfo(ctx, uid)
		if err != nil {
			return err
		}
		if addr.Email == "" {
			return errors.Codel(errors.ErrEmailEmpty)
		}
		uorg, err := s.GetUserRootOrg(ctx, usr.ID)
		if err != nil {
			return err
		}
		tid, err := s.getTopOrgIdByPath(uorg.Path)
		if err != nil {
			return err
		}
		// 增加重置密码邮件通知
		params := msg.PostableAlerts{
			{
				Annotations: map[string]string{
					"to":            addr.Email,
					"displayName":   usr.DisplayName,
					"principalName": usr.PrincipalName,
				},
				Alert: &msg.Alert{
					Labels: map[string]string{
						"receiver":  "email",
						"alertname": "ChangeUserPassword",
						"tenant":    strconv.Itoa(tid),
						"timestamp": strconv.Itoa(int(time.Now().Unix())),
					},
				},
			},
		}
		_ = s.postAlerts(ctx, params)
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
		ctx.Status(http.StatusUnauthorized)
		return "", err
	}
	var uid int
	cacheKey := forgetPwdBeginCachePrefix + id
	if err = s.cache.Get(ctx, cacheKey, &uid); err != nil {
		return "", err
	}
	// 生成验证码
	captchaId := captcha.NewLen(6)
	digits := s.captchaStore.Get(captchaId, true)
	captchaCode := ""
	for _, v := range digits {
		captchaCode = captchaCode + strconv.Itoa(int(v))
	}
	// 保存到redis
	captchaId = uuid.New().String()
	captchaKey := forgetPwdBeginCachePrefix + captchaId
	err = s.cache.Set(ctx, captchaKey, captchaCode, cache.WithTTL(s.CaptchaExpire))
	if err != nil {
		return "", err
	}
	usr, addr, err := s.getUserInfo(ctx, uid)
	if err != nil {
		return "", err
	}
	if addr.Email == "" {
		return "", errors.Codel(errors.ErrEmailEmpty)
	}
	uorg, err := s.GetUserRootOrg(ctx, usr.ID)
	if err != nil {
		return "", err
	}
	tid, err := s.getTopOrgIdByPath(uorg.Path)
	if err != nil {
		return "", err
	}
	params := msg.PostableAlerts{
		{
			Annotations: map[string]string{
				"to":            addr.Email,
				"displayName":   usr.DisplayName,
				"captchaCode":   captchaCode,
				"captchaExpire": strconv.Itoa(int(s.CaptchaExpire.Minutes())),
			},
			Alert: &msg.Alert{
				Labels: map[string]string{
					"receiver":  "email",
					"alertname": "SendCaptchaCode",
					"tenant":    strconv.Itoa(tid),
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
		ctx.Status(http.StatusUnauthorized)
		return nil, err
	}
	var uid int
	cacheKey := forgetPwdBeginCachePrefix + id
	if err = s.cache.Get(ctx, cacheKey, &uid); err != nil {
		return nil, err
	}
	// 验证验证码
	var captchaCode string
	captchaKey := forgetPwdBeginCachePrefix + req.CaptchaId
	if err = s.cache.Get(ctx, captchaKey, &captchaCode); err != nil {
		return nil, errors.Codel(errors.ErrCaptchaInvalid)
	}
	if captchaCode != req.Captcha {
		return nil, errors.Codel(errors.ErrCaptchaNotMatch)
	}
	sid := uuid.New().String()
	stateToken := createStateToken(sid, s.Options)
	err = s.cache.Set(ctx, forgetPwdVerifyCachePrefix+sid, uid, cache.WithTTL(s.Options.StateTokenTTL))
	if err != nil {
		return nil, err
	}
	s.cache.Del(ctx, cacheKey)
	s.cache.Del(ctx, captchaKey)
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
		ctx.Status(http.StatusUnauthorized)
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
			return nil, errors.Codel(errors.ErrMfaInvalidCode)
		}
	} else {
		return nil, errors.Codel(errors.ErrMfaDisable)
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

func (s *ServerImpl) PasswordPolicy(ctx *gin.Context) (*UserPasswordPolicy, error) {
	upp, err := s.getPasswordPolicy(ctx)
	if err != nil {
		return &UserPasswordPolicy{
			Length:               int(s.PwdPolicy.Length),
			IncludeElement:       int(s.PwdPolicy.IncludeElement),
			IncludeChar:          int(s.PwdPolicy.IncludeChar),
			AllowIncludeUserName: s.PwdPolicy.AllowIncludeUserName,
			InvalidDay:           int(s.PwdPolicy.InvalidDay),
			InvalidLoginLimit:    s.PwdPolicy.InvalidLoginLimit,
			Retry:                int(s.PwdPolicy.Retry),
			CaptchaTimes:         int(s.PwdPolicy.CaptchaTimes),
		}, nil
	}
	return &UserPasswordPolicy{
		Length:               int(upp.Length),
		IncludeElement:       int(upp.IncludeElement),
		IncludeChar:          int(upp.IncludeChar),
		AllowIncludeUserName: upp.AllowIncludeUserName,
		InvalidDay:           int(upp.InvalidDay),
		InvalidLoginLimit:    upp.InvalidLoginLimit,
		Retry:                int(upp.Retry),
		CaptchaTimes:         int(upp.CaptchaTimes),
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
		return errors.Codel(errors.ErrOrgNotFound)
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
		return nil, errors.Codel(errors.ErrInvalidSpm)
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

func (s *ServerImpl) getFileIdentity(c *gin.Context, bucket, endpoint string, tid int) (*ent.FileIdentity, error) {
	ctx := identity.WithTenantID(c, tid)
	var fi *ent.FileIdentity
	var err error
	if bucket != "" && endpoint != "" {
		// 传参取对应identity
		fi, err = s.db.FileIdentity.Query().Where(
			fileidentity.TenantID(tid),
			fileidentity.HasSourceWith(
				filesource.Endpoint(endpoint),
				filesource.Bucket(bucket),
			),
		).WithSource().Only(ctx)
	} else {
		// 不传参取默认值
		fi, err = s.db.FileIdentity.Query().Where(
			fileidentity.TenantID(tid),
			fileidentity.IsDefault(true),
		).WithSource().Only(ctx)
	}
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if fi == nil {
		t, err := s.db.Org.Get(ctx, tid)
		if err != nil {
			return nil, err
		}
		if t.ParentID == 0 {
			return nil, errors.Codel(errors.ErrFileIdentityIsNull)
		}
		return s.getFileIdentity(c, bucket, endpoint, t.ParentID)
	}
	return fi, nil
}

func (s *ServerImpl) GetSTS(c *gin.Context, req *GetSTSRequest) (*GetSTSResponse, error) {
	uid, err := identity.UserIDFromContext(c)
	if err != nil {
		return nil, err
	}
	tid, err := s.tryGetTenantID(c)
	if err != nil {
		return nil, err
	}
	fi, err := s.getFileIdentity(c, req.Bucket, req.Endpoint, tid)
	if err != nil {
		return nil, err
	}

	err = s.kosdk.Fs().RegistryProvider(s.toProviderConfig(fi), fs.GetProviderKey(s.toProviderConfig(fi)))
	if err != nil {
		return nil, err
	}
	provider, err := s.kosdk.Fs().GetProviderByBizKey(fs.GetProviderKey(s.toProviderConfig(fi)))
	if err != nil {
		return nil, err
	}
	usr, err := s.db.User.Get(c, uid)
	if err != nil {
		return nil, err
	}
	resp, err := provider.GetSTS(c, usr.PrincipalName)
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
	tid, err := s.tryGetTenantID(ctx)
	if err != nil {
		return nil, err
	}
	fi, path, err := s.convertUrlToFileSource(ctx, req, tid)
	if err != nil {
		return nil, err
	}
	err = s.kosdk.Fs().RegistryProvider(s.toProviderConfig(fi), fs.GetProviderKey(s.toProviderConfig(fi)))
	if err != nil {
		return nil, err
	}
	provider, err := s.kosdk.Fs().GetProviderByBizKey(fs.GetProviderKey(s.toProviderConfig(fi)))
	if err != nil {
		return nil, err
	}
	signUrl, err := provider.GetPreSignedURL(ctx, fi.Edges.Source.Bucket, path, time.Hour)
	if err != nil {
		return nil, err
	}
	return &GetPreSignUrlResponse{
		URL: signUrl,
	}, nil
}

// convertUrlToFileSource 将url转换为文件源
func (s *ServerImpl) convertUrlToFileSource(c *gin.Context, req *GetPreSignUrlRequest, tid int) (*ent.FileIdentity, string, error) {
	ctx := identity.WithTenantID(c, tid)
	var fis []*ent.FileIdentity
	// 取出组织对应的fileidentities
	var err error
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
		t, err := s.db.Org.Get(ctx, tid)
		if err != nil {
			return nil, "", err
		}
		if t.ParentID == 0 {
			return nil, "", errors.Codel(errors.ErrFileIdentityIsNull)
		}
		return s.convertUrlToFileSource(c, req, t.ParentID)
	}

	// 解析url的path
	path := ""
	u, err := url.Parse(req.URL)
	if fi.Edges.Source.Kind == filesource.KindMinio {
		path = strings.TrimPrefix(u.Path, "/"+fi.Edges.Source.Bucket)
	} else {
		path = u.Path
	}
	return fi, path, nil
}

func (s *ServerImpl) toProviderConfig(fi *ent.FileIdentity) *fs.ProviderConfig {
	return &fs.ProviderConfig{
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

func (s *ServerImpl) doCheckPermission(ctx context.Context, uid, tid int, action, appCode string) (bool, error) {
	has, err := s.db.AppAction.Query().Where(appaction.Name(action), appaction.HasAppWith(app.Code(appCode))).Exist(ctx)
	if err != nil {
		return false, err
	}
	if !has {
		return false, errors.Codel(errors.ErrInvalidPermission)
	}
	rule := []any{
		strconv.Itoa(uid),
		strconv.Itoa(tid),
		fmt.Sprintf("%s%s%s", appCode, authz.ArnSplit, action),
		authz.ActionTypeRead,
	}
	has, err = security.CheckUserPermission(rule...)
	if err != nil {
		return false, err
	}
	if !has {
		return false, nil
	}
	return true, nil
}

func (s *ServerImpl) getPasswordPolicy(ctx *gin.Context) (*ent.UserPasswordPolicy, error) {
	referer := ctx.GetHeader("Referer")
	if referer == "" {
		return s.defaultPwdPolicy(), nil
	}
	// 获取host
	u, err := url.Parse(referer)
	if err != nil {
		return s.defaultPwdPolicy(), nil
	}
	host := u.Hostname()
	// 先根据host找domain
	o, err := s.db.Org.Query().Where(org.Domain(host), org.ParentID(0)).Only(ctx)
	// 如果没找到，再找自定义域名
	if ent.IsNotFound(err) {
		orgs, _ := s.db.Org.Query().Where(org.ParentID(0)).All(ctx)
		for _, or := range orgs {
			// 获取自定义域名
			customDomains := or.CustomDomain
			for _, customDomain := range customDomains {
				if customDomain == host {
					o = or
					break
				}
			}
			if o != nil {
				break
			}
		}
	}
	if o == nil {
		return s.defaultPwdPolicy(), nil
	}
	upp, err := s.db.UserPasswordPolicy.Query().Where(userpasswordpolicy.TenantID(o.ID)).Only(ctx)
	if err != nil || upp == nil {
		return s.defaultPwdPolicy(), nil
	}
	return upp, nil
}

func (s *ServerImpl) defaultPwdPolicy() *ent.UserPasswordPolicy {
	var upp = ent.UserPasswordPolicy{
		Length:               s.PwdPolicy.Length,
		IncludeElement:       s.PwdPolicy.IncludeElement,
		IncludeChar:          s.PwdPolicy.IncludeChar,
		AllowIncludeUserName: s.PwdPolicy.AllowIncludeUserName,
		InvalidDay:           s.PwdPolicy.InvalidDay,
		InvalidLoginLimit:    s.PwdPolicy.InvalidLoginLimit,
		Retry:                s.PwdPolicy.Retry,
		CaptchaTimes:         s.PwdPolicy.CaptchaTimes,
	}
	return &upp
}

func (s *ServerImpl) GetDomain(ctx *gin.Context, req *GetDomainRequest) (*Domain, error) {
	o, err := s.db.Org.Get(ctx, req.OrgID)
	if err != nil {
		return nil, err
	}
	pID, err := s.getTopOrgIdByPath(o.Path)
	if err != nil {
		return nil, err
	}
	po, err := s.db.Org.Get(ctx, pID)
	if err != nil {
		return nil, err
	}
	return &Domain{
		ID:             o.ID,
		Name:           o.Name,
		LocalCurrency:  o.LocalCurrency,
		ParentID:       po.ID,
		ParentName:     po.Name,
		ParentCurrency: po.LocalCurrency,
	}, nil
}

func (s *ServerImpl) clearLoginTokensOfRedis(ctx context.Context, uid int) error {
	// 判断是否有redis实例
	if s.redisClient == nil {
		return nil
	}
	// 获取用户相关的token
	var cursor uint64
	allKeys := make([]string, 0)
	for {
		var keys []string
		var err error
		keys, cursor, err = s.redisClient.Scan(ctx, cursor, fmt.Sprintf("%s%d:*", tokenCachePrefix, uid), 100).Result()
		if err != nil {
			return err
		}
		allKeys = append(allKeys, keys...)
		if cursor == 0 {
			break
		}
	}
	if allKeys == nil {
		return nil
	}
	// keys从redis移除
	_, err := s.redisClient.Del(ctx, allKeys...).Result()
	if err != nil {
		return err
	}

	return nil
}
