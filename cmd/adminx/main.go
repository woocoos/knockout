package main

import (
	"ariga.io/entcache"
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/web"
	webhander "github.com/tsingsun/woocoo/web/handler"
	"github.com/tsingsun/woocoo/web/handler/authz"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/woocoos/entco/initx"
	"github.com/woocoos/entco/initx/oteldriver"
	"github.com/woocoos/entco/pkg/authorization"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/graph"
	"github.com/woocoos/knockout/graph/generated"
	"github.com/woocoos/knockout/service/resource"
	"go.opentelemetry.io/contrib/propagators/b3"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout/ent/runtime"
)

var (
	portalClient *ent.Client
)

func main() {
	app := woocoo.New()
	if app.AppConfiguration().IsSet("otel") {
		otelCnf := app.AppConfiguration().Sub("otel")
		otelcfg := telemetry.NewConfig(otelCnf,
			telemetry.WithPropagator(b3.New()),
		)
		defer otelcfg.Shutdown()
	}

	err := snowflake.SetDefaultNode(app.AppConfiguration().Sub("snowflake"))
	if err != nil {
		log.Fatal(err)
	}

	pd := oteldriver.BuildOTELDriver(app.AppConfiguration(), "store.portal")
	pd = initx.BuildEntCacheDriver(app.AppConfiguration(), pd)
	if app.AppConfiguration().Development {
		portalClient = ent.NewClient(ent.Driver(pd), ent.Debug())
	} else {
		portalClient = ent.NewClient(ent.Driver(pd))
	}

	buildCashbin(app.AppConfiguration(), pd)

	webSrv := buildWebServer(app.AppConfiguration())
	app.RegisterServer(webSrv)

	defer func() {
		portalClient.Close()
	}()
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func buildWebServer(cnf *conf.AppConfiguration) *web.Server {
	webSrv := web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		web.RegisterMiddleware(gql.New()),
		web.RegisterMiddleware(otelweb.NewMiddleware()),
		web.RegisterMiddleware(authz.New()),
	)
	webSrv.Router().NoRoute()

	gqlSrv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
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
	// gqlserver??????????????????
	if cnf.String("entcache.level") == "context" {
		// ????????????http request?????????
		useContextCache(gqlSrv)
	}
	// mutation??????
	gqlSrv.Use(entgql.Transactioner{TxOpener: portalClient})
	// ????????????,??????????????????????????????????????????.
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

func buildCashbin(cnf *conf.AppConfiguration, driver dialect.Driver) {
	_, err := authorization.SetAuthorization(cnf.Sub("authz"), driver)
	if err != nil {
		log.Fatal(err)
	}
}

func useContextCache(server *handler.Server) {
	server.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		if op := graphql.GetOperationContext(ctx).Operation; op != nil && op.Operation == ast.Query {
			ctx = entcache.NewContext(ctx)
		}
		return next(ctx)
	})
}
