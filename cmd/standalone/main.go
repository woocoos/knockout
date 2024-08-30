package main

import (
	"flag"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/telemetry"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/api/graphql"
	"github.com/woocoos/knockout/api/oas/auth"
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
	rmsSvr := graphql.NewServer(app)

	authcnf := &conf.AppConfiguration{
		Configuration: conf.New(conf.WithBaseDir(*authConfig), conf.WithGlobal(false)).Load(),
	}
	koapp.BuildCacheComponents(authcnf)
	app.AppConfiguration().Configuration = authcnf.Configuration
	authSrv := auth.NewServer(app)

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
