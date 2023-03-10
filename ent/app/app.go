// Code generated by ent, DO NOT EDIT.

package app

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/entco/schemax/typex"
)

const (
	// Label holds the string label denoting the app type in the database.
	Label = "app"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldRedirectURI holds the string denoting the redirect_uri field in the database.
	FieldRedirectURI = "redirect_uri"
	// FieldAppKey holds the string denoting the app_key field in the database.
	FieldAppKey = "app_key"
	// FieldAppSecret holds the string denoting the app_secret field in the database.
	FieldAppSecret = "app_secret"
	// FieldScopes holds the string denoting the scopes field in the database.
	FieldScopes = "scopes"
	// FieldTokenValidity holds the string denoting the token_validity field in the database.
	FieldTokenValidity = "token_validity"
	// FieldRefreshTokenValidity holds the string denoting the refresh_token_validity field in the database.
	FieldRefreshTokenValidity = "refresh_token_validity"
	// FieldLogo holds the string denoting the logo field in the database.
	FieldLogo = "logo"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedOrgID holds the string denoting the created_org_id field in the database.
	FieldCreatedOrgID = "created_org_id"
	// EdgeMenus holds the string denoting the menus edge name in mutations.
	EdgeMenus = "menus"
	// EdgeActions holds the string denoting the actions edge name in mutations.
	EdgeActions = "actions"
	// EdgeResources holds the string denoting the resources edge name in mutations.
	EdgeResources = "resources"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// EdgePolicies holds the string denoting the policies edge name in mutations.
	EdgePolicies = "policies"
	// EdgeOrgs holds the string denoting the orgs edge name in mutations.
	EdgeOrgs = "orgs"
	// EdgeOrgApp holds the string denoting the org_app edge name in mutations.
	EdgeOrgApp = "org_app"
	// Table holds the table name of the app in the database.
	Table = "app"
	// MenusTable is the table that holds the menus relation/edge.
	MenusTable = "app_menu"
	// MenusInverseTable is the table name for the AppMenu entity.
	// It exists in this package in order to avoid circular dependency with the "appmenu" package.
	MenusInverseTable = "app_menu"
	// MenusColumn is the table column denoting the menus relation/edge.
	MenusColumn = "app_id"
	// ActionsTable is the table that holds the actions relation/edge.
	ActionsTable = "app_action"
	// ActionsInverseTable is the table name for the AppAction entity.
	// It exists in this package in order to avoid circular dependency with the "appaction" package.
	ActionsInverseTable = "app_action"
	// ActionsColumn is the table column denoting the actions relation/edge.
	ActionsColumn = "app_id"
	// ResourcesTable is the table that holds the resources relation/edge.
	ResourcesTable = "app_res"
	// ResourcesInverseTable is the table name for the AppRes entity.
	// It exists in this package in order to avoid circular dependency with the "appres" package.
	ResourcesInverseTable = "app_res"
	// ResourcesColumn is the table column denoting the resources relation/edge.
	ResourcesColumn = "app_id"
	// RolesTable is the table that holds the roles relation/edge.
	RolesTable = "app_role"
	// RolesInverseTable is the table name for the AppRole entity.
	// It exists in this package in order to avoid circular dependency with the "approle" package.
	RolesInverseTable = "app_role"
	// RolesColumn is the table column denoting the roles relation/edge.
	RolesColumn = "app_id"
	// PoliciesTable is the table that holds the policies relation/edge.
	PoliciesTable = "app_policy"
	// PoliciesInverseTable is the table name for the AppPolicy entity.
	// It exists in this package in order to avoid circular dependency with the "apppolicy" package.
	PoliciesInverseTable = "app_policy"
	// PoliciesColumn is the table column denoting the policies relation/edge.
	PoliciesColumn = "app_id"
	// OrgsTable is the table that holds the orgs relation/edge. The primary key declared below.
	OrgsTable = "org_app"
	// OrgsInverseTable is the table name for the Org entity.
	// It exists in this package in order to avoid circular dependency with the "org" package.
	OrgsInverseTable = "org"
	// OrgAppTable is the table that holds the org_app relation/edge.
	OrgAppTable = "org_app"
	// OrgAppInverseTable is the table name for the OrgApp entity.
	// It exists in this package in order to avoid circular dependency with the "orgapp" package.
	OrgAppInverseTable = "org_app"
	// OrgAppColumn is the table column denoting the org_app relation/edge.
	OrgAppColumn = "app_id"
)

// Columns holds all SQL columns for app fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldName,
	FieldCode,
	FieldKind,
	FieldRedirectURI,
	FieldAppKey,
	FieldAppSecret,
	FieldScopes,
	FieldTokenValidity,
	FieldRefreshTokenValidity,
	FieldLogo,
	FieldComments,
	FieldStatus,
	FieldCreatedOrgID,
}

var (
	// OrgsPrimaryKey and OrgsColumn2 are the table columns denoting the
	// primary key for the orgs relation (M2M).
	OrgsPrimaryKey = []string{"org_id", "app_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/woocoos/knockout/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// RedirectURIValidator is a validator for the "redirect_uri" field. It is called by the builders before save.
	RedirectURIValidator func(string) error
	// AppSecretValidator is a validator for the "app_secret" field. It is called by the builders before save.
	AppSecretValidator func(string) error
	// ScopesValidator is a validator for the "scopes" field. It is called by the builders before save.
	ScopesValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int
)

// Kind defines the type for the "kind" enum field.
type Kind string

// Kind values.
const (
	KindWeb    Kind = "web"
	KindNative Kind = "native"
	KindServer Kind = "server"
)

func (k Kind) String() string {
	return string(k)
}

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindWeb, KindNative, KindServer:
		return nil
	default:
		return fmt.Errorf("app: invalid enum value for kind field: %q", k)
	}
}

const DefaultStatus typex.SimpleStatus = "active"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s typex.SimpleStatus) error {
	switch s.String() {
	case "active", "inactive", "processing":
		return nil
	default:
		return fmt.Errorf("app: invalid enum value for status field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Kind) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Kind) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Kind(str)
	if err := KindValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Kind", str)
	}
	return nil
}

var (
	// typex.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*typex.SimpleStatus)(nil)
	// typex.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*typex.SimpleStatus)(nil)
)
