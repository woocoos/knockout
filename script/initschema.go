//go:build ignore

package main

import (
	"context"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	entadapter "github.com/woocoos/casbin-ent-adapter"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/migrate"
	"github.com/woocoos/knockout/script/data"
	"log"
)

// receive two arguments: the migration name and the database dsn.
var (
	dsn  = flag.String("dsn", "root@tcp(localhost:3306)/portal?parseTime=true&loc=Local", "")
	name = flag.String("name", data.DefaultDriver, "driver name")
)

func main() {
	flag.Parse()
	data.ParseDNS(dsn)
	client, err := ent.Open(*name, *dsn)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	_, err = entadapter.NewAdapter(*name, *dsn, entadapter.WithMigration())
	if err != nil {
		log.Fatalf("failed creating casbin adapter: %v", err)
	}
}
