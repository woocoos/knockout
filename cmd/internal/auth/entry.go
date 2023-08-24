package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/cache/redisc"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/rpc/grpcx"
	"github.com/tsingsun/woocoo/web"
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
	Cnf        *conf.AppConfiguration
	RouteGroup *gin.RouterGroup
	GrpcSrv    *grpcx.Server
	service    *server.AuthService
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

	red, err := redisc.New(s.Cnf.Sub("cache.redis"))
	if err != nil {
		log.Panic(err)
	}
	if err = red.Register(); err != nil {
		log.Panic(err)
	}
	s.service = &server.AuthService{
		DB:    portalClient,
		Cache: cache.GetCache("redis"),
	}

	if err := s.service.Apply(cnf); err != nil {
		log.Fatal(err)
	}

	return s
}

func (s *Server) RegisterWebEngine(rg *gin.RouterGroup) {
	server.RegisterAuthHandlers(rg, s.service)
	server.RegisterHandlersManual(rg, s.service)
}

func (s *Server) BuildWebServer() *web.Server {
	webSrv := web.New(web.WithConfiguration(s.Cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(otelweb.NewMiddleware()),
	)
	return webSrv
}
