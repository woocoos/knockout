// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/file"
	"github.com/woocoos/knockout/ent/filesource"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (fc *FileCreate) SetCreatedBy(i int) *FileCreate {
	fc.mutation.SetCreatedBy(i)
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FileCreate) SetCreatedAt(t time.Time) *FileCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableCreatedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedBy sets the "updated_by" field.
func (fc *FileCreate) SetUpdatedBy(i int) *FileCreate {
	fc.mutation.SetUpdatedBy(i)
	return fc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fc *FileCreate) SetNillableUpdatedBy(i *int) *FileCreate {
	if i != nil {
		fc.SetUpdatedBy(*i)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FileCreate) SetUpdatedAt(t time.Time) *FileCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableUpdatedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetName sets the "name" field.
func (fc *FileCreate) SetName(s string) *FileCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetSourceID sets the "source_id" field.
func (fc *FileCreate) SetSourceID(i int) *FileCreate {
	fc.mutation.SetSourceID(i)
	return fc
}

// SetTenantID sets the "tenant_id" field.
func (fc *FileCreate) SetTenantID(i int) *FileCreate {
	fc.mutation.SetTenantID(i)
	return fc
}

// SetPath sets the "path" field.
func (fc *FileCreate) SetPath(s string) *FileCreate {
	fc.mutation.SetPath(s)
	return fc
}

// SetSize sets the "size" field.
func (fc *FileCreate) SetSize(i int) *FileCreate {
	fc.mutation.SetSize(i)
	return fc
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fc *FileCreate) SetNillableSize(i *int) *FileCreate {
	if i != nil {
		fc.SetSize(*i)
	}
	return fc
}

// SetMineType sets the "mine_type" field.
func (fc *FileCreate) SetMineType(s string) *FileCreate {
	fc.mutation.SetMineType(s)
	return fc
}

// SetNillableMineType sets the "mine_type" field if the given value is not nil.
func (fc *FileCreate) SetNillableMineType(s *string) *FileCreate {
	if s != nil {
		fc.SetMineType(*s)
	}
	return fc
}

// SetMd5 sets the "md5" field.
func (fc *FileCreate) SetMd5(s string) *FileCreate {
	fc.mutation.SetMd5(s)
	return fc
}

// SetNillableMd5 sets the "md5" field if the given value is not nil.
func (fc *FileCreate) SetNillableMd5(s *string) *FileCreate {
	if s != nil {
		fc.SetMd5(*s)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FileCreate) SetID(i int) *FileCreate {
	fc.mutation.SetID(i)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FileCreate) SetNillableID(i *int) *FileCreate {
	if i != nil {
		fc.SetID(*i)
	}
	return fc
}

// SetSource sets the "source" edge to the FileSource entity.
func (fc *FileCreate) SetSource(f *FileSource) *FileCreate {
	return fc.SetSourceID(f.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	if err := fc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FileCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FileCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		if file.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized file.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := file.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		if file.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized file.DefaultID (forgotten import ent/runtime?)")
		}
		v := file.DefaultID()
		fc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "File.created_by"`)}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "File.created_at"`)}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "File.name"`)}
	}
	if _, ok := fc.mutation.SourceID(); !ok {
		return &ValidationError{Name: "source_id", err: errors.New(`ent: missing required field "File.source_id"`)}
	}
	if _, ok := fc.mutation.TenantID(); !ok {
		return &ValidationError{Name: "tenant_id", err: errors.New(`ent: missing required field "File.tenant_id"`)}
	}
	if _, ok := fc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "File.path"`)}
	}
	if v, ok := fc.mutation.MineType(); ok {
		if err := file.MineTypeValidator(v); err != nil {
			return &ValidationError{Name: "mine_type", err: fmt.Errorf(`ent: validator failed for field "File.mine_type": %w`, err)}
		}
	}
	if v, ok := fc.mutation.Md5(); ok {
		if err := file.Md5Validator(v); err != nil {
			return &ValidationError{Name: "md5", err: fmt.Errorf(`ent: validator failed for field "File.md5": %w`, err)}
		}
	}
	if _, ok := fc.mutation.SourceID(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`ent: missing required edge "File.source"`)}
	}
	return nil
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FileCreate) createSpec() (*File, *sqlgraph.CreateSpec) {
	var (
		_node = &File{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(file.Table, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	)
	_spec.OnConflict = fc.conflict
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.CreatedBy(); ok {
		_spec.SetField(file.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(file.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedBy(); ok {
		_spec.SetField(file.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.TenantID(); ok {
		_spec.SetField(file.FieldTenantID, field.TypeInt, value)
		_node.TenantID = value
	}
	if value, ok := fc.mutation.Path(); ok {
		_spec.SetField(file.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := fc.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeInt, value)
		_node.Size = value
	}
	if value, ok := fc.mutation.MineType(); ok {
		_spec.SetField(file.FieldMineType, field.TypeString, value)
		_node.MineType = value
	}
	if value, ok := fc.mutation.Md5(); ok {
		_spec.SetField(file.FieldMd5, field.TypeString, value)
		_node.Md5 = value
	}
	if nodes := fc.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.SourceTable,
			Columns: []string{file.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(filesource.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SourceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.File.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (fc *FileCreate) OnConflict(opts ...sql.ConflictOption) *FileUpsertOne {
	fc.conflict = opts
	return &FileUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FileCreate) OnConflictColumns(columns ...string) *FileUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FileUpsertOne{
		create: fc,
	}
}

type (
	// FileUpsertOne is the builder for "upsert"-ing
	//  one File node.
	FileUpsertOne struct {
		create *FileCreate
	}

	// FileUpsert is the "OnConflict" setter.
	FileUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *FileUpsert) SetUpdatedBy(v int) *FileUpsert {
	u.Set(file.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *FileUpsert) UpdateUpdatedBy() *FileUpsert {
	u.SetExcluded(file.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *FileUpsert) AddUpdatedBy(v int) *FileUpsert {
	u.Add(file.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *FileUpsert) ClearUpdatedBy() *FileUpsert {
	u.SetNull(file.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FileUpsert) SetUpdatedAt(v time.Time) *FileUpsert {
	u.Set(file.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileUpsert) UpdateUpdatedAt() *FileUpsert {
	u.SetExcluded(file.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *FileUpsert) ClearUpdatedAt() *FileUpsert {
	u.SetNull(file.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *FileUpsert) SetName(v string) *FileUpsert {
	u.Set(file.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileUpsert) UpdateName() *FileUpsert {
	u.SetExcluded(file.FieldName)
	return u
}

// SetSourceID sets the "source_id" field.
func (u *FileUpsert) SetSourceID(v int) *FileUpsert {
	u.Set(file.FieldSourceID, v)
	return u
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *FileUpsert) UpdateSourceID() *FileUpsert {
	u.SetExcluded(file.FieldSourceID)
	return u
}

// SetTenantID sets the "tenant_id" field.
func (u *FileUpsert) SetTenantID(v int) *FileUpsert {
	u.Set(file.FieldTenantID, v)
	return u
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *FileUpsert) UpdateTenantID() *FileUpsert {
	u.SetExcluded(file.FieldTenantID)
	return u
}

// AddTenantID adds v to the "tenant_id" field.
func (u *FileUpsert) AddTenantID(v int) *FileUpsert {
	u.Add(file.FieldTenantID, v)
	return u
}

// SetPath sets the "path" field.
func (u *FileUpsert) SetPath(v string) *FileUpsert {
	u.Set(file.FieldPath, v)
	return u
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *FileUpsert) UpdatePath() *FileUpsert {
	u.SetExcluded(file.FieldPath)
	return u
}

// SetSize sets the "size" field.
func (u *FileUpsert) SetSize(v int) *FileUpsert {
	u.Set(file.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsert) UpdateSize() *FileUpsert {
	u.SetExcluded(file.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *FileUpsert) AddSize(v int) *FileUpsert {
	u.Add(file.FieldSize, v)
	return u
}

// ClearSize clears the value of the "size" field.
func (u *FileUpsert) ClearSize() *FileUpsert {
	u.SetNull(file.FieldSize)
	return u
}

// SetMineType sets the "mine_type" field.
func (u *FileUpsert) SetMineType(v string) *FileUpsert {
	u.Set(file.FieldMineType, v)
	return u
}

// UpdateMineType sets the "mine_type" field to the value that was provided on create.
func (u *FileUpsert) UpdateMineType() *FileUpsert {
	u.SetExcluded(file.FieldMineType)
	return u
}

// ClearMineType clears the value of the "mine_type" field.
func (u *FileUpsert) ClearMineType() *FileUpsert {
	u.SetNull(file.FieldMineType)
	return u
}

// SetMd5 sets the "md5" field.
func (u *FileUpsert) SetMd5(v string) *FileUpsert {
	u.Set(file.FieldMd5, v)
	return u
}

// UpdateMd5 sets the "md5" field to the value that was provided on create.
func (u *FileUpsert) UpdateMd5() *FileUpsert {
	u.SetExcluded(file.FieldMd5)
	return u
}

// ClearMd5 clears the value of the "md5" field.
func (u *FileUpsert) ClearMd5() *FileUpsert {
	u.SetNull(file.FieldMd5)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(file.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FileUpsertOne) UpdateNewValues() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(file.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(file.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(file.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.File.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FileUpsertOne) Ignore() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileUpsertOne) DoNothing() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileCreate.OnConflict
// documentation for more info.
func (u *FileUpsertOne) Update(set func(*FileUpsert)) *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *FileUpsertOne) SetUpdatedBy(v int) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *FileUpsertOne) AddUpdatedBy(v int) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateUpdatedBy() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *FileUpsertOne) ClearUpdatedBy() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FileUpsertOne) SetUpdatedAt(v time.Time) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateUpdatedAt() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *FileUpsertOne) ClearUpdatedAt() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *FileUpsertOne) SetName(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateName() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateName()
	})
}

// SetSourceID sets the "source_id" field.
func (u *FileUpsertOne) SetSourceID(v int) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetSourceID(v)
	})
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateSourceID() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSourceID()
	})
}

// SetTenantID sets the "tenant_id" field.
func (u *FileUpsertOne) SetTenantID(v int) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetTenantID(v)
	})
}

// AddTenantID adds v to the "tenant_id" field.
func (u *FileUpsertOne) AddTenantID(v int) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.AddTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateTenantID() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateTenantID()
	})
}

// SetPath sets the "path" field.
func (u *FileUpsertOne) SetPath(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *FileUpsertOne) UpdatePath() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdatePath()
	})
}

// SetSize sets the "size" field.
func (u *FileUpsertOne) SetSize(v int) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *FileUpsertOne) AddSize(v int) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateSize() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSize()
	})
}

// ClearSize clears the value of the "size" field.
func (u *FileUpsertOne) ClearSize() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.ClearSize()
	})
}

// SetMineType sets the "mine_type" field.
func (u *FileUpsertOne) SetMineType(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetMineType(v)
	})
}

// UpdateMineType sets the "mine_type" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateMineType() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateMineType()
	})
}

// ClearMineType clears the value of the "mine_type" field.
func (u *FileUpsertOne) ClearMineType() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.ClearMineType()
	})
}

// SetMd5 sets the "md5" field.
func (u *FileUpsertOne) SetMd5(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetMd5(v)
	})
}

// UpdateMd5 sets the "md5" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateMd5() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateMd5()
	})
}

// ClearMd5 clears the value of the "md5" field.
func (u *FileUpsertOne) ClearMd5() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.ClearMd5()
	})
}

// Exec executes the query.
func (u *FileUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FileUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FileUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	builders []*FileCreate
	conflict []sql.ConflictOption
}

// Save creates the File entities in the database.
func (fcb *FileCreateBulk) Save(ctx context.Context) ([]*File, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*File, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FileCreateBulk) SaveX(ctx context.Context) []*File {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FileCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FileCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.File.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (fcb *FileCreateBulk) OnConflict(opts ...sql.ConflictOption) *FileUpsertBulk {
	fcb.conflict = opts
	return &FileUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FileCreateBulk) OnConflictColumns(columns ...string) *FileUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FileUpsertBulk{
		create: fcb,
	}
}

// FileUpsertBulk is the builder for "upsert"-ing
// a bulk of File nodes.
type FileUpsertBulk struct {
	create *FileCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(file.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FileUpsertBulk) UpdateNewValues() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(file.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(file.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(file.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FileUpsertBulk) Ignore() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileUpsertBulk) DoNothing() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileCreateBulk.OnConflict
// documentation for more info.
func (u *FileUpsertBulk) Update(set func(*FileUpsert)) *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *FileUpsertBulk) SetUpdatedBy(v int) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *FileUpsertBulk) AddUpdatedBy(v int) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateUpdatedBy() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *FileUpsertBulk) ClearUpdatedBy() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FileUpsertBulk) SetUpdatedAt(v time.Time) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateUpdatedAt() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *FileUpsertBulk) ClearUpdatedAt() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *FileUpsertBulk) SetName(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateName() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateName()
	})
}

// SetSourceID sets the "source_id" field.
func (u *FileUpsertBulk) SetSourceID(v int) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetSourceID(v)
	})
}

// UpdateSourceID sets the "source_id" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateSourceID() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSourceID()
	})
}

// SetTenantID sets the "tenant_id" field.
func (u *FileUpsertBulk) SetTenantID(v int) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetTenantID(v)
	})
}

// AddTenantID adds v to the "tenant_id" field.
func (u *FileUpsertBulk) AddTenantID(v int) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.AddTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateTenantID() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateTenantID()
	})
}

// SetPath sets the "path" field.
func (u *FileUpsertBulk) SetPath(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdatePath() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdatePath()
	})
}

// SetSize sets the "size" field.
func (u *FileUpsertBulk) SetSize(v int) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *FileUpsertBulk) AddSize(v int) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateSize() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSize()
	})
}

// ClearSize clears the value of the "size" field.
func (u *FileUpsertBulk) ClearSize() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.ClearSize()
	})
}

// SetMineType sets the "mine_type" field.
func (u *FileUpsertBulk) SetMineType(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetMineType(v)
	})
}

// UpdateMineType sets the "mine_type" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateMineType() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateMineType()
	})
}

// ClearMineType clears the value of the "mine_type" field.
func (u *FileUpsertBulk) ClearMineType() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.ClearMineType()
	})
}

// SetMd5 sets the "md5" field.
func (u *FileUpsertBulk) SetMd5(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetMd5(v)
	})
}

// UpdateMd5 sets the "md5" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateMd5() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateMd5()
	})
}

// ClearMd5 clears the value of the "md5" field.
func (u *FileUpsertBulk) ClearMd5() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.ClearMd5()
	})
}

// Exec executes the query.
func (u *FileUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FileCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
