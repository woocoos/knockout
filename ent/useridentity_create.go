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
	"github.com/woocoos/knockout/ent/useridentity"
)

// UserIdentityCreate is the builder for creating a UserIdentity entity.
type UserIdentityCreate struct {
	config
	mutation *UserIdentityMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (uic *UserIdentityCreate) SetCreatedBy(i int) *UserIdentityCreate {
	uic.mutation.SetCreatedBy(i)
	return uic
}

// SetCreatedAt sets the "created_at" field.
func (uic *UserIdentityCreate) SetCreatedAt(t time.Time) *UserIdentityCreate {
	uic.mutation.SetCreatedAt(t)
	return uic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uic *UserIdentityCreate) SetNillableCreatedAt(t *time.Time) *UserIdentityCreate {
	if t != nil {
		uic.SetCreatedAt(*t)
	}
	return uic
}

// SetUpdatedBy sets the "updated_by" field.
func (uic *UserIdentityCreate) SetUpdatedBy(i int) *UserIdentityCreate {
	uic.mutation.SetUpdatedBy(i)
	return uic
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uic *UserIdentityCreate) SetNillableUpdatedBy(i *int) *UserIdentityCreate {
	if i != nil {
		uic.SetUpdatedBy(*i)
	}
	return uic
}

// SetUpdatedAt sets the "updated_at" field.
func (uic *UserIdentityCreate) SetUpdatedAt(t time.Time) *UserIdentityCreate {
	uic.mutation.SetUpdatedAt(t)
	return uic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uic *UserIdentityCreate) SetNillableUpdatedAt(t *time.Time) *UserIdentityCreate {
	if t != nil {
		uic.SetUpdatedAt(*t)
	}
	return uic
}

// SetUserID sets the "user_id" field.
func (uic *UserIdentityCreate) SetUserID(i int) *UserIdentityCreate {
	uic.mutation.SetUserID(i)
	return uic
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (uic *UserIdentityCreate) SetNillableUserID(i *int) *UserIdentityCreate {
	if i != nil {
		uic.SetUserID(*i)
	}
	return uic
}

// SetKind sets the "kind" field.
func (uic *UserIdentityCreate) SetKind(u useridentity.Kind) *UserIdentityCreate {
	uic.mutation.SetKind(u)
	return uic
}

// SetCode sets the "code" field.
func (uic *UserIdentityCreate) SetCode(s string) *UserIdentityCreate {
	uic.mutation.SetCode(s)
	return uic
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (uic *UserIdentityCreate) SetNillableCode(s *string) *UserIdentityCreate {
	if s != nil {
		uic.SetCode(*s)
	}
	return uic
}

// SetCodeExtend sets the "code_extend" field.
func (uic *UserIdentityCreate) SetCodeExtend(s string) *UserIdentityCreate {
	uic.mutation.SetCodeExtend(s)
	return uic
}

// SetNillableCodeExtend sets the "code_extend" field if the given value is not nil.
func (uic *UserIdentityCreate) SetNillableCodeExtend(s *string) *UserIdentityCreate {
	if s != nil {
		uic.SetCodeExtend(*s)
	}
	return uic
}

// SetStatus sets the "status" field.
func (uic *UserIdentityCreate) SetStatus(ts typex.SimpleStatus) *UserIdentityCreate {
	uic.mutation.SetStatus(ts)
	return uic
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uic *UserIdentityCreate) SetNillableStatus(ts *typex.SimpleStatus) *UserIdentityCreate {
	if ts != nil {
		uic.SetStatus(*ts)
	}
	return uic
}

// SetID sets the "id" field.
func (uic *UserIdentityCreate) SetID(i int) *UserIdentityCreate {
	uic.mutation.SetID(i)
	return uic
}

// SetUser sets the "user" edge to the User entity.
func (uic *UserIdentityCreate) SetUser(u *User) *UserIdentityCreate {
	return uic.SetUserID(u.ID)
}

// Mutation returns the UserIdentityMutation object of the builder.
func (uic *UserIdentityCreate) Mutation() *UserIdentityMutation {
	return uic.mutation
}

// Save creates the UserIdentity in the database.
func (uic *UserIdentityCreate) Save(ctx context.Context) (*UserIdentity, error) {
	if err := uic.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*UserIdentity, UserIdentityMutation](ctx, uic.sqlSave, uic.mutation, uic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uic *UserIdentityCreate) SaveX(ctx context.Context) *UserIdentity {
	v, err := uic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uic *UserIdentityCreate) Exec(ctx context.Context) error {
	_, err := uic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uic *UserIdentityCreate) ExecX(ctx context.Context) {
	if err := uic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uic *UserIdentityCreate) defaults() error {
	if _, ok := uic.mutation.CreatedAt(); !ok {
		if useridentity.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized useridentity.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := useridentity.DefaultCreatedAt()
		uic.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uic *UserIdentityCreate) check() error {
	if _, ok := uic.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "UserIdentity.created_by"`)}
	}
	if _, ok := uic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserIdentity.created_at"`)}
	}
	if _, ok := uic.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "UserIdentity.kind"`)}
	}
	if v, ok := uic.mutation.Kind(); ok {
		if err := useridentity.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "UserIdentity.kind": %w`, err)}
		}
	}
	if v, ok := uic.mutation.Status(); ok {
		if err := useridentity.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserIdentity.status": %w`, err)}
		}
	}
	return nil
}

func (uic *UserIdentityCreate) sqlSave(ctx context.Context) (*UserIdentity, error) {
	if err := uic.check(); err != nil {
		return nil, err
	}
	_node, _spec := uic.createSpec()
	if err := sqlgraph.CreateNode(ctx, uic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	uic.mutation.id = &_node.ID
	uic.mutation.done = true
	return _node, nil
}

func (uic *UserIdentityCreate) createSpec() (*UserIdentity, *sqlgraph.CreateSpec) {
	var (
		_node = &UserIdentity{config: uic.config}
		_spec = sqlgraph.NewCreateSpec(useridentity.Table, sqlgraph.NewFieldSpec(useridentity.FieldID, field.TypeInt))
	)
	if id, ok := uic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uic.mutation.CreatedBy(); ok {
		_spec.SetField(useridentity.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := uic.mutation.CreatedAt(); ok {
		_spec.SetField(useridentity.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uic.mutation.UpdatedBy(); ok {
		_spec.SetField(useridentity.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := uic.mutation.UpdatedAt(); ok {
		_spec.SetField(useridentity.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uic.mutation.Kind(); ok {
		_spec.SetField(useridentity.FieldKind, field.TypeEnum, value)
		_node.Kind = value
	}
	if value, ok := uic.mutation.Code(); ok {
		_spec.SetField(useridentity.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := uic.mutation.CodeExtend(); ok {
		_spec.SetField(useridentity.FieldCodeExtend, field.TypeString, value)
		_node.CodeExtend = value
	}
	if value, ok := uic.mutation.Status(); ok {
		_spec.SetField(useridentity.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := uic.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useridentity.UserTable,
			Columns: []string{useridentity.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
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

// UserIdentityCreateBulk is the builder for creating many UserIdentity entities in bulk.
type UserIdentityCreateBulk struct {
	config
	builders []*UserIdentityCreate
}

// Save creates the UserIdentity entities in the database.
func (uicb *UserIdentityCreateBulk) Save(ctx context.Context) ([]*UserIdentity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uicb.builders))
	nodes := make([]*UserIdentity, len(uicb.builders))
	mutators := make([]Mutator, len(uicb.builders))
	for i := range uicb.builders {
		func(i int, root context.Context) {
			builder := uicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserIdentityMutation)
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
					_, err = mutators[i+1].Mutate(root, uicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uicb *UserIdentityCreateBulk) SaveX(ctx context.Context) []*UserIdentity {
	v, err := uicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uicb *UserIdentityCreateBulk) Exec(ctx context.Context) error {
	_, err := uicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uicb *UserIdentityCreateBulk) ExecX(ctx context.Context) {
	if err := uicb.Exec(ctx); err != nil {
		panic(err)
	}
}
