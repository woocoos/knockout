package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/woocoos/knockout-go/ent/schemax"
)

// FileSource holds the schema definition for the FileSource entity.
type FileSource struct {
	ent.Schema
}

func (FileSource) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "file_source"},
		entgql.RelayConnection(),
		entgql.QueryField("fileSources").Description("文件来源"),
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
		schemax.NotifyMixin{},
	}
}

// Fields of the FileSource.
func (FileSource) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("kind").Values(
			"local", "minio", "aliOSS", "awsS3",
		).Comment("文件来源"),
		field.String("comments").Optional().Comment("备注").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.String("endpoint").MaxLen(255).Comment("对外服务的访问域名"),
		field.Bool("endpoint_immutable").Default(false).Comment("是否禁止修改endpoint，如果是自定义域名设为true"),
		field.String("sts_endpoint").MaxLen(255).Comment("sts服务的访问域名"),
		field.String("region").MaxLen(100).Comment("地域，数据存储的物理位置"),
		field.String("bucket").MaxLen(255).Comment("文件存储空间"),
		field.String("bucket_url").MaxLen(255).Comment("文件存储空间地址，用于匹配url"),
	}
}

// Edges of the FileSource.
func (FileSource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("identities", FileIdentity.Type).Comment("来源凭证").Annotations(entgql.Skip(entgql.SkipType)),
		edge.To("files", File.Type).Comment("所有文件").Annotations(entgql.RelayConnection()),
	}
}

// Indexes of the FileSource.
func (FileSource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("kind", "endpoint", "region", "bucket").Unique(),
	}
}
