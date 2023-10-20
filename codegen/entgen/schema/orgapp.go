package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
)

// OrgApp 组织采用的应用.组织不可自行添加应用.应用需要统一到app中.
type OrgApp struct {
	ent.Schema
}

func (OrgApp) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_app"},
		entgql.Skip(entgql.SkipAll),
	}
}

func (OrgApp) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the OrgApp.
func (OrgApp) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Comment("组织ID"),
		field.Int("app_id").Comment("应用ID"),
	}
}

// Edges of the OrgApp.
func (OrgApp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("app", App.Type).Unique().Required().Field("app_id"),
		edge.To("org", Org.Type).Unique().Required().Field("org_id"),
	}
}
