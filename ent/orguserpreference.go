// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orguserpreference"
	"github.com/woocoos/knockout/ent/user"
)

// OrgUserPreference is the model entity for the OrgUserPreference schema.
type OrgUserPreference struct {
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
	// 用户id
	UserID int `json:"user_id,omitempty"`
	// 组织ID
	OrgID int `json:"org_id,omitempty"`
	// 用户收藏菜单
	MenuFavorite []int `json:"menu_favorite,omitempty"`
	// 用户最近访问菜单
	MenuRecent []int `json:"menu_recent,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrgUserPreferenceQuery when eager-loading is set.
	Edges        OrgUserPreferenceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrgUserPreferenceEdges holds the relations/edges for other nodes in the graph.
type OrgUserPreferenceEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Org holds the value of the org edge.
	Org *Org `json:"org,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgUserPreferenceEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// OrgOrErr returns the Org value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgUserPreferenceEdges) OrgOrErr() (*Org, error) {
	if e.Org != nil {
		return e.Org, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: org.Label}
	}
	return nil, &NotLoadedError{edge: "org"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgUserPreference) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orguserpreference.FieldMenuFavorite, orguserpreference.FieldMenuRecent:
			values[i] = new([]byte)
		case orguserpreference.FieldID, orguserpreference.FieldCreatedBy, orguserpreference.FieldUpdatedBy, orguserpreference.FieldUserID, orguserpreference.FieldOrgID:
			values[i] = new(sql.NullInt64)
		case orguserpreference.FieldCreatedAt, orguserpreference.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgUserPreference fields.
func (oup *OrgUserPreference) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orguserpreference.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oup.ID = int(value.Int64)
		case orguserpreference.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				oup.CreatedBy = int(value.Int64)
			}
		case orguserpreference.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oup.CreatedAt = value.Time
			}
		case orguserpreference.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				oup.UpdatedBy = int(value.Int64)
			}
		case orguserpreference.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oup.UpdatedAt = value.Time
			}
		case orguserpreference.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				oup.UserID = int(value.Int64)
			}
		case orguserpreference.FieldOrgID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_id", values[i])
			} else if value.Valid {
				oup.OrgID = int(value.Int64)
			}
		case orguserpreference.FieldMenuFavorite:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field menu_favorite", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &oup.MenuFavorite); err != nil {
					return fmt.Errorf("unmarshal field menu_favorite: %w", err)
				}
			}
		case orguserpreference.FieldMenuRecent:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field menu_recent", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &oup.MenuRecent); err != nil {
					return fmt.Errorf("unmarshal field menu_recent: %w", err)
				}
			}
		default:
			oup.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrgUserPreference.
// This includes values selected through modifiers, order, etc.
func (oup *OrgUserPreference) Value(name string) (ent.Value, error) {
	return oup.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the OrgUserPreference entity.
func (oup *OrgUserPreference) QueryUser() *UserQuery {
	return NewOrgUserPreferenceClient(oup.config).QueryUser(oup)
}

// QueryOrg queries the "org" edge of the OrgUserPreference entity.
func (oup *OrgUserPreference) QueryOrg() *OrgQuery {
	return NewOrgUserPreferenceClient(oup.config).QueryOrg(oup)
}

// Update returns a builder for updating this OrgUserPreference.
// Note that you need to call OrgUserPreference.Unwrap() before calling this method if this OrgUserPreference
// was returned from a transaction, and the transaction was committed or rolled back.
func (oup *OrgUserPreference) Update() *OrgUserPreferenceUpdateOne {
	return NewOrgUserPreferenceClient(oup.config).UpdateOne(oup)
}

// Unwrap unwraps the OrgUserPreference entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oup *OrgUserPreference) Unwrap() *OrgUserPreference {
	_tx, ok := oup.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgUserPreference is not a transactional entity")
	}
	oup.config.driver = _tx.drv
	return oup
}

// String implements the fmt.Stringer.
func (oup *OrgUserPreference) String() string {
	var builder strings.Builder
	builder.WriteString("OrgUserPreference(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oup.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", oup.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(oup.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", oup.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(oup.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", oup.UserID))
	builder.WriteString(", ")
	builder.WriteString("org_id=")
	builder.WriteString(fmt.Sprintf("%v", oup.OrgID))
	builder.WriteString(", ")
	builder.WriteString("menu_favorite=")
	builder.WriteString(fmt.Sprintf("%v", oup.MenuFavorite))
	builder.WriteString(", ")
	builder.WriteString("menu_recent=")
	builder.WriteString(fmt.Sprintf("%v", oup.MenuRecent))
	builder.WriteByte(')')
	return builder.String()
}

// OrgUserPreferences is a parsable slice of OrgUserPreference.
type OrgUserPreferences []*OrgUserPreference
