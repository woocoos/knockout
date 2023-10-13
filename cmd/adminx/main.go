package main

import (
	"context"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/pkg/koapp"
	"github.com/woocoos/knockout/cmd/internal/rms"

	_ "github.com/woocoos/entco/pkg/snowflake"
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
