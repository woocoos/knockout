package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/hook"
)

// UserAddr holds the schema definition for the UserAddr entity.
type UserAddr struct {
	ent.Schema
}

func (UserAddr) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_addr"},
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (UserAddr) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the UserAddr.
func (UserAddr) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional().Immutable(),
		field.Enum("addr_type").NamedValues(
			"basic", "basic",
			"addr", "addr",
		).Comment("地址类型，basic：基本信息，addr：收货地址").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Int("region_id").Optional().Comment("地址地区"),
		field.String("addr").Optional().Comment("详细地址"),
		field.String("email").MaxLen(45).Optional().Comment("邮箱"),
		field.String("fax").MaxLen(45).Optional().Comment("传真"),
		field.String("tel").MaxLen(45).Optional().Comment("电话"),
		field.String("mobile").MaxLen(45).Optional().Comment("手机"),
		field.String("name").MaxLen(45).Optional().Comment("联系人名称"),
		field.Bool("is_default").Default(false).Comment("是否默认地址，类型为addr时使用"),
	}
}

// Edges of the UserAddr.
func (UserAddr) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("addrs").Field("user_id").Unique().Immutable(),
		edge.To("region", Region.Type).Field("region_id").Unique().Comment("地区信息"),
	}
}

// Hooks of the UserAddr.
func (UserAddr) Hooks() []ent.Hook {
	return []ent.Hook{
		basicAddrUnique(),
	}
}

func basicAddrUnique() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.UserAddrFunc(func(ctx context.Context, m *gen.UserAddrMutation) (ent.Value, error) {
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate|ent.OpUpdateOne)
}
