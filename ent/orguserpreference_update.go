// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/orguserpreference"
	"github.com/woocoos/knockout/ent/predicate"
	"github.com/woocoos/knockout/ent/user"
)

// OrgUserPreferenceUpdate is the builder for updating OrgUserPreference entities.
type OrgUserPreferenceUpdate struct {
	config
	hooks    []Hook
	mutation *OrgUserPreferenceMutation
}

// Where appends a list predicates to the OrgUserPreferenceUpdate builder.
func (oupu *OrgUserPreferenceUpdate) Where(ps ...predicate.OrgUserPreference) *OrgUserPreferenceUpdate {
	oupu.mutation.Where(ps...)
	return oupu
}

// SetUpdatedBy sets the "updated_by" field.
func (oupu *OrgUserPreferenceUpdate) SetUpdatedBy(i int) *OrgUserPreferenceUpdate {
	oupu.mutation.ResetUpdatedBy()
	oupu.mutation.SetUpdatedBy(i)
	return oupu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oupu *OrgUserPreferenceUpdate) SetNillableUpdatedBy(i *int) *OrgUserPreferenceUpdate {
	if i != nil {
		oupu.SetUpdatedBy(*i)
	}
	return oupu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (oupu *OrgUserPreferenceUpdate) AddUpdatedBy(i int) *OrgUserPreferenceUpdate {
	oupu.mutation.AddUpdatedBy(i)
	return oupu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (oupu *OrgUserPreferenceUpdate) ClearUpdatedBy() *OrgUserPreferenceUpdate {
	oupu.mutation.ClearUpdatedBy()
	return oupu
}

// SetUpdatedAt sets the "updated_at" field.
func (oupu *OrgUserPreferenceUpdate) SetUpdatedAt(t time.Time) *OrgUserPreferenceUpdate {
	oupu.mutation.SetUpdatedAt(t)
	return oupu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oupu *OrgUserPreferenceUpdate) SetNillableUpdatedAt(t *time.Time) *OrgUserPreferenceUpdate {
	if t != nil {
		oupu.SetUpdatedAt(*t)
	}
	return oupu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (oupu *OrgUserPreferenceUpdate) ClearUpdatedAt() *OrgUserPreferenceUpdate {
	oupu.mutation.ClearUpdatedAt()
	return oupu
}

// SetUserID sets the "user_id" field.
func (oupu *OrgUserPreferenceUpdate) SetUserID(i int) *OrgUserPreferenceUpdate {
	oupu.mutation.SetUserID(i)
	return oupu
}

// SetMenuFavorite sets the "menu_favorite" field.
func (oupu *OrgUserPreferenceUpdate) SetMenuFavorite(i []int) *OrgUserPreferenceUpdate {
	oupu.mutation.SetMenuFavorite(i)
	return oupu
}

// AppendMenuFavorite appends i to the "menu_favorite" field.
func (oupu *OrgUserPreferenceUpdate) AppendMenuFavorite(i []int) *OrgUserPreferenceUpdate {
	oupu.mutation.AppendMenuFavorite(i)
	return oupu
}

// ClearMenuFavorite clears the value of the "menu_favorite" field.
func (oupu *OrgUserPreferenceUpdate) ClearMenuFavorite() *OrgUserPreferenceUpdate {
	oupu.mutation.ClearMenuFavorite()
	return oupu
}

// SetMenuRecent sets the "menu_recent" field.
func (oupu *OrgUserPreferenceUpdate) SetMenuRecent(i []int) *OrgUserPreferenceUpdate {
	oupu.mutation.SetMenuRecent(i)
	return oupu
}

// AppendMenuRecent appends i to the "menu_recent" field.
func (oupu *OrgUserPreferenceUpdate) AppendMenuRecent(i []int) *OrgUserPreferenceUpdate {
	oupu.mutation.AppendMenuRecent(i)
	return oupu
}

// ClearMenuRecent clears the value of the "menu_recent" field.
func (oupu *OrgUserPreferenceUpdate) ClearMenuRecent() *OrgUserPreferenceUpdate {
	oupu.mutation.ClearMenuRecent()
	return oupu
}

// SetUser sets the "user" edge to the User entity.
func (oupu *OrgUserPreferenceUpdate) SetUser(u *User) *OrgUserPreferenceUpdate {
	return oupu.SetUserID(u.ID)
}

// Mutation returns the OrgUserPreferenceMutation object of the builder.
func (oupu *OrgUserPreferenceUpdate) Mutation() *OrgUserPreferenceMutation {
	return oupu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (oupu *OrgUserPreferenceUpdate) ClearUser() *OrgUserPreferenceUpdate {
	oupu.mutation.ClearUser()
	return oupu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oupu *OrgUserPreferenceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, oupu.sqlSave, oupu.mutation, oupu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oupu *OrgUserPreferenceUpdate) SaveX(ctx context.Context) int {
	affected, err := oupu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oupu *OrgUserPreferenceUpdate) Exec(ctx context.Context) error {
	_, err := oupu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oupu *OrgUserPreferenceUpdate) ExecX(ctx context.Context) {
	if err := oupu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oupu *OrgUserPreferenceUpdate) check() error {
	if _, ok := oupu.mutation.UserID(); oupu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrgUserPreference.user"`)
	}
	if _, ok := oupu.mutation.OrgID(); oupu.mutation.OrgCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrgUserPreference.org"`)
	}
	return nil
}

func (oupu *OrgUserPreferenceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oupu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orguserpreference.Table, orguserpreference.Columns, sqlgraph.NewFieldSpec(orguserpreference.FieldID, field.TypeInt))
	if ps := oupu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oupu.mutation.UpdatedBy(); ok {
		_spec.SetField(orguserpreference.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := oupu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(orguserpreference.FieldUpdatedBy, field.TypeInt, value)
	}
	if oupu.mutation.UpdatedByCleared() {
		_spec.ClearField(orguserpreference.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := oupu.mutation.UpdatedAt(); ok {
		_spec.SetField(orguserpreference.FieldUpdatedAt, field.TypeTime, value)
	}
	if oupu.mutation.UpdatedAtCleared() {
		_spec.ClearField(orguserpreference.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := oupu.mutation.MenuFavorite(); ok {
		_spec.SetField(orguserpreference.FieldMenuFavorite, field.TypeJSON, value)
	}
	if value, ok := oupu.mutation.AppendedMenuFavorite(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, orguserpreference.FieldMenuFavorite, value)
		})
	}
	if oupu.mutation.MenuFavoriteCleared() {
		_spec.ClearField(orguserpreference.FieldMenuFavorite, field.TypeJSON)
	}
	if value, ok := oupu.mutation.MenuRecent(); ok {
		_spec.SetField(orguserpreference.FieldMenuRecent, field.TypeJSON, value)
	}
	if value, ok := oupu.mutation.AppendedMenuRecent(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, orguserpreference.FieldMenuRecent, value)
		})
	}
	if oupu.mutation.MenuRecentCleared() {
		_spec.ClearField(orguserpreference.FieldMenuRecent, field.TypeJSON)
	}
	if oupu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oupu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oupu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orguserpreference.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oupu.mutation.done = true
	return n, nil
}

// OrgUserPreferenceUpdateOne is the builder for updating a single OrgUserPreference entity.
type OrgUserPreferenceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrgUserPreferenceMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (oupuo *OrgUserPreferenceUpdateOne) SetUpdatedBy(i int) *OrgUserPreferenceUpdateOne {
	oupuo.mutation.ResetUpdatedBy()
	oupuo.mutation.SetUpdatedBy(i)
	return oupuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oupuo *OrgUserPreferenceUpdateOne) SetNillableUpdatedBy(i *int) *OrgUserPreferenceUpdateOne {
	if i != nil {
		oupuo.SetUpdatedBy(*i)
	}
	return oupuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (oupuo *OrgUserPreferenceUpdateOne) AddUpdatedBy(i int) *OrgUserPreferenceUpdateOne {
	oupuo.mutation.AddUpdatedBy(i)
	return oupuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (oupuo *OrgUserPreferenceUpdateOne) ClearUpdatedBy() *OrgUserPreferenceUpdateOne {
	oupuo.mutation.ClearUpdatedBy()
	return oupuo
}

// SetUpdatedAt sets the "updated_at" field.
func (oupuo *OrgUserPreferenceUpdateOne) SetUpdatedAt(t time.Time) *OrgUserPreferenceUpdateOne {
	oupuo.mutation.SetUpdatedAt(t)
	return oupuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oupuo *OrgUserPreferenceUpdateOne) SetNillableUpdatedAt(t *time.Time) *OrgUserPreferenceUpdateOne {
	if t != nil {
		oupuo.SetUpdatedAt(*t)
	}
	return oupuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (oupuo *OrgUserPreferenceUpdateOne) ClearUpdatedAt() *OrgUserPreferenceUpdateOne {
	oupuo.mutation.ClearUpdatedAt()
	return oupuo
}

// SetUserID sets the "user_id" field.
func (oupuo *OrgUserPreferenceUpdateOne) SetUserID(i int) *OrgUserPreferenceUpdateOne {
	oupuo.mutation.SetUserID(i)
	return oupuo
}

// SetMenuFavorite sets the "menu_favorite" field.
func (oupuo *OrgUserPreferenceUpdateOne) SetMenuFavorite(i []int) *OrgUserPreferenceUpdateOne {
	oupuo.mutation.SetMenuFavorite(i)
	return oupuo
}

// AppendMenuFavorite appends i to the "menu_favorite" field.
func (oupuo *OrgUserPreferenceUpdateOne) AppendMenuFavorite(i []int) *OrgUserPreferenceUpdateOne {
	oupuo.mutation.AppendMenuFavorite(i)
	return oupuo
}

// ClearMenuFavorite clears the value of the "menu_favorite" field.
func (oupuo *OrgUserPreferenceUpdateOne) ClearMenuFavorite() *OrgUserPreferenceUpdateOne {
	oupuo.mutation.ClearMenuFavorite()
	return oupuo
}

// SetMenuRecent sets the "menu_recent" field.
func (oupuo *OrgUserPreferenceUpdateOne) SetMenuRecent(i []int) *OrgUserPreferenceUpdateOne {
	oupuo.mutation.SetMenuRecent(i)
	return oupuo
}

// AppendMenuRecent appends i to the "menu_recent" field.
func (oupuo *OrgUserPreferenceUpdateOne) AppendMenuRecent(i []int) *OrgUserPreferenceUpdateOne {
	oupuo.mutation.AppendMenuRecent(i)
	return oupuo
}

// ClearMenuRecent clears the value of the "menu_recent" field.
func (oupuo *OrgUserPreferenceUpdateOne) ClearMenuRecent() *OrgUserPreferenceUpdateOne {
	oupuo.mutation.ClearMenuRecent()
	return oupuo
}

// SetUser sets the "user" edge to the User entity.
func (oupuo *OrgUserPreferenceUpdateOne) SetUser(u *User) *OrgUserPreferenceUpdateOne {
	return oupuo.SetUserID(u.ID)
}

// Mutation returns the OrgUserPreferenceMutation object of the builder.
func (oupuo *OrgUserPreferenceUpdateOne) Mutation() *OrgUserPreferenceMutation {
	return oupuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (oupuo *OrgUserPreferenceUpdateOne) ClearUser() *OrgUserPreferenceUpdateOne {
	oupuo.mutation.ClearUser()
	return oupuo
}

// Where appends a list predicates to the OrgUserPreferenceUpdate builder.
func (oupuo *OrgUserPreferenceUpdateOne) Where(ps ...predicate.OrgUserPreference) *OrgUserPreferenceUpdateOne {
	oupuo.mutation.Where(ps...)
	return oupuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oupuo *OrgUserPreferenceUpdateOne) Select(field string, fields ...string) *OrgUserPreferenceUpdateOne {
	oupuo.fields = append([]string{field}, fields...)
	return oupuo
}

// Save executes the query and returns the updated OrgUserPreference entity.
func (oupuo *OrgUserPreferenceUpdateOne) Save(ctx context.Context) (*OrgUserPreference, error) {
	return withHooks(ctx, oupuo.sqlSave, oupuo.mutation, oupuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oupuo *OrgUserPreferenceUpdateOne) SaveX(ctx context.Context) *OrgUserPreference {
	node, err := oupuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oupuo *OrgUserPreferenceUpdateOne) Exec(ctx context.Context) error {
	_, err := oupuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oupuo *OrgUserPreferenceUpdateOne) ExecX(ctx context.Context) {
	if err := oupuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oupuo *OrgUserPreferenceUpdateOne) check() error {
	if _, ok := oupuo.mutation.UserID(); oupuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrgUserPreference.user"`)
	}
	if _, ok := oupuo.mutation.OrgID(); oupuo.mutation.OrgCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "OrgUserPreference.org"`)
	}
	return nil
}

func (oupuo *OrgUserPreferenceUpdateOne) sqlSave(ctx context.Context) (_node *OrgUserPreference, err error) {
	if err := oupuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orguserpreference.Table, orguserpreference.Columns, sqlgraph.NewFieldSpec(orguserpreference.FieldID, field.TypeInt))
	id, ok := oupuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrgUserPreference.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oupuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orguserpreference.FieldID)
		for _, f := range fields {
			if !orguserpreference.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orguserpreference.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oupuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oupuo.mutation.UpdatedBy(); ok {
		_spec.SetField(orguserpreference.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := oupuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(orguserpreference.FieldUpdatedBy, field.TypeInt, value)
	}
	if oupuo.mutation.UpdatedByCleared() {
		_spec.ClearField(orguserpreference.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := oupuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orguserpreference.FieldUpdatedAt, field.TypeTime, value)
	}
	if oupuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(orguserpreference.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := oupuo.mutation.MenuFavorite(); ok {
		_spec.SetField(orguserpreference.FieldMenuFavorite, field.TypeJSON, value)
	}
	if value, ok := oupuo.mutation.AppendedMenuFavorite(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, orguserpreference.FieldMenuFavorite, value)
		})
	}
	if oupuo.mutation.MenuFavoriteCleared() {
		_spec.ClearField(orguserpreference.FieldMenuFavorite, field.TypeJSON)
	}
	if value, ok := oupuo.mutation.MenuRecent(); ok {
		_spec.SetField(orguserpreference.FieldMenuRecent, field.TypeJSON, value)
	}
	if value, ok := oupuo.mutation.AppendedMenuRecent(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, orguserpreference.FieldMenuRecent, value)
		})
	}
	if oupuo.mutation.MenuRecentCleared() {
		_spec.ClearField(orguserpreference.FieldMenuRecent, field.TypeJSON)
	}
	if oupuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oupuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrgUserPreference{config: oupuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oupuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orguserpreference.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oupuo.mutation.done = true
	return _node, nil
}
