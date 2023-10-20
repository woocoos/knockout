// Code generated by ent, DO NOT EDIT.

package userdevice

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
	// Label holds the string label denoting the userdevice type in the database.
	Label = "user_device"
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
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldDeviceUID holds the string denoting the device_uid field in the database.
	FieldDeviceUID = "device_uid"
	// FieldDeviceName holds the string denoting the device_name field in the database.
	FieldDeviceName = "device_name"
	// FieldSystemName holds the string denoting the system_name field in the database.
	FieldSystemName = "system_name"
	// FieldSystemVersion holds the string denoting the system_version field in the database.
	FieldSystemVersion = "system_version"
	// FieldAppVersion holds the string denoting the app_version field in the database.
	FieldAppVersion = "app_version"
	// FieldDeviceModel holds the string denoting the device_model field in the database.
	FieldDeviceModel = "device_model"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the userdevice in the database.
	Table = "user_device"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_device"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "user"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for userdevice fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldUserID,
	FieldDeviceUID,
	FieldDeviceName,
	FieldSystemName,
	FieldSystemVersion,
	FieldAppVersion,
	FieldDeviceModel,
	FieldStatus,
	FieldComments,
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
	Hooks [2]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DeviceUIDValidator is a validator for the "device_uid" field. It is called by the builders before save.
	DeviceUIDValidator func(string) error
	// DeviceNameValidator is a validator for the "device_name" field. It is called by the builders before save.
	DeviceNameValidator func(string) error
	// SystemNameValidator is a validator for the "system_name" field. It is called by the builders before save.
	SystemNameValidator func(string) error
	// SystemVersionValidator is a validator for the "system_version" field. It is called by the builders before save.
	SystemVersionValidator func(string) error
	// AppVersionValidator is a validator for the "app_version" field. It is called by the builders before save.
	AppVersionValidator func(string) error
	// DeviceModelValidator is a validator for the "device_model" field. It is called by the builders before save.
	DeviceModelValidator func(string) error
)

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s typex.SimpleStatus) error {
	switch s.String() {
	case "active", "inactive", "processing", "disabled":
		return nil
	default:
		return fmt.Errorf("userdevice: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the UserDevice queries.
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

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByDeviceUID orders the results by the device_uid field.
func ByDeviceUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceUID, opts...).ToFunc()
}

// ByDeviceName orders the results by the device_name field.
func ByDeviceName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceName, opts...).ToFunc()
}

// BySystemName orders the results by the system_name field.
func BySystemName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSystemName, opts...).ToFunc()
}

// BySystemVersion orders the results by the system_version field.
func BySystemVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSystemVersion, opts...).ToFunc()
}

// ByAppVersion orders the results by the app_version field.
func ByAppVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppVersion, opts...).ToFunc()
}

// ByDeviceModel orders the results by the device_model field.
func ByDeviceModel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceModel, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByComments orders the results by the comments field.
func ByComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComments, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}

var (
	// typex.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*typex.SimpleStatus)(nil)
	// typex.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*typex.SimpleStatus)(nil)
)
