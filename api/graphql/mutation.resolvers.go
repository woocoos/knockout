package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	generated1 "github.com/woocoos/knockout/api/graphql/generated"
	"github.com/woocoos/knockout/api/graphql/model"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/user"
)

// EnableDirectory is the resolver for the enableDirectory field.
func (r *mutationResolver) EnableDirectory(ctx context.Context, input model.EnableDirectoryInput) (*ent.Org, error) {
	return r.Resource.EnableOrganization(ctx, input)
}

// CreateOrganization is the resolver for the createOrganization field.
func (r *mutationResolver) CreateOrganization(ctx context.Context, input ent.CreateOrgInput) (*ent.Org, error) {
	return r.Resource.CreateOrganization(ctx, input)
}

// UpdateOrganization is the resolver for the updateOrganization field.
func (r *mutationResolver) UpdateOrganization(ctx context.Context, orgID int, input ent.UpdateOrgInput) (*ent.Org, error) {
	return r.Resource.UpdateOrganization(ctx, orgID, input)
}

// DeleteOrganization is the resolver for the deleteOrganization field.
func (r *mutationResolver) DeleteOrganization(ctx context.Context, orgID int) (bool, error) {
	err := r.Resource.DeleteOrganization(ctx, orgID)
	if err != nil {
		return false, err
	}
	return true, nil
}

// MoveOrganization is the resolver for the moveOrganization field.
func (r *mutationResolver) MoveOrganization(ctx context.Context, sourceID int, targetID int, action model.TreeAction) (bool, error) {
	err := r.Resource.MoveOrganization(ctx, sourceID, targetID, action)
	return err == nil, err
}

// CreateOrganizationAccount is the resolver for the createOrganizationAccount field.
func (r *mutationResolver) CreateOrganizationAccount(ctx context.Context, rootOrgID int, input ent.CreateUserInput) (*ent.User, error) {
	return r.Resource.CreateOrganizationAccount(ctx, rootOrgID, input)
}

// CreateOrganizationUser is the resolver for the createOrganizationUser field.
func (r *mutationResolver) CreateOrganizationUser(ctx context.Context, rootOrgID int, input ent.CreateUserInput) (*ent.User, error) {
	return r.Resource.CreateOrganizationUser(ctx, rootOrgID, input, user.UserTypeMember)
}

// AllotOrganizationUser is the resolver for the allotOrganizationUser field.
func (r *mutationResolver) AllotOrganizationUser(ctx context.Context, input ent.CreateOrgUserInput) (bool, error) {
	err := r.Resource.AllotOrganizationUser(ctx, input)
	return err == nil, err
}

// RemoveOrganizationUser is the resolver for the removeOrganizationUser field.
func (r *mutationResolver) RemoveOrganizationUser(ctx context.Context, orgID int, userID int) (bool, error) {
	err := r.Resource.RemoveOrganizationUser(ctx, orgID, userID)
	return err == nil, err
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, userID int, input ent.UpdateUserInput) (*ent.User, error) {
	return r.Resource.UpdateUser(ctx, userID, input)
}

// UpdateLoginProfile is the resolver for the updateLoginProfile field.
func (r *mutationResolver) UpdateLoginProfile(ctx context.Context, input ent.UpdateUserLoginProfileInput) (*ent.UserLoginProfile, error) {
	return r.Resource.UpdateLoginProfile(ctx, input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID int) (bool, error) {
	err := r.Resource.DeleteOrganizationUser(ctx, userID)
	return err == nil, err
}

// BindUserIdentity is the resolver for the bindUserIdentity field.
func (r *mutationResolver) BindUserIdentity(ctx context.Context, input ent.CreateUserIdentityInput) (*ent.UserIdentity, error) {
	return ent.FromContext(ctx).UserIdentity.Create().SetInput(input).Save(ctx)
}

// DeleteUserIdentity is the resolver for the deleteUserIdentity field.
func (r *mutationResolver) DeleteUserIdentity(ctx context.Context, id int) (bool, error) {
	err := ent.FromContext(ctx).UserIdentity.DeleteOneID(id).Exec(ctx)
	return err == nil, err
}

// ChangePassword is the resolver for the changePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context, oldPwd string, newPwd string) (bool, error) {
	err := r.Resource.ChangePassword(ctx, oldPwd, newPwd)
	return err == nil, err
}

// ResetUserPasswordByEmail is the resolver for the resetUserPasswordByEmail field.
func (r *mutationResolver) ResetUserPasswordByEmail(ctx context.Context, userID int) (bool, error) {
	panic(fmt.Errorf("not implemented: ResetUserPasswordByEmail - resetUserPasswordByEmail"))
}

// CreateApp is the resolver for the createApp field.
func (r *mutationResolver) CreateApp(ctx context.Context, input ent.CreateAppInput) (*ent.App, error) {
	return r.Resource.CreateApp(ctx, input)
}

// UpdateApp is the resolver for the updateApp field.
func (r *mutationResolver) UpdateApp(ctx context.Context, appID int, input ent.UpdateAppInput) (*ent.App, error) {
	return r.Resource.UpdateApp(ctx, appID, input)
}

// DeleteApp is the resolver for the deleteApp field.
func (r *mutationResolver) DeleteApp(ctx context.Context, appID int) (bool, error) {
	err := r.Resource.DeleteApp(ctx, appID)
	return err == nil, err
}

// CreateAppActions is the resolver for the createAppActions field.
func (r *mutationResolver) CreateAppActions(ctx context.Context, appID int, input []*ent.CreateAppActionInput) (bool, error) {
	err := r.Resource.CreateAppActions(ctx, appID, input)
	return err == nil, err
}

// CreateAppPolicies is the resolver for the createAppPolicies field.
func (r *mutationResolver) CreateAppPolicies(ctx context.Context, appID int, input []*ent.CreateAppPolicyInput) (bool, error) {
	err := r.Resource.CreateAppPolicies(ctx, appID, input)
	return err == nil, err
}

// UpdateAppPolicy is the resolver for the updateAppPolicy field.
func (r *mutationResolver) UpdateAppPolicy(ctx context.Context, policyID int, input ent.UpdateAppPolicyInput) (*ent.AppPolicy, error) {
	return r.Resource.UpdateAppPolicy(ctx, policyID, input)
}

// DeleteAppPolicy is the resolver for the deleteAppPolicy field.
func (r *mutationResolver) DeleteAppPolicy(ctx context.Context, policyID int) (bool, error) {
	err := r.Client.AppRolePolicy.DeleteOneID(policyID).Exec(ctx)
	return err == nil, err
}

// CreateAppMenus is the resolver for the createAppMenus field.
func (r *mutationResolver) CreateAppMenus(ctx context.Context, appID int, input []*ent.CreateAppMenuInput) (bool, error) {
	err := r.Resource.CreateAppMenus(ctx, appID, input)
	return err == nil, err
}

// UpdateAppMenu is the resolver for the updateAppMenu field.
func (r *mutationResolver) UpdateAppMenu(ctx context.Context, menuID int, input ent.UpdateAppMenuInput) (*ent.AppMenu, error) {
	return ent.FromContext(ctx).AppMenu.UpdateOneID(menuID).SetInput(input).Save(ctx)
}

// MoveAppMenu is the resolver for the moveAppMenu field.
func (r *mutationResolver) MoveAppMenu(ctx context.Context, sourceID int, targetID int, action model.TreeAction) (bool, error) {
	err := r.Resource.MoveAppMenu(ctx, sourceID, targetID, action)
	return err == nil, err
}

// DeleteAppMenu is the resolver for the deleteAppMenu field.
func (r *mutationResolver) DeleteAppMenu(ctx context.Context, menuID int) (bool, error) {
	err := ent.FromContext(ctx).AppMenu.DeleteOneID(menuID).Exec(ctx)
	return err == nil, err
}

// CreateAppRole is the resolver for the createAppRole field.
func (r *mutationResolver) CreateAppRole(ctx context.Context, appID int, input ent.CreateAppRoleInput) (*ent.AppRole, error) {
	return ent.FromContext(ctx).AppRole.Create().SetInput(input).Save(ctx)
}

// UpdateAppRole is the resolver for the updateAppRole field.
func (r *mutationResolver) UpdateAppRole(ctx context.Context, roleID int, input ent.UpdateAppRoleInput) (*ent.AppRole, error) {
	return r.Resource.UpdateAppRole(ctx, roleID, input)
}

// DeleteAppRole is the resolver for the deleteAppRole field.
func (r *mutationResolver) DeleteAppRole(ctx context.Context, roleID int) (bool, error) {
	err := r.Resource.DeleteAppRole(ctx, roleID)
	return err == nil, err
}

// AssignOrganizationApp is the resolver for the assignOrganizationApp field.
func (r *mutationResolver) AssignOrganizationApp(ctx context.Context, orgID int, appID int) (bool, error) {
	err := r.Resource.AssignOrganizationApp(ctx, orgID, appID)
	return err == nil, err
}

// RevokeOrganizationApp is the resolver for the revokeOrganizationApp field.
func (r *mutationResolver) RevokeOrganizationApp(ctx context.Context, orgID int, appID int) (bool, error) {
	err := r.Resource.RevokeOrganizationApp(ctx, orgID, appID)
	return err == nil, err
}

// AssignOrganizationAppPolicy is the resolver for the assignOrganizationAppPolicy field.
func (r *mutationResolver) AssignOrganizationAppPolicy(ctx context.Context, orgID int, appPolicyID int) (bool, error) {
	err := r.Resource.AssignOrganizationAppPolicy(ctx, orgID, appPolicyID)
	return err == nil, err
}

// RevokeOrganizationAppPolicy is the resolver for the revokeOrganizationAppPolicy field.
func (r *mutationResolver) RevokeOrganizationAppPolicy(ctx context.Context, orgID int, appPolicyID int) (bool, error) {
	err := r.Resource.RevokeOrganizationAppPolicy(ctx, orgID, appPolicyID)
	return err == nil, err
}

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, input ent.CreateOrgRoleInput) (*ent.OrgRole, error) {
	return r.Resource.CreateRole(ctx, input)
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, roleID int, input ent.UpdateOrgRoleInput) (*ent.OrgRole, error) {
	return r.Resource.UpdateRole(ctx, roleID, input)
}

// DeleteRole is the resolver for the deleteRole field.
func (r *mutationResolver) DeleteRole(ctx context.Context, roleID int) (bool, error) {
	err := r.Resource.DeleteRole(ctx, roleID)
	return err == nil, err
}

// AssignRoleUser is the resolver for the assignRoleUser field.
func (r *mutationResolver) AssignRoleUser(ctx context.Context, input model.AssignRoleUserInput) (bool, error) {
	err := r.Resource.AssignRoleUser(ctx, input)
	return err == nil, err
}

// RevokeRoleUser is the resolver for the revokeRoleUser field.
func (r *mutationResolver) RevokeRoleUser(ctx context.Context, roleID int, userID int) (bool, error) {
	err := r.Resource.RevokeRoleUser(ctx, roleID, userID)
	return err == nil, err
}

// Grant is the resolver for the grant field.
func (r *mutationResolver) Grant(ctx context.Context, input ent.CreatePermissionInput) (*ent.Permission, error) {
	return r.Resource.Grant(ctx, input)
}

// UpdatePermission is the resolver for the updatePermission field.
func (r *mutationResolver) UpdatePermission(ctx context.Context, permissionID int, input ent.UpdatePermissionInput) (*ent.Permission, error) {
	return r.Resource.UpdatePermission(ctx, permissionID, input)
}

// Revoke is the resolver for the revoke field.
func (r *mutationResolver) Revoke(ctx context.Context, orgID int, permissionID int) (bool, error) {
	err := r.Resource.Revoke(ctx, orgID, permissionID)
	return err == nil, err
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
