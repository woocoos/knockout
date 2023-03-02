// Code generated by ent, DO NOT EDIT.

package appmenu

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the appmenu type in the database.
	Label = "app_menu"
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
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldActionID holds the string denoting the action_id field in the database.
	FieldActionID = "action_id"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldDisplaySort holds the string denoting the display_sort field in the database.
	FieldDisplaySort = "display_sort"
	// EdgeApp holds the string denoting the app edge name in mutations.
	EdgeApp = "app"
	// EdgeAction holds the string denoting the action edge name in mutations.
	EdgeAction = "action"
	// Table holds the table name of the appmenu in the database.
	Table = "app_menu"
	// AppTable is the table that holds the app relation/edge.
	AppTable = "app_menu"
	// AppInverseTable is the table name for the App entity.
	// It exists in this package in order to avoid circular dependency with the "app" package.
	AppInverseTable = "app"
	// AppColumn is the table column denoting the app relation/edge.
	AppColumn = "app_id"
	// ActionTable is the table that holds the action relation/edge.
	ActionTable = "app_menu"
	// ActionInverseTable is the table name for the AppAction entity.
	// It exists in this package in order to avoid circular dependency with the "appaction" package.
	ActionInverseTable = "app_action"
	// ActionColumn is the table column denoting the action relation/edge.
	ActionColumn = "action_id"
)

// Columns holds all SQL columns for appmenu fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldAppID,
	FieldParentID,
	FieldKind,
	FieldName,
	FieldActionID,
	FieldComments,
	FieldDisplaySort,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int
)

// Kind defines the type for the "kind" enum field.
type Kind string

// Kind values.
const (
	KindDir  Kind = "dir"
	KindMenu Kind = "menu"
)

func (k Kind) String() string {
	return string(k)
}

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindDir, KindMenu:
		return nil
	default:
		return fmt.Errorf("appmenu: invalid enum value for kind field: %q", k)
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
