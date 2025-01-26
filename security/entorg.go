package security

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/pkg/authz"
	"github.com/woocoos/knockout-go/pkg/identity"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/intercept"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/user"
	"slices"
	"strconv"
	"time"
)

import "errors"

var (
	ErrTenantIDNotAllow          = errors.New("not allow action in your tenant")
	ErrMutationOtherUserNotAllow = errors.New("not allow mutation other user")
)

var (
	AllOp        = ent.OpCreate | ent.OpUpdateOne | ent.OpUpdate | ent.OpDeleteOne | ent.OpDelete
	userOtherRes = fmt.Sprintf("%s%s%s", "resource", authz.ArnSplit, "allowMutationOtherUsers")
)

func RefTenantsCacheKey(tid int) string {
	return "RefTenants:" + strconv.Itoa(tid)
}

// IDer is used for Field Named ID
type IDer interface {
	ID() (int, bool)
}

// OrgIDer is used for Field Named OrgID
type OrgIDer interface {
	OrgID() (int, bool)
	OldOrgID(context.Context) (int, error)
}

// UserIDer is used for Field Named UserID
type UserIDer interface {
	UserID() (int, bool)
	OldUserID(context.Context) (int, error)
}

type EntHook struct {
	client         *gen.Client
	administrators *gen.OrgRole
}

// NewEntHook create a new EntHook.
//
// 注意, 此查询依赖OrgRole,如果有Hook时,需要注意OrgRole的加入顺序.
func NewEntHook(client *gen.Client) *EntHook {
	// 获取超级管理员角色
	or, err := client.OrgRole.Query().Where(orgrole.NameEQ("administrators")).Only(context.Background())
	if err != nil && !gen.IsNotFound(err) {
		panic(err)
	}
	return &EntHook{
		client:         client,
		administrators: or,
	}
}

// GetRefTenants 获取当前租户下的所有子租户ID,包括当前租户.
func GetRefTenants(ctx context.Context, client *gen.Client) (ids []int, err error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	ctx = schemax.SkipTenantPrivacy(ctx)
	path, err := client.Org.Query().Where(org.ID(tid)).Select(org.FieldPath).String(ctx)
	if err != nil {
		return nil, err
	}
	ids, err = client.Org.Query().Where(org.PathHasPrefix(path)).Select(org.FieldID).Ints(ctx)
	if ids != nil {
		err = cache.Set(ctx, RefTenantsCacheKey(tid), ids, cache.WithTTL(time.Minute*5))
	}
	return
}

func (e *EntHook) IsAdministrator(ctx context.Context) (bool, error) {
	if e.administrators == nil {
		return false, nil
	}
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	has, err := HasRoleForUser(uid, e.administrators.ID, tid)
	if err != nil {
		return false, err
	}
	if has {
		return true, nil
	}
	return false, nil
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
		if schemax.IfSkipTenantPrivacy(ctx) {
			return nil
		}
		// 判断是否root管理员
		if has, err := e.IsAdministrator(ctx); err != nil {
			return err
		} else if has {
			return nil
		}
		ids, err := GetCacheRefTenants(ctx, e.client)
		if err != nil {
			return err
		}
		orgin(ids, q)
		return nil
	})
}

func GetCacheRefTenants(ctx context.Context, client *gen.Client) ([]int, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	var ids []int
	if err := cache.Get(ctx, RefTenantsCacheKey(tid), &ids); err != nil {
		if ids, err = GetRefTenants(ctx, client); err != nil {
			return nil, err
		}
	}
	return ids, nil
}

// OrgMutationInAllowOrg 允许org mutation操作在当前租户下进行
//
// where 在根据不同的Schema对象实现不同的where方法
func (e *EntHook) OrgMutationInAllowOrg(op ent.Op, fieldName string) ent.Hook {
	isSelf := fieldName == org.FieldID
	var getOrgIdInOne = func(ctx context.Context, m ent.Mutation) (id int, err error) {
		pid, ok := m.Field(fieldName)
		if !ok {
			if isSelf {
				mi := m.(IDer)
				id, ok = mi.ID()
				if !ok {
					return 0, errors.New("org id not found")
				}
				return id, nil
			} else {
				mi := m.(OrgIDer)
				pid, err = mi.OldOrgID(schemax.SkipTenantPrivacy(ctx))
				if err != nil {
					return 0, err
				}
			}
		}
		id = pid.(int)
		return
	}
	var where = func(m ent.Mutation, ids []int) {
		w := m.(interface{ WhereP(...func(*sql.Selector)) })
		if len(ids) > 1 {
			w.WhereP(sql.FieldIn(fieldName, ids...))
		} else {
			w.WhereP(sql.FieldEQ(fieldName, ids[0]))
		}
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if !m.Op().Is(op) {
				return next.Mutate(ctx, m)
			}
			// 判断是否root管理员
			if has, err := e.IsAdministrator(ctx); err != nil {
				return nil, err
			} else if has {
				return next.Mutate(ctx, m)
			}
			ids, err := GetCacheRefTenants(ctx, e.client)
			if err != nil {
				return nil, err
			}

			switch m.Op() {
			case ent.OpCreate:
				var value ent.Value
				if isSelf {
					value, _ = m.Field(org.FieldParentID)
				} else {
					value, _ = m.Field(fieldName)
				}
				orgId := value.(int)
				if !slices.Contains(ids, orgId) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpUpdateOne:
				id, err := getOrgIdInOne(ctx, m)
				if err != nil {
					return nil, err
				}
				if !slices.Contains(ids, id) {
					return nil, ErrTenantIDNotAllow
				}
			case ent.OpUpdate, ent.OpDelete, ent.OpDeleteOne:
				where(m, ids)
			}
			return next.Mutate(ctx, m)
		})
	}
}

func (e *EntHook) allowMutationOtherUsers(ctx context.Context, uid int) (bool, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	return CheckUserPermission(
		strconv.Itoa(uid),
		strconv.Itoa(tid),
		userOtherRes,
		authz.ActionTypeRead)
}

// UserMutationAllow User更新操作，只允许有权限的更新其他用户信息
func (e *EntHook) UserMutationAllow(op ent.Op, fieldName string) ent.Hook {
	isSelf := fieldName == user.FieldID
	var getUserIdInOne = func(ctx context.Context, m ent.Mutation) (id int, err error) {
		uid, ok := m.Field(fieldName)
		if !ok {
			if isSelf {
				mi := m.(IDer)
				id, ok = mi.ID()
				if !ok {
					return 0, errors.New("org id not found")
				}
				return id, nil
			} else {
				mi := m.(UserIDer)
				uid, err = mi.OldUserID(ctx)
				if err != nil {
					return 0, err
				}
			}
		}
		id = uid.(int)
		return
	}

	var where = func(m ent.Mutation, uid int) {
		w := m.(interface{ WhereP(...func(*sql.Selector)) })
		w.WhereP(sql.FieldEQ(fieldName, uid))
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ev ent.Value, err error) {
			if !m.Op().Is(op) {
				return next.Mutate(ctx, m)
			}
			// 判断是否root管理员
			if has, err := e.IsAdministrator(ctx); err != nil {
				return nil, err
			} else if has {
				return next.Mutate(ctx, m)
			}
			uid, err := identity.UserIDFromContext(ctx)
			if err != nil {
				return nil, err
			}
			has, err := e.allowMutationOtherUsers(ctx, uid)
			if err != nil {
				return nil, err
			}

			switch m.Op() {
			case ent.OpCreate:
				var value ent.Value
				if isSelf {
					value, _ = m.Field(user.FieldID)
				} else {
					value, _ = m.Field(fieldName)
				}
				tid := value.(int)
				if tid != uid && !has {
					return nil, ErrMutationOtherUserNotAllow
				}
			case ent.OpUpdateOne:
				tid, err := getUserIdInOne(ctx, m)
				if err != nil {
					return nil, ErrMutationOtherUserNotAllow
				}
				if tid != uid && !has {
					return nil, ErrMutationOtherUserNotAllow
				}
			case ent.OpUpdate, ent.OpDelete, ent.OpDeleteOne:
				if !has {
					where(m, uid)
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}
