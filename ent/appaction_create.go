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
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appmenu"
)

// AppActionCreate is the builder for creating a AppAction entity.
type AppActionCreate struct {
	config
	mutation *AppActionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (aac *AppActionCreate) SetCreatedBy(i int) *AppActionCreate {
	aac.mutation.SetCreatedBy(i)
	return aac
}

// SetCreatedAt sets the "created_at" field.
func (aac *AppActionCreate) SetCreatedAt(t time.Time) *AppActionCreate {
	aac.mutation.SetCreatedAt(t)
	return aac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aac *AppActionCreate) SetNillableCreatedAt(t *time.Time) *AppActionCreate {
	if t != nil {
		aac.SetCreatedAt(*t)
	}
	return aac
}

// SetUpdatedBy sets the "updated_by" field.
func (aac *AppActionCreate) SetUpdatedBy(i int) *AppActionCreate {
	aac.mutation.SetUpdatedBy(i)
	return aac
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (aac *AppActionCreate) SetNillableUpdatedBy(i *int) *AppActionCreate {
	if i != nil {
		aac.SetUpdatedBy(*i)
	}
	return aac
}

// SetUpdatedAt sets the "updated_at" field.
func (aac *AppActionCreate) SetUpdatedAt(t time.Time) *AppActionCreate {
	aac.mutation.SetUpdatedAt(t)
	return aac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aac *AppActionCreate) SetNillableUpdatedAt(t *time.Time) *AppActionCreate {
	if t != nil {
		aac.SetUpdatedAt(*t)
	}
	return aac
}

// SetAppID sets the "app_id" field.
func (aac *AppActionCreate) SetAppID(i int) *AppActionCreate {
	aac.mutation.SetAppID(i)
	return aac
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (aac *AppActionCreate) SetNillableAppID(i *int) *AppActionCreate {
	if i != nil {
		aac.SetAppID(*i)
	}
	return aac
}

// SetName sets the "name" field.
func (aac *AppActionCreate) SetName(s string) *AppActionCreate {
	aac.mutation.SetName(s)
	return aac
}

// SetKind sets the "kind" field.
func (aac *AppActionCreate) SetKind(a appaction.Kind) *AppActionCreate {
	aac.mutation.SetKind(a)
	return aac
}

// SetMethod sets the "method" field.
func (aac *AppActionCreate) SetMethod(a appaction.Method) *AppActionCreate {
	aac.mutation.SetMethod(a)
	return aac
}

// SetComments sets the "comments" field.
func (aac *AppActionCreate) SetComments(s string) *AppActionCreate {
	aac.mutation.SetComments(s)
	return aac
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (aac *AppActionCreate) SetNillableComments(s *string) *AppActionCreate {
	if s != nil {
		aac.SetComments(*s)
	}
	return aac
}

// SetID sets the "id" field.
func (aac *AppActionCreate) SetID(i int) *AppActionCreate {
	aac.mutation.SetID(i)
	return aac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (aac *AppActionCreate) SetNillableID(i *int) *AppActionCreate {
	if i != nil {
		aac.SetID(*i)
	}
	return aac
}

// SetApp sets the "app" edge to the App entity.
func (aac *AppActionCreate) SetApp(a *App) *AppActionCreate {
	return aac.SetAppID(a.ID)
}

// AddMenuIDs adds the "menus" edge to the AppMenu entity by IDs.
func (aac *AppActionCreate) AddMenuIDs(ids ...int) *AppActionCreate {
	aac.mutation.AddMenuIDs(ids...)
	return aac
}

// AddMenus adds the "menus" edges to the AppMenu entity.
func (aac *AppActionCreate) AddMenus(a ...*AppMenu) *AppActionCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aac.AddMenuIDs(ids...)
}

// Mutation returns the AppActionMutation object of the builder.
func (aac *AppActionCreate) Mutation() *AppActionMutation {
	return aac.mutation
}

// Save creates the AppAction in the database.
func (aac *AppActionCreate) Save(ctx context.Context) (*AppAction, error) {
	if err := aac.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, aac.sqlSave, aac.mutation, aac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (aac *AppActionCreate) SaveX(ctx context.Context) *AppAction {
	v, err := aac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aac *AppActionCreate) Exec(ctx context.Context) error {
	_, err := aac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aac *AppActionCreate) ExecX(ctx context.Context) {
	if err := aac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aac *AppActionCreate) defaults() error {
	if _, ok := aac.mutation.CreatedAt(); !ok {
		if appaction.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appaction.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appaction.DefaultCreatedAt()
		aac.mutation.SetCreatedAt(v)
	}
	if _, ok := aac.mutation.ID(); !ok {
		if appaction.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized appaction.DefaultID (forgotten import ent/runtime?)")
		}
		v := appaction.DefaultID()
		aac.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (aac *AppActionCreate) check() error {
	if _, ok := aac.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "AppAction.created_by"`)}
	}
	if _, ok := aac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppAction.created_at"`)}
	}
	if _, ok := aac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AppAction.name"`)}
	}
	if v, ok := aac.mutation.Name(); ok {
		if err := appaction.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AppAction.name": %w`, err)}
		}
	}
	if _, ok := aac.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "AppAction.kind"`)}
	}
	if v, ok := aac.mutation.Kind(); ok {
		if err := appaction.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "AppAction.kind": %w`, err)}
		}
	}
	if _, ok := aac.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "AppAction.method"`)}
	}
	if v, ok := aac.mutation.Method(); ok {
		if err := appaction.MethodValidator(v); err != nil {
			return &ValidationError{Name: "method", err: fmt.Errorf(`ent: validator failed for field "AppAction.method": %w`, err)}
		}
	}
	return nil
}

func (aac *AppActionCreate) sqlSave(ctx context.Context) (*AppAction, error) {
	if err := aac.check(); err != nil {
		return nil, err
	}
	_node, _spec := aac.createSpec()
	if err := sqlgraph.CreateNode(ctx, aac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	aac.mutation.id = &_node.ID
	aac.mutation.done = true
	return _node, nil
}

func (aac *AppActionCreate) createSpec() (*AppAction, *sqlgraph.CreateSpec) {
	var (
		_node = &AppAction{config: aac.config}
		_spec = sqlgraph.NewCreateSpec(appaction.Table, sqlgraph.NewFieldSpec(appaction.FieldID, field.TypeInt))
	)
	_spec.OnConflict = aac.conflict
	if id, ok := aac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := aac.mutation.CreatedBy(); ok {
		_spec.SetField(appaction.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := aac.mutation.CreatedAt(); ok {
		_spec.SetField(appaction.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := aac.mutation.UpdatedBy(); ok {
		_spec.SetField(appaction.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := aac.mutation.UpdatedAt(); ok {
		_spec.SetField(appaction.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := aac.mutation.Name(); ok {
		_spec.SetField(appaction.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := aac.mutation.Kind(); ok {
		_spec.SetField(appaction.FieldKind, field.TypeEnum, value)
		_node.Kind = value
	}
	if value, ok := aac.mutation.Method(); ok {
		_spec.SetField(appaction.FieldMethod, field.TypeEnum, value)
		_node.Method = value
	}
	if value, ok := aac.mutation.Comments(); ok {
		_spec.SetField(appaction.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if nodes := aac.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appaction.AppTable,
			Columns: []string{appaction.AppColumn},
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
	if nodes := aac.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appaction.MenusTable,
			Columns: []string{appaction.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appmenu.FieldID, field.TypeInt),
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
//	client.AppAction.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppActionUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (aac *AppActionCreate) OnConflict(opts ...sql.ConflictOption) *AppActionUpsertOne {
	aac.conflict = opts
	return &AppActionUpsertOne{
		create: aac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppAction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (aac *AppActionCreate) OnConflictColumns(columns ...string) *AppActionUpsertOne {
	aac.conflict = append(aac.conflict, sql.ConflictColumns(columns...))
	return &AppActionUpsertOne{
		create: aac,
	}
}

type (
	// AppActionUpsertOne is the builder for "upsert"-ing
	//  one AppAction node.
	AppActionUpsertOne struct {
		create *AppActionCreate
	}

	// AppActionUpsert is the "OnConflict" setter.
	AppActionUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *AppActionUpsert) SetUpdatedBy(v int) *AppActionUpsert {
	u.Set(appaction.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppActionUpsert) UpdateUpdatedBy() *AppActionUpsert {
	u.SetExcluded(appaction.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppActionUpsert) AddUpdatedBy(v int) *AppActionUpsert {
	u.Add(appaction.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppActionUpsert) ClearUpdatedBy() *AppActionUpsert {
	u.SetNull(appaction.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppActionUpsert) SetUpdatedAt(v time.Time) *AppActionUpsert {
	u.Set(appaction.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppActionUpsert) UpdateUpdatedAt() *AppActionUpsert {
	u.SetExcluded(appaction.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppActionUpsert) ClearUpdatedAt() *AppActionUpsert {
	u.SetNull(appaction.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *AppActionUpsert) SetName(v string) *AppActionUpsert {
	u.Set(appaction.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppActionUpsert) UpdateName() *AppActionUpsert {
	u.SetExcluded(appaction.FieldName)
	return u
}

// SetKind sets the "kind" field.
func (u *AppActionUpsert) SetKind(v appaction.Kind) *AppActionUpsert {
	u.Set(appaction.FieldKind, v)
	return u
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *AppActionUpsert) UpdateKind() *AppActionUpsert {
	u.SetExcluded(appaction.FieldKind)
	return u
}

// SetMethod sets the "method" field.
func (u *AppActionUpsert) SetMethod(v appaction.Method) *AppActionUpsert {
	u.Set(appaction.FieldMethod, v)
	return u
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AppActionUpsert) UpdateMethod() *AppActionUpsert {
	u.SetExcluded(appaction.FieldMethod)
	return u
}

// SetComments sets the "comments" field.
func (u *AppActionUpsert) SetComments(v string) *AppActionUpsert {
	u.Set(appaction.FieldComments, v)
	return u
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppActionUpsert) UpdateComments() *AppActionUpsert {
	u.SetExcluded(appaction.FieldComments)
	return u
}

// ClearComments clears the value of the "comments" field.
func (u *AppActionUpsert) ClearComments() *AppActionUpsert {
	u.SetNull(appaction.FieldComments)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppAction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appaction.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppActionUpsertOne) UpdateNewValues() *AppActionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appaction.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(appaction.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(appaction.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.AppID(); exists {
			s.SetIgnore(appaction.FieldAppID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppAction.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppActionUpsertOne) Ignore() *AppActionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppActionUpsertOne) DoNothing() *AppActionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppActionCreate.OnConflict
// documentation for more info.
func (u *AppActionUpsertOne) Update(set func(*AppActionUpsert)) *AppActionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppActionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AppActionUpsertOne) SetUpdatedBy(v int) *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppActionUpsertOne) AddUpdatedBy(v int) *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppActionUpsertOne) UpdateUpdatedBy() *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppActionUpsertOne) ClearUpdatedBy() *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppActionUpsertOne) SetUpdatedAt(v time.Time) *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppActionUpsertOne) UpdateUpdatedAt() *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppActionUpsertOne) ClearUpdatedAt() *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *AppActionUpsertOne) SetName(v string) *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppActionUpsertOne) UpdateName() *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateName()
	})
}

// SetKind sets the "kind" field.
func (u *AppActionUpsertOne) SetKind(v appaction.Kind) *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.SetKind(v)
	})
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *AppActionUpsertOne) UpdateKind() *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateKind()
	})
}

// SetMethod sets the "method" field.
func (u *AppActionUpsertOne) SetMethod(v appaction.Method) *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AppActionUpsertOne) UpdateMethod() *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateMethod()
	})
}

// SetComments sets the "comments" field.
func (u *AppActionUpsertOne) SetComments(v string) *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppActionUpsertOne) UpdateComments() *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *AppActionUpsertOne) ClearComments() *AppActionUpsertOne {
	return u.Update(func(s *AppActionUpsert) {
		s.ClearComments()
	})
}

// Exec executes the query.
func (u *AppActionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppActionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppActionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppActionUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppActionUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppActionCreateBulk is the builder for creating many AppAction entities in bulk.
type AppActionCreateBulk struct {
	config
	err      error
	builders []*AppActionCreate
	conflict []sql.ConflictOption
}

// Save creates the AppAction entities in the database.
func (aacb *AppActionCreateBulk) Save(ctx context.Context) ([]*AppAction, error) {
	if aacb.err != nil {
		return nil, aacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(aacb.builders))
	nodes := make([]*AppAction, len(aacb.builders))
	mutators := make([]Mutator, len(aacb.builders))
	for i := range aacb.builders {
		func(i int, root context.Context) {
			builder := aacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppActionMutation)
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
					_, err = mutators[i+1].Mutate(root, aacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = aacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aacb *AppActionCreateBulk) SaveX(ctx context.Context) []*AppAction {
	v, err := aacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aacb *AppActionCreateBulk) Exec(ctx context.Context) error {
	_, err := aacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aacb *AppActionCreateBulk) ExecX(ctx context.Context) {
	if err := aacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppAction.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppActionUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (aacb *AppActionCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppActionUpsertBulk {
	aacb.conflict = opts
	return &AppActionUpsertBulk{
		create: aacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppAction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (aacb *AppActionCreateBulk) OnConflictColumns(columns ...string) *AppActionUpsertBulk {
	aacb.conflict = append(aacb.conflict, sql.ConflictColumns(columns...))
	return &AppActionUpsertBulk{
		create: aacb,
	}
}

// AppActionUpsertBulk is the builder for "upsert"-ing
// a bulk of AppAction nodes.
type AppActionUpsertBulk struct {
	create *AppActionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppAction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appaction.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppActionUpsertBulk) UpdateNewValues() *AppActionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appaction.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(appaction.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(appaction.FieldCreatedAt)
			}
			if _, exists := b.mutation.AppID(); exists {
				s.SetIgnore(appaction.FieldAppID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppAction.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppActionUpsertBulk) Ignore() *AppActionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppActionUpsertBulk) DoNothing() *AppActionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppActionCreateBulk.OnConflict
// documentation for more info.
func (u *AppActionUpsertBulk) Update(set func(*AppActionUpsert)) *AppActionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppActionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AppActionUpsertBulk) SetUpdatedBy(v int) *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppActionUpsertBulk) AddUpdatedBy(v int) *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppActionUpsertBulk) UpdateUpdatedBy() *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppActionUpsertBulk) ClearUpdatedBy() *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppActionUpsertBulk) SetUpdatedAt(v time.Time) *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppActionUpsertBulk) UpdateUpdatedAt() *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppActionUpsertBulk) ClearUpdatedAt() *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *AppActionUpsertBulk) SetName(v string) *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppActionUpsertBulk) UpdateName() *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateName()
	})
}

// SetKind sets the "kind" field.
func (u *AppActionUpsertBulk) SetKind(v appaction.Kind) *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.SetKind(v)
	})
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *AppActionUpsertBulk) UpdateKind() *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateKind()
	})
}

// SetMethod sets the "method" field.
func (u *AppActionUpsertBulk) SetMethod(v appaction.Method) *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AppActionUpsertBulk) UpdateMethod() *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateMethod()
	})
}

// SetComments sets the "comments" field.
func (u *AppActionUpsertBulk) SetComments(v string) *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppActionUpsertBulk) UpdateComments() *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *AppActionUpsertBulk) ClearComments() *AppActionUpsertBulk {
	return u.Update(func(s *AppActionUpsert) {
		s.ClearComments()
	})
}

// Exec executes the query.
func (u *AppActionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppActionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppActionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppActionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
