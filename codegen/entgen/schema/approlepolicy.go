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

// AppRolePolicy holds the schema definition for the AppRolePolicy entity.
type AppRolePolicy struct {
	ent.Schema
}

func (AppRolePolicy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_role_policy"},
		entgql.Skip(entgql.SkipType),
	}
}

func (AppRolePolicy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the AppRolePolicy.
func (AppRolePolicy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_role_id").Comment("应用角色ID"),
		field.Int("app_policy_id").Comment("策略ID"),
		field.Int("app_id").Comment("应用ID"),
	}
}

// Edges of the AppRolePolicy.
func (AppRolePolicy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", AppRole.Type).Unique().Required().Field("app_role_id"),
		edge.To("policy", AppPolicy.Type).Unique().Required().Field("app_policy_id"),
	}
}
