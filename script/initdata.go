//go:build ignore

package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout/ent/runtime"
	"github.com/woocoos/knockout/script/data"
)

// receive two arguments: the migration name and the database dsn.
var (
	dsn  = flag.String("dsn", "root@tcp(localhost:3306)/portal?parseTime=true&loc=Local", "")
	name = flag.String("name", data.DefaultDriver, "driver name")
)

func main() {
	flag.Parse()
	data.ParseDNS(dsn)
	data.InitBase(*name, *dsn)
}
