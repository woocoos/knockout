package graphql

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler/authz"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout-go/pkg/middleware"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/resource"
)

type Server struct {
	portalClient *ent.Client
	casbinClient *casbinent.Client
	webSrv       *web.Server
	kosdk        *api.SDK
}

func NewServer(app *woocoo.App) *Server {
	s := &Server{}
	cnf := app.AppConfiguration()
	ents := koapp.BuildEntComponents(cnf)
	drv := ents["portal"]
	if cnf.Development {
		s.portalClient = ent.NewClient(ent.Driver(drv), ent.Debug())
		s.casbinClient = casbinent.NewClient(casbinent.Driver(drv), casbinent.Debug())
	} else {
		s.portalClient = ent.NewClient(ent.Driver(drv))
		s.casbinClient = casbinent.NewClient(casbinent.Driver(drv))
	}
	buildCashbin(cnf, s.casbinClient)

	s.buildWebEngine(app)

	app.RegisterServer(s.webSrv)

	var err error
	s.kosdk, err = api.NewSDK(cnf.Sub("kosdk"))
	if err != nil {
		panic(err)
	}
	return s
}

func (s *Server) Start(ctx context.Context) error {
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.portalClient.Close()
	s.casbinClient.Close()
	return nil
}

func (s *Server) buildWebEngine(app *woocoo.App) {
	cnf := app.AppConfiguration()
	s.webSrv = web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		gql.RegisterMiddleware(),
		otelweb.RegisterMiddleware(),
		web.WithMiddlewareNewFunc("authz", authz.Middleware),
		middleware.RegisterTenantID(),
		middleware.RegisterTokenSigner(),
	)

	gqlSrv := handler.NewDefaultServer(NewSchema(WithClient(s.portalClient),
		WithResource(&resource.Service{Client: s.portalClient, KOSDK: s.kosdk}),
	))
	gqlSrv.AroundResponses(middleware.SimplePagination())
	// mutation transaction
	gqlSrv.Use(entgql.Transactioner{TxOpener: s.portalClient})

	if err := gql.RegisterGraphqlServer(s.webSrv, gqlSrv); err != nil {
		panic(err)
	}
}

func buildCashbin(cnf *conf.AppConfiguration, client *casbinent.Client) {
	err := casbin.SetAuthorizer(cnf.Sub("authz"), client)
	if err != nil {
		panic(err)
	}
}
