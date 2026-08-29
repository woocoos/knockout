package main

import (
	"github.com/tsingsun/woocoo/pkg/security"
	entadapter "github.com/woocoos/casbin-ent-adapter"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/ent/clientx"
	authzcasbin "github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/api/graphql"
	schemahook "github.com/woocoos/knockout/codegen/entgen/hook"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/job"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout-go/pkg/snowflake"
	_ "github.com/woocoos/knockout/ent/runtime"
)

func main() {
	app := koapp.New()
	cnf := app.AppConfiguration()

	ents := koapp.BuildEntComponents(cnf)
	drv := ents["portal"]
	portalClient := ent.NewClient(ent.Driver(drv))
	schemahook.RegisterAllHooks(portalClient)
	casbinClient := casbinent.NewClient(casbinent.Driver(drv))
	defer casbinClient.Close()
	if cnf.Development {
		portalClient = portalClient.Debug()
		casbinClient = casbinClient.Debug()
	}

	// 初始化全局 authorizer (casbin)
	if cnf.IsSet("authz") {
		adapter, err := entadapter.NewAdapterWithClient(casbinClient)
		if err != nil {
			panic(err)
		}
		authorizer, err := authzcasbin.NewAuthorizer(cnf.Sub("authz"), authzcasbin.WithAdapter(adapter))
		if err != nil {
			panic(err)
		}
		security.SetDefaultAuthorizer(authorizer)
	}

	// 初始化错误处理
	if err := fmterr.InitErrorHandler(cnf.Sub("errors")); err != nil {
		panic(err)
	}

	kosdk, err := api.NewSDK(cnf.Sub("kosdk"))
	if err != nil {
		panic(err)
	}

	rmsSvr := graphql.NewServer(cnf,
		graphql.WithPortalDB(portalClient),
		graphql.WithKOSdk(kosdk),
		graphql.WithDefaultCache(cnf),
	)

	// 调度
	if cnf.IsSet("job") {
		jobSrv, err := job.NewServer(cnf.Sub("job"))
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
