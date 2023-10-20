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
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orguserpreference"
	"github.com/woocoos/knockout/ent/user"
)

// OrgUserPreferenceCreate is the builder for creating a OrgUserPreference entity.
type OrgUserPreferenceCreate struct {
	config
	mutation *OrgUserPreferenceMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (oupc *OrgUserPreferenceCreate) SetCreatedBy(i int) *OrgUserPreferenceCreate {
	oupc.mutation.SetCreatedBy(i)
	return oupc
}

// SetCreatedAt sets the "created_at" field.
func (oupc *OrgUserPreferenceCreate) SetCreatedAt(t time.Time) *OrgUserPreferenceCreate {
	oupc.mutation.SetCreatedAt(t)
	return oupc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oupc *OrgUserPreferenceCreate) SetNillableCreatedAt(t *time.Time) *OrgUserPreferenceCreate {
	if t != nil {
		oupc.SetCreatedAt(*t)
	}
	return oupc
}

// SetUpdatedBy sets the "updated_by" field.
func (oupc *OrgUserPreferenceCreate) SetUpdatedBy(i int) *OrgUserPreferenceCreate {
	oupc.mutation.SetUpdatedBy(i)
	return oupc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oupc *OrgUserPreferenceCreate) SetNillableUpdatedBy(i *int) *OrgUserPreferenceCreate {
	if i != nil {
		oupc.SetUpdatedBy(*i)
	}
	return oupc
}

// SetUpdatedAt sets the "updated_at" field.
func (oupc *OrgUserPreferenceCreate) SetUpdatedAt(t time.Time) *OrgUserPreferenceCreate {
	oupc.mutation.SetUpdatedAt(t)
	return oupc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oupc *OrgUserPreferenceCreate) SetNillableUpdatedAt(t *time.Time) *OrgUserPreferenceCreate {
	if t != nil {
		oupc.SetUpdatedAt(*t)
	}
	return oupc
}

// SetUserID sets the "user_id" field.
func (oupc *OrgUserPreferenceCreate) SetUserID(i int) *OrgUserPreferenceCreate {
	oupc.mutation.SetUserID(i)
	return oupc
}

// SetOrgID sets the "org_id" field.
func (oupc *OrgUserPreferenceCreate) SetOrgID(i int) *OrgUserPreferenceCreate {
	oupc.mutation.SetOrgID(i)
	return oupc
}

// SetMenuFavorite sets the "menu_favorite" field.
func (oupc *OrgUserPreferenceCreate) SetMenuFavorite(i []int) *OrgUserPreferenceCreate {
	oupc.mutation.SetMenuFavorite(i)
	return oupc
}

// SetMenuRecent sets the "menu_recent" field.
func (oupc *OrgUserPreferenceCreate) SetMenuRecent(i []int) *OrgUserPreferenceCreate {
	oupc.mutation.SetMenuRecent(i)
	return oupc
}

// SetID sets the "id" field.
func (oupc *OrgUserPreferenceCreate) SetID(i int) *OrgUserPreferenceCreate {
	oupc.mutation.SetID(i)
	return oupc
}

// SetUser sets the "user" edge to the User entity.
func (oupc *OrgUserPreferenceCreate) SetUser(u *User) *OrgUserPreferenceCreate {
	return oupc.SetUserID(u.ID)
}

// SetOrg sets the "org" edge to the Org entity.
func (oupc *OrgUserPreferenceCreate) SetOrg(o *Org) *OrgUserPreferenceCreate {
	return oupc.SetOrgID(o.ID)
}

// Mutation returns the OrgUserPreferenceMutation object of the builder.
func (oupc *OrgUserPreferenceCreate) Mutation() *OrgUserPreferenceMutation {
	return oupc.mutation
}

// Save creates the OrgUserPreference in the database.
func (oupc *OrgUserPreferenceCreate) Save(ctx context.Context) (*OrgUserPreference, error) {
	if err := oupc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, oupc.sqlSave, oupc.mutation, oupc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oupc *OrgUserPreferenceCreate) SaveX(ctx context.Context) *OrgUserPreference {
	v, err := oupc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oupc *OrgUserPreferenceCreate) Exec(ctx context.Context) error {
	_, err := oupc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oupc *OrgUserPreferenceCreate) ExecX(ctx context.Context) {
	if err := oupc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oupc *OrgUserPreferenceCreate) defaults() error {
	if _, ok := oupc.mutation.CreatedAt(); !ok {
		if orguserpreference.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized orguserpreference.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := orguserpreference.DefaultCreatedAt()
		oupc.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (oupc *OrgUserPreferenceCreate) check() error {
	if _, ok := oupc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "OrgUserPreference.created_by"`)}
	}
	if _, ok := oupc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrgUserPreference.created_at"`)}
	}
	if _, ok := oupc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "OrgUserPreference.user_id"`)}
	}
	if _, ok := oupc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "OrgUserPreference.org_id"`)}
	}
	if _, ok := oupc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "OrgUserPreference.user"`)}
	}
	if _, ok := oupc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org", err: errors.New(`ent: missing required edge "OrgUserPreference.org"`)}
	}
	return nil
}

func (oupc *OrgUserPreferenceCreate) sqlSave(ctx context.Context) (*OrgUserPreference, error) {
	if err := oupc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oupc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oupc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	oupc.mutation.id = &_node.ID
	oupc.mutation.done = true
	return _node, nil
}

func (oupc *OrgUserPreferenceCreate) createSpec() (*OrgUserPreference, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgUserPreference{config: oupc.config}
		_spec = sqlgraph.NewCreateSpec(orguserpreference.Table, sqlgraph.NewFieldSpec(orguserpreference.FieldID, field.TypeInt))
	)
	_spec.OnConflict = oupc.conflict
	if id, ok := oupc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oupc.mutation.CreatedBy(); ok {
		_spec.SetField(orguserpreference.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := oupc.mutation.CreatedAt(); ok {
		_spec.SetField(orguserpreference.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oupc.mutation.UpdatedBy(); ok {
		_spec.SetField(orguserpreference.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := oupc.mutation.UpdatedAt(); ok {
		_spec.SetField(orguserpreference.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oupc.mutation.MenuFavorite(); ok {
		_spec.SetField(orguserpreference.FieldMenuFavorite, field.TypeJSON, value)
		_node.MenuFavorite = value
	}
	if value, ok := oupc.mutation.MenuRecent(); ok {
		_spec.SetField(orguserpreference.FieldMenuRecent, field.TypeJSON, value)
		_node.MenuRecent = value
	}
	if nodes := oupc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguserpreference.UserTable,
			Columns: []string{orguserpreference.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oupc.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguserpreference.OrgTable,
			Columns: []string{orguserpreference.OrgColumn},
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
//	client.OrgUserPreference.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgUserPreferenceUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (oupc *OrgUserPreferenceCreate) OnConflict(opts ...sql.ConflictOption) *OrgUserPreferenceUpsertOne {
	oupc.conflict = opts
	return &OrgUserPreferenceUpsertOne{
		create: oupc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgUserPreference.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oupc *OrgUserPreferenceCreate) OnConflictColumns(columns ...string) *OrgUserPreferenceUpsertOne {
	oupc.conflict = append(oupc.conflict, sql.ConflictColumns(columns...))
	return &OrgUserPreferenceUpsertOne{
		create: oupc,
	}
}

type (
	// OrgUserPreferenceUpsertOne is the builder for "upsert"-ing
	//  one OrgUserPreference node.
	OrgUserPreferenceUpsertOne struct {
		create *OrgUserPreferenceCreate
	}

	// OrgUserPreferenceUpsert is the "OnConflict" setter.
	OrgUserPreferenceUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *OrgUserPreferenceUpsert) SetUpdatedBy(v int) *OrgUserPreferenceUpsert {
	u.Set(orguserpreference.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsert) UpdateUpdatedBy() *OrgUserPreferenceUpsert {
	u.SetExcluded(orguserpreference.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *OrgUserPreferenceUpsert) AddUpdatedBy(v int) *OrgUserPreferenceUpsert {
	u.Add(orguserpreference.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OrgUserPreferenceUpsert) ClearUpdatedBy() *OrgUserPreferenceUpsert {
	u.SetNull(orguserpreference.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrgUserPreferenceUpsert) SetUpdatedAt(v time.Time) *OrgUserPreferenceUpsert {
	u.Set(orguserpreference.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsert) UpdateUpdatedAt() *OrgUserPreferenceUpsert {
	u.SetExcluded(orguserpreference.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OrgUserPreferenceUpsert) ClearUpdatedAt() *OrgUserPreferenceUpsert {
	u.SetNull(orguserpreference.FieldUpdatedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *OrgUserPreferenceUpsert) SetUserID(v int) *OrgUserPreferenceUpsert {
	u.Set(orguserpreference.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsert) UpdateUserID() *OrgUserPreferenceUpsert {
	u.SetExcluded(orguserpreference.FieldUserID)
	return u
}

// SetMenuFavorite sets the "menu_favorite" field.
func (u *OrgUserPreferenceUpsert) SetMenuFavorite(v []int) *OrgUserPreferenceUpsert {
	u.Set(orguserpreference.FieldMenuFavorite, v)
	return u
}

// UpdateMenuFavorite sets the "menu_favorite" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsert) UpdateMenuFavorite() *OrgUserPreferenceUpsert {
	u.SetExcluded(orguserpreference.FieldMenuFavorite)
	return u
}

// ClearMenuFavorite clears the value of the "menu_favorite" field.
func (u *OrgUserPreferenceUpsert) ClearMenuFavorite() *OrgUserPreferenceUpsert {
	u.SetNull(orguserpreference.FieldMenuFavorite)
	return u
}

// SetMenuRecent sets the "menu_recent" field.
func (u *OrgUserPreferenceUpsert) SetMenuRecent(v []int) *OrgUserPreferenceUpsert {
	u.Set(orguserpreference.FieldMenuRecent, v)
	return u
}

// UpdateMenuRecent sets the "menu_recent" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsert) UpdateMenuRecent() *OrgUserPreferenceUpsert {
	u.SetExcluded(orguserpreference.FieldMenuRecent)
	return u
}

// ClearMenuRecent clears the value of the "menu_recent" field.
func (u *OrgUserPreferenceUpsert) ClearMenuRecent() *OrgUserPreferenceUpsert {
	u.SetNull(orguserpreference.FieldMenuRecent)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrgUserPreference.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orguserpreference.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrgUserPreferenceUpsertOne) UpdateNewValues() *OrgUserPreferenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(orguserpreference.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(orguserpreference.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(orguserpreference.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.OrgID(); exists {
			s.SetIgnore(orguserpreference.FieldOrgID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrgUserPreference.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *OrgUserPreferenceUpsertOne) Ignore() *OrgUserPreferenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgUserPreferenceUpsertOne) DoNothing() *OrgUserPreferenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgUserPreferenceCreate.OnConflict
// documentation for more info.
func (u *OrgUserPreferenceUpsertOne) Update(set func(*OrgUserPreferenceUpsert)) *OrgUserPreferenceUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgUserPreferenceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *OrgUserPreferenceUpsertOne) SetUpdatedBy(v int) *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *OrgUserPreferenceUpsertOne) AddUpdatedBy(v int) *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsertOne) UpdateUpdatedBy() *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OrgUserPreferenceUpsertOne) ClearUpdatedBy() *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrgUserPreferenceUpsertOne) SetUpdatedAt(v time.Time) *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsertOne) UpdateUpdatedAt() *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OrgUserPreferenceUpsertOne) ClearUpdatedAt() *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrgUserPreferenceUpsertOne) SetUserID(v int) *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsertOne) UpdateUserID() *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.UpdateUserID()
	})
}

// SetMenuFavorite sets the "menu_favorite" field.
func (u *OrgUserPreferenceUpsertOne) SetMenuFavorite(v []int) *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.SetMenuFavorite(v)
	})
}

// UpdateMenuFavorite sets the "menu_favorite" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsertOne) UpdateMenuFavorite() *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.UpdateMenuFavorite()
	})
}

// ClearMenuFavorite clears the value of the "menu_favorite" field.
func (u *OrgUserPreferenceUpsertOne) ClearMenuFavorite() *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.ClearMenuFavorite()
	})
}

// SetMenuRecent sets the "menu_recent" field.
func (u *OrgUserPreferenceUpsertOne) SetMenuRecent(v []int) *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.SetMenuRecent(v)
	})
}

// UpdateMenuRecent sets the "menu_recent" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsertOne) UpdateMenuRecent() *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.UpdateMenuRecent()
	})
}

// ClearMenuRecent clears the value of the "menu_recent" field.
func (u *OrgUserPreferenceUpsertOne) ClearMenuRecent() *OrgUserPreferenceUpsertOne {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.ClearMenuRecent()
	})
}

// Exec executes the query.
func (u *OrgUserPreferenceUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrgUserPreferenceCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgUserPreferenceUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrgUserPreferenceUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrgUserPreferenceUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrgUserPreferenceCreateBulk is the builder for creating many OrgUserPreference entities in bulk.
type OrgUserPreferenceCreateBulk struct {
	config
	err      error
	builders []*OrgUserPreferenceCreate
	conflict []sql.ConflictOption
}

// Save creates the OrgUserPreference entities in the database.
func (oupcb *OrgUserPreferenceCreateBulk) Save(ctx context.Context) ([]*OrgUserPreference, error) {
	if oupcb.err != nil {
		return nil, oupcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oupcb.builders))
	nodes := make([]*OrgUserPreference, len(oupcb.builders))
	mutators := make([]Mutator, len(oupcb.builders))
	for i := range oupcb.builders {
		func(i int, root context.Context) {
			builder := oupcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgUserPreferenceMutation)
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
					_, err = mutators[i+1].Mutate(root, oupcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = oupcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oupcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oupcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oupcb *OrgUserPreferenceCreateBulk) SaveX(ctx context.Context) []*OrgUserPreference {
	v, err := oupcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oupcb *OrgUserPreferenceCreateBulk) Exec(ctx context.Context) error {
	_, err := oupcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oupcb *OrgUserPreferenceCreateBulk) ExecX(ctx context.Context) {
	if err := oupcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrgUserPreference.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrgUserPreferenceUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (oupcb *OrgUserPreferenceCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrgUserPreferenceUpsertBulk {
	oupcb.conflict = opts
	return &OrgUserPreferenceUpsertBulk{
		create: oupcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrgUserPreference.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (oupcb *OrgUserPreferenceCreateBulk) OnConflictColumns(columns ...string) *OrgUserPreferenceUpsertBulk {
	oupcb.conflict = append(oupcb.conflict, sql.ConflictColumns(columns...))
	return &OrgUserPreferenceUpsertBulk{
		create: oupcb,
	}
}

// OrgUserPreferenceUpsertBulk is the builder for "upsert"-ing
// a bulk of OrgUserPreference nodes.
type OrgUserPreferenceUpsertBulk struct {
	create *OrgUserPreferenceCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrgUserPreference.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orguserpreference.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *OrgUserPreferenceUpsertBulk) UpdateNewValues() *OrgUserPreferenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(orguserpreference.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(orguserpreference.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(orguserpreference.FieldCreatedAt)
			}
			if _, exists := b.mutation.OrgID(); exists {
				s.SetIgnore(orguserpreference.FieldOrgID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrgUserPreference.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *OrgUserPreferenceUpsertBulk) Ignore() *OrgUserPreferenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrgUserPreferenceUpsertBulk) DoNothing() *OrgUserPreferenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrgUserPreferenceCreateBulk.OnConflict
// documentation for more info.
func (u *OrgUserPreferenceUpsertBulk) Update(set func(*OrgUserPreferenceUpsert)) *OrgUserPreferenceUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrgUserPreferenceUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *OrgUserPreferenceUpsertBulk) SetUpdatedBy(v int) *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *OrgUserPreferenceUpsertBulk) AddUpdatedBy(v int) *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsertBulk) UpdateUpdatedBy() *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *OrgUserPreferenceUpsertBulk) ClearUpdatedBy() *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrgUserPreferenceUpsertBulk) SetUpdatedAt(v time.Time) *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsertBulk) UpdateUpdatedAt() *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *OrgUserPreferenceUpsertBulk) ClearUpdatedAt() *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrgUserPreferenceUpsertBulk) SetUserID(v int) *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsertBulk) UpdateUserID() *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.UpdateUserID()
	})
}

// SetMenuFavorite sets the "menu_favorite" field.
func (u *OrgUserPreferenceUpsertBulk) SetMenuFavorite(v []int) *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.SetMenuFavorite(v)
	})
}

// UpdateMenuFavorite sets the "menu_favorite" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsertBulk) UpdateMenuFavorite() *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.UpdateMenuFavorite()
	})
}

// ClearMenuFavorite clears the value of the "menu_favorite" field.
func (u *OrgUserPreferenceUpsertBulk) ClearMenuFavorite() *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.ClearMenuFavorite()
	})
}

// SetMenuRecent sets the "menu_recent" field.
func (u *OrgUserPreferenceUpsertBulk) SetMenuRecent(v []int) *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.SetMenuRecent(v)
	})
}

// UpdateMenuRecent sets the "menu_recent" field to the value that was provided on create.
func (u *OrgUserPreferenceUpsertBulk) UpdateMenuRecent() *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.UpdateMenuRecent()
	})
}

// ClearMenuRecent clears the value of the "menu_recent" field.
func (u *OrgUserPreferenceUpsertBulk) ClearMenuRecent() *OrgUserPreferenceUpsertBulk {
	return u.Update(func(s *OrgUserPreferenceUpsert) {
		s.ClearMenuRecent()
	})
}

// Exec executes the query.
func (u *OrgUserPreferenceUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrgUserPreferenceCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrgUserPreferenceCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrgUserPreferenceUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
