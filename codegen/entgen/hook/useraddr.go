package hook

import (
	"context"
	"entgo.io/ent"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/hook"
)

// UserAddrContactUniqueHook 联系人地址唯一性hook.
func UserAddrContactUniqueHook() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.UserAddrFunc(func(ctx context.Context, m *gen.UserAddrMutation) (gen.Value, error) {
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate|ent.OpUpdateOne)
}