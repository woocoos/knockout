package hook

import (
	"context"
	"entgo.io/ent"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/apppolicyview"
	"github.com/woocoos/knockout/ent/hook"
	"strconv"
)

// AppPolicyViewPathHook 创建策略视图路径hook.
func AppPolicyViewPathHook() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.AppPolicyViewFunc(func(ctx context.Context, mutation *gen.AppPolicyViewMutation) (gen.Value, error) {
				if _, ok := mutation.Path(); ok {
					return next.Mutate(ctx, mutation)
				}
				if pid, ok := mutation.ParentID(); ok {
					id, _ := mutation.ID()
					path := strconv.FormatInt(int64(id), 36)
					if pid == 0 {
						mutation.SetPath(path)
					} else {
						parentPath := ""
						prow, err := mutation.Client().AppPolicyView.Query().Where(apppolicyview.ID(pid)).
							Select(apppolicyview.FieldPath).Only(ctx)
						if err != nil {
							if !gen.IsNotFound(err) {
								return nil, err
							}
							parentPath = ""
						} else {
							parentPath = prow.Path + "/"
						}
						mutation.SetPath(parentPath + path)
					}
				}
				return next.Mutate(ctx, mutation)
			})
		}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}