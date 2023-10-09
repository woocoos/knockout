package main

import (
	"context"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/knockout/cmd/internal/otel"
	"github.com/woocoos/knockout/cmd/internal/rms"
)

func main() {
	app := woocoo.New()
	otelStop := otel.Apply(app.AppConfiguration())
	defer otelStop()

	err := snowflake.SetDefaultNode(app.AppConfiguration().Sub("snowflake"))
	if err != nil {
		log.Panic(err)
	}

	rmsSvr := rms.NewServer(app.AppConfiguration())
	// 如果需要使用由run来关闭服务,可以使用app.RegisterServer
	defer rmsSvr.Stop(context.Background())
	webEngine := rmsSvr.BuildWebEngine()
	app.RegisterServer(webEngine, ecx.ChangeSet)
	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
