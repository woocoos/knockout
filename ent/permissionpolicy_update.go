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
	"github.com/woocoos/knockout/ent/organization"
	"github.com/woocoos/knockout/ent/permissionpolicy"
	"github.com/woocoos/knockout/ent/predicate"
	"github.com/woocoos/knockout/graph/entgen/types"
)

// PermissionPolicyUpdate is the builder for updating PermissionPolicy entities.
type PermissionPolicyUpdate struct {
	config
	hooks    []Hook
	mutation *PermissionPolicyMutation
}

// Where appends a list predicates to the PermissionPolicyUpdate builder.
func (ppu *PermissionPolicyUpdate) Where(ps ...predicate.PermissionPolicy) *PermissionPolicyUpdate {
	ppu.mutation.Where(ps...)
	return ppu
}

// SetUpdatedBy sets the "updated_by" field.
func (ppu *PermissionPolicyUpdate) SetUpdatedBy(i int) *PermissionPolicyUpdate {
	ppu.mutation.ResetUpdatedBy()
	ppu.mutation.SetUpdatedBy(i)
	return ppu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ppu *PermissionPolicyUpdate) SetNillableUpdatedBy(i *int) *PermissionPolicyUpdate {
	if i != nil {
		ppu.SetUpdatedBy(*i)
	}
	return ppu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ppu *PermissionPolicyUpdate) AddUpdatedBy(i int) *PermissionPolicyUpdate {
	ppu.mutation.AddUpdatedBy(i)
	return ppu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ppu *PermissionPolicyUpdate) ClearUpdatedBy() *PermissionPolicyUpdate {
	ppu.mutation.ClearUpdatedBy()
	return ppu
}

// SetUpdatedAt sets the "updated_at" field.
func (ppu *PermissionPolicyUpdate) SetUpdatedAt(t time.Time) *PermissionPolicyUpdate {
	ppu.mutation.SetUpdatedAt(t)
	return ppu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ppu *PermissionPolicyUpdate) SetNillableUpdatedAt(t *time.Time) *PermissionPolicyUpdate {
	if t != nil {
		ppu.SetUpdatedAt(*t)
	}
	return ppu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ppu *PermissionPolicyUpdate) ClearUpdatedAt() *PermissionPolicyUpdate {
	ppu.mutation.ClearUpdatedAt()
	return ppu
}

// SetOrgID sets the "org_id" field.
func (ppu *PermissionPolicyUpdate) SetOrgID(i int) *PermissionPolicyUpdate {
	ppu.mutation.SetOrgID(i)
	return ppu
}

// SetAppID sets the "app_id" field.
func (ppu *PermissionPolicyUpdate) SetAppID(i int) *PermissionPolicyUpdate {
	ppu.mutation.ResetAppID()
	ppu.mutation.SetAppID(i)
	return ppu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ppu *PermissionPolicyUpdate) SetNillableAppID(i *int) *PermissionPolicyUpdate {
	if i != nil {
		ppu.SetAppID(*i)
	}
	return ppu
}

// AddAppID adds i to the "app_id" field.
func (ppu *PermissionPolicyUpdate) AddAppID(i int) *PermissionPolicyUpdate {
	ppu.mutation.AddAppID(i)
	return ppu
}

// ClearAppID clears the value of the "app_id" field.
func (ppu *PermissionPolicyUpdate) ClearAppID() *PermissionPolicyUpdate {
	ppu.mutation.ClearAppID()
	return ppu
}

// SetAppPolicyID sets the "app_policy_id" field.
func (ppu *PermissionPolicyUpdate) SetAppPolicyID(i int) *PermissionPolicyUpdate {
	ppu.mutation.ResetAppPolicyID()
	ppu.mutation.SetAppPolicyID(i)
	return ppu
}

// SetNillableAppPolicyID sets the "app_policy_id" field if the given value is not nil.
func (ppu *PermissionPolicyUpdate) SetNillableAppPolicyID(i *int) *PermissionPolicyUpdate {
	if i != nil {
		ppu.SetAppPolicyID(*i)
	}
	return ppu
}

// AddAppPolicyID adds i to the "app_policy_id" field.
func (ppu *PermissionPolicyUpdate) AddAppPolicyID(i int) *PermissionPolicyUpdate {
	ppu.mutation.AddAppPolicyID(i)
	return ppu
}

// ClearAppPolicyID clears the value of the "app_policy_id" field.
func (ppu *PermissionPolicyUpdate) ClearAppPolicyID() *PermissionPolicyUpdate {
	ppu.mutation.ClearAppPolicyID()
	return ppu
}

// SetName sets the "name" field.
func (ppu *PermissionPolicyUpdate) SetName(s string) *PermissionPolicyUpdate {
	ppu.mutation.SetName(s)
	return ppu
}

// SetComments sets the "comments" field.
func (ppu *PermissionPolicyUpdate) SetComments(s string) *PermissionPolicyUpdate {
	ppu.mutation.SetComments(s)
	return ppu
}

// SetRules sets the "rules" field.
func (ppu *PermissionPolicyUpdate) SetRules(tr []types.PolicyRule) *PermissionPolicyUpdate {
	ppu.mutation.SetRules(tr)
	return ppu
}

// AppendRules appends tr to the "rules" field.
func (ppu *PermissionPolicyUpdate) AppendRules(tr []types.PolicyRule) *PermissionPolicyUpdate {
	ppu.mutation.AppendRules(tr)
	return ppu
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (ppu *PermissionPolicyUpdate) SetOrganizationID(id int) *PermissionPolicyUpdate {
	ppu.mutation.SetOrganizationID(id)
	return ppu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (ppu *PermissionPolicyUpdate) SetOrganization(o *Organization) *PermissionPolicyUpdate {
	return ppu.SetOrganizationID(o.ID)
}

// Mutation returns the PermissionPolicyMutation object of the builder.
func (ppu *PermissionPolicyUpdate) Mutation() *PermissionPolicyMutation {
	return ppu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (ppu *PermissionPolicyUpdate) ClearOrganization() *PermissionPolicyUpdate {
	ppu.mutation.ClearOrganization()
	return ppu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ppu *PermissionPolicyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PermissionPolicyMutation](ctx, ppu.sqlSave, ppu.mutation, ppu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ppu *PermissionPolicyUpdate) SaveX(ctx context.Context) int {
	affected, err := ppu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ppu *PermissionPolicyUpdate) Exec(ctx context.Context) error {
	_, err := ppu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppu *PermissionPolicyUpdate) ExecX(ctx context.Context) {
	if err := ppu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ppu *PermissionPolicyUpdate) check() error {
	if _, ok := ppu.mutation.OrganizationID(); ppu.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PermissionPolicy.organization"`)
	}
	return nil
}

func (ppu *PermissionPolicyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ppu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(permissionpolicy.Table, permissionpolicy.Columns, sqlgraph.NewFieldSpec(permissionpolicy.FieldID, field.TypeInt))
	if ps := ppu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ppu.mutation.UpdatedBy(); ok {
		_spec.SetField(permissionpolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ppu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(permissionpolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if ppu.mutation.UpdatedByCleared() {
		_spec.ClearField(permissionpolicy.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := ppu.mutation.UpdatedAt(); ok {
		_spec.SetField(permissionpolicy.FieldUpdatedAt, field.TypeTime, value)
	}
	if ppu.mutation.UpdatedAtCleared() {
		_spec.ClearField(permissionpolicy.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ppu.mutation.AppID(); ok {
		_spec.SetField(permissionpolicy.FieldAppID, field.TypeInt, value)
	}
	if value, ok := ppu.mutation.AddedAppID(); ok {
		_spec.AddField(permissionpolicy.FieldAppID, field.TypeInt, value)
	}
	if ppu.mutation.AppIDCleared() {
		_spec.ClearField(permissionpolicy.FieldAppID, field.TypeInt)
	}
	if value, ok := ppu.mutation.AppPolicyID(); ok {
		_spec.SetField(permissionpolicy.FieldAppPolicyID, field.TypeInt, value)
	}
	if value, ok := ppu.mutation.AddedAppPolicyID(); ok {
		_spec.AddField(permissionpolicy.FieldAppPolicyID, field.TypeInt, value)
	}
	if ppu.mutation.AppPolicyIDCleared() {
		_spec.ClearField(permissionpolicy.FieldAppPolicyID, field.TypeInt)
	}
	if value, ok := ppu.mutation.Name(); ok {
		_spec.SetField(permissionpolicy.FieldName, field.TypeString, value)
	}
	if value, ok := ppu.mutation.Comments(); ok {
		_spec.SetField(permissionpolicy.FieldComments, field.TypeString, value)
	}
	if value, ok := ppu.mutation.Rules(); ok {
		_spec.SetField(permissionpolicy.FieldRules, field.TypeJSON, value)
	}
	if value, ok := ppu.mutation.AppendedRules(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, permissionpolicy.FieldRules, value)
		})
	}
	if ppu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permissionpolicy.OrganizationTable,
			Columns: []string{permissionpolicy.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permissionpolicy.OrganizationTable,
			Columns: []string{permissionpolicy.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ppu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permissionpolicy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ppu.mutation.done = true
	return n, nil
}

// PermissionPolicyUpdateOne is the builder for updating a single PermissionPolicy entity.
type PermissionPolicyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PermissionPolicyMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (ppuo *PermissionPolicyUpdateOne) SetUpdatedBy(i int) *PermissionPolicyUpdateOne {
	ppuo.mutation.ResetUpdatedBy()
	ppuo.mutation.SetUpdatedBy(i)
	return ppuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ppuo *PermissionPolicyUpdateOne) SetNillableUpdatedBy(i *int) *PermissionPolicyUpdateOne {
	if i != nil {
		ppuo.SetUpdatedBy(*i)
	}
	return ppuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ppuo *PermissionPolicyUpdateOne) AddUpdatedBy(i int) *PermissionPolicyUpdateOne {
	ppuo.mutation.AddUpdatedBy(i)
	return ppuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ppuo *PermissionPolicyUpdateOne) ClearUpdatedBy() *PermissionPolicyUpdateOne {
	ppuo.mutation.ClearUpdatedBy()
	return ppuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ppuo *PermissionPolicyUpdateOne) SetUpdatedAt(t time.Time) *PermissionPolicyUpdateOne {
	ppuo.mutation.SetUpdatedAt(t)
	return ppuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ppuo *PermissionPolicyUpdateOne) SetNillableUpdatedAt(t *time.Time) *PermissionPolicyUpdateOne {
	if t != nil {
		ppuo.SetUpdatedAt(*t)
	}
	return ppuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ppuo *PermissionPolicyUpdateOne) ClearUpdatedAt() *PermissionPolicyUpdateOne {
	ppuo.mutation.ClearUpdatedAt()
	return ppuo
}

// SetOrgID sets the "org_id" field.
func (ppuo *PermissionPolicyUpdateOne) SetOrgID(i int) *PermissionPolicyUpdateOne {
	ppuo.mutation.SetOrgID(i)
	return ppuo
}

// SetAppID sets the "app_id" field.
func (ppuo *PermissionPolicyUpdateOne) SetAppID(i int) *PermissionPolicyUpdateOne {
	ppuo.mutation.ResetAppID()
	ppuo.mutation.SetAppID(i)
	return ppuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (ppuo *PermissionPolicyUpdateOne) SetNillableAppID(i *int) *PermissionPolicyUpdateOne {
	if i != nil {
		ppuo.SetAppID(*i)
	}
	return ppuo
}

// AddAppID adds i to the "app_id" field.
func (ppuo *PermissionPolicyUpdateOne) AddAppID(i int) *PermissionPolicyUpdateOne {
	ppuo.mutation.AddAppID(i)
	return ppuo
}

// ClearAppID clears the value of the "app_id" field.
func (ppuo *PermissionPolicyUpdateOne) ClearAppID() *PermissionPolicyUpdateOne {
	ppuo.mutation.ClearAppID()
	return ppuo
}

// SetAppPolicyID sets the "app_policy_id" field.
func (ppuo *PermissionPolicyUpdateOne) SetAppPolicyID(i int) *PermissionPolicyUpdateOne {
	ppuo.mutation.ResetAppPolicyID()
	ppuo.mutation.SetAppPolicyID(i)
	return ppuo
}

// SetNillableAppPolicyID sets the "app_policy_id" field if the given value is not nil.
func (ppuo *PermissionPolicyUpdateOne) SetNillableAppPolicyID(i *int) *PermissionPolicyUpdateOne {
	if i != nil {
		ppuo.SetAppPolicyID(*i)
	}
	return ppuo
}

// AddAppPolicyID adds i to the "app_policy_id" field.
func (ppuo *PermissionPolicyUpdateOne) AddAppPolicyID(i int) *PermissionPolicyUpdateOne {
	ppuo.mutation.AddAppPolicyID(i)
	return ppuo
}

// ClearAppPolicyID clears the value of the "app_policy_id" field.
func (ppuo *PermissionPolicyUpdateOne) ClearAppPolicyID() *PermissionPolicyUpdateOne {
	ppuo.mutation.ClearAppPolicyID()
	return ppuo
}

// SetName sets the "name" field.
func (ppuo *PermissionPolicyUpdateOne) SetName(s string) *PermissionPolicyUpdateOne {
	ppuo.mutation.SetName(s)
	return ppuo
}

// SetComments sets the "comments" field.
func (ppuo *PermissionPolicyUpdateOne) SetComments(s string) *PermissionPolicyUpdateOne {
	ppuo.mutation.SetComments(s)
	return ppuo
}

// SetRules sets the "rules" field.
func (ppuo *PermissionPolicyUpdateOne) SetRules(tr []types.PolicyRule) *PermissionPolicyUpdateOne {
	ppuo.mutation.SetRules(tr)
	return ppuo
}

// AppendRules appends tr to the "rules" field.
func (ppuo *PermissionPolicyUpdateOne) AppendRules(tr []types.PolicyRule) *PermissionPolicyUpdateOne {
	ppuo.mutation.AppendRules(tr)
	return ppuo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (ppuo *PermissionPolicyUpdateOne) SetOrganizationID(id int) *PermissionPolicyUpdateOne {
	ppuo.mutation.SetOrganizationID(id)
	return ppuo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (ppuo *PermissionPolicyUpdateOne) SetOrganization(o *Organization) *PermissionPolicyUpdateOne {
	return ppuo.SetOrganizationID(o.ID)
}

// Mutation returns the PermissionPolicyMutation object of the builder.
func (ppuo *PermissionPolicyUpdateOne) Mutation() *PermissionPolicyMutation {
	return ppuo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (ppuo *PermissionPolicyUpdateOne) ClearOrganization() *PermissionPolicyUpdateOne {
	ppuo.mutation.ClearOrganization()
	return ppuo
}

// Where appends a list predicates to the PermissionPolicyUpdate builder.
func (ppuo *PermissionPolicyUpdateOne) Where(ps ...predicate.PermissionPolicy) *PermissionPolicyUpdateOne {
	ppuo.mutation.Where(ps...)
	return ppuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ppuo *PermissionPolicyUpdateOne) Select(field string, fields ...string) *PermissionPolicyUpdateOne {
	ppuo.fields = append([]string{field}, fields...)
	return ppuo
}

// Save executes the query and returns the updated PermissionPolicy entity.
func (ppuo *PermissionPolicyUpdateOne) Save(ctx context.Context) (*PermissionPolicy, error) {
	return withHooks[*PermissionPolicy, PermissionPolicyMutation](ctx, ppuo.sqlSave, ppuo.mutation, ppuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ppuo *PermissionPolicyUpdateOne) SaveX(ctx context.Context) *PermissionPolicy {
	node, err := ppuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ppuo *PermissionPolicyUpdateOne) Exec(ctx context.Context) error {
	_, err := ppuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppuo *PermissionPolicyUpdateOne) ExecX(ctx context.Context) {
	if err := ppuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ppuo *PermissionPolicyUpdateOne) check() error {
	if _, ok := ppuo.mutation.OrganizationID(); ppuo.mutation.OrganizationCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PermissionPolicy.organization"`)
	}
	return nil
}

func (ppuo *PermissionPolicyUpdateOne) sqlSave(ctx context.Context) (_node *PermissionPolicy, err error) {
	if err := ppuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(permissionpolicy.Table, permissionpolicy.Columns, sqlgraph.NewFieldSpec(permissionpolicy.FieldID, field.TypeInt))
	id, ok := ppuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PermissionPolicy.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ppuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permissionpolicy.FieldID)
		for _, f := range fields {
			if !permissionpolicy.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permissionpolicy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ppuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ppuo.mutation.UpdatedBy(); ok {
		_spec.SetField(permissionpolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ppuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(permissionpolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if ppuo.mutation.UpdatedByCleared() {
		_spec.ClearField(permissionpolicy.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := ppuo.mutation.UpdatedAt(); ok {
		_spec.SetField(permissionpolicy.FieldUpdatedAt, field.TypeTime, value)
	}
	if ppuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(permissionpolicy.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ppuo.mutation.AppID(); ok {
		_spec.SetField(permissionpolicy.FieldAppID, field.TypeInt, value)
	}
	if value, ok := ppuo.mutation.AddedAppID(); ok {
		_spec.AddField(permissionpolicy.FieldAppID, field.TypeInt, value)
	}
	if ppuo.mutation.AppIDCleared() {
		_spec.ClearField(permissionpolicy.FieldAppID, field.TypeInt)
	}
	if value, ok := ppuo.mutation.AppPolicyID(); ok {
		_spec.SetField(permissionpolicy.FieldAppPolicyID, field.TypeInt, value)
	}
	if value, ok := ppuo.mutation.AddedAppPolicyID(); ok {
		_spec.AddField(permissionpolicy.FieldAppPolicyID, field.TypeInt, value)
	}
	if ppuo.mutation.AppPolicyIDCleared() {
		_spec.ClearField(permissionpolicy.FieldAppPolicyID, field.TypeInt)
	}
	if value, ok := ppuo.mutation.Name(); ok {
		_spec.SetField(permissionpolicy.FieldName, field.TypeString, value)
	}
	if value, ok := ppuo.mutation.Comments(); ok {
		_spec.SetField(permissionpolicy.FieldComments, field.TypeString, value)
	}
	if value, ok := ppuo.mutation.Rules(); ok {
		_spec.SetField(permissionpolicy.FieldRules, field.TypeJSON, value)
	}
	if value, ok := ppuo.mutation.AppendedRules(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, permissionpolicy.FieldRules, value)
		})
	}
	if ppuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permissionpolicy.OrganizationTable,
			Columns: []string{permissionpolicy.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permissionpolicy.OrganizationTable,
			Columns: []string{permissionpolicy.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PermissionPolicy{config: ppuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ppuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permissionpolicy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ppuo.mutation.done = true
	return _node, nil
}
