package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tsingsun/woocoo/web"
	ecx "github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/cmd/internal/auth"
	_ "github.com/woocoos/knockout/ent/runtime"
)

func main() {
	app := koapp.New()

	authSrv := auth.NewServer(app.AppConfiguration())

	webEngine := web.New(web.WithConfiguration(app.AppConfiguration().Sub("web")))
	authSrv.RegisterWebEngine(webEngine.Router().FindGroup("/").Group)

	app.RegisterServer(webEngine, authSrv.GrpcSrv, ecx.ChangeSet)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
