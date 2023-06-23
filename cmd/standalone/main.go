package main

import (
	"flag"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/knockout/cmd/internal/auth"
	"github.com/woocoos/knockout/cmd/internal/otel"
	"github.com/woocoos/knockout/cmd/internal/rms"
)

var (
	rmcConfig  = flag.String("r", "../adminx", "rms etc dir")
	authConfig = flag.String("a", "../auth", "auth etc dir")
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
	app.RegisterServer(rmsSvr)

	authcnf := conf.New(conf.WithBaseDir(*authConfig), conf.WithGlobal(false)).Load()
	authSrv := auth.NewServer(&conf.AppConfiguration{Configuration: authcnf})
	app.RegisterServer(authSrv.WebSrv, authSrv.GrpcSrv)
	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
