package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
)

// FileIdentity holds the schema definition for the FileIdentity entity.
type FileIdentity struct {
	ent.Schema
}

func (FileIdentity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "file_identity"},
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (FileIdentity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the FileIdentity.
func (FileIdentity) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tenant_id").Comment("组织ID").Annotations(entgql.Type("ID")),
		field.String("access_key_id").MaxLen(255).Comment("accesskey id"),
		field.String("access_key_secret").MaxLen(255).Comment("accesskey secret").Annotations(entgql.Skip(entgql.SkipType)),
		field.Int("file_source_id").Comment("文件来源ID"),
		field.String("role_arn").MaxLen(255).Comment("角色的资源名称(ARN)，用于STS"),
		field.String("policy").Optional().Comment("指定返回的STS令牌的权限的策略"),
		field.Int("duration_seconds").Optional().Comment("STS令牌的有效期，默认3600s"),
		field.Bool("is_default").Default(false).Comment("租户默认的凭证"),
		field.String("comments").Optional().Comment("备注").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the FileIdentity.
func (FileIdentity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("source", FileSource.Type).Ref("identities").Required().Unique().Field("file_source_id"),
	}
}
