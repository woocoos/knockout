package hook

import (
	"context"
	"entgo.io/ent"
	"github.com/gin-gonic/gin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/appdictitem"
	"github.com/woocoos/knockout/ent/hook"
)

// AppDictDeleteHook 删除字典前检查是否有子项.
func AppDictDeleteHook() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.AppDictFunc(func(ctx context.Context, m *gen.AppDictMutation) (gen.Value, error) {
			id, _ := m.ID()
			has, err := m.Client().AppDictItem.Query().Where(appdictitem.DictID(id)).Exist(ctx)
			if err != nil {
				return nil, err
			}
			if has {
				return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "has items,please remove items first")
			}
			return next.Mutate(ctx, m)
		})
	}, ent.OpDeleteOne)
}