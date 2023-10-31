package main

import (
	"github.com/woocoos/knockout-go/ent/clientx"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/api/oas/file"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout-go/pkg/snowflake"
	_ "github.com/woocoos/knockout/ent/runtime"
)

func main() {
	app := koapp.New()

	srv := file.NewServer(app)
	app.RegisterServer(srv, clientx.ChangeSet)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
