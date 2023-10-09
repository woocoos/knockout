package main

import (
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/knockout/cmd/internal/files"
	"github.com/woocoos/knockout/cmd/internal/otel"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout/ent/runtime"
)

func main() {
	app := woocoo.New()
	otelStop := otel.Apply(app.AppConfiguration())
	defer otelStop()

	err := snowflake.SetDefaultNode(app.AppConfiguration().Sub("snowflake"))
	if err != nil {
		log.Panic(err)
	}

	authSrv := files.NewServer(app.AppConfiguration())
	webEngine := authSrv.BuildWebServer()
	authSrv.RegisterWebEngine(webEngine.Router().FindGroup("/").Group)
	app.RegisterServer(webEngine, ecx.ChangeSet)
	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
