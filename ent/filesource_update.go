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
	"github.com/woocoos/knockout/ent/fileidentity"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/predicate"
)

// FileSourceUpdate is the builder for updating FileSource entities.
type FileSourceUpdate struct {
	config
	hooks    []Hook
	mutation *FileSourceMutation
}

// Where appends a list predicates to the FileSourceUpdate builder.
func (fsu *FileSourceUpdate) Where(ps ...predicate.FileSource) *FileSourceUpdate {
	fsu.mutation.Where(ps...)
	return fsu
}

// SetUpdatedBy sets the "updated_by" field.
func (fsu *FileSourceUpdate) SetUpdatedBy(i int) *FileSourceUpdate {
	fsu.mutation.ResetUpdatedBy()
	fsu.mutation.SetUpdatedBy(i)
	return fsu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fsu *FileSourceUpdate) SetNillableUpdatedBy(i *int) *FileSourceUpdate {
	if i != nil {
		fsu.SetUpdatedBy(*i)
	}
	return fsu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (fsu *FileSourceUpdate) AddUpdatedBy(i int) *FileSourceUpdate {
	fsu.mutation.AddUpdatedBy(i)
	return fsu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (fsu *FileSourceUpdate) ClearUpdatedBy() *FileSourceUpdate {
	fsu.mutation.ClearUpdatedBy()
	return fsu
}

// SetUpdatedAt sets the "updated_at" field.
func (fsu *FileSourceUpdate) SetUpdatedAt(t time.Time) *FileSourceUpdate {
	fsu.mutation.SetUpdatedAt(t)
	return fsu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fsu *FileSourceUpdate) SetNillableUpdatedAt(t *time.Time) *FileSourceUpdate {
	if t != nil {
		fsu.SetUpdatedAt(*t)
	}
	return fsu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fsu *FileSourceUpdate) ClearUpdatedAt() *FileSourceUpdate {
	fsu.mutation.ClearUpdatedAt()
	return fsu
}

// SetKind sets the "kind" field.
func (fsu *FileSourceUpdate) SetKind(f filesource.Kind) *FileSourceUpdate {
	fsu.mutation.SetKind(f)
	return fsu
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (fsu *FileSourceUpdate) SetNillableKind(f *filesource.Kind) *FileSourceUpdate {
	if f != nil {
		fsu.SetKind(*f)
	}
	return fsu
}

// SetComments sets the "comments" field.
func (fsu *FileSourceUpdate) SetComments(s string) *FileSourceUpdate {
	fsu.mutation.SetComments(s)
	return fsu
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (fsu *FileSourceUpdate) SetNillableComments(s *string) *FileSourceUpdate {
	if s != nil {
		fsu.SetComments(*s)
	}
	return fsu
}

// ClearComments clears the value of the "comments" field.
func (fsu *FileSourceUpdate) ClearComments() *FileSourceUpdate {
	fsu.mutation.ClearComments()
	return fsu
}

// SetEndpoint sets the "endpoint" field.
func (fsu *FileSourceUpdate) SetEndpoint(s string) *FileSourceUpdate {
	fsu.mutation.SetEndpoint(s)
	return fsu
}

// SetNillableEndpoint sets the "endpoint" field if the given value is not nil.
func (fsu *FileSourceUpdate) SetNillableEndpoint(s *string) *FileSourceUpdate {
	if s != nil {
		fsu.SetEndpoint(*s)
	}
	return fsu
}

// SetEndpointImmutable sets the "endpoint_immutable" field.
func (fsu *FileSourceUpdate) SetEndpointImmutable(b bool) *FileSourceUpdate {
	fsu.mutation.SetEndpointImmutable(b)
	return fsu
}

// SetNillableEndpointImmutable sets the "endpoint_immutable" field if the given value is not nil.
func (fsu *FileSourceUpdate) SetNillableEndpointImmutable(b *bool) *FileSourceUpdate {
	if b != nil {
		fsu.SetEndpointImmutable(*b)
	}
	return fsu
}

// SetStsEndpoint sets the "sts_endpoint" field.
func (fsu *FileSourceUpdate) SetStsEndpoint(s string) *FileSourceUpdate {
	fsu.mutation.SetStsEndpoint(s)
	return fsu
}

// SetNillableStsEndpoint sets the "sts_endpoint" field if the given value is not nil.
func (fsu *FileSourceUpdate) SetNillableStsEndpoint(s *string) *FileSourceUpdate {
	if s != nil {
		fsu.SetStsEndpoint(*s)
	}
	return fsu
}

// SetRegion sets the "region" field.
func (fsu *FileSourceUpdate) SetRegion(s string) *FileSourceUpdate {
	fsu.mutation.SetRegion(s)
	return fsu
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (fsu *FileSourceUpdate) SetNillableRegion(s *string) *FileSourceUpdate {
	if s != nil {
		fsu.SetRegion(*s)
	}
	return fsu
}

// SetBucket sets the "bucket" field.
func (fsu *FileSourceUpdate) SetBucket(s string) *FileSourceUpdate {
	fsu.mutation.SetBucket(s)
	return fsu
}

// SetNillableBucket sets the "bucket" field if the given value is not nil.
func (fsu *FileSourceUpdate) SetNillableBucket(s *string) *FileSourceUpdate {
	if s != nil {
		fsu.SetBucket(*s)
	}
	return fsu
}

// SetBucketURL sets the "bucket_url" field.
func (fsu *FileSourceUpdate) SetBucketURL(s string) *FileSourceUpdate {
	fsu.mutation.SetBucketURL(s)
	return fsu
}

// SetNillableBucketURL sets the "bucket_url" field if the given value is not nil.
func (fsu *FileSourceUpdate) SetNillableBucketURL(s *string) *FileSourceUpdate {
	if s != nil {
		fsu.SetBucketURL(*s)
	}
	return fsu
}

// AddIdentityIDs adds the "identities" edge to the FileIdentity entity by IDs.
func (fsu *FileSourceUpdate) AddIdentityIDs(ids ...int) *FileSourceUpdate {
	fsu.mutation.AddIdentityIDs(ids...)
	return fsu
}

// AddIdentities adds the "identities" edges to the FileIdentity entity.
func (fsu *FileSourceUpdate) AddIdentities(f ...*FileIdentity) *FileSourceUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsu.AddIdentityIDs(ids...)
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (fsu *FileSourceUpdate) AddFileIDs(ids ...int) *FileSourceUpdate {
	fsu.mutation.AddFileIDs(ids...)
	return fsu
}

// AddFiles adds the "files" edges to the File entity.
func (fsu *FileSourceUpdate) AddFiles(f ...*File) *FileSourceUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsu.AddFileIDs(ids...)
}

// Mutation returns the FileSourceMutation object of the builder.
func (fsu *FileSourceUpdate) Mutation() *FileSourceMutation {
	return fsu.mutation
}

// ClearIdentities clears all "identities" edges to the FileIdentity entity.
func (fsu *FileSourceUpdate) ClearIdentities() *FileSourceUpdate {
	fsu.mutation.ClearIdentities()
	return fsu
}

// RemoveIdentityIDs removes the "identities" edge to FileIdentity entities by IDs.
func (fsu *FileSourceUpdate) RemoveIdentityIDs(ids ...int) *FileSourceUpdate {
	fsu.mutation.RemoveIdentityIDs(ids...)
	return fsu
}

// RemoveIdentities removes "identities" edges to FileIdentity entities.
func (fsu *FileSourceUpdate) RemoveIdentities(f ...*FileIdentity) *FileSourceUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsu.RemoveIdentityIDs(ids...)
}

// ClearFiles clears all "files" edges to the File entity.
func (fsu *FileSourceUpdate) ClearFiles() *FileSourceUpdate {
	fsu.mutation.ClearFiles()
	return fsu
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (fsu *FileSourceUpdate) RemoveFileIDs(ids ...int) *FileSourceUpdate {
	fsu.mutation.RemoveFileIDs(ids...)
	return fsu
}

// RemoveFiles removes "files" edges to File entities.
func (fsu *FileSourceUpdate) RemoveFiles(f ...*File) *FileSourceUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fsu *FileSourceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fsu.sqlSave, fsu.mutation, fsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fsu *FileSourceUpdate) SaveX(ctx context.Context) int {
	affected, err := fsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fsu *FileSourceUpdate) Exec(ctx context.Context) error {
	_, err := fsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsu *FileSourceUpdate) ExecX(ctx context.Context) {
	if err := fsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fsu *FileSourceUpdate) check() error {
	if v, ok := fsu.mutation.Kind(); ok {
		if err := filesource.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "FileSource.kind": %w`, err)}
		}
	}
	if v, ok := fsu.mutation.Endpoint(); ok {
		if err := filesource.EndpointValidator(v); err != nil {
			return &ValidationError{Name: "endpoint", err: fmt.Errorf(`ent: validator failed for field "FileSource.endpoint": %w`, err)}
		}
	}
	if v, ok := fsu.mutation.StsEndpoint(); ok {
		if err := filesource.StsEndpointValidator(v); err != nil {
			return &ValidationError{Name: "sts_endpoint", err: fmt.Errorf(`ent: validator failed for field "FileSource.sts_endpoint": %w`, err)}
		}
	}
	if v, ok := fsu.mutation.Region(); ok {
		if err := filesource.RegionValidator(v); err != nil {
			return &ValidationError{Name: "region", err: fmt.Errorf(`ent: validator failed for field "FileSource.region": %w`, err)}
		}
	}
	if v, ok := fsu.mutation.Bucket(); ok {
		if err := filesource.BucketValidator(v); err != nil {
			return &ValidationError{Name: "bucket", err: fmt.Errorf(`ent: validator failed for field "FileSource.bucket": %w`, err)}
		}
	}
	if v, ok := fsu.mutation.BucketURL(); ok {
		if err := filesource.BucketURLValidator(v); err != nil {
			return &ValidationError{Name: "bucket_url", err: fmt.Errorf(`ent: validator failed for field "FileSource.bucket_url": %w`, err)}
		}
	}
	return nil
}

func (fsu *FileSourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(filesource.Table, filesource.Columns, sqlgraph.NewFieldSpec(filesource.FieldID, field.TypeInt))
	if ps := fsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fsu.mutation.UpdatedBy(); ok {
		_spec.SetField(filesource.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := fsu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(filesource.FieldUpdatedBy, field.TypeInt, value)
	}
	if fsu.mutation.UpdatedByCleared() {
		_spec.ClearField(filesource.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := fsu.mutation.UpdatedAt(); ok {
		_spec.SetField(filesource.FieldUpdatedAt, field.TypeTime, value)
	}
	if fsu.mutation.UpdatedAtCleared() {
		_spec.ClearField(filesource.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := fsu.mutation.Kind(); ok {
		_spec.SetField(filesource.FieldKind, field.TypeEnum, value)
	}
	if value, ok := fsu.mutation.Comments(); ok {
		_spec.SetField(filesource.FieldComments, field.TypeString, value)
	}
	if fsu.mutation.CommentsCleared() {
		_spec.ClearField(filesource.FieldComments, field.TypeString)
	}
	if value, ok := fsu.mutation.Endpoint(); ok {
		_spec.SetField(filesource.FieldEndpoint, field.TypeString, value)
	}
	if value, ok := fsu.mutation.EndpointImmutable(); ok {
		_spec.SetField(filesource.FieldEndpointImmutable, field.TypeBool, value)
	}
	if value, ok := fsu.mutation.StsEndpoint(); ok {
		_spec.SetField(filesource.FieldStsEndpoint, field.TypeString, value)
	}
	if value, ok := fsu.mutation.Region(); ok {
		_spec.SetField(filesource.FieldRegion, field.TypeString, value)
	}
	if value, ok := fsu.mutation.Bucket(); ok {
		_spec.SetField(filesource.FieldBucket, field.TypeString, value)
	}
	if value, ok := fsu.mutation.BucketURL(); ok {
		_spec.SetField(filesource.FieldBucketURL, field.TypeString, value)
	}
	if fsu.mutation.IdentitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.IdentitiesTable,
			Columns: []string{filesource.IdentitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fileidentity.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsu.mutation.RemovedIdentitiesIDs(); len(nodes) > 0 && !fsu.mutation.IdentitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.IdentitiesTable,
			Columns: []string{filesource.IdentitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fileidentity.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsu.mutation.IdentitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.IdentitiesTable,
			Columns: []string{filesource.IdentitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fileidentity.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fsu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.FilesTable,
			Columns: []string{filesource.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsu.mutation.RemovedFilesIDs(); len(nodes) > 0 && !fsu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.FilesTable,
			Columns: []string{filesource.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsu.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.FilesTable,
			Columns: []string{filesource.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filesource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fsu.mutation.done = true
	return n, nil
}

// FileSourceUpdateOne is the builder for updating a single FileSource entity.
type FileSourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileSourceMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (fsuo *FileSourceUpdateOne) SetUpdatedBy(i int) *FileSourceUpdateOne {
	fsuo.mutation.ResetUpdatedBy()
	fsuo.mutation.SetUpdatedBy(i)
	return fsuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fsuo *FileSourceUpdateOne) SetNillableUpdatedBy(i *int) *FileSourceUpdateOne {
	if i != nil {
		fsuo.SetUpdatedBy(*i)
	}
	return fsuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (fsuo *FileSourceUpdateOne) AddUpdatedBy(i int) *FileSourceUpdateOne {
	fsuo.mutation.AddUpdatedBy(i)
	return fsuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (fsuo *FileSourceUpdateOne) ClearUpdatedBy() *FileSourceUpdateOne {
	fsuo.mutation.ClearUpdatedBy()
	return fsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fsuo *FileSourceUpdateOne) SetUpdatedAt(t time.Time) *FileSourceUpdateOne {
	fsuo.mutation.SetUpdatedAt(t)
	return fsuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fsuo *FileSourceUpdateOne) SetNillableUpdatedAt(t *time.Time) *FileSourceUpdateOne {
	if t != nil {
		fsuo.SetUpdatedAt(*t)
	}
	return fsuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fsuo *FileSourceUpdateOne) ClearUpdatedAt() *FileSourceUpdateOne {
	fsuo.mutation.ClearUpdatedAt()
	return fsuo
}

// SetKind sets the "kind" field.
func (fsuo *FileSourceUpdateOne) SetKind(f filesource.Kind) *FileSourceUpdateOne {
	fsuo.mutation.SetKind(f)
	return fsuo
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (fsuo *FileSourceUpdateOne) SetNillableKind(f *filesource.Kind) *FileSourceUpdateOne {
	if f != nil {
		fsuo.SetKind(*f)
	}
	return fsuo
}

// SetComments sets the "comments" field.
func (fsuo *FileSourceUpdateOne) SetComments(s string) *FileSourceUpdateOne {
	fsuo.mutation.SetComments(s)
	return fsuo
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (fsuo *FileSourceUpdateOne) SetNillableComments(s *string) *FileSourceUpdateOne {
	if s != nil {
		fsuo.SetComments(*s)
	}
	return fsuo
}

// ClearComments clears the value of the "comments" field.
func (fsuo *FileSourceUpdateOne) ClearComments() *FileSourceUpdateOne {
	fsuo.mutation.ClearComments()
	return fsuo
}

// SetEndpoint sets the "endpoint" field.
func (fsuo *FileSourceUpdateOne) SetEndpoint(s string) *FileSourceUpdateOne {
	fsuo.mutation.SetEndpoint(s)
	return fsuo
}

// SetNillableEndpoint sets the "endpoint" field if the given value is not nil.
func (fsuo *FileSourceUpdateOne) SetNillableEndpoint(s *string) *FileSourceUpdateOne {
	if s != nil {
		fsuo.SetEndpoint(*s)
	}
	return fsuo
}

// SetEndpointImmutable sets the "endpoint_immutable" field.
func (fsuo *FileSourceUpdateOne) SetEndpointImmutable(b bool) *FileSourceUpdateOne {
	fsuo.mutation.SetEndpointImmutable(b)
	return fsuo
}

// SetNillableEndpointImmutable sets the "endpoint_immutable" field if the given value is not nil.
func (fsuo *FileSourceUpdateOne) SetNillableEndpointImmutable(b *bool) *FileSourceUpdateOne {
	if b != nil {
		fsuo.SetEndpointImmutable(*b)
	}
	return fsuo
}

// SetStsEndpoint sets the "sts_endpoint" field.
func (fsuo *FileSourceUpdateOne) SetStsEndpoint(s string) *FileSourceUpdateOne {
	fsuo.mutation.SetStsEndpoint(s)
	return fsuo
}

// SetNillableStsEndpoint sets the "sts_endpoint" field if the given value is not nil.
func (fsuo *FileSourceUpdateOne) SetNillableStsEndpoint(s *string) *FileSourceUpdateOne {
	if s != nil {
		fsuo.SetStsEndpoint(*s)
	}
	return fsuo
}

// SetRegion sets the "region" field.
func (fsuo *FileSourceUpdateOne) SetRegion(s string) *FileSourceUpdateOne {
	fsuo.mutation.SetRegion(s)
	return fsuo
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (fsuo *FileSourceUpdateOne) SetNillableRegion(s *string) *FileSourceUpdateOne {
	if s != nil {
		fsuo.SetRegion(*s)
	}
	return fsuo
}

// SetBucket sets the "bucket" field.
func (fsuo *FileSourceUpdateOne) SetBucket(s string) *FileSourceUpdateOne {
	fsuo.mutation.SetBucket(s)
	return fsuo
}

// SetNillableBucket sets the "bucket" field if the given value is not nil.
func (fsuo *FileSourceUpdateOne) SetNillableBucket(s *string) *FileSourceUpdateOne {
	if s != nil {
		fsuo.SetBucket(*s)
	}
	return fsuo
}

// SetBucketURL sets the "bucket_url" field.
func (fsuo *FileSourceUpdateOne) SetBucketURL(s string) *FileSourceUpdateOne {
	fsuo.mutation.SetBucketURL(s)
	return fsuo
}

// SetNillableBucketURL sets the "bucket_url" field if the given value is not nil.
func (fsuo *FileSourceUpdateOne) SetNillableBucketURL(s *string) *FileSourceUpdateOne {
	if s != nil {
		fsuo.SetBucketURL(*s)
	}
	return fsuo
}

// AddIdentityIDs adds the "identities" edge to the FileIdentity entity by IDs.
func (fsuo *FileSourceUpdateOne) AddIdentityIDs(ids ...int) *FileSourceUpdateOne {
	fsuo.mutation.AddIdentityIDs(ids...)
	return fsuo
}

// AddIdentities adds the "identities" edges to the FileIdentity entity.
func (fsuo *FileSourceUpdateOne) AddIdentities(f ...*FileIdentity) *FileSourceUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsuo.AddIdentityIDs(ids...)
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (fsuo *FileSourceUpdateOne) AddFileIDs(ids ...int) *FileSourceUpdateOne {
	fsuo.mutation.AddFileIDs(ids...)
	return fsuo
}

// AddFiles adds the "files" edges to the File entity.
func (fsuo *FileSourceUpdateOne) AddFiles(f ...*File) *FileSourceUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsuo.AddFileIDs(ids...)
}

// Mutation returns the FileSourceMutation object of the builder.
func (fsuo *FileSourceUpdateOne) Mutation() *FileSourceMutation {
	return fsuo.mutation
}

// ClearIdentities clears all "identities" edges to the FileIdentity entity.
func (fsuo *FileSourceUpdateOne) ClearIdentities() *FileSourceUpdateOne {
	fsuo.mutation.ClearIdentities()
	return fsuo
}

// RemoveIdentityIDs removes the "identities" edge to FileIdentity entities by IDs.
func (fsuo *FileSourceUpdateOne) RemoveIdentityIDs(ids ...int) *FileSourceUpdateOne {
	fsuo.mutation.RemoveIdentityIDs(ids...)
	return fsuo
}

// RemoveIdentities removes "identities" edges to FileIdentity entities.
func (fsuo *FileSourceUpdateOne) RemoveIdentities(f ...*FileIdentity) *FileSourceUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsuo.RemoveIdentityIDs(ids...)
}

// ClearFiles clears all "files" edges to the File entity.
func (fsuo *FileSourceUpdateOne) ClearFiles() *FileSourceUpdateOne {
	fsuo.mutation.ClearFiles()
	return fsuo
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (fsuo *FileSourceUpdateOne) RemoveFileIDs(ids ...int) *FileSourceUpdateOne {
	fsuo.mutation.RemoveFileIDs(ids...)
	return fsuo
}

// RemoveFiles removes "files" edges to File entities.
func (fsuo *FileSourceUpdateOne) RemoveFiles(f ...*File) *FileSourceUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fsuo.RemoveFileIDs(ids...)
}

// Where appends a list predicates to the FileSourceUpdate builder.
func (fsuo *FileSourceUpdateOne) Where(ps ...predicate.FileSource) *FileSourceUpdateOne {
	fsuo.mutation.Where(ps...)
	return fsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fsuo *FileSourceUpdateOne) Select(field string, fields ...string) *FileSourceUpdateOne {
	fsuo.fields = append([]string{field}, fields...)
	return fsuo
}

// Save executes the query and returns the updated FileSource entity.
func (fsuo *FileSourceUpdateOne) Save(ctx context.Context) (*FileSource, error) {
	return withHooks(ctx, fsuo.sqlSave, fsuo.mutation, fsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fsuo *FileSourceUpdateOne) SaveX(ctx context.Context) *FileSource {
	node, err := fsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fsuo *FileSourceUpdateOne) Exec(ctx context.Context) error {
	_, err := fsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fsuo *FileSourceUpdateOne) ExecX(ctx context.Context) {
	if err := fsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fsuo *FileSourceUpdateOne) check() error {
	if v, ok := fsuo.mutation.Kind(); ok {
		if err := filesource.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "FileSource.kind": %w`, err)}
		}
	}
	if v, ok := fsuo.mutation.Endpoint(); ok {
		if err := filesource.EndpointValidator(v); err != nil {
			return &ValidationError{Name: "endpoint", err: fmt.Errorf(`ent: validator failed for field "FileSource.endpoint": %w`, err)}
		}
	}
	if v, ok := fsuo.mutation.StsEndpoint(); ok {
		if err := filesource.StsEndpointValidator(v); err != nil {
			return &ValidationError{Name: "sts_endpoint", err: fmt.Errorf(`ent: validator failed for field "FileSource.sts_endpoint": %w`, err)}
		}
	}
	if v, ok := fsuo.mutation.Region(); ok {
		if err := filesource.RegionValidator(v); err != nil {
			return &ValidationError{Name: "region", err: fmt.Errorf(`ent: validator failed for field "FileSource.region": %w`, err)}
		}
	}
	if v, ok := fsuo.mutation.Bucket(); ok {
		if err := filesource.BucketValidator(v); err != nil {
			return &ValidationError{Name: "bucket", err: fmt.Errorf(`ent: validator failed for field "FileSource.bucket": %w`, err)}
		}
	}
	if v, ok := fsuo.mutation.BucketURL(); ok {
		if err := filesource.BucketURLValidator(v); err != nil {
			return &ValidationError{Name: "bucket_url", err: fmt.Errorf(`ent: validator failed for field "FileSource.bucket_url": %w`, err)}
		}
	}
	return nil
}

func (fsuo *FileSourceUpdateOne) sqlSave(ctx context.Context) (_node *FileSource, err error) {
	if err := fsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(filesource.Table, filesource.Columns, sqlgraph.NewFieldSpec(filesource.FieldID, field.TypeInt))
	id, ok := fsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FileSource.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filesource.FieldID)
		for _, f := range fields {
			if !filesource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != filesource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fsuo.mutation.UpdatedBy(); ok {
		_spec.SetField(filesource.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := fsuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(filesource.FieldUpdatedBy, field.TypeInt, value)
	}
	if fsuo.mutation.UpdatedByCleared() {
		_spec.ClearField(filesource.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := fsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(filesource.FieldUpdatedAt, field.TypeTime, value)
	}
	if fsuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(filesource.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := fsuo.mutation.Kind(); ok {
		_spec.SetField(filesource.FieldKind, field.TypeEnum, value)
	}
	if value, ok := fsuo.mutation.Comments(); ok {
		_spec.SetField(filesource.FieldComments, field.TypeString, value)
	}
	if fsuo.mutation.CommentsCleared() {
		_spec.ClearField(filesource.FieldComments, field.TypeString)
	}
	if value, ok := fsuo.mutation.Endpoint(); ok {
		_spec.SetField(filesource.FieldEndpoint, field.TypeString, value)
	}
	if value, ok := fsuo.mutation.EndpointImmutable(); ok {
		_spec.SetField(filesource.FieldEndpointImmutable, field.TypeBool, value)
	}
	if value, ok := fsuo.mutation.StsEndpoint(); ok {
		_spec.SetField(filesource.FieldStsEndpoint, field.TypeString, value)
	}
	if value, ok := fsuo.mutation.Region(); ok {
		_spec.SetField(filesource.FieldRegion, field.TypeString, value)
	}
	if value, ok := fsuo.mutation.Bucket(); ok {
		_spec.SetField(filesource.FieldBucket, field.TypeString, value)
	}
	if value, ok := fsuo.mutation.BucketURL(); ok {
		_spec.SetField(filesource.FieldBucketURL, field.TypeString, value)
	}
	if fsuo.mutation.IdentitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.IdentitiesTable,
			Columns: []string{filesource.IdentitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fileidentity.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsuo.mutation.RemovedIdentitiesIDs(); len(nodes) > 0 && !fsuo.mutation.IdentitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.IdentitiesTable,
			Columns: []string{filesource.IdentitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fileidentity.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsuo.mutation.IdentitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.IdentitiesTable,
			Columns: []string{filesource.IdentitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fileidentity.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fsuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.FilesTable,
			Columns: []string{filesource.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsuo.mutation.RemovedFilesIDs(); len(nodes) > 0 && !fsuo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.FilesTable,
			Columns: []string{filesource.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fsuo.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filesource.FilesTable,
			Columns: []string{filesource.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FileSource{config: fsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filesource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fsuo.mutation.done = true
	return _node, nil
}
