package main

import (
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/log"
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

	app.RegisterServer(rmsSvr)
	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
