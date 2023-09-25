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
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/user"
)

// PermissionCreate is the builder for creating a Permission entity.
type PermissionCreate struct {
	config
	mutation *PermissionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (pc *PermissionCreate) SetCreatedBy(i int) *PermissionCreate {
	pc.mutation.SetCreatedBy(i)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PermissionCreate) SetCreatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableCreatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedBy sets the "updated_by" field.
func (pc *PermissionCreate) SetUpdatedBy(i int) *PermissionCreate {
	pc.mutation.SetUpdatedBy(i)
	return pc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableUpdatedBy(i *int) *PermissionCreate {
	if i != nil {
		pc.SetUpdatedBy(*i)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PermissionCreate) SetUpdatedAt(t time.Time) *PermissionCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableUpdatedAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetOrgID sets the "org_id" field.
func (pc *PermissionCreate) SetOrgID(i int) *PermissionCreate {
	pc.mutation.SetOrgID(i)
	return pc
}

// SetPrincipalKind sets the "principal_kind" field.
func (pc *PermissionCreate) SetPrincipalKind(pk permission.PrincipalKind) *PermissionCreate {
	pc.mutation.SetPrincipalKind(pk)
	return pc
}

// SetUserID sets the "user_id" field.
func (pc *PermissionCreate) SetUserID(i int) *PermissionCreate {
	pc.mutation.SetUserID(i)
	return pc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableUserID(i *int) *PermissionCreate {
	if i != nil {
		pc.SetUserID(*i)
	}
	return pc
}

// SetRoleID sets the "role_id" field.
func (pc *PermissionCreate) SetRoleID(i int) *PermissionCreate {
	pc.mutation.SetRoleID(i)
	return pc
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableRoleID(i *int) *PermissionCreate {
	if i != nil {
		pc.SetRoleID(*i)
	}
	return pc
}

// SetOrgPolicyID sets the "org_policy_id" field.
func (pc *PermissionCreate) SetOrgPolicyID(i int) *PermissionCreate {
	pc.mutation.SetOrgPolicyID(i)
	return pc
}

// SetStartAt sets the "start_at" field.
func (pc *PermissionCreate) SetStartAt(t time.Time) *PermissionCreate {
	pc.mutation.SetStartAt(t)
	return pc
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableStartAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetStartAt(*t)
	}
	return pc
}

// SetEndAt sets the "end_at" field.
func (pc *PermissionCreate) SetEndAt(t time.Time) *PermissionCreate {
	pc.mutation.SetEndAt(t)
	return pc
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableEndAt(t *time.Time) *PermissionCreate {
	if t != nil {
		pc.SetEndAt(*t)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *PermissionCreate) SetStatus(ts typex.SimpleStatus) *PermissionCreate {
	pc.mutation.SetStatus(ts)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableStatus(ts *typex.SimpleStatus) *PermissionCreate {
	if ts != nil {
		pc.SetStatus(*ts)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PermissionCreate) SetID(i int) *PermissionCreate {
	pc.mutation.SetID(i)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PermissionCreate) SetNillableID(i *int) *PermissionCreate {
	if i != nil {
		pc.SetID(*i)
	}
	return pc
}

// SetOrg sets the "org" edge to the Org entity.
func (pc *PermissionCreate) SetOrg(o *Org) *PermissionCreate {
	return pc.SetOrgID(o.ID)
}

// SetUser sets the "user" edge to the User entity.
func (pc *PermissionCreate) SetUser(u *User) *PermissionCreate {
	return pc.SetUserID(u.ID)
}

// SetRole sets the "role" edge to the OrgRole entity.
func (pc *PermissionCreate) SetRole(o *OrgRole) *PermissionCreate {
	return pc.SetRoleID(o.ID)
}

// SetOrgPolicy sets the "org_policy" edge to the OrgPolicy entity.
func (pc *PermissionCreate) SetOrgPolicy(o *OrgPolicy) *PermissionCreate {
	return pc.SetOrgPolicyID(o.ID)
}

// Mutation returns the PermissionMutation object of the builder.
func (pc *PermissionCreate) Mutation() *PermissionMutation {
	return pc.mutation
}

// Save creates the Permission in the database.
func (pc *PermissionCreate) Save(ctx context.Context) (*Permission, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PermissionCreate) SaveX(ctx context.Context) *Permission {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PermissionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PermissionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PermissionCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if permission.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized permission.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := permission.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		if permission.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized permission.DefaultID (forgotten import ent/runtime?)")
		}
		v := permission.DefaultID()
		pc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PermissionCreate) check() error {
	if _, ok := pc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Permission.created_by"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Permission.created_at"`)}
	}
	if _, ok := pc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "Permission.org_id"`)}
	}
	if _, ok := pc.mutation.PrincipalKind(); !ok {
		return &ValidationError{Name: "principal_kind", err: errors.New(`ent: missing required field "Permission.principal_kind"`)}
	}
	if v, ok := pc.mutation.PrincipalKind(); ok {
		if err := permission.PrincipalKindValidator(v); err != nil {
			return &ValidationError{Name: "principal_kind", err: fmt.Errorf(`ent: validator failed for field "Permission.principal_kind": %w`, err)}
		}
	}
	if _, ok := pc.mutation.OrgPolicyID(); !ok {
		return &ValidationError{Name: "org_policy_id", err: errors.New(`ent: missing required field "Permission.org_policy_id"`)}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := permission.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Permission.status": %w`, err)}
		}
	}
	if _, ok := pc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org", err: errors.New(`ent: missing required edge "Permission.org"`)}
	}
	if _, ok := pc.mutation.OrgPolicyID(); !ok {
		return &ValidationError{Name: "org_policy", err: errors.New(`ent: missing required edge "Permission.org_policy"`)}
	}
	return nil
}

func (pc *PermissionCreate) sqlSave(ctx context.Context) (*Permission, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PermissionCreate) createSpec() (*Permission, *sqlgraph.CreateSpec) {
	var (
		_node = &Permission{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(permission.Table, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreatedBy(); ok {
		_spec.SetField(permission.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(permission.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedBy(); ok {
		_spec.SetField(permission.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.PrincipalKind(); ok {
		_spec.SetField(permission.FieldPrincipalKind, field.TypeEnum, value)
		_node.PrincipalKind = value
	}
	if value, ok := pc.mutation.StartAt(); ok {
		_spec.SetField(permission.FieldStartAt, field.TypeTime, value)
		_node.StartAt = value
	}
	if value, ok := pc.mutation.EndAt(); ok {
		_spec.SetField(permission.FieldEndAt, field.TypeTime, value)
		_node.EndAt = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(permission.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := pc.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.OrgTable,
			Columns: []string{permission.OrgColumn},
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
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permission.UserTable,
			Columns: []string{permission.UserColumn},
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
	if nodes := pc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   permission.RoleTable,
			Columns: []string{permission.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RoleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OrgPolicyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.OrgPolicyTable,
			Columns: []string{permission.OrgPolicyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgpolicy.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrgPolicyID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Permission.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PermissionUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (pc *PermissionCreate) OnConflict(opts ...sql.ConflictOption) *PermissionUpsertOne {
	pc.conflict = opts
	return &PermissionUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PermissionCreate) OnConflictColumns(columns ...string) *PermissionUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PermissionUpsertOne{
		create: pc,
	}
}

type (
	// PermissionUpsertOne is the builder for "upsert"-ing
	//  one Permission node.
	PermissionUpsertOne struct {
		create *PermissionCreate
	}

	// PermissionUpsert is the "OnConflict" setter.
	PermissionUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedBy sets the "updated_by" field.
func (u *PermissionUpsert) SetUpdatedBy(v int) *PermissionUpsert {
	u.Set(permission.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateUpdatedBy() *PermissionUpsert {
	u.SetExcluded(permission.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *PermissionUpsert) AddUpdatedBy(v int) *PermissionUpsert {
	u.Add(permission.FieldUpdatedBy, v)
	return u
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *PermissionUpsert) ClearUpdatedBy() *PermissionUpsert {
	u.SetNull(permission.FieldUpdatedBy)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PermissionUpsert) SetUpdatedAt(v time.Time) *PermissionUpsert {
	u.Set(permission.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateUpdatedAt() *PermissionUpsert {
	u.SetExcluded(permission.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *PermissionUpsert) ClearUpdatedAt() *PermissionUpsert {
	u.SetNull(permission.FieldUpdatedAt)
	return u
}

// SetStartAt sets the "start_at" field.
func (u *PermissionUpsert) SetStartAt(v time.Time) *PermissionUpsert {
	u.Set(permission.FieldStartAt, v)
	return u
}

// UpdateStartAt sets the "start_at" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateStartAt() *PermissionUpsert {
	u.SetExcluded(permission.FieldStartAt)
	return u
}

// ClearStartAt clears the value of the "start_at" field.
func (u *PermissionUpsert) ClearStartAt() *PermissionUpsert {
	u.SetNull(permission.FieldStartAt)
	return u
}

// SetEndAt sets the "end_at" field.
func (u *PermissionUpsert) SetEndAt(v time.Time) *PermissionUpsert {
	u.Set(permission.FieldEndAt, v)
	return u
}

// UpdateEndAt sets the "end_at" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateEndAt() *PermissionUpsert {
	u.SetExcluded(permission.FieldEndAt)
	return u
}

// ClearEndAt clears the value of the "end_at" field.
func (u *PermissionUpsert) ClearEndAt() *PermissionUpsert {
	u.SetNull(permission.FieldEndAt)
	return u
}

// SetStatus sets the "status" field.
func (u *PermissionUpsert) SetStatus(v typex.SimpleStatus) *PermissionUpsert {
	u.Set(permission.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PermissionUpsert) UpdateStatus() *PermissionUpsert {
	u.SetExcluded(permission.FieldStatus)
	return u
}

// ClearStatus clears the value of the "status" field.
func (u *PermissionUpsert) ClearStatus() *PermissionUpsert {
	u.SetNull(permission.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(permission.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PermissionUpsertOne) UpdateNewValues() *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(permission.FieldID)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(permission.FieldCreatedBy)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(permission.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.OrgID(); exists {
			s.SetIgnore(permission.FieldOrgID)
		}
		if _, exists := u.create.mutation.PrincipalKind(); exists {
			s.SetIgnore(permission.FieldPrincipalKind)
		}
		if _, exists := u.create.mutation.UserID(); exists {
			s.SetIgnore(permission.FieldUserID)
		}
		if _, exists := u.create.mutation.RoleID(); exists {
			s.SetIgnore(permission.FieldRoleID)
		}
		if _, exists := u.create.mutation.OrgPolicyID(); exists {
			s.SetIgnore(permission.FieldOrgPolicyID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Permission.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PermissionUpsertOne) Ignore() *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PermissionUpsertOne) DoNothing() *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PermissionCreate.OnConflict
// documentation for more info.
func (u *PermissionUpsertOne) Update(set func(*PermissionUpsert)) *PermissionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PermissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *PermissionUpsertOne) SetUpdatedBy(v int) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *PermissionUpsertOne) AddUpdatedBy(v int) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateUpdatedBy() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *PermissionUpsertOne) ClearUpdatedBy() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PermissionUpsertOne) SetUpdatedAt(v time.Time) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateUpdatedAt() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *PermissionUpsertOne) ClearUpdatedAt() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetStartAt sets the "start_at" field.
func (u *PermissionUpsertOne) SetStartAt(v time.Time) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetStartAt(v)
	})
}

// UpdateStartAt sets the "start_at" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateStartAt() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateStartAt()
	})
}

// ClearStartAt clears the value of the "start_at" field.
func (u *PermissionUpsertOne) ClearStartAt() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearStartAt()
	})
}

// SetEndAt sets the "end_at" field.
func (u *PermissionUpsertOne) SetEndAt(v time.Time) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetEndAt(v)
	})
}

// UpdateEndAt sets the "end_at" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateEndAt() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateEndAt()
	})
}

// ClearEndAt clears the value of the "end_at" field.
func (u *PermissionUpsertOne) ClearEndAt() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearEndAt()
	})
}

// SetStatus sets the "status" field.
func (u *PermissionUpsertOne) SetStatus(v typex.SimpleStatus) *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PermissionUpsertOne) UpdateStatus() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *PermissionUpsertOne) ClearStatus() *PermissionUpsertOne {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *PermissionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PermissionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PermissionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PermissionUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PermissionUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PermissionCreateBulk is the builder for creating many Permission entities in bulk.
type PermissionCreateBulk struct {
	config
	err      error
	builders []*PermissionCreate
	conflict []sql.ConflictOption
}

// Save creates the Permission entities in the database.
func (pcb *PermissionCreateBulk) Save(ctx context.Context) ([]*Permission, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Permission, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PermissionMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PermissionCreateBulk) SaveX(ctx context.Context) []*Permission {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PermissionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PermissionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Permission.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PermissionUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (pcb *PermissionCreateBulk) OnConflict(opts ...sql.ConflictOption) *PermissionUpsertBulk {
	pcb.conflict = opts
	return &PermissionUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PermissionCreateBulk) OnConflictColumns(columns ...string) *PermissionUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PermissionUpsertBulk{
		create: pcb,
	}
}

// PermissionUpsertBulk is the builder for "upsert"-ing
// a bulk of Permission nodes.
type PermissionUpsertBulk struct {
	create *PermissionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(permission.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PermissionUpsertBulk) UpdateNewValues() *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(permission.FieldID)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(permission.FieldCreatedBy)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(permission.FieldCreatedAt)
			}
			if _, exists := b.mutation.OrgID(); exists {
				s.SetIgnore(permission.FieldOrgID)
			}
			if _, exists := b.mutation.PrincipalKind(); exists {
				s.SetIgnore(permission.FieldPrincipalKind)
			}
			if _, exists := b.mutation.UserID(); exists {
				s.SetIgnore(permission.FieldUserID)
			}
			if _, exists := b.mutation.RoleID(); exists {
				s.SetIgnore(permission.FieldRoleID)
			}
			if _, exists := b.mutation.OrgPolicyID(); exists {
				s.SetIgnore(permission.FieldOrgPolicyID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Permission.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PermissionUpsertBulk) Ignore() *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PermissionUpsertBulk) DoNothing() *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PermissionCreateBulk.OnConflict
// documentation for more info.
func (u *PermissionUpsertBulk) Update(set func(*PermissionUpsert)) *PermissionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PermissionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *PermissionUpsertBulk) SetUpdatedBy(v int) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *PermissionUpsertBulk) AddUpdatedBy(v int) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateUpdatedBy() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateUpdatedBy()
	})
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (u *PermissionUpsertBulk) ClearUpdatedBy() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PermissionUpsertBulk) SetUpdatedAt(v time.Time) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateUpdatedAt() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *PermissionUpsertBulk) ClearUpdatedAt() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetStartAt sets the "start_at" field.
func (u *PermissionUpsertBulk) SetStartAt(v time.Time) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetStartAt(v)
	})
}

// UpdateStartAt sets the "start_at" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateStartAt() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateStartAt()
	})
}

// ClearStartAt clears the value of the "start_at" field.
func (u *PermissionUpsertBulk) ClearStartAt() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearStartAt()
	})
}

// SetEndAt sets the "end_at" field.
func (u *PermissionUpsertBulk) SetEndAt(v time.Time) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetEndAt(v)
	})
}

// UpdateEndAt sets the "end_at" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateEndAt() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateEndAt()
	})
}

// ClearEndAt clears the value of the "end_at" field.
func (u *PermissionUpsertBulk) ClearEndAt() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearEndAt()
	})
}

// SetStatus sets the "status" field.
func (u *PermissionUpsertBulk) SetStatus(v typex.SimpleStatus) *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PermissionUpsertBulk) UpdateStatus() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.UpdateStatus()
	})
}

// ClearStatus clears the value of the "status" field.
func (u *PermissionUpsertBulk) ClearStatus() *PermissionUpsertBulk {
	return u.Update(func(s *PermissionUpsert) {
		s.ClearStatus()
	})
}

// Exec executes the query.
func (u *PermissionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PermissionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PermissionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PermissionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
