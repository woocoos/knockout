// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
)

// AppAction is the model entity for the AppAction schema.
type AppAction struct {
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
	// 名称
	Name string `json:"name,omitempty"`
	// restful,graphql,rpc
	Kind appaction.Kind `json:"kind,omitempty"`
	// 操作方法:读,写,列表
	Method appaction.Method `json:"method,omitempty"`
	// 备注
	Comments string `json:"comments,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppActionQuery when eager-loading is set.
	Edges        AppActionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AppActionEdges holds the relations/edges for other nodes in the graph.
type AppActionEdges struct {
	// App holds the value of the app edge.
	App *App `json:"app,omitempty"`
	// 被引用的菜单项
	Menus []*AppMenu `json:"menus,omitempty"`
	// 引用的资源
	Resources []*AppRes `json:"resources,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedMenus     map[string][]*AppMenu
	namedResources map[string][]*AppRes
}

// AppOrErr returns the App value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppActionEdges) AppOrErr() (*App, error) {
	if e.loadedTypes[0] {
		if e.App == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: app.Label}
		}
		return e.App, nil
	}
	return nil, &NotLoadedError{edge: "app"}
}

// MenusOrErr returns the Menus value or an error if the edge
// was not loaded in eager-loading.
func (e AppActionEdges) MenusOrErr() ([]*AppMenu, error) {
	if e.loadedTypes[1] {
		return e.Menus, nil
	}
	return nil, &NotLoadedError{edge: "menus"}
}

// ResourcesOrErr returns the Resources value or an error if the edge
// was not loaded in eager-loading.
func (e AppActionEdges) ResourcesOrErr() ([]*AppRes, error) {
	if e.loadedTypes[2] {
		return e.Resources, nil
	}
	return nil, &NotLoadedError{edge: "resources"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppAction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case appaction.FieldID, appaction.FieldCreatedBy, appaction.FieldUpdatedBy, appaction.FieldAppID:
			values[i] = new(sql.NullInt64)
		case appaction.FieldName, appaction.FieldKind, appaction.FieldMethod, appaction.FieldComments:
			values[i] = new(sql.NullString)
		case appaction.FieldCreatedAt, appaction.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppAction fields.
func (aa *AppAction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aa.ID = int(value.Int64)
		case appaction.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				aa.CreatedBy = int(value.Int64)
			}
		case appaction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				aa.CreatedAt = value.Time
			}
		case appaction.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				aa.UpdatedBy = int(value.Int64)
			}
		case appaction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				aa.UpdatedAt = value.Time
			}
		case appaction.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				aa.AppID = int(value.Int64)
			}
		case appaction.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				aa.Name = value.String
			}
		case appaction.FieldKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kind", values[i])
			} else if value.Valid {
				aa.Kind = appaction.Kind(value.String)
			}
		case appaction.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				aa.Method = appaction.Method(value.String)
			}
		case appaction.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				aa.Comments = value.String
			}
		default:
			aa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AppAction.
// This includes values selected through modifiers, order, etc.
func (aa *AppAction) Value(name string) (ent.Value, error) {
	return aa.selectValues.Get(name)
}

// QueryApp queries the "app" edge of the AppAction entity.
func (aa *AppAction) QueryApp() *AppQuery {
	return NewAppActionClient(aa.config).QueryApp(aa)
}

// QueryMenus queries the "menus" edge of the AppAction entity.
func (aa *AppAction) QueryMenus() *AppMenuQuery {
	return NewAppActionClient(aa.config).QueryMenus(aa)
}

// QueryResources queries the "resources" edge of the AppAction entity.
func (aa *AppAction) QueryResources() *AppResQuery {
	return NewAppActionClient(aa.config).QueryResources(aa)
}

// Update returns a builder for updating this AppAction.
// Note that you need to call AppAction.Unwrap() before calling this method if this AppAction
// was returned from a transaction, and the transaction was committed or rolled back.
func (aa *AppAction) Update() *AppActionUpdateOne {
	return NewAppActionClient(aa.config).UpdateOne(aa)
}

// Unwrap unwraps the AppAction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aa *AppAction) Unwrap() *AppAction {
	_tx, ok := aa.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppAction is not a transactional entity")
	}
	aa.config.driver = _tx.drv
	return aa
}

// String implements the fmt.Stringer.
func (aa *AppAction) String() string {
	var builder strings.Builder
	builder.WriteString("AppAction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", aa.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", aa.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(aa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", aa.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(aa.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", aa.AppID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(aa.Name)
	builder.WriteString(", ")
	builder.WriteString("kind=")
	builder.WriteString(fmt.Sprintf("%v", aa.Kind))
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(fmt.Sprintf("%v", aa.Method))
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(aa.Comments)
	builder.WriteByte(')')
	return builder.String()
}

// NamedMenus returns the Menus named value or an error if the edge was not
// loaded in eager-loading with this name.
func (aa *AppAction) NamedMenus(name string) ([]*AppMenu, error) {
	if aa.Edges.namedMenus == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := aa.Edges.namedMenus[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (aa *AppAction) appendNamedMenus(name string, edges ...*AppMenu) {
	if aa.Edges.namedMenus == nil {
		aa.Edges.namedMenus = make(map[string][]*AppMenu)
	}
	if len(edges) == 0 {
		aa.Edges.namedMenus[name] = []*AppMenu{}
	} else {
		aa.Edges.namedMenus[name] = append(aa.Edges.namedMenus[name], edges...)
	}
}

// NamedResources returns the Resources named value or an error if the edge was not
// loaded in eager-loading with this name.
func (aa *AppAction) NamedResources(name string) ([]*AppRes, error) {
	if aa.Edges.namedResources == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := aa.Edges.namedResources[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (aa *AppAction) appendNamedResources(name string, edges ...*AppRes) {
	if aa.Edges.namedResources == nil {
		aa.Edges.namedResources = make(map[string][]*AppRes)
	}
	if len(edges) == 0 {
		aa.Edges.namedResources[name] = []*AppRes{}
	} else {
		aa.Edges.namedResources[name] = append(aa.Edges.namedResources[name], edges...)
	}
}

// AppActions is a parsable slice of AppAction.
type AppActions []*AppAction
