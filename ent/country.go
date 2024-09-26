// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/ent/country"
)

// Country is the model entity for the Country schema.
type Country struct {
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
	// 国家中文名称
	Name string `json:"name,omitempty"`
	// 国家英文名称
	NameEn string `json:"name_en,omitempty"`
	// 国家代码
	Code string `json:"code,omitempty"`
	// DisplaySort holds the value of the "display_sort" field.
	DisplaySort int32 `json:"display_sort,omitempty"`
	// 状态
	Status typex.SimpleStatus `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CountryQuery when eager-loading is set.
	Edges        CountryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CountryEdges holds the relations/edges for other nodes in the graph.
type CountryEdges struct {
	// 地区信息
	Regions []*Region `json:"regions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedRegions map[string][]*Region
}

// RegionsOrErr returns the Regions value or an error if the edge
// was not loaded in eager-loading.
func (e CountryEdges) RegionsOrErr() ([]*Region, error) {
	if e.loadedTypes[0] {
		return e.Regions, nil
	}
	return nil, &NotLoadedError{edge: "regions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Country) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case country.FieldID, country.FieldCreatedBy, country.FieldUpdatedBy, country.FieldDisplaySort:
			values[i] = new(sql.NullInt64)
		case country.FieldName, country.FieldNameEn, country.FieldCode, country.FieldStatus:
			values[i] = new(sql.NullString)
		case country.FieldCreatedAt, country.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Country fields.
func (c *Country) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case country.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case country.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				c.CreatedBy = int(value.Int64)
			}
		case country.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case country.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				c.UpdatedBy = int(value.Int64)
			}
		case country.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case country.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case country.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				c.NameEn = value.String
			}
		case country.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				c.Code = value.String
			}
		case country.FieldDisplaySort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field display_sort", values[i])
			} else if value.Valid {
				c.DisplaySort = int32(value.Int64)
			}
		case country.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = typex.SimpleStatus(value.String)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Country.
// This includes values selected through modifiers, order, etc.
func (c *Country) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryRegions queries the "regions" edge of the Country entity.
func (c *Country) QueryRegions() *RegionQuery {
	return NewCountryClient(c.config).QueryRegions(c)
}

// Update returns a builder for updating this Country.
// Note that you need to call Country.Unwrap() before calling this method if this Country
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Country) Update() *CountryUpdateOne {
	return NewCountryClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Country entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Country) Unwrap() *Country {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Country is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Country) String() string {
	var builder strings.Builder
	builder.WriteString("Country(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("name_en=")
	builder.WriteString(c.NameEn)
	builder.WriteString(", ")
	builder.WriteString("code=")
	builder.WriteString(c.Code)
	builder.WriteString(", ")
	builder.WriteString("display_sort=")
	builder.WriteString(fmt.Sprintf("%v", c.DisplaySort))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteByte(')')
	return builder.String()
}

// NamedRegions returns the Regions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Country) NamedRegions(name string) ([]*Region, error) {
	if c.Edges.namedRegions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedRegions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Country) appendNamedRegions(name string, edges ...*Region) {
	if c.Edges.namedRegions == nil {
		c.Edges.namedRegions = make(map[string][]*Region)
	}
	if len(edges) == 0 {
		c.Edges.namedRegions[name] = []*Region{}
	} else {
		c.Edges.namedRegions[name] = append(c.Edges.namedRegions[name], edges...)
	}
}

// Countries is a parsable slice of Country.
type Countries []*Country