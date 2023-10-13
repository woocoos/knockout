package main

import (
	"flag"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/contrib/telemetry"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/pkg/koapp"
	"github.com/woocoos/knockout/cmd/internal/auth"
	"github.com/woocoos/knockout/cmd/internal/files"
	"github.com/woocoos/knockout/cmd/internal/rms"
	"go.opentelemetry.io/contrib/propagators/b3"
)

var (
	rmcConfig  = flag.String("r", "../adminx", "rms etc dir")
	authConfig = flag.String("a", "../auth", "auth etc dir")
	fileConfig = flag.String("f", "../files", "files etc dir")
)

func main() {
	flag.Parse()
	app := woocoo.New()

	rmscnf := &conf.AppConfiguration{
		Configuration: conf.New(conf.WithBaseDir(*rmcConfig), conf.WithGlobal(false)).Load(),
	}
	otelStop := applyOTEL(rmscnf)
	defer otelStop()

	koapp.BuildCacheComponents(rmscnf)
	rmsSvr := rms.NewServer(rmscnf)
	rmsEngine := rmsSvr.BuildWebEngine()

	authcnf := &conf.AppConfiguration{
		Configuration: conf.New(conf.WithBaseDir(*authConfig), conf.WithGlobal(false)).Load(),
	}
	koapp.BuildCacheComponents(authcnf)
	authSrv := auth.NewServer(authcnf)
	authEngine := authSrv.BuildWebServer()
	authSrv.RegisterWebEngine(authEngine.Router().FindGroup("/").Group)

	filecnf := &conf.AppConfiguration{
		Configuration: conf.New(conf.WithBaseDir(*fileConfig), conf.WithGlobal(false)).Load(),
	}
	koapp.BuildCacheComponents(filecnf)
	fileSrv := files.NewServer(filecnf)
	fileEngine := fileSrv.BuildWebServer()
	fileSrv.RegisterWebEngine(fileEngine.Router().FindGroup("/").Group)

	app.RegisterServer(rmsEngine, authEngine, authSrv.GrpcSrv, fileEngine, ecx.ChangeSet)

	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}

// Apply 尝试注册otel,如果配置中有otel配置,则注册.并返回关闭函数
func applyOTEL(cnf *conf.AppConfiguration) func() {
	if cnf.IsSet("otel") {
		otelCnf := cnf.Sub("otel")
		otelcfg := telemetry.NewConfig(otelCnf,
			telemetry.WithPropagator(b3.New()),
		)
		return func() {
			otelcfg.Shutdown()
		}
	}
	return func() {}
}
