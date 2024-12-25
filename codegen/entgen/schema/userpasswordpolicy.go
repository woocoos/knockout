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

// UserPasswordPolicy 用户密码策略
type UserPasswordPolicy struct {
	ent.Schema
}

func (UserPasswordPolicy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_password_policy"},
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (UserPasswordPolicy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the UserPasswordPolicy.
func (UserPasswordPolicy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tenant_id").Optional().Immutable().Comment("租户id"),
		field.Int32("length").Optional().Default(6).Min(6).Max(32).Comment("密码最短长度，长度应在6-32位之间"),
		field.Int32("include_element").Optional().Comment("必须包含的元素，异或：1-小写字母，2-大写字母，4-数字，8-符号"),
		field.Int32("include_char").Max(8).Optional().Comment("最少包含的不同字符数，最多8个，0代表不限制"),
		field.Bool("allow_include_user_name").Optional().Comment("是否允许包含用户名"),
		field.Int32("invalid_day").Max(1095).Optional().Comment("有效天数，最大1095天，0代表不过期"),
		field.Bool("invalid_login_limit").Optional().Comment("过期后是否限制登录"),
		field.Int32("retry").Max(32).Optional().Comment("一小时内密码错误最多尝试次数，最大32次，0代表不限次数"),
		field.Int32("captcha_times").Max(5).Optional().Comment("密码错误多少次出现验证码，最大5次，0代表不出现验证码"),
	}
}

// Edges of the UserPasswordPolicy.
func (UserPasswordPolicy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("org", Org.Type).Ref("user_password_policy").Unique().Immutable().Field("tenant_id"),
	}
}
