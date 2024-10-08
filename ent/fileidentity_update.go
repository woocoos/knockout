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
	"github.com/woocoos/knockout/ent/fileidentity"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/predicate"
)

// FileIdentityUpdate is the builder for updating FileIdentity entities.
type FileIdentityUpdate struct {
	config
	hooks    []Hook
	mutation *FileIdentityMutation
}

// Where appends a list predicates to the FileIdentityUpdate builder.
func (fiu *FileIdentityUpdate) Where(ps ...predicate.FileIdentity) *FileIdentityUpdate {
	fiu.mutation.Where(ps...)
	return fiu
}

// SetUpdatedBy sets the "updated_by" field.
func (fiu *FileIdentityUpdate) SetUpdatedBy(i int) *FileIdentityUpdate {
	fiu.mutation.ResetUpdatedBy()
	fiu.mutation.SetUpdatedBy(i)
	return fiu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fiu *FileIdentityUpdate) SetNillableUpdatedBy(i *int) *FileIdentityUpdate {
	if i != nil {
		fiu.SetUpdatedBy(*i)
	}
	return fiu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (fiu *FileIdentityUpdate) AddUpdatedBy(i int) *FileIdentityUpdate {
	fiu.mutation.AddUpdatedBy(i)
	return fiu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (fiu *FileIdentityUpdate) ClearUpdatedBy() *FileIdentityUpdate {
	fiu.mutation.ClearUpdatedBy()
	return fiu
}

// SetUpdatedAt sets the "updated_at" field.
func (fiu *FileIdentityUpdate) SetUpdatedAt(t time.Time) *FileIdentityUpdate {
	fiu.mutation.SetUpdatedAt(t)
	return fiu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fiu *FileIdentityUpdate) SetNillableUpdatedAt(t *time.Time) *FileIdentityUpdate {
	if t != nil {
		fiu.SetUpdatedAt(*t)
	}
	return fiu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fiu *FileIdentityUpdate) ClearUpdatedAt() *FileIdentityUpdate {
	fiu.mutation.ClearUpdatedAt()
	return fiu
}

// SetAccessKeyID sets the "access_key_id" field.
func (fiu *FileIdentityUpdate) SetAccessKeyID(s string) *FileIdentityUpdate {
	fiu.mutation.SetAccessKeyID(s)
	return fiu
}

// SetNillableAccessKeyID sets the "access_key_id" field if the given value is not nil.
func (fiu *FileIdentityUpdate) SetNillableAccessKeyID(s *string) *FileIdentityUpdate {
	if s != nil {
		fiu.SetAccessKeyID(*s)
	}
	return fiu
}

// SetAccessKeySecret sets the "access_key_secret" field.
func (fiu *FileIdentityUpdate) SetAccessKeySecret(s string) *FileIdentityUpdate {
	fiu.mutation.SetAccessKeySecret(s)
	return fiu
}

// SetNillableAccessKeySecret sets the "access_key_secret" field if the given value is not nil.
func (fiu *FileIdentityUpdate) SetNillableAccessKeySecret(s *string) *FileIdentityUpdate {
	if s != nil {
		fiu.SetAccessKeySecret(*s)
	}
	return fiu
}

// SetFileSourceID sets the "file_source_id" field.
func (fiu *FileIdentityUpdate) SetFileSourceID(i int) *FileIdentityUpdate {
	fiu.mutation.SetFileSourceID(i)
	return fiu
}

// SetNillableFileSourceID sets the "file_source_id" field if the given value is not nil.
func (fiu *FileIdentityUpdate) SetNillableFileSourceID(i *int) *FileIdentityUpdate {
	if i != nil {
		fiu.SetFileSourceID(*i)
	}
	return fiu
}

// SetRoleArn sets the "role_arn" field.
func (fiu *FileIdentityUpdate) SetRoleArn(s string) *FileIdentityUpdate {
	fiu.mutation.SetRoleArn(s)
	return fiu
}

// SetNillableRoleArn sets the "role_arn" field if the given value is not nil.
func (fiu *FileIdentityUpdate) SetNillableRoleArn(s *string) *FileIdentityUpdate {
	if s != nil {
		fiu.SetRoleArn(*s)
	}
	return fiu
}

// SetPolicy sets the "policy" field.
func (fiu *FileIdentityUpdate) SetPolicy(s string) *FileIdentityUpdate {
	fiu.mutation.SetPolicy(s)
	return fiu
}

// SetNillablePolicy sets the "policy" field if the given value is not nil.
func (fiu *FileIdentityUpdate) SetNillablePolicy(s *string) *FileIdentityUpdate {
	if s != nil {
		fiu.SetPolicy(*s)
	}
	return fiu
}

// ClearPolicy clears the value of the "policy" field.
func (fiu *FileIdentityUpdate) ClearPolicy() *FileIdentityUpdate {
	fiu.mutation.ClearPolicy()
	return fiu
}

// SetDurationSeconds sets the "duration_seconds" field.
func (fiu *FileIdentityUpdate) SetDurationSeconds(i int) *FileIdentityUpdate {
	fiu.mutation.ResetDurationSeconds()
	fiu.mutation.SetDurationSeconds(i)
	return fiu
}

// SetNillableDurationSeconds sets the "duration_seconds" field if the given value is not nil.
func (fiu *FileIdentityUpdate) SetNillableDurationSeconds(i *int) *FileIdentityUpdate {
	if i != nil {
		fiu.SetDurationSeconds(*i)
	}
	return fiu
}

// AddDurationSeconds adds i to the "duration_seconds" field.
func (fiu *FileIdentityUpdate) AddDurationSeconds(i int) *FileIdentityUpdate {
	fiu.mutation.AddDurationSeconds(i)
	return fiu
}

// ClearDurationSeconds clears the value of the "duration_seconds" field.
func (fiu *FileIdentityUpdate) ClearDurationSeconds() *FileIdentityUpdate {
	fiu.mutation.ClearDurationSeconds()
	return fiu
}

// SetIsDefault sets the "is_default" field.
func (fiu *FileIdentityUpdate) SetIsDefault(b bool) *FileIdentityUpdate {
	fiu.mutation.SetIsDefault(b)
	return fiu
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (fiu *FileIdentityUpdate) SetNillableIsDefault(b *bool) *FileIdentityUpdate {
	if b != nil {
		fiu.SetIsDefault(*b)
	}
	return fiu
}

// SetComments sets the "comments" field.
func (fiu *FileIdentityUpdate) SetComments(s string) *FileIdentityUpdate {
	fiu.mutation.SetComments(s)
	return fiu
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (fiu *FileIdentityUpdate) SetNillableComments(s *string) *FileIdentityUpdate {
	if s != nil {
		fiu.SetComments(*s)
	}
	return fiu
}

// ClearComments clears the value of the "comments" field.
func (fiu *FileIdentityUpdate) ClearComments() *FileIdentityUpdate {
	fiu.mutation.ClearComments()
	return fiu
}

// SetSourceID sets the "source" edge to the FileSource entity by ID.
func (fiu *FileIdentityUpdate) SetSourceID(id int) *FileIdentityUpdate {
	fiu.mutation.SetSourceID(id)
	return fiu
}

// SetSource sets the "source" edge to the FileSource entity.
func (fiu *FileIdentityUpdate) SetSource(f *FileSource) *FileIdentityUpdate {
	return fiu.SetSourceID(f.ID)
}

// Mutation returns the FileIdentityMutation object of the builder.
func (fiu *FileIdentityUpdate) Mutation() *FileIdentityMutation {
	return fiu.mutation
}

// ClearSource clears the "source" edge to the FileSource entity.
func (fiu *FileIdentityUpdate) ClearSource() *FileIdentityUpdate {
	fiu.mutation.ClearSource()
	return fiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fiu *FileIdentityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fiu.sqlSave, fiu.mutation, fiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fiu *FileIdentityUpdate) SaveX(ctx context.Context) int {
	affected, err := fiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fiu *FileIdentityUpdate) Exec(ctx context.Context) error {
	_, err := fiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fiu *FileIdentityUpdate) ExecX(ctx context.Context) {
	if err := fiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fiu *FileIdentityUpdate) check() error {
	if v, ok := fiu.mutation.AccessKeyID(); ok {
		if err := fileidentity.AccessKeyIDValidator(v); err != nil {
			return &ValidationError{Name: "access_key_id", err: fmt.Errorf(`ent: validator failed for field "FileIdentity.access_key_id": %w`, err)}
		}
	}
	if v, ok := fiu.mutation.AccessKeySecret(); ok {
		if err := fileidentity.AccessKeySecretValidator(v); err != nil {
			return &ValidationError{Name: "access_key_secret", err: fmt.Errorf(`ent: validator failed for field "FileIdentity.access_key_secret": %w`, err)}
		}
	}
	if v, ok := fiu.mutation.RoleArn(); ok {
		if err := fileidentity.RoleArnValidator(v); err != nil {
			return &ValidationError{Name: "role_arn", err: fmt.Errorf(`ent: validator failed for field "FileIdentity.role_arn": %w`, err)}
		}
	}
	if fiu.mutation.SourceCleared() && len(fiu.mutation.SourceIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "FileIdentity.source"`)
	}
	if fiu.mutation.OrgCleared() && len(fiu.mutation.OrgIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "FileIdentity.org"`)
	}
	return nil
}

func (fiu *FileIdentityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(fileidentity.Table, fileidentity.Columns, sqlgraph.NewFieldSpec(fileidentity.FieldID, field.TypeInt))
	if ps := fiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fiu.mutation.UpdatedBy(); ok {
		_spec.SetField(fileidentity.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := fiu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(fileidentity.FieldUpdatedBy, field.TypeInt, value)
	}
	if fiu.mutation.UpdatedByCleared() {
		_spec.ClearField(fileidentity.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := fiu.mutation.UpdatedAt(); ok {
		_spec.SetField(fileidentity.FieldUpdatedAt, field.TypeTime, value)
	}
	if fiu.mutation.UpdatedAtCleared() {
		_spec.ClearField(fileidentity.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := fiu.mutation.AccessKeyID(); ok {
		_spec.SetField(fileidentity.FieldAccessKeyID, field.TypeString, value)
	}
	if value, ok := fiu.mutation.AccessKeySecret(); ok {
		_spec.SetField(fileidentity.FieldAccessKeySecret, field.TypeString, value)
	}
	if value, ok := fiu.mutation.RoleArn(); ok {
		_spec.SetField(fileidentity.FieldRoleArn, field.TypeString, value)
	}
	if value, ok := fiu.mutation.Policy(); ok {
		_spec.SetField(fileidentity.FieldPolicy, field.TypeString, value)
	}
	if fiu.mutation.PolicyCleared() {
		_spec.ClearField(fileidentity.FieldPolicy, field.TypeString)
	}
	if value, ok := fiu.mutation.DurationSeconds(); ok {
		_spec.SetField(fileidentity.FieldDurationSeconds, field.TypeInt, value)
	}
	if value, ok := fiu.mutation.AddedDurationSeconds(); ok {
		_spec.AddField(fileidentity.FieldDurationSeconds, field.TypeInt, value)
	}
	if fiu.mutation.DurationSecondsCleared() {
		_spec.ClearField(fileidentity.FieldDurationSeconds, field.TypeInt)
	}
	if value, ok := fiu.mutation.IsDefault(); ok {
		_spec.SetField(fileidentity.FieldIsDefault, field.TypeBool, value)
	}
	if value, ok := fiu.mutation.Comments(); ok {
		_spec.SetField(fileidentity.FieldComments, field.TypeString, value)
	}
	if fiu.mutation.CommentsCleared() {
		_spec.ClearField(fileidentity.FieldComments, field.TypeString)
	}
	if fiu.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fileidentity.SourceTable,
			Columns: []string{fileidentity.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(filesource.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiu.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fileidentity.SourceTable,
			Columns: []string{fileidentity.SourceColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fileidentity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fiu.mutation.done = true
	return n, nil
}

// FileIdentityUpdateOne is the builder for updating a single FileIdentity entity.
type FileIdentityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileIdentityMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (fiuo *FileIdentityUpdateOne) SetUpdatedBy(i int) *FileIdentityUpdateOne {
	fiuo.mutation.ResetUpdatedBy()
	fiuo.mutation.SetUpdatedBy(i)
	return fiuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fiuo *FileIdentityUpdateOne) SetNillableUpdatedBy(i *int) *FileIdentityUpdateOne {
	if i != nil {
		fiuo.SetUpdatedBy(*i)
	}
	return fiuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (fiuo *FileIdentityUpdateOne) AddUpdatedBy(i int) *FileIdentityUpdateOne {
	fiuo.mutation.AddUpdatedBy(i)
	return fiuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (fiuo *FileIdentityUpdateOne) ClearUpdatedBy() *FileIdentityUpdateOne {
	fiuo.mutation.ClearUpdatedBy()
	return fiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fiuo *FileIdentityUpdateOne) SetUpdatedAt(t time.Time) *FileIdentityUpdateOne {
	fiuo.mutation.SetUpdatedAt(t)
	return fiuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fiuo *FileIdentityUpdateOne) SetNillableUpdatedAt(t *time.Time) *FileIdentityUpdateOne {
	if t != nil {
		fiuo.SetUpdatedAt(*t)
	}
	return fiuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fiuo *FileIdentityUpdateOne) ClearUpdatedAt() *FileIdentityUpdateOne {
	fiuo.mutation.ClearUpdatedAt()
	return fiuo
}

// SetAccessKeyID sets the "access_key_id" field.
func (fiuo *FileIdentityUpdateOne) SetAccessKeyID(s string) *FileIdentityUpdateOne {
	fiuo.mutation.SetAccessKeyID(s)
	return fiuo
}

// SetNillableAccessKeyID sets the "access_key_id" field if the given value is not nil.
func (fiuo *FileIdentityUpdateOne) SetNillableAccessKeyID(s *string) *FileIdentityUpdateOne {
	if s != nil {
		fiuo.SetAccessKeyID(*s)
	}
	return fiuo
}

// SetAccessKeySecret sets the "access_key_secret" field.
func (fiuo *FileIdentityUpdateOne) SetAccessKeySecret(s string) *FileIdentityUpdateOne {
	fiuo.mutation.SetAccessKeySecret(s)
	return fiuo
}

// SetNillableAccessKeySecret sets the "access_key_secret" field if the given value is not nil.
func (fiuo *FileIdentityUpdateOne) SetNillableAccessKeySecret(s *string) *FileIdentityUpdateOne {
	if s != nil {
		fiuo.SetAccessKeySecret(*s)
	}
	return fiuo
}

// SetFileSourceID sets the "file_source_id" field.
func (fiuo *FileIdentityUpdateOne) SetFileSourceID(i int) *FileIdentityUpdateOne {
	fiuo.mutation.SetFileSourceID(i)
	return fiuo
}

// SetNillableFileSourceID sets the "file_source_id" field if the given value is not nil.
func (fiuo *FileIdentityUpdateOne) SetNillableFileSourceID(i *int) *FileIdentityUpdateOne {
	if i != nil {
		fiuo.SetFileSourceID(*i)
	}
	return fiuo
}

// SetRoleArn sets the "role_arn" field.
func (fiuo *FileIdentityUpdateOne) SetRoleArn(s string) *FileIdentityUpdateOne {
	fiuo.mutation.SetRoleArn(s)
	return fiuo
}

// SetNillableRoleArn sets the "role_arn" field if the given value is not nil.
func (fiuo *FileIdentityUpdateOne) SetNillableRoleArn(s *string) *FileIdentityUpdateOne {
	if s != nil {
		fiuo.SetRoleArn(*s)
	}
	return fiuo
}

// SetPolicy sets the "policy" field.
func (fiuo *FileIdentityUpdateOne) SetPolicy(s string) *FileIdentityUpdateOne {
	fiuo.mutation.SetPolicy(s)
	return fiuo
}

// SetNillablePolicy sets the "policy" field if the given value is not nil.
func (fiuo *FileIdentityUpdateOne) SetNillablePolicy(s *string) *FileIdentityUpdateOne {
	if s != nil {
		fiuo.SetPolicy(*s)
	}
	return fiuo
}

// ClearPolicy clears the value of the "policy" field.
func (fiuo *FileIdentityUpdateOne) ClearPolicy() *FileIdentityUpdateOne {
	fiuo.mutation.ClearPolicy()
	return fiuo
}

// SetDurationSeconds sets the "duration_seconds" field.
func (fiuo *FileIdentityUpdateOne) SetDurationSeconds(i int) *FileIdentityUpdateOne {
	fiuo.mutation.ResetDurationSeconds()
	fiuo.mutation.SetDurationSeconds(i)
	return fiuo
}

// SetNillableDurationSeconds sets the "duration_seconds" field if the given value is not nil.
func (fiuo *FileIdentityUpdateOne) SetNillableDurationSeconds(i *int) *FileIdentityUpdateOne {
	if i != nil {
		fiuo.SetDurationSeconds(*i)
	}
	return fiuo
}

// AddDurationSeconds adds i to the "duration_seconds" field.
func (fiuo *FileIdentityUpdateOne) AddDurationSeconds(i int) *FileIdentityUpdateOne {
	fiuo.mutation.AddDurationSeconds(i)
	return fiuo
}

// ClearDurationSeconds clears the value of the "duration_seconds" field.
func (fiuo *FileIdentityUpdateOne) ClearDurationSeconds() *FileIdentityUpdateOne {
	fiuo.mutation.ClearDurationSeconds()
	return fiuo
}

// SetIsDefault sets the "is_default" field.
func (fiuo *FileIdentityUpdateOne) SetIsDefault(b bool) *FileIdentityUpdateOne {
	fiuo.mutation.SetIsDefault(b)
	return fiuo
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (fiuo *FileIdentityUpdateOne) SetNillableIsDefault(b *bool) *FileIdentityUpdateOne {
	if b != nil {
		fiuo.SetIsDefault(*b)
	}
	return fiuo
}

// SetComments sets the "comments" field.
func (fiuo *FileIdentityUpdateOne) SetComments(s string) *FileIdentityUpdateOne {
	fiuo.mutation.SetComments(s)
	return fiuo
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (fiuo *FileIdentityUpdateOne) SetNillableComments(s *string) *FileIdentityUpdateOne {
	if s != nil {
		fiuo.SetComments(*s)
	}
	return fiuo
}

// ClearComments clears the value of the "comments" field.
func (fiuo *FileIdentityUpdateOne) ClearComments() *FileIdentityUpdateOne {
	fiuo.mutation.ClearComments()
	return fiuo
}

// SetSourceID sets the "source" edge to the FileSource entity by ID.
func (fiuo *FileIdentityUpdateOne) SetSourceID(id int) *FileIdentityUpdateOne {
	fiuo.mutation.SetSourceID(id)
	return fiuo
}

// SetSource sets the "source" edge to the FileSource entity.
func (fiuo *FileIdentityUpdateOne) SetSource(f *FileSource) *FileIdentityUpdateOne {
	return fiuo.SetSourceID(f.ID)
}

// Mutation returns the FileIdentityMutation object of the builder.
func (fiuo *FileIdentityUpdateOne) Mutation() *FileIdentityMutation {
	return fiuo.mutation
}

// ClearSource clears the "source" edge to the FileSource entity.
func (fiuo *FileIdentityUpdateOne) ClearSource() *FileIdentityUpdateOne {
	fiuo.mutation.ClearSource()
	return fiuo
}

// Where appends a list predicates to the FileIdentityUpdate builder.
func (fiuo *FileIdentityUpdateOne) Where(ps ...predicate.FileIdentity) *FileIdentityUpdateOne {
	fiuo.mutation.Where(ps...)
	return fiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fiuo *FileIdentityUpdateOne) Select(field string, fields ...string) *FileIdentityUpdateOne {
	fiuo.fields = append([]string{field}, fields...)
	return fiuo
}

// Save executes the query and returns the updated FileIdentity entity.
func (fiuo *FileIdentityUpdateOne) Save(ctx context.Context) (*FileIdentity, error) {
	return withHooks(ctx, fiuo.sqlSave, fiuo.mutation, fiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fiuo *FileIdentityUpdateOne) SaveX(ctx context.Context) *FileIdentity {
	node, err := fiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fiuo *FileIdentityUpdateOne) Exec(ctx context.Context) error {
	_, err := fiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fiuo *FileIdentityUpdateOne) ExecX(ctx context.Context) {
	if err := fiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fiuo *FileIdentityUpdateOne) check() error {
	if v, ok := fiuo.mutation.AccessKeyID(); ok {
		if err := fileidentity.AccessKeyIDValidator(v); err != nil {
			return &ValidationError{Name: "access_key_id", err: fmt.Errorf(`ent: validator failed for field "FileIdentity.access_key_id": %w`, err)}
		}
	}
	if v, ok := fiuo.mutation.AccessKeySecret(); ok {
		if err := fileidentity.AccessKeySecretValidator(v); err != nil {
			return &ValidationError{Name: "access_key_secret", err: fmt.Errorf(`ent: validator failed for field "FileIdentity.access_key_secret": %w`, err)}
		}
	}
	if v, ok := fiuo.mutation.RoleArn(); ok {
		if err := fileidentity.RoleArnValidator(v); err != nil {
			return &ValidationError{Name: "role_arn", err: fmt.Errorf(`ent: validator failed for field "FileIdentity.role_arn": %w`, err)}
		}
	}
	if fiuo.mutation.SourceCleared() && len(fiuo.mutation.SourceIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "FileIdentity.source"`)
	}
	if fiuo.mutation.OrgCleared() && len(fiuo.mutation.OrgIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "FileIdentity.org"`)
	}
	return nil
}

func (fiuo *FileIdentityUpdateOne) sqlSave(ctx context.Context) (_node *FileIdentity, err error) {
	if err := fiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(fileidentity.Table, fileidentity.Columns, sqlgraph.NewFieldSpec(fileidentity.FieldID, field.TypeInt))
	id, ok := fiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FileIdentity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fileidentity.FieldID)
		for _, f := range fields {
			if !fileidentity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fileidentity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fiuo.mutation.UpdatedBy(); ok {
		_spec.SetField(fileidentity.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := fiuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(fileidentity.FieldUpdatedBy, field.TypeInt, value)
	}
	if fiuo.mutation.UpdatedByCleared() {
		_spec.ClearField(fileidentity.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := fiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(fileidentity.FieldUpdatedAt, field.TypeTime, value)
	}
	if fiuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(fileidentity.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := fiuo.mutation.AccessKeyID(); ok {
		_spec.SetField(fileidentity.FieldAccessKeyID, field.TypeString, value)
	}
	if value, ok := fiuo.mutation.AccessKeySecret(); ok {
		_spec.SetField(fileidentity.FieldAccessKeySecret, field.TypeString, value)
	}
	if value, ok := fiuo.mutation.RoleArn(); ok {
		_spec.SetField(fileidentity.FieldRoleArn, field.TypeString, value)
	}
	if value, ok := fiuo.mutation.Policy(); ok {
		_spec.SetField(fileidentity.FieldPolicy, field.TypeString, value)
	}
	if fiuo.mutation.PolicyCleared() {
		_spec.ClearField(fileidentity.FieldPolicy, field.TypeString)
	}
	if value, ok := fiuo.mutation.DurationSeconds(); ok {
		_spec.SetField(fileidentity.FieldDurationSeconds, field.TypeInt, value)
	}
	if value, ok := fiuo.mutation.AddedDurationSeconds(); ok {
		_spec.AddField(fileidentity.FieldDurationSeconds, field.TypeInt, value)
	}
	if fiuo.mutation.DurationSecondsCleared() {
		_spec.ClearField(fileidentity.FieldDurationSeconds, field.TypeInt)
	}
	if value, ok := fiuo.mutation.IsDefault(); ok {
		_spec.SetField(fileidentity.FieldIsDefault, field.TypeBool, value)
	}
	if value, ok := fiuo.mutation.Comments(); ok {
		_spec.SetField(fileidentity.FieldComments, field.TypeString, value)
	}
	if fiuo.mutation.CommentsCleared() {
		_spec.ClearField(fileidentity.FieldComments, field.TypeString)
	}
	if fiuo.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fileidentity.SourceTable,
			Columns: []string{fileidentity.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(filesource.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fiuo.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   fileidentity.SourceTable,
			Columns: []string{fileidentity.SourceColumn},
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
	_node = &FileIdentity{config: fiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fileidentity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fiuo.mutation.done = true
	return _node, nil
}
