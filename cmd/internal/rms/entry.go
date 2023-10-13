package rms

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/httpx"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler/authz"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/entco/gqlx"
	"github.com/woocoos/entco/pkg/authorization"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/pkg/koapp"
	"github.com/woocoos/knockout/api/graphql"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/resource"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout/ent/runtime"
)

var (
	portalClient *ent.Client
	casbinClient *casbinent.Client
)

type Server struct {
	Cnf    *conf.AppConfiguration
	webSrv *web.Server
}

func NewServer(cnf *conf.AppConfiguration) *Server {
	s := &Server{
		Cnf: cnf,
	}
	ents := koapp.BuildEntComponents(s.Cnf)
	drv := ents["portal"]
	if s.Cnf.Development {
		portalClient = ent.NewClient(ent.Driver(drv), ent.Debug())
		casbinClient = casbinent.NewClient(casbinent.Driver(drv), casbinent.Debug())
	} else {
		portalClient = ent.NewClient(ent.Driver(drv))
		casbinClient = casbinent.NewClient(casbinent.Driver(drv))
	}
	buildCashbin(s.Cnf, casbinClient)

	return s
}

func (s *Server) Start(ctx context.Context) error {
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	portalClient.Close()
	casbinClient.Close()
	return nil
}

func (s *Server) BuildWebEngine() *web.Server {
	webSrv := web.New(web.WithConfiguration(s.Cnf.Sub("web")),
		web.WithGracefulStop(),
		gql.RegistryMiddleware(),
		otelweb.RegisterMiddleware(),
		web.WithMiddlewareNewFunc("authz", authz.Middleware),
		identity.RegistryTenantIDMiddleware(),
	)

	// 初始化httpx
	cfg, err := httpx.NewClientConfig(s.Cnf.Sub("oauth-with-cache"))
	if err != nil {
		log.Panic(err)
	}
	httpClient, err := cfg.Client(context.Background(), nil)
	if err != nil {
		log.Panic(err)
	}

	oasOptions := resource.OASOptions{}
	err = s.Cnf.Sub("oas").Unmarshal(&oasOptions)
	if err != nil {
		log.Panic(err)
	}
	gqlSrv := handler.NewDefaultServer(graphql.NewSchema(graphql.WithClient(portalClient),
		graphql.WithResource(&resource.Service{Client: portalClient, HttpClient: httpClient, OASOptions: oasOptions}),
	))
	gqlSrv.AroundResponses(gqlx.SimplePagination())
	// mutation事务
	gqlSrv.Use(entgql.Transactioner{TxOpener: portalClient})

	if err := gql.RegisterGraphqlServer(webSrv, gqlSrv); err != nil {
		log.Fatal(err)
	}

	return webSrv
}

func buildCashbin(cnf *conf.AppConfiguration, client *casbinent.Client) {
	authorizer, err := authorization.SetAuthorization(cnf.Sub("authz"), client)
	if err != nil {
		log.Panic(err)
	}
	// 关闭缓存
	// authorizer.Enforcer.(*casbin.CachedEnforcer).EnableCache(false)
	authorizer.RequestParser = func(ctx context.Context, id security.Identity, item *security.PermissionItem) []any {
		gctx := ctx.Value(gin.ContextKey).(*gin.Context)
		domain := gctx.GetHeader(identity.TenantHeaderKey)
		p := item.AppCode + ":" + item.Action
		return []any{id.Name(), domain, p, "read"}
	}
	if err != nil {
		log.Panic(err)
	}
}
