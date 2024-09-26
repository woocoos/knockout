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
	"github.com/woocoos/knockout/ent/region"
)

// CountryCreate is the builder for creating a Country entity.
type CountryCreate struct {
	config
	mutation *CountryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (cc *CountryCreate) SetCreatedBy(i int) *CountryCreate {
	cc.mutation.SetCreatedBy(i)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CountryCreate) SetCreatedAt(t time.Time) *CountryCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CountryCreate) SetNillableCreatedAt(t *time.Time) *CountryCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedBy sets the "updated_by" field.
func (cc *CountryCreate) SetUpdatedBy(i int) *CountryCreate {
	cc.mutation.SetUpdatedBy(i)
	return cc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cc *CountryCreate) SetNillableUpdatedBy(i *int) *CountryCreate {
	if i != nil {
		cc.SetUpdatedBy(*i)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CountryCreate) SetUpdatedAt(t time.Time) *CountryCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CountryCreate) SetNillableUpdatedAt(t *time.Time) *CountryCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CountryCreate) SetName(s string) *CountryCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *CountryCreate) SetNillableName(s *string) *CountryCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetNameEn sets the "name_en" field.
func (cc *CountryCreate) SetNameEn(s string) *CountryCreate {
	cc.mutation.SetNameEn(s)
	return cc
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (cc *CountryCreate) SetNillableNameEn(s *string) *CountryCreate {
	if s != nil {
		cc.SetNameEn(*s)
	}
	return cc
}

// SetCode sets the "code" field.
func (cc *CountryCreate) SetCode(s string) *CountryCreate {
	cc.mutation.SetCode(s)
	return cc
}

// SetDisplaySort sets the "display_sort" field.
func (cc *CountryCreate) SetDisplaySort(i int32) *CountryCreate {
	cc.mutation.SetDisplaySort(i)
	return cc
}

// SetNillableDisplaySort sets the "display_sort" field if the given value is not nil.
func (cc *CountryCreate) SetNillableDisplaySort(i *int32) *CountryCreate {
	if i != nil {
		cc.SetDisplaySort(*i)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CountryCreate) SetStatus(ts typex.SimpleStatus) *CountryCreate {
	cc.mutation.SetStatus(ts)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CountryCreate) SetNillableStatus(ts *typex.SimpleStatus) *CountryCreate {
	if ts != nil {
		cc.SetStatus(*ts)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CountryCreate) SetID(i int) *CountryCreate {
	cc.mutation.SetID(i)
	return cc
}

// AddRegionIDs adds the "regions" edge to the Region entity by IDs.
func (cc *CountryCreate) AddRegionIDs(ids ...int) *CountryCreate {
	cc.mutation.AddRegionIDs(ids...)
	return cc
}

// AddRegions adds the "regions" edges to the Region entity.
func (cc *CountryCreate) AddRegions(r ...*Region) *CountryCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cc.AddRegionIDs(ids...)
}

// Mutation returns the CountryMutation object of the builder.
func (cc *CountryCreate) Mutation() *CountryMutation {
	return cc.mutation
}

// Save creates the Country in the database.
func (cc *CountryCreate) Save(ctx context.Context) (*Country, error) {
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CountryCreate) SaveX(ctx context.Context) *Country {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CountryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CountryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CountryCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if country.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized country.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := country.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := country.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CountryCreate) check() error {
	if _, ok := cc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Country.created_by"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Country.created_at"`)}
	}
	if _, ok := cc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Country.code"`)}
	}
	if v, ok := cc.mutation.Code(); ok {
		if err := country.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Country.code": %w`, err)}
		}
	}
	if v, ok := cc.mutation.Status(); ok {
		if err := country.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Country.status": %w`, err)}
		}
	}
	return nil
}

func (cc *CountryCreate) sqlSave(ctx context.Context) (*Country, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CountryCreate) createSpec() (*Country, *sqlgraph.CreateSpec) {
	var (
		_node = &Country{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(country.Table, sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt))
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedBy(); ok {
		_spec.SetField(country.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(country.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedBy(); ok {
		_spec.SetField(country.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(country.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(country.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.NameEn(); ok {
		_spec.SetField(country.FieldNameEn, field.TypeString, value)
		_node.NameEn = value
	}
	if value, ok := cc.mutation.Code(); ok {
		_spec.SetField(country.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := cc.mutation.DisplaySort(); ok {
		_spec.SetField(country.FieldDisplaySort, field.TypeInt32, value)
		_node.DisplaySort = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(country.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := cc.mutation.RegionsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Country.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CountryUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (cc *CountryCreate) OnConflict(opts ...sql.ConflictOption) *CountryUpsertOne {
	cc.conflict = opts
	return &CountryUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Country.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cc *CountryCreate) OnConflictColumns(columns ...string) *CountryUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CountryUpsertOne{
		create: cc,
	}
}

type (
	// CountryUpsertOne is the builder for "upsert"-ing
	//  one Country node.
	CountryUpsertOne struct {
		create *CountryCreate
	}

	// CountryUpsert is the "OnConflict" setter.
	CountryUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *CountryUpsert) SetUpdatedBy(v int) *CountryUpsert {
	u.Set(country.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *CountryUpsert) UpdateUpdatedBy() *CountryUpsert {
	u.SetExcluded(country.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *CountryUpsert) AddUpdatedBy(v int) *CountryUpsert {
	u.Add(country.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *CountryUpsert) ClearUpdatedBy() *CountryUpsert {
	u.SetNull(country.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CountryUpsert) SetUpdatedAt(v time.Time) *CountryUpsert {
	u.Set(country.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CountryUpsert) UpdateUpdatedAt() *CountryUpsert {
	u.SetExcluded(country.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *CountryUpsert) ClearUpdatedAt() *CountryUpsert {
	u.SetNull(country.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *CountryUpsert) SetName(v string) *CountryUpsert {
	u.Set(country.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CountryUpsert) UpdateName() *CountryUpsert {
	u.SetExcluded(country.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *CountryUpsert) ClearName() *CountryUpsert {
	u.SetNull(country.FieldName)
	return u
}

// SetNameEn sets the "name_en" field.
func (u *CountryUpsert) SetNameEn(v string) *CountryUpsert {
	u.Set(country.FieldNameEn, v)
	return u
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *CountryUpsert) UpdateNameEn() *CountryUpsert {
	u.SetExcluded(country.FieldNameEn)
	return u
}

// ClearNameEn clears the value of the "name_en" field.
func (u *CountryUpsert) ClearNameEn() *CountryUpsert {
	u.SetNull(country.FieldNameEn)
	return u
}

// SetCode sets the "code" field.
func (u *CountryUpsert) SetCode(v string) *CountryUpsert {
	u.Set(country.FieldCode, v)
	return u
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *CountryUpsert) UpdateCode() *CountryUpsert {
	u.SetExcluded(country.FieldCode)
	return u
}

// SetDisplaySort sets the "display_sort" field.
func (u *CountryUpsert) SetDisplaySort(v int32) *CountryUpsert {
	u.Set(country.FieldDisplaySort, v)
	return u
}

// UpdateDisplaySort sets the "display_sort" field to the value that was provided on create.
func (u *CountryUpsert) UpdateDisplaySort() *CountryUpsert {
	u.SetExcluded(country.FieldDisplaySort)
	return u
}

// AddDisplaySort adds v to the "display_sort" field.
func (u *CountryUpsert) AddDisplaySort(v int32) *CountryUpsert {
	u.Add(country.FieldDisplaySort, v)
	return u
}

// ClearDisplaySort clears the value of the "display_sort" field.
func (u *CountryUpsert) ClearDisplaySort() *CountryUpsert {
	u.SetNull(country.FieldDisplaySort)
	return u
}

// SetStatus sets the "status" field.
func (u *CountryUpsert) SetStatus(v typex.SimpleStatus) *CountryUpsert {
	u.Set(country.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CountryUpsert) UpdateStatus() *CountryUpsert {
	u.SetExcluded(country.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *CountryUpsert) ClearStatus() *CountryUpsert {
	u.SetNull(country.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Country.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(country.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CountryUpsertOne) UpdateNewValues() *CountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(country.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(country.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(country.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Country.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CountryUpsertOne) Ignore() *CountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CountryUpsertOne) DoNothing() *CountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CountryCreate.OnConflict
// documentation for more info.
func (u *CountryUpsertOne) Update(set func(*CountryUpsert)) *CountryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CountryUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *CountryUpsertOne) SetUpdatedBy(v int) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *CountryUpsertOne) AddUpdatedBy(v int) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateUpdatedBy() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *CountryUpsertOne) ClearUpdatedBy() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CountryUpsertOne) SetUpdatedAt(v time.Time) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateUpdatedAt() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *CountryUpsertOne) ClearUpdatedAt() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *CountryUpsertOne) SetName(v string) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateName() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *CountryUpsertOne) ClearName() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.ClearName()
	})
}

// SetNameEn sets the "name_en" field.
func (u *CountryUpsertOne) SetNameEn(v string) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetNameEn(v)
	})
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateNameEn() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateNameEn()
	})
}

// ClearNameEn clears the value of the "name_en" field.
func (u *CountryUpsertOne) ClearNameEn() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.ClearNameEn()
	})
}

// SetCode sets the "code" field.
func (u *CountryUpsertOne) SetCode(v string) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateCode() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateCode()
	})
}

// SetDisplaySort sets the "display_sort" field.
func (u *CountryUpsertOne) SetDisplaySort(v int32) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetDisplaySort(v)
	})
}

// AddDisplaySort adds v to the "display_sort" field.
func (u *CountryUpsertOne) AddDisplaySort(v int32) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.AddDisplaySort(v)
	})
}

// UpdateDisplaySort sets the "display_sort" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateDisplaySort() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateDisplaySort()
	})
}

// ClearDisplaySort clears the value of the "display_sort" field.
func (u *CountryUpsertOne) ClearDisplaySort() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.ClearDisplaySort()
	})
}

// SetStatus sets the "status" field.
func (u *CountryUpsertOne) SetStatus(v typex.SimpleStatus) *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CountryUpsertOne) UpdateStatus() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *CountryUpsertOne) ClearStatus() *CountryUpsertOne {
	return u.Update(func(s *CountryUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *CountryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CountryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CountryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CountryUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CountryUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CountryCreateBulk is the builder for creating many Country entities in bulk.
type CountryCreateBulk struct {
	config
	err      error
	builders []*CountryCreate
	conflict []sql.ConflictOption
}

// Save creates the Country entities in the database.
func (ccb *CountryCreateBulk) Save(ctx context.Context) ([]*Country, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Country, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CountryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CountryCreateBulk) SaveX(ctx context.Context) []*Country {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CountryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CountryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Country.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CountryUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (ccb *CountryCreateBulk) OnConflict(opts ...sql.ConflictOption) *CountryUpsertBulk {
	ccb.conflict = opts
	return &CountryUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Country.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ccb *CountryCreateBulk) OnConflictColumns(columns ...string) *CountryUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CountryUpsertBulk{
		create: ccb,
	}
}

// CountryUpsertBulk is the builder for "upsert"-ing
// a bulk of Country nodes.
type CountryUpsertBulk struct {
	create *CountryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Country.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(country.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CountryUpsertBulk) UpdateNewValues() *CountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(country.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(country.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(country.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Country.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CountryUpsertBulk) Ignore() *CountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CountryUpsertBulk) DoNothing() *CountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CountryCreateBulk.OnConflict
// documentation for more info.
func (u *CountryUpsertBulk) Update(set func(*CountryUpsert)) *CountryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CountryUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *CountryUpsertBulk) SetUpdatedBy(v int) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *CountryUpsertBulk) AddUpdatedBy(v int) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateUpdatedBy() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *CountryUpsertBulk) ClearUpdatedBy() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CountryUpsertBulk) SetUpdatedAt(v time.Time) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateUpdatedAt() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *CountryUpsertBulk) ClearUpdatedAt() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *CountryUpsertBulk) SetName(v string) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateName() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *CountryUpsertBulk) ClearName() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.ClearName()
	})
}

// SetNameEn sets the "name_en" field.
func (u *CountryUpsertBulk) SetNameEn(v string) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetNameEn(v)
	})
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateNameEn() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateNameEn()
	})
}

// ClearNameEn clears the value of the "name_en" field.
func (u *CountryUpsertBulk) ClearNameEn() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.ClearNameEn()
	})
}

// SetCode sets the "code" field.
func (u *CountryUpsertBulk) SetCode(v string) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateCode() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateCode()
	})
}

// SetDisplaySort sets the "display_sort" field.
func (u *CountryUpsertBulk) SetDisplaySort(v int32) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetDisplaySort(v)
	})
}

// AddDisplaySort adds v to the "display_sort" field.
func (u *CountryUpsertBulk) AddDisplaySort(v int32) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.AddDisplaySort(v)
	})
}

// UpdateDisplaySort sets the "display_sort" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateDisplaySort() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateDisplaySort()
	})
}

// ClearDisplaySort clears the value of the "display_sort" field.
func (u *CountryUpsertBulk) ClearDisplaySort() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.ClearDisplaySort()
	})
}

// SetStatus sets the "status" field.
func (u *CountryUpsertBulk) SetStatus(v typex.SimpleStatus) *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *CountryUpsertBulk) UpdateStatus() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *CountryUpsertBulk) ClearStatus() *CountryUpsertBulk {
	return u.Update(func(s *CountryUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *CountryUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CountryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CountryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CountryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}