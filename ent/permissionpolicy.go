// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/knockout/ent/organization"
	"github.com/woocoos/knockout/ent/permissionpolicy"
	"github.com/woocoos/knockout/graph/entgen/types"
)

// PermissionPolicy is the model entity for the PermissionPolicy schema.
type PermissionPolicy struct {
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
	// 组织ID
	OrgID int `json:"org_id,omitempty"`
	// 所属应用
	AppID int `json:"app_id,omitempty"`
	// 所属应用策略,如果是自定义应用策略,则为空
	AppPolicyID int `json:"app_policy_id,omitempty"`
	// 策略名称
	Name string `json:"name,omitempty"`
	// 描述
	Comments string `json:"comments,omitempty"`
	// 策略规则,如果是应用策略,则为空
	Rules []types.PolicyRule `json:"rules,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PermissionPolicyQuery when eager-loading is set.
	Edges PermissionPolicyEdges `json:"edges"`
}

// PermissionPolicyEdges holds the relations/edges for other nodes in the graph.
type PermissionPolicyEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PermissionPolicyEdges) OrganizationOrErr() (*Organization, error) {
	if e.loadedTypes[0] {
		if e.Organization == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PermissionPolicy) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case permissionpolicy.FieldRules:
			values[i] = new([]byte)
		case permissionpolicy.FieldID, permissionpolicy.FieldCreatedBy, permissionpolicy.FieldUpdatedBy, permissionpolicy.FieldOrgID, permissionpolicy.FieldAppID, permissionpolicy.FieldAppPolicyID:
			values[i] = new(sql.NullInt64)
		case permissionpolicy.FieldName, permissionpolicy.FieldComments:
			values[i] = new(sql.NullString)
		case permissionpolicy.FieldCreatedAt, permissionpolicy.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PermissionPolicy", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PermissionPolicy fields.
func (pp *PermissionPolicy) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case permissionpolicy.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pp.ID = int(value.Int64)
		case permissionpolicy.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pp.CreatedBy = int(value.Int64)
			}
		case permissionpolicy.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pp.CreatedAt = value.Time
			}
		case permissionpolicy.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pp.UpdatedBy = int(value.Int64)
			}
		case permissionpolicy.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pp.UpdatedAt = value.Time
			}
		case permissionpolicy.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				pp.OrgID = int(value.Int64)
			}
		case permissionpolicy.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				pp.AppID = int(value.Int64)
			}
		case permissionpolicy.FieldAppPolicyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_policy_id", values[i])
			} else if value.Valid {
				pp.AppPolicyID = int(value.Int64)
			}
		case permissionpolicy.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pp.Name = value.String
			}
		case permissionpolicy.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				pp.Comments = value.String
			}
		case permissionpolicy.FieldRules:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field rules", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pp.Rules); err != nil {
					return fmt.Errorf("unmarshal field rules: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryOrganization queries the "organization" edge of the PermissionPolicy entity.
func (pp *PermissionPolicy) QueryOrganization() *OrganizationQuery {
	return NewPermissionPolicyClient(pp.config).QueryOrganization(pp)
}

// Update returns a builder for updating this PermissionPolicy.
// Note that you need to call PermissionPolicy.Unwrap() before calling this method if this PermissionPolicy
// was returned from a transaction, and the transaction was committed or rolled back.
func (pp *PermissionPolicy) Update() *PermissionPolicyUpdateOne {
	return NewPermissionPolicyClient(pp.config).UpdateOne(pp)
}

// Unwrap unwraps the PermissionPolicy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pp *PermissionPolicy) Unwrap() *PermissionPolicy {
	_tx, ok := pp.config.driver.(*txDriver)
	if !ok {
		panic("ent: PermissionPolicy is not a transactional entity")
	}
	pp.config.driver = _tx.drv
	return pp
}

// String implements the fmt.Stringer.
func (pp *PermissionPolicy) String() string {
	var builder strings.Builder
	builder.WriteString("PermissionPolicy(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pp.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", pp.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", pp.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", pp.OrgID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", pp.AppID))
	builder.WriteString(", ")
	builder.WriteString("app_policy_id=")
	builder.WriteString(fmt.Sprintf("%v", pp.AppPolicyID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pp.Name)
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(pp.Comments)
	builder.WriteString(", ")
	builder.WriteString("rules=")
	builder.WriteString(fmt.Sprintf("%v", pp.Rules))
	builder.WriteByte(')')
	return builder.String()
}

// PermissionPolicies is a parsable slice of PermissionPolicy.
type PermissionPolicies []*PermissionPolicy
