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
	"github.com/woocoos/knockout/ent/appres"
	"github.com/woocoos/knockout/ent/predicate"
)

// AppResUpdate is the builder for updating AppRes entities.
type AppResUpdate struct {
	config
	hooks    []Hook
	mutation *AppResMutation
}

// Where appends a list predicates to the AppResUpdate builder.
func (aru *AppResUpdate) Where(ps ...predicate.AppRes) *AppResUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetUpdatedBy sets the "updated_by" field.
func (aru *AppResUpdate) SetUpdatedBy(i int) *AppResUpdate {
	aru.mutation.ResetUpdatedBy()
	aru.mutation.SetUpdatedBy(i)
	return aru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (aru *AppResUpdate) SetNillableUpdatedBy(i *int) *AppResUpdate {
	if i != nil {
		aru.SetUpdatedBy(*i)
	}
	return aru
}

// AddUpdatedBy adds i to the "updated_by" field.
func (aru *AppResUpdate) AddUpdatedBy(i int) *AppResUpdate {
	aru.mutation.AddUpdatedBy(i)
	return aru
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (aru *AppResUpdate) ClearUpdatedBy() *AppResUpdate {
	aru.mutation.ClearUpdatedBy()
	return aru
}

// SetUpdatedAt sets the "updated_at" field.
func (aru *AppResUpdate) SetUpdatedAt(t time.Time) *AppResUpdate {
	aru.mutation.SetUpdatedAt(t)
	return aru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aru *AppResUpdate) SetNillableUpdatedAt(t *time.Time) *AppResUpdate {
	if t != nil {
		aru.SetUpdatedAt(*t)
	}
	return aru
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (aru *AppResUpdate) ClearUpdatedAt() *AppResUpdate {
	aru.mutation.ClearUpdatedAt()
	return aru
}

// SetName sets the "name" field.
func (aru *AppResUpdate) SetName(s string) *AppResUpdate {
	aru.mutation.SetName(s)
	return aru
}

// SetTypeName sets the "type_name" field.
func (aru *AppResUpdate) SetTypeName(s string) *AppResUpdate {
	aru.mutation.SetTypeName(s)
	return aru
}

// SetArnPattern sets the "arn_pattern" field.
func (aru *AppResUpdate) SetArnPattern(s string) *AppResUpdate {
	aru.mutation.SetArnPattern(s)
	return aru
}

// Mutation returns the AppResMutation object of the builder.
func (aru *AppResUpdate) Mutation() *AppResMutation {
	return aru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *AppResUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aru.sqlSave, aru.mutation, aru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aru *AppResUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *AppResUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *AppResUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aru *AppResUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(appres.Table, appres.Columns, sqlgraph.NewFieldSpec(appres.FieldID, field.TypeInt))
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.UpdatedBy(); ok {
		_spec.SetField(appres.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := aru.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(appres.FieldUpdatedBy, field.TypeInt, value)
	}
	if aru.mutation.UpdatedByCleared() {
		_spec.ClearField(appres.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := aru.mutation.UpdatedAt(); ok {
		_spec.SetField(appres.FieldUpdatedAt, field.TypeTime, value)
	}
	if aru.mutation.UpdatedAtCleared() {
		_spec.ClearField(appres.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := aru.mutation.Name(); ok {
		_spec.SetField(appres.FieldName, field.TypeString, value)
	}
	if value, ok := aru.mutation.TypeName(); ok {
		_spec.SetField(appres.FieldTypeName, field.TypeString, value)
	}
	if value, ok := aru.mutation.ArnPattern(); ok {
		_spec.SetField(appres.FieldArnPattern, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appres.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aru.mutation.done = true
	return n, nil
}

// AppResUpdateOne is the builder for updating a single AppRes entity.
type AppResUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppResMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (aruo *AppResUpdateOne) SetUpdatedBy(i int) *AppResUpdateOne {
	aruo.mutation.ResetUpdatedBy()
	aruo.mutation.SetUpdatedBy(i)
	return aruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (aruo *AppResUpdateOne) SetNillableUpdatedBy(i *int) *AppResUpdateOne {
	if i != nil {
		aruo.SetUpdatedBy(*i)
	}
	return aruo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (aruo *AppResUpdateOne) AddUpdatedBy(i int) *AppResUpdateOne {
	aruo.mutation.AddUpdatedBy(i)
	return aruo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (aruo *AppResUpdateOne) ClearUpdatedBy() *AppResUpdateOne {
	aruo.mutation.ClearUpdatedBy()
	return aruo
}

// SetUpdatedAt sets the "updated_at" field.
func (aruo *AppResUpdateOne) SetUpdatedAt(t time.Time) *AppResUpdateOne {
	aruo.mutation.SetUpdatedAt(t)
	return aruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aruo *AppResUpdateOne) SetNillableUpdatedAt(t *time.Time) *AppResUpdateOne {
	if t != nil {
		aruo.SetUpdatedAt(*t)
	}
	return aruo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (aruo *AppResUpdateOne) ClearUpdatedAt() *AppResUpdateOne {
	aruo.mutation.ClearUpdatedAt()
	return aruo
}

// SetName sets the "name" field.
func (aruo *AppResUpdateOne) SetName(s string) *AppResUpdateOne {
	aruo.mutation.SetName(s)
	return aruo
}

// SetTypeName sets the "type_name" field.
func (aruo *AppResUpdateOne) SetTypeName(s string) *AppResUpdateOne {
	aruo.mutation.SetTypeName(s)
	return aruo
}

// SetArnPattern sets the "arn_pattern" field.
func (aruo *AppResUpdateOne) SetArnPattern(s string) *AppResUpdateOne {
	aruo.mutation.SetArnPattern(s)
	return aruo
}

// Mutation returns the AppResMutation object of the builder.
func (aruo *AppResUpdateOne) Mutation() *AppResMutation {
	return aruo.mutation
}

// Where appends a list predicates to the AppResUpdate builder.
func (aruo *AppResUpdateOne) Where(ps ...predicate.AppRes) *AppResUpdateOne {
	aruo.mutation.Where(ps...)
	return aruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *AppResUpdateOne) Select(field string, fields ...string) *AppResUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated AppRes entity.
func (aruo *AppResUpdateOne) Save(ctx context.Context) (*AppRes, error) {
	return withHooks(ctx, aruo.sqlSave, aruo.mutation, aruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *AppResUpdateOne) SaveX(ctx context.Context) *AppRes {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *AppResUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *AppResUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aruo *AppResUpdateOne) sqlSave(ctx context.Context) (_node *AppRes, err error) {
	_spec := sqlgraph.NewUpdateSpec(appres.Table, appres.Columns, sqlgraph.NewFieldSpec(appres.FieldID, field.TypeInt))
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppRes.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appres.FieldID)
		for _, f := range fields {
			if !appres.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appres.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruo.mutation.UpdatedBy(); ok {
		_spec.SetField(appres.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := aruo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(appres.FieldUpdatedBy, field.TypeInt, value)
	}
	if aruo.mutation.UpdatedByCleared() {
		_spec.ClearField(appres.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := aruo.mutation.UpdatedAt(); ok {
		_spec.SetField(appres.FieldUpdatedAt, field.TypeTime, value)
	}
	if aruo.mutation.UpdatedAtCleared() {
		_spec.ClearField(appres.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := aruo.mutation.Name(); ok {
		_spec.SetField(appres.FieldName, field.TypeString, value)
	}
	if value, ok := aruo.mutation.TypeName(); ok {
		_spec.SetField(appres.FieldTypeName, field.TypeString, value)
	}
	if value, ok := aruo.mutation.ArnPattern(); ok {
		_spec.SetField(appres.FieldArnPattern, field.TypeString, value)
	}
	_node = &AppRes{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appres.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aruo.mutation.done = true
	return _node, nil
}
