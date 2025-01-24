package resource

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/woocoos/knockout-go/pkg/authz"
	"github.com/woocoos/knockout-go/pkg/identity"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/fileidentity"
	"github.com/woocoos/knockout/ent/intercept"
	"github.com/woocoos/knockout/ent/oauthclient"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/userdevice"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpasswordpolicy"
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
		ids, err := e.GetCacheRefTenants(ctx)
		if err != nil {
			return err
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

func (e *EntHook) allowMutationOtherUsers(ctx context.Context) (bool, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	return security.CheckUserPermission([]any{
		strconv.Itoa(uid),
		strconv.Itoa(tid),
		fmt.Sprintf("%s%s%s", "resource", ArnSplit, "allowMutationOtherUsers"),
		authz.ActionTypeRead,
	}...)
}

func RefTenantsCacheKey(tid int) string {
	return "RefTenants:" + strconv.Itoa(tid)
}

func (e *EntHook) GetCacheRefTenants(ctx context.Context) ([]int, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var ids []int
	if err := cache.Get(ctx, RefTenantsCacheKey(tid), &ids); err != nil {
		if ids, err = e.service.GetRefTenants(ctx); err != nil {
			return nil, err
		}
	}
	return ids, nil
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
				ids, err := e.GetCacheRefTenants(ctx)
				if err != nil {
					return nil, err
				}
				value, _ := m.Field(org.FieldParentID)
				orgId := value.(int)
				if !slices.Contains(ids, orgId) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpUpdateOne, ent.OpDeleteOne:
				mi := m.(interface {
					ID() (v int, ok bool)
					ParentID() (v int, ok bool)
					OldParentID(ctx context.Context) (v int, err error)
				})
				pid, ok := mi.ParentID()
				if !ok {
					if m.Op() == ent.OpUpdateOne {
						pid, err = mi.OldParentID(ctx)
						if err != nil {
							return nil, err
						}
					} else {
						oid, ok := mi.ID()
						if !ok {
							return nil, ErrTenantIDNotAllow
						}
						pid, err = e.service.Client.Org.Query().Where(org.ID(oid)).Select(org.FieldParentID).Int(ctx)
						if err != nil {
							return nil, err
						}
					}
				}
				ids, err := e.GetCacheRefTenants(ctx)
				if err != nil {
					return nil, err
				}
				if !slices.Contains(ids, pid) {
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
				ids, err := e.GetCacheRefTenants(ctx)
				if err != nil {
					return nil, err
				}
				value, _ := m.Field(app.FieldOwnerOrgID)
				orgId := value.(int)
				if !slices.Contains(ids, orgId) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpUpdateOne, ent.OpDeleteOne:
				mi := m.(interface {
					ID() (v int, ok bool)
					OwnerOrgID() (v int, ok bool)
					OldOwnerOrgID(ctx context.Context) (v int, err error)
				})
				oid, ok := mi.OwnerOrgID()
				if !ok {
					if m.Op() == ent.OpUpdateOne {
						oid, err = mi.OldOwnerOrgID(ctx)
						if err != nil {
							return nil, err
						}
					} else {
						aid, ok := mi.ID()
						if !ok {
							return nil, ErrTenantIDNotAllow
						}
						oid, err = e.service.Client.App.Query().Where(app.ID(aid)).Select(app.FieldOwnerOrgID).Int(ctx)
						if err != nil {
							return nil, err
						}
					}
				}
				ids, err := e.GetCacheRefTenants(ctx)
				if err != nil {
					return nil, err
				}
				if !slices.Contains(ids, oid) {
					return nil, ErrTenantIDNotAllow
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}

func (e *EntHook) UPPMutationInAllowOrg() ent.Hook {
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
				ids, err := e.GetCacheRefTenants(ctx)
				if err != nil {
					return nil, err
				}
				value, _ := m.Field(userpasswordpolicy.FieldTenantID)
				orgId := value.(int)
				if !slices.Contains(ids, orgId) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpUpdateOne, ent.OpDeleteOne:
				mi := m.(interface {
					ID() (v int, ok bool)
					TenantID() (v int, ok bool)
					OldTenantID(ctx context.Context) (v int, err error)
				})
				oid, ok := mi.TenantID()
				if !ok {
					if m.Op() == ent.OpUpdateOne {
						oid, err = mi.OldTenantID(ctx)
						if err != nil {
							return nil, err
						}
					} else {
						uppid, ok := mi.ID()
						if !ok {
							return nil, ErrTenantIDNotAllow
						}
						oid, err = e.service.Client.UserPasswordPolicy.Query().Where(userpasswordpolicy.ID(uppid)).Select(userpasswordpolicy.FieldTenantID).Int(ctx)
						if err != nil {
							return nil, err
						}
					}
				}
				ids, err := e.GetCacheRefTenants(ctx)
				if err != nil {
					return nil, err
				}
				if !slices.Contains(ids, oid) {
					return nil, ErrTenantIDNotAllow
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}

type OrgIDCallback func(ctx context.Context, id int, client *gen.Client) (int, error)

func (e *EntHook) MutationInAllowOrg(orgIDCallback OrgIDCallback) ent.Hook {
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
				ids, err := e.GetCacheRefTenants(ctx)
				if err != nil {
					return nil, err
				}
				value, _ := m.Field("org_id")
				orgId := value.(int)
				if !slices.Contains(ids, orgId) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpUpdateOne, ent.OpDeleteOne:
				mi := m.(interface {
					ID() (v int, ok bool)
					OrgID() (v int, ok bool)
					OldOrgID(ctx context.Context) (v int, err error)
				})
				oid, ok := mi.OrgID()
				if !ok {
					if m.Op() == ent.OpUpdateOne {
						oid, err = mi.OldOrgID(ctx)
						if err != nil {
							return nil, err
						}
					} else {
						id, ok := mi.ID()
						if !ok {
							return nil, ErrTenantIDNotAllow
						}
						oid, err = orgIDCallback(ctx, id, e.service.Client)
						if err != nil {
							return nil, err
						}
					}
				}
				ids, err := e.GetCacheRefTenants(ctx)
				if err != nil {
					return nil, err
				}
				if !slices.Contains(ids, oid) {
					return nil, ErrTenantIDNotAllow
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}

func (e *EntHook) FileIdentityMutationInAllowOrg() ent.Hook {
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
				ids, err := e.GetCacheRefTenants(ctx)
				if err != nil {
					return nil, err
				}
				value, _ := m.Field(fileidentity.FieldTenantID)
				orgId := value.(int)
				if !slices.Contains(ids, orgId) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpUpdateOne, ent.OpDeleteOne:
				mi := m.(interface {
					ID() (v int, ok bool)
					TenantID() (v int, ok bool)
					OldTenantID(ctx context.Context) (v int, err error)
				})
				oid, ok := mi.TenantID()
				if !ok {
					if m.Op() == ent.OpUpdateOne {
						oid, err = mi.OldTenantID(ctx)
						if err != nil {
							return nil, err
						}
					} else {
						aid, ok := mi.ID()
						if !ok {
							return nil, ErrTenantIDNotAllow
						}
						oid, err = e.service.Client.FileIdentity.Query().Where(fileidentity.ID(aid)).Select(fileidentity.FieldTenantID).Int(ctx)
						if err != nil {
							return nil, err
						}
					}
				}
				ids, err := e.GetCacheRefTenants(ctx)
				if err != nil {
					return nil, err
				}
				if !slices.Contains(ids, oid) {
					return nil, ErrTenantIDNotAllow
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}

// UserMutationAllowAll User更新操作，只允许有权限的更新其他用户信息
func (e *EntHook) UserMutationAllowAll() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ev ent.Value, err error) {
			// 判断是否root管理员
			if has, err := e.IsAdministrator(ctx); err != nil {
				return nil, err
			} else if has {
				return next.Mutate(ctx, m)
			}
			switch m.Op() {
			case ent.OpUpdateOne, ent.OpDeleteOne:
				uid, err := identity.UserIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				mi := m.(interface {
					ID() (v int, ok bool)
				})
				oid, ok := mi.ID()
				if !ok {
					return nil, ErrMutationOtherUserNotAllow
				}
				if oid != uid {
					has, err := e.allowMutationOtherUsers(ctx)
					if err != nil {
						return nil, err
					}
					if !has {
						return nil, ErrMutationOtherUserNotAllow
					}
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}

// UserLoginProfileMutationAllowAll UserLoginProfile更新操作，只允许有权限的更新其他用户信息
func (e *EntHook) UserLoginProfileMutationAllowAll() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ev ent.Value, err error) {
			// 判断是否root管理员
			if has, err := e.IsAdministrator(ctx); err != nil {
				return nil, err
			} else if has {
				return next.Mutate(ctx, m)
			}
			switch m.Op() {
			case ent.OpUpdateOne, ent.OpDeleteOne:
				uid, err := identity.UserIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				mi := m.(interface {
					ID() (v int, ok bool)
					UserID() (v int, ok bool)
					OldUserID(ctx context.Context) (v int, err error)
				})
				oid, ok := mi.UserID()
				if !ok {
					if m.Op() == ent.OpUpdateOne {
						oid, err = mi.OldUserID(ctx)
						if err != nil {
							return nil, err
						}
					} else {
						id, ok := mi.ID()
						if !ok {
							return nil, ErrTenantIDNotAllow
						}
						oid, err = e.service.Client.UserLoginProfile.Query().Where(userloginprofile.ID(id)).Select(userloginprofile.FieldUserID).Int(ctx)
						if err != nil {
							return nil, err
						}
					}
				}
				if oid != uid {
					has, err := e.allowMutationOtherUsers(ctx)
					if err != nil {
						return nil, err
					}
					if !has {
						return nil, ErrMutationOtherUserNotAllow
					}
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}

// UserIdentityMutationAllowAll UserIdentity更新操作，只允许有权限的更新其他用户信息
func (e *EntHook) UserIdentityMutationAllowAll() ent.Hook {
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
				uid, err := identity.UserIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				value, _ := m.Field(useridentity.FieldUserID)
				oid := value.(int)
				if oid != uid {
					has, err := e.allowMutationOtherUsers(ctx)
					if err != nil {
						return nil, err
					}
					if !has {
						return nil, ErrMutationOtherUserNotAllow
					}
				}
			case ent.OpUpdateOne, ent.OpDeleteOne:
				uid, err := identity.UserIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				mi := m.(interface {
					ID() (v int, ok bool)
					UserID() (v int, ok bool)
					OldUserID(ctx context.Context) (v int, err error)
				})
				oid, ok := mi.UserID()
				if !ok {
					if m.Op() == ent.OpUpdateOne {
						oid, err = mi.OldUserID(ctx)
						if err != nil {
							return nil, err
						}
					} else {
						id, ok := mi.ID()
						if !ok {
							return nil, ErrTenantIDNotAllow
						}
						oid, err = e.service.Client.UserIdentity.Query().Where(useridentity.ID(id)).Select(useridentity.FieldUserID).Int(ctx)
						if err != nil {
							return nil, err
						}
					}
				}
				if oid != uid {
					has, err := e.allowMutationOtherUsers(ctx)
					if err != nil {
						return nil, err
					}
					if !has {
						return nil, ErrMutationOtherUserNotAllow
					}
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}

// UserDeviceMutationAllowAll UserDevice更新操作，只允许有权限的更新其他用户信息
func (e *EntHook) UserDeviceMutationAllowAll() ent.Hook {
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
				uid, err := identity.UserIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				value, _ := m.Field(userdevice.FieldUserID)
				oid := value.(int)
				if oid != uid {
					has, err := e.allowMutationOtherUsers(ctx)
					if err != nil {
						return nil, err
					}
					if !has {
						return nil, ErrMutationOtherUserNotAllow
					}
				}
			case ent.OpUpdateOne, ent.OpDeleteOne:
				uid, err := identity.UserIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				mi := m.(interface {
					ID() (v int, ok bool)
					UserID() (v int, ok bool)
					OldUserID(ctx context.Context) (v int, err error)
				})
				oid, ok := mi.UserID()
				if !ok {
					if m.Op() == ent.OpUpdateOne {
						oid, err = mi.OldUserID(ctx)
						if err != nil {
							return nil, err
						}
					} else {
						id, ok := mi.ID()
						if !ok {
							return nil, ErrTenantIDNotAllow
						}
						oid, err = e.service.Client.UserDevice.Query().Where(userdevice.ID(id)).Select(userdevice.FieldUserID).Int(ctx)
						if err != nil {
							return nil, err
						}
					}
				}
				if oid != uid {
					has, err := e.allowMutationOtherUsers(ctx)
					if err != nil {
						return nil, err
					}
					if !has {
						return nil, ErrMutationOtherUserNotAllow
					}
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}

func (e *EntHook) OauthClientMutationAllowAll() ent.Hook {
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
				uid, err := identity.UserIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				value, _ := m.Field(oauthclient.FieldUserID)
				oid := value.(int)
				if oid != uid {
					has, err := e.allowMutationOtherUsers(ctx)
					if err != nil {
						return nil, err
					}
					if !has {
						return nil, ErrMutationOtherUserNotAllow
					}
				}
			case ent.OpUpdateOne, ent.OpDeleteOne:
				uid, err := identity.UserIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				mi := m.(interface {
					ID() (v int, ok bool)
					UserID() (v int, ok bool)
					OldUserID(ctx context.Context) (v int, err error)
				})
				oid, ok := mi.UserID()
				if !ok {
					if m.Op() == ent.OpUpdateOne {
						oid, err = mi.OldUserID(ctx)
						if err != nil {
							return nil, err
						}
					} else {
						id, ok := mi.ID()
						if !ok {
							return nil, ErrTenantIDNotAllow
						}
						oid, err = e.service.Client.OauthClient.Query().Where(oauthclient.ID(id)).Select(oauthclient.FieldUserID).Int(ctx)
						if err != nil {
							return nil, err
						}
					}
				}
				if oid != uid {
					has, err := e.allowMutationOtherUsers(ctx)
					if err != nil {
						return nil, err
					}
					if !has {
						return nil, ErrMutationOtherUserNotAllow
					}
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}
