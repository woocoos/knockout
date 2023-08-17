package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/woocoos/entco/schemax"
	"regexp"
)

// AppAction holds the schema definition for the AppAction entity.
//
// Kind:
//
//	function: 功能性权限,此时需要明确该操作的类型(读,写,列表)
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
		field.Int("app_id").Optional().Immutable().Comment("所属应用"),
		field.String("name").Comment("名称").
			Match(regexp.MustCompile("[a-zA-Z0-9/]+$")),
		field.Enum("kind").Values("restful", "graphql", "rpc", "function", "route").Comment("restful,graphql,rpc,function"),
		field.Enum("method").Values("read", "write", "list").Comment("操作方法:读,写,列表"),
		field.String("comments").Optional().Comment("备注").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the AppAction.
func (AppAction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("actions").Unique().Immutable().Field("app_id"),
		edge.To("menus", AppMenu.Type).Comment("被引用的菜单项"),
	}
}

func (AppAction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "name").Unique(),
	}
}
