package server

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/woocoos/entco/pkg/koapp"
	"github.com/woocoos/knockout/ent/useridentity"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/api/oas"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/migrate"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
	"github.com/woocoos/knockout/service/resource"
	"github.com/woocoos/knockout/status"
	"github.com/woocoos/knockout/test"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/knockout/ent/runtime"
)

var (
	adminToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI2N2E4NzQ4MmU5MWY0ZjJlOTIyMGY1MTM3NjE4NWI3ZSIsInN1YiI6IjEiLCJuYW1lIjoiYWRtaW4iLCJpYXQiOjE1MTYyMzkwMjJ9.MzfOsaGAtHj0IIoVDgpfUD1GMtgLTNbY7D_CkEoR9CQ"
	adminTokenJTI = "67a87482e91f4f2e9220f51376185b7e"
	appCnf        *conf.AppConfiguration
)

type ServiceSuite struct {
	AuthService *AuthService
	FileService *FileService
	redisServer *miniredis.Miniredis

	server *web.Server
}

func (s *ServiceSuite) SetupSuite(t *testing.T) {
	s.redisServer = miniredis.RunT(t)
	require.NoError(t, s.redisServer.Set(adminTokenJTI, "1"))

	file := filepath.Join(test.BaseDir(), "testdata", "etc", "app.yaml")
	bs, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	cnf := conf.NewFromBytes(bs, conf.WithBaseDir(test.BaseDir())).AsGlobal()
	cnf.Parser().Set("cache.redis.addrs", []string{s.redisServer.Addr()})

	app := koapp.New(woocoo.WithAppConfiguration(cnf))
	appCnf = app.AppConfiguration()

	s.server = web.New(web.WithConfiguration(cnf.Sub("web")))

	s.AuthService = &AuthService{}
	err = s.AuthService.Apply(app.AppConfiguration())
	if err != nil {
		t.Fatal(err)
	}
	drv := koapp.BuildEntComponents(app.AppConfiguration())["test"]
	s.AuthService.DB = ent.NewClient(ent.Driver(drv), ent.Debug())
	//s.AuthService.DB, err = ent.Open(appCnf.String("store.test.driverName"), appCnf.String("store.test.dsn"), ent.Debug())
	s.AuthService.Cache, err = cache.GetCache("redis")
	require.NoError(t, err)
	s.FileService = &FileService{
		DB:       s.AuthService.DB,
		BaseDir:  appCnf.Abs(appCnf.String("files.local.baseDir")),
		Endpoint: appCnf.String("files.local.endpoint"),
	}

	RegisterHandlersManual(&s.server.Router().RouterGroup, s.AuthService)
	RegisterAuthHandlers(&s.server.Router().RouterGroup, s.AuthService)
	RegisterFileHandlers(&s.server.Router().RouterGroup, s.FileService)
	jwtmdl, ok := s.server.HandlerManager().GetMiddleware(web.GetMiddlewareKey(s.server.Router().RouterGroup.BasePath(), "jwt"))
	require.True(t, ok)
	s.AuthService.LogoutHandler = jwtmdl.(*handler.JWTMiddleware).Config.LogoutHandler

	err = s.AuthService.DB.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false))
	assert.NoError(t, err)
}

type loginFlowSuite struct {
	suite.Suite
	*ServiceSuite
}

// loginFlowSuite runs all the tests in the suite.
func TestLoginFlow(t *testing.T) {
	service := ServiceSuite{}
	service.SetupSuite(t)
	suite.Run(t, &loginFlowSuite{
		ServiceSuite: &service,
	})
}

func (ts *loginFlowSuite) SetupSuite() {
	ts.redisServer.FlushAll()
	db := ts.AuthService.DB
	adm := db.User.Create().SetCreatedBy(1).SetID(1).SetPrincipalName("admin").
		SetStatus(typex.SimpleStatusActive).SetDisplayName("admin").SetUserType(user.UserTypeAccount).
		SetCreationType(user.CreationTypeManual).SetRegisterIP("")
	_, err := db.User.CreateBulk(
		adm,
	).Save(context.Background())
	ts.Require().NoError(err)

	err = db.UserIdentity.Create().SetCreatedBy(1).SetUserID(1).SetCode("admin").
		SetStatus(typex.SimpleStatusActive).SetKind(useridentity.KindName).Exec(context.Background())
	ts.Require().NoError(err)
	err = db.UserPassword.Create().SetCreatedBy(1).SetUserID(1).SetStatus(typex.SimpleStatusActive).SetSalt("123456").
		SetPassword("9b1063951d443cfac15cc879efb4054f4f4fd599e1b1a9aee67b0301e19e40fe").SetScene(userpassword.SceneLogin).
		Exec(context.Background())
	ts.Require().NoError(err)

	err = db.UserLoginProfile.Create().SetCreatedBy(1).SetUserID(1).SetSetKind(userloginprofile.SetKindKeep).
		SetMfaEnabled(true).SetMfaSecret("UWZLIIUMPX53NYXB").SetCanLogin(true).
		SetMfaStatus(typex.SimpleStatusActive).SetVerifyDevice(true).
		Exec(context.Background())
	ts.Require().NoError(err)

	err = db.Org.Create().SetID(1).SetCreatedBy(1).SetName("test").SetDomain("test.com").SetCode("test").
		SetKind(org.KindRoot).SetStatus(typex.SimpleStatusActive).SetParentID(0).
		Exec(context.Background())
	ts.Require().NoError(err)

	err = db.OrgUser.Create().SetUserID(1).SetCreatedBy(1).SetOrgID(1).
		SetDisplayName("admin").Exec(context.Background())
	ts.Require().NoError(err)
}

func (ts *loginFlowSuite) Test_AuthNoFlow() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, err := ts.AuthService.Login(ctx, &oas.LoginRequest{
		Body: oas.LoginRequestBody{Password: "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", Username: "admin"},
	})
	ts.Require().NoError(err)
	ts.Require().NotNil(res)
}

func (ts *loginFlowSuite) Test_AuthMFAFlow() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, err := ts.AuthService.Login(ctx, &oas.LoginRequest{
		Body: oas.LoginRequestBody{Password: "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", Username: "admin"},
	})
	ts.Require().NoError(err)
	ts.Contains(res.CallbackUrl, "verifyFactor")
	ts.NotEmptyf(res.StateToken, "state token should not be empty")
}

func (ts *loginFlowSuite) Test_AuthFail() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	for i := 0; i < ts.AuthService.CaptchaTimes-1; i++ {
		res, err := ts.AuthService.Login(ctx, &oas.LoginRequest{
			Body: oas.LoginRequestBody{Password: "error", Username: "admin"},
		})
		ts.Require().ErrorIs(err, status.ErrMismatchPWD)
		ts.Nil(res)
	}
	res, err := ts.AuthService.Login(ctx, &oas.LoginRequest{
		Body: oas.LoginRequestBody{Password: "error", Username: "admin"},
	})
	ts.Require().ErrorIs(err, status.ErrMismatchPWD)
	ts.Equal(res.CallbackUrl, "/captcha")
	for i := ts.AuthService.CaptchaTimes; i < ts.AuthService.LoginFailTimes; i++ {
		res, err := ts.AuthService.Login(ctx, &oas.LoginRequest{
			Body: oas.LoginRequestBody{Password: "error", Username: "admin", Captcha: ""},
		})
		ts.Require().ErrorIs(err, status.ErrCaptchaNotMatch)
		ts.Nil(res)
	}
	res, err = ts.AuthService.Login(ctx, &oas.LoginRequest{
		Body: oas.LoginRequestBody{Password: "error", Username: "admin"},
	})
	// TODO 验证码准确性测试
	ts.Require().ErrorIs(err, status.ErrLoginFailUpperLimit)
	ts.Nil(res)
}

func (ts *loginFlowSuite) Test_VerifyFactor() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	// use admin token as state token
	err := ts.AuthService.Cache.Set(context.Background(), mfaCachePrefix+"67a87482e91f4f2e9220f51376185b7e", 1)
	ts.Require().NoError(err)

	ctx.Request = httptest.NewRequest("POST", "/verifyFactor", nil)
	pwd := GeneratePassCode("UWZLIIUMPX53NYXB")
	time.Sleep(1 * time.Second)
	res, err := ts.AuthService.VerifyFactor(ctx, &oas.VerifyFactorRequest{
		Body: oas.VerifyFactorRequestBody{
			OtpToken:   pwd,
			StateToken: adminToken,
		},
	})
	ts.Require().NoError(err)
	ts.Equal(res.User.ID, 1)
}

// Demo function, not used in main
// Generates Passcode using a UTF-8 (not base32) secret and custom parameters
func GeneratePassCode(base32string string) string {
	passcode, err := totp.GenerateCodeCustom(base32string, time.Now(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		panic(err)
	}
	return passcode
}

func (ts *loginFlowSuite) Test_Logout() {
	err := ts.AuthService.Cache.Set(context.Background(), "67a87482e91f4f2e9220f51376185b7e", "1")
	ts.Require().NoError(err)

	req := httptest.NewRequest("POST", "/logout", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	res := httptest.NewRecorder()
	ts.server.Router().ServeHTTP(res, req)

	ts.Equal(res.Code, 200)
}

func (ts *loginFlowSuite) Test_BindMfaFlow() {
	// 绑定mfa前置数据
	req := httptest.NewRequest("POST", "/mfa/bind-prepare", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("X-Tenant-ID", "1")
	res := httptest.NewRecorder()
	ts.server.Router().ServeHTTP(res, req)
	ts.Require().Equal(res.Code, 200)

	var bp map[string]any
	err := json.Unmarshal([]byte(res.Body.String()), &bp)
	ts.Require().NoError(err)

	// 绑定mfa
	pwd := GeneratePassCode(bp["secret"].(string))
	payload := strings.NewReader(`{
		"stateToken": "` + bp["stateToken"].(string) + `",
		"otpToken": "` + pwd + `"
	}`)
	req = httptest.NewRequest("POST", "/mfa/bind", payload)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("Content-Type", "application/json")
	res = httptest.NewRecorder()
	ts.server.Router().ServeHTTP(res, req)
	ts.Equal(res.Code, 200)

	// 解绑mfa
	pwd = GeneratePassCode(bp["secret"].(string))
	payload = strings.NewReader(`{
		"otpToken": "` + pwd + `"
	}`)
	req = httptest.NewRequest("POST", "/mfa/unbind", payload)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("Content-Type", "application/json")
	res = httptest.NewRecorder()
	ts.server.Router().ServeHTTP(res, req)
	ts.Equal(res.Code, 200)
}

func (ts *loginFlowSuite) Test_SpmFlow() {
	// 绑定mfa前置数据
	req := httptest.NewRequest("POST", "/spm/create", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("X-Tenant-ID", "1")
	res := httptest.NewRecorder()
	ts.server.Router().ServeHTTP(res, req)
	ts.Equal(res.Code, 200)

	spmKey := res.Body.String()
	// 获取登录信息
	payload := strings.NewReader(`{
		"spm": ` + spmKey + `
	}`)
	req = httptest.NewRequest("POST", "/spm/auth", payload)
	req.Header.Set("Content-Type", "application/json")
	res = httptest.NewRecorder()
	ts.server.Router().ServeHTTP(res, req)
	fmt.Println(res.Body.String())
	ts.Equal(res.Code, 200)
}

func (ts *loginFlowSuite) Test_ResetPassword() {
	ts.Require().NoError(ts.AuthService.Cache.Set(context.Background(),
		resetCachePrefix+"67a87482e91f4f2e9220f51376185b7e", 1))
	ctx := security.WithContext(context.Background(), security.NewGenericPrincipalByClaims(jwt.MapClaims{
		"sub": "1",
	}))
	gctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	gctx.Request = httptest.NewRequest("POST", "/resetPassword", nil).WithContext(ctx)
	//ctx = context.WithValue(ctx, gin.ContextKey, gctx)
	res, err := ts.AuthService.ResetPassword(gctx, &oas.ResetPasswordRequest{
		Body: oas.ResetPasswordRequestBody{
			NewPassword: "234567",
			StateToken:  adminToken,
		},
	})
	ts.Require().NoError(err)
	ts.False(ts.redisServer.Exists(resetCachePrefix + "123456"))
	pwd := ts.AuthService.DB.UserPassword.Query().Where(userpassword.UserID(1),
		userpassword.SceneEQ(userpassword.SceneLogin),
	).OnlyX(context.Background())

	ts.Equal(pwd.Password, resource.SaltSecret("234567", "123456"))
	ts.Equal(res.User.ID, 1)
}

func TestPwd(t *testing.T) {
	// 随机字符串
	req := sha256.New()
	req.Write([]byte("234567"))
	param := hex.EncodeToString(req.Sum(nil))
	t.Log(param)
	sha := sha256.New()
	sha.Write([]byte(param))
	sha.Write([]byte("123456"))
	given := hex.EncodeToString(sha.Sum(nil))
	t.Log(given)
}
