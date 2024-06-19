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
)

// OrgRole is the model entity for the OrgRole schema.
type OrgRole struct {
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
	// 类型,group:组,role:角色
	Kind orgrole.Kind `json:"kind,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 角色ID,如有表示该角色来源于应用角色
	AppRoleID int `json:"app_role_id,omitempty"`
	// 备注
	Comments string `json:"comments,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrgRoleQuery when eager-loading is set.
	Edges        OrgRoleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrgRoleEdges holds the relations/edges for other nodes in the graph.
type OrgRoleEdges struct {
	// Org holds the value of the org edge.
	Org *Org `json:"org,omitempty"`
	// 组用户
	OrgUsers []*OrgUser `json:"org_users,omitempty"`
	// OrgRoleUser holds the value of the org_role_user edge.
	OrgRoleUser []*OrgRoleUser `json:"org_role_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool

	namedOrgUsers    map[string][]*OrgUser
	namedOrgRoleUser map[string][]*OrgRoleUser
}

// OrgOrErr returns the Org value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgRoleEdges) OrgOrErr() (*Org, error) {
	if e.Org != nil {
		return e.Org, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: org.Label}
	}
	return nil, &NotLoadedError{edge: "org"}
}

// OrgUsersOrErr returns the OrgUsers value or an error if the edge
// was not loaded in eager-loading.
func (e OrgRoleEdges) OrgUsersOrErr() ([]*OrgUser, error) {
	if e.loadedTypes[1] {
		return e.OrgUsers, nil
	}
	return nil, &NotLoadedError{edge: "org_users"}
}

// OrgRoleUserOrErr returns the OrgRoleUser value or an error if the edge
// was not loaded in eager-loading.
func (e OrgRoleEdges) OrgRoleUserOrErr() ([]*OrgRoleUser, error) {
	if e.loadedTypes[2] {
		return e.OrgRoleUser, nil
	}
	return nil, &NotLoadedError{edge: "org_role_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgRole) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orgrole.FieldID, orgrole.FieldCreatedBy, orgrole.FieldUpdatedBy, orgrole.FieldOrgID, orgrole.FieldAppRoleID:
			values[i] = new(sql.NullInt64)
		case orgrole.FieldKind, orgrole.FieldName, orgrole.FieldComments:
			values[i] = new(sql.NullString)
		case orgrole.FieldCreatedAt, orgrole.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgRole fields.
func (or *OrgRole) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orgrole.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			or.ID = int(value.Int64)
		case orgrole.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				or.CreatedBy = int(value.Int64)
			}
		case orgrole.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				or.CreatedAt = value.Time
			}
		case orgrole.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				or.UpdatedBy = int(value.Int64)
			}
		case orgrole.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				or.UpdatedAt = value.Time
			}
		case orgrole.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				or.OrgID = int(value.Int64)
			}
		case orgrole.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				or.Kind = orgrole.Kind(value.String)
			}
		case orgrole.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				or.Name = value.String
			}
		case orgrole.FieldAppRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_role_id", values[i])
			} else if value.Valid {
				or.AppRoleID = int(value.Int64)
			}
		case orgrole.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				or.Comments = value.String
			}
		default:
			or.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrgRole.
// This includes values selected through modifiers, order, etc.
func (or *OrgRole) Value(name string) (ent.Value, error) {
	return or.selectValues.Get(name)
}

// QueryOrg queries the "org" edge of the OrgRole entity.
func (or *OrgRole) QueryOrg() *OrgQuery {
	return NewOrgRoleClient(or.config).QueryOrg(or)
}

// QueryOrgUsers queries the "org_users" edge of the OrgRole entity.
func (or *OrgRole) QueryOrgUsers() *OrgUserQuery {
	return NewOrgRoleClient(or.config).QueryOrgUsers(or)
}

// QueryOrgRoleUser queries the "org_role_user" edge of the OrgRole entity.
func (or *OrgRole) QueryOrgRoleUser() *OrgRoleUserQuery {
	return NewOrgRoleClient(or.config).QueryOrgRoleUser(or)
}

// Update returns a builder for updating this OrgRole.
// Note that you need to call OrgRole.Unwrap() before calling this method if this OrgRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (or *OrgRole) Update() *OrgRoleUpdateOne {
	return NewOrgRoleClient(or.config).UpdateOne(or)
}

// Unwrap unwraps the OrgRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (or *OrgRole) Unwrap() *OrgRole {
	_tx, ok := or.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgRole is not a transactional entity")
	}
	or.config.driver = _tx.drv
	return or
}

// String implements the fmt.Stringer.
func (or *OrgRole) String() string {
	var builder strings.Builder
	builder.WriteString("OrgRole(")
	builder.WriteString(fmt.Sprintf("id=%v, ", or.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", or.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(or.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", or.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(or.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", or.OrgID))
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(fmt.Sprintf("%v", or.Kind))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(or.Name)
	builder.WriteString(", ")
	builder.WriteString("app_role_id=")
	builder.WriteString(fmt.Sprintf("%v", or.AppRoleID))
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(or.Comments)
	builder.WriteByte(')')
	return builder.String()
}

// NamedOrgUsers returns the OrgUsers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (or *OrgRole) NamedOrgUsers(name string) ([]*OrgUser, error) {
	if or.Edges.namedOrgUsers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := or.Edges.namedOrgUsers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (or *OrgRole) appendNamedOrgUsers(name string, edges ...*OrgUser) {
	if or.Edges.namedOrgUsers == nil {
		or.Edges.namedOrgUsers = make(map[string][]*OrgUser)
	}
	if len(edges) == 0 {
		or.Edges.namedOrgUsers[name] = []*OrgUser{}
	} else {
		or.Edges.namedOrgUsers[name] = append(or.Edges.namedOrgUsers[name], edges...)
	}
}

// NamedOrgRoleUser returns the OrgRoleUser named value or an error if the edge was not
// loaded in eager-loading with this name.
func (or *OrgRole) NamedOrgRoleUser(name string) ([]*OrgRoleUser, error) {
	if or.Edges.namedOrgRoleUser == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := or.Edges.namedOrgRoleUser[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (or *OrgRole) appendNamedOrgRoleUser(name string, edges ...*OrgRoleUser) {
	if or.Edges.namedOrgRoleUser == nil {
		or.Edges.namedOrgRoleUser = make(map[string][]*OrgRoleUser)
	}
	if len(edges) == 0 {
		or.Edges.namedOrgRoleUser[name] = []*OrgRoleUser{}
	} else {
		or.Edges.namedOrgRoleUser[name] = append(or.Edges.namedOrgRoleUser[name], edges...)
	}
}

// OrgRoles is a parsable slice of OrgRole.
type OrgRoles []*OrgRole
