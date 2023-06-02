package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/pkg/conf"
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
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/knockout/api/graphql"
	"github.com/woocoos/knockout/api/graphql/generated"
	"github.com/woocoos/knockout/cmd/internal/otel"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/resource"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout/ent/runtime"
)

var (
	portalClient *ent.Client
	casbinClient *casbinent.Client
)

func main() {
	app := woocoo.New()
	otelStop := otel.Apply(app.AppConfiguration())

	err := snowflake.SetDefaultNode(app.AppConfiguration().Sub("snowflake"))
	if err != nil {
		log.Fatal(err)
	}

	pd := oteldriver.BuildOTELDriver(app.AppConfiguration(), "store.portal")
	pd = ecx.BuildEntCacheDriver(app.AppConfiguration(), pd)
	if app.AppConfiguration().Development {
		portalClient = ent.NewClient(ent.Driver(pd), ent.Debug())
		casbinClient = casbinent.NewClient(casbinent.Driver(pd), casbinent.Debug())
	} else {
		portalClient = ent.NewClient(ent.Driver(pd))
		casbinClient = casbinent.NewClient(casbinent.Driver(pd))
	}

	buildCashbin(app.AppConfiguration(), casbinClient)

	webSrv := buildWebServer(app.AppConfiguration())
	app.RegisterServer(webSrv)

	defer func() {
		portalClient.Close()
		casbinClient.Close()
		otelStop()
	}()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func buildWebServer(cnf *conf.AppConfiguration) *web.Server {
	webSrv := web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(gql.New()),
		//web.RegisterMiddleware(otelweb.NewMiddleware()),
		web.RegisterMiddleware(authz.New()),
		identity.RegistryTenantIDMiddleware(),
	)

	gqlSrv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql.Resolver{
			Client:   portalClient,
			Resource: &resource.Service{Client: portalClient},
		},
	}))
	// gqlserver的中间件处理
	if cnf.String("entcache.level") == "context" {
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
	// 关闭缓存
	//authorizer.Enforcer.(*casbin.CachedEnforcer).EnableCache(false)
	authorizer.RequestParser = func(ctx context.Context, id security.Identity, item *security.PermissionItem) []any {
		gctx := ctx.Value(gin.ContextKey).(*gin.Context)
		domain := gctx.GetHeader(identity.TenantHeaderKey)
		p := item.AppCode + ":" + item.Action
		return []any{id.Name(), domain, p, "read"}
	}
	if err != nil {
		log.Fatal(err)
	}
}
