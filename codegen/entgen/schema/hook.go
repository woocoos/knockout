package schema

import (
	"context"
	"entgo.io/ent"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/appdictitem"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/hook"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/region"
)

// InitDisplaySortHook 初始化displaySort字段, 表需要有parent_id字段.
func InitDisplaySortHook(table string) ent.Hook {
	return InitDisplaySortHookEx(table, org.FieldParentID)
}

func InitDisplaySortHookEx(table, parentField string) ent.Hook {
	displayField := org.FieldDisplaySort
	type displaySorter interface {
		SetDisplaySort(int32)
		Client() *gen.Client
	}
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				pid, _ := m.Field(parentField)
				if mx, ok := m.(displaySorter); ok {
					var old int
					switch table {
					case org.Table:
						old, _ = mx.Client().Org.Query().Where(org.ParentID(pid.(int))).
							Aggregate(gen.Max(displayField)).Int(ctx)
					case appmenu.Table:
						old, _ = mx.Client().AppMenu.Query().Where(appmenu.ParentID(pid.(int))).
							Aggregate(gen.Max(displayField)).Int(ctx)
					case appdictitem.Table:
						old, _ = mx.Client().AppDictItem.Query().Where(appdictitem.DictID(pid.(int))).
							Aggregate(gen.Max(displayField)).Int(ctx)
					case region.Table:
						old, _ = mx.Client().Region.Query().Where(region.ParentID(pid.(int))).
							Aggregate(gen.Max(displayField)).Int(ctx)
					}
					mx.SetDisplaySort(int32(old + 1))
				}
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate)
}
