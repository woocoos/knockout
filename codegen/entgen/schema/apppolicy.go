package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/codegen/entgen/types"
)

// AppPolicy 应用定义的策略.
//
// 对于应用来说,策略定义是一种模板.在关联账户后,对该账户进行默认授权即授权范围为账户的授权.
type AppPolicy struct {
	ent.Schema
}

func (AppPolicy) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_policy"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (AppPolicy) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the AppPolicy.
func (AppPolicy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Optional().Immutable().Comment("所属应用"),
		field.Enum("kind").Default("app").Values("app", "view").Comment("分类：app-应用策略、view-策略视图"),
		field.String("name").Comment("策略名称"),
		field.String("comments").Optional().Comment("描述"),
		field.JSON("rules", []*types.PolicyRule{}).Comment("策略规则"),
		field.String("version").Default("1").Comment("版本号").Annotations(entgql.Skip(entgql.SkipAll)),
		field.Bool("auto_grant").Default(false).Comment("标识是否自动授予到账户"),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusActive.String()).
			Optional().Comment("状态"),
	}
}

// Edges of the AppPolicy.
func (AppPolicy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("policies").Unique().Immutable().Field("app_id"),
		edge.From("roles", AppRole.Type).Ref("policies").Through("app_role_policy", AppRolePolicy.Type),
		edge.To("org_policies", OrgPolicy.Type).Comment("策略授权的组织策略"),
		edge.To("policy_views", AppPolicyView.Type).Comment("策略视图"),
	}
}
