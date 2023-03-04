//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	graph, err := entc.LoadGraph("./graph/entgen/schema",
		&gen.Config{
			Target:  "./api",
			Package: "github.com/woocoos/knockout/api",
		})
	if err != nil {
		log.Fatalf("entproto: failed loading ent graph: %v", err)
	}
	if err := entproto.GenerateWithEntPackage(graph, "github.com/woocoos/knockout/ent"); err != nil {
		log.Fatalf("entproto: failed generating protos: %s", err)
	}
}
