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

// OrganizationApp 组织采用的应用.组织不可自行添加应用.应用需要统一到app中.
type OrganizationApp struct {
	ent.Schema
}

func (OrganizationApp) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "organization_app"},
		field.ID("org_id", "app_id"),
		entgql.Skip(entgql.SkipAll),
	}
}

func (OrganizationApp) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.AuditMixin{},
	}
}

// Fields of the OrganizationApp.
func (OrganizationApp) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Comment("组织ID"),
		field.Int("app_id").Comment("应用ID"),
	}
}

// Edges of the OrganizationApp.
func (OrganizationApp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Unique().Required().Field("app_id"),
		edge.To("organization", Organization.Type).Unique().Required().Field("org_id"),
	}
}
