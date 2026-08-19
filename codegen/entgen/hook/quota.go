package hook

import (
	"context"
	"entgo.io/ent"
	"github.com/gin-gonic/gin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/hook"
)

// QuotaTenantOrOrgIDHook 验证 tenant_id 和 user_id 至少有一个有值.
func QuotaTenantOrOrgIDHook() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.QuotaFunc(func(ctx context.Context, mutation *gen.QuotaMutation) (gen.Value, error) {
				_, tidOk := mutation.TenantID()
				_, uidOk := mutation.UserID()
				if mutation.Op() == ent.OpCreate {
					if !tidOk && !uidOk {
						return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "at least one of tenant_id and org_id has a value")
					}
				}
				return next.Mutate(ctx, mutation)
			})
		}, ent.OpCreate)
}