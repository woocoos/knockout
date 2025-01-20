package graphql

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler/authz"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout-go/pkg/middleware"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/resource"
)

type ServerOptions struct {
	portalDB *ent.Client
	casbinDB *casbinent.Client
}

type Server struct {
	ServerOptions
	webSrv   *web.Server
	kosdk    *api.SDK
	resolver *Resolver
}

func NewServer(cnf *conf.AppConfiguration, opts ...ServerOption) *Server {
	s := &Server{
		ServerOptions: ServerOptions{},
	}

	for _, opt := range opts {
		opt(&s.ServerOptions)
	}

	buildCasbin(cnf, s.casbinDB)

	var err error
	s.kosdk, err = api.NewSDK(cnf.Sub("kosdk"))
	if err != nil {
		panic(err)
	}

	rs := resource.NewService(
		resource.WithClient(s.portalDB),
		resource.WithKOSDK(s.kosdk),
		resource.WithCfg(cnf))
	buildPortalHook(s.portalDB, rs)
	s.resolver = NewResolver(WithClient(s.portalDB),
		WithResource(rs))
	s.buildWebEngine(cnf)

	return s
}

func (s *Server) Start(ctx context.Context) error {
	return s.webSrv.Start(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	s.portalDB.Close()
	s.casbinDB.Close()
	return nil
}

func (s *Server) buildWebEngine(cnf *conf.AppConfiguration) {
	s.webSrv = web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		gql.RegisterMiddleware(),
		otelweb.RegisterMiddleware(),
		web.WithMiddlewareNewFunc("authz", authz.Middleware),
		middleware.RegisterTenantID(),
		middleware.RegisterTokenSigner(),
	)

	gqlSrv := handler.NewDefaultServer(NewSchema(s.resolver))
	gqlSrv.AroundResponses(middleware.SimplePagination())
	// mutation transaction
	gqlSrv.Use(entgql.Transactioner{TxOpener: s.portalDB})

	if err := gql.RegisterGraphqlServer(s.webSrv, gqlSrv); err != nil {
		panic(err)
	}
}

func buildCasbin(cnf *conf.AppConfiguration, client *casbinent.Client) {
	err := casbin.SetAuthorizer(cnf.Sub("authz"), client)
	if err != nil {
		panic(err)
	}
}

func buildPortalHook(db *ent.Client, ss *resource.Service) {
	hook := resource.NewEntHook(ss)
	db.Org.Intercept(hook.OrgTraverseFunc())
}
