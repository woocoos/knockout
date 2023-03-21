package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/knockout/ent/appmenu"
)

// AppMenu holds the schema definition for the AppMenu entity.
type AppMenu struct {
	ent.Schema
}

func (AppMenu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_menu"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (AppMenu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the AppMenu.
func (AppMenu) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Comment("所属应用"),
		field.Int("parent_id").Comment("父级ID"),
		field.Enum("kind").Values("dir", "menu").Comment("目录,菜单项"),
		field.String("name").MaxLen(100).Optional().Comment("菜单名称"),
		field.Int("action_id").Optional().Nillable().Comment("操作ID"),
		field.String("comments").Optional().Comment("备注").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Int32("display_sort").Optional().
			Annotations(entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the AppMenu.
func (AppMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("menus").Unique().Required().Field("app_id"),
		edge.From("action", AppAction.Type).Ref("menus").Unique().Field("action_id").
			Comment("需要权限控制时对应的权限"),
	}
}

func (AppMenu) Hooks() []ent.Hook {
	return []ent.Hook{
		InitDisplaySortHook(appmenu.Table),
	}
}
