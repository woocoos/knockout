package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout/codegen/entgen/types"
)

// OrgPolicy 组织中的策略.基本包括来源于应用初始化的策略和组织自定义的策略.
//
// 在策略制定UI时,原则上策略应该对应一个应用,但是也可以是多个应用策略集合在一个策略中.因此策略中的app_id和app_policy_id都是可选的.
type OrgPolicy struct {
	ent.Schema
}

func (OrgPolicy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_policy"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (OrgPolicy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the OrgPolicy.
func (OrgPolicy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Optional().Immutable().Comment("租户ID"),
		field.Int("app_id").Optional().Comment("所属应用").Annotations(entgql.Skip(entgql.SkipAll)),
		field.Int("app_policy_id").Optional().Nillable().Comment("所属应用策略,如果是自定义应用策略,则为空"),
		field.String("name").Comment("策略名称"),
		field.String("comments").Optional().Comment("描述"),
		field.JSON("rules", []*types.PolicyRule{}).Comment("策略规则"),
	}
}

// Edges of the OrgPolicy.
func (OrgPolicy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("org", Org.Type).Ref("policies").Unique().Immutable().Field("org_id"),
		edge.To("permissions", Permission.Type),
		edge.From("app_policy", AppPolicy.Type).Unique().Ref("org_policies").Field("app_policy_id"),
		edge.To("app", App.Type).Unique().Field("app_id"),
	}
}
