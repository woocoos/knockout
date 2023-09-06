// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/knockout/ent/file"
	"github.com/woocoos/knockout/ent/filesource"
)

// File is the model entity for the File schema.
type File struct {
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
	// 文件名称
	Name string `json:"name,omitempty"`
	// 文件来源
	SourceID int `json:"source_id,omitempty"`
	// 租户ID
	TenantID int `json:"tenant_id,omitempty"`
	// 业务引用次数
	RefCount int `json:"ref_count,omitempty"`
	// 文件相对路径
	Path string `json:"path,omitempty"`
	// 文件大小，单位为B
	Size int `json:"size,omitempty"`
	// 媒体类型，如：image/png
	MineType string `json:"mine_type,omitempty"`
	// md5值
	Md5 string `json:"md5,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges        FileEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// 文件来源
	Source *FileSource `json:"source,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// SourceOrErr returns the Source value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) SourceOrErr() (*FileSource, error) {
	if e.loadedTypes[0] {
		if e.Source == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: filesource.Label}
		}
		return e.Source, nil
	}
	return nil, &NotLoadedError{edge: "source"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldID, file.FieldCreatedBy, file.FieldUpdatedBy, file.FieldSourceID, file.FieldTenantID, file.FieldRefCount, file.FieldSize:
			values[i] = new(sql.NullInt64)
		case file.FieldName, file.FieldPath, file.FieldMineType, file.FieldMd5:
			values[i] = new(sql.NullString)
		case file.FieldCreatedAt, file.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case file.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				f.CreatedBy = int(value.Int64)
			}
		case file.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case file.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				f.UpdatedBy = int(value.Int64)
			}
		case file.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case file.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case file.FieldSourceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source_id", values[i])
			} else if value.Valid {
				f.SourceID = int(value.Int64)
			}
		case file.FieldTenantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				f.TenantID = int(value.Int64)
			}
		case file.FieldRefCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ref_count", values[i])
			} else if value.Valid {
				f.RefCount = int(value.Int64)
			}
		case file.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				f.Path = value.String
			}
		case file.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				f.Size = int(value.Int64)
			}
		case file.FieldMineType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mine_type", values[i])
			} else if value.Valid {
				f.MineType = value.String
			}
		case file.FieldMd5:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field md5", values[i])
			} else if value.Valid {
				f.Md5 = value.String
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the File.
// This includes values selected through modifiers, order, etc.
func (f *File) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QuerySource queries the "source" edge of the File entity.
func (f *File) QuerySource() *FileSourceQuery {
	return NewFileClient(f.config).QuerySource(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return NewFileClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", f.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", f.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	builder.WriteString("source_id=")
	builder.WriteString(fmt.Sprintf("%v", f.SourceID))
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(fmt.Sprintf("%v", f.TenantID))
	builder.WriteString(", ")
	builder.WriteString("ref_count=")
	builder.WriteString(fmt.Sprintf("%v", f.RefCount))
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(f.Path)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", ")
	builder.WriteString("mine_type=")
	builder.WriteString(f.MineType)
	builder.WriteString(", ")
	builder.WriteString("md5=")
	builder.WriteString(f.Md5)
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File
