package resource

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/woocoos/knockout-go/pkg/identity"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/intercept"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/security"
	"slices"
	"strconv"
)

type EntHook struct {
	service        *Service
	administrators *gen.OrgRole
}

type tenantTraverseKey struct{}

// SkipTenantTraverse returns a new context that skips the tenant traverse.
func SkipTenantTraverse(parent context.Context) context.Context {
	return context.WithValue(parent, tenantTraverseKey{}, true)
}

// isSkipTenantTraverse returns true if the tenant traverse. should be skipped.
func isSkipTenantTraverse(ctx context.Context) bool {
	skip, _ := ctx.Value(tenantTraverseKey{}).(bool)
	return skip
}

func NewEntHook(service *Service) *EntHook {
	// 获取超级管理员角色
	or, err := service.Client.OrgRole.Query().Where(orgrole.NameEQ("administrators")).Only(context.Background())
	if err != nil && !gen.IsNotFound(err) {
		panic(err)
	}
	return &EntHook{
		service:        service,
		administrators: or,
	}
}

// OrgTraverseFunc 只允许查询当前租户下的资源
func (e *EntHook) OrgTraverseFunc(orgField string) ent.Interceptor {
	var orgin = func(ids []int, w interface{ WhereP(...func(*sql.Selector)) }) {
		if len(ids) == 1 {
			w.WhereP(sql.FieldEQ(orgField, ids[0]))
		} else {
			w.WhereP(sql.FieldIn(orgField, ids...))
		}
	}
	return intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
		if isSkipTenantTraverse(ctx) {
			return nil
		}
		// 判断是否root管理员
		if has, err := e.IsAdministrator(ctx); err != nil {
			return err
		} else if has {
			return nil
		}
		tid, err := identity.TenantIDFromContext(ctx)
		if err != nil {
			return err
		}
		var ids []int
		if err := cache.Get(ctx, RefTenantsCacheKey(tid), &ids); err != nil {
			if ids, err = e.service.GetRefTenants(ctx); err != nil {
				return err
			}
		}
		orgin(ids, q)
		return nil
	})
}

func (e *EntHook) IsAdministrator(ctx context.Context) (bool, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	has, err := security.HasRoleForUser(uid, e.administrators.ID, tid)
	if err != nil {
		return false, err
	}
	if has {
		return true, nil
	}
	return false, nil
}

func RefTenantsCacheKey(tid int) string {
	return "RefTenants:" + strconv.Itoa(tid)
}

func (e *EntHook) OrgMutationInAllowOrg() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ev ent.Value, err error) {
			// 判断是否root管理员
			if has, err := e.IsAdministrator(ctx); err != nil {
				return nil, err
			} else if has {
				return next.Mutate(ctx, m)
			}
			switch m.Op() {
			case ent.OpCreate:
				tid, err := identity.TenantIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				var ids []int
				if err = cache.Get(ctx, RefTenantsCacheKey(tid), &ids); err != nil {
					if ids, err = e.service.GetRefTenants(ctx); err != nil {
						return nil, err
					}
				}
				value, _ := m.Field("parent_id")
				orgId := value.(int)
				if !slices.Contains(ids, orgId) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpUpdateOne, ent.OpDeleteOne:
				mi := m.(interface {
					ParentID() (v int, ok bool)
					OldParentID(ctx context.Context) (v int, err error)
				})
				oid, ok := mi.ParentID()
				if !ok {
					oid, err = mi.OldParentID(ctx)
					if err != nil {
						return nil, err
					}
				}
				tid, err := identity.TenantIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				var ids []int
				if err = cache.Get(ctx, RefTenantsCacheKey(tid), &ids); err != nil {
					if ids, err = e.service.GetRefTenants(ctx); err != nil {
						return nil, err
					}
				}
				if !slices.Contains(ids, oid) {
					return nil, ErrTenantIDNotAllow
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}

func (e *EntHook) AppMutationInAllowOrg() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ev ent.Value, err error) {
			// 判断是否root管理员
			if has, err := e.IsAdministrator(ctx); err != nil {
				return nil, err
			} else if has {
				return next.Mutate(ctx, m)
			}
			switch m.Op() {
			case ent.OpCreate:
				var tid, err = identity.TenantIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				var ids []int
				if err = cache.Get(ctx, RefTenantsCacheKey(tid), &ids); err != nil {
					if ids, err = e.service.GetRefTenants(ctx); err != nil {
						return nil, err
					}
				}
				value, _ := m.Field("owner_org_id")
				orgId := value.(int)
				if !slices.Contains(ids, orgId) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpUpdateOne, ent.OpDeleteOne:
				mi := m.(interface {
					OwnerOrgID() (v int, ok bool)
					OldOwnerOrgID(ctx context.Context) (v int, err error)
				})
				oid, ok := mi.OwnerOrgID()
				if !ok {
					oid, err = mi.OldOwnerOrgID(ctx)
					if err != nil {
						return nil, err
					}
				}
				tid, err := identity.TenantIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				var ids []int
				if err = cache.Get(ctx, RefTenantsCacheKey(tid), &ids); err != nil {
					if ids, err = e.service.GetRefTenants(ctx); err != nil {
						return nil, err
					}
				}
				if !slices.Contains(ids, oid) {
					return nil, ErrTenantIDNotAllow
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}

func (e *EntHook) MutationInAllowOrg() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ev ent.Value, err error) {
			// 判断是否root管理员
			if has, err := e.IsAdministrator(ctx); err != nil {
				return nil, err
			} else if has {
				return next.Mutate(ctx, m)
			}
			switch m.Op() {
			case ent.OpCreate:
				var tid, err = identity.TenantIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				var ids []int
				if err = cache.Get(ctx, RefTenantsCacheKey(tid), &ids); err != nil {
					if ids, err = e.service.GetRefTenants(ctx); err != nil {
						return nil, err
					}
				}
				value, _ := m.Field("org_id")
				orgId := value.(int)
				if !slices.Contains(ids, orgId) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpUpdateOne, ent.OpDeleteOne:
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
				if err = cache.Get(ctx, RefTenantsCacheKey(tid), &ids); err != nil {
					if ids, err = e.service.GetRefTenants(ctx); err != nil {
						return nil, err
					}
				}
				if !slices.Contains(ids, oid) {
					return nil, ErrTenantIDNotAllow
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}
