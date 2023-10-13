package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/httpx"
	"github.com/tsingsun/woocoo/rpc/grpcx"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/entco/pkg/koapp"
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
	ents := koapp.BuildEntComponents(s.Cnf)
	if s.Cnf.Development {
		portalClient = ent.NewClient(ent.Driver(ents["portal"]), ent.Debug())
	} else {
		portalClient = ent.NewClient(ent.Driver(ents["portal"]))
	}

	// 初始化httpx
	cfg, err := httpx.NewClientConfig(s.Cnf.Sub("oauth-with-cache"))
	if err != nil {
		panic(err)
	}
	httpClient, err := cfg.Client(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	us := entpb.NewUserService(portalClient)
	srv := grpcx.New(grpcx.WithConfiguration(cnf.Sub("grpc")), grpcx.WithGracefulStop())
	entpb.RegisterUserServiceServer(srv.Engine(), us)
	s.GrpcSrv = srv

	s.service = &server.AuthService{
		DB:         portalClient,
		HttpClient: httpClient,
	}
	s.service.Cache, err = cache.GetCache("redis")
	if err != nil {
		panic(err)
	}
	if err := s.service.Apply(cnf); err != nil {
		panic(err)
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
		otelweb.RegisterMiddleware(),
	)
	return webSrv
}
