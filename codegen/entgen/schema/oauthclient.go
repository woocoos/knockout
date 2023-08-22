package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/entco/schemax/typex"
)

// OauthClient holds the schema definition for the OauthClient entity.
type OauthClient struct {
	ent.Schema
}

func (OauthClient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "oauth_client"},
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (OauthClient) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the OauthClient.
func (OauthClient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(100).Comment("名称"),
		field.String("client_id").MaxLen(40).Unique().Immutable().Comment("id").
			Annotations(entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("client_secret").MaxLen(40).Immutable().Comment("密钥").
			Annotations(entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Enum("grant_types").NamedValues(
			"client_credentials", "client_credentials",
		).Comment("授权类型"),
		field.Int("user_id").Comment("关联用户id"),
		field.Time("last_auth_at").Optional().Comment("最后认证时间").
			Annotations(entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusActive.String()).Comment("状态").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the OauthClient.
func (OauthClient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("oauth_clients").Field("user_id").Required().Unique(),
	}
}
