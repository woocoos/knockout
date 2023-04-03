package main

import (
	"entgo.io/ent/dialect"
	"github.com/gin-contrib/cors"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/cache/redisc"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/rpc/grpcx"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/pkg/authorization"
	"github.com/woocoos/knockout/api/oas/server"
	"github.com/woocoos/knockout/api/proto/entpb"
	"github.com/woocoos/knockout/ent"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout/ent/runtime"
)

func main() {
	app := woocoo.New()
	//if app.AppConfiguration().IsSet("otel") {
	//	otelCnf := app.AppConfiguration().Sub("otel")
	//	otelcfg := telemetry.NewConfig(otelCnf,
	//		telemetry.WithPropagator(b3.New()),
	//	)
	//	defer otelcfg.Shutdown()
	//}
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

	if err := redisc.NewBuiltIn().Register(); err != nil {
		log.Fatal(err)
	}
	webservice := &server.Service{
		DB:    portalClient,
		Cache: cache.GetCache("redis"),
	}
	server.Cnf = app.AppConfiguration()
	if err := webservice.Apply(app.AppConfiguration()); err != nil {
		log.Fatal(err)
	}
	websrv := buildWebServer(app.AppConfiguration(), webservice)

	app.RegisterServer(websrv, srv)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func buildWebServer(cnf *conf.AppConfiguration, ws *server.Service) *web.Server {
	webSrv := web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		//web.RegisterMiddleware(otelweb.NewMiddleware()),
	)
	router := webSrv.Router()
	router.Use(cors.Default())
	if mdl, ok := webSrv.HandlerManager().Get("jwt"); ok {
		ws.LogoutHandler = mdl.(*handler.JWTMiddleware).Config.LogoutHandler
	}
	server.RegisterHandlers(&webSrv.Router().RouterGroup, ws)
	server.RegisterHandlersManual(webSrv.Router().Engine, ws)
	return webSrv
}

func buildCashbin(cnf *conf.AppConfiguration, driver dialect.Driver) {
	_, err := authorization.SetAuthorization(cnf.Sub("authz"), driver)
	if err != nil {
		log.Fatal(err)
	}
}
