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

// UserLoginProfile 用户登陆配置
//
// 用户登陆配置
type UserLoginProfile struct {
	ent.Schema
}

func (UserLoginProfile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_login_profile"},
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (UserLoginProfile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the UserLoginProfile.
func (UserLoginProfile) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional().Immutable(),
		field.String("last_login_ip").Optional().
			Annotations(entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("last_login_at").Optional().Comment("最后登陆时间").
			Annotations(entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Bool("can_login").Optional().Comment("是否允许使用密码登陆控制台"),
		field.Enum("set_kind").Values("keep", "customer", "auto").Comment("设置密码:keep-保持不变,customer-客户自行设置,auto-自动生成"),
		field.Bool("password_reset").Optional().Comment("下次登陆时需要重置密码"),
		field.Bool("verify_device").Comment("是否开启设备认证"),
		field.Bool("mfa_enabled").Optional().Comment("是否开启多因素验证").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("mfa_secret").Optional().MaxLen(100).Comment("多因素验证密钥BASE32").Sensitive().
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Enum("mfa_status").Optional().
			GoType(typex.SimpleStatus("")).Comment("多因素验证状态").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the UserLoginProfile.
func (UserLoginProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("login_profile").Field("user_id").Immutable().Unique(),
	}
}
