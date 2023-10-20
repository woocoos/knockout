package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"fmt"
	"github.com/woocoos/knockout-go/ent/schemax"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/appdictitem"
	"github.com/woocoos/knockout/ent/hook"
)

// AppDict holds the schema definition for the AppDict entity.
type AppDict struct {
	ent.Schema
}

func (AppDict) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_dict"},
		entgql.QueryField().Description("数据字典查询"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (AppDict) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the AppDict.
func (AppDict) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Optional().Immutable().Comment("所属应用"),
		field.String("code").MinLen(3).MaxLen(20).Immutable().
			Comment("用于标识应用资源的唯一代码,尽量简短"),
		field.String("name").NotEmpty().MaxLen(45).Comment("名称"),
		field.String("comments").Optional().Comment("备注").Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the AppDict.
func (AppDict) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("dicts").Unique().Immutable().Field("app_id"),
		edge.To("items", AppDictItem.Type),
	}
}

func (AppDict) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "code").Unique(),
	}
}

func (AppDict) Hooks() []ent.Hook {
	return []ent.Hook{
		// check subs.
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.AppDictFunc(func(ctx context.Context, m *gen.AppDictMutation) (gen.Value, error) {
				id, _ := m.ID()
				has, err := m.Client().AppDictItem.Query().Where(appdictitem.DictID(id)).Exist(ctx)
				if err != nil {
					return nil, err
				}
				if has {
					return nil, fmt.Errorf("has items,please remove items first")
				}
				return next.Mutate(ctx, m)
			})
		}, ent.OpDeleteOne),
	}
}
