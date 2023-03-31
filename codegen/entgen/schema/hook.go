package schema

import (
	"context"
	"entgo.io/ent"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/hook"
	"github.com/woocoos/knockout/ent/org"
)

// InitDisplaySortHook 初始化displaySort字段, 表需要有parent_id字段.
func InitDisplaySortHook(table string) ent.Hook {
	parentField := org.FieldParentID
	displayField := org.FieldDisplaySort
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				pid, _ := m.Field(parentField)
				if mx, ok := m.(interface {
					SetDisplaySort(int32)
					Client() *gen.Client
				}); ok {
					var old int
					switch table {
					case org.Table:
						old, _ = mx.Client().Org.Query().Where(org.ParentID(pid.(int))).
							Aggregate(gen.Max(displayField)).Int(ctx)
					case appmenu.Table:
						old, _ = mx.Client().AppMenu.Query().Where(appmenu.ParentID(pid.(int))).
							Aggregate(gen.Max(displayField)).Int(ctx)
					}
					mx.SetDisplaySort(int32(old + 1))
				}
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate)
}
