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
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/codegen/entgen/types"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/hook"
	"strings"
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
	}
}

// Fields of the AppPolicy.
func (AppPolicy) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Optional().Immutable().Comment("所属应用"),
		field.String("name").Comment("策略名称"),
		field.String("comments").Optional().Comment("描述"),
		field.JSON("rules", []types.PolicyRule{}).Comment("策略规则"),
		field.String("version").Comment("版本号"),
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
	}
}

func (AppPolicy) Hooks() []ent.Hook {
	return []ent.Hook{
		// check rules
		appRulesHook(),
	}
}

// appRulesHook 检查规则.
func appRulesHook() ent.Hook {
	return hook.If(
		func(next ent.Mutator) ent.Mutator {
			return hook.AppPolicyFunc(func(ctx context.Context, m *gen.AppPolicyMutation) (ent.Value, error) {
				rules, ok := m.Rules()
				if !ok {
					return next.Mutate(ctx, m)
				}

				var (
					appcode string
					err     error
					acs     = make(map[string][]string)
				)
				appID, _ := m.AppID()
				if appID == 0 {
					id, _ := m.ID()
					appcode, err = m.Client().AppPolicy.Query().Where(apppolicy.ID(id)).QueryApp().Select(app.FieldCode).String(ctx)
					if err != nil {
						return nil, err
					}
				} else {
					appcode, err = m.Client().App.Query().Where(app.ID(appID)).Select(app.FieldCode).String(ctx)
					if err != nil {
						return nil, err
					}
				}

				for _, rule := range rules {
					for _, action := range rule.Actions {
						if action == "*" {
							continue
						}
						if !strings.HasPrefix(action, appcode+":") {
							return nil, fmt.Errorf("action %s must start with app code %s", action, appcode)
						}
						// 分离出appcode和action
						parts := strings.SplitN(action, ":", 2)
						if len(parts) != 2 {
							return nil, fmt.Errorf("invalid action %s", action)
						}
						acs[appcode] = append(acs[appcode], parts[1])
					}
				}
				// 检查action是否存在
				for cd, actions := range acs {
					// 检查action是否存在
					count, err := m.Client().AppAction.Query().Where(appaction.NameIn(actions...), appaction.HasAppWith(app.Code(cd))).Count(ctx)
					if err != nil {
						return nil, err
					}
					if count != len(actions) {
						return nil, fmt.Errorf("invalid action in %s", actions)
					}
				}
				return next.Mutate(ctx, m)
			})
		}, hook.HasFields("rules"),
	)
}
