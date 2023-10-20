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
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/approle"
	"github.com/woocoos/knockout/ent/approlepolicy"
)

// AppPolicyCreate is the builder for creating a AppPolicy entity.
type AppPolicyCreate struct {
	config
	mutation *AppPolicyMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (apc *AppPolicyCreate) SetCreatedBy(i int) *AppPolicyCreate {
	apc.mutation.SetCreatedBy(i)
	return apc
}

// SetCreatedAt sets the "created_at" field.
func (apc *AppPolicyCreate) SetCreatedAt(t time.Time) *AppPolicyCreate {
	apc.mutation.SetCreatedAt(t)
	return apc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableCreatedAt(t *time.Time) *AppPolicyCreate {
	if t != nil {
		apc.SetCreatedAt(*t)
	}
	return apc
}

// SetUpdatedBy sets the "updated_by" field.
func (apc *AppPolicyCreate) SetUpdatedBy(i int) *AppPolicyCreate {
	apc.mutation.SetUpdatedBy(i)
	return apc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableUpdatedBy(i *int) *AppPolicyCreate {
	if i != nil {
		apc.SetUpdatedBy(*i)
	}
	return apc
}

// SetUpdatedAt sets the "updated_at" field.
func (apc *AppPolicyCreate) SetUpdatedAt(t time.Time) *AppPolicyCreate {
	apc.mutation.SetUpdatedAt(t)
	return apc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableUpdatedAt(t *time.Time) *AppPolicyCreate {
	if t != nil {
		apc.SetUpdatedAt(*t)
	}
	return apc
}

// SetAppID sets the "app_id" field.
func (apc *AppPolicyCreate) SetAppID(i int) *AppPolicyCreate {
	apc.mutation.SetAppID(i)
	return apc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableAppID(i *int) *AppPolicyCreate {
	if i != nil {
		apc.SetAppID(*i)
	}
	return apc
}

// SetName sets the "name" field.
func (apc *AppPolicyCreate) SetName(s string) *AppPolicyCreate {
	apc.mutation.SetName(s)
	return apc
}

// SetComments sets the "comments" field.
func (apc *AppPolicyCreate) SetComments(s string) *AppPolicyCreate {
	apc.mutation.SetComments(s)
	return apc
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableComments(s *string) *AppPolicyCreate {
	if s != nil {
		apc.SetComments(*s)
	}
	return apc
}

// SetRules sets the "rules" field.
func (apc *AppPolicyCreate) SetRules(tr []*types.PolicyRule) *AppPolicyCreate {
	apc.mutation.SetRules(tr)
	return apc
}

// SetVersion sets the "version" field.
func (apc *AppPolicyCreate) SetVersion(s string) *AppPolicyCreate {
	apc.mutation.SetVersion(s)
	return apc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableVersion(s *string) *AppPolicyCreate {
	if s != nil {
		apc.SetVersion(*s)
	}
	return apc
}

// SetAutoGrant sets the "auto_grant" field.
func (apc *AppPolicyCreate) SetAutoGrant(b bool) *AppPolicyCreate {
	apc.mutation.SetAutoGrant(b)
	return apc
}

// SetNillableAutoGrant sets the "auto_grant" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableAutoGrant(b *bool) *AppPolicyCreate {
	if b != nil {
		apc.SetAutoGrant(*b)
	}
	return apc
}

// SetStatus sets the "status" field.
func (apc *AppPolicyCreate) SetStatus(ts typex.SimpleStatus) *AppPolicyCreate {
	apc.mutation.SetStatus(ts)
	return apc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableStatus(ts *typex.SimpleStatus) *AppPolicyCreate {
	if ts != nil {
		apc.SetStatus(*ts)
	}
	return apc
}

// SetID sets the "id" field.
func (apc *AppPolicyCreate) SetID(i int) *AppPolicyCreate {
	apc.mutation.SetID(i)
	return apc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (apc *AppPolicyCreate) SetNillableID(i *int) *AppPolicyCreate {
	if i != nil {
		apc.SetID(*i)
	}
	return apc
}

// SetApp sets the "app" edge to the App entity.
func (apc *AppPolicyCreate) SetApp(a *App) *AppPolicyCreate {
	return apc.SetAppID(a.ID)
}

// AddRoleIDs adds the "roles" edge to the AppRole entity by IDs.
func (apc *AppPolicyCreate) AddRoleIDs(ids ...int) *AppPolicyCreate {
	apc.mutation.AddRoleIDs(ids...)
	return apc
}

// AddRoles adds the "roles" edges to the AppRole entity.
func (apc *AppPolicyCreate) AddRoles(a ...*AppRole) *AppPolicyCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apc.AddRoleIDs(ids...)
}

// AddAppRolePolicyIDs adds the "app_role_policy" edge to the AppRolePolicy entity by IDs.
func (apc *AppPolicyCreate) AddAppRolePolicyIDs(ids ...int) *AppPolicyCreate {
	apc.mutation.AddAppRolePolicyIDs(ids...)
	return apc
}

// AddAppRolePolicy adds the "app_role_policy" edges to the AppRolePolicy entity.
func (apc *AppPolicyCreate) AddAppRolePolicy(a ...*AppRolePolicy) *AppPolicyCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apc.AddAppRolePolicyIDs(ids...)
}

// Mutation returns the AppPolicyMutation object of the builder.
func (apc *AppPolicyCreate) Mutation() *AppPolicyMutation {
	return apc.mutation
}

// Save creates the AppPolicy in the database.
func (apc *AppPolicyCreate) Save(ctx context.Context) (*AppPolicy, error) {
	if err := apc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, apc.sqlSave, apc.mutation, apc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (apc *AppPolicyCreate) SaveX(ctx context.Context) *AppPolicy {
	v, err := apc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apc *AppPolicyCreate) Exec(ctx context.Context) error {
	_, err := apc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apc *AppPolicyCreate) ExecX(ctx context.Context) {
	if err := apc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apc *AppPolicyCreate) defaults() error {
	if _, ok := apc.mutation.CreatedAt(); !ok {
		if apppolicy.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized apppolicy.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := apppolicy.DefaultCreatedAt()
		apc.mutation.SetCreatedAt(v)
	}
	if _, ok := apc.mutation.Version(); !ok {
		v := apppolicy.DefaultVersion
		apc.mutation.SetVersion(v)
	}
	if _, ok := apc.mutation.AutoGrant(); !ok {
		v := apppolicy.DefaultAutoGrant
		apc.mutation.SetAutoGrant(v)
	}
	if _, ok := apc.mutation.Status(); !ok {
		v := apppolicy.DefaultStatus
		apc.mutation.SetStatus(v)
	}
	if _, ok := apc.mutation.ID(); !ok {
		if apppolicy.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized apppolicy.DefaultID (forgotten import ent/runtime?)")
		}
		v := apppolicy.DefaultID()
		apc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (apc *AppPolicyCreate) check() error {
	if _, ok := apc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "AppPolicy.created_by"`)}
	}
	if _, ok := apc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppPolicy.created_at"`)}
	}
	if _, ok := apc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AppPolicy.name"`)}
	}
	if _, ok := apc.mutation.Rules(); !ok {
		return &ValidationError{Name: "rules", err: errors.New(`ent: missing required field "AppPolicy.rules"`)}
	}
	if _, ok := apc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "AppPolicy.version"`)}
	}
	if _, ok := apc.mutation.AutoGrant(); !ok {
		return &ValidationError{Name: "auto_grant", err: errors.New(`ent: missing required field "AppPolicy.auto_grant"`)}
	}
	if v, ok := apc.mutation.Status(); ok {
		if err := apppolicy.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "AppPolicy.status": %w`, err)}
		}
	}
	return nil
}

func (apc *AppPolicyCreate) sqlSave(ctx context.Context) (*AppPolicy, error) {
	if err := apc.check(); err != nil {
		return nil, err
	}
	_node, _spec := apc.createSpec()
	if err := sqlgraph.CreateNode(ctx, apc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	apc.mutation.id = &_node.ID
	apc.mutation.done = true
	return _node, nil
}

func (apc *AppPolicyCreate) createSpec() (*AppPolicy, *sqlgraph.CreateSpec) {
	var (
		_node = &AppPolicy{config: apc.config}
		_spec = sqlgraph.NewCreateSpec(apppolicy.Table, sqlgraph.NewFieldSpec(apppolicy.FieldID, field.TypeInt))
	)
	_spec.OnConflict = apc.conflict
	if id, ok := apc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := apc.mutation.CreatedBy(); ok {
		_spec.SetField(apppolicy.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := apc.mutation.CreatedAt(); ok {
		_spec.SetField(apppolicy.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := apc.mutation.UpdatedBy(); ok {
		_spec.SetField(apppolicy.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := apc.mutation.UpdatedAt(); ok {
		_spec.SetField(apppolicy.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := apc.mutation.Name(); ok {
		_spec.SetField(apppolicy.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := apc.mutation.Comments(); ok {
		_spec.SetField(apppolicy.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if value, ok := apc.mutation.Rules(); ok {
		_spec.SetField(apppolicy.FieldRules, field.TypeJSON, value)
		_node.Rules = value
	}
	if value, ok := apc.mutation.Version(); ok {
		_spec.SetField(apppolicy.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := apc.mutation.AutoGrant(); ok {
		_spec.SetField(apppolicy.FieldAutoGrant, field.TypeBool, value)
		_node.AutoGrant = value
	}
	if value, ok := apc.mutation.Status(); ok {
		_spec.SetField(apppolicy.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := apc.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apppolicy.AppTable,
			Columns: []string{apppolicy.AppColumn},
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
	if nodes := apc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   apppolicy.RolesTable,
			Columns: apppolicy.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppRolePolicyCreate{config: apc.config, mutation: newAppRolePolicyMutation(apc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := apc.mutation.AppRolePolicyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   apppolicy.AppRolePolicyTable,
			Columns: []string{apppolicy.AppRolePolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(approlepolicy.FieldID, field.TypeInt),
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
//	client.AppPolicy.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppPolicyUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (apc *AppPolicyCreate) OnConflict(opts ...sql.ConflictOption) *AppPolicyUpsertOne {
	apc.conflict = opts
	return &AppPolicyUpsertOne{
		create: apc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppPolicy.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (apc *AppPolicyCreate) OnConflictColumns(columns ...string) *AppPolicyUpsertOne {
	apc.conflict = append(apc.conflict, sql.ConflictColumns(columns...))
	return &AppPolicyUpsertOne{
		create: apc,
	}
}

type (
	// AppPolicyUpsertOne is the builder for "upsert"-ing
	//  one AppPolicy node.
	AppPolicyUpsertOne struct {
		create *AppPolicyCreate
	}

	// AppPolicyUpsert is the "OnConflict" setter.
	AppPolicyUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *AppPolicyUpsert) SetUpdatedBy(v int) *AppPolicyUpsert {
	u.Set(apppolicy.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppPolicyUpsert) UpdateUpdatedBy() *AppPolicyUpsert {
	u.SetExcluded(apppolicy.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppPolicyUpsert) AddUpdatedBy(v int) *AppPolicyUpsert {
	u.Add(apppolicy.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppPolicyUpsert) ClearUpdatedBy() *AppPolicyUpsert {
	u.SetNull(apppolicy.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppPolicyUpsert) SetUpdatedAt(v time.Time) *AppPolicyUpsert {
	u.Set(apppolicy.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppPolicyUpsert) UpdateUpdatedAt() *AppPolicyUpsert {
	u.SetExcluded(apppolicy.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppPolicyUpsert) ClearUpdatedAt() *AppPolicyUpsert {
	u.SetNull(apppolicy.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *AppPolicyUpsert) SetName(v string) *AppPolicyUpsert {
	u.Set(apppolicy.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppPolicyUpsert) UpdateName() *AppPolicyUpsert {
	u.SetExcluded(apppolicy.FieldName)
	return u
}

// SetComments sets the "comments" field.
func (u *AppPolicyUpsert) SetComments(v string) *AppPolicyUpsert {
	u.Set(apppolicy.FieldComments, v)
	return u
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppPolicyUpsert) UpdateComments() *AppPolicyUpsert {
	u.SetExcluded(apppolicy.FieldComments)
	return u
}

// ClearComments clears the value of the "comments" field.
func (u *AppPolicyUpsert) ClearComments() *AppPolicyUpsert {
	u.SetNull(apppolicy.FieldComments)
	return u
}

// SetRules sets the "rules" field.
func (u *AppPolicyUpsert) SetRules(v []*types.PolicyRule) *AppPolicyUpsert {
	u.Set(apppolicy.FieldRules, v)
	return u
}

// UpdateRules sets the "rules" field to the value that was provided on create.
func (u *AppPolicyUpsert) UpdateRules() *AppPolicyUpsert {
	u.SetExcluded(apppolicy.FieldRules)
	return u
}

// SetVersion sets the "version" field.
func (u *AppPolicyUpsert) SetVersion(v string) *AppPolicyUpsert {
	u.Set(apppolicy.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *AppPolicyUpsert) UpdateVersion() *AppPolicyUpsert {
	u.SetExcluded(apppolicy.FieldVersion)
	return u
}

// SetAutoGrant sets the "auto_grant" field.
func (u *AppPolicyUpsert) SetAutoGrant(v bool) *AppPolicyUpsert {
	u.Set(apppolicy.FieldAutoGrant, v)
	return u
}

// UpdateAutoGrant sets the "auto_grant" field to the value that was provided on create.
func (u *AppPolicyUpsert) UpdateAutoGrant() *AppPolicyUpsert {
	u.SetExcluded(apppolicy.FieldAutoGrant)
	return u
}

// SetStatus sets the "status" field.
func (u *AppPolicyUpsert) SetStatus(v typex.SimpleStatus) *AppPolicyUpsert {
	u.Set(apppolicy.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AppPolicyUpsert) UpdateStatus() *AppPolicyUpsert {
	u.SetExcluded(apppolicy.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *AppPolicyUpsert) ClearStatus() *AppPolicyUpsert {
	u.SetNull(apppolicy.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppPolicy.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apppolicy.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppPolicyUpsertOne) UpdateNewValues() *AppPolicyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(apppolicy.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(apppolicy.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(apppolicy.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.AppID(); exists {
			s.SetIgnore(apppolicy.FieldAppID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppPolicy.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppPolicyUpsertOne) Ignore() *AppPolicyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppPolicyUpsertOne) DoNothing() *AppPolicyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppPolicyCreate.OnConflict
// documentation for more info.
func (u *AppPolicyUpsertOne) Update(set func(*AppPolicyUpsert)) *AppPolicyUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppPolicyUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AppPolicyUpsertOne) SetUpdatedBy(v int) *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppPolicyUpsertOne) AddUpdatedBy(v int) *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppPolicyUpsertOne) UpdateUpdatedBy() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppPolicyUpsertOne) ClearUpdatedBy() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppPolicyUpsertOne) SetUpdatedAt(v time.Time) *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppPolicyUpsertOne) UpdateUpdatedAt() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppPolicyUpsertOne) ClearUpdatedAt() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *AppPolicyUpsertOne) SetName(v string) *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppPolicyUpsertOne) UpdateName() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateName()
	})
}

// SetComments sets the "comments" field.
func (u *AppPolicyUpsertOne) SetComments(v string) *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppPolicyUpsertOne) UpdateComments() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *AppPolicyUpsertOne) ClearComments() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.ClearComments()
	})
}

// SetRules sets the "rules" field.
func (u *AppPolicyUpsertOne) SetRules(v []*types.PolicyRule) *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetRules(v)
	})
}

// UpdateRules sets the "rules" field to the value that was provided on create.
func (u *AppPolicyUpsertOne) UpdateRules() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateRules()
	})
}

// SetVersion sets the "version" field.
func (u *AppPolicyUpsertOne) SetVersion(v string) *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *AppPolicyUpsertOne) UpdateVersion() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateVersion()
	})
}

// SetAutoGrant sets the "auto_grant" field.
func (u *AppPolicyUpsertOne) SetAutoGrant(v bool) *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetAutoGrant(v)
	})
}

// UpdateAutoGrant sets the "auto_grant" field to the value that was provided on create.
func (u *AppPolicyUpsertOne) UpdateAutoGrant() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateAutoGrant()
	})
}

// SetStatus sets the "status" field.
func (u *AppPolicyUpsertOne) SetStatus(v typex.SimpleStatus) *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AppPolicyUpsertOne) UpdateStatus() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *AppPolicyUpsertOne) ClearStatus() *AppPolicyUpsertOne {
	return u.Update(func(s *AppPolicyUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *AppPolicyUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppPolicyCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppPolicyUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppPolicyUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppPolicyUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppPolicyCreateBulk is the builder for creating many AppPolicy entities in bulk.
type AppPolicyCreateBulk struct {
	config
	err      error
	builders []*AppPolicyCreate
	conflict []sql.ConflictOption
}

// Save creates the AppPolicy entities in the database.
func (apcb *AppPolicyCreateBulk) Save(ctx context.Context) ([]*AppPolicy, error) {
	if apcb.err != nil {
		return nil, apcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(apcb.builders))
	nodes := make([]*AppPolicy, len(apcb.builders))
	mutators := make([]Mutator, len(apcb.builders))
	for i := range apcb.builders {
		func(i int, root context.Context) {
			builder := apcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppPolicyMutation)
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
					_, err = mutators[i+1].Mutate(root, apcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = apcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, apcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, apcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (apcb *AppPolicyCreateBulk) SaveX(ctx context.Context) []*AppPolicy {
	v, err := apcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (apcb *AppPolicyCreateBulk) Exec(ctx context.Context) error {
	_, err := apcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apcb *AppPolicyCreateBulk) ExecX(ctx context.Context) {
	if err := apcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppPolicy.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppPolicyUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (apcb *AppPolicyCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppPolicyUpsertBulk {
	apcb.conflict = opts
	return &AppPolicyUpsertBulk{
		create: apcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppPolicy.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (apcb *AppPolicyCreateBulk) OnConflictColumns(columns ...string) *AppPolicyUpsertBulk {
	apcb.conflict = append(apcb.conflict, sql.ConflictColumns(columns...))
	return &AppPolicyUpsertBulk{
		create: apcb,
	}
}

// AppPolicyUpsertBulk is the builder for "upsert"-ing
// a bulk of AppPolicy nodes.
type AppPolicyUpsertBulk struct {
	create *AppPolicyCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppPolicy.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(apppolicy.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AppPolicyUpsertBulk) UpdateNewValues() *AppPolicyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(apppolicy.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(apppolicy.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(apppolicy.FieldCreatedAt)
			}
			if _, exists := b.mutation.AppID(); exists {
				s.SetIgnore(apppolicy.FieldAppID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppPolicy.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppPolicyUpsertBulk) Ignore() *AppPolicyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppPolicyUpsertBulk) DoNothing() *AppPolicyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppPolicyCreateBulk.OnConflict
// documentation for more info.
func (u *AppPolicyUpsertBulk) Update(set func(*AppPolicyUpsert)) *AppPolicyUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppPolicyUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AppPolicyUpsertBulk) SetUpdatedBy(v int) *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AppPolicyUpsertBulk) AddUpdatedBy(v int) *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AppPolicyUpsertBulk) UpdateUpdatedBy() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *AppPolicyUpsertBulk) ClearUpdatedBy() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppPolicyUpsertBulk) SetUpdatedAt(v time.Time) *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppPolicyUpsertBulk) UpdateUpdatedAt() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *AppPolicyUpsertBulk) ClearUpdatedAt() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *AppPolicyUpsertBulk) SetName(v string) *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppPolicyUpsertBulk) UpdateName() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateName()
	})
}

// SetComments sets the "comments" field.
func (u *AppPolicyUpsertBulk) SetComments(v string) *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetComments(v)
	})
}

// UpdateComments sets the "comments" field to the value that was provided on create.
func (u *AppPolicyUpsertBulk) UpdateComments() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateComments()
	})
}

// ClearComments clears the value of the "comments" field.
func (u *AppPolicyUpsertBulk) ClearComments() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.ClearComments()
	})
}

// SetRules sets the "rules" field.
func (u *AppPolicyUpsertBulk) SetRules(v []*types.PolicyRule) *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetRules(v)
	})
}

// UpdateRules sets the "rules" field to the value that was provided on create.
func (u *AppPolicyUpsertBulk) UpdateRules() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateRules()
	})
}

// SetVersion sets the "version" field.
func (u *AppPolicyUpsertBulk) SetVersion(v string) *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *AppPolicyUpsertBulk) UpdateVersion() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateVersion()
	})
}

// SetAutoGrant sets the "auto_grant" field.
func (u *AppPolicyUpsertBulk) SetAutoGrant(v bool) *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetAutoGrant(v)
	})
}

// UpdateAutoGrant sets the "auto_grant" field to the value that was provided on create.
func (u *AppPolicyUpsertBulk) UpdateAutoGrant() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateAutoGrant()
	})
}

// SetStatus sets the "status" field.
func (u *AppPolicyUpsertBulk) SetStatus(v typex.SimpleStatus) *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AppPolicyUpsertBulk) UpdateStatus() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *AppPolicyUpsertBulk) ClearStatus() *AppPolicyUpsertBulk {
	return u.Update(func(s *AppPolicyUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *AppPolicyUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppPolicyCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppPolicyCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppPolicyUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
