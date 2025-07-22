package types

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent/schema"
	"fmt"
	"io"
	"strconv"
)

type PolicyEffect string

const (
	PolicyEffectAllow PolicyEffect = "allow"
	PolicyEffectDeny  PolicyEffect = "deny"
)

func (e PolicyEffect) String() string {
	return string(e)
}

// PolicyEffectValidator is a validator for the "status" field enum values. It is called by the builders before save.
func PolicyEffectValidator(st PolicyEffect) error {
	switch st {
	case PolicyEffectAllow, PolicyEffectDeny:
		return nil
	default:
		return fmt.Errorf("status: invalid enum value for status field: %q", st)
	}
}

// Values implements field.EnumValues interface
func (PolicyEffect) Values() []string {
	return []string{
		PolicyEffectAllow.String(),
		PolicyEffectDeny.String(),
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e PolicyEffect) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *PolicyEffect) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = PolicyEffect(str)
	if err := PolicyEffectValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid PolicyEffect", str)
	}
	return nil
}

// UserStatus 用于简单型状态字段枚举型
type UserStatus string

// UserType values.
const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusLocked   UserStatus = "locked"
	UserStatusDisabled UserStatus = "disabled"
)

func (st UserStatus) String() string {
	return string(st)
}

// UserStatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func UserStatusValidator(st UserStatus) error {
	switch st {
	case UserStatusActive, UserStatusInactive, UserStatusLocked, UserStatusDisabled:
		return nil
	default:
		return fmt.Errorf("status: invalid enum value for status field: %q", st)
	}
}

// Values implements field.EnumValues interface
func (UserStatus) Values() []string {
	return []string{
		UserStatusActive.String(),
		UserStatusInactive.String(),
		UserStatusLocked.String(),
		UserStatusDisabled.String(),
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (st UserStatus) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(st.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (st *UserStatus) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*st = UserStatus(str)
	if err := UserStatusValidator(*st); err != nil {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (st UserStatus) ProtoAnnotation() schema.Annotation {
	return entproto.Enum(map[string]int32{
		UserStatusActive.String():   1,
		UserStatusInactive.String(): 2,
		UserStatusLocked.String():   3,
		UserStatusDisabled.String(): 4,
	})
}
