package main

import (
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/api/graphql"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/job"

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
	// 初始化错误处理
	if err := fmterr.InitErrorHandler(app.AppConfiguration().Sub("errors.errorCodeMap")); err != nil {
		panic(err)
	}
	var err error
	kosdk, err := api.NewSDK(app.AppConfiguration().Sub("kosdk"))
	if err != nil {
		panic(err)
	}
	rmsSvr := graphql.NewServer(app.AppConfiguration(),
		graphql.WithCasbinDB(casbinClient), graphql.WithPortalDB(portalClient), graphql.WithKOSdk(kosdk))

	// 调度
	if app.AppConfiguration().IsSet("job") {
		jobSrv, err := job.NewServer(app.AppConfiguration().Sub("job"))
		if err != nil {
			panic(err)
		}
		err = jobSrv.InitJobs(portalClient, kosdk)
		if err != nil {
			panic(err)
		}
		app.RegisterServer(jobSrv)
	}
	app.RegisterServer(rmsSvr, clientx.ChangeSet)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
