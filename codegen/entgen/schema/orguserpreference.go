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

// OrgUserPreference holds the schema definition for the OrgUserPreference entity.
type OrgUserPreference struct {
	ent.Schema
}

func (OrgUserPreference) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_user_preference"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (OrgUserPreference) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the OrgUserPreference.
func (OrgUserPreference) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("用户id"),
		field.Int("org_id").Immutable().Comment("组织ID").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.JSON("menu_favorite", []int{}).Optional().Comment("用户收藏菜单").Annotations(entgql.Type("[ID!]")),
		field.JSON("menu_recent", []int{}).Optional().Comment("用户最近访问菜单").Annotations(entgql.Type("[ID!]")),
	}
}

// Edges of the OrgUserPreference.
func (OrgUserPreference) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required().Field("user_id").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		edge.To("org", Org.Type).Immutable().Unique().Required().Field("org_id").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}
