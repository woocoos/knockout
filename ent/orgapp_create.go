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
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgapp"
)

// OrgAppCreate is the builder for creating a OrgApp entity.
type OrgAppCreate struct {
	config
	mutation *OrgAppMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (oac *OrgAppCreate) SetCreatedBy(i int) *OrgAppCreate {
	oac.mutation.SetCreatedBy(i)
	return oac
}

// SetCreatedAt sets the "created_at" field.
func (oac *OrgAppCreate) SetCreatedAt(t time.Time) *OrgAppCreate {
	oac.mutation.SetCreatedAt(t)
	return oac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oac *OrgAppCreate) SetNillableCreatedAt(t *time.Time) *OrgAppCreate {
	if t != nil {
		oac.SetCreatedAt(*t)
	}
	return oac
}

// SetUpdatedBy sets the "updated_by" field.
func (oac *OrgAppCreate) SetUpdatedBy(i int) *OrgAppCreate {
	oac.mutation.SetUpdatedBy(i)
	return oac
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oac *OrgAppCreate) SetNillableUpdatedBy(i *int) *OrgAppCreate {
	if i != nil {
		oac.SetUpdatedBy(*i)
	}
	return oac
}

// SetUpdatedAt sets the "updated_at" field.
func (oac *OrgAppCreate) SetUpdatedAt(t time.Time) *OrgAppCreate {
	oac.mutation.SetUpdatedAt(t)
	return oac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oac *OrgAppCreate) SetNillableUpdatedAt(t *time.Time) *OrgAppCreate {
	if t != nil {
		oac.SetUpdatedAt(*t)
	}
	return oac
}

// SetOrgID sets the "org_id" field.
func (oac *OrgAppCreate) SetOrgID(i int) *OrgAppCreate {
	oac.mutation.SetOrgID(i)
	return oac
}

// SetAppID sets the "app_id" field.
func (oac *OrgAppCreate) SetAppID(i int) *OrgAppCreate {
	oac.mutation.SetAppID(i)
	return oac
}

// SetID sets the "id" field.
func (oac *OrgAppCreate) SetID(i int) *OrgAppCreate {
	oac.mutation.SetID(i)
	return oac
}

// SetApp sets the "app" edge to the App entity.
func (oac *OrgAppCreate) SetApp(a *App) *OrgAppCreate {
	return oac.SetAppID(a.ID)
}

// SetOrg sets the "org" edge to the Org entity.
func (oac *OrgAppCreate) SetOrg(o *Org) *OrgAppCreate {
	return oac.SetOrgID(o.ID)
}

// Mutation returns the OrgAppMutation object of the builder.
func (oac *OrgAppCreate) Mutation() *OrgAppMutation {
	return oac.mutation
}

// Save creates the OrgApp in the database.
func (oac *OrgAppCreate) Save(ctx context.Context) (*OrgApp, error) {
	if err := oac.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, oac.sqlSave, oac.mutation, oac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oac *OrgAppCreate) SaveX(ctx context.Context) *OrgApp {
	v, err := oac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oac *OrgAppCreate) Exec(ctx context.Context) error {
	_, err := oac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oac *OrgAppCreate) ExecX(ctx context.Context) {
	if err := oac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oac *OrgAppCreate) defaults() error {
	if _, ok := oac.mutation.CreatedAt(); !ok {
		if orgapp.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized orgapp.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := orgapp.DefaultCreatedAt()
		oac.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (oac *OrgAppCreate) check() error {
	if _, ok := oac.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "OrgApp.created_by"`)}
	}
	if _, ok := oac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrgApp.created_at"`)}
	}
	if _, ok := oac.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "OrgApp.org_id"`)}
	}
	if _, ok := oac.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "OrgApp.app_id"`)}
	}
	if len(oac.mutation.AppIDs()) == 0 {
		return &ValidationError{Name: "app", err: errors.New(`ent: missing required edge "OrgApp.app"`)}
	}
	if len(oac.mutation.OrgIDs()) == 0 {
		return &ValidationError{Name: "org", err: errors.New(`ent: missing required edge "OrgApp.org"`)}
	}
	return nil
}

func (oac *OrgAppCreate) sqlSave(ctx context.Context) (*OrgApp, error) {
	if err := oac.check(); err != nil {
		return nil, err
	}
	_node, _spec := oac.createSpec()
	if err := sqlgraph.CreateNode(ctx, oac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	oac.mutation.id = &_node.ID
	oac.mutation.done = true
	return _node, nil
}

func (oac *OrgAppCreate) createSpec() (*OrgApp, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgApp{config: oac.config}
		_spec = sqlgraph.NewCreateSpec(orgapp.Table, sqlgraph.NewFieldSpec(orgapp.FieldID, field.TypeInt))
	)
	_spec.OnConflict = oac.conflict
	if id, ok := oac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oac.mutation.CreatedBy(); ok {
		_spec.SetField(orgapp.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := oac.mutation.CreatedAt(); ok {
		_spec.SetField(orgapp.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oac.mutation.UpdatedBy(); ok {
		_spec.SetField(orgapp.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := oac.mutation.UpdatedAt(); ok {
		_spec.SetField(orgapp.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := oac.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgapp.AppTable,
			Columns: []string{orgapp.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(app.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AppID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oac.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgapp.OrgTable,
			Columns: []string{orgapp.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(org.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrgID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrgApp.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgAppUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (oac *OrgAppCreate) OnConflict(opts ...sql.ConflictOption) *OrgAppUpsertOne {
	oac.conflict = opts
	return &OrgAppUpsertOne{
		create: oac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgApp.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oac *OrgAppCreate) OnConflictColumns(columns ...string) *OrgAppUpsertOne {
	oac.conflict = append(oac.conflict, sql.ConflictColumns(columns...))
	return &OrgAppUpsertOne{
		create: oac,
	}
}

type (
	// OrgAppUpsertOne is the builder for "upsert"-ing
	//  one OrgApp node.
	OrgAppUpsertOne struct {
		create *OrgAppCreate
	}

	// OrgAppUpsert is the "OnConflict" setter.
	OrgAppUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *OrgAppUpsert) SetUpdatedBy(v int) *OrgAppUpsert {
	u.Set(orgapp.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OrgAppUpsert) UpdateUpdatedBy() *OrgAppUpsert {
	u.SetExcluded(orgapp.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *OrgAppUpsert) AddUpdatedBy(v int) *OrgAppUpsert {
	u.Add(orgapp.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OrgAppUpsert) ClearUpdatedBy() *OrgAppUpsert {
	u.SetNull(orgapp.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrgAppUpsert) SetUpdatedAt(v time.Time) *OrgAppUpsert {
	u.Set(orgapp.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrgAppUpsert) UpdateUpdatedAt() *OrgAppUpsert {
	u.SetExcluded(orgapp.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OrgAppUpsert) ClearUpdatedAt() *OrgAppUpsert {
	u.SetNull(orgapp.FieldUpdatedAt)
	return u
}

// SetOrgID sets the "org_id" field.
func (u *OrgAppUpsert) SetOrgID(v int) *OrgAppUpsert {
	u.Set(orgapp.FieldOrgID, v)
	return u
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgAppUpsert) UpdateOrgID() *OrgAppUpsert {
	u.SetExcluded(orgapp.FieldOrgID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *OrgAppUpsert) SetAppID(v int) *OrgAppUpsert {
	u.Set(orgapp.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *OrgAppUpsert) UpdateAppID() *OrgAppUpsert {
	u.SetExcluded(orgapp.FieldAppID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrgApp.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orgapp.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrgAppUpsertOne) UpdateNewValues() *OrgAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(orgapp.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(orgapp.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(orgapp.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrgApp.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrgAppUpsertOne) Ignore() *OrgAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgAppUpsertOne) DoNothing() *OrgAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgAppCreate.OnConflict
// documentation for more info.
func (u *OrgAppUpsertOne) Update(set func(*OrgAppUpsert)) *OrgAppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgAppUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *OrgAppUpsertOne) SetUpdatedBy(v int) *OrgAppUpsertOne {
	return u.Update(func(s *OrgAppUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *OrgAppUpsertOne) AddUpdatedBy(v int) *OrgAppUpsertOne {
	return u.Update(func(s *OrgAppUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OrgAppUpsertOne) UpdateUpdatedBy() *OrgAppUpsertOne {
	return u.Update(func(s *OrgAppUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OrgAppUpsertOne) ClearUpdatedBy() *OrgAppUpsertOne {
	return u.Update(func(s *OrgAppUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrgAppUpsertOne) SetUpdatedAt(v time.Time) *OrgAppUpsertOne {
	return u.Update(func(s *OrgAppUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrgAppUpsertOne) UpdateUpdatedAt() *OrgAppUpsertOne {
	return u.Update(func(s *OrgAppUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OrgAppUpsertOne) ClearUpdatedAt() *OrgAppUpsertOne {
	return u.Update(func(s *OrgAppUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetOrgID sets the "org_id" field.
func (u *OrgAppUpsertOne) SetOrgID(v int) *OrgAppUpsertOne {
	return u.Update(func(s *OrgAppUpsert) {
		s.SetOrgID(v)
	})
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgAppUpsertOne) UpdateOrgID() *OrgAppUpsertOne {
	return u.Update(func(s *OrgAppUpsert) {
		s.UpdateOrgID()
	})
}

// SetAppID sets the "app_id" field.
func (u *OrgAppUpsertOne) SetAppID(v int) *OrgAppUpsertOne {
	return u.Update(func(s *OrgAppUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *OrgAppUpsertOne) UpdateAppID() *OrgAppUpsertOne {
	return u.Update(func(s *OrgAppUpsert) {
		s.UpdateAppID()
	})
}

// Exec executes the query.
func (u *OrgAppUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrgAppCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgAppUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrgAppUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrgAppUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrgAppCreateBulk is the builder for creating many OrgApp entities in bulk.
type OrgAppCreateBulk struct {
	config
	err      error
	builders []*OrgAppCreate
	conflict []sql.ConflictOption
}

// Save creates the OrgApp entities in the database.
func (oacb *OrgAppCreateBulk) Save(ctx context.Context) ([]*OrgApp, error) {
	if oacb.err != nil {
		return nil, oacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oacb.builders))
	nodes := make([]*OrgApp, len(oacb.builders))
	mutators := make([]Mutator, len(oacb.builders))
	for i := range oacb.builders {
		func(i int, root context.Context) {
			builder := oacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgAppMutation)
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
					_, err = mutators[i+1].Mutate(root, oacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = oacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oacb *OrgAppCreateBulk) SaveX(ctx context.Context) []*OrgApp {
	v, err := oacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oacb *OrgAppCreateBulk) Exec(ctx context.Context) error {
	_, err := oacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oacb *OrgAppCreateBulk) ExecX(ctx context.Context) {
	if err := oacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrgApp.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgAppUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (oacb *OrgAppCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrgAppUpsertBulk {
	oacb.conflict = opts
	return &OrgAppUpsertBulk{
		create: oacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgApp.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oacb *OrgAppCreateBulk) OnConflictColumns(columns ...string) *OrgAppUpsertBulk {
	oacb.conflict = append(oacb.conflict, sql.ConflictColumns(columns...))
	return &OrgAppUpsertBulk{
		create: oacb,
	}
}

// OrgAppUpsertBulk is the builder for "upsert"-ing
// a bulk of OrgApp nodes.
type OrgAppUpsertBulk struct {
	create *OrgAppCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrgApp.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orgapp.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrgAppUpsertBulk) UpdateNewValues() *OrgAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(orgapp.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(orgapp.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(orgapp.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrgApp.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrgAppUpsertBulk) Ignore() *OrgAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgAppUpsertBulk) DoNothing() *OrgAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgAppCreateBulk.OnConflict
// documentation for more info.
func (u *OrgAppUpsertBulk) Update(set func(*OrgAppUpsert)) *OrgAppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgAppUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *OrgAppUpsertBulk) SetUpdatedBy(v int) *OrgAppUpsertBulk {
	return u.Update(func(s *OrgAppUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *OrgAppUpsertBulk) AddUpdatedBy(v int) *OrgAppUpsertBulk {
	return u.Update(func(s *OrgAppUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OrgAppUpsertBulk) UpdateUpdatedBy() *OrgAppUpsertBulk {
	return u.Update(func(s *OrgAppUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OrgAppUpsertBulk) ClearUpdatedBy() *OrgAppUpsertBulk {
	return u.Update(func(s *OrgAppUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrgAppUpsertBulk) SetUpdatedAt(v time.Time) *OrgAppUpsertBulk {
	return u.Update(func(s *OrgAppUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrgAppUpsertBulk) UpdateUpdatedAt() *OrgAppUpsertBulk {
	return u.Update(func(s *OrgAppUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OrgAppUpsertBulk) ClearUpdatedAt() *OrgAppUpsertBulk {
	return u.Update(func(s *OrgAppUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetOrgID sets the "org_id" field.
func (u *OrgAppUpsertBulk) SetOrgID(v int) *OrgAppUpsertBulk {
	return u.Update(func(s *OrgAppUpsert) {
		s.SetOrgID(v)
	})
}

// UpdateOrgID sets the "org_id" field to the value that was provided on create.
func (u *OrgAppUpsertBulk) UpdateOrgID() *OrgAppUpsertBulk {
	return u.Update(func(s *OrgAppUpsert) {
		s.UpdateOrgID()
	})
}

// SetAppID sets the "app_id" field.
func (u *OrgAppUpsertBulk) SetAppID(v int) *OrgAppUpsertBulk {
	return u.Update(func(s *OrgAppUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *OrgAppUpsertBulk) UpdateAppID() *OrgAppUpsertBulk {
	return u.Update(func(s *OrgAppUpsert) {
		s.UpdateAppID()
	})
}

// Exec executes the query.
func (u *OrgAppUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrgAppCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrgAppCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgAppUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
