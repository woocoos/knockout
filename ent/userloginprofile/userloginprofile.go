// Code generated by ent, DO NOT EDIT.

package userloginprofile

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
	// Label holds the string label denoting the userloginprofile type in the database.
	Label = "user_login_profile"
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
	// FieldLastLoginIP holds the string denoting the last_login_ip field in the database.
	FieldLastLoginIP = "last_login_ip"
	// FieldLastLoginAt holds the string denoting the last_login_at field in the database.
	FieldLastLoginAt = "last_login_at"
	// FieldCanLogin holds the string denoting the can_login field in the database.
	FieldCanLogin = "can_login"
	// FieldSetKind holds the string denoting the set_kind field in the database.
	FieldSetKind = "set_kind"
	// FieldPasswordReset holds the string denoting the password_reset field in the database.
	FieldPasswordReset = "password_reset"
	// FieldVerifyDevice holds the string denoting the verify_device field in the database.
	FieldVerifyDevice = "verify_device"
	// FieldMfaEnabled holds the string denoting the mfa_enabled field in the database.
	FieldMfaEnabled = "mfa_enabled"
	// FieldMfaSecret holds the string denoting the mfa_secret field in the database.
	FieldMfaSecret = "mfa_secret"
	// FieldMfaStatus holds the string denoting the mfa_status field in the database.
	FieldMfaStatus = "mfa_status"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the userloginprofile in the database.
	Table = "user_login_profile"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_login_profile"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "user"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for userloginprofile fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldUserID,
	FieldLastLoginIP,
	FieldLastLoginAt,
	FieldCanLogin,
	FieldSetKind,
	FieldPasswordReset,
	FieldVerifyDevice,
	FieldMfaEnabled,
	FieldMfaSecret,
	FieldMfaStatus,
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
	// MfaSecretValidator is a validator for the "mfa_secret" field. It is called by the builders before save.
	MfaSecretValidator func(string) error
)

// SetKind defines the type for the "set_kind" enum field.
type SetKind string

// SetKind values.
const (
	SetKindKeep     SetKind = "keep"
	SetKindCustomer SetKind = "customer"
	SetKindAuto     SetKind = "auto"
)

func (sk SetKind) String() string {
	return string(sk)
}

// SetKindValidator is a validator for the "set_kind" field enum values. It is called by the builders before save.
func SetKindValidator(sk SetKind) error {
	switch sk {
	case SetKindKeep, SetKindCustomer, SetKindAuto:
		return nil
	default:
		return fmt.Errorf("userloginprofile: invalid enum value for set_kind field: %q", sk)
	}
}

// MfaStatusValidator is a validator for the "mfa_status" field enum values. It is called by the builders before save.
func MfaStatusValidator(ms typex.SimpleStatus) error {
	switch ms.String() {
	case "active", "inactive", "processing":
		return nil
	default:
		return fmt.Errorf("userloginprofile: invalid enum value for mfa_status field: %q", ms)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e SetKind) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *SetKind) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = SetKind(str)
	if err := SetKindValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid SetKind", str)
	}
	return nil
}

var (
	// typex.SimpleStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*typex.SimpleStatus)(nil)
	// typex.SimpleStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*typex.SimpleStatus)(nil)
)
