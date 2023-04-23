package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
)

// AppRole holds the schema definition for the AppRole entity.
type AppRole struct {
	ent.Schema
}

func (AppRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_role"},
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (AppRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the AppRole.
func (AppRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Optional().Immutable().Comment("所属应用"),
		field.String("name").Unique().Comment("角色名称"),
		field.String("comments").Optional().Comment("备注").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Bool("auto_grant").Default(false).Comment("标识是否自动授予到账户"),
		field.Bool("editable").Default(false).Comment("授权后是否可编辑"),
	}
}

// Edges of the AppRole.
func (AppRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("roles").Unique().Immutable().Field("app_id"),
		edge.To("policies", AppPolicy.Type).
			Annotations(entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipMutationCreateInput)).Comment("权限授权策略").
			Through("app_role_policy", AppRolePolicy.Type),
	}
}
