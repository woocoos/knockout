// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
	"github.com/woocoos/knockout/graph/entgen/types"
)

// CreateAppInput represents a mutation input for creating apps.
type CreateAppInput struct {
	ID                   *int
	Name                 string
	Code                 string
	Kind                 app.Kind
	RedirectURI          *string
	AppKey               *string
	AppSecret            *string
	Scopes               *string
	TokenValidity        *int32
	RefreshTokenValidity *int32
	Logo                 *string
	Comments             *string
	Status               *typex.SimpleStatus
	MenuIDs              []int
	ActionIDs            []int
	ResourceIDs          []int
	RoleIDs              []int
	PolicyIDs            []int
	OrgIDs               []int
}

// Mutate applies the CreateAppInput on the AppMutation builder.
func (i *CreateAppInput) Mutate(m *AppMutation) {
	m.SetName(i.Name)
	m.SetCode(i.Code)
	m.SetKind(i.Kind)
	if v := i.RedirectURI; v != nil {
		m.SetRedirectURI(*v)
	}
	if v := i.AppKey; v != nil {
		m.SetAppKey(*v)
	}
	if v := i.AppSecret; v != nil {
		m.SetAppSecret(*v)
	}
	if v := i.Scopes; v != nil {
		m.SetScopes(*v)
	}
	if v := i.TokenValidity; v != nil {
		m.SetTokenValidity(*v)
	}
	if v := i.RefreshTokenValidity; v != nil {
		m.SetRefreshTokenValidity(*v)
	}
	if v := i.Logo; v != nil {
		m.SetLogo(*v)
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.MenuIDs; len(v) > 0 {
		m.AddMenuIDs(v...)
	}
	if v := i.ActionIDs; len(v) > 0 {
		m.AddActionIDs(v...)
	}
	if v := i.ResourceIDs; len(v) > 0 {
		m.AddResourceIDs(v...)
	}
	if v := i.RoleIDs; len(v) > 0 {
		m.AddRoleIDs(v...)
	}
	if v := i.PolicyIDs; len(v) > 0 {
		m.AddPolicyIDs(v...)
	}
	if v := i.OrgIDs; len(v) > 0 {
		m.AddOrgIDs(v...)
	}
}

// SetInput applies the change-set in the CreateAppInput on the AppCreate builder.
func (c *AppCreate) SetInput(i CreateAppInput) *AppCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateAppInput represents a mutation input for updating apps.
type UpdateAppInput struct {
	Name                      *string
	Code                      *string
	Kind                      *app.Kind
	ClearRedirectURI          bool
	RedirectURI               *string
	ClearAppKey               bool
	AppKey                    *string
	ClearAppSecret            bool
	AppSecret                 *string
	ClearScopes               bool
	Scopes                    *string
	ClearTokenValidity        bool
	TokenValidity             *int32
	ClearRefreshTokenValidity bool
	RefreshTokenValidity      *int32
	ClearLogo                 bool
	Logo                      *string
	ClearComments             bool
	Comments                  *string
	ClearStatus               bool
	Status                    *typex.SimpleStatus
	ClearMenus                bool
	AddMenuIDs                []int
	RemoveMenuIDs             []int
	ClearActions              bool
	AddActionIDs              []int
	RemoveActionIDs           []int
	ClearResources            bool
	AddResourceIDs            []int
	RemoveResourceIDs         []int
	ClearRoles                bool
	AddRoleIDs                []int
	RemoveRoleIDs             []int
	ClearPolicies             bool
	AddPolicyIDs              []int
	RemovePolicyIDs           []int
	ClearOrgs                 bool
	AddOrgIDs                 []int
	RemoveOrgIDs              []int
}

// Mutate applies the UpdateAppInput on the AppMutation builder.
func (i *UpdateAppInput) Mutate(m *AppMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Code; v != nil {
		m.SetCode(*v)
	}
	if v := i.Kind; v != nil {
		m.SetKind(*v)
	}
	if i.ClearRedirectURI {
		m.ClearRedirectURI()
	}
	if v := i.RedirectURI; v != nil {
		m.SetRedirectURI(*v)
	}
	if i.ClearAppKey {
		m.ClearAppKey()
	}
	if v := i.AppKey; v != nil {
		m.SetAppKey(*v)
	}
	if i.ClearAppSecret {
		m.ClearAppSecret()
	}
	if v := i.AppSecret; v != nil {
		m.SetAppSecret(*v)
	}
	if i.ClearScopes {
		m.ClearScopes()
	}
	if v := i.Scopes; v != nil {
		m.SetScopes(*v)
	}
	if i.ClearTokenValidity {
		m.ClearTokenValidity()
	}
	if v := i.TokenValidity; v != nil {
		m.SetTokenValidity(*v)
	}
	if i.ClearRefreshTokenValidity {
		m.ClearRefreshTokenValidity()
	}
	if v := i.RefreshTokenValidity; v != nil {
		m.SetRefreshTokenValidity(*v)
	}
	if i.ClearLogo {
		m.ClearLogo()
	}
	if v := i.Logo; v != nil {
		m.SetLogo(*v)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if i.ClearStatus {
		m.ClearStatus()
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearMenus {
		m.ClearMenus()
	}
	if v := i.AddMenuIDs; len(v) > 0 {
		m.AddMenuIDs(v...)
	}
	if v := i.RemoveMenuIDs; len(v) > 0 {
		m.RemoveMenuIDs(v...)
	}
	if i.ClearActions {
		m.ClearActions()
	}
	if v := i.AddActionIDs; len(v) > 0 {
		m.AddActionIDs(v...)
	}
	if v := i.RemoveActionIDs; len(v) > 0 {
		m.RemoveActionIDs(v...)
	}
	if i.ClearResources {
		m.ClearResources()
	}
	if v := i.AddResourceIDs; len(v) > 0 {
		m.AddResourceIDs(v...)
	}
	if v := i.RemoveResourceIDs; len(v) > 0 {
		m.RemoveResourceIDs(v...)
	}
	if i.ClearRoles {
		m.ClearRoles()
	}
	if v := i.AddRoleIDs; len(v) > 0 {
		m.AddRoleIDs(v...)
	}
	if v := i.RemoveRoleIDs; len(v) > 0 {
		m.RemoveRoleIDs(v...)
	}
	if i.ClearPolicies {
		m.ClearPolicies()
	}
	if v := i.AddPolicyIDs; len(v) > 0 {
		m.AddPolicyIDs(v...)
	}
	if v := i.RemovePolicyIDs; len(v) > 0 {
		m.RemovePolicyIDs(v...)
	}
	if i.ClearOrgs {
		m.ClearOrgs()
	}
	if v := i.AddOrgIDs; len(v) > 0 {
		m.AddOrgIDs(v...)
	}
	if v := i.RemoveOrgIDs; len(v) > 0 {
		m.RemoveOrgIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateAppInput on the AppUpdate builder.
func (c *AppUpdate) SetInput(i UpdateAppInput) *AppUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateAppInput on the AppUpdateOne builder.
func (c *AppUpdateOne) SetInput(i UpdateAppInput) *AppUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateAppActionInput represents a mutation input for creating appactions.
type CreateAppActionInput struct {
	ID          *int
	Name        string
	Kind        appaction.Kind
	Method      appaction.Method
	Comments    *string
	AppID       int
	MenuIDs     []int
	ResourceIDs []int
}

// Mutate applies the CreateAppActionInput on the AppActionMutation builder.
func (i *CreateAppActionInput) Mutate(m *AppActionMutation) {
	m.SetName(i.Name)
	m.SetKind(i.Kind)
	m.SetMethod(i.Method)
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	m.SetAppID(i.AppID)
	if v := i.MenuIDs; len(v) > 0 {
		m.AddMenuIDs(v...)
	}
	if v := i.ResourceIDs; len(v) > 0 {
		m.AddResourceIDs(v...)
	}
}

// SetInput applies the change-set in the CreateAppActionInput on the AppActionCreate builder.
func (c *AppActionCreate) SetInput(i CreateAppActionInput) *AppActionCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateAppActionInput represents a mutation input for updating appactions.
type UpdateAppActionInput struct {
	Name              *string
	Kind              *appaction.Kind
	Method            *appaction.Method
	ClearComments     bool
	Comments          *string
	ClearApp          bool
	AppID             *int
	ClearMenus        bool
	AddMenuIDs        []int
	RemoveMenuIDs     []int
	ClearResources    bool
	AddResourceIDs    []int
	RemoveResourceIDs []int
}

// Mutate applies the UpdateAppActionInput on the AppActionMutation builder.
func (i *UpdateAppActionInput) Mutate(m *AppActionMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Kind; v != nil {
		m.SetKind(*v)
	}
	if v := i.Method; v != nil {
		m.SetMethod(*v)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if i.ClearApp {
		m.ClearApp()
	}
	if v := i.AppID; v != nil {
		m.SetAppID(*v)
	}
	if i.ClearMenus {
		m.ClearMenus()
	}
	if v := i.AddMenuIDs; len(v) > 0 {
		m.AddMenuIDs(v...)
	}
	if v := i.RemoveMenuIDs; len(v) > 0 {
		m.RemoveMenuIDs(v...)
	}
	if i.ClearResources {
		m.ClearResources()
	}
	if v := i.AddResourceIDs; len(v) > 0 {
		m.AddResourceIDs(v...)
	}
	if v := i.RemoveResourceIDs; len(v) > 0 {
		m.RemoveResourceIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateAppActionInput on the AppActionUpdate builder.
func (c *AppActionUpdate) SetInput(i UpdateAppActionInput) *AppActionUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateAppActionInput on the AppActionUpdateOne builder.
func (c *AppActionUpdateOne) SetInput(i UpdateAppActionInput) *AppActionUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateAppMenuInput represents a mutation input for creating appmenus.
type CreateAppMenuInput struct {
	ID       *int
	ParentID int
	Kind     appmenu.Kind
	Name     *string
	Comments *string
	AppID    int
	ActionID *int
}

// Mutate applies the CreateAppMenuInput on the AppMenuMutation builder.
func (i *CreateAppMenuInput) Mutate(m *AppMenuMutation) {
	m.SetParentID(i.ParentID)
	m.SetKind(i.Kind)
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	m.SetAppID(i.AppID)
	if v := i.ActionID; v != nil {
		m.SetActionID(*v)
	}
}

// SetInput applies the change-set in the CreateAppMenuInput on the AppMenuCreate builder.
func (c *AppMenuCreate) SetInput(i CreateAppMenuInput) *AppMenuCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateAppMenuInput represents a mutation input for updating appmenus.
type UpdateAppMenuInput struct {
	ParentID      *int
	Kind          *appmenu.Kind
	ClearName     bool
	Name          *string
	ClearComments bool
	Comments      *string
	ClearApp      bool
	AppID         *int
	ClearAction   bool
	ActionID      *int
}

// Mutate applies the UpdateAppMenuInput on the AppMenuMutation builder.
func (i *UpdateAppMenuInput) Mutate(m *AppMenuMutation) {
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if v := i.Kind; v != nil {
		m.SetKind(*v)
	}
	if i.ClearName {
		m.ClearName()
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if i.ClearApp {
		m.ClearApp()
	}
	if v := i.AppID; v != nil {
		m.SetAppID(*v)
	}
	if i.ClearAction {
		m.ClearAction()
	}
	if v := i.ActionID; v != nil {
		m.SetActionID(*v)
	}
}

// SetInput applies the change-set in the UpdateAppMenuInput on the AppMenuUpdate builder.
func (c *AppMenuUpdate) SetInput(i UpdateAppMenuInput) *AppMenuUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateAppMenuInput on the AppMenuUpdateOne builder.
func (c *AppMenuUpdateOne) SetInput(i UpdateAppMenuInput) *AppMenuUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateAppPolicyInput represents a mutation input for creating apppolicies.
type CreateAppPolicyInput struct {
	ID        *int
	Name      string
	Comments  string
	Rules     []types.PolicyRule
	Version   string
	AutoGrant *bool
	Status    *typex.SimpleStatus
	AppID     int
	RoleIDs   []int
}

// Mutate applies the CreateAppPolicyInput on the AppPolicyMutation builder.
func (i *CreateAppPolicyInput) Mutate(m *AppPolicyMutation) {
	m.SetName(i.Name)
	m.SetComments(i.Comments)
	if v := i.Rules; v != nil {
		m.SetRules(v)
	}
	m.SetVersion(i.Version)
	if v := i.AutoGrant; v != nil {
		m.SetAutoGrant(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	m.SetAppID(i.AppID)
	if v := i.RoleIDs; len(v) > 0 {
		m.AddRoleIDs(v...)
	}
}

// SetInput applies the change-set in the CreateAppPolicyInput on the AppPolicyCreate builder.
func (c *AppPolicyCreate) SetInput(i CreateAppPolicyInput) *AppPolicyCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateAppPolicyInput represents a mutation input for updating apppolicies.
type UpdateAppPolicyInput struct {
	Name          *string
	Comments      *string
	Rules         []types.PolicyRule
	AppendRules   []types.PolicyRule
	Version       *string
	AutoGrant     *bool
	ClearStatus   bool
	Status        *typex.SimpleStatus
	ClearApp      bool
	AppID         *int
	ClearRoles    bool
	AddRoleIDs    []int
	RemoveRoleIDs []int
}

// Mutate applies the UpdateAppPolicyInput on the AppPolicyMutation builder.
func (i *UpdateAppPolicyInput) Mutate(m *AppPolicyMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.Rules; v != nil {
		m.SetRules(v)
	}
	if i.AppendRules != nil {
		m.AppendRules(i.Rules)
	}
	if v := i.Version; v != nil {
		m.SetVersion(*v)
	}
	if v := i.AutoGrant; v != nil {
		m.SetAutoGrant(*v)
	}
	if i.ClearStatus {
		m.ClearStatus()
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearApp {
		m.ClearApp()
	}
	if v := i.AppID; v != nil {
		m.SetAppID(*v)
	}
	if i.ClearRoles {
		m.ClearRoles()
	}
	if v := i.AddRoleIDs; len(v) > 0 {
		m.AddRoleIDs(v...)
	}
	if v := i.RemoveRoleIDs; len(v) > 0 {
		m.RemoveRoleIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateAppPolicyInput on the AppPolicyUpdate builder.
func (c *AppPolicyUpdate) SetInput(i UpdateAppPolicyInput) *AppPolicyUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateAppPolicyInput on the AppPolicyUpdateOne builder.
func (c *AppPolicyUpdateOne) SetInput(i UpdateAppPolicyInput) *AppPolicyUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateAppResInput represents a mutation input for creating appresslice.
type CreateAppResInput struct {
	ID         *int
	Name       string
	TypeName   string
	ArnPattern string
	AppID      int
}

// Mutate applies the CreateAppResInput on the AppResMutation builder.
func (i *CreateAppResInput) Mutate(m *AppResMutation) {
	m.SetName(i.Name)
	m.SetTypeName(i.TypeName)
	m.SetArnPattern(i.ArnPattern)
	m.SetAppID(i.AppID)
}

// SetInput applies the change-set in the CreateAppResInput on the AppResCreate builder.
func (c *AppResCreate) SetInput(i CreateAppResInput) *AppResCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateAppResInput represents a mutation input for updating appresslice.
type UpdateAppResInput struct {
	Name       *string
	TypeName   *string
	ArnPattern *string
	ClearApp   bool
	AppID      *int
}

// Mutate applies the UpdateAppResInput on the AppResMutation builder.
func (i *UpdateAppResInput) Mutate(m *AppResMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.TypeName; v != nil {
		m.SetTypeName(*v)
	}
	if v := i.ArnPattern; v != nil {
		m.SetArnPattern(*v)
	}
	if i.ClearApp {
		m.ClearApp()
	}
	if v := i.AppID; v != nil {
		m.SetAppID(*v)
	}
}

// SetInput applies the change-set in the UpdateAppResInput on the AppResUpdate builder.
func (c *AppResUpdate) SetInput(i UpdateAppResInput) *AppResUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateAppResInput on the AppResUpdateOne builder.
func (c *AppResUpdateOne) SetInput(i UpdateAppResInput) *AppResUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateAppRoleInput represents a mutation input for creating approles.
type CreateAppRoleInput struct {
	ID        *int
	Name      string
	Comments  *string
	AutoGrant *bool
	Editable  *bool
	AppID     int
	PolicyIDs []int
}

// Mutate applies the CreateAppRoleInput on the AppRoleMutation builder.
func (i *CreateAppRoleInput) Mutate(m *AppRoleMutation) {
	m.SetName(i.Name)
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.AutoGrant; v != nil {
		m.SetAutoGrant(*v)
	}
	if v := i.Editable; v != nil {
		m.SetEditable(*v)
	}
	m.SetAppID(i.AppID)
	if v := i.PolicyIDs; len(v) > 0 {
		m.AddPolicyIDs(v...)
	}
}

// SetInput applies the change-set in the CreateAppRoleInput on the AppRoleCreate builder.
func (c *AppRoleCreate) SetInput(i CreateAppRoleInput) *AppRoleCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateAppRoleInput represents a mutation input for updating approles.
type UpdateAppRoleInput struct {
	Name            *string
	ClearComments   bool
	Comments        *string
	AutoGrant       *bool
	Editable        *bool
	ClearApp        bool
	AppID           *int
	ClearPolicies   bool
	AddPolicyIDs    []int
	RemovePolicyIDs []int
}

// Mutate applies the UpdateAppRoleInput on the AppRoleMutation builder.
func (i *UpdateAppRoleInput) Mutate(m *AppRoleMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.AutoGrant; v != nil {
		m.SetAutoGrant(*v)
	}
	if v := i.Editable; v != nil {
		m.SetEditable(*v)
	}
	if i.ClearApp {
		m.ClearApp()
	}
	if v := i.AppID; v != nil {
		m.SetAppID(*v)
	}
	if i.ClearPolicies {
		m.ClearPolicies()
	}
	if v := i.AddPolicyIDs; len(v) > 0 {
		m.AddPolicyIDs(v...)
	}
	if v := i.RemovePolicyIDs; len(v) > 0 {
		m.RemovePolicyIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateAppRoleInput on the AppRoleUpdate builder.
func (c *AppRoleUpdate) SetInput(i UpdateAppRoleInput) *AppRoleUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateAppRoleInput on the AppRoleUpdateOne builder.
func (c *AppRoleUpdateOne) SetInput(i UpdateAppRoleInput) *AppRoleUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateOrgInput represents a mutation input for creating orgs.
type CreateOrgInput struct {
	ID               *int
	Domain           *string
	Name             string
	Profile          *string
	Status           *typex.SimpleStatus
	CountryCode      *string
	Timezone         *string
	ParentID         int
	ChildIDs         []int
	OwnerID          *int
	UserIDs          []int
	RolesAndGroupIDs []int
	PermissionIDs    []int
	PolicyIDs        []int
	AppIDs           []int
}

// Mutate applies the CreateOrgInput on the OrgMutation builder.
func (i *CreateOrgInput) Mutate(m *OrgMutation) {
	if v := i.Domain; v != nil {
		m.SetDomain(*v)
	}
	m.SetName(i.Name)
	if v := i.Profile; v != nil {
		m.SetProfile(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.CountryCode; v != nil {
		m.SetCountryCode(*v)
	}
	if v := i.Timezone; v != nil {
		m.SetTimezone(*v)
	}
	m.SetParentID(i.ParentID)
	if v := i.ChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
	if v := i.UserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
	if v := i.RolesAndGroupIDs; len(v) > 0 {
		m.AddRolesAndGroupIDs(v...)
	}
	if v := i.PermissionIDs; len(v) > 0 {
		m.AddPermissionIDs(v...)
	}
	if v := i.PolicyIDs; len(v) > 0 {
		m.AddPolicyIDs(v...)
	}
	if v := i.AppIDs; len(v) > 0 {
		m.AddAppIDs(v...)
	}
}

// SetInput applies the change-set in the CreateOrgInput on the OrgCreate builder.
func (c *OrgCreate) SetInput(i CreateOrgInput) *OrgCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateOrgInput represents a mutation input for updating orgs.
type UpdateOrgInput struct {
	ClearDomain            bool
	Domain                 *string
	Name                   *string
	ClearProfile           bool
	Profile                *string
	ClearStatus            bool
	Status                 *typex.SimpleStatus
	ClearCountryCode       bool
	CountryCode            *string
	ClearTimezone          bool
	Timezone               *string
	ClearParent            bool
	ParentID               *int
	ClearChildren          bool
	AddChildIDs            []int
	RemoveChildIDs         []int
	ClearOwner             bool
	OwnerID                *int
	ClearUsers             bool
	AddUserIDs             []int
	RemoveUserIDs          []int
	ClearRolesAndGroups    bool
	AddRolesAndGroupIDs    []int
	RemoveRolesAndGroupIDs []int
	ClearPermissions       bool
	AddPermissionIDs       []int
	RemovePermissionIDs    []int
	ClearPolicies          bool
	AddPolicyIDs           []int
	RemovePolicyIDs        []int
	ClearApps              bool
	AddAppIDs              []int
	RemoveAppIDs           []int
}

// Mutate applies the UpdateOrgInput on the OrgMutation builder.
func (i *UpdateOrgInput) Mutate(m *OrgMutation) {
	if i.ClearDomain {
		m.ClearDomain()
	}
	if v := i.Domain; v != nil {
		m.SetDomain(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearProfile {
		m.ClearProfile()
	}
	if v := i.Profile; v != nil {
		m.SetProfile(*v)
	}
	if i.ClearStatus {
		m.ClearStatus()
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearCountryCode {
		m.ClearCountryCode()
	}
	if v := i.CountryCode; v != nil {
		m.SetCountryCode(*v)
	}
	if i.ClearTimezone {
		m.ClearTimezone()
	}
	if v := i.Timezone; v != nil {
		m.SetTimezone(*v)
	}
	if i.ClearParent {
		m.ClearParent()
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if i.ClearChildren {
		m.ClearChildren()
	}
	if v := i.AddChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.RemoveChildIDs; len(v) > 0 {
		m.RemoveChildIDs(v...)
	}
	if i.ClearOwner {
		m.ClearOwner()
	}
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
	if i.ClearUsers {
		m.ClearUsers()
	}
	if v := i.AddUserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
	if v := i.RemoveUserIDs; len(v) > 0 {
		m.RemoveUserIDs(v...)
	}
	if i.ClearRolesAndGroups {
		m.ClearRolesAndGroups()
	}
	if v := i.AddRolesAndGroupIDs; len(v) > 0 {
		m.AddRolesAndGroupIDs(v...)
	}
	if v := i.RemoveRolesAndGroupIDs; len(v) > 0 {
		m.RemoveRolesAndGroupIDs(v...)
	}
	if i.ClearPermissions {
		m.ClearPermissions()
	}
	if v := i.AddPermissionIDs; len(v) > 0 {
		m.AddPermissionIDs(v...)
	}
	if v := i.RemovePermissionIDs; len(v) > 0 {
		m.RemovePermissionIDs(v...)
	}
	if i.ClearPolicies {
		m.ClearPolicies()
	}
	if v := i.AddPolicyIDs; len(v) > 0 {
		m.AddPolicyIDs(v...)
	}
	if v := i.RemovePolicyIDs; len(v) > 0 {
		m.RemovePolicyIDs(v...)
	}
	if i.ClearApps {
		m.ClearApps()
	}
	if v := i.AddAppIDs; len(v) > 0 {
		m.AddAppIDs(v...)
	}
	if v := i.RemoveAppIDs; len(v) > 0 {
		m.RemoveAppIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateOrgInput on the OrgUpdate builder.
func (c *OrgUpdate) SetInput(i UpdateOrgInput) *OrgUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateOrgInput on the OrgUpdateOne builder.
func (c *OrgUpdateOne) SetInput(i UpdateOrgInput) *OrgUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateOrgPolicyInput represents a mutation input for creating orgpolicies.
type CreateOrgPolicyInput struct {
	ID          *int
	AppPolicyID *int
	Name        string
	Comments    string
	Rules       []types.PolicyRule
	OrgID       int
}

// Mutate applies the CreateOrgPolicyInput on the OrgPolicyMutation builder.
func (i *CreateOrgPolicyInput) Mutate(m *OrgPolicyMutation) {
	if v := i.AppPolicyID; v != nil {
		m.SetAppPolicyID(*v)
	}
	m.SetName(i.Name)
	m.SetComments(i.Comments)
	if v := i.Rules; v != nil {
		m.SetRules(v)
	}
	m.SetOrgID(i.OrgID)
}

// SetInput applies the change-set in the CreateOrgPolicyInput on the OrgPolicyCreate builder.
func (c *OrgPolicyCreate) SetInput(i CreateOrgPolicyInput) *OrgPolicyCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateOrgPolicyInput represents a mutation input for updating orgpolicies.
type UpdateOrgPolicyInput struct {
	ClearAppPolicyID bool
	AppPolicyID      *int
	Name             *string
	Comments         *string
	Rules            []types.PolicyRule
	AppendRules      []types.PolicyRule
	ClearOrg         bool
	OrgID            *int
}

// Mutate applies the UpdateOrgPolicyInput on the OrgPolicyMutation builder.
func (i *UpdateOrgPolicyInput) Mutate(m *OrgPolicyMutation) {
	if i.ClearAppPolicyID {
		m.ClearAppPolicyID()
	}
	if v := i.AppPolicyID; v != nil {
		m.SetAppPolicyID(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.Rules; v != nil {
		m.SetRules(v)
	}
	if i.AppendRules != nil {
		m.AppendRules(i.Rules)
	}
	if i.ClearOrg {
		m.ClearOrg()
	}
	if v := i.OrgID; v != nil {
		m.SetOrgID(*v)
	}
}

// SetInput applies the change-set in the UpdateOrgPolicyInput on the OrgPolicyUpdate builder.
func (c *OrgPolicyUpdate) SetInput(i UpdateOrgPolicyInput) *OrgPolicyUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateOrgPolicyInput on the OrgPolicyUpdateOne builder.
func (c *OrgPolicyUpdateOne) SetInput(i UpdateOrgPolicyInput) *OrgPolicyUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreatePermissionInput represents a mutation input for creating permissions.
type CreatePermissionInput struct {
	ID            *int
	PrincipalKind permission.PrincipalKind
	RoleID        *int
	OrgPolicyID   int
	StartAt       *time.Time
	EndAt         *time.Time
	OrgID         int
	UserID        *int
}

// Mutate applies the CreatePermissionInput on the PermissionMutation builder.
func (i *CreatePermissionInput) Mutate(m *PermissionMutation) {
	m.SetPrincipalKind(i.PrincipalKind)
	if v := i.RoleID; v != nil {
		m.SetRoleID(*v)
	}
	m.SetOrgPolicyID(i.OrgPolicyID)
	if v := i.StartAt; v != nil {
		m.SetStartAt(*v)
	}
	if v := i.EndAt; v != nil {
		m.SetEndAt(*v)
	}
	m.SetOrgID(i.OrgID)
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreatePermissionInput on the PermissionCreate builder.
func (c *PermissionCreate) SetInput(i CreatePermissionInput) *PermissionCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePermissionInput represents a mutation input for updating permissions.
type UpdatePermissionInput struct {
	PrincipalKind *permission.PrincipalKind
	ClearRoleID   bool
	RoleID        *int
	OrgPolicyID   *int
	ClearStartAt  bool
	StartAt       *time.Time
	ClearEndAt    bool
	EndAt         *time.Time
	ClearStatus   bool
	Status        *typex.SimpleStatus
	ClearOrg      bool
	OrgID         *int
	ClearUser     bool
	UserID        *int
}

// Mutate applies the UpdatePermissionInput on the PermissionMutation builder.
func (i *UpdatePermissionInput) Mutate(m *PermissionMutation) {
	if v := i.PrincipalKind; v != nil {
		m.SetPrincipalKind(*v)
	}
	if i.ClearRoleID {
		m.ClearRoleID()
	}
	if v := i.RoleID; v != nil {
		m.SetRoleID(*v)
	}
	if v := i.OrgPolicyID; v != nil {
		m.SetOrgPolicyID(*v)
	}
	if i.ClearStartAt {
		m.ClearStartAt()
	}
	if v := i.StartAt; v != nil {
		m.SetStartAt(*v)
	}
	if i.ClearEndAt {
		m.ClearEndAt()
	}
	if v := i.EndAt; v != nil {
		m.SetEndAt(*v)
	}
	if i.ClearStatus {
		m.ClearStatus()
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearOrg {
		m.ClearOrg()
	}
	if v := i.OrgID; v != nil {
		m.SetOrgID(*v)
	}
	if i.ClearUser {
		m.ClearUser()
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the UpdatePermissionInput on the PermissionUpdate builder.
func (c *PermissionUpdate) SetInput(i UpdatePermissionInput) *PermissionUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePermissionInput on the PermissionUpdateOne builder.
func (c *PermissionUpdateOne) SetInput(i UpdatePermissionInput) *PermissionUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	ID             *int
	PrincipalName  string
	DisplayName    string
	Email          *string
	Mobile         *string
	Status         *typex.SimpleStatus
	Comments       *string
	IdentityIDs    []int
	LoginProfileID *int
	PasswordIDs    []int
	DeviceIDs      []int
	PermissionIDs  []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetPrincipalName(i.PrincipalName)
	m.SetDisplayName(i.DisplayName)
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Mobile; v != nil {
		m.SetMobile(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.IdentityIDs; len(v) > 0 {
		m.AddIdentityIDs(v...)
	}
	if v := i.LoginProfileID; v != nil {
		m.SetLoginProfileID(*v)
	}
	if v := i.PasswordIDs; len(v) > 0 {
		m.AddPasswordIDs(v...)
	}
	if v := i.DeviceIDs; len(v) > 0 {
		m.AddDeviceIDs(v...)
	}
	if v := i.PermissionIDs; len(v) > 0 {
		m.AddPermissionIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	PrincipalName       *string
	DisplayName         *string
	ClearEmail          bool
	Email               *string
	ClearMobile         bool
	Mobile              *string
	ClearComments       bool
	Comments            *string
	ClearIdentities     bool
	AddIdentityIDs      []int
	RemoveIdentityIDs   []int
	ClearLoginProfile   bool
	LoginProfileID      *int
	ClearPasswords      bool
	AddPasswordIDs      []int
	RemovePasswordIDs   []int
	ClearDevices        bool
	AddDeviceIDs        []int
	RemoveDeviceIDs     []int
	ClearPermissions    bool
	AddPermissionIDs    []int
	RemovePermissionIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.PrincipalName; v != nil {
		m.SetPrincipalName(*v)
	}
	if v := i.DisplayName; v != nil {
		m.SetDisplayName(*v)
	}
	if i.ClearEmail {
		m.ClearEmail()
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if i.ClearMobile {
		m.ClearMobile()
	}
	if v := i.Mobile; v != nil {
		m.SetMobile(*v)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if i.ClearIdentities {
		m.ClearIdentities()
	}
	if v := i.AddIdentityIDs; len(v) > 0 {
		m.AddIdentityIDs(v...)
	}
	if v := i.RemoveIdentityIDs; len(v) > 0 {
		m.RemoveIdentityIDs(v...)
	}
	if i.ClearLoginProfile {
		m.ClearLoginProfile()
	}
	if v := i.LoginProfileID; v != nil {
		m.SetLoginProfileID(*v)
	}
	if i.ClearPasswords {
		m.ClearPasswords()
	}
	if v := i.AddPasswordIDs; len(v) > 0 {
		m.AddPasswordIDs(v...)
	}
	if v := i.RemovePasswordIDs; len(v) > 0 {
		m.RemovePasswordIDs(v...)
	}
	if i.ClearDevices {
		m.ClearDevices()
	}
	if v := i.AddDeviceIDs; len(v) > 0 {
		m.AddDeviceIDs(v...)
	}
	if v := i.RemoveDeviceIDs; len(v) > 0 {
		m.RemoveDeviceIDs(v...)
	}
	if i.ClearPermissions {
		m.ClearPermissions()
	}
	if v := i.AddPermissionIDs; len(v) > 0 {
		m.AddPermissionIDs(v...)
	}
	if v := i.RemovePermissionIDs; len(v) > 0 {
		m.RemovePermissionIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserIdentityInput represents a mutation input for creating useridentities.
type CreateUserIdentityInput struct {
	ID         *int
	Kind       useridentity.Kind
	Code       *string
	CodeExtend *string
	Status     *typex.SimpleStatus
	UserID     *int
}

// Mutate applies the CreateUserIdentityInput on the UserIdentityMutation builder.
func (i *CreateUserIdentityInput) Mutate(m *UserIdentityMutation) {
	m.SetKind(i.Kind)
	if v := i.Code; v != nil {
		m.SetCode(*v)
	}
	if v := i.CodeExtend; v != nil {
		m.SetCodeExtend(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreateUserIdentityInput on the UserIdentityCreate builder.
func (c *UserIdentityCreate) SetInput(i CreateUserIdentityInput) *UserIdentityCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserIdentityInput represents a mutation input for updating useridentities.
type UpdateUserIdentityInput struct {
	Kind            *useridentity.Kind
	ClearCode       bool
	Code            *string
	ClearCodeExtend bool
	CodeExtend      *string
	ClearStatus     bool
	Status          *typex.SimpleStatus
}

// Mutate applies the UpdateUserIdentityInput on the UserIdentityMutation builder.
func (i *UpdateUserIdentityInput) Mutate(m *UserIdentityMutation) {
	if v := i.Kind; v != nil {
		m.SetKind(*v)
	}
	if i.ClearCode {
		m.ClearCode()
	}
	if v := i.Code; v != nil {
		m.SetCode(*v)
	}
	if i.ClearCodeExtend {
		m.ClearCodeExtend()
	}
	if v := i.CodeExtend; v != nil {
		m.SetCodeExtend(*v)
	}
	if i.ClearStatus {
		m.ClearStatus()
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
}

// SetInput applies the change-set in the UpdateUserIdentityInput on the UserIdentityUpdate builder.
func (c *UserIdentityUpdate) SetInput(i UpdateUserIdentityInput) *UserIdentityUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserIdentityInput on the UserIdentityUpdateOne builder.
func (c *UserIdentityUpdateOne) SetInput(i UpdateUserIdentityInput) *UserIdentityUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserLoginProfileInput represents a mutation input for creating userloginprofiles.
type CreateUserLoginProfileInput struct {
	ID            *int
	CanLogin      *bool
	SetKind       userloginprofile.SetKind
	PasswordReset *bool
	VerifyDevice  bool
	MfaEnabled    *bool
	MfaSecret     *string
	MfaStatus     *typex.SimpleStatus
	UserID        *int
}

// Mutate applies the CreateUserLoginProfileInput on the UserLoginProfileMutation builder.
func (i *CreateUserLoginProfileInput) Mutate(m *UserLoginProfileMutation) {
	if v := i.CanLogin; v != nil {
		m.SetCanLogin(*v)
	}
	m.SetSetKind(i.SetKind)
	if v := i.PasswordReset; v != nil {
		m.SetPasswordReset(*v)
	}
	m.SetVerifyDevice(i.VerifyDevice)
	if v := i.MfaEnabled; v != nil {
		m.SetMfaEnabled(*v)
	}
	if v := i.MfaSecret; v != nil {
		m.SetMfaSecret(*v)
	}
	if v := i.MfaStatus; v != nil {
		m.SetMfaStatus(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreateUserLoginProfileInput on the UserLoginProfileCreate builder.
func (c *UserLoginProfileCreate) SetInput(i CreateUserLoginProfileInput) *UserLoginProfileCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserLoginProfileInput represents a mutation input for updating userloginprofiles.
type UpdateUserLoginProfileInput struct {
	ClearCanLogin      bool
	CanLogin           *bool
	SetKind            *userloginprofile.SetKind
	ClearPasswordReset bool
	PasswordReset      *bool
	VerifyDevice       *bool
	ClearMfaEnabled    bool
	MfaEnabled         *bool
	ClearMfaSecret     bool
	MfaSecret          *string
	ClearMfaStatus     bool
	MfaStatus          *typex.SimpleStatus
}

// Mutate applies the UpdateUserLoginProfileInput on the UserLoginProfileMutation builder.
func (i *UpdateUserLoginProfileInput) Mutate(m *UserLoginProfileMutation) {
	if i.ClearCanLogin {
		m.ClearCanLogin()
	}
	if v := i.CanLogin; v != nil {
		m.SetCanLogin(*v)
	}
	if v := i.SetKind; v != nil {
		m.SetSetKind(*v)
	}
	if i.ClearPasswordReset {
		m.ClearPasswordReset()
	}
	if v := i.PasswordReset; v != nil {
		m.SetPasswordReset(*v)
	}
	if v := i.VerifyDevice; v != nil {
		m.SetVerifyDevice(*v)
	}
	if i.ClearMfaEnabled {
		m.ClearMfaEnabled()
	}
	if v := i.MfaEnabled; v != nil {
		m.SetMfaEnabled(*v)
	}
	if i.ClearMfaSecret {
		m.ClearMfaSecret()
	}
	if v := i.MfaSecret; v != nil {
		m.SetMfaSecret(*v)
	}
	if i.ClearMfaStatus {
		m.ClearMfaStatus()
	}
	if v := i.MfaStatus; v != nil {
		m.SetMfaStatus(*v)
	}
}

// SetInput applies the change-set in the UpdateUserLoginProfileInput on the UserLoginProfileUpdate builder.
func (c *UserLoginProfileUpdate) SetInput(i UpdateUserLoginProfileInput) *UserLoginProfileUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserLoginProfileInput on the UserLoginProfileUpdateOne builder.
func (c *UserLoginProfileUpdateOne) SetInput(i UpdateUserLoginProfileInput) *UserLoginProfileUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserPasswordInput represents a mutation input for creating userpasswords.
type CreateUserPasswordInput struct {
	ID       *int
	Scene    *userpassword.Scene
	Password *string
	Status   *typex.SimpleStatus
	Memo     *string
	UserID   *int
}

// Mutate applies the CreateUserPasswordInput on the UserPasswordMutation builder.
func (i *CreateUserPasswordInput) Mutate(m *UserPasswordMutation) {
	if v := i.Scene; v != nil {
		m.SetScene(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreateUserPasswordInput on the UserPasswordCreate builder.
func (c *UserPasswordCreate) SetInput(i CreateUserPasswordInput) *UserPasswordCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserPasswordInput represents a mutation input for updating userpasswords.
type UpdateUserPasswordInput struct {
	ClearScene    bool
	Scene         *userpassword.Scene
	ClearPassword bool
	Password      *string
	ClearStatus   bool
	Status        *typex.SimpleStatus
	ClearMemo     bool
	Memo          *string
}

// Mutate applies the UpdateUserPasswordInput on the UserPasswordMutation builder.
func (i *UpdateUserPasswordInput) Mutate(m *UserPasswordMutation) {
	if i.ClearScene {
		m.ClearScene()
	}
	if v := i.Scene; v != nil {
		m.SetScene(*v)
	}
	if i.ClearPassword {
		m.ClearPassword()
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if i.ClearStatus {
		m.ClearStatus()
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearMemo {
		m.ClearMemo()
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
}

// SetInput applies the change-set in the UpdateUserPasswordInput on the UserPasswordUpdate builder.
func (c *UserPasswordUpdate) SetInput(i UpdateUserPasswordInput) *UserPasswordUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserPasswordInput on the UserPasswordUpdateOne builder.
func (c *UserPasswordUpdateOne) SetInput(i UpdateUserPasswordInput) *UserPasswordUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
