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

// OrgUser holds the schema definition for the OrgUser entity.
type OrgUser struct {
	ent.Schema
}

func (OrgUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_user"},
		entgql.Skip(entgql.SkipType),
	}
}

func (OrgUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the OrgUser.
func (OrgUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Comment("组织ID"),
		field.Int("user_id").Comment("用户ID"),
		field.String("display_name").Comment("在组织内的显示名称"),
	}
}

// Edges of the OrgUser.
func (OrgUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("org", Org.Type).Unique().Required().Field("org_id").
			Annotations(entgql.Skip(entgql.SkipAll)),
		edge.To("user", User.Type).Unique().Required().Field("user_id").
			Annotations(entgql.Skip(entgql.SkipAll)),
	}
}
