package main

import (
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/web"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/knockout/cmd/internal/auth"
	"github.com/woocoos/knockout/cmd/internal/otel"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout/ent/runtime"
)

func main() {
	app := woocoo.New()
	otelStop := otel.Apply(app.AppConfiguration())
	defer otelStop()

	authSrv := auth.NewServer(app.AppConfiguration())
	webEngine := web.New(web.WithConfiguration(app.AppConfiguration().Sub("web")))
	authSrv.RegisterWebEngine(webEngine.Router().FindGroup("/").Group)
	app.RegisterServer(webEngine, authSrv.GrpcSrv, ecx.ChangeSet)
	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
