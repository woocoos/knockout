package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/woocoos/knockout-go/ent/schemax"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/hook"
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
		field.Int32("length").Optional().Comment("密码长度"),
		field.Int32("include_element").Optional().Comment("必须包含的元素，异或：1-小写字母，2-大写字母，4-数字，8-符号"),
		field.Int32("include_char").Optional().Comment("最少包含的不同字符数"),
		field.Bool("allow_include_user_name").Optional().Comment("是否允许包含用户名"),
		field.Int32("invalid_day").Optional().Comment("有效天数"),
		field.Bool("invalid_login_limit").Optional().Comment("过期后是否限制登录"),
		field.Int32("retry").Optional().Comment("一小时内密码错误最多尝试次数"),
		field.Int32("captcha_times").Optional().Comment("密码错误多少次出现验证码"),
	}
}

// Edges of the UserPasswordPolicy.
func (UserPasswordPolicy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("org", Org.Type).Ref("user_password_policy").Unique().Immutable().Field("tenant_id"),
	}
}

func (UserPasswordPolicy) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(CheckValueHook(), ent.OpCreate|ent.OpUpdateOne|ent.OpUpdate),
	}
}

func CheckValueHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.UserPasswordPolicyFunc(func(ctx context.Context, mutation *gen.UserPasswordPolicyMutation) (gen.Value, error) {
			if length, ok := mutation.Length(); ok {
				if length < 6 || length > 32 {
					return nil, fmt.Errorf("密码长度必须在6至32位之间")
				}
			}
			if includeChar, ok := mutation.IncludeChar(); ok {
				if includeChar > 8 {
					return nil, fmt.Errorf("密码包含的不同字符数不能大于8个")
				}
			}
			if invalidDay, ok := mutation.InvalidDay(); ok {
				if invalidDay > 1095 {
					return nil, fmt.Errorf("密码有效期不能大于1095天")
				}
			}
			if retry, ok := mutation.Retry(); ok {
				if retry > 32 {
					return nil, fmt.Errorf("密码重试次数不能大于32次")
				}
			}
			return next.Mutate(ctx, mutation)
		})
	}
}
