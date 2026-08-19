package job

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout-go/pkg/koapp"
	schemahook "github.com/woocoos/knockout/codegen/entgen/hook"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/enttest"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/userpassword"
	"testing"
	"time"
)

type passwordExpiredJobSuite struct {
	suite.Suite
	db     *ent.Client
	server *Server
	pej    *PasswordExpiredJob
}

func TestPasswordExpiredJobSuite(t *testing.T) {
	suite.Run(t, new(passwordExpiredJobSuite))
}

func (t *passwordExpiredJobSuite) SetupSuite() {
	app := koapp.New(woocoo.WithAppConfiguration(conf.New(conf.WithBaseDir("../../test/testdata"), conf.WithGlobal(true)).Load()))
	var err error
	t.server, err = NewServer(app.AppConfiguration().Sub("job"))
	t.Require().NoError(err)
	t.db = enttest.Open(t.T(), "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1",
		enttest.WithOptions(ent.Debug()), enttest.WithMigrateOptions(schema.WithForeignKeys(false)))
	schemahook.RegisterAllHooks(t.db)
	t.initDbData(t.newTestCtx(t.db), t.db)
	// 初始化job
	var pej *PasswordExpiredJob
	pej, err = NewPasswordExpiredJob(app.AppConfiguration().Sub("job.password"))
	if err != nil {
		return
	}
	pej.db = t.db
	pej.kosdk = nil
	t.pej = pej
	go func() {
		if err := app.Run(); err != nil {
			panic(err)
		}
	}()
}

func (t *passwordExpiredJobSuite) initDbData(ctx context.Context, client *ent.Client) {
	// 用户
	usrs := make([]*ent.UserCreate, 0)
	usr := client.User.Create().SetID(1).SetUserType(user.UserTypeAccount).SetCreationType(user.CreationTypeManual).
		SetRegisterIP("").SetPrincipalName("user").SetDisplayName("user").
		SetStatus(types.UserStatusActive).SetCreatedBy(1)
	usrs = append(usrs, usr)
	client.User.CreateBulk(usrs...).ExecX(ctx)
	// 密码
	usrPwds := make([]*ent.UserPasswordCreate, 0)
	pwd := client.UserPassword.Create().SetID(1).SetUserID(1).SetCreatedBy(1).SetScene(userpassword.SceneLogin).
		SetStatus(typex.SimpleStatusActive).SetPassword("123456").SetSalt("123456").
		SetPassword("9b1063951d443cfac15cc879efb4054f4f4fd599e1b1a9aee67b0301e19e40fe")
	usrPwds = append(usrPwds, pwd)
	client.UserPassword.CreateBulk(usrPwds...).ExecX(ctx)
	// 组织
	orgs := make([]*ent.OrgCreate, 0)
	or := client.Org.Create().SetID(1).SetKind(org.KindRoot).SetParentID(0).SetStatus(typex.SimpleStatusActive).
		SetCreatedBy(1).SetUpdatedBy(1).SetName("org")
	orgs = append(orgs, or)
	client.Org.CreateBulk(orgs...).ExecX(ctx)
	// 组织用户
	orgUsers := make([]*ent.OrgUserCreate, 0)
	orgUser := client.OrgUser.Create().SetOrgID(1).SetUserID(1).SetCreatedBy(1).SetDisplayName("user")
	orgUsers = append(orgUsers, orgUser)
	client.OrgUser.CreateBulk(orgUsers...).ExecX(ctx)
}

func (t *passwordExpiredJobSuite) newTestCtx(db *ent.Client) context.Context {
	ctx := ent.NewContext(context.Background(), db)
	// with identity
	ctx = security.WithContext(ctx, security.NewGenericPrincipalByClaims(jwt.MapClaims{"sub": "1"}))
	ctx = identity.WithTenantID(ctx, 1000)
	return ctx
}

func (t *passwordExpiredJobSuite) TestPwdExpired() {
	ctx := t.newTestCtx(t.db)
	checkTime, err := time.Parse("2006-01-02 15:04:05", "2025-12-18 12:01:20")
	t.NoError(err)
	ups := make([]*ent.UserPassword, 0)
	up, err := t.db.UserPassword.Get(ctx, 1)
	t.NoError(err)
	// 已过期
	pwdTime, err := time.Parse("2006-01-02 15:04:05", "2024-12-18 12:00:20")
	t.NoError(err)
	up.UpdatedAt = pwdTime
	ups = append(ups, up)
	t.pej.checkPwd(ctx, ups, checkTime)
}

func (t *passwordExpiredJobSuite) TestPwdExpiring() {
	ctx := t.newTestCtx(t.db)
	checkTime, err := time.Parse("2006-01-02 15:04:05", "2025-12-18 12:01:20")
	t.NoError(err)
	ups := make([]*ent.UserPassword, 0)
	up, err := t.db.UserPassword.Get(ctx, 1)
	t.NoError(err)
	// 差一天过期
	pwdTime, err := time.Parse("2006-01-02 15:04:05", "2024-12-19 12:00:20")
	t.NoError(err)
	up.UpdatedAt = pwdTime
	ups = append(ups, up)
	t.pej.checkPwd(ctx, ups, checkTime)
}

func (t *passwordExpiredJobSuite) TestPwdRemind() {
	ctx := t.newTestCtx(t.db)
	checkTime, err := time.Parse("2006-01-02 15:04:05", "2025-12-18 12:01:20")
	t.NoError(err)
	ups := make([]*ent.UserPassword, 0)
	up, err := t.db.UserPassword.Get(ctx, 1)
	t.NoError(err)
	// 未过期常规提醒
	pwdTime, err := time.Parse("2006-01-02 15:04:05", "2025-03-18 12:01:20")
	t.NoError(err)
	up.UpdatedAt = pwdTime
	ups = append(ups, up)
	t.pej.checkPwd(ctx, ups, checkTime)
}
