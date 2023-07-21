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
	"github.com/woocoos/knockout/ent/predicate"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where appends a list predicates to the FileUpdate builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetUpdatedBy sets the "updated_by" field.
func (fu *FileUpdate) SetUpdatedBy(i int) *FileUpdate {
	fu.mutation.ResetUpdatedBy()
	fu.mutation.SetUpdatedBy(i)
	return fu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fu *FileUpdate) SetNillableUpdatedBy(i *int) *FileUpdate {
	if i != nil {
		fu.SetUpdatedBy(*i)
	}
	return fu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (fu *FileUpdate) AddUpdatedBy(i int) *FileUpdate {
	fu.mutation.AddUpdatedBy(i)
	return fu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (fu *FileUpdate) ClearUpdatedBy() *FileUpdate {
	fu.mutation.ClearUpdatedBy()
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FileUpdate) SetUpdatedAt(t time.Time) *FileUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fu *FileUpdate) SetNillableUpdatedAt(t *time.Time) *FileUpdate {
	if t != nil {
		fu.SetUpdatedAt(*t)
	}
	return fu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fu *FileUpdate) ClearUpdatedAt() *FileUpdate {
	fu.mutation.ClearUpdatedAt()
	return fu
}

// SetName sets the "name" field.
func (fu *FileUpdate) SetName(s string) *FileUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetSourceID sets the "source_id" field.
func (fu *FileUpdate) SetSourceID(i int) *FileUpdate {
	fu.mutation.SetSourceID(i)
	return fu
}

// SetTenantID sets the "tenant_id" field.
func (fu *FileUpdate) SetTenantID(i int) *FileUpdate {
	fu.mutation.ResetTenantID()
	fu.mutation.SetTenantID(i)
	return fu
}

// AddTenantID adds i to the "tenant_id" field.
func (fu *FileUpdate) AddTenantID(i int) *FileUpdate {
	fu.mutation.AddTenantID(i)
	return fu
}

// SetPath sets the "path" field.
func (fu *FileUpdate) SetPath(s string) *FileUpdate {
	fu.mutation.SetPath(s)
	return fu
}

// SetSize sets the "size" field.
func (fu *FileUpdate) SetSize(i int) *FileUpdate {
	fu.mutation.ResetSize()
	fu.mutation.SetSize(i)
	return fu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fu *FileUpdate) SetNillableSize(i *int) *FileUpdate {
	if i != nil {
		fu.SetSize(*i)
	}
	return fu
}

// AddSize adds i to the "size" field.
func (fu *FileUpdate) AddSize(i int) *FileUpdate {
	fu.mutation.AddSize(i)
	return fu
}

// ClearSize clears the value of the "size" field.
func (fu *FileUpdate) ClearSize() *FileUpdate {
	fu.mutation.ClearSize()
	return fu
}

// SetMineType sets the "mine_type" field.
func (fu *FileUpdate) SetMineType(s string) *FileUpdate {
	fu.mutation.SetMineType(s)
	return fu
}

// SetNillableMineType sets the "mine_type" field if the given value is not nil.
func (fu *FileUpdate) SetNillableMineType(s *string) *FileUpdate {
	if s != nil {
		fu.SetMineType(*s)
	}
	return fu
}

// ClearMineType clears the value of the "mine_type" field.
func (fu *FileUpdate) ClearMineType() *FileUpdate {
	fu.mutation.ClearMineType()
	return fu
}

// SetMd5 sets the "md5" field.
func (fu *FileUpdate) SetMd5(s string) *FileUpdate {
	fu.mutation.SetMd5(s)
	return fu
}

// SetNillableMd5 sets the "md5" field if the given value is not nil.
func (fu *FileUpdate) SetNillableMd5(s *string) *FileUpdate {
	if s != nil {
		fu.SetMd5(*s)
	}
	return fu
}

// ClearMd5 clears the value of the "md5" field.
func (fu *FileUpdate) ClearMd5() *FileUpdate {
	fu.mutation.ClearMd5()
	return fu
}

// SetSource sets the "source" edge to the FileSource entity.
func (fu *FileUpdate) SetSource(f *FileSource) *FileUpdate {
	return fu.SetSourceID(f.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fu *FileUpdate) Mutation() *FileMutation {
	return fu.mutation
}

// ClearSource clears the "source" edge to the FileSource entity.
func (fu *FileUpdate) ClearSource() *FileUpdate {
	fu.mutation.ClearSource()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FileUpdate) check() error {
	if v, ok := fu.mutation.MineType(); ok {
		if err := file.MineTypeValidator(v); err != nil {
			return &ValidationError{Name: "mine_type", err: fmt.Errorf(`ent: validator failed for field "File.mine_type": %w`, err)}
		}
	}
	if v, ok := fu.mutation.Md5(); ok {
		if err := file.Md5Validator(v); err != nil {
			return &ValidationError{Name: "md5", err: fmt.Errorf(`ent: validator failed for field "File.md5": %w`, err)}
		}
	}
	if _, ok := fu.mutation.SourceID(); fu.mutation.SourceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "File.source"`)
	}
	return nil
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.UpdatedBy(); ok {
		_spec.SetField(file.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := fu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(file.FieldUpdatedBy, field.TypeInt, value)
	}
	if fu.mutation.UpdatedByCleared() {
		_spec.ClearField(file.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
	}
	if fu.mutation.UpdatedAtCleared() {
		_spec.ClearField(file.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
	}
	if value, ok := fu.mutation.TenantID(); ok {
		_spec.SetField(file.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := fu.mutation.AddedTenantID(); ok {
		_spec.AddField(file.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := fu.mutation.Path(); ok {
		_spec.SetField(file.FieldPath, field.TypeString, value)
	}
	if value, ok := fu.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeInt, value)
	}
	if value, ok := fu.mutation.AddedSize(); ok {
		_spec.AddField(file.FieldSize, field.TypeInt, value)
	}
	if fu.mutation.SizeCleared() {
		_spec.ClearField(file.FieldSize, field.TypeInt)
	}
	if value, ok := fu.mutation.MineType(); ok {
		_spec.SetField(file.FieldMineType, field.TypeString, value)
	}
	if fu.mutation.MineTypeCleared() {
		_spec.ClearField(file.FieldMineType, field.TypeString)
	}
	if value, ok := fu.mutation.Md5(); ok {
		_spec.SetField(file.FieldMd5, field.TypeString, value)
	}
	if fu.mutation.Md5Cleared() {
		_spec.ClearField(file.FieldMd5, field.TypeString)
	}
	if fu.mutation.SourceCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.SourceIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (fuo *FileUpdateOne) SetUpdatedBy(i int) *FileUpdateOne {
	fuo.mutation.ResetUpdatedBy()
	fuo.mutation.SetUpdatedBy(i)
	return fuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableUpdatedBy(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetUpdatedBy(*i)
	}
	return fuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (fuo *FileUpdateOne) AddUpdatedBy(i int) *FileUpdateOne {
	fuo.mutation.AddUpdatedBy(i)
	return fuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (fuo *FileUpdateOne) ClearUpdatedBy() *FileUpdateOne {
	fuo.mutation.ClearUpdatedBy()
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FileUpdateOne) SetUpdatedAt(t time.Time) *FileUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableUpdatedAt(t *time.Time) *FileUpdateOne {
	if t != nil {
		fuo.SetUpdatedAt(*t)
	}
	return fuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fuo *FileUpdateOne) ClearUpdatedAt() *FileUpdateOne {
	fuo.mutation.ClearUpdatedAt()
	return fuo
}

// SetName sets the "name" field.
func (fuo *FileUpdateOne) SetName(s string) *FileUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetSourceID sets the "source_id" field.
func (fuo *FileUpdateOne) SetSourceID(i int) *FileUpdateOne {
	fuo.mutation.SetSourceID(i)
	return fuo
}

// SetTenantID sets the "tenant_id" field.
func (fuo *FileUpdateOne) SetTenantID(i int) *FileUpdateOne {
	fuo.mutation.ResetTenantID()
	fuo.mutation.SetTenantID(i)
	return fuo
}

// AddTenantID adds i to the "tenant_id" field.
func (fuo *FileUpdateOne) AddTenantID(i int) *FileUpdateOne {
	fuo.mutation.AddTenantID(i)
	return fuo
}

// SetPath sets the "path" field.
func (fuo *FileUpdateOne) SetPath(s string) *FileUpdateOne {
	fuo.mutation.SetPath(s)
	return fuo
}

// SetSize sets the "size" field.
func (fuo *FileUpdateOne) SetSize(i int) *FileUpdateOne {
	fuo.mutation.ResetSize()
	fuo.mutation.SetSize(i)
	return fuo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableSize(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetSize(*i)
	}
	return fuo
}

// AddSize adds i to the "size" field.
func (fuo *FileUpdateOne) AddSize(i int) *FileUpdateOne {
	fuo.mutation.AddSize(i)
	return fuo
}

// ClearSize clears the value of the "size" field.
func (fuo *FileUpdateOne) ClearSize() *FileUpdateOne {
	fuo.mutation.ClearSize()
	return fuo
}

// SetMineType sets the "mine_type" field.
func (fuo *FileUpdateOne) SetMineType(s string) *FileUpdateOne {
	fuo.mutation.SetMineType(s)
	return fuo
}

// SetNillableMineType sets the "mine_type" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableMineType(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetMineType(*s)
	}
	return fuo
}

// ClearMineType clears the value of the "mine_type" field.
func (fuo *FileUpdateOne) ClearMineType() *FileUpdateOne {
	fuo.mutation.ClearMineType()
	return fuo
}

// SetMd5 sets the "md5" field.
func (fuo *FileUpdateOne) SetMd5(s string) *FileUpdateOne {
	fuo.mutation.SetMd5(s)
	return fuo
}

// SetNillableMd5 sets the "md5" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableMd5(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetMd5(*s)
	}
	return fuo
}

// ClearMd5 clears the value of the "md5" field.
func (fuo *FileUpdateOne) ClearMd5() *FileUpdateOne {
	fuo.mutation.ClearMd5()
	return fuo
}

// SetSource sets the "source" edge to the FileSource entity.
func (fuo *FileUpdateOne) SetSource(f *FileSource) *FileUpdateOne {
	return fuo.SetSourceID(f.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fuo *FileUpdateOne) Mutation() *FileMutation {
	return fuo.mutation
}

// ClearSource clears the "source" edge to the FileSource entity.
func (fuo *FileUpdateOne) ClearSource() *FileUpdateOne {
	fuo.mutation.ClearSource()
	return fuo
}

// Where appends a list predicates to the FileUpdate builder.
func (fuo *FileUpdateOne) Where(ps ...predicate.File) *FileUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated File entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FileUpdateOne) check() error {
	if v, ok := fuo.mutation.MineType(); ok {
		if err := file.MineTypeValidator(v); err != nil {
			return &ValidationError{Name: "mine_type", err: fmt.Errorf(`ent: validator failed for field "File.mine_type": %w`, err)}
		}
	}
	if v, ok := fuo.mutation.Md5(); ok {
		if err := file.Md5Validator(v); err != nil {
			return &ValidationError{Name: "md5", err: fmt.Errorf(`ent: validator failed for field "File.md5": %w`, err)}
		}
	}
	if _, ok := fuo.mutation.SourceID(); fuo.mutation.SourceCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "File.source"`)
	}
	return nil
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (_node *File, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "File.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for _, f := range fields {
			if !file.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != file.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.UpdatedBy(); ok {
		_spec.SetField(file.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(file.FieldUpdatedBy, field.TypeInt, value)
	}
	if fuo.mutation.UpdatedByCleared() {
		_spec.ClearField(file.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
	}
	if fuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(file.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.TenantID(); ok {
		_spec.SetField(file.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.AddedTenantID(); ok {
		_spec.AddField(file.FieldTenantID, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.Path(); ok {
		_spec.SetField(file.FieldPath, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.AddedSize(); ok {
		_spec.AddField(file.FieldSize, field.TypeInt, value)
	}
	if fuo.mutation.SizeCleared() {
		_spec.ClearField(file.FieldSize, field.TypeInt)
	}
	if value, ok := fuo.mutation.MineType(); ok {
		_spec.SetField(file.FieldMineType, field.TypeString, value)
	}
	if fuo.mutation.MineTypeCleared() {
		_spec.ClearField(file.FieldMineType, field.TypeString)
	}
	if value, ok := fuo.mutation.Md5(); ok {
		_spec.SetField(file.FieldMd5, field.TypeString, value)
	}
	if fuo.mutation.Md5Cleared() {
		_spec.ClearField(file.FieldMd5, field.TypeString)
	}
	if fuo.mutation.SourceCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.SourceIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &File{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}