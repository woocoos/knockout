package main

import (
	"github.com/tsingsun/woocoo"
	"github.com/woocoos/knockout/cmd/internal/auth"
	"github.com/woocoos/knockout/cmd/internal/otel"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout/ent/runtime"
)

func main() {
	app := woocoo.New()
	otelStop := otel.Apply(app.AppConfiguration())
	defer otelStop()

	authSrv := auth.NewServer(app.AppConfiguration())
	app.RegisterServer(authSrv.WebSrv, authSrv.GrpcSrv)
	if err := app.Run(); err != nil {
		log.Panic(err)
	}
}
