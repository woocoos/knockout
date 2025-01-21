package resource

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout/ent/intercept"
	"github.com/woocoos/knockout/ent/org"
	"slices"
	"strconv"
)

type EntHook struct {
	service *Service
}

func NewEntHook(service *Service) *EntHook {
	return &EntHook{
		service: service,
	}
}

// OrgTraverseFunc 只允许查询当前租户下的org
func (e *EntHook) OrgTraverseFunc() ent.Interceptor {
	var orgin = func(ids []int, w interface{ WhereP(...func(*sql.Selector)) }) {
		if len(ids) == 1 {
			w.WhereP(sql.FieldEQ(org.FieldID, ids[0]))
		} else {
			w.WhereP(sql.FieldIn(org.FieldID, ids...))
		}
	}
	return intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
		ids, err := e.service.GetRefTenants(ctx)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		orgin(ids, q)
		return nil
	})
}

func (e *EntHook) RefTenantsCacheKey(tid int) string {
	return "RefTenants:" + strconv.Itoa(tid)
}

func (e *EntHook) MutationInAllowOrg() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ev ent.Value, err error) {
			switch m.Op() {
			case ent.OpCreate:
				tid, err := identity.TenantIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				var ids []int
				if err = cache.Get(ctx, e.RefTenantsCacheKey(tid), ids); err != nil {
					if ids, err = e.service.GetRefTenants(ctx); err != nil {
						return nil, err
					}
				}
				value, _ := m.Field("org_id")
				orgId := value.(int)
				if !slices.Contains(ids, orgId) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpUpdateOne:
				mi := m.(interface {
					OrgID() (v int, ok bool)
					OldOrgID(ctx context.Context) (v int, err error)
				})
				oid, ok := mi.OrgID()
				if !ok {
					oid, err = mi.OldOrgID(ctx)
					if err != nil {
						return nil, err
					}
				}
				tid, err := identity.TenantIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				var ids []int
				if err = cache.Get(ctx, e.RefTenantsCacheKey(tid), ids); err != nil {
					if ids, err = e.service.GetRefTenants(ctx); err != nil {
						return nil, err
					}
				}
				if !slices.Contains(ids, oid) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpDeleteOne:
				// TODO
			}
			return next.Mutate(ctx, m)
		})
	}
}
