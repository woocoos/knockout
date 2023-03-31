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
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/approle"
	"github.com/woocoos/knockout/ent/approlepolicy"
	"github.com/woocoos/knockout/ent/predicate"
)

// AppRolePolicyUpdate is the builder for updating AppRolePolicy entities.
type AppRolePolicyUpdate struct {
	config
	hooks    []Hook
	mutation *AppRolePolicyMutation
}

// Where appends a list predicates to the AppRolePolicyUpdate builder.
func (arpu *AppRolePolicyUpdate) Where(ps ...predicate.AppRolePolicy) *AppRolePolicyUpdate {
	arpu.mutation.Where(ps...)
	return arpu
}

// SetUpdatedBy sets the "updated_by" field.
func (arpu *AppRolePolicyUpdate) SetUpdatedBy(i int) *AppRolePolicyUpdate {
	arpu.mutation.ResetUpdatedBy()
	arpu.mutation.SetUpdatedBy(i)
	return arpu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (arpu *AppRolePolicyUpdate) SetNillableUpdatedBy(i *int) *AppRolePolicyUpdate {
	if i != nil {
		arpu.SetUpdatedBy(*i)
	}
	return arpu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (arpu *AppRolePolicyUpdate) AddUpdatedBy(i int) *AppRolePolicyUpdate {
	arpu.mutation.AddUpdatedBy(i)
	return arpu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (arpu *AppRolePolicyUpdate) ClearUpdatedBy() *AppRolePolicyUpdate {
	arpu.mutation.ClearUpdatedBy()
	return arpu
}

// SetUpdatedAt sets the "updated_at" field.
func (arpu *AppRolePolicyUpdate) SetUpdatedAt(t time.Time) *AppRolePolicyUpdate {
	arpu.mutation.SetUpdatedAt(t)
	return arpu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (arpu *AppRolePolicyUpdate) SetNillableUpdatedAt(t *time.Time) *AppRolePolicyUpdate {
	if t != nil {
		arpu.SetUpdatedAt(*t)
	}
	return arpu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (arpu *AppRolePolicyUpdate) ClearUpdatedAt() *AppRolePolicyUpdate {
	arpu.mutation.ClearUpdatedAt()
	return arpu
}

// SetAppRoleID sets the "app_role_id" field.
func (arpu *AppRolePolicyUpdate) SetAppRoleID(i int) *AppRolePolicyUpdate {
	arpu.mutation.SetAppRoleID(i)
	return arpu
}

// SetAppPolicyID sets the "app_policy_id" field.
func (arpu *AppRolePolicyUpdate) SetAppPolicyID(i int) *AppRolePolicyUpdate {
	arpu.mutation.SetAppPolicyID(i)
	return arpu
}

// SetAppID sets the "app_id" field.
func (arpu *AppRolePolicyUpdate) SetAppID(i int) *AppRolePolicyUpdate {
	arpu.mutation.ResetAppID()
	arpu.mutation.SetAppID(i)
	return arpu
}

// AddAppID adds i to the "app_id" field.
func (arpu *AppRolePolicyUpdate) AddAppID(i int) *AppRolePolicyUpdate {
	arpu.mutation.AddAppID(i)
	return arpu
}

// SetRoleID sets the "role" edge to the AppRole entity by ID.
func (arpu *AppRolePolicyUpdate) SetRoleID(id int) *AppRolePolicyUpdate {
	arpu.mutation.SetRoleID(id)
	return arpu
}

// SetRole sets the "role" edge to the AppRole entity.
func (arpu *AppRolePolicyUpdate) SetRole(a *AppRole) *AppRolePolicyUpdate {
	return arpu.SetRoleID(a.ID)
}

// SetPolicyID sets the "policy" edge to the AppPolicy entity by ID.
func (arpu *AppRolePolicyUpdate) SetPolicyID(id int) *AppRolePolicyUpdate {
	arpu.mutation.SetPolicyID(id)
	return arpu
}

// SetPolicy sets the "policy" edge to the AppPolicy entity.
func (arpu *AppRolePolicyUpdate) SetPolicy(a *AppPolicy) *AppRolePolicyUpdate {
	return arpu.SetPolicyID(a.ID)
}

// Mutation returns the AppRolePolicyMutation object of the builder.
func (arpu *AppRolePolicyUpdate) Mutation() *AppRolePolicyMutation {
	return arpu.mutation
}

// ClearRole clears the "role" edge to the AppRole entity.
func (arpu *AppRolePolicyUpdate) ClearRole() *AppRolePolicyUpdate {
	arpu.mutation.ClearRole()
	return arpu
}

// ClearPolicy clears the "policy" edge to the AppPolicy entity.
func (arpu *AppRolePolicyUpdate) ClearPolicy() *AppRolePolicyUpdate {
	arpu.mutation.ClearPolicy()
	return arpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (arpu *AppRolePolicyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AppRolePolicyMutation](ctx, arpu.sqlSave, arpu.mutation, arpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (arpu *AppRolePolicyUpdate) SaveX(ctx context.Context) int {
	affected, err := arpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (arpu *AppRolePolicyUpdate) Exec(ctx context.Context) error {
	_, err := arpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arpu *AppRolePolicyUpdate) ExecX(ctx context.Context) {
	if err := arpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arpu *AppRolePolicyUpdate) check() error {
	if _, ok := arpu.mutation.RoleID(); arpu.mutation.RoleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AppRolePolicy.role"`)
	}
	if _, ok := arpu.mutation.PolicyID(); arpu.mutation.PolicyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AppRolePolicy.policy"`)
	}
	return nil
}

func (arpu *AppRolePolicyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := arpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(approlepolicy.Table, approlepolicy.Columns, sqlgraph.NewFieldSpec(approlepolicy.FieldID, field.TypeInt))
	if ps := arpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := arpu.mutation.UpdatedBy(); ok {
		_spec.SetField(approlepolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := arpu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(approlepolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if arpu.mutation.UpdatedByCleared() {
		_spec.ClearField(approlepolicy.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := arpu.mutation.UpdatedAt(); ok {
		_spec.SetField(approlepolicy.FieldUpdatedAt, field.TypeTime, value)
	}
	if arpu.mutation.UpdatedAtCleared() {
		_spec.ClearField(approlepolicy.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := arpu.mutation.AppID(); ok {
		_spec.SetField(approlepolicy.FieldAppID, field.TypeInt, value)
	}
	if value, ok := arpu.mutation.AddedAppID(); ok {
		_spec.AddField(approlepolicy.FieldAppID, field.TypeInt, value)
	}
	if arpu.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approlepolicy.RoleTable,
			Columns: []string{approlepolicy.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approle.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := arpu.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approlepolicy.RoleTable,
			Columns: []string{approlepolicy.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if arpu.mutation.PolicyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approlepolicy.PolicyTable,
			Columns: []string{approlepolicy.PolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := arpu.mutation.PolicyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approlepolicy.PolicyTable,
			Columns: []string{approlepolicy.PolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, arpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{approlepolicy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	arpu.mutation.done = true
	return n, nil
}

// AppRolePolicyUpdateOne is the builder for updating a single AppRolePolicy entity.
type AppRolePolicyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppRolePolicyMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (arpuo *AppRolePolicyUpdateOne) SetUpdatedBy(i int) *AppRolePolicyUpdateOne {
	arpuo.mutation.ResetUpdatedBy()
	arpuo.mutation.SetUpdatedBy(i)
	return arpuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (arpuo *AppRolePolicyUpdateOne) SetNillableUpdatedBy(i *int) *AppRolePolicyUpdateOne {
	if i != nil {
		arpuo.SetUpdatedBy(*i)
	}
	return arpuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (arpuo *AppRolePolicyUpdateOne) AddUpdatedBy(i int) *AppRolePolicyUpdateOne {
	arpuo.mutation.AddUpdatedBy(i)
	return arpuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (arpuo *AppRolePolicyUpdateOne) ClearUpdatedBy() *AppRolePolicyUpdateOne {
	arpuo.mutation.ClearUpdatedBy()
	return arpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (arpuo *AppRolePolicyUpdateOne) SetUpdatedAt(t time.Time) *AppRolePolicyUpdateOne {
	arpuo.mutation.SetUpdatedAt(t)
	return arpuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (arpuo *AppRolePolicyUpdateOne) SetNillableUpdatedAt(t *time.Time) *AppRolePolicyUpdateOne {
	if t != nil {
		arpuo.SetUpdatedAt(*t)
	}
	return arpuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (arpuo *AppRolePolicyUpdateOne) ClearUpdatedAt() *AppRolePolicyUpdateOne {
	arpuo.mutation.ClearUpdatedAt()
	return arpuo
}

// SetAppRoleID sets the "app_role_id" field.
func (arpuo *AppRolePolicyUpdateOne) SetAppRoleID(i int) *AppRolePolicyUpdateOne {
	arpuo.mutation.SetAppRoleID(i)
	return arpuo
}

// SetAppPolicyID sets the "app_policy_id" field.
func (arpuo *AppRolePolicyUpdateOne) SetAppPolicyID(i int) *AppRolePolicyUpdateOne {
	arpuo.mutation.SetAppPolicyID(i)
	return arpuo
}

// SetAppID sets the "app_id" field.
func (arpuo *AppRolePolicyUpdateOne) SetAppID(i int) *AppRolePolicyUpdateOne {
	arpuo.mutation.ResetAppID()
	arpuo.mutation.SetAppID(i)
	return arpuo
}

// AddAppID adds i to the "app_id" field.
func (arpuo *AppRolePolicyUpdateOne) AddAppID(i int) *AppRolePolicyUpdateOne {
	arpuo.mutation.AddAppID(i)
	return arpuo
}

// SetRoleID sets the "role" edge to the AppRole entity by ID.
func (arpuo *AppRolePolicyUpdateOne) SetRoleID(id int) *AppRolePolicyUpdateOne {
	arpuo.mutation.SetRoleID(id)
	return arpuo
}

// SetRole sets the "role" edge to the AppRole entity.
func (arpuo *AppRolePolicyUpdateOne) SetRole(a *AppRole) *AppRolePolicyUpdateOne {
	return arpuo.SetRoleID(a.ID)
}

// SetPolicyID sets the "policy" edge to the AppPolicy entity by ID.
func (arpuo *AppRolePolicyUpdateOne) SetPolicyID(id int) *AppRolePolicyUpdateOne {
	arpuo.mutation.SetPolicyID(id)
	return arpuo
}

// SetPolicy sets the "policy" edge to the AppPolicy entity.
func (arpuo *AppRolePolicyUpdateOne) SetPolicy(a *AppPolicy) *AppRolePolicyUpdateOne {
	return arpuo.SetPolicyID(a.ID)
}

// Mutation returns the AppRolePolicyMutation object of the builder.
func (arpuo *AppRolePolicyUpdateOne) Mutation() *AppRolePolicyMutation {
	return arpuo.mutation
}

// ClearRole clears the "role" edge to the AppRole entity.
func (arpuo *AppRolePolicyUpdateOne) ClearRole() *AppRolePolicyUpdateOne {
	arpuo.mutation.ClearRole()
	return arpuo
}

// ClearPolicy clears the "policy" edge to the AppPolicy entity.
func (arpuo *AppRolePolicyUpdateOne) ClearPolicy() *AppRolePolicyUpdateOne {
	arpuo.mutation.ClearPolicy()
	return arpuo
}

// Where appends a list predicates to the AppRolePolicyUpdate builder.
func (arpuo *AppRolePolicyUpdateOne) Where(ps ...predicate.AppRolePolicy) *AppRolePolicyUpdateOne {
	arpuo.mutation.Where(ps...)
	return arpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (arpuo *AppRolePolicyUpdateOne) Select(field string, fields ...string) *AppRolePolicyUpdateOne {
	arpuo.fields = append([]string{field}, fields...)
	return arpuo
}

// Save executes the query and returns the updated AppRolePolicy entity.
func (arpuo *AppRolePolicyUpdateOne) Save(ctx context.Context) (*AppRolePolicy, error) {
	return withHooks[*AppRolePolicy, AppRolePolicyMutation](ctx, arpuo.sqlSave, arpuo.mutation, arpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (arpuo *AppRolePolicyUpdateOne) SaveX(ctx context.Context) *AppRolePolicy {
	node, err := arpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (arpuo *AppRolePolicyUpdateOne) Exec(ctx context.Context) error {
	_, err := arpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arpuo *AppRolePolicyUpdateOne) ExecX(ctx context.Context) {
	if err := arpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (arpuo *AppRolePolicyUpdateOne) check() error {
	if _, ok := arpuo.mutation.RoleID(); arpuo.mutation.RoleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AppRolePolicy.role"`)
	}
	if _, ok := arpuo.mutation.PolicyID(); arpuo.mutation.PolicyCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AppRolePolicy.policy"`)
	}
	return nil
}

func (arpuo *AppRolePolicyUpdateOne) sqlSave(ctx context.Context) (_node *AppRolePolicy, err error) {
	if err := arpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(approlepolicy.Table, approlepolicy.Columns, sqlgraph.NewFieldSpec(approlepolicy.FieldID, field.TypeInt))
	id, ok := arpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppRolePolicy.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := arpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, approlepolicy.FieldID)
		for _, f := range fields {
			if !approlepolicy.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != approlepolicy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := arpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := arpuo.mutation.UpdatedBy(); ok {
		_spec.SetField(approlepolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := arpuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(approlepolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if arpuo.mutation.UpdatedByCleared() {
		_spec.ClearField(approlepolicy.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := arpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(approlepolicy.FieldUpdatedAt, field.TypeTime, value)
	}
	if arpuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(approlepolicy.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := arpuo.mutation.AppID(); ok {
		_spec.SetField(approlepolicy.FieldAppID, field.TypeInt, value)
	}
	if value, ok := arpuo.mutation.AddedAppID(); ok {
		_spec.AddField(approlepolicy.FieldAppID, field.TypeInt, value)
	}
	if arpuo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approlepolicy.RoleTable,
			Columns: []string{approlepolicy.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approle.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := arpuo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approlepolicy.RoleTable,
			Columns: []string{approlepolicy.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if arpuo.mutation.PolicyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approlepolicy.PolicyTable,
			Columns: []string{approlepolicy.PolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := arpuo.mutation.PolicyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   approlepolicy.PolicyTable,
			Columns: []string{approlepolicy.PolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AppRolePolicy{config: arpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, arpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{approlepolicy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	arpuo.mutation.done = true
	return _node, nil
}
