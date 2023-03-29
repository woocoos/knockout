package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	"time"
)

// OrgUser 用户组织关系.
//
// 用户与组织至少有一个关系(用户-根组织),然后用户可与根组织下的组织建立关系.概念可为:用户加入公司,用户分配入部门.
type OrgUser struct {
	ent.Schema
}

func (OrgUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_user"},
		entgql.Skip(entgql.SkipType),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
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
		field.Time("joined_at").Default(time.Now).Comment("加入时间"),
		field.String("display_name").Comment("在组织内的显示名称"),
	}
}

// Edges of the OrgUser.
func (OrgUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("org", Org.Type).Unique().Required().Field("org_id"),
		edge.To("user", User.Type).Unique().Required().Field("user_id"),
		edge.From("org_roles", OrgRole.Type).Ref("org_users").
			Annotations(entgql.Skip(entgql.SkipAll)),
	}
}
