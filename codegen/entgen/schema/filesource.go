package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/woocoos/entco/schemax"
)

// FileSource holds the schema definition for the FileSource entity.
type FileSource struct {
	ent.Schema
}

func (FileSource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "file_source"},
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (FileSource) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the FileSource.
func (FileSource) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("kind").Values(
			"local", "alioss",
		).Comment("文件来源"),
		field.String("comments").Optional().Comment("备注").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.String("endpoint").Optional().Comment("对外服务的访问域名"),
		field.String("region").MaxLen(100).Optional().Comment("地域，数据存储的物理位置。本地存储为：localhost"),
		field.String("bucket").MaxLen(100).Optional().Comment("文件存储空间。本地存储为：assets"),
	}
}

// Edges of the FileSource.
func (FileSource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("files", File.Type).Comment("所有文件").Annotations(entgql.RelayConnection()),
	}
}

// Indexes of the FileSource.
func (FileSource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("kind", "endpoint", "region", "bucket").Unique(),
	}
}
