package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/approlepolicy"
	"github.com/woocoos/knockout/ent/orgroleuser"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/permission"
)

// IsGrantAppRole is the resolver for the isGrantAppRole field.
func (r *appPolicyResolver) IsGrantAppRole(ctx context.Context, obj *ent.AppPolicy, appRoleID int) (bool, error) {
	exist, err := r.Client.AppRolePolicy.Query().Where(
		approlepolicy.AppID(obj.AppID), approlepolicy.AppRoleID(appRoleID), approlepolicy.AppPolicyID(obj.ID),
	).Exist(ctx)
	if err != nil {
		return false, err
	}
	return exist, nil
}

// IsAllowRevokeAppPolicy is the resolver for the isAllowRevokeAppPolicy field.
func (r *orgResolver) IsAllowRevokeAppPolicy(ctx context.Context, obj *ent.Org, appPolicyID int) (bool, error) {
	return r.Resource.IsAllowRevokeAppPolicy(ctx, obj.ID, appPolicyID)
}

// IsGrantRole is the resolver for the isGrantRole field.
func (r *orgPolicyResolver) IsGrantRole(ctx context.Context, obj *ent.OrgPolicy, roleID int) (bool, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	exist, err := r.Client.Permission.Query().Where(
		permission.PrincipalKindEQ(permission.PrincipalKindRole),
		permission.RoleID(roleID), permission.OrgPolicyID(obj.ID), permission.OrgID(tid),
	).Exist(ctx)
	if err != nil {
		return false, err
	}
	return exist, nil
}

// IsGrantUser is the resolver for the isGrantUser field.
func (r *orgPolicyResolver) IsGrantUser(ctx context.Context, obj *ent.OrgPolicy, userID int) (bool, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	exist, err := r.Client.Permission.Query().Where(
		permission.PrincipalKindEQ(permission.PrincipalKindUser),
		permission.UserID(userID), permission.OrgPolicyID(obj.ID), permission.OrgID(tid),
	).Exist(ctx)
	if err != nil {
		return false, err
	}
	return exist, nil
}

// IsAppRole is the resolver for the isAppRole field.
func (r *orgRoleResolver) IsAppRole(ctx context.Context, obj *ent.OrgRole) (bool, error) {
	return obj.AppRoleID > 0, nil
}

// IsGrantUser is the resolver for the isGrantUser field.
func (r *orgRoleResolver) IsGrantUser(ctx context.Context, obj *ent.OrgRole, userID int) (bool, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	has, err := r.Client.OrgRoleUser.Query().Where(
		orgroleuser.OrgRoleID(obj.ID),
		orgroleuser.HasOrgUserWith(orguser.UserID(userID), orguser.OrgID(tid)),
	).Exist(ctx)
	if err != nil {
		return false, err
	}
	return has, nil
}

// IsAllowRevoke is the resolver for the isAllowRevoke field.
func (r *permissionResolver) IsAllowRevoke(ctx context.Context, obj *ent.Permission) (bool, error) {
	return r.Resource.IsAllowRevokePermission(ctx, obj)
}

// IsAssignOrgRole is the resolver for the isAssignOrgRole field.
func (r *userResolver) IsAssignOrgRole(ctx context.Context, obj *ent.User, orgRoleID int) (bool, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	ouid, err := r.Client.OrgUser.Query().Where(orguser.OrgID(tid), orguser.UserID(obj.ID)).Select(orguser.FieldID).Int(ctx)
	if err != nil {
		return false, err
	}
	exist, err := r.Client.OrgRoleUser.Query().Where(orgroleuser.OrgUserID(ouid), orgroleuser.OrgRoleID(orgRoleID)).Exist(ctx)
	if err != nil {
		return false, err
	}
	return exist, nil
}

// IsAllowRevokeRole is the resolver for the isAllowRevokeRole field.
func (r *userResolver) IsAllowRevokeRole(ctx context.Context, obj *ent.User, orgRoleID int) (bool, error) {
	return r.Resource.IsAllowRevokeOrgRole(ctx, obj.ID, orgRoleID)
}

// LoginProfile is the resolver for the loginProfile field.
func (r *createUserInputResolver) LoginProfile(ctx context.Context, obj *ent.CreateUserInput, data *ent.CreateUserLoginProfileInput) error {
	if data != nil {
		row, err := ent.FromContext(ctx).UserLoginProfile.Create().SetInput(*data).Save(ctx)
		if err != nil {
			return err
		}
		obj.LoginProfileID = &row.ID
	}
	return nil
}

// Password is the resolver for the password field.
func (r *createUserInputResolver) Password(ctx context.Context, obj *ent.CreateUserInput, data *ent.CreateUserPasswordInput) error {
	if data != nil {
		row, err := r.Resource.CreateUserPassword(ctx, data)
		if err != nil {
			return err
		}
		obj.PasswordIDs = append(obj.PasswordIDs, row.ID)
	}
	return nil
}
