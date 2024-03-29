// Code generated by ent, DO NOT EDIT.

package apppolicy

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
)

const (
	// Label holds the string label denoting the apppolicy type in the database.
	Label = "app_policy"
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
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldRules holds the string denoting the rules field in the database.
	FieldRules = "rules"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldAutoGrant holds the string denoting the auto_grant field in the database.
	FieldAutoGrant = "auto_grant"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "app"
	// EdgeRoles holds the string denoting the roles edge name in mutations.
	EdgeRoles = "roles"
	// EdgeAppRolePolicy holds the string denoting the app_role_policy edge name in mutations.
	EdgeAppRolePolicy = "app_role_policy"
	// Table holds the table name of the apppolicy in the database.
	Table = "app_policy"
	// AppTable is the table that holds the app relation/edge.
	AppTable = "app_policy"
	// AppInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppInverseTable = "app"
	// AppColumn is the table column denoting the app relation/edge.
	AppColumn = "app_id"
	// RolesTable is the table that holds the roles relation/edge. The primary key declared below.
	RolesTable = "app_role_policy"
	// RolesInverseTable is the table name for the AppRole entity.
	// It exists in this package in order to avoid circular dependency with the "approle" package.
	RolesInverseTable = "app_role"
	// AppRolePolicyTable is the table that holds the app_role_policy relation/edge.
	AppRolePolicyTable = "app_role_policy"
	// AppRolePolicyInverseTable is the table name for the AppRolePolicy entity.
	// It exists in this package in order to avoid circular dependency with the "approlepolicy" package.
	AppRolePolicyInverseTable = "app_role_policy"
	// AppRolePolicyColumn is the table column denoting the app_role_policy relation/edge.
	AppRolePolicyColumn = "app_policy_id"
)

// Columns holds all SQL columns for apppolicy fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldAppID,
	FieldName,
	FieldComments,
	FieldRules,
	FieldVersion,
	FieldAutoGrant,
	FieldStatus,
}

var (
	// RolesPrimaryKey and RolesColumn2 are the table columns denoting the
	// primary key for the roles relation (M2M).
	RolesPrimaryKey = []string{"app_role_id", "app_policy_id"}
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
	Hooks [3]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion string
	// DefaultAutoGrant holds the default value on creation for the "auto_grant" field.
	DefaultAutoGrant bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int
)

const DefaultStatus typex.SimpleStatus = "active"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s typex.SimpleStatus) error {
	switch s.String() {
	case "active", "inactive", "processing", "disabled":
		return nil
	default:
		return fmt.Errorf("apppolicy: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the AppPolicy queries.
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

// ByAppID orders the results by the app_id field.
func ByAppID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByComments orders the results by the comments field.
func ByComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComments, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByAutoGrant orders the results by the auto_grant field.
func ByAutoGrant(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAutoGrant, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByAppField orders the results by app field.
func ByAppField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppStep(), sql.OrderByField(field, opts...))
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

// ByAppRolePolicyCount orders the results by app_role_policy count.
func ByAppRolePolicyCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAppRolePolicyStep(), opts...)
	}
}

// ByAppRolePolicy orders the results by app_role_policy terms.
func ByAppRolePolicy(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAppRolePolicyStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAppStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AppTable, AppColumn),
	)
}
func newRolesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RolesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, RolesTable, RolesPrimaryKey...),
	)
}
func newAppRolePolicyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AppRolePolicyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, AppRolePolicyTable, AppRolePolicyColumn),
	)
}

var (
	// typex.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*typex.SimpleStatus)(nil)
	// typex.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*typex.SimpleStatus)(nil)
)
