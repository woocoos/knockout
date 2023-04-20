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
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/predicate"
)

// OrgPolicyUpdate is the builder for updating OrgPolicy entities.
type OrgPolicyUpdate struct {
	config
	hooks    []Hook
	mutation *OrgPolicyMutation
}

// Where appends a list predicates to the OrgPolicyUpdate builder.
func (opu *OrgPolicyUpdate) Where(ps ...predicate.OrgPolicy) *OrgPolicyUpdate {
	opu.mutation.Where(ps...)
	return opu
}

// SetUpdatedBy sets the "updated_by" field.
func (opu *OrgPolicyUpdate) SetUpdatedBy(i int) *OrgPolicyUpdate {
	opu.mutation.ResetUpdatedBy()
	opu.mutation.SetUpdatedBy(i)
	return opu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (opu *OrgPolicyUpdate) SetNillableUpdatedBy(i *int) *OrgPolicyUpdate {
	if i != nil {
		opu.SetUpdatedBy(*i)
	}
	return opu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (opu *OrgPolicyUpdate) AddUpdatedBy(i int) *OrgPolicyUpdate {
	opu.mutation.AddUpdatedBy(i)
	return opu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (opu *OrgPolicyUpdate) ClearUpdatedBy() *OrgPolicyUpdate {
	opu.mutation.ClearUpdatedBy()
	return opu
}

// SetUpdatedAt sets the "updated_at" field.
func (opu *OrgPolicyUpdate) SetUpdatedAt(t time.Time) *OrgPolicyUpdate {
	opu.mutation.SetUpdatedAt(t)
	return opu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (opu *OrgPolicyUpdate) SetNillableUpdatedAt(t *time.Time) *OrgPolicyUpdate {
	if t != nil {
		opu.SetUpdatedAt(*t)
	}
	return opu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (opu *OrgPolicyUpdate) ClearUpdatedAt() *OrgPolicyUpdate {
	opu.mutation.ClearUpdatedAt()
	return opu
}

// SetAppID sets the "app_id" field.
func (opu *OrgPolicyUpdate) SetAppID(i int) *OrgPolicyUpdate {
	opu.mutation.ResetAppID()
	opu.mutation.SetAppID(i)
	return opu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (opu *OrgPolicyUpdate) SetNillableAppID(i *int) *OrgPolicyUpdate {
	if i != nil {
		opu.SetAppID(*i)
	}
	return opu
}

// AddAppID adds i to the "app_id" field.
func (opu *OrgPolicyUpdate) AddAppID(i int) *OrgPolicyUpdate {
	opu.mutation.AddAppID(i)
	return opu
}

// ClearAppID clears the value of the "app_id" field.
func (opu *OrgPolicyUpdate) ClearAppID() *OrgPolicyUpdate {
	opu.mutation.ClearAppID()
	return opu
}

// SetAppPolicyID sets the "app_policy_id" field.
func (opu *OrgPolicyUpdate) SetAppPolicyID(i int) *OrgPolicyUpdate {
	opu.mutation.ResetAppPolicyID()
	opu.mutation.SetAppPolicyID(i)
	return opu
}

// SetNillableAppPolicyID sets the "app_policy_id" field if the given value is not nil.
func (opu *OrgPolicyUpdate) SetNillableAppPolicyID(i *int) *OrgPolicyUpdate {
	if i != nil {
		opu.SetAppPolicyID(*i)
	}
	return opu
}

// AddAppPolicyID adds i to the "app_policy_id" field.
func (opu *OrgPolicyUpdate) AddAppPolicyID(i int) *OrgPolicyUpdate {
	opu.mutation.AddAppPolicyID(i)
	return opu
}

// ClearAppPolicyID clears the value of the "app_policy_id" field.
func (opu *OrgPolicyUpdate) ClearAppPolicyID() *OrgPolicyUpdate {
	opu.mutation.ClearAppPolicyID()
	return opu
}

// SetName sets the "name" field.
func (opu *OrgPolicyUpdate) SetName(s string) *OrgPolicyUpdate {
	opu.mutation.SetName(s)
	return opu
}

// SetComments sets the "comments" field.
func (opu *OrgPolicyUpdate) SetComments(s string) *OrgPolicyUpdate {
	opu.mutation.SetComments(s)
	return opu
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (opu *OrgPolicyUpdate) SetNillableComments(s *string) *OrgPolicyUpdate {
	if s != nil {
		opu.SetComments(*s)
	}
	return opu
}

// ClearComments clears the value of the "comments" field.
func (opu *OrgPolicyUpdate) ClearComments() *OrgPolicyUpdate {
	opu.mutation.ClearComments()
	return opu
}

// SetRules sets the "rules" field.
func (opu *OrgPolicyUpdate) SetRules(tr []types.PolicyRule) *OrgPolicyUpdate {
	opu.mutation.SetRules(tr)
	return opu
}

// AppendRules appends tr to the "rules" field.
func (opu *OrgPolicyUpdate) AppendRules(tr []types.PolicyRule) *OrgPolicyUpdate {
	opu.mutation.AppendRules(tr)
	return opu
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (opu *OrgPolicyUpdate) AddPermissionIDs(ids ...int) *OrgPolicyUpdate {
	opu.mutation.AddPermissionIDs(ids...)
	return opu
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (opu *OrgPolicyUpdate) AddPermissions(p ...*Permission) *OrgPolicyUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return opu.AddPermissionIDs(ids...)
}

// Mutation returns the OrgPolicyMutation object of the builder.
func (opu *OrgPolicyUpdate) Mutation() *OrgPolicyMutation {
	return opu.mutation
}

// ClearPermissions clears all "permissions" edges to the Permission entity.
func (opu *OrgPolicyUpdate) ClearPermissions() *OrgPolicyUpdate {
	opu.mutation.ClearPermissions()
	return opu
}

// RemovePermissionIDs removes the "permissions" edge to Permission entities by IDs.
func (opu *OrgPolicyUpdate) RemovePermissionIDs(ids ...int) *OrgPolicyUpdate {
	opu.mutation.RemovePermissionIDs(ids...)
	return opu
}

// RemovePermissions removes "permissions" edges to Permission entities.
func (opu *OrgPolicyUpdate) RemovePermissions(p ...*Permission) *OrgPolicyUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return opu.RemovePermissionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (opu *OrgPolicyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OrgPolicyMutation](ctx, opu.sqlSave, opu.mutation, opu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opu *OrgPolicyUpdate) SaveX(ctx context.Context) int {
	affected, err := opu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (opu *OrgPolicyUpdate) Exec(ctx context.Context) error {
	_, err := opu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opu *OrgPolicyUpdate) ExecX(ctx context.Context) {
	if err := opu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (opu *OrgPolicyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(orgpolicy.Table, orgpolicy.Columns, sqlgraph.NewFieldSpec(orgpolicy.FieldID, field.TypeInt))
	if ps := opu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opu.mutation.UpdatedBy(); ok {
		_spec.SetField(orgpolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := opu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(orgpolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if opu.mutation.UpdatedByCleared() {
		_spec.ClearField(orgpolicy.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := opu.mutation.UpdatedAt(); ok {
		_spec.SetField(orgpolicy.FieldUpdatedAt, field.TypeTime, value)
	}
	if opu.mutation.UpdatedAtCleared() {
		_spec.ClearField(orgpolicy.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := opu.mutation.AppID(); ok {
		_spec.SetField(orgpolicy.FieldAppID, field.TypeInt, value)
	}
	if value, ok := opu.mutation.AddedAppID(); ok {
		_spec.AddField(orgpolicy.FieldAppID, field.TypeInt, value)
	}
	if opu.mutation.AppIDCleared() {
		_spec.ClearField(orgpolicy.FieldAppID, field.TypeInt)
	}
	if value, ok := opu.mutation.AppPolicyID(); ok {
		_spec.SetField(orgpolicy.FieldAppPolicyID, field.TypeInt, value)
	}
	if value, ok := opu.mutation.AddedAppPolicyID(); ok {
		_spec.AddField(orgpolicy.FieldAppPolicyID, field.TypeInt, value)
	}
	if opu.mutation.AppPolicyIDCleared() {
		_spec.ClearField(orgpolicy.FieldAppPolicyID, field.TypeInt)
	}
	if value, ok := opu.mutation.Name(); ok {
		_spec.SetField(orgpolicy.FieldName, field.TypeString, value)
	}
	if value, ok := opu.mutation.Comments(); ok {
		_spec.SetField(orgpolicy.FieldComments, field.TypeString, value)
	}
	if opu.mutation.CommentsCleared() {
		_spec.ClearField(orgpolicy.FieldComments, field.TypeString)
	}
	if value, ok := opu.mutation.Rules(); ok {
		_spec.SetField(orgpolicy.FieldRules, field.TypeJSON, value)
	}
	if value, ok := opu.mutation.AppendedRules(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, orgpolicy.FieldRules, value)
		})
	}
	if opu.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orgpolicy.PermissionsTable,
			Columns: []string{orgpolicy.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !opu.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orgpolicy.PermissionsTable,
			Columns: []string{orgpolicy.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orgpolicy.PermissionsTable,
			Columns: []string{orgpolicy.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, opu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgpolicy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	opu.mutation.done = true
	return n, nil
}

// OrgPolicyUpdateOne is the builder for updating a single OrgPolicy entity.
type OrgPolicyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrgPolicyMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (opuo *OrgPolicyUpdateOne) SetUpdatedBy(i int) *OrgPolicyUpdateOne {
	opuo.mutation.ResetUpdatedBy()
	opuo.mutation.SetUpdatedBy(i)
	return opuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (opuo *OrgPolicyUpdateOne) SetNillableUpdatedBy(i *int) *OrgPolicyUpdateOne {
	if i != nil {
		opuo.SetUpdatedBy(*i)
	}
	return opuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (opuo *OrgPolicyUpdateOne) AddUpdatedBy(i int) *OrgPolicyUpdateOne {
	opuo.mutation.AddUpdatedBy(i)
	return opuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (opuo *OrgPolicyUpdateOne) ClearUpdatedBy() *OrgPolicyUpdateOne {
	opuo.mutation.ClearUpdatedBy()
	return opuo
}

// SetUpdatedAt sets the "updated_at" field.
func (opuo *OrgPolicyUpdateOne) SetUpdatedAt(t time.Time) *OrgPolicyUpdateOne {
	opuo.mutation.SetUpdatedAt(t)
	return opuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (opuo *OrgPolicyUpdateOne) SetNillableUpdatedAt(t *time.Time) *OrgPolicyUpdateOne {
	if t != nil {
		opuo.SetUpdatedAt(*t)
	}
	return opuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (opuo *OrgPolicyUpdateOne) ClearUpdatedAt() *OrgPolicyUpdateOne {
	opuo.mutation.ClearUpdatedAt()
	return opuo
}

// SetAppID sets the "app_id" field.
func (opuo *OrgPolicyUpdateOne) SetAppID(i int) *OrgPolicyUpdateOne {
	opuo.mutation.ResetAppID()
	opuo.mutation.SetAppID(i)
	return opuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (opuo *OrgPolicyUpdateOne) SetNillableAppID(i *int) *OrgPolicyUpdateOne {
	if i != nil {
		opuo.SetAppID(*i)
	}
	return opuo
}

// AddAppID adds i to the "app_id" field.
func (opuo *OrgPolicyUpdateOne) AddAppID(i int) *OrgPolicyUpdateOne {
	opuo.mutation.AddAppID(i)
	return opuo
}

// ClearAppID clears the value of the "app_id" field.
func (opuo *OrgPolicyUpdateOne) ClearAppID() *OrgPolicyUpdateOne {
	opuo.mutation.ClearAppID()
	return opuo
}

// SetAppPolicyID sets the "app_policy_id" field.
func (opuo *OrgPolicyUpdateOne) SetAppPolicyID(i int) *OrgPolicyUpdateOne {
	opuo.mutation.ResetAppPolicyID()
	opuo.mutation.SetAppPolicyID(i)
	return opuo
}

// SetNillableAppPolicyID sets the "app_policy_id" field if the given value is not nil.
func (opuo *OrgPolicyUpdateOne) SetNillableAppPolicyID(i *int) *OrgPolicyUpdateOne {
	if i != nil {
		opuo.SetAppPolicyID(*i)
	}
	return opuo
}

// AddAppPolicyID adds i to the "app_policy_id" field.
func (opuo *OrgPolicyUpdateOne) AddAppPolicyID(i int) *OrgPolicyUpdateOne {
	opuo.mutation.AddAppPolicyID(i)
	return opuo
}

// ClearAppPolicyID clears the value of the "app_policy_id" field.
func (opuo *OrgPolicyUpdateOne) ClearAppPolicyID() *OrgPolicyUpdateOne {
	opuo.mutation.ClearAppPolicyID()
	return opuo
}

// SetName sets the "name" field.
func (opuo *OrgPolicyUpdateOne) SetName(s string) *OrgPolicyUpdateOne {
	opuo.mutation.SetName(s)
	return opuo
}

// SetComments sets the "comments" field.
func (opuo *OrgPolicyUpdateOne) SetComments(s string) *OrgPolicyUpdateOne {
	opuo.mutation.SetComments(s)
	return opuo
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (opuo *OrgPolicyUpdateOne) SetNillableComments(s *string) *OrgPolicyUpdateOne {
	if s != nil {
		opuo.SetComments(*s)
	}
	return opuo
}

// ClearComments clears the value of the "comments" field.
func (opuo *OrgPolicyUpdateOne) ClearComments() *OrgPolicyUpdateOne {
	opuo.mutation.ClearComments()
	return opuo
}

// SetRules sets the "rules" field.
func (opuo *OrgPolicyUpdateOne) SetRules(tr []types.PolicyRule) *OrgPolicyUpdateOne {
	opuo.mutation.SetRules(tr)
	return opuo
}

// AppendRules appends tr to the "rules" field.
func (opuo *OrgPolicyUpdateOne) AppendRules(tr []types.PolicyRule) *OrgPolicyUpdateOne {
	opuo.mutation.AppendRules(tr)
	return opuo
}

// AddPermissionIDs adds the "permissions" edge to the Permission entity by IDs.
func (opuo *OrgPolicyUpdateOne) AddPermissionIDs(ids ...int) *OrgPolicyUpdateOne {
	opuo.mutation.AddPermissionIDs(ids...)
	return opuo
}

// AddPermissions adds the "permissions" edges to the Permission entity.
func (opuo *OrgPolicyUpdateOne) AddPermissions(p ...*Permission) *OrgPolicyUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return opuo.AddPermissionIDs(ids...)
}

// Mutation returns the OrgPolicyMutation object of the builder.
func (opuo *OrgPolicyUpdateOne) Mutation() *OrgPolicyMutation {
	return opuo.mutation
}

// ClearPermissions clears all "permissions" edges to the Permission entity.
func (opuo *OrgPolicyUpdateOne) ClearPermissions() *OrgPolicyUpdateOne {
	opuo.mutation.ClearPermissions()
	return opuo
}

// RemovePermissionIDs removes the "permissions" edge to Permission entities by IDs.
func (opuo *OrgPolicyUpdateOne) RemovePermissionIDs(ids ...int) *OrgPolicyUpdateOne {
	opuo.mutation.RemovePermissionIDs(ids...)
	return opuo
}

// RemovePermissions removes "permissions" edges to Permission entities.
func (opuo *OrgPolicyUpdateOne) RemovePermissions(p ...*Permission) *OrgPolicyUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return opuo.RemovePermissionIDs(ids...)
}

// Where appends a list predicates to the OrgPolicyUpdate builder.
func (opuo *OrgPolicyUpdateOne) Where(ps ...predicate.OrgPolicy) *OrgPolicyUpdateOne {
	opuo.mutation.Where(ps...)
	return opuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (opuo *OrgPolicyUpdateOne) Select(field string, fields ...string) *OrgPolicyUpdateOne {
	opuo.fields = append([]string{field}, fields...)
	return opuo
}

// Save executes the query and returns the updated OrgPolicy entity.
func (opuo *OrgPolicyUpdateOne) Save(ctx context.Context) (*OrgPolicy, error) {
	return withHooks[*OrgPolicy, OrgPolicyMutation](ctx, opuo.sqlSave, opuo.mutation, opuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opuo *OrgPolicyUpdateOne) SaveX(ctx context.Context) *OrgPolicy {
	node, err := opuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (opuo *OrgPolicyUpdateOne) Exec(ctx context.Context) error {
	_, err := opuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opuo *OrgPolicyUpdateOne) ExecX(ctx context.Context) {
	if err := opuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (opuo *OrgPolicyUpdateOne) sqlSave(ctx context.Context) (_node *OrgPolicy, err error) {
	_spec := sqlgraph.NewUpdateSpec(orgpolicy.Table, orgpolicy.Columns, sqlgraph.NewFieldSpec(orgpolicy.FieldID, field.TypeInt))
	id, ok := opuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrgPolicy.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := opuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgpolicy.FieldID)
		for _, f := range fields {
			if !orgpolicy.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orgpolicy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := opuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opuo.mutation.UpdatedBy(); ok {
		_spec.SetField(orgpolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := opuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(orgpolicy.FieldUpdatedBy, field.TypeInt, value)
	}
	if opuo.mutation.UpdatedByCleared() {
		_spec.ClearField(orgpolicy.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := opuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orgpolicy.FieldUpdatedAt, field.TypeTime, value)
	}
	if opuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(orgpolicy.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := opuo.mutation.AppID(); ok {
		_spec.SetField(orgpolicy.FieldAppID, field.TypeInt, value)
	}
	if value, ok := opuo.mutation.AddedAppID(); ok {
		_spec.AddField(orgpolicy.FieldAppID, field.TypeInt, value)
	}
	if opuo.mutation.AppIDCleared() {
		_spec.ClearField(orgpolicy.FieldAppID, field.TypeInt)
	}
	if value, ok := opuo.mutation.AppPolicyID(); ok {
		_spec.SetField(orgpolicy.FieldAppPolicyID, field.TypeInt, value)
	}
	if value, ok := opuo.mutation.AddedAppPolicyID(); ok {
		_spec.AddField(orgpolicy.FieldAppPolicyID, field.TypeInt, value)
	}
	if opuo.mutation.AppPolicyIDCleared() {
		_spec.ClearField(orgpolicy.FieldAppPolicyID, field.TypeInt)
	}
	if value, ok := opuo.mutation.Name(); ok {
		_spec.SetField(orgpolicy.FieldName, field.TypeString, value)
	}
	if value, ok := opuo.mutation.Comments(); ok {
		_spec.SetField(orgpolicy.FieldComments, field.TypeString, value)
	}
	if opuo.mutation.CommentsCleared() {
		_spec.ClearField(orgpolicy.FieldComments, field.TypeString)
	}
	if value, ok := opuo.mutation.Rules(); ok {
		_spec.SetField(orgpolicy.FieldRules, field.TypeJSON, value)
	}
	if value, ok := opuo.mutation.AppendedRules(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, orgpolicy.FieldRules, value)
		})
	}
	if opuo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orgpolicy.PermissionsTable,
			Columns: []string{orgpolicy.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.RemovedPermissionsIDs(); len(nodes) > 0 && !opuo.mutation.PermissionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orgpolicy.PermissionsTable,
			Columns: []string{orgpolicy.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.PermissionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orgpolicy.PermissionsTable,
			Columns: []string{orgpolicy.PermissionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrgPolicy{config: opuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, opuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgpolicy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	opuo.mutation.done = true
	return _node, nil
}
