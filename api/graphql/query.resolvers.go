package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"
	"strconv"

	"entgo.io/contrib/entgql"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appres"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgapp"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orguser"
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
	return r.Client.OrgRole.Query().Where(orgrole.OrgID(tid), orgrole.KindEQ(orgrole.KindGroup)).Paginate(ctx, after, first, before, last,
		ent.WithOrgRoleOrder(orderBy), ent.WithOrgRoleFilter(where.Filter))
}

// OrgRoleUsers is the resolver for the orgRoleUsers field.
func (r *queryResolver) OrgRoleUsers(ctx context.Context, roleID int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	uIds, err := r.Resource.GetRoleUserIds(ctx, roleID)
	if err != nil {
		return nil, err
	}
	return r.Client.User.Query().Where(user.IDIn(uIds...)).Paginate(ctx, after, first, before, last,
		ent.WithUserOrder(orderBy), ent.WithUserFilter(where.Filter))
}

// OrgRoles is the resolver for the orgRoles field.
func (r *queryResolver) OrgRoles(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.OrgRoleOrder, where *ent.OrgRoleWhereInput) (*ent.OrgRoleConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Client.OrgRole.Query().Where(orgrole.OrgID(tid), orgrole.KindEQ(orgrole.KindRole)).Paginate(ctx, after, first, before, last,
		ent.WithOrgRoleOrder(orderBy), ent.WithOrgRoleFilter(where.Filter))
}

// AppRoleAssignedToOrgs is the resolver for the appRoleAssignedToOrgs field.
func (r *queryResolver) AppRoleAssignedToOrgs(ctx context.Context, roleID int, where *ent.OrgWhereInput) ([]*ent.Org, error) {
	oIds, err := r.Client.OrgRole.Query().Where(orgrole.AppRoleID(roleID)).Select(orgrole.FieldOrgID).Ints(ctx)
	if err != nil {
		return nil, err
	}
	q := r.Client.Org.Query()
	q, err = where.Filter(q)
	if err != nil {
		return nil, err
	}
	return q.Where(org.IDIn(oIds...)).All(ctx)
}

// AppPolicyAssignedToOrgs is the resolver for the appPolicyAssignedToOrgs field.
func (r *queryResolver) AppPolicyAssignedToOrgs(ctx context.Context, policyID int, where *ent.OrgWhereInput) ([]*ent.Org, error) {
	oIds, err := r.Client.OrgPolicy.Query().Where(orgpolicy.AppPolicyID(policyID)).Select(orgrole.FieldOrgID).Ints(ctx)
	if err != nil {
		return nil, err
	}
	q := r.Client.Org.Query()
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
	has, err := r.Client.OrgPolicy.Query().Where(orgpolicy.ID(policyID), orgpolicy.OrgID(tid)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("policy not exist")
	}
	return r.Client.Permission.Query().Where(permission.OrgID(tid), permission.OrgPolicyID(policyID)).Paginate(ctx, after, first, before, last, ent.WithPermissionOrder(orderBy), ent.WithPermissionFilter(where.Filter))
}

// AppResources is the resolver for the appResources field.
func (r *queryResolver) AppResources(ctx context.Context, appID int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.AppResOrder, where *ent.AppResWhereInput) (*ent.AppResConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Client.AppRes.Query().Where(appres.AppID(appID), appres.HasAppWith(app.OwnerOrgID(tid))).Paginate(ctx, after, first, before, last, ent.WithAppResOrder(orderBy), ent.WithAppResFilter(where.Filter))
}

// OrgAppResources is the resolver for the orgAppResources field.
func (r *queryResolver) OrgAppResources(ctx context.Context, appID int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.AppResOrder, where *ent.AppResWhereInput) (*ent.AppResConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Client.AppRes.Query().Where(appres.AppID(appID), appres.HasAppWith(app.HasOrgAppWith(orgapp.AppID(appID), orgapp.OrgID(tid)))).Paginate(ctx, after, first, before, last, ent.WithAppResOrder(orderBy), ent.WithAppResFilter(where.Filter))
}

// UserGroups is the resolver for the userGroups field.
func (r *queryResolver) UserGroups(ctx context.Context, userID int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.OrgRoleOrder, where *ent.OrgRoleWhereInput) (*ent.OrgRoleConnection, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.Client.OrgRole.Query().Where(orgrole.HasOrgUsersWith(orguser.OrgID(tid), orguser.UserID(userID)), orgrole.KindIn(orgrole.KindGroup)).
		Paginate(ctx, after, first, before, last, ent.WithOrgRoleOrder(orderBy), ent.WithOrgRoleFilter(where.Filter))
}
