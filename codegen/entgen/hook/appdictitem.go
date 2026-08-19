package hook

import (
	"context"
	"entgo.io/ent"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	"github.com/woocoos/knockout-go/pkg/identity"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/appdict"
	"github.com/woocoos/knockout/ent/appdictitem"
	"github.com/woocoos/knockout/ent/hook"
)

// AppDictItemCodeUniqueHook 检查code唯一性.
func AppDictItemCodeUniqueHook() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.AppDictItemFunc(func(ctx context.Context, m *gen.AppDictItemMutation) (gen.Value, error) {
			code, cok := m.Code()
			id, iok := m.ID()
			dictid, dok := m.DictID()
			if cok && iok && dok {
				return next.Mutate(ctx, m)
			}
			oid, ook := m.OrgID()
			var err error
			var has bool
			if ook {
				has, err = m.Client().AppDictItem.Query().Where(
					appdictitem.DictID(dictid), appdictitem.Code(code), appdictitem.IDNEQ(id), appdictitem.OrgID(oid)).Exist(ctx)
			} else {
				has, err = m.Client().AppDictItem.Query().Where(
					appdictitem.DictID(dictid), appdictitem.Code(code), appdictitem.IDNEQ(id), appdictitem.OrgIDIsNil()).Exist(ctx)
			}
			if err != nil {
				return nil, err
			}
			if has {
				return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "code exists:%s", code)
			}
			return next.Mutate(ctx, m)
		})
	}, ent.OpCreate|ent.OpUpdateOne)
}

// AppDictItemRefCodeHook 设置ref_code.
func AppDictItemRefCodeHook() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.AppDictItemFunc(func(ctx context.Context, m *gen.AppDictItemMutation) (gen.Value, error) {
			dict, _ := m.DictID()
			dr, err := m.Client().AppDict.Query().Where(appdict.ID(dict)).WithApp().Only(schemax.SkipTenantPrivacy(ctx))
			if err != nil {
				return nil, err
			}
			app, err := dr.App(schemax.SkipTenantPrivacy(ctx))
			if err != nil {
				return nil, err
			}
			m.SetRefCode(fmt.Sprintf("%s:%s", app.Code, dr.Code))
			return next.Mutate(ctx, m)
		})
	}, ent.OpCreate)
}

// AppDictItemOrgIDHook 设置org_id.
func AppDictItemOrgIDHook() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.AppDictItemFunc(func(ctx context.Context, m *gen.AppDictItemMutation) (gen.Value, error) {
			tid, err := identity.TenantIDFromContext(ctx)
			if err != nil {
				return nil, err
			}
			if m.Op().Is(ent.OpCreate) {
				org, _ := m.OrgID()
				if org != 0 {
					m.SetOrgID(tid)
				} else {
					// TODO check pemission
				}
			}
			return next.Mutate(ctx, m)
		})
	}, ent.OpCreate|ent.OpUpdateOne)
}

// AppDictItemDeleteStatusHook 检查删除时的状态.
func AppDictItemDeleteStatusHook() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.AppDictItemFunc(func(ctx context.Context, m *gen.AppDictItemMutation) (gen.Value, error) {
			id, _ := m.ID()
			status, ok := m.Status()
			if !ok {
				row, err := m.Client().AppDictItem.Get(ctx, id)
				if err != nil {
					return nil, err
				}
				status = row.Status
			}
			if status != typex.SimpleStatusInactive {
				return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "can't not delete not active stauts")
			}
			return next.Mutate(ctx, m)
		})
	}, ent.OpDeleteOne)
}