// Code generated by ent, DO NOT EDIT.

package appdictitem

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/entco/schemax/typex"
)

const (
	// Label holds the string label denoting the appdictitem type in the database.
	Label = "app_dict_item"
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
	// FieldOrgID holds the string denoting the org_id field in the database.
	FieldOrgID = "org_id"
	// FieldDictID holds the string denoting the dict_id field in the database.
	FieldDictID = "dict_id"
	// FieldRefCode holds the string denoting the ref_code field in the database.
	FieldRefCode = "ref_code"
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldDisplaySort holds the string denoting the display_sort field in the database.
	FieldDisplaySort = "display_sort"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeDict holds the string denoting the dict edge name in mutations.
	EdgeDict = "dict"
	// Table holds the table name of the appdictitem in the database.
	Table = "app_dict_item"
	// DictTable is the table that holds the dict relation/edge.
	DictTable = "app_dict_item"
	// DictInverseTable is the table name for the AppDict entity.
	// It exists in this package in order to avoid circular dependency with the "appdict" package.
	DictInverseTable = "app_dict"
	// DictColumn is the table column denoting the dict relation/edge.
	DictColumn = "dict_id"
)

// Columns holds all SQL columns for appdictitem fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldOrgID,
	FieldDictID,
	FieldRefCode,
	FieldCode,
	FieldName,
	FieldComments,
	FieldDisplaySort,
	FieldStatus,
}

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
	Hooks [6]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

const DefaultStatus typex.SimpleStatus = "inactive"

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s typex.SimpleStatus) error {
	switch s.String() {
	case "active", "inactive", "processing", "disabled":
		return nil
	default:
		return fmt.Errorf("appdictitem: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the AppDictItem queries.
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

// ByOrgID orders the results by the org_id field.
func ByOrgID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrgID, opts...).ToFunc()
}

// ByDictID orders the results by the dict_id field.
func ByDictID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDictID, opts...).ToFunc()
}

// ByRefCode orders the results by the ref_code field.
func ByRefCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefCode, opts...).ToFunc()
}

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByComments orders the results by the comments field.
func ByComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComments, opts...).ToFunc()
}

// ByDisplaySort orders the results by the display_sort field.
func ByDisplaySort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplaySort, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByDictField orders the results by dict field.
func ByDictField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDictStep(), sql.OrderByField(field, opts...))
	}
}
func newDictStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DictInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DictTable, DictColumn),
	)
}

var (
	// typex.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*typex.SimpleStatus)(nil)
	// typex.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*typex.SimpleStatus)(nil)
)
