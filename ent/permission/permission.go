// Code generated by ent, DO NOT EDIT.

package permission

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
	// Label holds the string label denoting the permission type in the database.
	Label = "permission"
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
	// FieldPrincipalKind holds the string denoting the principal_kind field in the database.
	FieldPrincipalKind = "principal_kind"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldOrgPolicyID holds the string denoting the org_policy_id field in the database.
	FieldOrgPolicyID = "org_policy_id"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the permission in the database.
	Table = "permission"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "permission"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organization"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "org_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "permission"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "user"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for permission fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldOrgID,
	FieldPrincipalKind,
	FieldUserID,
	FieldRoleID,
	FieldOrgPolicyID,
	FieldStartAt,
	FieldEndAt,
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
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int
)

// PrincipalKind defines the type for the "principal_kind" enum field.
type PrincipalKind string

// PrincipalKind values.
const (
	PrincipalKindUser PrincipalKind = "user"
	PrincipalKindRole PrincipalKind = "role"
)

func (pk PrincipalKind) String() string {
	return string(pk)
}

// PrincipalKindValidator is a validator for the "principal_kind" field enum values. It is called by the builders before save.
func PrincipalKindValidator(pk PrincipalKind) error {
	switch pk {
	case PrincipalKindUser, PrincipalKindRole:
		return nil
	default:
		return fmt.Errorf("permission: invalid enum value for principal_kind field: %q", pk)
	}
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s typex.SimpleStatus) error {
	switch s.String() {
	case "active", "inactive", "processing":
		return nil
	default:
		return fmt.Errorf("permission: invalid enum value for status field: %q", s)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e PrincipalKind) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *PrincipalKind) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = PrincipalKind(str)
	if err := PrincipalKindValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid PrincipalKind", str)
	}
	return nil
}

var (
	// typex.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*typex.SimpleStatus)(nil)
	// typex.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*typex.SimpleStatus)(nil)
)
