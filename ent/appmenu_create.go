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

// AppMenuCreate is the builder for creating a AppMenu entity.
type AppMenuCreate struct {
	config
	mutation *AppMenuMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (amc *AppMenuCreate) SetCreatedBy(i int) *AppMenuCreate {
	amc.mutation.SetCreatedBy(i)
	return amc
}

// SetCreatedAt sets the "created_at" field.
func (amc *AppMenuCreate) SetCreatedAt(t time.Time) *AppMenuCreate {
	amc.mutation.SetCreatedAt(t)
	return amc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (amc *AppMenuCreate) SetNillableCreatedAt(t *time.Time) *AppMenuCreate {
	if t != nil {
		amc.SetCreatedAt(*t)
	}
	return amc
}

// SetUpdatedBy sets the "updated_by" field.
func (amc *AppMenuCreate) SetUpdatedBy(i int) *AppMenuCreate {
	amc.mutation.SetUpdatedBy(i)
	return amc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (amc *AppMenuCreate) SetNillableUpdatedBy(i *int) *AppMenuCreate {
	if i != nil {
		amc.SetUpdatedBy(*i)
	}
	return amc
}

// SetUpdatedAt sets the "updated_at" field.
func (amc *AppMenuCreate) SetUpdatedAt(t time.Time) *AppMenuCreate {
	amc.mutation.SetUpdatedAt(t)
	return amc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (amc *AppMenuCreate) SetNillableUpdatedAt(t *time.Time) *AppMenuCreate {
	if t != nil {
		amc.SetUpdatedAt(*t)
	}
	return amc
}

// SetAppID sets the "app_id" field.
func (amc *AppMenuCreate) SetAppID(i int) *AppMenuCreate {
	amc.mutation.SetAppID(i)
	return amc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (amc *AppMenuCreate) SetNillableAppID(i *int) *AppMenuCreate {
	if i != nil {
		amc.SetAppID(*i)
	}
	return amc
}

// SetParentID sets the "parent_id" field.
func (amc *AppMenuCreate) SetParentID(i int) *AppMenuCreate {
	amc.mutation.SetParentID(i)
	return amc
}

// SetKind sets the "kind" field.
func (amc *AppMenuCreate) SetKind(a appmenu.Kind) *AppMenuCreate {
	amc.mutation.SetKind(a)
	return amc
}

// SetName sets the "name" field.
func (amc *AppMenuCreate) SetName(s string) *AppMenuCreate {
	amc.mutation.SetName(s)
	return amc
}

// SetIcon sets the "icon" field.
func (amc *AppMenuCreate) SetIcon(s string) *AppMenuCreate {
	amc.mutation.SetIcon(s)
	return amc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (amc *AppMenuCreate) SetNillableIcon(s *string) *AppMenuCreate {
	if s != nil {
		amc.SetIcon(*s)
	}
	return amc
}

// SetRoute sets the "route" field.
func (amc *AppMenuCreate) SetRoute(s string) *AppMenuCreate {
	amc.mutation.SetRoute(s)
	return amc
}

// SetNillableRoute sets the "route" field if the given value is not nil.
func (amc *AppMenuCreate) SetNillableRoute(s *string) *AppMenuCreate {
	if s != nil {
		amc.SetRoute(*s)
	}
	return amc
}

// SetActionID sets the "action_id" field.
func (amc *AppMenuCreate) SetActionID(i int) *AppMenuCreate {
	amc.mutation.SetActionID(i)
	return amc
}

// SetNillableActionID sets the "action_id" field if the given value is not nil.
func (amc *AppMenuCreate) SetNillableActionID(i *int) *AppMenuCreate {
	if i != nil {
		amc.SetActionID(*i)
	}
	return amc
}

// SetComments sets the "comments" field.
func (amc *AppMenuCreate) SetComments(s string) *AppMenuCreate {
	amc.mutation.SetComments(s)
	return amc
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (amc *AppMenuCreate) SetNillableComments(s *string) *AppMenuCreate {
	if s != nil {
		amc.SetComments(*s)
	}
	return amc
}

// SetDisplaySort sets the "display_sort" field.
func (amc *AppMenuCreate) SetDisplaySort(i int32) *AppMenuCreate {
	amc.mutation.SetDisplaySort(i)
	return amc
}

// SetNillableDisplaySort sets the "display_sort" field if the given value is not nil.
func (amc *AppMenuCreate) SetNillableDisplaySort(i *int32) *AppMenuCreate {
	if i != nil {
		amc.SetDisplaySort(*i)
	}
	return amc
}

// SetID sets the "id" field.
func (amc *AppMenuCreate) SetID(i int) *AppMenuCreate {
	amc.mutation.SetID(i)
	return amc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (amc *AppMenuCreate) SetNillableID(i *int) *AppMenuCreate {
	if i != nil {
		amc.SetID(*i)
	}
	return amc
}

// SetApp sets the "app" edge to the App entity.
func (amc *AppMenuCreate) SetApp(a *App) *AppMenuCreate {
	return amc.SetAppID(a.ID)
}

// SetAction sets the "action" edge to the AppAction entity.
func (amc *AppMenuCreate) SetAction(a *AppAction) *AppMenuCreate {
	return amc.SetActionID(a.ID)
}

// Mutation returns the AppMenuMutation object of the builder.
func (amc *AppMenuCreate) Mutation() *AppMenuMutation {
	return amc.mutation
}

// Save creates the AppMenu in the database.
func (amc *AppMenuCreate) Save(ctx context.Context) (*AppMenu, error) {
	if err := amc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, amc.sqlSave, amc.mutation, amc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (amc *AppMenuCreate) SaveX(ctx context.Context) *AppMenu {
	v, err := amc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (amc *AppMenuCreate) Exec(ctx context.Context) error {
	_, err := amc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amc *AppMenuCreate) ExecX(ctx context.Context) {
	if err := amc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (amc *AppMenuCreate) defaults() error {
	if _, ok := amc.mutation.CreatedAt(); !ok {
		if appmenu.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appmenu.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appmenu.DefaultCreatedAt()
		amc.mutation.SetCreatedAt(v)
	}
	if _, ok := amc.mutation.ID(); !ok {
		if appmenu.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized appmenu.DefaultID (forgotten import ent/runtime?)")
		}
		v := appmenu.DefaultID()
		amc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (amc *AppMenuCreate) check() error {
	if _, ok := amc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "AppMenu.created_by"`)}
	}
	if _, ok := amc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppMenu.created_at"`)}
	}
	if _, ok := amc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent_id", err: errors.New(`ent: missing required field "AppMenu.parent_id"`)}
	}
	if _, ok := amc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "AppMenu.kind"`)}
	}
	if v, ok := amc.mutation.Kind(); ok {
		if err := appmenu.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "AppMenu.kind": %w`, err)}
		}
	}
	if _, ok := amc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AppMenu.name"`)}
	}
	return nil
}

func (amc *AppMenuCreate) sqlSave(ctx context.Context) (*AppMenu, error) {
	if err := amc.check(); err != nil {
		return nil, err
	}
	_node, _spec := amc.createSpec()
	if err := sqlgraph.CreateNode(ctx, amc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	amc.mutation.id = &_node.ID
	amc.mutation.done = true
	return _node, nil
}

func (amc *AppMenuCreate) createSpec() (*AppMenu, *sqlgraph.CreateSpec) {
	var (
		_node = &AppMenu{config: amc.config}
		_spec = sqlgraph.NewCreateSpec(appmenu.Table, sqlgraph.NewFieldSpec(appmenu.FieldID, field.TypeInt))
	)
	_spec.OnConflict = amc.conflict
	if id, ok := amc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := amc.mutation.CreatedBy(); ok {
		_spec.SetField(appmenu.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := amc.mutation.CreatedAt(); ok {
		_spec.SetField(appmenu.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := amc.mutation.UpdatedBy(); ok {
		_spec.SetField(appmenu.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := amc.mutation.UpdatedAt(); ok {
		_spec.SetField(appmenu.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := amc.mutation.ParentID(); ok {
		_spec.SetField(appmenu.FieldParentID, field.TypeInt, value)
		_node.ParentID = value
	}
	if value, ok := amc.mutation.Kind(); ok {
		_spec.SetField(appmenu.FieldKind, field.TypeEnum, value)
		_node.Kind = value
	}
	if value, ok := amc.mutation.Name(); ok {
		_spec.SetField(appmenu.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := amc.mutation.Icon(); ok {
		_spec.SetField(appmenu.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if value, ok := amc.mutation.Route(); ok {
		_spec.SetField(appmenu.FieldRoute, field.TypeString, value)
		_node.Route = value
	}
	if value, ok := amc.mutation.Comments(); ok {
		_spec.SetField(appmenu.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if value, ok := amc.mutation.DisplaySort(); ok {
		_spec.SetField(appmenu.FieldDisplaySort, field.TypeInt32, value)
		_node.DisplaySort = value
	}
	if nodes := amc.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appmenu.AppTable,
			Columns: []string{appmenu.AppColumn},
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
	if nodes := amc.mutation.ActionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appmenu.ActionTable,
			Columns: []string{appmenu.ActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ActionID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppMenu.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppMenuUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (amc *AppMenuCreate) OnConflict(opts ...sql.ConflictOption) *AppMenuUpsertOne {
	amc.conflict = opts
	return &AppMenuUpsertOne{
		create: amc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppMenu.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (amc *AppMenuCreate) OnConflictColumns(columns ...string) *AppMenuUpsertOne {
	amc.conflict = append(amc.conflict, sql.ConflictColumns(columns...))
	return &AppMenuUpsertOne{
		create: amc,
	}
}

type (
	// AppMenuUpsertOne is the builder for "upsert"-ing
	//  one AppMenu node.
	AppMenuUpsertOne struct {
		create *AppMenuCreate
	}

	// AppMenuUpsert is the "OnConflict" setter.
	AppMenuUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *AppMenuUpsert) SetUpdatedBy(v int) *AppMenuUpsert {
	u.Set(appmenu.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppMenuUpsert) UpdateUpdatedBy() *AppMenuUpsert {
	u.SetExcluded(appmenu.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppMenuUpsert) AddUpdatedBy(v int) *AppMenuUpsert {
	u.Add(appmenu.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppMenuUpsert) ClearUpdatedBy() *AppMenuUpsert {
	u.SetNull(appmenu.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppMenuUpsert) SetUpdatedAt(v time.Time) *AppMenuUpsert {
	u.Set(appmenu.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppMenuUpsert) UpdateUpdatedAt() *AppMenuUpsert {
	u.SetExcluded(appmenu.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppMenuUpsert) ClearUpdatedAt() *AppMenuUpsert {
	u.SetNull(appmenu.FieldUpdatedAt)
	return u
}

// SetParentID sets the "parent_id" field.
func (u *AppMenuUpsert) SetParentID(v int) *AppMenuUpsert {
	u.Set(appmenu.FieldParentID, v)
	return u
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *AppMenuUpsert) UpdateParentID() *AppMenuUpsert {
	u.SetExcluded(appmenu.FieldParentID)
	return u
}

// AddParentID adds v to the "parent_id" field.
func (u *AppMenuUpsert) AddParentID(v int) *AppMenuUpsert {
	u.Add(appmenu.FieldParentID, v)
	return u
}

// SetKind sets the "kind" field.
func (u *AppMenuUpsert) SetKind(v appmenu.Kind) *AppMenuUpsert {
	u.Set(appmenu.FieldKind, v)
	return u
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *AppMenuUpsert) UpdateKind() *AppMenuUpsert {
	u.SetExcluded(appmenu.FieldKind)
	return u
}

// SetName sets the "name" field.
func (u *AppMenuUpsert) SetName(v string) *AppMenuUpsert {
	u.Set(appmenu.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppMenuUpsert) UpdateName() *AppMenuUpsert {
	u.SetExcluded(appmenu.FieldName)
	return u
}

// SetIcon sets the "icon" field.
func (u *AppMenuUpsert) SetIcon(v string) *AppMenuUpsert {
	u.Set(appmenu.FieldIcon, v)
	return u
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *AppMenuUpsert) UpdateIcon() *AppMenuUpsert {
	u.SetExcluded(appmenu.FieldIcon)
	return u
}

// ClearIcon clears the value of the "icon" field.
func (u *AppMenuUpsert) ClearIcon() *AppMenuUpsert {
	u.SetNull(appmenu.FieldIcon)
	return u
}

// SetRoute sets the "route" field.
func (u *AppMenuUpsert) SetRoute(v string) *AppMenuUpsert {
	u.Set(appmenu.FieldRoute, v)
	return u
}

// UpdateRoute sets the "route" field to the value that was provided on create.
func (u *AppMenuUpsert) UpdateRoute() *AppMenuUpsert {
	u.SetExcluded(appmenu.FieldRoute)
	return u
}

// ClearRoute clears the value of the "route" field.
func (u *AppMenuUpsert) ClearRoute() *AppMenuUpsert {
	u.SetNull(appmenu.FieldRoute)
	return u
}

// SetActionID sets the "action_id" field.
func (u *AppMenuUpsert) SetActionID(v int) *AppMenuUpsert {
	u.Set(appmenu.FieldActionID, v)
	return u
}

// UpdateActionID sets the "action_id" field to the value that was provided on create.
func (u *AppMenuUpsert) UpdateActionID() *AppMenuUpsert {
	u.SetExcluded(appmenu.FieldActionID)
	return u
}

// ClearActionID clears the value of the "action_id" field.
func (u *AppMenuUpsert) ClearActionID() *AppMenuUpsert {
	u.SetNull(appmenu.FieldActionID)
	return u
}

// SetComments sets the "comments" field.
func (u *AppMenuUpsert) SetComments(v string) *AppMenuUpsert {
	u.Set(appmenu.FieldComments, v)
	return u
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppMenuUpsert) UpdateComments() *AppMenuUpsert {
	u.SetExcluded(appmenu.FieldComments)
	return u
}

// ClearComments clears the value of the "comments" field.
func (u *AppMenuUpsert) ClearComments() *AppMenuUpsert {
	u.SetNull(appmenu.FieldComments)
	return u
}

// SetDisplaySort sets the "display_sort" field.
func (u *AppMenuUpsert) SetDisplaySort(v int32) *AppMenuUpsert {
	u.Set(appmenu.FieldDisplaySort, v)
	return u
}

// UpdateDisplaySort sets the "display_sort" field to the value that was provided on create.
func (u *AppMenuUpsert) UpdateDisplaySort() *AppMenuUpsert {
	u.SetExcluded(appmenu.FieldDisplaySort)
	return u
}

// AddDisplaySort adds v to the "display_sort" field.
func (u *AppMenuUpsert) AddDisplaySort(v int32) *AppMenuUpsert {
	u.Add(appmenu.FieldDisplaySort, v)
	return u
}

// ClearDisplaySort clears the value of the "display_sort" field.
func (u *AppMenuUpsert) ClearDisplaySort() *AppMenuUpsert {
	u.SetNull(appmenu.FieldDisplaySort)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppMenu.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appmenu.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppMenuUpsertOne) UpdateNewValues() *AppMenuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appmenu.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(appmenu.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(appmenu.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.AppID(); exists {
			s.SetIgnore(appmenu.FieldAppID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppMenu.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppMenuUpsertOne) Ignore() *AppMenuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppMenuUpsertOne) DoNothing() *AppMenuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppMenuCreate.OnConflict
// documentation for more info.
func (u *AppMenuUpsertOne) Update(set func(*AppMenuUpsert)) *AppMenuUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppMenuUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AppMenuUpsertOne) SetUpdatedBy(v int) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppMenuUpsertOne) AddUpdatedBy(v int) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppMenuUpsertOne) UpdateUpdatedBy() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppMenuUpsertOne) ClearUpdatedBy() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppMenuUpsertOne) SetUpdatedAt(v time.Time) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppMenuUpsertOne) UpdateUpdatedAt() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppMenuUpsertOne) ClearUpdatedAt() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetParentID sets the "parent_id" field.
func (u *AppMenuUpsertOne) SetParentID(v int) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetParentID(v)
	})
}

// AddParentID adds v to the "parent_id" field.
func (u *AppMenuUpsertOne) AddParentID(v int) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.AddParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *AppMenuUpsertOne) UpdateParentID() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateParentID()
	})
}

// SetKind sets the "kind" field.
func (u *AppMenuUpsertOne) SetKind(v appmenu.Kind) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetKind(v)
	})
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *AppMenuUpsertOne) UpdateKind() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateKind()
	})
}

// SetName sets the "name" field.
func (u *AppMenuUpsertOne) SetName(v string) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppMenuUpsertOne) UpdateName() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateName()
	})
}

// SetIcon sets the "icon" field.
func (u *AppMenuUpsertOne) SetIcon(v string) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetIcon(v)
	})
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *AppMenuUpsertOne) UpdateIcon() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateIcon()
	})
}

// ClearIcon clears the value of the "icon" field.
func (u *AppMenuUpsertOne) ClearIcon() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearIcon()
	})
}

// SetRoute sets the "route" field.
func (u *AppMenuUpsertOne) SetRoute(v string) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetRoute(v)
	})
}

// UpdateRoute sets the "route" field to the value that was provided on create.
func (u *AppMenuUpsertOne) UpdateRoute() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateRoute()
	})
}

// ClearRoute clears the value of the "route" field.
func (u *AppMenuUpsertOne) ClearRoute() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearRoute()
	})
}

// SetActionID sets the "action_id" field.
func (u *AppMenuUpsertOne) SetActionID(v int) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetActionID(v)
	})
}

// UpdateActionID sets the "action_id" field to the value that was provided on create.
func (u *AppMenuUpsertOne) UpdateActionID() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateActionID()
	})
}

// ClearActionID clears the value of the "action_id" field.
func (u *AppMenuUpsertOne) ClearActionID() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearActionID()
	})
}

// SetComments sets the "comments" field.
func (u *AppMenuUpsertOne) SetComments(v string) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppMenuUpsertOne) UpdateComments() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *AppMenuUpsertOne) ClearComments() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearComments()
	})
}

// SetDisplaySort sets the "display_sort" field.
func (u *AppMenuUpsertOne) SetDisplaySort(v int32) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetDisplaySort(v)
	})
}

// AddDisplaySort adds v to the "display_sort" field.
func (u *AppMenuUpsertOne) AddDisplaySort(v int32) *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.AddDisplaySort(v)
	})
}

// UpdateDisplaySort sets the "display_sort" field to the value that was provided on create.
func (u *AppMenuUpsertOne) UpdateDisplaySort() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateDisplaySort()
	})
}

// ClearDisplaySort clears the value of the "display_sort" field.
func (u *AppMenuUpsertOne) ClearDisplaySort() *AppMenuUpsertOne {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearDisplaySort()
	})
}

// Exec executes the query.
func (u *AppMenuUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppMenuCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppMenuUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppMenuUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppMenuUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppMenuCreateBulk is the builder for creating many AppMenu entities in bulk.
type AppMenuCreateBulk struct {
	config
	builders []*AppMenuCreate
	conflict []sql.ConflictOption
}

// Save creates the AppMenu entities in the database.
func (amcb *AppMenuCreateBulk) Save(ctx context.Context) ([]*AppMenu, error) {
	specs := make([]*sqlgraph.CreateSpec, len(amcb.builders))
	nodes := make([]*AppMenu, len(amcb.builders))
	mutators := make([]Mutator, len(amcb.builders))
	for i := range amcb.builders {
		func(i int, root context.Context) {
			builder := amcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppMenuMutation)
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
					_, err = mutators[i+1].Mutate(root, amcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = amcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, amcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, amcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (amcb *AppMenuCreateBulk) SaveX(ctx context.Context) []*AppMenu {
	v, err := amcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (amcb *AppMenuCreateBulk) Exec(ctx context.Context) error {
	_, err := amcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amcb *AppMenuCreateBulk) ExecX(ctx context.Context) {
	if err := amcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppMenu.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppMenuUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (amcb *AppMenuCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppMenuUpsertBulk {
	amcb.conflict = opts
	return &AppMenuUpsertBulk{
		create: amcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppMenu.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (amcb *AppMenuCreateBulk) OnConflictColumns(columns ...string) *AppMenuUpsertBulk {
	amcb.conflict = append(amcb.conflict, sql.ConflictColumns(columns...))
	return &AppMenuUpsertBulk{
		create: amcb,
	}
}

// AppMenuUpsertBulk is the builder for "upsert"-ing
// a bulk of AppMenu nodes.
type AppMenuUpsertBulk struct {
	create *AppMenuCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppMenu.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appmenu.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppMenuUpsertBulk) UpdateNewValues() *AppMenuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appmenu.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(appmenu.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(appmenu.FieldCreatedAt)
			}
			if _, exists := b.mutation.AppID(); exists {
				s.SetIgnore(appmenu.FieldAppID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppMenu.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppMenuUpsertBulk) Ignore() *AppMenuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppMenuUpsertBulk) DoNothing() *AppMenuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppMenuCreateBulk.OnConflict
// documentation for more info.
func (u *AppMenuUpsertBulk) Update(set func(*AppMenuUpsert)) *AppMenuUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppMenuUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AppMenuUpsertBulk) SetUpdatedBy(v int) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppMenuUpsertBulk) AddUpdatedBy(v int) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppMenuUpsertBulk) UpdateUpdatedBy() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppMenuUpsertBulk) ClearUpdatedBy() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppMenuUpsertBulk) SetUpdatedAt(v time.Time) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppMenuUpsertBulk) UpdateUpdatedAt() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppMenuUpsertBulk) ClearUpdatedAt() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetParentID sets the "parent_id" field.
func (u *AppMenuUpsertBulk) SetParentID(v int) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetParentID(v)
	})
}

// AddParentID adds v to the "parent_id" field.
func (u *AppMenuUpsertBulk) AddParentID(v int) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.AddParentID(v)
	})
}

// UpdateParentID sets the "parent_id" field to the value that was provided on create.
func (u *AppMenuUpsertBulk) UpdateParentID() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateParentID()
	})
}

// SetKind sets the "kind" field.
func (u *AppMenuUpsertBulk) SetKind(v appmenu.Kind) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetKind(v)
	})
}

// UpdateKind sets the "kind" field to the value that was provided on create.
func (u *AppMenuUpsertBulk) UpdateKind() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateKind()
	})
}

// SetName sets the "name" field.
func (u *AppMenuUpsertBulk) SetName(v string) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppMenuUpsertBulk) UpdateName() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateName()
	})
}

// SetIcon sets the "icon" field.
func (u *AppMenuUpsertBulk) SetIcon(v string) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetIcon(v)
	})
}

// UpdateIcon sets the "icon" field to the value that was provided on create.
func (u *AppMenuUpsertBulk) UpdateIcon() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateIcon()
	})
}

// ClearIcon clears the value of the "icon" field.
func (u *AppMenuUpsertBulk) ClearIcon() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearIcon()
	})
}

// SetRoute sets the "route" field.
func (u *AppMenuUpsertBulk) SetRoute(v string) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetRoute(v)
	})
}

// UpdateRoute sets the "route" field to the value that was provided on create.
func (u *AppMenuUpsertBulk) UpdateRoute() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateRoute()
	})
}

// ClearRoute clears the value of the "route" field.
func (u *AppMenuUpsertBulk) ClearRoute() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearRoute()
	})
}

// SetActionID sets the "action_id" field.
func (u *AppMenuUpsertBulk) SetActionID(v int) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetActionID(v)
	})
}

// UpdateActionID sets the "action_id" field to the value that was provided on create.
func (u *AppMenuUpsertBulk) UpdateActionID() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateActionID()
	})
}

// ClearActionID clears the value of the "action_id" field.
func (u *AppMenuUpsertBulk) ClearActionID() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearActionID()
	})
}

// SetComments sets the "comments" field.
func (u *AppMenuUpsertBulk) SetComments(v string) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppMenuUpsertBulk) UpdateComments() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *AppMenuUpsertBulk) ClearComments() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearComments()
	})
}

// SetDisplaySort sets the "display_sort" field.
func (u *AppMenuUpsertBulk) SetDisplaySort(v int32) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.SetDisplaySort(v)
	})
}

// AddDisplaySort adds v to the "display_sort" field.
func (u *AppMenuUpsertBulk) AddDisplaySort(v int32) *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.AddDisplaySort(v)
	})
}

// UpdateDisplaySort sets the "display_sort" field to the value that was provided on create.
func (u *AppMenuUpsertBulk) UpdateDisplaySort() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.UpdateDisplaySort()
	})
}

// ClearDisplaySort clears the value of the "display_sort" field.
func (u *AppMenuUpsertBulk) ClearDisplaySort() *AppMenuUpsertBulk {
	return u.Update(func(s *AppMenuUpsert) {
		s.ClearDisplaySort()
	})
}

// Exec executes the query.
func (u *AppMenuUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppMenuCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppMenuCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppMenuUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
