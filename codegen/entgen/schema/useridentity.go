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
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/entco/schemax/typex"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/hook"
	"github.com/woocoos/knockout/ent/useridentity"
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
		field.String("code").Optional().Comment("用户名、邮箱、手机、unionid、qq"),
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

// Hooks of the UserIdentity.
func (UserIdentity) Hooks() []ent.Hook {
	return []ent.Hook{
		codeUnique(),
	}
}

func codeUnique() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.UserIdentityFunc(func(ctx context.Context, m *gen.UserIdentityMutation) (ent.Value, error) {
				nc, ok := m.Code()
				if !ok {
					return next.Mutate(ctx, m)
				}
				// 检查code是否唯一
				has, err := m.Client().UserIdentity.Query().Where(useridentity.Code(nc), useridentity.UserIDNotNil()).
					Exist(ctx)
				if err != nil {
					return nil, err
				}
				if has {
					return nil, fmt.Errorf("code %s already exists", nc)
				}
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate|ent.OpUpdateOne)
}
