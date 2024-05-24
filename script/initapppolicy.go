//go:build ignore

package main

import (
	"context"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/woocoos/knockout/ent"
	_ "github.com/woocoos/knockout/ent/runtime"
	"github.com/woocoos/knockout/script/data"
	"log"
)

// receive two arguments: the migration name and the database dsn.
var (
	dsn  = flag.String("dsn", "", "")
	name = flag.String("name", data.DefaultDriver, "driver name")
)

func main() {
	flag.Parse()
	data.ParseDNS(dsn)
	client, err := ent.Open(*name, *dsn, ent.Debug())
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	tx, err := client.Tx(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()
	data.InitResourcePolicy(tx)
}
