package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/knockout/codegen/entgen/types"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/hook"
	"strings"
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
	}
}

// Fields of the OrgPolicy.
func (OrgPolicy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Optional().Immutable().Comment("组织ID"),
		field.Int("app_id").Optional().Comment("所属应用").Annotations(entgql.Skip(entgql.SkipAll)),
		field.Int("app_policy_id").Optional().Comment("所属应用策略,如果是自定义应用策略,则为空"),
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
	}
}

func (OrgPolicy) Hooks() []ent.Hook {
	return []ent.Hook{
		// check rules
		rulesHook(),
	}
}

// rulesHook 检查策略规则是否合法,以error的方式返回
// 由于在策略制定UI中,策略规则是以json的形式输入,因此需要在保存时检查规则是否合法.
//  1. 检查action是否存在; 2. 检查action是否重复
//
// 在编辑阶段,只是做基本的action存在验证,并未验证对应App是否在租户空间内,在生成租户的授权时,需要过滤掉不存在的App.
func rulesHook() ent.Hook {
	rulesField := "rules"
	return hook.If(
		func(next ent.Mutator) ent.Mutator {
			return hook.OrgPolicyFunc(func(ctx context.Context, m *gen.OrgPolicyMutation) (ent.Value, error) {
				rules, ok := m.Rules()
				if !ok {
					return next.Mutate(ctx, m)
				}
				acs := make(map[string][]string)
				for _, rule := range rules {
					for _, action := range rule.Actions {
						if action == "*" {
							return nil, fmt.Errorf("missing app code %s", action)
						}
						// 分离出appcode和action
						parts := strings.SplitN(action, ":", 2)
						if len(parts) != 2 {
							return nil, fmt.Errorf("invalid action %s", action)
						}
						if parts[1] != "*" {
							appcode := parts[0]
							acs[appcode] = append(acs[appcode], parts[1])
						}
					}
				}
				// 检查action是否存在
				for appcode, actions := range acs {
					// 检查action是否存在
					count, err := m.Client().AppAction.Query().Where(
						appaction.NameIn(actions...),
						appaction.HasAppWith(app.Code(appcode))).Count(ctx)
					if err != nil {
						return nil, err
					}
					if count != len(actions) {
						return nil, fmt.Errorf("invalid action in %s", actions)
					}
				}
				return next.Mutate(ctx, m)
			})
		}, hook.HasFields(rulesField),
	)
}
