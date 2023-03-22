package main

import (
	"entgo.io/ent/dialect"
	"github.com/gin-contrib/cors"
	"github.com/go-redis/redis/v8"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/telemetry"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/rpc/grpcx"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/pkg/authorization"
	"github.com/woocoos/knockout/api/oas/server"
	"github.com/woocoos/knockout/api/proto/entpb"
	"github.com/woocoos/knockout/ent"
	"go.opentelemetry.io/contrib/propagators/b3"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := woocoo.New()
	if app.AppConfiguration().IsSet("otel") {
		otelCnf := app.AppConfiguration().Sub("otel")
		otelcfg := telemetry.NewConfig(otelCnf,
			telemetry.WithPropagator(b3.New()),
		)
		defer otelcfg.Shutdown()
	}
	pd := oteldriver.BuildOTELDriver(app.AppConfiguration(), "store.portal")
	pd = ecx.BuildEntCacheDriver(app.AppConfiguration(), pd)
	var portalClient *ent.Client

	if app.AppConfiguration().Development {
		portalClient = ent.NewClient(ent.Driver(pd), ent.Debug())
	} else {
		portalClient = ent.NewClient(ent.Driver(pd))
	}

	buildCashbin(app.AppConfiguration(), pd)

	us := entpb.NewUserService(portalClient)
	srv := grpcx.New(grpcx.WithGracefulStop())
	entpb.RegisterUserServiceServer(srv.Engine(), us)

	webservice := &server.Service{
		DB: portalClient,
		Redis: redis.NewClient(&redis.Options{
			Addr: app.AppConfiguration().String("store.redis.addr"),
		}),
	}
	server.Cnf = app.AppConfiguration()
	websrv := buildWebServer(app.AppConfiguration(), webservice)

	app.RegisterServer(websrv, srv)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func buildWebServer(cnf *conf.AppConfiguration, ws *server.Service) *web.Server {
	webSrv := web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(otelweb.NewMiddleware()),
	)
	router := webSrv.Router()
	router.Use(cors.Default())
	server.RegisterHandlers(&webSrv.Router().RouterGroup, ws)
	return webSrv
}

func buildCashbin(cnf *conf.AppConfiguration, driver dialect.Driver) {
	_, err := authorization.SetAuthorization(cnf.Sub("authz"), driver)
	if err != nil {
		log.Fatal(err)
	}
}
