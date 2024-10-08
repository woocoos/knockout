// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/approle"
)

// AppRole is the model entity for the AppRole schema.
type AppRole struct {
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
	// 所属应用
	AppID int `json:"app_id,omitempty"`
	// 角色名称
	Name string `json:"name,omitempty"`
	// 备注
	Comments string `json:"comments,omitempty"`
	// 标识是否自动授予到账户
	AutoGrant bool `json:"auto_grant,omitempty"`
	// 授权后是否可编辑
	Editable bool `json:"editable,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppRoleQuery when eager-loading is set.
	Edges        AppRoleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AppRoleEdges holds the relations/edges for other nodes in the graph.
type AppRoleEdges struct {
	// App holds the value of the app edge.
	App *App `json:"app,omitempty"`
	// 权限授权策略
	Policies []*AppPolicy `json:"policies,omitempty"`
	// AppRolePolicy holds the value of the app_role_policy edge.
	AppRolePolicy []*AppRolePolicy `json:"app_role_policy,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedPolicies      map[string][]*AppPolicy
	namedAppRolePolicy map[string][]*AppRolePolicy
}

// AppOrErr returns the App value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppRoleEdges) AppOrErr() (*App, error) {
	if e.App != nil {
		return e.App, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: app.Label}
	}
	return nil, &NotLoadedError{edge: "app"}
}

// PoliciesOrErr returns the Policies value or an error if the edge
// was not loaded in eager-loading.
func (e AppRoleEdges) PoliciesOrErr() ([]*AppPolicy, error) {
	if e.loadedTypes[1] {
		return e.Policies, nil
	}
	return nil, &NotLoadedError{edge: "policies"}
}

// AppRolePolicyOrErr returns the AppRolePolicy value or an error if the edge
// was not loaded in eager-loading.
func (e AppRoleEdges) AppRolePolicyOrErr() ([]*AppRolePolicy, error) {
	if e.loadedTypes[2] {
		return e.AppRolePolicy, nil
	}
	return nil, &NotLoadedError{edge: "app_role_policy"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppRole) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case approle.FieldAutoGrant, approle.FieldEditable:
			values[i] = new(sql.NullBool)
		case approle.FieldID, approle.FieldCreatedBy, approle.FieldUpdatedBy, approle.FieldAppID:
			values[i] = new(sql.NullInt64)
		case approle.FieldName, approle.FieldComments:
			values[i] = new(sql.NullString)
		case approle.FieldCreatedAt, approle.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppRole fields.
func (ar *AppRole) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case approle.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ar.ID = int(value.Int64)
		case approle.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ar.CreatedBy = int(value.Int64)
			}
		case approle.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ar.CreatedAt = value.Time
			}
		case approle.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ar.UpdatedBy = int(value.Int64)
			}
		case approle.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ar.UpdatedAt = value.Time
			}
		case approle.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				ar.AppID = int(value.Int64)
			}
		case approle.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ar.Name = value.String
			}
		case approle.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				ar.Comments = value.String
			}
		case approle.FieldAutoGrant:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field auto_grant", values[i])
			} else if value.Valid {
				ar.AutoGrant = value.Bool
			}
		case approle.FieldEditable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field editable", values[i])
			} else if value.Valid {
				ar.Editable = value.Bool
			}
		default:
			ar.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AppRole.
// This includes values selected through modifiers, order, etc.
func (ar *AppRole) Value(name string) (ent.Value, error) {
	return ar.selectValues.Get(name)
}

// QueryApp queries the "app" edge of the AppRole entity.
func (ar *AppRole) QueryApp() *AppQuery {
	return NewAppRoleClient(ar.config).QueryApp(ar)
}

// QueryPolicies queries the "policies" edge of the AppRole entity.
func (ar *AppRole) QueryPolicies() *AppPolicyQuery {
	return NewAppRoleClient(ar.config).QueryPolicies(ar)
}

// QueryAppRolePolicy queries the "app_role_policy" edge of the AppRole entity.
func (ar *AppRole) QueryAppRolePolicy() *AppRolePolicyQuery {
	return NewAppRoleClient(ar.config).QueryAppRolePolicy(ar)
}

// Update returns a builder for updating this AppRole.
// Note that you need to call AppRole.Unwrap() before calling this method if this AppRole
// was returned from a transaction, and the transaction was committed or rolled back.
func (ar *AppRole) Update() *AppRoleUpdateOne {
	return NewAppRoleClient(ar.config).UpdateOne(ar)
}

// Unwrap unwraps the AppRole entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ar *AppRole) Unwrap() *AppRole {
	_tx, ok := ar.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppRole is not a transactional entity")
	}
	ar.config.driver = _tx.drv
	return ar
}

// String implements the fmt.Stringer.
func (ar *AppRole) String() string {
	var builder strings.Builder
	builder.WriteString("AppRole(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ar.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ar.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ar.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ar.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ar.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ar.AppID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ar.Name)
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(ar.Comments)
	builder.WriteString(", ")
	builder.WriteString("auto_grant=")
	builder.WriteString(fmt.Sprintf("%v", ar.AutoGrant))
	builder.WriteString(", ")
	builder.WriteString("editable=")
	builder.WriteString(fmt.Sprintf("%v", ar.Editable))
	builder.WriteByte(')')
	return builder.String()
}

// NamedPolicies returns the Policies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ar *AppRole) NamedPolicies(name string) ([]*AppPolicy, error) {
	if ar.Edges.namedPolicies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ar.Edges.namedPolicies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ar *AppRole) appendNamedPolicies(name string, edges ...*AppPolicy) {
	if ar.Edges.namedPolicies == nil {
		ar.Edges.namedPolicies = make(map[string][]*AppPolicy)
	}
	if len(edges) == 0 {
		ar.Edges.namedPolicies[name] = []*AppPolicy{}
	} else {
		ar.Edges.namedPolicies[name] = append(ar.Edges.namedPolicies[name], edges...)
	}
}

// NamedAppRolePolicy returns the AppRolePolicy named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ar *AppRole) NamedAppRolePolicy(name string) ([]*AppRolePolicy, error) {
	if ar.Edges.namedAppRolePolicy == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ar.Edges.namedAppRolePolicy[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ar *AppRole) appendNamedAppRolePolicy(name string, edges ...*AppRolePolicy) {
	if ar.Edges.namedAppRolePolicy == nil {
		ar.Edges.namedAppRolePolicy = make(map[string][]*AppRolePolicy)
	}
	if len(edges) == 0 {
		ar.Edges.namedAppRolePolicy[name] = []*AppRolePolicy{}
	} else {
		ar.Edges.namedAppRolePolicy[name] = append(ar.Edges.namedAppRolePolicy[name], edges...)
	}
}

// AppRoles is a parsable slice of AppRole.
type AppRoles []*AppRole
