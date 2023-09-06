// Code generated by ent, DO NOT EDIT.

package app

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// FieldLogoFileID holds the string denoting the logo_file_id field in the database.
	FieldLogoFileID = "logo_file_id"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPrivate holds the string denoting the private field in the database.
	FieldPrivate = "private"
	// FieldOwnerOrgID holds the string denoting the owner_org_id field in the database.
	FieldOwnerOrgID = "owner_org_id"
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
	FieldLogoFileID,
	FieldComments,
	FieldStatus,
	FieldPrivate,
	FieldOwnerOrgID,
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
	// DefaultPrivate holds the default value on creation for the "private" field.
	DefaultPrivate bool
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

// OrderOption defines the ordering options for the App queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByKind orders the results by the kind field.
func ByKind(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKind, opts...).ToFunc()
}

// ByRedirectURI orders the results by the redirect_uri field.
func ByRedirectURI(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRedirectURI, opts...).ToFunc()
}

// ByAppKey orders the results by the app_key field.
func ByAppKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppKey, opts...).ToFunc()
}

// ByAppSecret orders the results by the app_secret field.
func ByAppSecret(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppSecret, opts...).ToFunc()
}

// ByScopes orders the results by the scopes field.
func ByScopes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScopes, opts...).ToFunc()
}

// ByTokenValidity orders the results by the token_validity field.
func ByTokenValidity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTokenValidity, opts...).ToFunc()
}

// ByRefreshTokenValidity orders the results by the refresh_token_validity field.
func ByRefreshTokenValidity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefreshTokenValidity, opts...).ToFunc()
}

// ByLogoFileID orders the results by the logo_file_id field.
func ByLogoFileID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogoFileID, opts...).ToFunc()
}

// ByComments orders the results by the comments field.
func ByComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComments, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPrivate orders the results by the private field.
func ByPrivate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivate, opts...).ToFunc()
}

// ByOwnerOrgID orders the results by the owner_org_id field.
func ByOwnerOrgID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerOrgID, opts...).ToFunc()
}

// ByMenusCount orders the results by menus count.
func ByMenusCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMenusStep(), opts...)
	}
}

// ByMenus orders the results by menus terms.
func ByMenus(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMenusStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByActionsCount orders the results by actions count.
func ByActionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newActionsStep(), opts...)
	}
}

// ByActions orders the results by actions terms.
func ByActions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByResourcesCount orders the results by resources count.
func ByResourcesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newResourcesStep(), opts...)
	}
}

// ByResources orders the results by resources terms.
func ByResources(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newResourcesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRolesCount orders the results by roles count.
func ByRolesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRolesStep(), opts...)
	}
}

// ByRoles orders the results by roles terms.
func ByRoles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRolesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPoliciesCount orders the results by policies count.
func ByPoliciesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPoliciesStep(), opts...)
	}
}

// ByPolicies orders the results by policies terms.
func ByPolicies(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPoliciesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrgsCount orders the results by orgs count.
func ByOrgsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrgsStep(), opts...)
	}
}

// ByOrgs orders the results by orgs terms.
func ByOrgs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrgsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrgAppCount orders the results by org_app count.
func ByOrgAppCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrgAppStep(), opts...)
	}
}

// ByOrgApp orders the results by org_app terms.
func ByOrgApp(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrgAppStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMenusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MenusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MenusTable, MenusColumn),
	)
}
func newActionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ActionsTable, ActionsColumn),
	)
}
func newResourcesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ResourcesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ResourcesTable, ResourcesColumn),
	)
}
func newRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RolesTable, RolesColumn),
	)
}
func newPoliciesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PoliciesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PoliciesTable, PoliciesColumn),
	)
}
func newOrgsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrgsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, OrgsTable, OrgsPrimaryKey...),
	)
}
func newOrgAppStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrgAppInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, OrgAppTable, OrgAppColumn),
	)
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
