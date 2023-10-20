package main

import (
	"context"
	"github.com/tsingsun/woocoo/pkg/log"
	ecx "github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/cmd/internal/rms"

	_ "github.com/woocoos/knockout-go/pkg/snowflake"
)

func main() {
	app := koapp.New()

	rmsSvr := rms.NewServer(app.AppConfiguration())
	// 如果需要使用由run来关闭服务,可以使用app.RegisterServer
	defer rmsSvr.Stop(context.Background())
	webEngine := rmsSvr.BuildWebEngine()
	app.RegisterServer(webEngine, ecx.ChangeSet)
	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
