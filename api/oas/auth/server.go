package auth

import (
	"context"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/ent"
)

type Server struct {
	webServer *web.Server
	service   *ServerImpl
	authDb    *casbinent.Client
}

func NewServer(app *woocoo.App) *Server {
	var srv Server

	cnf := app.AppConfiguration()
	srv.service = NewServerImpl(cnf)
	// 初始化错误处理
	if err := fmterr.InitErrorHandler(app.AppConfiguration().Sub("errors.errorCodeMap")); err != nil {
		panic(err)
	}

	ents := koapp.BuildEntComponents(cnf)
	drv := ents["portal"]
	if cnf.Development {
		srv.service.db = ent.NewClient(ent.Driver(drv), ent.Debug())
		srv.authDb = casbinent.NewClient(casbinent.Driver(drv), casbinent.Debug())
	} else {
		srv.service.db = ent.NewClient(ent.Driver(drv))
		srv.authDb = casbinent.NewClient(casbinent.Driver(drv))
	}

	buildCashbin(cnf, srv.authDb)

	srv.buildWebServer(app.AppConfiguration())

	return &srv
}

func (s *Server) buildWebServer(cnf *conf.AppConfiguration) *web.Server {
	s.webServer = web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		otelweb.RegisterMiddleware(),
	)
	// default group is '/'
	dr := s.webServer.Router().FindGroup("/").Group
	if mdl, ok := s.webServer.HandlerManager().GetMiddleware(web.GetMiddlewareKey("/", "jwt")); ok {
		s.service.LogoutHandler = mdl.(*handler.JWTMiddleware).Config.LogoutHandler
	}
	RegisterAuthHandlers(dr, s.service)
	RegisterHandlersManual(dr, s.service)
	return s.webServer
}

func buildCashbin(cnf *conf.AppConfiguration, client *casbinent.Client) {
	err := casbin.SetAuthorizer(cnf.Sub("authz"), client)
	if err != nil {
		panic(err)
	}
}

// Start implements woocoo.Server but do noting in start, the web server has registered by NewServer.
func (s *Server) Start(ctx context.Context) error {
	return s.webServer.Start(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	s.service.db.Close()
	return s.authDb.Close()
}
