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

// UserIdentity 用户登陆身份
type UserIdentity struct {
	ent.Schema
}

func (UserIdentity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_identity"},
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (UserIdentity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the UserIdentity.
func (UserIdentity) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional().Immutable(),
		field.Enum("kind").Values("name", "email", "phone", "wechat", "qq").
			Comment("身份标识类型 手机、邮箱、用户名、微信、qq"),
		field.String("code").Optional().Unique().Comment("用户名、邮箱、手机、unionid、qq"),
		field.String("code_extend").Optional().Comment("扩展标识码,比如微信的openID"),
		field.Enum("status").GoType(typex.SimpleStatus("")).Optional().
			Comment("状态,部分登陆方式需要验证通过才可启用"),
	}
}

// Edges of the UserIdentity.
func (UserIdentity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("identities").Field("user_id").Immutable().Unique(),
	}
}
