package main

import (
	"flag"

	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/telemetry"
	"github.com/tsingsun/woocoo/pkg/conf"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/koapp"
	schemahook "github.com/woocoos/knockout/codegen/entgen/hook"
	"github.com/woocoos/knockout/api/graphql"
	"github.com/woocoos/knockout/api/oas/auth"
	"github.com/woocoos/knockout/ent"
	"go.opentelemetry.io/contrib/propagators/b3"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout-go/pkg/snowflake"
	_ "github.com/woocoos/knockout/ent/runtime"
)

var (
	rmcConfig  = flag.String("r", "../adminx", "rms etc dir")
	authConfig = flag.String("a", "../auth", "auth etc dir")
)

// Notice: reuse the app instance to initial servers,
// please guarantee the configuration of app no using after NewServer.
func main() {
	flag.Parse()
	app := woocoo.New()

	rmscnf := &conf.AppConfiguration{
		Configuration: conf.New(conf.WithBaseDir(*rmcConfig), conf.WithGlobal(false)).Load(),
	}
	otelStop := applyOTEL(rmscnf)
	defer otelStop()

	koapp.BuildCacheComponents(rmscnf)
	app.AppConfiguration().Configuration = rmscnf.Configuration
	ents := koapp.BuildEntComponents(app.AppConfiguration())
	drv := ents["portal"]
	portalClient := ent.NewClient(ent.Driver(drv))
	schemahook.RegisterAllHooks(portalClient)
	casbinClient := casbinent.NewClient(casbinent.Driver(drv))
	if app.AppConfiguration().Development {
		portalClient = portalClient.Debug()
		casbinClient = casbinClient.Debug()
	}
	rmsSvr := graphql.NewServer(rmscnf, graphql.WithCasbinDB(casbinClient), graphql.WithPortalDB(portalClient))

	authcnf := &conf.AppConfiguration{
		Configuration: conf.New(conf.WithBaseDir(*authConfig), conf.WithGlobal(false)).Load(),
	}
	koapp.BuildCacheComponents(authcnf)
	app.AppConfiguration().Configuration = authcnf.Configuration
	authSrv, err := auth.NewServer(app.AppConfiguration(), auth.WithAuthDB(drv))
	if err != nil {
		panic(err)
	}
	app.RegisterServer(rmsSvr, authSrv, clientx.ChangeSet)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

// Apply 尝试注册otel,如果配置中有otel配置,则注册.并返回关闭函数
func applyOTEL(cnf *conf.AppConfiguration) func() {
	if cnf.IsSet("otel") {
		otelCnf := cnf.Sub("otel")
		otelcfg := telemetry.NewConfig(otelCnf,
			telemetry.WithPropagator(b3.New()),
		)
		return func() {
			otelcfg.Shutdown()
		}
	}
	return func() {}
}
