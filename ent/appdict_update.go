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
	"github.com/woocoos/knockout/ent/appdict"
	"github.com/woocoos/knockout/ent/appdictitem"
	"github.com/woocoos/knockout/ent/predicate"
)

// AppDictUpdate is the builder for updating AppDict entities.
type AppDictUpdate struct {
	config
	hooks    []Hook
	mutation *AppDictMutation
}

// Where appends a list predicates to the AppDictUpdate builder.
func (adu *AppDictUpdate) Where(ps ...predicate.AppDict) *AppDictUpdate {
	adu.mutation.Where(ps...)
	return adu
}

// SetUpdatedBy sets the "updated_by" field.
func (adu *AppDictUpdate) SetUpdatedBy(i int) *AppDictUpdate {
	adu.mutation.ResetUpdatedBy()
	adu.mutation.SetUpdatedBy(i)
	return adu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (adu *AppDictUpdate) SetNillableUpdatedBy(i *int) *AppDictUpdate {
	if i != nil {
		adu.SetUpdatedBy(*i)
	}
	return adu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (adu *AppDictUpdate) AddUpdatedBy(i int) *AppDictUpdate {
	adu.mutation.AddUpdatedBy(i)
	return adu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (adu *AppDictUpdate) ClearUpdatedBy() *AppDictUpdate {
	adu.mutation.ClearUpdatedBy()
	return adu
}

// SetUpdatedAt sets the "updated_at" field.
func (adu *AppDictUpdate) SetUpdatedAt(t time.Time) *AppDictUpdate {
	adu.mutation.SetUpdatedAt(t)
	return adu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (adu *AppDictUpdate) SetNillableUpdatedAt(t *time.Time) *AppDictUpdate {
	if t != nil {
		adu.SetUpdatedAt(*t)
	}
	return adu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (adu *AppDictUpdate) ClearUpdatedAt() *AppDictUpdate {
	adu.mutation.ClearUpdatedAt()
	return adu
}

// SetName sets the "name" field.
func (adu *AppDictUpdate) SetName(s string) *AppDictUpdate {
	adu.mutation.SetName(s)
	return adu
}

// SetComments sets the "comments" field.
func (adu *AppDictUpdate) SetComments(s string) *AppDictUpdate {
	adu.mutation.SetComments(s)
	return adu
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (adu *AppDictUpdate) SetNillableComments(s *string) *AppDictUpdate {
	if s != nil {
		adu.SetComments(*s)
	}
	return adu
}

// ClearComments clears the value of the "comments" field.
func (adu *AppDictUpdate) ClearComments() *AppDictUpdate {
	adu.mutation.ClearComments()
	return adu
}

// AddItemIDs adds the "items" edge to the AppDictItem entity by IDs.
func (adu *AppDictUpdate) AddItemIDs(ids ...int) *AppDictUpdate {
	adu.mutation.AddItemIDs(ids...)
	return adu
}

// AddItems adds the "items" edges to the AppDictItem entity.
func (adu *AppDictUpdate) AddItems(a ...*AppDictItem) *AppDictUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return adu.AddItemIDs(ids...)
}

// Mutation returns the AppDictMutation object of the builder.
func (adu *AppDictUpdate) Mutation() *AppDictMutation {
	return adu.mutation
}

// ClearItems clears all "items" edges to the AppDictItem entity.
func (adu *AppDictUpdate) ClearItems() *AppDictUpdate {
	adu.mutation.ClearItems()
	return adu
}

// RemoveItemIDs removes the "items" edge to AppDictItem entities by IDs.
func (adu *AppDictUpdate) RemoveItemIDs(ids ...int) *AppDictUpdate {
	adu.mutation.RemoveItemIDs(ids...)
	return adu
}

// RemoveItems removes "items" edges to AppDictItem entities.
func (adu *AppDictUpdate) RemoveItems(a ...*AppDictItem) *AppDictUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return adu.RemoveItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (adu *AppDictUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, adu.sqlSave, adu.mutation, adu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (adu *AppDictUpdate) SaveX(ctx context.Context) int {
	affected, err := adu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (adu *AppDictUpdate) Exec(ctx context.Context) error {
	_, err := adu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adu *AppDictUpdate) ExecX(ctx context.Context) {
	if err := adu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (adu *AppDictUpdate) check() error {
	if v, ok := adu.mutation.Name(); ok {
		if err := appdict.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AppDict.name": %w`, err)}
		}
	}
	return nil
}

func (adu *AppDictUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := adu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(appdict.Table, appdict.Columns, sqlgraph.NewFieldSpec(appdict.FieldID, field.TypeInt))
	if ps := adu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := adu.mutation.UpdatedBy(); ok {
		_spec.SetField(appdict.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := adu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(appdict.FieldUpdatedBy, field.TypeInt, value)
	}
	if adu.mutation.UpdatedByCleared() {
		_spec.ClearField(appdict.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := adu.mutation.UpdatedAt(); ok {
		_spec.SetField(appdict.FieldUpdatedAt, field.TypeTime, value)
	}
	if adu.mutation.UpdatedAtCleared() {
		_spec.ClearField(appdict.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := adu.mutation.Name(); ok {
		_spec.SetField(appdict.FieldName, field.TypeString, value)
	}
	if value, ok := adu.mutation.Comments(); ok {
		_spec.SetField(appdict.FieldComments, field.TypeString, value)
	}
	if adu.mutation.CommentsCleared() {
		_spec.ClearField(appdict.FieldComments, field.TypeString)
	}
	if adu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appdict.ItemsTable,
			Columns: []string{appdict.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appdictitem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := adu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !adu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appdict.ItemsTable,
			Columns: []string{appdict.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appdictitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := adu.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appdict.ItemsTable,
			Columns: []string{appdict.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appdictitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, adu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appdict.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	adu.mutation.done = true
	return n, nil
}

// AppDictUpdateOne is the builder for updating a single AppDict entity.
type AppDictUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppDictMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (aduo *AppDictUpdateOne) SetUpdatedBy(i int) *AppDictUpdateOne {
	aduo.mutation.ResetUpdatedBy()
	aduo.mutation.SetUpdatedBy(i)
	return aduo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (aduo *AppDictUpdateOne) SetNillableUpdatedBy(i *int) *AppDictUpdateOne {
	if i != nil {
		aduo.SetUpdatedBy(*i)
	}
	return aduo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (aduo *AppDictUpdateOne) AddUpdatedBy(i int) *AppDictUpdateOne {
	aduo.mutation.AddUpdatedBy(i)
	return aduo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (aduo *AppDictUpdateOne) ClearUpdatedBy() *AppDictUpdateOne {
	aduo.mutation.ClearUpdatedBy()
	return aduo
}

// SetUpdatedAt sets the "updated_at" field.
func (aduo *AppDictUpdateOne) SetUpdatedAt(t time.Time) *AppDictUpdateOne {
	aduo.mutation.SetUpdatedAt(t)
	return aduo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aduo *AppDictUpdateOne) SetNillableUpdatedAt(t *time.Time) *AppDictUpdateOne {
	if t != nil {
		aduo.SetUpdatedAt(*t)
	}
	return aduo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (aduo *AppDictUpdateOne) ClearUpdatedAt() *AppDictUpdateOne {
	aduo.mutation.ClearUpdatedAt()
	return aduo
}

// SetName sets the "name" field.
func (aduo *AppDictUpdateOne) SetName(s string) *AppDictUpdateOne {
	aduo.mutation.SetName(s)
	return aduo
}

// SetComments sets the "comments" field.
func (aduo *AppDictUpdateOne) SetComments(s string) *AppDictUpdateOne {
	aduo.mutation.SetComments(s)
	return aduo
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (aduo *AppDictUpdateOne) SetNillableComments(s *string) *AppDictUpdateOne {
	if s != nil {
		aduo.SetComments(*s)
	}
	return aduo
}

// ClearComments clears the value of the "comments" field.
func (aduo *AppDictUpdateOne) ClearComments() *AppDictUpdateOne {
	aduo.mutation.ClearComments()
	return aduo
}

// AddItemIDs adds the "items" edge to the AppDictItem entity by IDs.
func (aduo *AppDictUpdateOne) AddItemIDs(ids ...int) *AppDictUpdateOne {
	aduo.mutation.AddItemIDs(ids...)
	return aduo
}

// AddItems adds the "items" edges to the AppDictItem entity.
func (aduo *AppDictUpdateOne) AddItems(a ...*AppDictItem) *AppDictUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aduo.AddItemIDs(ids...)
}

// Mutation returns the AppDictMutation object of the builder.
func (aduo *AppDictUpdateOne) Mutation() *AppDictMutation {
	return aduo.mutation
}

// ClearItems clears all "items" edges to the AppDictItem entity.
func (aduo *AppDictUpdateOne) ClearItems() *AppDictUpdateOne {
	aduo.mutation.ClearItems()
	return aduo
}

// RemoveItemIDs removes the "items" edge to AppDictItem entities by IDs.
func (aduo *AppDictUpdateOne) RemoveItemIDs(ids ...int) *AppDictUpdateOne {
	aduo.mutation.RemoveItemIDs(ids...)
	return aduo
}

// RemoveItems removes "items" edges to AppDictItem entities.
func (aduo *AppDictUpdateOne) RemoveItems(a ...*AppDictItem) *AppDictUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aduo.RemoveItemIDs(ids...)
}

// Where appends a list predicates to the AppDictUpdate builder.
func (aduo *AppDictUpdateOne) Where(ps ...predicate.AppDict) *AppDictUpdateOne {
	aduo.mutation.Where(ps...)
	return aduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aduo *AppDictUpdateOne) Select(field string, fields ...string) *AppDictUpdateOne {
	aduo.fields = append([]string{field}, fields...)
	return aduo
}

// Save executes the query and returns the updated AppDict entity.
func (aduo *AppDictUpdateOne) Save(ctx context.Context) (*AppDict, error) {
	return withHooks(ctx, aduo.sqlSave, aduo.mutation, aduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aduo *AppDictUpdateOne) SaveX(ctx context.Context) *AppDict {
	node, err := aduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aduo *AppDictUpdateOne) Exec(ctx context.Context) error {
	_, err := aduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aduo *AppDictUpdateOne) ExecX(ctx context.Context) {
	if err := aduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aduo *AppDictUpdateOne) check() error {
	if v, ok := aduo.mutation.Name(); ok {
		if err := appdict.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AppDict.name": %w`, err)}
		}
	}
	return nil
}

func (aduo *AppDictUpdateOne) sqlSave(ctx context.Context) (_node *AppDict, err error) {
	if err := aduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(appdict.Table, appdict.Columns, sqlgraph.NewFieldSpec(appdict.FieldID, field.TypeInt))
	id, ok := aduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppDict.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appdict.FieldID)
		for _, f := range fields {
			if !appdict.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appdict.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aduo.mutation.UpdatedBy(); ok {
		_spec.SetField(appdict.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := aduo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(appdict.FieldUpdatedBy, field.TypeInt, value)
	}
	if aduo.mutation.UpdatedByCleared() {
		_spec.ClearField(appdict.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := aduo.mutation.UpdatedAt(); ok {
		_spec.SetField(appdict.FieldUpdatedAt, field.TypeTime, value)
	}
	if aduo.mutation.UpdatedAtCleared() {
		_spec.ClearField(appdict.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := aduo.mutation.Name(); ok {
		_spec.SetField(appdict.FieldName, field.TypeString, value)
	}
	if value, ok := aduo.mutation.Comments(); ok {
		_spec.SetField(appdict.FieldComments, field.TypeString, value)
	}
	if aduo.mutation.CommentsCleared() {
		_spec.ClearField(appdict.FieldComments, field.TypeString)
	}
	if aduo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appdict.ItemsTable,
			Columns: []string{appdict.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appdictitem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aduo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !aduo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appdict.ItemsTable,
			Columns: []string{appdict.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appdictitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aduo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appdict.ItemsTable,
			Columns: []string{appdict.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appdictitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AppDict{config: aduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appdict.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aduo.mutation.done = true
	return _node, nil
}
