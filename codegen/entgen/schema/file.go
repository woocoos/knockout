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

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "file"},
		entgql.RelayConnection(),
	}
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("文件名称"),
		field.Int("source_id").Comment("文件来源"),
		field.Int("tenant_id").Comment("租户ID"),
		field.String("path").Comment("文件相对路径"),
		field.Int("size").Optional().Comment("文件大小，单位为B"),
		field.String("mine_type").MaxLen(100).Optional().Comment("媒体类型，如：image/png").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.String("md5").MaxLen(100).Optional().Comment("md5值").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("source", FileSource.Type).Ref("files").
			Field("source_id").Comment("文件来源").Unique().Required(),
	}
}

// Indexes of the FileSource.
func (File) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_id", "source_id", "path").Unique(),
	}
}
