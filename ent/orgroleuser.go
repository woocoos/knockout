// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orgroleuser"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/user"
)

// OrgRoleUser is the model entity for the OrgRoleUser schema.
type OrgRoleUser struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int `json:"created_by,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int `json:"updated_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 组织角色ID
	OrgRoleID int `json:"org_role_id,omitempty"`
	// 组织用户ID
	OrgUserID int `json:"org_user_id,omitempty"`
	// 用户ID
	UserID int `json:"user_id,omitempty"`
	// 组织ID
	OrgID int `json:"org_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrgRoleUserQuery when eager-loading is set.
	Edges        OrgRoleUserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrgRoleUserEdges holds the relations/edges for other nodes in the graph.
type OrgRoleUserEdges struct {
	// OrgRole holds the value of the org_role edge.
	OrgRole *OrgRole `json:"org_role,omitempty"`
	// OrgUser holds the value of the org_user edge.
	OrgUser *OrgUser `json:"org_user,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Org holds the value of the org edge.
	Org *Org `json:"org,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int
}

// OrgRoleOrErr returns the OrgRole value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgRoleUserEdges) OrgRoleOrErr() (*OrgRole, error) {
	if e.OrgRole != nil {
		return e.OrgRole, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: orgrole.Label}
	}
	return nil, &NotLoadedError{edge: "org_role"}
}

// OrgUserOrErr returns the OrgUser value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgRoleUserEdges) OrgUserOrErr() (*OrgUser, error) {
	if e.OrgUser != nil {
		return e.OrgUser, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: orguser.Label}
	}
	return nil, &NotLoadedError{edge: "org_user"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgRoleUserEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// OrgOrErr returns the Org value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgRoleUserEdges) OrgOrErr() (*Org, error) {
	if e.Org != nil {
		return e.Org, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: org.Label}
	}
	return nil, &NotLoadedError{edge: "org"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgRoleUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orgroleuser.FieldID, orgroleuser.FieldCreatedBy, orgroleuser.FieldUpdatedBy, orgroleuser.FieldOrgRoleID, orgroleuser.FieldOrgUserID, orgroleuser.FieldUserID, orgroleuser.FieldOrgID:
			values[i] = new(sql.NullInt64)
		case orgroleuser.FieldCreatedAt, orgroleuser.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgRoleUser fields.
func (oru *OrgRoleUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orgroleuser.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oru.ID = int(value.Int64)
		case orgroleuser.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				oru.CreatedBy = int(value.Int64)
			}
		case orgroleuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oru.CreatedAt = value.Time
			}
		case orgroleuser.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				oru.UpdatedBy = int(value.Int64)
			}
		case orgroleuser.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oru.UpdatedAt = value.Time
			}
		case orgroleuser.FieldOrgRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_role_id", values[i])
			} else if value.Valid {
				oru.OrgRoleID = int(value.Int64)
			}
		case orgroleuser.FieldOrgUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_user_id", values[i])
			} else if value.Valid {
				oru.OrgUserID = int(value.Int64)
			}
		case orgroleuser.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				oru.UserID = int(value.Int64)
			}
		case orgroleuser.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				oru.OrgID = int(value.Int64)
			}
		default:
			oru.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrgRoleUser.
// This includes values selected through modifiers, order, etc.
func (oru *OrgRoleUser) Value(name string) (ent.Value, error) {
	return oru.selectValues.Get(name)
}

// QueryOrgRole queries the "org_role" edge of the OrgRoleUser entity.
func (oru *OrgRoleUser) QueryOrgRole() *OrgRoleQuery {
	return NewOrgRoleUserClient(oru.config).QueryOrgRole(oru)
}

// QueryOrgUser queries the "org_user" edge of the OrgRoleUser entity.
func (oru *OrgRoleUser) QueryOrgUser() *OrgUserQuery {
	return NewOrgRoleUserClient(oru.config).QueryOrgUser(oru)
}

// QueryUser queries the "user" edge of the OrgRoleUser entity.
func (oru *OrgRoleUser) QueryUser() *UserQuery {
	return NewOrgRoleUserClient(oru.config).QueryUser(oru)
}

// QueryOrg queries the "org" edge of the OrgRoleUser entity.
func (oru *OrgRoleUser) QueryOrg() *OrgQuery {
	return NewOrgRoleUserClient(oru.config).QueryOrg(oru)
}

// Update returns a builder for updating this OrgRoleUser.
// Note that you need to call OrgRoleUser.Unwrap() before calling this method if this OrgRoleUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (oru *OrgRoleUser) Update() *OrgRoleUserUpdateOne {
	return NewOrgRoleUserClient(oru.config).UpdateOne(oru)
}

// Unwrap unwraps the OrgRoleUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oru *OrgRoleUser) Unwrap() *OrgRoleUser {
	_tx, ok := oru.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgRoleUser is not a transactional entity")
	}
	oru.config.driver = _tx.drv
	return oru
}

// String implements the fmt.Stringer.
func (oru *OrgRoleUser) String() string {
	var builder strings.Builder
	builder.WriteString("OrgRoleUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oru.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", oru.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(oru.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", oru.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(oru.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("org_role_id=")
	builder.WriteString(fmt.Sprintf("%v", oru.OrgRoleID))
	builder.WriteString(", ")
	builder.WriteString("org_user_id=")
	builder.WriteString(fmt.Sprintf("%v", oru.OrgUserID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", oru.UserID))
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", oru.OrgID))
	builder.WriteByte(')')
	return builder.String()
}

// OrgRoleUsers is a parsable slice of OrgRoleUser.
type OrgRoleUsers []*OrgRoleUser
