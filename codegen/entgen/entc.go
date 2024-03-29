//go:build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	entcachegen "github.com/woocoos/entcache/gen"
	"github.com/woocoos/knockout-go/codegen/entx"
	"log"
	"os"
)

func main() {
	ex, err := entgql.NewExtension(
		entx.WithGqlWithTemplates(),
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true), // 需要放在 WithGqlWithTemplates 之后
		entgql.WithConfigPath("codegen/gqlgen/gqlgen.yaml"),
		entgql.WithSchemaPath("api/graphql/ent.graphql"),
		entgql.WithSchemaHook(entx.ChangeRelayNodeType()),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	protoExtension, err := entproto.NewExtension(
		entproto.WithProtoDir("./api/proto"),
	)
	os.MkdirAll("./api/graphql", os.ModePerm)
	opts := []entc.Option{
		entc.Extensions(ex, protoExtension),
		//entc.FeatureNames("privacy", "schema/snapshot"),
		entx.GlobalID(),
		entx.SimplePagination(),
		entcachegen.QueryCache(),
	}
	err = entc.Generate("./codegen/entgen/schema", &gen.Config{
		Package:  "github.com/woocoos/knockout/ent",
		Features: []gen.Feature{gen.FeatureVersionedMigration, gen.FeatureUpsert, gen.FeatureIntercept},
		Target:   "./ent",
	}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
