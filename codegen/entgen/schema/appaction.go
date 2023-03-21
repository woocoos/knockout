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
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/hook"
	"regexp"
)

// AppAction holds the schema definition for the AppAction entity.
type AppAction struct {
	ent.Schema
}

func (AppAction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_action"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (AppAction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the AppAction.
func (AppAction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Comment("所属应用"),
		field.String("name").Comment("名称").
			Match(regexp.MustCompile("[a-z/]+$")),
		field.Enum("kind").Values("restful", "graphql", "rpc").Comment("restful,graphql,rpc"),
		field.Enum("method").Values("read", "write", "list").Comment("操作方法:读,写,列表"),
		field.String("comments").Optional().Comment("备注").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the AppAction.
func (AppAction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("actions").Unique().Required().Field("app_id"),
		edge.To("menus", AppMenu.Type).Comment("被引用的菜单项"),
		edge.To("resources", AppRes.Type).Comment("引用的资源"),
	}
}

func (AppAction) Hooks() []ent.Hook {
	// app name unique
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.AppActionFunc(func(ctx context.Context, m *gen.AppActionMutation) (gen.Value, error) {
				if m.Op() == ent.OpCreate {
					appid, _ := m.AppID()
					n, _ := m.Name()
					exists, err := m.Client().AppAction.Query().
						Where(appaction.AppID(appid), appaction.Name(n)).Exist(ctx)
					if err != nil {
						return nil, err
					}
					if exists {
						return nil, fmt.Errorf("app permission name %s already exists", n)
					}
				} else if m.Op() == ent.OpUpdateOne || m.Op() == ent.OpUpdate {
					appid, _ := m.AppID()
					n, _ := m.Name()
					c, err := m.Client().AppAction.Query().
						Where(appaction.AppID(appid), appaction.Name(n)).Count(ctx)
					if err != nil {
						return nil, err
					}
					if c > 1 {
						return nil, fmt.Errorf("app permission name %s already exists", n)
					}
				}
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne),
	}
}
