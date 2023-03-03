package main

import (
	"ariga.io/entcache"
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/XSAM/otelsql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/pkg/store/sqlx"
	"github.com/tsingsun/woocoo/web"
	webhander "github.com/tsingsun/woocoo/web/handler"
	"github.com/tsingsun/woocoo/web/handler/authz"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/woocoos/entco/pkg/authorization"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/knockout/ent"
	_ "github.com/woocoos/knockout/ent/runtime"
	"github.com/woocoos/knockout/graph"
	"github.com/woocoos/knockout/service/resource"
	"go.opentelemetry.io/contrib/propagators/b3"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"time"
)

type entCacheLevel int

const (
	entCacheLevelContext entCacheLevel = 1
	entCacheLevelDB      entCacheLevel = 2
)

var (
	portalClient *ent.Client
	cacheLevel   entCacheLevel
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
	if err := snowflake.SetDefaultNode(app.AppConfiguration().Sub("snowflake")); err != nil {
		log.Fatal(err)
	}

	buildCashbin(app.AppConfiguration())

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
	client, err := open(conf.Global(), "store.portal")
	if err != nil {
		log.Fatal(err)
	}

	gqlSrv := handler.New(graph.NewSchema(
		graph.RegisterEntClient(client),
		graph.RegistryService(&resource.Service{Client: client}),
	))
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
	if err = gql.RegisterGraphqlServer(webSrv, gqlSrv); err != nil {
		log.Fatal(err)
	}
	// gqlserver的中间件处理
	if cacheLevel == entCacheLevelContext {
		// 启用针对http request的缓存
		useContextCache(gqlSrv)
	}
	// mutation事务
	gqlSrv.Use(entgql.Transactioner{TxOpener: client})
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

// 注意:
//  1. otelsql.AllowRoot()目前关闭,不需要记录的场景(warmup,如果开启,会导致大量的记录)
//  2. 在不需要记录的场景下,可以使用context.Background()代替
func buildEntDriver(cnf *conf.AppConfiguration, storekey string) dialect.Driver {
	var err error
	storeCfg := cnf.Sub(storekey)
	driverName := storeCfg.String("driverName")
	if cnf.IsSet("otel") {
		// Register the otelsql wrapper for the provided postgres driver.
		driverName, err = otelsql.Register("mysql",
			otelsql.WithAttributes(semconv.DBSystemMySQL),
			otelsql.WithAttributes(semconv.DBNameKey.String(storekey)),
			otelsql.WithSpanOptions(otelsql.SpanOptions{
				DisableErrSkip:  true,
				OmitRows:        true,
				OmitConnPrepare: true,
			}),
		)
		if err != nil {
			panic(err)
		}
		storeCfg.Parser().Set("driverName", driverName)
	}
	db := sqlx.NewSqlDB(storeCfg)
	drvori := entsql.OpenDB(driverName, db)
	// 如果需要设置缓存级别,可以使用entcache.ContextLevel()设置
	cacheLevel = entCacheLevel(cnf.Int("entcache.level"))
	var cacheOpts []entcache.Option
	switch cacheLevel {
	case entCacheLevelContext:
		// 使用Context缓存,但是不使用缓存的ttl
		cacheOpts = append(cacheOpts, entcache.ContextLevel())
	case entCacheLevelDB:
		// 使用db缓存,如不设置TTL.
		if entCacheTTL := cnf.Duration("entcache.ttl"); entCacheTTL > 0 {
			cacheOpts = append(cacheOpts, entcache.TTL(entCacheTTL))
		}
	default:
		return drvori // no cache
	}
	return entcache.NewDriver(drvori, cacheOpts...)
}

func open(cnf *conf.AppConfiguration, storekey string) (*ent.Client, error) {
	drv := buildEntDriver(cnf, storekey)
	var opts = []ent.Option{ent.Driver(drv)}
	if cnf.Development {
		opts = append(opts, ent.Debug())
	}
	portalClient = ent.NewClient(opts...)
	return portalClient, nil
}

func buildCashbin(cnf *conf.AppConfiguration) {
	drv := buildEntDriver(cnf, "store.portal")
	_, err := authorization.SetAuthorization(cnf.Sub("authz"), drv)
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
