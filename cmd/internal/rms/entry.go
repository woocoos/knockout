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
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/gqlx"
	"github.com/woocoos/entco/pkg/authorization"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/knockout/api/graphql"
	"github.com/woocoos/knockout/api/graphql/generated"
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
	pd := oteldriver.BuildOTELDriver(s.Cnf, "store.portal")
	pd = ecx.BuildEntCacheDriver(s.Cnf, pd)
	if s.Cnf.Development {
		portalClient = ent.NewClient(ent.Driver(pd), ent.Debug())
		casbinClient = casbinent.NewClient(casbinent.Driver(pd), casbinent.Debug())
	} else {
		portalClient = ent.NewClient(ent.Driver(pd))
		casbinClient = casbinent.NewClient(casbinent.Driver(pd))
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
		web.RegisterMiddleware(gql.New()),
		web.RegisterMiddleware(otelweb.NewMiddleware()),
		web.RegisterMiddleware(authz.New()),
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

	fileOptions := resource.FileOptions{}
	err = s.Cnf.Sub("files").Unmarshal(&fileOptions)
	if err != nil {
		log.Panic(err)
	}
	gqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql.Resolver{
			Client:   portalClient,
			Resource: &resource.Service{Client: portalClient, HttpClient: httpClient, FileOptions: fileOptions},
		},
	}))
	// gqlserver的中间件处理
	if s.Cnf.String("entcache.level") == "context" {
		gqlSrv.AroundResponses(gqlx.ContextCache())
	}
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
