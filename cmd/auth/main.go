package main

import (
	"github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/api/oas/auth"
	"github.com/woocoos/knockout/cmd/auth/casbin"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout/ent/runtime"
)

func main() {
	app := koapp.New()
	ents := koapp.BuildEntComponents(app.AppConfiguration())
	drv := ents["portal"]
	defer drv.Close()
	authSrv, err := auth.NewServer(app.AppConfiguration(), auth.WithAuthDB(drv))
	if err != nil {
		panic(err)
	}

	casbinSrv, err := casbin.NewServer(app.AppConfiguration(), casbin.WithAuthDB(drv))
	if err != nil {
		panic(err)
	}

	app.RegisterServer(authSrv, casbinSrv, clientx.ChangeSet)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
