package main

import (
	"ariga.io/entcache"
	"context"
	"entgo.io/contrib/entgql"
	gqlgen "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/tsingsun/woocoo/web"
	webhander "github.com/tsingsun/woocoo/web/handler"
	"github.com/tsingsun/woocoo/web/handler/authz"
	"github.com/vektah/gqlparser/v2/ast"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/ecx/oteldriver"
	"github.com/woocoos/entco/pkg/authorization"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/knockout/api/graphql"
	"github.com/woocoos/knockout/api/graphql/generated"
	"github.com/woocoos/knockout/cmd/internal/otel"
	"github.com/woocoos/knockout/ent"
	_ "github.com/woocoos/knockout/ent/runtime"
	"github.com/woocoos/knockout/service/resource"
	"time"
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

	gqlSrv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graphql.Resolver{
			Client:   portalClient,
			Resource: &resource.Service{Client: portalClient},
		},
	}))
	gqlSrv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	gqlSrv.AddTransport(transport.Options{})
	gqlSrv.AddTransport(transport.GET{})
	gqlSrv.AddTransport(transport.POST{})
	gqlSrv.AddTransport(transport.MultipartForm{})
	gqlSrv.SetQueryCache(lru.New(1000))

	gqlSrv.Use(extension.Introspection{})
	gqlSrv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	if err := gql.RegisterGraphqlServer(webSrv, gqlSrv); err != nil {
		log.Fatal(err)
	}
	// gqlserver的中间件处理
	if cnf.String("entcache.level") == "context" {
		// 启用针对http request的缓存
		useContextCache(gqlSrv)
		useSimplePagination(gqlSrv)
	}
	// mutation事务
	gqlSrv.Use(entgql.Transactioner{TxOpener: portalClient})
	// 订阅支持,如果握手阶段开发模式可以允许.
	if jwt, ok := webSrv.HandlerManager().Get("jwt"); ok {
		if h, ok := jwt.(*webhander.JWTMiddleware); ok {
			h.Config.ErrorHandler = func(c *gin.Context, err error) error {
				if c.IsWebsocket() && cnf.Development {
					return nil
				}
				return err
			}
		}
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

func useContextCache(server *handler.Server) {
	server.AroundResponses(func(ctx context.Context, next gqlgen.ResponseHandler) *gqlgen.Response {
		if op := gqlgen.GetOperationContext(ctx).Operation; op != nil && op.Operation == ast.Query {
			ctx = entcache.NewContext(ctx)
		}
		return next(ctx)
	})
}

func useSimplePagination(server *handler.Server) {
	server.AroundResponses(func(ctx context.Context, next gqlgen.ResponseHandler) *gqlgen.Response {
		gctx, _ := gql.FromIncomingContext(ctx)
		if gctx != nil {
			sp, err := ent.NewSimplePagination(gctx.Query("p"), gctx.Query("c"))
			if err != nil {
				return gqlgen.ErrorResponse(ctx, "pagination error:%v", err)
			}
			if sp != nil {
				ctx = ent.WithSimplePagination(ctx, sp)
			}
		}
		return next(ctx)
	})
}
