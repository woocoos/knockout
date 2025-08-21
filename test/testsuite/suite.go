package testsuite

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/alicebob/miniredis/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/security"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	ecx "github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/migrate"
	"github.com/woocoos/knockout/test"
	"os"
	"path/filepath"
	"strconv"
)

type BaseSuite struct {
	suite.Suite
	Cnf             *conf.AppConfiguration
	DSN, DriverName string
	Client          *ent.Client
	CacheClient     *ent.Client
	Redis           *miniredis.Miniredis
	App             *woocoo.App
	AuthDbClient    *casbinent.Client
}

func (o *BaseSuite) Setup() error {
	o.App = initTestApp()
	o.Cnf = o.App.AppConfiguration()
	o.Redis = initMiniRedis(o.Cnf)

	if o.DSN == "" && o.DriverName == "" {
		o.DriverName = "sqlite3"
		o.DSN = "file:portalLite?mode=memory&cache=shared&_fk=1"
	}
	o.Cnf.Parser().Set("store.portal.driverName", o.DriverName)
	o.Cnf.Parser().Set("store.portal.dsn", o.DSN)

	drv, err := sql.Open(o.DriverName, o.DSN)
	o.Require().NoError(err)
	o.Client = ent.NewClient(ent.Driver(drv))
	o.AuthDbClient = casbinent.NewClient(casbinent.Driver(drv))

	td, _ := ecx.BuildEntCacheDriver(o.Cnf.Sub("entcache"), drv)
	o.CacheClient = ent.NewClient(ent.Driver(td)).Debug()

	err = o.Client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false))
	o.Require().NoError(err)

	err = o.AuthDbClient.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false))
	o.Require().NoError(err)
	return nil
}

func initTestApp() *woocoo.App {
	file := filepath.Join(test.BaseDir(), "testdata", "etc", "app.yaml")
	bs, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	app := woocoo.New(woocoo.WithAppConfiguration(
		conf.NewFromBytes(bs, conf.WithBaseDir(test.BaseDir())).AsGlobal()),
	)
	return app
}

func initMiniRedis(cnf *conf.AppConfiguration) *miniredis.Miniredis {
	db, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	cnf.Parser().Set("store.redis.addrs", []string{db.Addr()})
	cnf.Parser().Set("authz.watcherOptions.options.addr", db.Addr())

	koapp.BuildCacheComponents(cnf)
	return db
}

func (o *BaseSuite) BearToken() string {
	const Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIiwiZXhwIjoyMDcxMjIwMzA4LCJpYXQiOjE3NTU2ODAzMDgsImp0aSI6InRva2VuOjE6OWFmNGVhNTktN2EzMS00NjU4LThhMmUtNGIwZDRmODQ5Y2RhIn0.u1anRaDXmMBRjBNLJawLZ4RM98HjdWKodXuAsp3-DJ8"
	return Token
}

func (o *BaseSuite) NewTestCtx(uid, oid int) context.Context {
	return NewTestCtx(uid, oid, o.Client)
}

func NewTestCtx(uid, oid int, client *ent.Client) context.Context {
	ctx := ent.NewContext(context.Background(), client)
	// with identity
	ctx = security.WithContext(ctx, security.NewGenericPrincipalByClaims(jwt.MapClaims{"sub": strconv.Itoa(uid)}))
	if oid != 0 {
		ctx = identity.WithTenantID(ctx, oid)
	}
	return ctx
}
