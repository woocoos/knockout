package files

import (
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/knockout/api/oas/server"
	"github.com/woocoos/knockout/ent"
)

var (
	portalClient *ent.Client
)

type Server struct {
	Cnf        *conf.AppConfiguration
	RouteGroup *gin.RouterGroup
	Service    *server.FileService
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
	s.Service = &server.FileService{
		DB:       portalClient,
		BaseDir:  s.Cnf.Abs(s.Cnf.String("files.local.baseDir")),
		Endpoint: s.Cnf.Abs(s.Cnf.String("files.local.endpoint")),
	}
	return s
}

func (s *Server) BuildWebServer() *web.Server {
	webSrv := web.New(web.WithConfiguration(s.Cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(otelweb.NewMiddleware()),
	)
	return webSrv
}

func (s *Server) RegisterWebEngine(rg *gin.RouterGroup) {
	server.RegisterFileHandlers(rg, s.Service)
}
