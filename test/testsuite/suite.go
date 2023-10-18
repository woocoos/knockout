package testsuite

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/alicebob/miniredis/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/security"
	entadapter "github.com/woocoos/casbin-ent-adapter"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/pkg/identity"
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
	redis           *miniredis.Miniredis
}

func (o *BaseSuite) Setup() error {
	app := initTestApp()
	o.Cnf = app.AppConfiguration()
	o.redis = initMiniRedis(o.Cnf)

	if o.DSN == "" && o.DriverName == "" {
		o.DriverName = "sqlite3"
		o.DSN = "file:msgcenter?mode=memory&cache=shared&_fk=1"
	}
	client, err := open(context.Background(), o.DriverName, o.DSN)
	o.Require().NoError(err)
	o.Client = client.Debug()

	drv, err := sql.Open(o.DriverName, o.DSN)
	o.Require().NoError(err)
	td, _ := ecx.BuildEntCacheDriver(o.Cnf.Sub("entcache"), drv)
	o.CacheClient = ent.NewClient(ent.Driver(td)).Debug()
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
	return db
}

// some as scripts/intidb.go as possible
func open(ctx context.Context, driverName, dsn string) (*ent.Client, error) {
	client, err := ent.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
	); err != nil {
		return nil, err
	}

	_, err = entadapter.NewAdapter(driverName, dsn, entadapter.WithMigration())
	return client, nil
}

func (o *BaseSuite) NewTestCtx(uid, oid int) context.Context {
	ctx := ent.NewContext(context.Background(), o.Client)
	// with identity
	ctx = security.WithContext(ctx, security.NewGenericPrincipalByClaims(jwt.MapClaims{"sub": strconv.Itoa(uid)}))
	if oid != 0 {
		ctx = identity.WithTenantID(ctx, oid)
	}
	return ctx
}
