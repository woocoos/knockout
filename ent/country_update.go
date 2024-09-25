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
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/ent/country"
	"github.com/woocoos/knockout/ent/predicate"
	"github.com/woocoos/knockout/ent/region"
)

// CountryUpdate is the builder for updating Country entities.
type CountryUpdate struct {
	config
	hooks    []Hook
	mutation *CountryMutation
}

// Where appends a list predicates to the CountryUpdate builder.
func (cu *CountryUpdate) Where(ps ...predicate.Country) *CountryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedBy sets the "updated_by" field.
func (cu *CountryUpdate) SetUpdatedBy(i int) *CountryUpdate {
	cu.mutation.ResetUpdatedBy()
	cu.mutation.SetUpdatedBy(i)
	return cu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableUpdatedBy(i *int) *CountryUpdate {
	if i != nil {
		cu.SetUpdatedBy(*i)
	}
	return cu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (cu *CountryUpdate) AddUpdatedBy(i int) *CountryUpdate {
	cu.mutation.AddUpdatedBy(i)
	return cu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (cu *CountryUpdate) ClearUpdatedBy() *CountryUpdate {
	cu.mutation.ClearUpdatedBy()
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CountryUpdate) SetUpdatedAt(t time.Time) *CountryUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableUpdatedAt(t *time.Time) *CountryUpdate {
	if t != nil {
		cu.SetUpdatedAt(*t)
	}
	return cu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cu *CountryUpdate) ClearUpdatedAt() *CountryUpdate {
	cu.mutation.ClearUpdatedAt()
	return cu
}

// SetName sets the "name" field.
func (cu *CountryUpdate) SetName(s string) *CountryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableName(s *string) *CountryUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of the "name" field.
func (cu *CountryUpdate) ClearName() *CountryUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetNameEn sets the "name_en" field.
func (cu *CountryUpdate) SetNameEn(s string) *CountryUpdate {
	cu.mutation.SetNameEn(s)
	return cu
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableNameEn(s *string) *CountryUpdate {
	if s != nil {
		cu.SetNameEn(*s)
	}
	return cu
}

// ClearNameEn clears the value of the "name_en" field.
func (cu *CountryUpdate) ClearNameEn() *CountryUpdate {
	cu.mutation.ClearNameEn()
	return cu
}

// SetCode sets the "code" field.
func (cu *CountryUpdate) SetCode(s string) *CountryUpdate {
	cu.mutation.SetCode(s)
	return cu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableCode(s *string) *CountryUpdate {
	if s != nil {
		cu.SetCode(*s)
	}
	return cu
}

// SetDisplaySort sets the "display_sort" field.
func (cu *CountryUpdate) SetDisplaySort(i int32) *CountryUpdate {
	cu.mutation.ResetDisplaySort()
	cu.mutation.SetDisplaySort(i)
	return cu
}

// SetNillableDisplaySort sets the "display_sort" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableDisplaySort(i *int32) *CountryUpdate {
	if i != nil {
		cu.SetDisplaySort(*i)
	}
	return cu
}

// AddDisplaySort adds i to the "display_sort" field.
func (cu *CountryUpdate) AddDisplaySort(i int32) *CountryUpdate {
	cu.mutation.AddDisplaySort(i)
	return cu
}

// ClearDisplaySort clears the value of the "display_sort" field.
func (cu *CountryUpdate) ClearDisplaySort() *CountryUpdate {
	cu.mutation.ClearDisplaySort()
	return cu
}

// SetStatus sets the "status" field.
func (cu *CountryUpdate) SetStatus(ts typex.SimpleStatus) *CountryUpdate {
	cu.mutation.SetStatus(ts)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *CountryUpdate) SetNillableStatus(ts *typex.SimpleStatus) *CountryUpdate {
	if ts != nil {
		cu.SetStatus(*ts)
	}
	return cu
}

// ClearStatus clears the value of the "status" field.
func (cu *CountryUpdate) ClearStatus() *CountryUpdate {
	cu.mutation.ClearStatus()
	return cu
}

// AddRegionIDs adds the "regions" edge to the Region entity by IDs.
func (cu *CountryUpdate) AddRegionIDs(ids ...int) *CountryUpdate {
	cu.mutation.AddRegionIDs(ids...)
	return cu
}

// AddRegions adds the "regions" edges to the Region entity.
func (cu *CountryUpdate) AddRegions(r ...*Region) *CountryUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.AddRegionIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cu *CountryUpdate) Mutation() *CountryMutation {
	return cu.mutation
}

// ClearRegions clears all "regions" edges to the Region entity.
func (cu *CountryUpdate) ClearRegions() *CountryUpdate {
	cu.mutation.ClearRegions()
	return cu
}

// RemoveRegionIDs removes the "regions" edge to Region entities by IDs.
func (cu *CountryUpdate) RemoveRegionIDs(ids ...int) *CountryUpdate {
	cu.mutation.RemoveRegionIDs(ids...)
	return cu
}

// RemoveRegions removes "regions" edges to Region entities.
func (cu *CountryUpdate) RemoveRegions(r ...*Region) *CountryUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.RemoveRegionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CountryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CountryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CountryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CountryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CountryUpdate) check() error {
	if v, ok := cu.mutation.Code(); ok {
		if err := country.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Country.code": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Status(); ok {
		if err := country.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Country.status": %w`, err)}
		}
	}
	return nil
}

func (cu *CountryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(country.Table, country.Columns, sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedBy(); ok {
		_spec.SetField(country.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(country.FieldUpdatedBy, field.TypeInt, value)
	}
	if cu.mutation.UpdatedByCleared() {
		_spec.ClearField(country.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(country.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.UpdatedAtCleared() {
		_spec.ClearField(country.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(country.FieldName, field.TypeString, value)
	}
	if cu.mutation.NameCleared() {
		_spec.ClearField(country.FieldName, field.TypeString)
	}
	if value, ok := cu.mutation.NameEn(); ok {
		_spec.SetField(country.FieldNameEn, field.TypeString, value)
	}
	if cu.mutation.NameEnCleared() {
		_spec.ClearField(country.FieldNameEn, field.TypeString)
	}
	if value, ok := cu.mutation.Code(); ok {
		_spec.SetField(country.FieldCode, field.TypeString, value)
	}
	if value, ok := cu.mutation.DisplaySort(); ok {
		_spec.SetField(country.FieldDisplaySort, field.TypeInt32, value)
	}
	if value, ok := cu.mutation.AddedDisplaySort(); ok {
		_spec.AddField(country.FieldDisplaySort, field.TypeInt32, value)
	}
	if cu.mutation.DisplaySortCleared() {
		_spec.ClearField(country.FieldDisplaySort, field.TypeInt32)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(country.FieldStatus, field.TypeEnum, value)
	}
	if cu.mutation.StatusCleared() {
		_spec.ClearField(country.FieldStatus, field.TypeEnum)
	}
	if cu.mutation.RegionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.RegionsTable,
			Columns: []string{country.RegionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(region.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedRegionsIDs(); len(nodes) > 0 && !cu.mutation.RegionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.RegionsTable,
			Columns: []string{country.RegionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(region.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RegionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.RegionsTable,
			Columns: []string{country.RegionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(region.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CountryUpdateOne is the builder for updating a single Country entity.
type CountryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CountryMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (cuo *CountryUpdateOne) SetUpdatedBy(i int) *CountryUpdateOne {
	cuo.mutation.ResetUpdatedBy()
	cuo.mutation.SetUpdatedBy(i)
	return cuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableUpdatedBy(i *int) *CountryUpdateOne {
	if i != nil {
		cuo.SetUpdatedBy(*i)
	}
	return cuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (cuo *CountryUpdateOne) AddUpdatedBy(i int) *CountryUpdateOne {
	cuo.mutation.AddUpdatedBy(i)
	return cuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (cuo *CountryUpdateOne) ClearUpdatedBy() *CountryUpdateOne {
	cuo.mutation.ClearUpdatedBy()
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CountryUpdateOne) SetUpdatedAt(t time.Time) *CountryUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableUpdatedAt(t *time.Time) *CountryUpdateOne {
	if t != nil {
		cuo.SetUpdatedAt(*t)
	}
	return cuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cuo *CountryUpdateOne) ClearUpdatedAt() *CountryUpdateOne {
	cuo.mutation.ClearUpdatedAt()
	return cuo
}

// SetName sets the "name" field.
func (cuo *CountryUpdateOne) SetName(s string) *CountryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableName(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of the "name" field.
func (cuo *CountryUpdateOne) ClearName() *CountryUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetNameEn sets the "name_en" field.
func (cuo *CountryUpdateOne) SetNameEn(s string) *CountryUpdateOne {
	cuo.mutation.SetNameEn(s)
	return cuo
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableNameEn(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetNameEn(*s)
	}
	return cuo
}

// ClearNameEn clears the value of the "name_en" field.
func (cuo *CountryUpdateOne) ClearNameEn() *CountryUpdateOne {
	cuo.mutation.ClearNameEn()
	return cuo
}

// SetCode sets the "code" field.
func (cuo *CountryUpdateOne) SetCode(s string) *CountryUpdateOne {
	cuo.mutation.SetCode(s)
	return cuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableCode(s *string) *CountryUpdateOne {
	if s != nil {
		cuo.SetCode(*s)
	}
	return cuo
}

// SetDisplaySort sets the "display_sort" field.
func (cuo *CountryUpdateOne) SetDisplaySort(i int32) *CountryUpdateOne {
	cuo.mutation.ResetDisplaySort()
	cuo.mutation.SetDisplaySort(i)
	return cuo
}

// SetNillableDisplaySort sets the "display_sort" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableDisplaySort(i *int32) *CountryUpdateOne {
	if i != nil {
		cuo.SetDisplaySort(*i)
	}
	return cuo
}

// AddDisplaySort adds i to the "display_sort" field.
func (cuo *CountryUpdateOne) AddDisplaySort(i int32) *CountryUpdateOne {
	cuo.mutation.AddDisplaySort(i)
	return cuo
}

// ClearDisplaySort clears the value of the "display_sort" field.
func (cuo *CountryUpdateOne) ClearDisplaySort() *CountryUpdateOne {
	cuo.mutation.ClearDisplaySort()
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CountryUpdateOne) SetStatus(ts typex.SimpleStatus) *CountryUpdateOne {
	cuo.mutation.SetStatus(ts)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *CountryUpdateOne) SetNillableStatus(ts *typex.SimpleStatus) *CountryUpdateOne {
	if ts != nil {
		cuo.SetStatus(*ts)
	}
	return cuo
}

// ClearStatus clears the value of the "status" field.
func (cuo *CountryUpdateOne) ClearStatus() *CountryUpdateOne {
	cuo.mutation.ClearStatus()
	return cuo
}

// AddRegionIDs adds the "regions" edge to the Region entity by IDs.
func (cuo *CountryUpdateOne) AddRegionIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.AddRegionIDs(ids...)
	return cuo
}

// AddRegions adds the "regions" edges to the Region entity.
func (cuo *CountryUpdateOne) AddRegions(r ...*Region) *CountryUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.AddRegionIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cuo *CountryUpdateOne) Mutation() *CountryMutation {
	return cuo.mutation
}

// ClearRegions clears all "regions" edges to the Region entity.
func (cuo *CountryUpdateOne) ClearRegions() *CountryUpdateOne {
	cuo.mutation.ClearRegions()
	return cuo
}

// RemoveRegionIDs removes the "regions" edge to Region entities by IDs.
func (cuo *CountryUpdateOne) RemoveRegionIDs(ids ...int) *CountryUpdateOne {
	cuo.mutation.RemoveRegionIDs(ids...)
	return cuo
}

// RemoveRegions removes "regions" edges to Region entities.
func (cuo *CountryUpdateOne) RemoveRegions(r ...*Region) *CountryUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.RemoveRegionIDs(ids...)
}

// Where appends a list predicates to the CountryUpdate builder.
func (cuo *CountryUpdateOne) Where(ps ...predicate.Country) *CountryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CountryUpdateOne) Select(field string, fields ...string) *CountryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Country entity.
func (cuo *CountryUpdateOne) Save(ctx context.Context) (*Country, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CountryUpdateOne) SaveX(ctx context.Context) *Country {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CountryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CountryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CountryUpdateOne) check() error {
	if v, ok := cuo.mutation.Code(); ok {
		if err := country.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Country.code": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Status(); ok {
		if err := country.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Country.status": %w`, err)}
		}
	}
	return nil
}

func (cuo *CountryUpdateOne) sqlSave(ctx context.Context) (_node *Country, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(country.Table, country.Columns, sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Country.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, country.FieldID)
		for _, f := range fields {
			if !country.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != country.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedBy(); ok {
		_spec.SetField(country.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(country.FieldUpdatedBy, field.TypeInt, value)
	}
	if cuo.mutation.UpdatedByCleared() {
		_spec.ClearField(country.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(country.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(country.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(country.FieldName, field.TypeString, value)
	}
	if cuo.mutation.NameCleared() {
		_spec.ClearField(country.FieldName, field.TypeString)
	}
	if value, ok := cuo.mutation.NameEn(); ok {
		_spec.SetField(country.FieldNameEn, field.TypeString, value)
	}
	if cuo.mutation.NameEnCleared() {
		_spec.ClearField(country.FieldNameEn, field.TypeString)
	}
	if value, ok := cuo.mutation.Code(); ok {
		_spec.SetField(country.FieldCode, field.TypeString, value)
	}
	if value, ok := cuo.mutation.DisplaySort(); ok {
		_spec.SetField(country.FieldDisplaySort, field.TypeInt32, value)
	}
	if value, ok := cuo.mutation.AddedDisplaySort(); ok {
		_spec.AddField(country.FieldDisplaySort, field.TypeInt32, value)
	}
	if cuo.mutation.DisplaySortCleared() {
		_spec.ClearField(country.FieldDisplaySort, field.TypeInt32)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(country.FieldStatus, field.TypeEnum, value)
	}
	if cuo.mutation.StatusCleared() {
		_spec.ClearField(country.FieldStatus, field.TypeEnum)
	}
	if cuo.mutation.RegionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.RegionsTable,
			Columns: []string{country.RegionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(region.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedRegionsIDs(); len(nodes) > 0 && !cuo.mutation.RegionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.RegionsTable,
			Columns: []string{country.RegionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(region.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RegionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   country.RegionsTable,
			Columns: []string{country.RegionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(region.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Country{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{country.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
