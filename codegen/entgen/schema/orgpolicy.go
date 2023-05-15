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
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/knockout/codegen/entgen/types"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/hook"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/security"
	"github.com/woocoos/knockout/service/resource"
	"strconv"
	"strings"
)

const splitPolicyEffect = "&&"

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
				// 所有actions及resources
				var naas, nars []string
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
						naas = append(naas, action+splitPolicyEffect+rule.Effect.String())
					}
					for _, res := range rule.Resources {
						nars = append(nars, res+splitPolicyEffect+rule.Effect.String())
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

				// 判断是否更新policy,更新policy需同步更新至casbin
				if m.Op() == ent.OpUpdateOne || m.Op() == ent.OpUpdate {
					//
					pid, ok := m.ID()
					if ok {
						tid, err := identity.TenantIDFromContext(ctx)
						if err != nil {
							return nil, err
						}
						// 查询数据库policy
						op, err := m.Client().OrgPolicy.Query().Where(orgpolicy.ID(pid)).Only(ctx)
						if err != nil {
							return nil, err
						}
						// 旧的actions、resources
						var oaas, oars []string
						for _, rule := range op.Rules {
							for _, action := range rule.Actions {
								oaas = append(oaas, action+splitPolicyEffect+rule.Effect.String())
							}
							for _, res := range rule.Resources {
								oars = append(oars, res+splitPolicyEffect+rule.Effect.String())
							}
						}
						// 更新policy时,获取 addActions：新增actions、delActions：删除actions、addResources：新增resources、delResources：删除resources
						addActions, delActions := resource.DiffArrays(naas, oaas)
						addResources, delResources := resource.DiffArrays(nars, oars)
						domain, err := m.Client().Org.Query().Where(org.ID(tid)).Select(org.FieldDomain).String(ctx)
						if err != nil {
							return nil, err
						}
						if domain == "" {
							return nil, fmt.Errorf("organization %d domain is empty", tid)
						}
						// 查询policy授权的users、roles
						ps, err := m.Client().Permission.Query().Where(permission.OrgPolicyID(pid), permission.OrgID(tid)).All(ctx)
						if err != nil {
							return nil, err
						}
						var uids, rids []int
						for _, p := range ps {
							if p.PrincipalKind == permission.PrincipalKindUser {
								uids = append(uids, p.UserID)
							} else if p.PrincipalKind == permission.PrincipalKindRole {
								rids = append(rids, p.RoleID)
							}
						}

						// 查询授权给user的policy
						ups, err := m.Client().Permission.Query().Where(permission.PrincipalKindEQ(permission.PrincipalKindUser), permission.UserIDIn(uids...), permission.OrgID(tid)).WithOrgPolicy().All(ctx)
						if err != nil {
							return nil, err
						}
						uops := make(map[string][]*gen.OrgPolicy)
						// 根据用户授权分组
						for _, p := range ups {
							if uops[strconv.Itoa(p.UserID)] == nil {
								uops[strconv.Itoa(p.UserID)] = []*gen.OrgPolicy{
									p.Edges.OrgPolicy,
								}
							} else {
								uops[strconv.Itoa(p.UserID)] = append(uops[strconv.Itoa(p.UserID)], p.Edges.OrgPolicy)
							}
						}
						// 更新权限
						err = updatePolicyToCasbin(pid, uops, domain, addActions, addResources, delActions, delResources, permission.PrincipalKindUser)
						if err != nil {
							return nil, err
						}

						// 查询授权给role的policy
						rps, err := m.Client().Permission.Query().Where(permission.PrincipalKindEQ(permission.PrincipalKindRole), permission.RoleIDIn(rids...), permission.OrgID(tid)).WithOrgPolicy().All(ctx)
						if err != nil {
							return nil, err
						}
						rops := make(map[string][]*gen.OrgPolicy)
						// 根据用户授权分组
						for _, p := range rps {
							if rops[strconv.Itoa(p.RoleID)] == nil {
								rops[strconv.Itoa(p.RoleID)] = []*gen.OrgPolicy{
									p.Edges.OrgPolicy,
								}
							} else {
								rops[strconv.Itoa(p.RoleID)] = append(rops[strconv.Itoa(p.RoleID)], p.Edges.OrgPolicy)
							}
						}
						// 更新权限
						err = updatePolicyToCasbin(pid, rops, domain, addActions, addResources, delActions, delResources, permission.PrincipalKindRole)
						if err != nil {
							return nil, err
						}
					}
				}
				return next.Mutate(ctx, m)
			})
		}, hook.HasFields(rulesField),
	)
}

func splitPolicyActionsOrResources(data []string) ([]string, []string) {
	var allows, denys []string
	for _, res := range data {
		parts := strings.SplitN(res, splitPolicyEffect, 2)
		if parts[1] == types.PolicyEffectAllow.String() {
			allows = append(allows, parts[0])
		} else if parts[1] == types.PolicyEffectDeny.String() {
			denys = append(denys, parts[0])
		}
	}
	return allows, denys
}

func updatePolicyToCasbin(pid int, uops map[string][]*gen.OrgPolicy, domain string, addActions []string, addResources []string, delActions []string, delResources []string, principalKind permission.PrincipalKind) error {
	// 根据用户授权的policy来判断需要移除的actions及resources
	for uid, uop := range uops {
		var otherPolicyActions, otherPolicyResources []string
		for _, p := range uop {
			// 过滤当前policy id
			if p.ID == pid {
				continue
			}
			for _, rule := range p.Rules {
				// actions
				for _, action := range rule.Actions {
					otherPolicyActions = append(otherPolicyActions, action+splitPolicyEffect+rule.Effect.String())
				}
				// resource
				for _, res := range rule.Actions {
					otherPolicyResources = append(otherPolicyResources, res+splitPolicyEffect+rule.Effect.String())
				}
			}
		}
		// 移除的actions与用户其他策略的actions对比，不存在其他策略则删除
		actuallyDelActions, _ := resource.DiffArrays(delActions, otherPolicyActions)
		// 移除的resources与用户其他策略的resources对比，不存在其他策略则删除
		actuallyDelResources, _ := resource.DiffArrays(delResources, otherPolicyResources)

		var addPolicyRule []*types.PolicyRule
		var delPolicyRule []*types.PolicyRule
		splitAddAllowActions, splitAddDenyActions := splitPolicyActionsOrResources(addActions)
		splitAddAllowResources, splitAddDenyResources := splitPolicyActionsOrResources(addResources)
		splitDelAllowActions, splitDelDenyActions := splitPolicyActionsOrResources(actuallyDelActions)
		splitDelAllowResources, splitDelDenyResources := splitPolicyActionsOrResources(actuallyDelResources)
		addAllowRule := types.PolicyRule{
			Effect:    types.PolicyEffectAllow,
			Actions:   splitAddAllowActions,
			Resources: splitAddAllowResources,
		}
		addDenyRule := types.PolicyRule{
			Effect:    types.PolicyEffectDeny,
			Actions:   splitAddDenyActions,
			Resources: splitAddDenyResources,
		}
		addPolicyRule = append(addPolicyRule, &addAllowRule)
		addPolicyRule = append(addPolicyRule, &addDenyRule)
		// 新增的授权
		err := security.GrantPolicy(addPolicyRule, uid, domain, principalKind)
		if err != nil {
			return err
		}
		delAllowRule := types.PolicyRule{
			Effect:    types.PolicyEffectAllow,
			Actions:   splitDelAllowActions,
			Resources: splitDelAllowResources,
		}
		delDenyRule := types.PolicyRule{
			Effect:    types.PolicyEffectDeny,
			Actions:   splitDelDenyActions,
			Resources: splitDelDenyResources,
		}
		delPolicyRule = append(delPolicyRule, &delAllowRule)
		delPolicyRule = append(delPolicyRule, &delDenyRule)
		// 删除的授权
		err = security.RevokePolicy(delPolicyRule, uid, domain, principalKind)
		if err != nil {
			return err
		}
	}
	return nil
}
