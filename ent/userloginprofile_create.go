// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/userloginprofile"
)

// UserLoginProfileCreate is the builder for creating a UserLoginProfile entity.
type UserLoginProfileCreate struct {
	config
	mutation *UserLoginProfileMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (ulpc *UserLoginProfileCreate) SetCreatedBy(i int) *UserLoginProfileCreate {
	ulpc.mutation.SetCreatedBy(i)
	return ulpc
}

// SetCreatedAt sets the "created_at" field.
func (ulpc *UserLoginProfileCreate) SetCreatedAt(t time.Time) *UserLoginProfileCreate {
	ulpc.mutation.SetCreatedAt(t)
	return ulpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ulpc *UserLoginProfileCreate) SetNillableCreatedAt(t *time.Time) *UserLoginProfileCreate {
	if t != nil {
		ulpc.SetCreatedAt(*t)
	}
	return ulpc
}

// SetUpdatedBy sets the "updated_by" field.
func (ulpc *UserLoginProfileCreate) SetUpdatedBy(i int) *UserLoginProfileCreate {
	ulpc.mutation.SetUpdatedBy(i)
	return ulpc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ulpc *UserLoginProfileCreate) SetNillableUpdatedBy(i *int) *UserLoginProfileCreate {
	if i != nil {
		ulpc.SetUpdatedBy(*i)
	}
	return ulpc
}

// SetUpdatedAt sets the "updated_at" field.
func (ulpc *UserLoginProfileCreate) SetUpdatedAt(t time.Time) *UserLoginProfileCreate {
	ulpc.mutation.SetUpdatedAt(t)
	return ulpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ulpc *UserLoginProfileCreate) SetNillableUpdatedAt(t *time.Time) *UserLoginProfileCreate {
	if t != nil {
		ulpc.SetUpdatedAt(*t)
	}
	return ulpc
}

// SetUserID sets the "user_id" field.
func (ulpc *UserLoginProfileCreate) SetUserID(i int) *UserLoginProfileCreate {
	ulpc.mutation.SetUserID(i)
	return ulpc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ulpc *UserLoginProfileCreate) SetNillableUserID(i *int) *UserLoginProfileCreate {
	if i != nil {
		ulpc.SetUserID(*i)
	}
	return ulpc
}

// SetLastLoginIP sets the "last_login_ip" field.
func (ulpc *UserLoginProfileCreate) SetLastLoginIP(s string) *UserLoginProfileCreate {
	ulpc.mutation.SetLastLoginIP(s)
	return ulpc
}

// SetNillableLastLoginIP sets the "last_login_ip" field if the given value is not nil.
func (ulpc *UserLoginProfileCreate) SetNillableLastLoginIP(s *string) *UserLoginProfileCreate {
	if s != nil {
		ulpc.SetLastLoginIP(*s)
	}
	return ulpc
}

// SetLastLoginAt sets the "last_login_at" field.
func (ulpc *UserLoginProfileCreate) SetLastLoginAt(t time.Time) *UserLoginProfileCreate {
	ulpc.mutation.SetLastLoginAt(t)
	return ulpc
}

// SetNillableLastLoginAt sets the "last_login_at" field if the given value is not nil.
func (ulpc *UserLoginProfileCreate) SetNillableLastLoginAt(t *time.Time) *UserLoginProfileCreate {
	if t != nil {
		ulpc.SetLastLoginAt(*t)
	}
	return ulpc
}

// SetCanLogin sets the "can_login" field.
func (ulpc *UserLoginProfileCreate) SetCanLogin(b bool) *UserLoginProfileCreate {
	ulpc.mutation.SetCanLogin(b)
	return ulpc
}

// SetNillableCanLogin sets the "can_login" field if the given value is not nil.
func (ulpc *UserLoginProfileCreate) SetNillableCanLogin(b *bool) *UserLoginProfileCreate {
	if b != nil {
		ulpc.SetCanLogin(*b)
	}
	return ulpc
}

// SetSetKind sets the "set_kind" field.
func (ulpc *UserLoginProfileCreate) SetSetKind(uk userloginprofile.SetKind) *UserLoginProfileCreate {
	ulpc.mutation.SetSetKind(uk)
	return ulpc
}

// SetPasswordReset sets the "password_reset" field.
func (ulpc *UserLoginProfileCreate) SetPasswordReset(b bool) *UserLoginProfileCreate {
	ulpc.mutation.SetPasswordReset(b)
	return ulpc
}

// SetNillablePasswordReset sets the "password_reset" field if the given value is not nil.
func (ulpc *UserLoginProfileCreate) SetNillablePasswordReset(b *bool) *UserLoginProfileCreate {
	if b != nil {
		ulpc.SetPasswordReset(*b)
	}
	return ulpc
}

// SetVerifyDevice sets the "verify_device" field.
func (ulpc *UserLoginProfileCreate) SetVerifyDevice(b bool) *UserLoginProfileCreate {
	ulpc.mutation.SetVerifyDevice(b)
	return ulpc
}

// SetMfaEnabled sets the "mfa_enabled" field.
func (ulpc *UserLoginProfileCreate) SetMfaEnabled(b bool) *UserLoginProfileCreate {
	ulpc.mutation.SetMfaEnabled(b)
	return ulpc
}

// SetNillableMfaEnabled sets the "mfa_enabled" field if the given value is not nil.
func (ulpc *UserLoginProfileCreate) SetNillableMfaEnabled(b *bool) *UserLoginProfileCreate {
	if b != nil {
		ulpc.SetMfaEnabled(*b)
	}
	return ulpc
}

// SetMfaSecret sets the "mfa_secret" field.
func (ulpc *UserLoginProfileCreate) SetMfaSecret(s string) *UserLoginProfileCreate {
	ulpc.mutation.SetMfaSecret(s)
	return ulpc
}

// SetNillableMfaSecret sets the "mfa_secret" field if the given value is not nil.
func (ulpc *UserLoginProfileCreate) SetNillableMfaSecret(s *string) *UserLoginProfileCreate {
	if s != nil {
		ulpc.SetMfaSecret(*s)
	}
	return ulpc
}

// SetMfaStatus sets the "mfa_status" field.
func (ulpc *UserLoginProfileCreate) SetMfaStatus(ts typex.SimpleStatus) *UserLoginProfileCreate {
	ulpc.mutation.SetMfaStatus(ts)
	return ulpc
}

// SetNillableMfaStatus sets the "mfa_status" field if the given value is not nil.
func (ulpc *UserLoginProfileCreate) SetNillableMfaStatus(ts *typex.SimpleStatus) *UserLoginProfileCreate {
	if ts != nil {
		ulpc.SetMfaStatus(*ts)
	}
	return ulpc
}

// SetID sets the "id" field.
func (ulpc *UserLoginProfileCreate) SetID(i int) *UserLoginProfileCreate {
	ulpc.mutation.SetID(i)
	return ulpc
}

// SetUser sets the "user" edge to the User entity.
func (ulpc *UserLoginProfileCreate) SetUser(u *User) *UserLoginProfileCreate {
	return ulpc.SetUserID(u.ID)
}

// Mutation returns the UserLoginProfileMutation object of the builder.
func (ulpc *UserLoginProfileCreate) Mutation() *UserLoginProfileMutation {
	return ulpc.mutation
}

// Save creates the UserLoginProfile in the database.
func (ulpc *UserLoginProfileCreate) Save(ctx context.Context) (*UserLoginProfile, error) {
	if err := ulpc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*UserLoginProfile, UserLoginProfileMutation](ctx, ulpc.sqlSave, ulpc.mutation, ulpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ulpc *UserLoginProfileCreate) SaveX(ctx context.Context) *UserLoginProfile {
	v, err := ulpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ulpc *UserLoginProfileCreate) Exec(ctx context.Context) error {
	_, err := ulpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulpc *UserLoginProfileCreate) ExecX(ctx context.Context) {
	if err := ulpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ulpc *UserLoginProfileCreate) defaults() error {
	if _, ok := ulpc.mutation.CreatedAt(); !ok {
		if userloginprofile.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized userloginprofile.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := userloginprofile.DefaultCreatedAt()
		ulpc.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ulpc *UserLoginProfileCreate) check() error {
	if _, ok := ulpc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "UserLoginProfile.created_by"`)}
	}
	if _, ok := ulpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserLoginProfile.created_at"`)}
	}
	if _, ok := ulpc.mutation.SetKind(); !ok {
		return &ValidationError{Name: "set_kind", err: errors.New(`ent: missing required field "UserLoginProfile.set_kind"`)}
	}
	if v, ok := ulpc.mutation.SetKind(); ok {
		if err := userloginprofile.SetKindValidator(v); err != nil {
			return &ValidationError{Name: "set_kind", err: fmt.Errorf(`ent: validator failed for field "UserLoginProfile.set_kind": %w`, err)}
		}
	}
	if _, ok := ulpc.mutation.VerifyDevice(); !ok {
		return &ValidationError{Name: "verify_device", err: errors.New(`ent: missing required field "UserLoginProfile.verify_device"`)}
	}
	if v, ok := ulpc.mutation.MfaSecret(); ok {
		if err := userloginprofile.MfaSecretValidator(v); err != nil {
			return &ValidationError{Name: "mfa_secret", err: fmt.Errorf(`ent: validator failed for field "UserLoginProfile.mfa_secret": %w`, err)}
		}
	}
	if v, ok := ulpc.mutation.MfaStatus(); ok {
		if err := userloginprofile.MfaStatusValidator(v); err != nil {
			return &ValidationError{Name: "mfa_status", err: fmt.Errorf(`ent: validator failed for field "UserLoginProfile.mfa_status": %w`, err)}
		}
	}
	return nil
}

func (ulpc *UserLoginProfileCreate) sqlSave(ctx context.Context) (*UserLoginProfile, error) {
	if err := ulpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ulpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ulpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ulpc.mutation.id = &_node.ID
	ulpc.mutation.done = true
	return _node, nil
}

func (ulpc *UserLoginProfileCreate) createSpec() (*UserLoginProfile, *sqlgraph.CreateSpec) {
	var (
		_node = &UserLoginProfile{config: ulpc.config}
		_spec = sqlgraph.NewCreateSpec(userloginprofile.Table, sqlgraph.NewFieldSpec(userloginprofile.FieldID, field.TypeInt))
	)
	if id, ok := ulpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ulpc.mutation.CreatedBy(); ok {
		_spec.SetField(userloginprofile.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := ulpc.mutation.CreatedAt(); ok {
		_spec.SetField(userloginprofile.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ulpc.mutation.UpdatedBy(); ok {
		_spec.SetField(userloginprofile.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := ulpc.mutation.UpdatedAt(); ok {
		_spec.SetField(userloginprofile.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ulpc.mutation.LastLoginIP(); ok {
		_spec.SetField(userloginprofile.FieldLastLoginIP, field.TypeString, value)
		_node.LastLoginIP = value
	}
	if value, ok := ulpc.mutation.LastLoginAt(); ok {
		_spec.SetField(userloginprofile.FieldLastLoginAt, field.TypeTime, value)
		_node.LastLoginAt = value
	}
	if value, ok := ulpc.mutation.CanLogin(); ok {
		_spec.SetField(userloginprofile.FieldCanLogin, field.TypeBool, value)
		_node.CanLogin = value
	}
	if value, ok := ulpc.mutation.SetKind(); ok {
		_spec.SetField(userloginprofile.FieldSetKind, field.TypeEnum, value)
		_node.SetKind = value
	}
	if value, ok := ulpc.mutation.PasswordReset(); ok {
		_spec.SetField(userloginprofile.FieldPasswordReset, field.TypeBool, value)
		_node.PasswordReset = value
	}
	if value, ok := ulpc.mutation.VerifyDevice(); ok {
		_spec.SetField(userloginprofile.FieldVerifyDevice, field.TypeBool, value)
		_node.VerifyDevice = value
	}
	if value, ok := ulpc.mutation.MfaEnabled(); ok {
		_spec.SetField(userloginprofile.FieldMfaEnabled, field.TypeBool, value)
		_node.MfaEnabled = value
	}
	if value, ok := ulpc.mutation.MfaSecret(); ok {
		_spec.SetField(userloginprofile.FieldMfaSecret, field.TypeString, value)
		_node.MfaSecret = value
	}
	if value, ok := ulpc.mutation.MfaStatus(); ok {
		_spec.SetField(userloginprofile.FieldMfaStatus, field.TypeEnum, value)
		_node.MfaStatus = value
	}
	if nodes := ulpc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   userloginprofile.UserTable,
			Columns: []string{userloginprofile.UserColumn},
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
	return _node, _spec
}

// UserLoginProfileCreateBulk is the builder for creating many UserLoginProfile entities in bulk.
type UserLoginProfileCreateBulk struct {
	config
	builders []*UserLoginProfileCreate
}

// Save creates the UserLoginProfile entities in the database.
func (ulpcb *UserLoginProfileCreateBulk) Save(ctx context.Context) ([]*UserLoginProfile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ulpcb.builders))
	nodes := make([]*UserLoginProfile, len(ulpcb.builders))
	mutators := make([]Mutator, len(ulpcb.builders))
	for i := range ulpcb.builders {
		func(i int, root context.Context) {
			builder := ulpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserLoginProfileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ulpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ulpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ulpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ulpcb *UserLoginProfileCreateBulk) SaveX(ctx context.Context) []*UserLoginProfile {
	v, err := ulpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ulpcb *UserLoginProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := ulpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulpcb *UserLoginProfileCreateBulk) ExecX(ctx context.Context) {
	if err := ulpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
