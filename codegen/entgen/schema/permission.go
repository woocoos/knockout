package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/entco/schemax/typex"
)

// Permission 授权产生的权限信息.
type Permission struct {
	ent.Schema
}

func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "permission"},
		entgql.RelayConnection(),
		//entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Immutable().Comment("授权的域根组织"),
		field.Enum("principal_kind").Immutable().Values("user", "role").Comment("授权类型:角色,用户"),
		field.Int("user_id").Optional().Immutable().Comment("授权类型为用户的ID"),
		field.Int("role_id").Optional().Immutable().Comment("授权类型为角色或用户组的ID"),
		field.Int("org_policy_id").Immutable().Comment("策略"),
		field.Time("start_at").Optional().Comment("生效开始时间"),
		field.Time("end_at").Optional().Comment("生效结束时间"),
		field.Enum("status").GoType(typex.SimpleStatus("")).Optional().Comment("状态").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("org", Org.Type).Ref("permissions").
			Unique().Required().Immutable().Field("org_id"),
		edge.To("user", User.Type).Unique().Immutable().Field("user_id"),
		edge.To("role", OrgRole.Type).Unique().Immutable().Field("role_id"),
		edge.From("org_policy", OrgPolicy.Type).Ref("permissions").Unique().Required().Immutable().Field("org_policy_id"),
	}
}
