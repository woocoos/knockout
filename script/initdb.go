//go:build ignore

package main

import (
	"context"
	"flag"
	entadapter "github.com/woocoos/casbin-ent-adapter"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/migrate"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// receive two arguments: the migration name and the database dsn.
var (
	dsn  = flag.String("dsn", "root:@tcp(localhost:3306)/portal", "")
	name = flag.String("name", "mysql", "driver name")
)

func main() {
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
