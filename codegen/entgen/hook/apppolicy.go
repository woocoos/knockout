package hook

import (
	"context"
	"entgo.io/ent"
	"github.com/gin-gonic/gin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/hook"
	"strings"
)

// AppPolicyRulesHook 检查规则是否合法.
func AppPolicyRulesHook() ent.Hook {
	return hook.If(
		func(next ent.Mutator) ent.Mutator {
			return hook.AppPolicyFunc(func(ctx context.Context, m *gen.AppPolicyMutation) (ent.Value, error) {
				rules, ok := m.Rules()
				if !ok {
					return next.Mutate(ctx, m)
				}

				acs := make(map[string][]string)
				for _, rule := range rules {
					for _, action := range rule.Actions {
						if action == "*" {
							return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "missing app code %s", action)
						}
						// 分离出appcode和action
						parts := strings.SplitN(action, ":", 2)
						if len(parts) != 2 {
							return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "invalid action %s", action)
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
						return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "invalid action in %s", actions)
					}
				}
				return next.Mutate(ctx, m)
			})
		}, hook.HasFields("rules"),
	)
}