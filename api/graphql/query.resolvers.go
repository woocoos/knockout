package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"fmt"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appdict"
	"github.com/woocoos/knockout/ent/appdictitem"
	"github.com/woocoos/knockout/ent/appres"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgapp"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/orguserpreference"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/user"
)

// GlobalID is the resolver for the globalID field.
func (r *queryResolver) GlobalID(ctx context.Context, typeArg string, id int) (*string, error) {
	s, err := ent.GlobalID(typeArg, strconv.Itoa(id))
	return &s, err
}

// OrgGroups is the resolver for the orgGroups field.
func (r *queryResolver) OrgGroups(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.OrgRoleOrder, where *ent.OrgRoleWhereInput) (*ent.OrgRoleConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.OrgRole.Query().Where(orgrole.OrgID(tid), orgrole.KindEQ(orgrole.KindGroup)).Paginate(ctx, after, first, before, last,
		ent.WithOrgRoleOrder(orderBy), ent.WithOrgRoleFilter(where.Filter))
}

// OrgRoleUsers is the resolver for the orgRoleUsers field.
func (r *queryResolver) OrgRoleUsers(ctx context.Context, roleID int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	uIds, err := r.resource.GetRoleUserIds(ctx, roleID)
	if err != nil {
		return nil, err
	}
	return r.client.User.Query().Where(user.IDIn(uIds...)).Paginate(ctx, after, first, before, last,
		ent.WithUserOrder(orderBy), ent.WithUserFilter(where.Filter))
}

// OrgRoles is the resolver for the orgRoles field.
func (r *queryResolver) OrgRoles(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.OrgRoleOrder, where *ent.OrgRoleWhereInput) (*ent.OrgRoleConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.OrgRole.Query().Where(orgrole.OrgID(tid), orgrole.KindEQ(orgrole.KindRole)).Paginate(ctx, after, first, before, last,
		ent.WithOrgRoleOrder(orderBy), ent.WithOrgRoleFilter(where.Filter))
}

// AppRoleAssignedToOrgs is the resolver for the appRoleAssignedToOrgs field.
func (r *queryResolver) AppRoleAssignedToOrgs(ctx context.Context, roleID int, where *ent.OrgWhereInput) ([]*ent.Org, error) {
	oIds, err := r.client.OrgRole.Query().Where(orgrole.AppRoleID(roleID)).Select(orgrole.FieldOrgID).Ints(ctx)
	if err != nil {
		return nil, err
	}
	q := r.client.Org.Query()
	q, err = where.Filter(q)
	if err != nil {
		return nil, err
	}
	return q.Where(org.IDIn(oIds...)).All(ctx)
}

// AppPolicyAssignedToOrgs is the resolver for the appPolicyAssignedToOrgs field.
func (r *queryResolver) AppPolicyAssignedToOrgs(ctx context.Context, policyID int, where *ent.OrgWhereInput) ([]*ent.Org, error) {
	oIds, err := r.client.OrgPolicy.Query().Where(orgpolicy.AppPolicyID(policyID)).Select(orgrole.FieldOrgID).Ints(ctx)
	if err != nil {
		return nil, err
	}
	q := r.client.Org.Query()
	q, err = where.Filter(q)
	if err != nil {
		return nil, err
	}
	return q.Where(org.IDIn(oIds...)).All(ctx)
}

// OrgPolicyReferences is the resolver for the orgPolicyReferences field.
func (r *queryResolver) OrgPolicyReferences(ctx context.Context, policyID int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.PermissionOrder, where *ent.PermissionWhereInput) (*ent.PermissionConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	has, err := r.client.OrgPolicy.Query().Where(orgpolicy.ID(policyID), orgpolicy.OrgID(tid)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("policy not exist")
	}
	return r.client.Permission.Query().Where(permission.OrgID(tid), permission.OrgPolicyID(policyID)).Paginate(ctx, after, first, before, last, ent.WithPermissionOrder(orderBy), ent.WithPermissionFilter(where.Filter))
}

// AppResources is the resolver for the appResources field.
func (r *queryResolver) AppResources(ctx context.Context, appID int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.AppResOrder, where *ent.AppResWhereInput) (*ent.AppResConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.AppRes.Query().Where(appres.AppID(appID), appres.HasAppWith(app.OwnerOrgID(tid))).Paginate(ctx, after, first, before, last, ent.WithAppResOrder(orderBy), ent.WithAppResFilter(where.Filter))
}

// OrgAppResources is the resolver for the orgAppResources field.
func (r *queryResolver) OrgAppResources(ctx context.Context, appID int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.AppResOrder, where *ent.AppResWhereInput) (*ent.AppResConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.AppRes.Query().Where(appres.AppID(appID), appres.HasAppWith(app.HasOrgAppWith(orgapp.AppID(appID), orgapp.OrgID(tid)))).Paginate(ctx, after, first, before, last, ent.WithAppResOrder(orderBy), ent.WithAppResFilter(where.Filter))
}

// UserGroups is the resolver for the userGroups field.
func (r *queryResolver) UserGroups(ctx context.Context, userID int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.OrgRoleOrder, where *ent.OrgRoleWhereInput) (*ent.OrgRoleConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.OrgRole.Query().Where(orgrole.HasOrgUsersWith(orguser.OrgID(tid), orguser.UserID(userID)), orgrole.KindIn(orgrole.KindGroup)).
		Paginate(ctx, after, first, before, last, ent.WithOrgRoleOrder(orderBy), ent.WithOrgRoleFilter(where.Filter))
}

// UserExtendGroupPolicies is the resolver for the userExtendGroupPolicies field.
func (r *queryResolver) UserExtendGroupPolicies(ctx context.Context, userID int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.PermissionOrder, where *ent.PermissionWhereInput) (*ent.PermissionConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.Permission.Query().Where(permission.OrgID(tid), permission.PrincipalKindEQ(permission.PrincipalKindRole),
		permission.HasRoleWith(orgrole.HasOrgUsersWith(orguser.UserID(userID), orguser.OrgID(tid)))).
		Paginate(ctx, after, first, before, last, ent.WithPermissionOrder(orderBy), ent.WithPermissionFilter(where.Filter))
}

// UserMenus is the resolver for the userMenus field.
func (r *queryResolver) UserMenus(ctx context.Context, appCode string) ([]*ent.AppMenu, error) {
	return r.resource.GetUserMenus(ctx, appCode)
}

// UserPermissions is the resolver for the userPermissions field.
func (r *queryResolver) UserPermissions(ctx context.Context, where *ent.AppActionWhereInput) ([]*ent.AppAction, error) {
	return r.resource.GetUserPermissions(ctx, where)
}

// CheckPermission is the resolver for the checkPermission field.
func (r *queryResolver) CheckPermission(ctx context.Context, permission string) (bool, error) {
	return r.resource.CheckPermission(ctx, permission)
}

// OrgAppActions is the resolver for the orgAppActions field.
func (r *queryResolver) OrgAppActions(ctx context.Context, appCode string) ([]*ent.AppAction, error) {
	//获取跟用户ID
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	rootOrg, err := r.resource.GetRootOrgByUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	// 获取根用户所有权限
	return r.resource.GetUserPermissionsByUserID(ctx, *rootOrg.OwnerID, &ent.AppActionWhereInput{
		HasAppWith: []*ent.AppWhereInput{{Code: &appCode}},
	})
}

// UserRootOrgs is the resolver for the userRootOrgs field.
func (r *queryResolver) UserRootOrgs(ctx context.Context) ([]*ent.Org, error) {
	//获取跟用户ID
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.Org.Query().Where(
		org.HasOrgUserWith(orguser.UserID(uid)),
		org.StatusEQ(typex.SimpleStatusActive),
		org.DomainNotNil(),
		org.KindEQ(org.KindRoot),
	).All(ctx)
}

// OrgRecycleUsers is the resolver for the orgRecycleUsers field.
func (r *queryResolver) OrgRecycleUsers(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	o, err := r.client.Org.Query().Where(org.ID(tid), org.DomainNotNil(), org.KindEQ(org.KindRoot)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.User.Query().Where(
		user.PrincipalNameHasSuffix("@"+o.Domain),
		user.DeletedAtNotNil(),
		user.StatusEQ(typex.SimpleStatusInactive),
	).Paginate(schemax.SkipSoftDelete(ctx), after, first, before, last, ent.WithUserOrder(orderBy), ent.WithUserFilter(where.Filter))
}

// OrgUserPreference is the resolver for the orgUserPreference field.
func (r *queryResolver) OrgUserPreference(ctx context.Context) (*ent.OrgUserPreference, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	oup, err := r.client.OrgUserPreference.Query().Where(orguserpreference.UserID(uid), orguserpreference.OrgID(tid)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return oup, nil
}

// UserApps is the resolver for the userApps field.
func (r *queryResolver) UserApps(ctx context.Context) ([]*ent.App, error) {
	return r.resource.GetUserApps(ctx)
}

// AppDictByRefCode is the resolver for the appDictByRefCode field.
func (r *queryResolver) AppDictByRefCode(ctx context.Context, refCodes []string) ([]*ent.AppDict, error) {
	return r.client.AppDict.Query().Where(
		appdict.HasItemsWith(
			appdictitem.RefCodeIn(refCodes...),
			appdictitem.StatusEQ(typex.SimpleStatusActive),
		)).
		WithNamedItems(appdict.EdgeItems, func(query *ent.AppDictItemQuery) {
			query.Where(appdictitem.StatusEQ(typex.SimpleStatusActive)).
				Order(appdictitem.ByDisplaySort(sql.OrderAsc()))
		}).All(ctx)
}

// AppDictItemByRefCode is the resolver for the appDictItemByRefCode field.
func (r *queryResolver) AppDictItemByRefCode(ctx context.Context, refCode string) ([]*ent.AppDictItem, error) {
	return r.client.AppDictItem.Query().Where(
		appdictitem.RefCode(refCode),
		appdictitem.StatusEQ(typex.SimpleStatusActive),
	).Order(appdictitem.ByDisplaySort(sql.OrderAsc())).All(ctx)
}
