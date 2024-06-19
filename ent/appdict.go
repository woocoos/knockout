// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appdict"
)

// AppDict is the model entity for the AppDict schema.
type AppDict struct {
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
	// 用于标识应用资源的唯一代码,尽量简短
	Code string `json:"code,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 备注
	Comments string `json:"comments,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AppDictQuery when eager-loading is set.
	Edges        AppDictEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AppDictEdges holds the relations/edges for other nodes in the graph.
type AppDictEdges struct {
	// App holds the value of the app edge.
	App *App `json:"app,omitempty"`
	// Items holds the value of the items edge.
	Items []*AppDictItem `json:"items,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedItems map[string][]*AppDictItem
}

// AppOrErr returns the App value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AppDictEdges) AppOrErr() (*App, error) {
	if e.App != nil {
		return e.App, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: app.Label}
	}
	return nil, &NotLoadedError{edge: "app"}
}

// ItemsOrErr returns the Items value or an error if the edge
// was not loaded in eager-loading.
func (e AppDictEdges) ItemsOrErr() ([]*AppDictItem, error) {
	if e.loadedTypes[1] {
		return e.Items, nil
	}
	return nil, &NotLoadedError{edge: "items"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppDict) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case appdict.FieldID, appdict.FieldCreatedBy, appdict.FieldUpdatedBy, appdict.FieldAppID:
			values[i] = new(sql.NullInt64)
		case appdict.FieldCode, appdict.FieldName, appdict.FieldComments:
			values[i] = new(sql.NullString)
		case appdict.FieldCreatedAt, appdict.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppDict fields.
func (ad *AppDict) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appdict.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ad.ID = int(value.Int64)
		case appdict.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ad.CreatedBy = int(value.Int64)
			}
		case appdict.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ad.CreatedAt = value.Time
			}
		case appdict.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ad.UpdatedBy = int(value.Int64)
			}
		case appdict.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ad.UpdatedAt = value.Time
			}
		case appdict.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				ad.AppID = int(value.Int64)
			}
		case appdict.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				ad.Code = value.String
			}
		case appdict.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ad.Name = value.String
			}
		case appdict.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				ad.Comments = value.String
			}
		default:
			ad.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AppDict.
// This includes values selected through modifiers, order, etc.
func (ad *AppDict) Value(name string) (ent.Value, error) {
	return ad.selectValues.Get(name)
}

// QueryApp queries the "app" edge of the AppDict entity.
func (ad *AppDict) QueryApp() *AppQuery {
	return NewAppDictClient(ad.config).QueryApp(ad)
}

// QueryItems queries the "items" edge of the AppDict entity.
func (ad *AppDict) QueryItems() *AppDictItemQuery {
	return NewAppDictClient(ad.config).QueryItems(ad)
}

// Update returns a builder for updating this AppDict.
// Note that you need to call AppDict.Unwrap() before calling this method if this AppDict
// was returned from a transaction, and the transaction was committed or rolled back.
func (ad *AppDict) Update() *AppDictUpdateOne {
	return NewAppDictClient(ad.config).UpdateOne(ad)
}

// Unwrap unwraps the AppDict entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ad *AppDict) Unwrap() *AppDict {
	_tx, ok := ad.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppDict is not a transactional entity")
	}
	ad.config.driver = _tx.drv
	return ad
}

// String implements the fmt.Stringer.
func (ad *AppDict) String() string {
	var builder strings.Builder
	builder.WriteString("AppDict(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ad.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ad.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(ad.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ad.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ad.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ad.AppID))
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(ad.Code)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ad.Name)
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(ad.Comments)
	builder.WriteByte(')')
	return builder.String()
}

// NamedItems returns the Items named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ad *AppDict) NamedItems(name string) ([]*AppDictItem, error) {
	if ad.Edges.namedItems == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ad.Edges.namedItems[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ad *AppDict) appendNamedItems(name string, edges ...*AppDictItem) {
	if ad.Edges.namedItems == nil {
		ad.Edges.namedItems = make(map[string][]*AppDictItem)
	}
	if len(edges) == 0 {
		ad.Edges.namedItems[name] = []*AppDictItem{}
	} else {
		ad.Edges.namedItems[name] = append(ad.Edges.namedItems[name], edges...)
	}
}

// AppDicts is a parsable slice of AppDict.
type AppDicts []*AppDict
