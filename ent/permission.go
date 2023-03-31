// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/user"
)

// Permission is the model entity for the Permission schema.
type Permission struct {
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
	// 授权的域根组织
	OrgID int `json:"org_id,omitempty"`
	// 授权类型:角色,用户
	PrincipalKind permission.PrincipalKind `json:"principal_kind,omitempty"`
	// 授权类型为用户的ID
	UserID int `json:"user_id,omitempty"`
	// 授权类型为角色或用户组的ID
	RoleID int `json:"role_id,omitempty"`
	// 策略
	OrgPolicyID int `json:"org_policy_id,omitempty"`
	// 生效开始时间
	StartAt time.Time `json:"start_at,omitempty"`
	// 生效结束时间
	EndAt time.Time `json:"end_at,omitempty"`
	// 状态
	Status typex.SimpleStatus `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionQuery when eager-loading is set.
	Edges PermissionEdges `json:"edges"`
}

// PermissionEdges holds the relations/edges for other nodes in the graph.
type PermissionEdges struct {
	// Org holds the value of the org edge.
	Org *Org `json:"org,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// OrgPolicy holds the value of the org_policy edge.
	OrgPolicy *OrgPolicy `json:"org_policy,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int
}

// OrgOrErr returns the Org value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionEdges) OrgOrErr() (*Org, error) {
	if e.loadedTypes[0] {
		if e.Org == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: org.Label}
		}
		return e.Org, nil
	}
	return nil, &NotLoadedError{edge: "org"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// OrgPolicyOrErr returns the OrgPolicy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionEdges) OrgPolicyOrErr() (*OrgPolicy, error) {
	if e.loadedTypes[2] {
		if e.OrgPolicy == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: orgpolicy.Label}
		}
		return e.OrgPolicy, nil
	}
	return nil, &NotLoadedError{edge: "org_policy"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Permission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permission.FieldID, permission.FieldCreatedBy, permission.FieldUpdatedBy, permission.FieldOrgID, permission.FieldUserID, permission.FieldRoleID, permission.FieldOrgPolicyID:
			values[i] = new(sql.NullInt64)
		case permission.FieldPrincipalKind, permission.FieldStatus:
			values[i] = new(sql.NullString)
		case permission.FieldCreatedAt, permission.FieldUpdatedAt, permission.FieldStartAt, permission.FieldEndAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Permission", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Permission fields.
func (pe *Permission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case permission.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pe.CreatedBy = int(value.Int64)
			}
		case permission.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pe.CreatedAt = value.Time
			}
		case permission.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pe.UpdatedBy = int(value.Int64)
			}
		case permission.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pe.UpdatedAt = value.Time
			}
		case permission.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				pe.OrgID = int(value.Int64)
			}
		case permission.FieldPrincipalKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field principal_kind", values[i])
			} else if value.Valid {
				pe.PrincipalKind = permission.PrincipalKind(value.String)
			}
		case permission.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				pe.UserID = int(value.Int64)
			}
		case permission.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				pe.RoleID = int(value.Int64)
			}
		case permission.FieldOrgPolicyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_policy_id", values[i])
			} else if value.Valid {
				pe.OrgPolicyID = int(value.Int64)
			}
		case permission.FieldStartAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				pe.StartAt = value.Time
			}
		case permission.FieldEndAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_at", values[i])
			} else if value.Valid {
				pe.EndAt = value.Time
			}
		case permission.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pe.Status = typex.SimpleStatus(value.String)
			}
		}
	}
	return nil
}

// QueryOrg queries the "org" edge of the Permission entity.
func (pe *Permission) QueryOrg() *OrgQuery {
	return NewPermissionClient(pe.config).QueryOrg(pe)
}

// QueryUser queries the "user" edge of the Permission entity.
func (pe *Permission) QueryUser() *UserQuery {
	return NewPermissionClient(pe.config).QueryUser(pe)
}

// QueryOrgPolicy queries the "org_policy" edge of the Permission entity.
func (pe *Permission) QueryOrgPolicy() *OrgPolicyQuery {
	return NewPermissionClient(pe.config).QueryOrgPolicy(pe)
}

// Update returns a builder for updating this Permission.
// Note that you need to call Permission.Unwrap() before calling this method if this Permission
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Permission) Update() *PermissionUpdateOne {
	return NewPermissionClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Permission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Permission) Unwrap() *Permission {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Permission is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Permission) String() string {
	var builder strings.Builder
	builder.WriteString("Permission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", pe.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", pe.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pe.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.OrgID))
	builder.WriteString(", ")
	builder.WriteString("principal_kind=")
	builder.WriteString(fmt.Sprintf("%v", pe.PrincipalKind))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.UserID))
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.RoleID))
	builder.WriteString(", ")
	builder.WriteString("org_policy_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.OrgPolicyID))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(pe.StartAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_at=")
	builder.WriteString(pe.EndAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pe.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Permissions is a parsable slice of Permission.
type Permissions []*Permission
