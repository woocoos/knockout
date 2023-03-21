package types

import (
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
