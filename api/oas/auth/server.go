package auth

import (
	"context"

	"entgo.io/ent/dialect"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	"github.com/woocoos/knockout/ent"
)

type ServerOption func(*Server)

func WithAuthDB(drv dialect.Driver) ServerOption {
	return func(srv *Server) {
		srv.drv = drv
	}
}

func WithCache(c cache.Cache) ServerOption {
	return func(srv *Server) {
		srv.cache = c
	}
}

type Server struct {
	webServer *web.Server
	service   *ServerImpl
	drv       dialect.Driver
	cache     cache.Cache
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
	if cnf.Development {
		srv.service.db = srv.service.db.Debug()
	}
	// 初始化 cache 组件: 优先使用传入的 cache, 否则从配置获取默认
	if srv.cache != nil {
		srv.service.cache = srv.cache
	} else {
		cacheCnf := cnf.Sub("cache")
		if driverName := cacheCnf.String("default"); driverName != "" {
			if c, err := cache.GetCache(driverName); err == nil {
				srv.service.cache = c
			}
		} else {
			var (
				c     cache.Cache
				count int
			)
			cnf.Map("cache", func(root string, sub *conf.Configuration) {
				if root == "default" {
					return
				}
				count++
				if c == nil {
					driverName := sub.String("driverName")
					if driverName == "" {
						driverName = root
					}
					if inst, err := cache.GetCache(driverName); err == nil {
						c = inst
					}
				}
			})
			if count == 1 && c != nil {
				srv.service.cache = c
			}
		}
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

// Start implements woocoo.Server but do noting in start, the web server has registered by NewServer.
func (s *Server) Start(ctx context.Context) error {
	return s.webServer.Start(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.service.db.Close()
}
