package auth

import (
	"context"

	"entgo.io/ent/dialect"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/store/redisx"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler"
	entadapter "github.com/woocoos/casbin-ent-adapter"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	"github.com/woocoos/knockout/ent"
)

type ServerOption func(*Server)

func WithAuthDB(drv dialect.Driver) ServerOption {
	return func(srv *Server) {
		srv.drv = drv
	}
}

type Server struct {
	webServer *web.Server
	service   *ServerImpl
	drv       dialect.Driver
	authDb    *casbinent.Client
}

func NewServer(cnf *conf.AppConfiguration, opts ...ServerOption) (*Server, error) {
	var srv Server
	for _, opt := range opts {
		opt(&srv)
	}
	srv.service = NewServerImpl(cnf)
	// 初始化错误处理
	if err := fmterr.InitErrorHandler(cnf.Sub("errors")); err != nil {
		return nil, err
	}
	srv.service.db = ent.NewClient(ent.Driver(srv.drv))
	srv.authDb = casbinent.NewClient(casbinent.Driver(srv.drv))
	if cnf.Development {
		srv.service.db = srv.service.db.Debug()
		srv.authDb = srv.authDb.Debug()
	}
	// 初始化redis客户端
	srv.service.redisClient = buildRedis(cnf)

	if err := buildCasbin(cnf, srv.authDb); err != nil {
		return nil, err
	}

	srv.buildWebServer(cnf)

	return &srv, nil
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

func buildCasbin(cnf *conf.AppConfiguration, client *casbinent.Client) error {
	adapter, err := entadapter.NewAdapterWithClient(client)
	if err != nil {
		return err
	}
	err = casbin.SetAuthorizer(cnf.Sub("authz"), casbin.WithAdapter(adapter))
	return err
}

func buildRedis(cnf *conf.AppConfiguration) *redisx.Client {
	if cnf.IsSet("store.redis") {
		cli, err := redisx.NewClient(cnf.Sub("store.redis"))
		if err != nil {
			panic(err)
		}
		return cli
	}
	return nil
}

// Start implements woocoo.Server but do noting in start, the web server has registered by NewServer.
func (s *Server) Start(ctx context.Context) error {
	return s.webServer.Start(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	s.service.db.Close()
	return s.authDb.Close()
}
