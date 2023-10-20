package main

import (
	_ "github.com/go-sql-driver/mysql"
	ecx "github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/koapp"
	_ "github.com/woocoos/knockout-go/pkg/snowflake"
	"github.com/woocoos/knockout/cmd/internal/files"
	_ "github.com/woocoos/knockout/ent/runtime"
)

func main() {
	app := koapp.New()

	srv := files.NewServer(app.AppConfiguration())
	webEngine := srv.BuildWebServer()
	srv.RegisterWebEngine(webEngine.Router().FindGroup("/").Group)

	app.RegisterServer(webEngine, ecx.ChangeSet)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
