package main

import (
	"flag"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/knockout/cmd/internal/auth"
	"github.com/woocoos/knockout/cmd/internal/files"
	"github.com/woocoos/knockout/cmd/internal/otel"
	"github.com/woocoos/knockout/cmd/internal/rms"
)

var (
	rmcConfig  = flag.String("r", "../adminx", "rms etc dir")
	authConfig = flag.String("a", "../auth", "auth etc dir")
	fileConfig = flag.String("f", "../files", "files etc dir")
)

func main() {
	flag.Parse()
	rmsc := conf.New(conf.WithBaseDir(*rmcConfig), conf.WithGlobal(false)).Load()
	rmscnf := &conf.AppConfiguration{Configuration: rmsc}
	app := woocoo.New()
	otelStop := otel.Apply(rmscnf)
	defer otelStop()

	err := snowflake.SetDefaultNode(rmscnf.Sub("snowflake"))
	if err != nil {
		log.Panic(err)
	}
	rmsSvr := rms.NewServer(rmscnf)
	rmsEngine := rmsSvr.BuildWebEngine()

	authcnf := &conf.AppConfiguration{
		Configuration: conf.New(conf.WithBaseDir(*authConfig), conf.WithGlobal(false)).Load(),
	}
	authSrv := auth.NewServer(authcnf)
	authEngine := authSrv.BuildWebServer()
	authSrv.RegisterWebEngine(authEngine.Router().FindGroup("/").Group)

	filecnf := &conf.AppConfiguration{
		Configuration: conf.New(conf.WithBaseDir(*fileConfig), conf.WithGlobal(false)).Load(),
	}
	fileSrv := files.NewServer(filecnf)
	fileEngine := fileSrv.BuildWebServer()
	fileSrv.RegisterWebEngine(fileEngine.Router().FindGroup("/").Group)

	app.RegisterServer(rmsEngine, authEngine, authSrv.GrpcSrv, fileEngine, ecx.ChangeSet)

	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
