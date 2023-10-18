package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"fmt"

	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/entco/schemax/typex"
	generated1 "github.com/woocoos/knockout/api/graphql/generated"
	"github.com/woocoos/knockout/api/graphql/model"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/oauthclient"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/service/resource"
)

// EnableDirectory is the resolver for the enableDirectory field.
func (r *mutationResolver) EnableDirectory(ctx context.Context, input model.EnableDirectoryInput) (*ent.Org, error) {
	return r.resource.EnableOrganization(ctx, input)
}

// CreateRoot is the resolver for the createRoot field.
func (r *mutationResolver) CreateRoot(ctx context.Context, input ent.CreateOrgInput) (*ent.Org, error) {
	return r.resource.CreateRoot(ctx, input)
}

// CreateOrganization is the resolver for the createOrganization field.
func (r *mutationResolver) CreateOrganization(ctx context.Context, input ent.CreateOrgInput) (*ent.Org, error) {
	return r.resource.CreateOrganization(ctx, input)
}

// UpdateOrganization is the resolver for the updateOrganization field.
func (r *mutationResolver) UpdateOrganization(ctx context.Context, orgID int, input ent.UpdateOrgInput) (*ent.Org, error) {
	return r.client.Org.UpdateOneID(orgID).SetInput(input).Save(ctx)
}

// DeleteOrganization is the resolver for the deleteOrganization field.
func (r *mutationResolver) DeleteOrganization(ctx context.Context, orgID int) (bool, error) {
	err := r.resource.DeleteOrganization(ctx, orgID)
	if err != nil {
		return false, err
	}
	return true, nil
}

// MoveOrganization is the resolver for the moveOrganization field.
func (r *mutationResolver) MoveOrganization(ctx context.Context, sourceID int, targetID int, action model.TreeAction) (bool, error) {
	err := r.resource.MoveOrganization(ctx, sourceID, targetID, action)
	return err == nil, err
}

// CreateOrganizationAccount is the resolver for the createOrganizationAccount field.
func (r *mutationResolver) CreateOrganizationAccount(ctx context.Context, rootOrgID int, input ent.CreateUserInput) (*ent.User, error) {
	return r.resource.CreateOrganizationAccount(ctx, rootOrgID, input)
}

// CreateOrganizationUser is the resolver for the createOrganizationUser field.
func (r *mutationResolver) CreateOrganizationUser(ctx context.Context, rootOrgID int, input ent.CreateUserInput) (*ent.User, error) {
	return r.resource.CreateOrganizationUser(ctx, rootOrgID, input, user.UserTypeMember)
}

// AllotOrganizationUser is the resolver for the allotOrganizationUser field.
func (r *mutationResolver) AllotOrganizationUser(ctx context.Context, input ent.CreateOrgUserInput) (bool, error) {
	err := r.resource.AllotOrganizationUser(ctx, input)
	return err == nil, err
}

// RemoveOrganizationUser is the resolver for the removeOrganizationUser field.
func (r *mutationResolver) RemoveOrganizationUser(ctx context.Context, orgID int, userID int) (bool, error) {
	err := r.resource.RemoveOrganizationUser(ctx, orgID, userID)
	return err == nil, err
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, userID int, input ent.UpdateUserInput) (*ent.User, error) {
	return r.resource.UpdateUser(ctx, userID, input)
}

// UpdateLoginProfile is the resolver for the updateLoginProfile field.
func (r *mutationResolver) UpdateLoginProfile(ctx context.Context, userID int, input ent.UpdateUserLoginProfileInput) (*ent.UserLoginProfile, error) {
	return r.resource.UpdateLoginProfile(ctx, userID, input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID int) (bool, error) {
	err := r.resource.DeleteOrganizationUser(ctx, userID)
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
	err := r.resource.ChangePassword(ctx, oldPwd, newPwd)
	return err == nil, err
}

// ResetUserPasswordByEmail is the resolver for the resetUserPasswordByEmail field.
func (r *mutationResolver) ResetUserPasswordByEmail(ctx context.Context, userID int) (bool, error) {
	err := r.resource.ResetUserPasswordByEmail(ctx, userID)
	return err == nil, err
}

// CreateApp is the resolver for the createApp field.
func (r *mutationResolver) CreateApp(ctx context.Context, input ent.CreateAppInput) (*ent.App, error) {
	return r.resource.CreateApp(ctx, input)
}

// UpdateApp is the resolver for the updateApp field.
func (r *mutationResolver) UpdateApp(ctx context.Context, appID int, input ent.UpdateAppInput) (*ent.App, error) {
	return r.resource.UpdateApp(ctx, appID, input)
}

// DeleteApp is the resolver for the deleteApp field.
func (r *mutationResolver) DeleteApp(ctx context.Context, appID int) (bool, error) {
	err := r.client.App.DeleteOneID(appID).Exec(ctx)
	return err == nil, err
}

// CreateAppActions is the resolver for the createAppActions field.
func (r *mutationResolver) CreateAppActions(ctx context.Context, appID int, input []*ent.CreateAppActionInput) ([]*ent.AppAction, error) {
	return r.resource.CreateAppActions(ctx, appID, input)
}

// UpdateAppAction is the resolver for the updateAppAction field.
func (r *mutationResolver) UpdateAppAction(ctx context.Context, actionID int, input ent.UpdateAppActionInput) (*ent.AppAction, error) {
	//return r.client.AppAction.UpdateOneID(actionID).SetInput(input).Save(ctx)
	return r.resource.UpdateAppAction(ctx, actionID, input)
}

// DeleteAppAction is the resolver for the deleteAppAction field.
func (r *mutationResolver) DeleteAppAction(ctx context.Context, actionID int) (bool, error) {
	err := r.resource.DeleteAppAction(ctx, actionID)
	return err == nil, err
}

// CreateAppPolicy is the resolver for the createAppPolicy field.
func (r *mutationResolver) CreateAppPolicy(ctx context.Context, appID int, input ent.CreateAppPolicyInput) (*ent.AppPolicy, error) {
	return r.resource.CreateAppPolicy(ctx, appID, input)
}

// UpdateAppPolicy is the resolver for the updateAppPolicy field.
func (r *mutationResolver) UpdateAppPolicy(ctx context.Context, policyID int, input ent.UpdateAppPolicyInput) (*ent.AppPolicy, error) {
	return r.resource.UpdateAppPolicy(ctx, policyID, input)
}

// DeleteAppPolicy is the resolver for the deleteAppPolicy field.
func (r *mutationResolver) DeleteAppPolicy(ctx context.Context, policyID int) (bool, error) {
	err := r.resource.DeleteAppPolicy(ctx, policyID)
	return err == nil, err
}

// CreateAppMenus is the resolver for the createAppMenus field.
func (r *mutationResolver) CreateAppMenus(ctx context.Context, appID int, input []*ent.CreateAppMenuInput) ([]*ent.AppMenu, error) {
	return r.resource.CreateAppMenus(ctx, appID, input)
}

// UpdateAppMenu is the resolver for the updateAppMenu field.
func (r *mutationResolver) UpdateAppMenu(ctx context.Context, menuID int, input ent.UpdateAppMenuInput) (*ent.AppMenu, error) {
	return r.resource.UpdateAppMenu(ctx, menuID, input)
}

// MoveAppMenu is the resolver for the moveAppMenu field.
func (r *mutationResolver) MoveAppMenu(ctx context.Context, sourceID int, targetID int, action model.TreeAction) (bool, error) {
	err := r.resource.MoveAppMenu(ctx, sourceID, targetID, action)
	return err == nil, err
}

// DeleteAppMenu is the resolver for the deleteAppMenu field.
func (r *mutationResolver) DeleteAppMenu(ctx context.Context, menuID int) (bool, error) {
	err := r.resource.DeleteAppMenu(ctx, menuID)
	return err == nil, err
}

// CreateAppRole is the resolver for the createAppRole field.
func (r *mutationResolver) CreateAppRole(ctx context.Context, appID int, input ent.CreateAppRoleInput) (*ent.AppRole, error) {
	return ent.FromContext(ctx).AppRole.Create().SetInput(input).Save(ctx)
}

// UpdateAppRole is the resolver for the updateAppRole field.
func (r *mutationResolver) UpdateAppRole(ctx context.Context, roleID int, input ent.UpdateAppRoleInput) (*ent.AppRole, error) {
	return r.resource.UpdateAppRole(ctx, roleID, input)
}

// DeleteAppRole is the resolver for the deleteAppRole field.
func (r *mutationResolver) DeleteAppRole(ctx context.Context, roleID int) (bool, error) {
	err := r.resource.DeleteAppRole(ctx, roleID)
	return err == nil, err
}

// CreateAppDict is the resolver for the createAppDict field.
func (r *mutationResolver) CreateAppDict(ctx context.Context, appID int, input ent.CreateAppDictInput) (*ent.AppDict, error) {
	return ent.FromContext(ctx).AppDict.Create().SetAppID(appID).SetInput(input).Save(ctx)
}

// UpdateAppDict is the resolver for the updateAppDict field.
func (r *mutationResolver) UpdateAppDict(ctx context.Context, dictID int, input ent.UpdateAppDictInput) (*ent.AppDict, error) {
	return ent.FromContext(ctx).AppDict.UpdateOneID(dictID).SetInput(input).Save(ctx)
}

// DeleteAppDict is the resolver for the deleteAppDict field.
func (r *mutationResolver) DeleteAppDict(ctx context.Context, dictID int) (bool, error) {
	err := ent.FromContext(ctx).AppDict.DeleteOneID(dictID).Exec(ctx)
	return err == nil, err
}

// CreateAppDictItem is the resolver for the createAppDictItem field.
func (r *mutationResolver) CreateAppDictItem(ctx context.Context, dictID int, input ent.CreateAppDictItemInput) (*ent.AppDictItem, error) {
	return ent.FromContext(ctx).AppDictItem.Create().SetDictID(dictID).SetInput(input).Save(ctx)
}

// UpdateAppDictItem is the resolver for the updateAppDictItem field.
func (r *mutationResolver) UpdateAppDictItem(ctx context.Context, itemID int, input ent.UpdateAppDictItemInput) (*ent.AppDictItem, error) {
	return ent.FromContext(ctx).AppDictItem.UpdateOneID(itemID).SetInput(input).Save(ctx)
}

// DeleteAppDictItem is the resolver for the deleteAppDictItem field.
func (r *mutationResolver) DeleteAppDictItem(ctx context.Context, itemID int) (bool, error) {
	err := ent.FromContext(ctx).AppDictItem.DeleteOneID(itemID).Exec(ctx)
	return err == nil, err
}

// MoveAppDictItem is the resolver for the moveAppDictItem field.
func (r *mutationResolver) MoveAppDictItem(ctx context.Context, sourceID int, targetID int, action model.TreeAction) (bool, error) {
	err := r.resource.MoveAppDictItem(ctx, sourceID, targetID, action)
	return err == nil, err
}

// AssignOrganizationAppRole is the resolver for the assignOrganizationAppRole field.
func (r *mutationResolver) AssignOrganizationAppRole(ctx context.Context, orgID int, appRoleID int) (bool, error) {
	err := r.resource.AssignOrganizationAppRole(ctx, orgID, appRoleID)
	return err == nil, err
}

// RevokeOrganizationAppRole is the resolver for the revokeOrganizationAppRole field.
func (r *mutationResolver) RevokeOrganizationAppRole(ctx context.Context, orgID int, appRoleID int) (bool, error) {
	err := r.resource.RevokeOrganizationAppRole(ctx, orgID, appRoleID)
	return err == nil, err
}

// AssignAppRolePolicy is the resolver for the assignAppRolePolicy field.
func (r *mutationResolver) AssignAppRolePolicy(ctx context.Context, appID int, roleID int, policyIDs []int) (bool, error) {
	err := r.resource.AssignAppRolePolicy(ctx, appID, roleID, policyIDs)
	return err == nil, err
}

// RevokeAppRolePolicy is the resolver for the revokeAppRolePolicy field.
func (r *mutationResolver) RevokeAppRolePolicy(ctx context.Context, appID int, roleID int, policyIDs []int) (bool, error) {
	err := r.resource.RevokeAppRolePolicy(ctx, appID, roleID, policyIDs)
	return err == nil, err
}

// AssignOrganizationApp is the resolver for the assignOrganizationApp field.
func (r *mutationResolver) AssignOrganizationApp(ctx context.Context, orgID int, appID int) (bool, error) {
	err := r.resource.AssignOrganizationApp(ctx, orgID, appID)
	return err == nil, err
}

// RevokeOrganizationApp is the resolver for the revokeOrganizationApp field.
func (r *mutationResolver) RevokeOrganizationApp(ctx context.Context, orgID int, appID int) (bool, error) {
	err := r.resource.RevokeOrganizationApp(ctx, orgID, appID)
	return err == nil, err
}

// AssignOrganizationAppPolicy is the resolver for the assignOrganizationAppPolicy field.
func (r *mutationResolver) AssignOrganizationAppPolicy(ctx context.Context, orgID int, appPolicyID int) (bool, error) {
	err := r.resource.AssignOrganizationAppPolicy(ctx, orgID, appPolicyID)
	return err == nil, err
}

// RevokeOrganizationAppPolicy is the resolver for the revokeOrganizationAppPolicy field.
func (r *mutationResolver) RevokeOrganizationAppPolicy(ctx context.Context, orgID int, appPolicyID int) (bool, error) {
	err := r.resource.RevokeOrganizationAppPolicy(ctx, orgID, appPolicyID)
	return err == nil, err
}

// CreateOrganizationPolicy is the resolver for the createOrganizationPolicy field.
func (r *mutationResolver) CreateOrganizationPolicy(ctx context.Context, input ent.CreateOrgPolicyInput) (*ent.OrgPolicy, error) {
	return r.resource.CreateOrganizationPolicy(ctx, input)
}

// UpdateOrganizationPolicy is the resolver for the updateOrganizationPolicy field.
func (r *mutationResolver) UpdateOrganizationPolicy(ctx context.Context, orgPolicyID int, input ent.UpdateOrgPolicyInput) (*ent.OrgPolicy, error) {
	return r.resource.UpdateOrganizationPolicy(ctx, orgPolicyID, input)
}

// DeleteOrganizationPolicy is the resolver for the deleteOrganizationPolicy field.
func (r *mutationResolver) DeleteOrganizationPolicy(ctx context.Context, orgPolicyID int) (bool, error) {
	err := r.resource.DeleteOrganizationPolicy(ctx, orgPolicyID)
	return err == nil, err
}

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, input ent.CreateOrgRoleInput) (*ent.OrgRole, error) {
	return r.resource.CreateRole(ctx, input)
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, roleID int, input ent.UpdateOrgRoleInput) (*ent.OrgRole, error) {
	return r.resource.UpdateRole(ctx, roleID, input)
}

// DeleteRole is the resolver for the deleteRole field.
func (r *mutationResolver) DeleteRole(ctx context.Context, roleID int) (bool, error) {
	err := r.resource.DeleteRole(ctx, roleID)
	return err == nil, err
}

// AssignRoleUser is the resolver for the assignRoleUser field.
func (r *mutationResolver) AssignRoleUser(ctx context.Context, input model.AssignRoleUserInput) (bool, error) {
	err := r.resource.AssignRoleUser(ctx, input)
	return err == nil, err
}

// RevokeRoleUser is the resolver for the revokeRoleUser field.
func (r *mutationResolver) RevokeRoleUser(ctx context.Context, roleID int, userID int) (bool, error) {
	err := r.resource.RevokeRoleUser(ctx, roleID, userID)
	return err == nil, err
}

// Grant is the resolver for the grant field.
func (r *mutationResolver) Grant(ctx context.Context, input ent.CreatePermissionInput) (*ent.Permission, error) {
	return r.resource.Grant(ctx, input)
}

// UpdatePermission is the resolver for the updatePermission field.
func (r *mutationResolver) UpdatePermission(ctx context.Context, permissionID int, input ent.UpdatePermissionInput) (*ent.Permission, error) {
	return r.resource.UpdatePermission(ctx, permissionID, input)
}

// Revoke is the resolver for the revoke field.
func (r *mutationResolver) Revoke(ctx context.Context, orgID int, permissionID int) (bool, error) {
	err := r.resource.Revoke(ctx, orgID, permissionID)
	return err == nil, err
}

// EnableMfa is the resolver for the enableMFA field.
func (r *mutationResolver) EnableMfa(ctx context.Context, userID int) (*model.Mfa, error) {
	return r.resource.EnableMFA(ctx, userID)
}

// DisableMfa is the resolver for the disableMFA field.
func (r *mutationResolver) DisableMfa(ctx context.Context, userID int) (bool, error) {
	err := r.resource.DisableMFA(ctx, userID)
	return err == nil, err
}

// SendMFAToUserByEmail is the resolver for the sendMFAToUserByEmail field.
func (r *mutationResolver) SendMFAToUserByEmail(ctx context.Context, userID int) (bool, error) {
	err := r.resource.SendMFAToUserByEmail(ctx, userID)
	return err == nil, err
}

// UpdateAppRes is the resolver for the updateAppRes field.
func (r *mutationResolver) UpdateAppRes(ctx context.Context, appResID int, input ent.UpdateAppResInput) (*ent.AppRes, error) {
	return ent.FromContext(ctx).AppRes.UpdateOneID(appResID).SetInput(input).Save(ctx)
}

// RecoverOrgUser is the resolver for the recoverOrgUser field.
func (r *mutationResolver) RecoverOrgUser(ctx context.Context, userID int, userInput ent.UpdateUserInput, pwdKind userloginprofile.SetKind, pwdInput *ent.CreateUserPasswordInput) (*ent.User, error) {
	return r.resource.RecoverOrgUser(ctx, userID, userInput, pwdKind, pwdInput)
}

// CreateFileSource is the resolver for the createFileSource field.
func (r *mutationResolver) CreateFileSource(ctx context.Context, input ent.CreateFileSourceInput) (*ent.FileSource, error) {
	return ent.FromContext(ctx).FileSource.Create().SetInput(input).Save(ctx)
}

// UpdateFileSource is the resolver for the updateFileSource field.
func (r *mutationResolver) UpdateFileSource(ctx context.Context, fsID int, input ent.UpdateFileSourceInput) (*ent.FileSource, error) {
	return ent.FromContext(ctx).FileSource.UpdateOneID(fsID).SetInput(input).Save(ctx)
}

// DeleteFileSource is the resolver for the deleteFileSource field.
func (r *mutationResolver) DeleteFileSource(ctx context.Context, fsID int) (bool, error) {
	client := ent.FromContext(ctx)
	has, err := client.FileSource.Query().Where(filesource.ID(fsID), filesource.HasFiles()).Exist(ctx)
	if err != nil {
		return false, err
	}
	if has {
		return false, fmt.Errorf("filesource: %d has be referenced, cannot be deleted", fsID)
	}
	err = client.FileSource.DeleteOneID(fsID).Exec(ctx)
	return err == nil, err
}

// CreateOauthClient is the resolver for the createOauthClient field.
func (r *mutationResolver) CreateOauthClient(ctx context.Context, input ent.CreateOauthClientInput) (*ent.OauthClient, error) {
	clientId := snowflake.New().String()
	clientSecret := resource.RandomStr(32)
	return ent.FromContext(ctx).OauthClient.Create().SetInput(input).SetClientID(clientId).SetClientSecret(clientSecret).Save(ctx)
}

// EnableOauthClient is the resolver for the enableOauthClient field.
func (r *mutationResolver) EnableOauthClient(ctx context.Context, id int) (*ent.OauthClient, error) {
	return ent.FromContext(ctx).OauthClient.UpdateOneID(id).SetStatus(typex.SimpleStatusActive).Save(ctx)
}

// DisableOauthClient is the resolver for the disableOauthClient field.
func (r *mutationResolver) DisableOauthClient(ctx context.Context, id int) (*ent.OauthClient, error) {
	return ent.FromContext(ctx).OauthClient.UpdateOneID(id).SetStatus(typex.SimpleStatusInactive).Save(ctx)
}

// DeleteOauthClient is the resolver for the deleteOauthClient field.
func (r *mutationResolver) DeleteOauthClient(ctx context.Context, id int) (bool, error) {
	client := ent.FromContext(ctx)
	oc, err := client.OauthClient.Query().Where(oauthclient.ID(id)).Only(ctx)
	if err != nil {
		return false, err
	}
	if oc.Status == typex.SimpleStatusActive {
		return false, fmt.Errorf("the active status cannot be deleted")
	}
	err = client.OauthClient.DeleteOneID(id).Exec(ctx)
	return err == nil, err
}

// SaveOrgUserPreference is the resolver for the saveOrgUserPreference field.
func (r *mutationResolver) SaveOrgUserPreference(ctx context.Context, input model.OrgUserPreferenceInput) (*ent.OrgUserPreference, error) {
	return r.resource.SaveOrgUserPreference(ctx, input)
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
