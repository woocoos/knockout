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

// OrgRole holds the schema definition for the OrgRole entity.
type OrgRole struct {
	ent.Schema
}

func (OrgRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_role"},
		entgql.Skip(entgql.SkipType),
	}
}

func (OrgRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.AuditMixin{},
	}
}

// Fields of the OrgRole.
func (OrgRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Comment("组织ID"),
		field.Enum("kind").Values("group", "role").Comment("类型,group:组,role:角色"),
		field.String("name").Comment("名称"),
		field.Int("app_role_id").Optional().Comment("角色ID,如有表示该角色来源于应用角色").Annotations(entgql.Skip(entgql.SkipAll)),
		field.String("comments").Optional().Comment("备注").Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the OrgRole.
func (OrgRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("org", Org.Type).Ref("roles_and_groups").Unique().Required().Field("org_id").
			Annotations(entgql.Skip(entgql.SkipAll)),
	}
}
