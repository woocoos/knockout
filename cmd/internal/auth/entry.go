package auth

import (
	"github.com/gin-contrib/cors"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/cache/redisc"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/rpc/grpcx"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/knockout/api/oas/server"
	"github.com/woocoos/knockout/api/proto/entpb"
	"github.com/woocoos/knockout/ent"
)

var (
	portalClient *ent.Client
)

type Server struct {
	Cnf     *conf.AppConfiguration
	WebSrv  *web.Server
	GrpcSrv *grpcx.Server
}

func NewServer(cnf *conf.AppConfiguration) *Server {
	s := &Server{
		Cnf: cnf,
	}
	pd := oteldriver.BuildOTELDriver(s.Cnf, "store.portal")
	pd = ecx.BuildEntCacheDriver(s.Cnf, pd)
	if s.Cnf.Development {
		portalClient = ent.NewClient(ent.Driver(pd), ent.Debug())
	} else {
		portalClient = ent.NewClient(ent.Driver(pd))
	}

	us := entpb.NewUserService(portalClient)
	srv := grpcx.New(grpcx.WithConfiguration(cnf.Sub("grpc")), grpcx.WithGracefulStop())
	entpb.RegisterUserServiceServer(srv.Engine(), us)
	s.GrpcSrv = srv

	if err := redisc.New(s.Cnf.Sub("cache.redis")).Register(); err != nil {
		log.Panic(err)
	}
	webservice := &server.Service{
		DB:    portalClient,
		Cache: cache.GetCache("redis"),
	}

	if err := webservice.Apply(cnf); err != nil {
		log.Fatal(err)
	}

	s.WebSrv = buildWebServer(s.Cnf, webservice)

	return s
}

func buildWebServer(cnf *conf.AppConfiguration, ws *server.Service) *web.Server {
	webSrv := web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(otelweb.NewMiddleware()),
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
