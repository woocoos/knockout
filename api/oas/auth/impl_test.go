package auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/alicebob/miniredis/v2"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler"
	"github.com/woocoos/entcache"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/migrate"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
	"github.com/woocoos/knockout/internal/status"
	"github.com/woocoos/knockout/service/resource"
	"github.com/woocoos/knockout/test"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/knockout-go/pkg/snowflake"
	_ "github.com/woocoos/knockout/ent/runtime"
)

var (
	adminToken    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI2N2E4NzQ4MmU5MWY0ZjJlOTIyMGY1MTM3NjE4NWI3ZSIsInN1YiI6IjEiLCJuYW1lIjoiYWRtaW4iLCJpYXQiOjE1MTYyMzkwMjJ9.MzfOsaGAtHj0IIoVDgpfUD1GMtgLTNbY7D_CkEoR9CQ"
	adminTokenJTI = "67a87482e91f4f2e9220f51376185b7e"
)

type authSuite struct {
	suite.Suite
	AuthServer  *ServerImpl
	redisServer *miniredis.Miniredis

	server *web.Server
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, &authSuite{})
}

func (t *authSuite) TestForgetPwdSendEmail() {

}

func (t *authSuite) SetupSuite() {
	t.redisServer = miniredis.RunT(t.T())
	t.Require().NoError(t.redisServer.Set(adminTokenJTI, "1"))

	file := filepath.Join(test.BaseDir(), "testdata", "etc", "app.yaml")
	bs, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	cnf := conf.NewFromBytes(bs, conf.WithBaseDir(test.BaseDir()))
	cnf.Parser().Set("cache.redis.addrs", []string{t.redisServer.Addr()})
	cache.UnRegisterCache("redis")
	println(t.redisServer.Addr())
	app := koapp.New(woocoo.WithAppConfiguration(cnf))
	t.AuthServer = NewServer(app)
	t.server = t.AuthServer.webServer

	jwtmdl, ok := t.server.HandlerManager().GetMiddleware(web.GetMiddlewareKey(t.server.Router().RouterGroup.BasePath(), "jwt"))
	t.Require().True(ok)
	t.AuthServer.LogoutHandler = jwtmdl.(*handler.JWTMiddleware).Config.LogoutHandler

	err = t.AuthServer.db.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false))
	t.Require().NoError(err)
}

type loginFlowSuite struct {
	*authSuite
}

// loginFlowSuite runs all the tests in the suite.
func TestLoginFlow(t *testing.T) {
	service := authSuite{}
	suite.Run(t, &loginFlowSuite{
		authSuite: &service,
	})
}

func (ts *loginFlowSuite) SetupSuite() {
	ts.authSuite.SetupSuite()
	ts.redisServer.FlushAll()
	db := ts.AuthServer.db
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

	err = db.FileSource.Create().SetID(1).SetCreatedBy(1).SetKind(filesource.KindMinio).
		SetEndpoint("http://localhost:9000").SetStsEndpoint("http://localhost:9000").SetRegion("cn-east-1").SetBucket("test2").SetBucketURL("http://localhost:9000/test2").Exec(context.Background())
	ts.Require().NoError(err)

	err = db.FileIdentity.Create().SetID(1).SetCreatedBy(1).SetTenantID(1).SetAccessKeyID("test1").SetAccessKeySecret("test1234").
		SetRoleArn("arn:aws:s3:::*").SetDurationSeconds(3600).SetIsDefault(true).SetFileSourceID(1).Exec(context.Background())
	ts.Require().NoError(err)
}

func (ts *loginFlowSuite) Test_AuthNoFlow() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	res, err := ts.AuthServer.Login(ctx, &LoginRequest{
		Password: "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", Username: "admin",
	})
	ts.Require().NoError(err)
	ts.Require().NotNil(res)
}

func (ts *loginFlowSuite) Test_AuthMFAFlow() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ts.AuthServer.cache.Set(ctx, loginFailCachePrefix+"admin", 0)
	res, err := ts.AuthServer.Login(ctx, &LoginRequest{
		Password: "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", Username: "admin",
	})
	ts.Require().NoError(err)
	ts.Contains(res.CallbackUrl, "/login/verify-factor")
	ts.NotEmptyf(res.StateToken, "state token should not be empty")
}

func Test_CreateToken(t *testing.T) {
	type jwtOpt struct {
		SigningMethod   string        `json:"signingMethod"`
		SigningKey      string        `json:"signingKey"`
		TokenTTL        time.Duration `json:"tokenTTL"`
		RefreshTokenTTL time.Duration `json:"refreshTokenTTL"`
	}
	opts := Options{
		JWT: jwtOpt{
			SigningMethod:   "HS256",
			SigningKey:      "secret",
			TokenTTL:        time.Hour * 10000,
			RefreshTokenTTL: time.Hour * 10000,
		},
	}
	_, token, err := createToken("2", opts, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}

func (ts *loginFlowSuite) Test_AuthFail() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	for i := 0; i < ts.AuthServer.CaptchaTimes-1; i++ {
		res, err := ts.AuthServer.Login(ctx, &LoginRequest{
			Password: "error", Username: "admin",
		})
		ts.Require().ErrorIs(err, status.ErrMismatchPWD)
		ts.Nil(res)
	}
	res, err := ts.AuthServer.Login(ctx, &LoginRequest{
		Password: "error", Username: "admin",
	})
	ts.Require().ErrorIs(err, status.ErrMismatchPWD)
	ts.Equal(res.CallbackUrl, "/captcha")
	for i := ts.AuthServer.CaptchaTimes; i < ts.AuthServer.LoginFailTimes; i++ {
		res, err := ts.AuthServer.Login(ctx, &LoginRequest{
			Password: "error", Username: "admin", Captcha: "",
		})
		ts.Require().ErrorIs(err, status.ErrCaptchaNotMatch)
		ts.Nil(res)
	}

	// 生成验证码
	captchaId := captcha.NewLen(6)
	digits := ts.AuthServer.captchaStore.Get(captchaId, false)
	captchaCode := ""
	for _, v := range digits {
		captchaCode = captchaCode + strconv.Itoa(int(v))
	}
	res, err = ts.AuthServer.Login(ctx, &LoginRequest{
		Password: "error", Username: "admin", Captcha: captchaCode, CaptchaId: captchaId,
	})
	// TODO 验证码准确性测试
	ts.Require().ErrorIs(err, status.ErrLoginFailUpperLimit)
	ts.Nil(res)
}

func (ts *loginFlowSuite) Test_VerifyFactor() {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	// use admin token as state token
	err := ts.AuthServer.cache.Set(context.Background(), mfaCachePrefix+adminTokenJTI, 1)
	ts.Require().NoError(err)

	ctx.Request = httptest.NewRequest("POST", "/verifyFactor", nil)
	pwd := GeneratePassCode("UWZLIIUMPX53NYXB")
	time.Sleep(1 * time.Second)
	res, err := ts.AuthServer.VerifyFactor(ctx, &VerifyFactorRequest{
		OtpToken:   pwd,
		StateToken: adminToken,
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
	err := ts.AuthServer.cache.Set(context.Background(), adminTokenJTI, "1")
	ts.Require().NoError(err)

	req := httptest.NewRequest("POST", "/logout", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	res := httptest.NewRecorder()
	ts.server.Router().ServeHTTP(res, req)

	ts.Equal(res.Code, 200)
}

func (ts *loginFlowSuite) Test_ZBindMfaFlow() {
	// 绑定mfa前置数据
	err := ts.AuthServer.cache.Set(context.Background(), adminTokenJTI, "1", cache.WithTTL(ts.AuthServer.Options.JWT.TokenTTL))
	ts.NoError(err)
	req := httptest.NewRequest("POST", "/mfa/bind-prepare", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("X-Tenant-ID", "1")
	res := httptest.NewRecorder()
	ts.server.Router().ServeHTTP(res, req)
	ts.Require().Equal(res.Code, 200)

	var bp map[string]any
	err = json.Unmarshal([]byte(res.Body.String()), &bp)
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
	err := ts.AuthServer.cache.Set(context.Background(), adminTokenJTI, "1", cache.WithTTL(ts.AuthServer.Options.JWT.TokenTTL))
	ts.NoError(err)
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
	ts.Require().NoError(ts.AuthServer.cache.Set(context.Background(),
		resetCachePrefix+adminTokenJTI, 1))
	ctx := security.WithContext(context.Background(), security.NewGenericPrincipalByClaims(jwt.MapClaims{
		"sub": "1",
	}))
	gctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	gctx.Request = httptest.NewRequest("POST", "/resetPassword", nil).WithContext(ctx)
	//ctx = context.WithValue(ctx, gin.ContextKey, gctx)
	res, err := ts.AuthServer.ResetPassword(gctx, &ResetPasswordRequest{
		NewPassword: "234567",
		StateToken:  adminToken,
	})
	ts.Require().NoError(err)
	ts.False(ts.redisServer.Exists(resetCachePrefix + adminTokenJTI))
	pwd := ts.AuthServer.db.UserPassword.Query().Where(userpassword.UserID(1),
		userpassword.SceneEQ(userpassword.SceneLogin),
	).OnlyX(entcache.Skip(context.Background()))

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

func (ts *loginFlowSuite) Test_GetOssSts() {
	err := ts.AuthServer.cache.Set(context.Background(), adminTokenJTI, "1", cache.WithTTL(ts.AuthServer.Options.JWT.TokenTTL))
	ts.NoError(err)
	payload := strings.NewReader(`{}`)
	req := httptest.NewRequest("POST", "/oss/sts", payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("X-Tenant-ID", "1")
	res := httptest.NewRecorder()
	ts.server.Router().ServeHTTP(res, req)
	ts.Require().Equal(res.Code, 200)

	var bp map[string]any
	err = json.Unmarshal([]byte(res.Body.String()), &bp)
	fmt.Print(bp)
	ts.Require().NoError(err)
}

func (ts *loginFlowSuite) Test_GetPreSignUrl() {
	err := ts.AuthServer.cache.Set(context.Background(), adminTokenJTI, "1", cache.WithTTL(ts.AuthServer.Options.JWT.TokenTTL))
	ts.NoError(err)
	payload := strings.NewReader(`{
		"url": "http://localhost:9000/test2/63f91559865b894cbe885e9a49d92a29.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=test1%2F20240618%2Fcn-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240618T085259Z&X-Amz-Expires=3600&X-Amz-SignedHeaders=host&X-Amz-Signature=cf0ac630cc5f39b106721a200ff4d178b53860d5c64d0e57fd59769b0825b035"
	}`)
	req := httptest.NewRequest("POST", "/oss/presignurl", payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("X-Tenant-ID", "1")
	res := httptest.NewRecorder()
	ts.server.Router().ServeHTTP(res, req)
	ts.Require().Equal(res.Code, 200)

	var bp map[string]any
	err = json.Unmarshal([]byte(res.Body.String()), &bp)
	fmt.Print(bp)
	ts.Require().NoError(err)
}
