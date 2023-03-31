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

// AppRoleUpdate is the builder for updating AppRole entities.
type AppRoleUpdate struct {
	config
	hooks    []Hook
	mutation *AppRoleMutation
}

// Where appends a list predicates to the AppRoleUpdate builder.
func (aru *AppRoleUpdate) Where(ps ...predicate.AppRole) *AppRoleUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetUpdatedBy sets the "updated_by" field.
func (aru *AppRoleUpdate) SetUpdatedBy(i int) *AppRoleUpdate {
	aru.mutation.ResetUpdatedBy()
	aru.mutation.SetUpdatedBy(i)
	return aru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (aru *AppRoleUpdate) SetNillableUpdatedBy(i *int) *AppRoleUpdate {
	if i != nil {
		aru.SetUpdatedBy(*i)
	}
	return aru
}

// AddUpdatedBy adds i to the "updated_by" field.
func (aru *AppRoleUpdate) AddUpdatedBy(i int) *AppRoleUpdate {
	aru.mutation.AddUpdatedBy(i)
	return aru
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (aru *AppRoleUpdate) ClearUpdatedBy() *AppRoleUpdate {
	aru.mutation.ClearUpdatedBy()
	return aru
}

// SetUpdatedAt sets the "updated_at" field.
func (aru *AppRoleUpdate) SetUpdatedAt(t time.Time) *AppRoleUpdate {
	aru.mutation.SetUpdatedAt(t)
	return aru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aru *AppRoleUpdate) SetNillableUpdatedAt(t *time.Time) *AppRoleUpdate {
	if t != nil {
		aru.SetUpdatedAt(*t)
	}
	return aru
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (aru *AppRoleUpdate) ClearUpdatedAt() *AppRoleUpdate {
	aru.mutation.ClearUpdatedAt()
	return aru
}

// SetName sets the "name" field.
func (aru *AppRoleUpdate) SetName(s string) *AppRoleUpdate {
	aru.mutation.SetName(s)
	return aru
}

// SetComments sets the "comments" field.
func (aru *AppRoleUpdate) SetComments(s string) *AppRoleUpdate {
	aru.mutation.SetComments(s)
	return aru
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (aru *AppRoleUpdate) SetNillableComments(s *string) *AppRoleUpdate {
	if s != nil {
		aru.SetComments(*s)
	}
	return aru
}

// ClearComments clears the value of the "comments" field.
func (aru *AppRoleUpdate) ClearComments() *AppRoleUpdate {
	aru.mutation.ClearComments()
	return aru
}

// SetAutoGrant sets the "auto_grant" field.
func (aru *AppRoleUpdate) SetAutoGrant(b bool) *AppRoleUpdate {
	aru.mutation.SetAutoGrant(b)
	return aru
}

// SetNillableAutoGrant sets the "auto_grant" field if the given value is not nil.
func (aru *AppRoleUpdate) SetNillableAutoGrant(b *bool) *AppRoleUpdate {
	if b != nil {
		aru.SetAutoGrant(*b)
	}
	return aru
}

// SetEditable sets the "editable" field.
func (aru *AppRoleUpdate) SetEditable(b bool) *AppRoleUpdate {
	aru.mutation.SetEditable(b)
	return aru
}

// SetNillableEditable sets the "editable" field if the given value is not nil.
func (aru *AppRoleUpdate) SetNillableEditable(b *bool) *AppRoleUpdate {
	if b != nil {
		aru.SetEditable(*b)
	}
	return aru
}

// AddPolicyIDs adds the "policies" edge to the AppPolicy entity by IDs.
func (aru *AppRoleUpdate) AddPolicyIDs(ids ...int) *AppRoleUpdate {
	aru.mutation.AddPolicyIDs(ids...)
	return aru
}

// AddPolicies adds the "policies" edges to the AppPolicy entity.
func (aru *AppRoleUpdate) AddPolicies(a ...*AppPolicy) *AppRoleUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aru.AddPolicyIDs(ids...)
}

// AddAppRolePolicyIDs adds the "app_role_policy" edge to the AppRolePolicy entity by IDs.
func (aru *AppRoleUpdate) AddAppRolePolicyIDs(ids ...int) *AppRoleUpdate {
	aru.mutation.AddAppRolePolicyIDs(ids...)
	return aru
}

// AddAppRolePolicy adds the "app_role_policy" edges to the AppRolePolicy entity.
func (aru *AppRoleUpdate) AddAppRolePolicy(a ...*AppRolePolicy) *AppRoleUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aru.AddAppRolePolicyIDs(ids...)
}

// Mutation returns the AppRoleMutation object of the builder.
func (aru *AppRoleUpdate) Mutation() *AppRoleMutation {
	return aru.mutation
}

// ClearPolicies clears all "policies" edges to the AppPolicy entity.
func (aru *AppRoleUpdate) ClearPolicies() *AppRoleUpdate {
	aru.mutation.ClearPolicies()
	return aru
}

// RemovePolicyIDs removes the "policies" edge to AppPolicy entities by IDs.
func (aru *AppRoleUpdate) RemovePolicyIDs(ids ...int) *AppRoleUpdate {
	aru.mutation.RemovePolicyIDs(ids...)
	return aru
}

// RemovePolicies removes "policies" edges to AppPolicy entities.
func (aru *AppRoleUpdate) RemovePolicies(a ...*AppPolicy) *AppRoleUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aru.RemovePolicyIDs(ids...)
}

// ClearAppRolePolicy clears all "app_role_policy" edges to the AppRolePolicy entity.
func (aru *AppRoleUpdate) ClearAppRolePolicy() *AppRoleUpdate {
	aru.mutation.ClearAppRolePolicy()
	return aru
}

// RemoveAppRolePolicyIDs removes the "app_role_policy" edge to AppRolePolicy entities by IDs.
func (aru *AppRoleUpdate) RemoveAppRolePolicyIDs(ids ...int) *AppRoleUpdate {
	aru.mutation.RemoveAppRolePolicyIDs(ids...)
	return aru
}

// RemoveAppRolePolicy removes "app_role_policy" edges to AppRolePolicy entities.
func (aru *AppRoleUpdate) RemoveAppRolePolicy(a ...*AppRolePolicy) *AppRoleUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aru.RemoveAppRolePolicyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *AppRoleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AppRoleMutation](ctx, aru.sqlSave, aru.mutation, aru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aru *AppRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *AppRoleUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *AppRoleUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aru *AppRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(approle.Table, approle.Columns, sqlgraph.NewFieldSpec(approle.FieldID, field.TypeInt))
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.UpdatedBy(); ok {
		_spec.SetField(approle.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := aru.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(approle.FieldUpdatedBy, field.TypeInt, value)
	}
	if aru.mutation.UpdatedByCleared() {
		_spec.ClearField(approle.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := aru.mutation.UpdatedAt(); ok {
		_spec.SetField(approle.FieldUpdatedAt, field.TypeTime, value)
	}
	if aru.mutation.UpdatedAtCleared() {
		_spec.ClearField(approle.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := aru.mutation.Name(); ok {
		_spec.SetField(approle.FieldName, field.TypeString, value)
	}
	if value, ok := aru.mutation.Comments(); ok {
		_spec.SetField(approle.FieldComments, field.TypeString, value)
	}
	if aru.mutation.CommentsCleared() {
		_spec.ClearField(approle.FieldComments, field.TypeString)
	}
	if value, ok := aru.mutation.AutoGrant(); ok {
		_spec.SetField(approle.FieldAutoGrant, field.TypeBool, value)
	}
	if value, ok := aru.mutation.Editable(); ok {
		_spec.SetField(approle.FieldEditable, field.TypeBool, value)
	}
	if aru.mutation.PoliciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   approle.PoliciesTable,
			Columns: approle.PoliciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt),
			},
		}
		createE := &AppRolePolicyCreate{config: aru.config, mutation: newAppRolePolicyMutation(aru.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.RemovedPoliciesIDs(); len(nodes) > 0 && !aru.mutation.PoliciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   approle.PoliciesTable,
			Columns: approle.PoliciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppRolePolicyCreate{config: aru.config, mutation: newAppRolePolicyMutation(aru.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.PoliciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   approle.PoliciesTable,
			Columns: approle.PoliciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppRolePolicyCreate{config: aru.config, mutation: newAppRolePolicyMutation(aru.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aru.mutation.AppRolePolicyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   approle.AppRolePolicyTable,
			Columns: []string{approle.AppRolePolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approlepolicy.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.RemovedAppRolePolicyIDs(); len(nodes) > 0 && !aru.mutation.AppRolePolicyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   approle.AppRolePolicyTable,
			Columns: []string{approle.AppRolePolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approlepolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aru.mutation.AppRolePolicyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   approle.AppRolePolicyTable,
			Columns: []string{approle.AppRolePolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approlepolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{approle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aru.mutation.done = true
	return n, nil
}

// AppRoleUpdateOne is the builder for updating a single AppRole entity.
type AppRoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppRoleMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (aruo *AppRoleUpdateOne) SetUpdatedBy(i int) *AppRoleUpdateOne {
	aruo.mutation.ResetUpdatedBy()
	aruo.mutation.SetUpdatedBy(i)
	return aruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (aruo *AppRoleUpdateOne) SetNillableUpdatedBy(i *int) *AppRoleUpdateOne {
	if i != nil {
		aruo.SetUpdatedBy(*i)
	}
	return aruo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (aruo *AppRoleUpdateOne) AddUpdatedBy(i int) *AppRoleUpdateOne {
	aruo.mutation.AddUpdatedBy(i)
	return aruo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (aruo *AppRoleUpdateOne) ClearUpdatedBy() *AppRoleUpdateOne {
	aruo.mutation.ClearUpdatedBy()
	return aruo
}

// SetUpdatedAt sets the "updated_at" field.
func (aruo *AppRoleUpdateOne) SetUpdatedAt(t time.Time) *AppRoleUpdateOne {
	aruo.mutation.SetUpdatedAt(t)
	return aruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aruo *AppRoleUpdateOne) SetNillableUpdatedAt(t *time.Time) *AppRoleUpdateOne {
	if t != nil {
		aruo.SetUpdatedAt(*t)
	}
	return aruo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (aruo *AppRoleUpdateOne) ClearUpdatedAt() *AppRoleUpdateOne {
	aruo.mutation.ClearUpdatedAt()
	return aruo
}

// SetName sets the "name" field.
func (aruo *AppRoleUpdateOne) SetName(s string) *AppRoleUpdateOne {
	aruo.mutation.SetName(s)
	return aruo
}

// SetComments sets the "comments" field.
func (aruo *AppRoleUpdateOne) SetComments(s string) *AppRoleUpdateOne {
	aruo.mutation.SetComments(s)
	return aruo
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (aruo *AppRoleUpdateOne) SetNillableComments(s *string) *AppRoleUpdateOne {
	if s != nil {
		aruo.SetComments(*s)
	}
	return aruo
}

// ClearComments clears the value of the "comments" field.
func (aruo *AppRoleUpdateOne) ClearComments() *AppRoleUpdateOne {
	aruo.mutation.ClearComments()
	return aruo
}

// SetAutoGrant sets the "auto_grant" field.
func (aruo *AppRoleUpdateOne) SetAutoGrant(b bool) *AppRoleUpdateOne {
	aruo.mutation.SetAutoGrant(b)
	return aruo
}

// SetNillableAutoGrant sets the "auto_grant" field if the given value is not nil.
func (aruo *AppRoleUpdateOne) SetNillableAutoGrant(b *bool) *AppRoleUpdateOne {
	if b != nil {
		aruo.SetAutoGrant(*b)
	}
	return aruo
}

// SetEditable sets the "editable" field.
func (aruo *AppRoleUpdateOne) SetEditable(b bool) *AppRoleUpdateOne {
	aruo.mutation.SetEditable(b)
	return aruo
}

// SetNillableEditable sets the "editable" field if the given value is not nil.
func (aruo *AppRoleUpdateOne) SetNillableEditable(b *bool) *AppRoleUpdateOne {
	if b != nil {
		aruo.SetEditable(*b)
	}
	return aruo
}

// AddPolicyIDs adds the "policies" edge to the AppPolicy entity by IDs.
func (aruo *AppRoleUpdateOne) AddPolicyIDs(ids ...int) *AppRoleUpdateOne {
	aruo.mutation.AddPolicyIDs(ids...)
	return aruo
}

// AddPolicies adds the "policies" edges to the AppPolicy entity.
func (aruo *AppRoleUpdateOne) AddPolicies(a ...*AppPolicy) *AppRoleUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aruo.AddPolicyIDs(ids...)
}

// AddAppRolePolicyIDs adds the "app_role_policy" edge to the AppRolePolicy entity by IDs.
func (aruo *AppRoleUpdateOne) AddAppRolePolicyIDs(ids ...int) *AppRoleUpdateOne {
	aruo.mutation.AddAppRolePolicyIDs(ids...)
	return aruo
}

// AddAppRolePolicy adds the "app_role_policy" edges to the AppRolePolicy entity.
func (aruo *AppRoleUpdateOne) AddAppRolePolicy(a ...*AppRolePolicy) *AppRoleUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aruo.AddAppRolePolicyIDs(ids...)
}

// Mutation returns the AppRoleMutation object of the builder.
func (aruo *AppRoleUpdateOne) Mutation() *AppRoleMutation {
	return aruo.mutation
}

// ClearPolicies clears all "policies" edges to the AppPolicy entity.
func (aruo *AppRoleUpdateOne) ClearPolicies() *AppRoleUpdateOne {
	aruo.mutation.ClearPolicies()
	return aruo
}

// RemovePolicyIDs removes the "policies" edge to AppPolicy entities by IDs.
func (aruo *AppRoleUpdateOne) RemovePolicyIDs(ids ...int) *AppRoleUpdateOne {
	aruo.mutation.RemovePolicyIDs(ids...)
	return aruo
}

// RemovePolicies removes "policies" edges to AppPolicy entities.
func (aruo *AppRoleUpdateOne) RemovePolicies(a ...*AppPolicy) *AppRoleUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aruo.RemovePolicyIDs(ids...)
}

// ClearAppRolePolicy clears all "app_role_policy" edges to the AppRolePolicy entity.
func (aruo *AppRoleUpdateOne) ClearAppRolePolicy() *AppRoleUpdateOne {
	aruo.mutation.ClearAppRolePolicy()
	return aruo
}

// RemoveAppRolePolicyIDs removes the "app_role_policy" edge to AppRolePolicy entities by IDs.
func (aruo *AppRoleUpdateOne) RemoveAppRolePolicyIDs(ids ...int) *AppRoleUpdateOne {
	aruo.mutation.RemoveAppRolePolicyIDs(ids...)
	return aruo
}

// RemoveAppRolePolicy removes "app_role_policy" edges to AppRolePolicy entities.
func (aruo *AppRoleUpdateOne) RemoveAppRolePolicy(a ...*AppRolePolicy) *AppRoleUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return aruo.RemoveAppRolePolicyIDs(ids...)
}

// Where appends a list predicates to the AppRoleUpdate builder.
func (aruo *AppRoleUpdateOne) Where(ps ...predicate.AppRole) *AppRoleUpdateOne {
	aruo.mutation.Where(ps...)
	return aruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *AppRoleUpdateOne) Select(field string, fields ...string) *AppRoleUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated AppRole entity.
func (aruo *AppRoleUpdateOne) Save(ctx context.Context) (*AppRole, error) {
	return withHooks[*AppRole, AppRoleMutation](ctx, aruo.sqlSave, aruo.mutation, aruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *AppRoleUpdateOne) SaveX(ctx context.Context) *AppRole {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *AppRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *AppRoleUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aruo *AppRoleUpdateOne) sqlSave(ctx context.Context) (_node *AppRole, err error) {
	_spec := sqlgraph.NewUpdateSpec(approle.Table, approle.Columns, sqlgraph.NewFieldSpec(approle.FieldID, field.TypeInt))
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppRole.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, approle.FieldID)
		for _, f := range fields {
			if !approle.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != approle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruo.mutation.UpdatedBy(); ok {
		_spec.SetField(approle.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := aruo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(approle.FieldUpdatedBy, field.TypeInt, value)
	}
	if aruo.mutation.UpdatedByCleared() {
		_spec.ClearField(approle.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := aruo.mutation.UpdatedAt(); ok {
		_spec.SetField(approle.FieldUpdatedAt, field.TypeTime, value)
	}
	if aruo.mutation.UpdatedAtCleared() {
		_spec.ClearField(approle.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := aruo.mutation.Name(); ok {
		_spec.SetField(approle.FieldName, field.TypeString, value)
	}
	if value, ok := aruo.mutation.Comments(); ok {
		_spec.SetField(approle.FieldComments, field.TypeString, value)
	}
	if aruo.mutation.CommentsCleared() {
		_spec.ClearField(approle.FieldComments, field.TypeString)
	}
	if value, ok := aruo.mutation.AutoGrant(); ok {
		_spec.SetField(approle.FieldAutoGrant, field.TypeBool, value)
	}
	if value, ok := aruo.mutation.Editable(); ok {
		_spec.SetField(approle.FieldEditable, field.TypeBool, value)
	}
	if aruo.mutation.PoliciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   approle.PoliciesTable,
			Columns: approle.PoliciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt),
			},
		}
		createE := &AppRolePolicyCreate{config: aruo.config, mutation: newAppRolePolicyMutation(aruo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.RemovedPoliciesIDs(); len(nodes) > 0 && !aruo.mutation.PoliciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   approle.PoliciesTable,
			Columns: approle.PoliciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppRolePolicyCreate{config: aruo.config, mutation: newAppRolePolicyMutation(aruo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.PoliciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   approle.PoliciesTable,
			Columns: approle.PoliciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppRolePolicyCreate{config: aruo.config, mutation: newAppRolePolicyMutation(aruo.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aruo.mutation.AppRolePolicyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   approle.AppRolePolicyTable,
			Columns: []string{approle.AppRolePolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approlepolicy.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.RemovedAppRolePolicyIDs(); len(nodes) > 0 && !aruo.mutation.AppRolePolicyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   approle.AppRolePolicyTable,
			Columns: []string{approle.AppRolePolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approlepolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aruo.mutation.AppRolePolicyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   approle.AppRolePolicyTable,
			Columns: []string{approle.AppRolePolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approlepolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AppRole{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{approle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aruo.mutation.done = true
	return _node, nil
}
