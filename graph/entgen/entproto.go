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
			Target: "./api",
		})
	if err != nil {
		log.Fatalf("entproto: failed loading ent graph: %v", err)
	}
	if err := entproto.Generate(graph); err != nil {
		log.Fatalf("entproto: failed generating protos: %s", err)
	}
}
