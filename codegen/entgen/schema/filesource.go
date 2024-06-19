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
		field.Int("tenant_id").Comment("组织ID").Annotations(entgql.Type("ID")),
		field.Enum("kind").Values(
			"local", "minio", "aliOSS",
		).Comment("文件来源"),
		field.String("comments").Optional().Comment("备注").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.String("access_key_id").MaxLen(255).Comment("accesskey id"),
		field.String("access_key_secret").MaxLen(255).Comment("accesskey secret"),
		field.String("endpoint").MaxLen(255).Comment("对外服务的访问域名"),
		field.String("sts_endpoint").MaxLen(255).Comment("sts服务的访问域名"),
		field.String("region").MaxLen(100).Comment("地域，数据存储的物理位置"),
		field.String("bucket").MaxLen(255).Comment("文件存储空间"),
		field.String("bucketUrl").MaxLen(255).Optional().Comment("文件存储空间地址，用于匹配url"),
		field.String("role_arn").MaxLen(255).Comment("角色的资源名称(ARN)，用于STS"),
		field.String("policy").Optional().Comment("指定返回的STS令牌的权限的策略"),
		field.Int("duration_seconds").Optional().Comment("STS令牌的有效期，默认3600s"),
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
