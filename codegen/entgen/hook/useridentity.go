package hook

import (
	"context"
	"entgo.io/ent"
	"github.com/gin-gonic/gin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/hook"
	"github.com/woocoos/knockout/ent/useridentity"
)

// UserIdentityCodeUniqueHook 检查code唯一性.
func UserIdentityCodeUniqueHook() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.UserIdentityFunc(func(ctx context.Context, m *gen.UserIdentityMutation) (ent.Value, error) {
				nc, ok := m.Code()
				if !ok {
					return next.Mutate(ctx, m)
				}
				// 检查code是否唯一
				has, err := m.Client().UserIdentity.Query().Where(useridentity.Code(nc), useridentity.UserIDNotNil()).
					Exist(ctx)
				if err != nil {
					return nil, err
				}
				if has {
					return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "code %s already exists", nc)
				}
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate|ent.OpUpdateOne)
}