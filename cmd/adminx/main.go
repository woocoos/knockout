package main

import (
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/api/graphql"
	"github.com/woocoos/knockout/ent"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout-go/pkg/snowflake"
	_ "github.com/woocoos/knockout/ent/runtime"
)

func main() {
	app := koapp.New()

	ents := koapp.BuildEntComponents(app.AppConfiguration())
	drv := ents["portal"]
	portalClient := ent.NewClient(ent.Driver(drv))
	casbinClient := casbinent.NewClient(casbinent.Driver(drv))
	if app.AppConfiguration().Development {
		portalClient = portalClient.Debug()
		casbinClient = casbinClient.Debug()
	}
	rmsSvr := graphql.NewServer(app.AppConfiguration(),
		graphql.WithCasbinDB(casbinClient), graphql.WithPortalDB(portalClient))

	app.RegisterServer(rmsSvr, clientx.ChangeSet)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
