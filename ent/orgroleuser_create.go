// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orgroleuser"
	"github.com/woocoos/knockout/ent/orguser"
)

// OrgRoleUserCreate is the builder for creating a OrgRoleUser entity.
type OrgRoleUserCreate struct {
	config
	mutation *OrgRoleUserMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (oruc *OrgRoleUserCreate) SetCreatedBy(i int) *OrgRoleUserCreate {
	oruc.mutation.SetCreatedBy(i)
	return oruc
}

// SetCreatedAt sets the "created_at" field.
func (oruc *OrgRoleUserCreate) SetCreatedAt(t time.Time) *OrgRoleUserCreate {
	oruc.mutation.SetCreatedAt(t)
	return oruc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oruc *OrgRoleUserCreate) SetNillableCreatedAt(t *time.Time) *OrgRoleUserCreate {
	if t != nil {
		oruc.SetCreatedAt(*t)
	}
	return oruc
}

// SetUpdatedBy sets the "updated_by" field.
func (oruc *OrgRoleUserCreate) SetUpdatedBy(i int) *OrgRoleUserCreate {
	oruc.mutation.SetUpdatedBy(i)
	return oruc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oruc *OrgRoleUserCreate) SetNillableUpdatedBy(i *int) *OrgRoleUserCreate {
	if i != nil {
		oruc.SetUpdatedBy(*i)
	}
	return oruc
}

// SetUpdatedAt sets the "updated_at" field.
func (oruc *OrgRoleUserCreate) SetUpdatedAt(t time.Time) *OrgRoleUserCreate {
	oruc.mutation.SetUpdatedAt(t)
	return oruc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oruc *OrgRoleUserCreate) SetNillableUpdatedAt(t *time.Time) *OrgRoleUserCreate {
	if t != nil {
		oruc.SetUpdatedAt(*t)
	}
	return oruc
}

// SetOrgRoleID sets the "org_role_id" field.
func (oruc *OrgRoleUserCreate) SetOrgRoleID(i int) *OrgRoleUserCreate {
	oruc.mutation.SetOrgRoleID(i)
	return oruc
}

// SetOrgUserID sets the "org_user_id" field.
func (oruc *OrgRoleUserCreate) SetOrgUserID(i int) *OrgRoleUserCreate {
	oruc.mutation.SetOrgUserID(i)
	return oruc
}

// SetOrgRole sets the "org_role" edge to the OrgRole entity.
func (oruc *OrgRoleUserCreate) SetOrgRole(o *OrgRole) *OrgRoleUserCreate {
	return oruc.SetOrgRoleID(o.ID)
}

// SetOrgUser sets the "org_user" edge to the OrgUser entity.
func (oruc *OrgRoleUserCreate) SetOrgUser(o *OrgUser) *OrgRoleUserCreate {
	return oruc.SetOrgUserID(o.ID)
}

// Mutation returns the OrgRoleUserMutation object of the builder.
func (oruc *OrgRoleUserCreate) Mutation() *OrgRoleUserMutation {
	return oruc.mutation
}

// Save creates the OrgRoleUser in the database.
func (oruc *OrgRoleUserCreate) Save(ctx context.Context) (*OrgRoleUser, error) {
	if err := oruc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*OrgRoleUser, OrgRoleUserMutation](ctx, oruc.sqlSave, oruc.mutation, oruc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oruc *OrgRoleUserCreate) SaveX(ctx context.Context) *OrgRoleUser {
	v, err := oruc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oruc *OrgRoleUserCreate) Exec(ctx context.Context) error {
	_, err := oruc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oruc *OrgRoleUserCreate) ExecX(ctx context.Context) {
	if err := oruc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oruc *OrgRoleUserCreate) defaults() error {
	if _, ok := oruc.mutation.CreatedAt(); !ok {
		if orgroleuser.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized orgroleuser.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := orgroleuser.DefaultCreatedAt()
		oruc.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (oruc *OrgRoleUserCreate) check() error {
	if _, ok := oruc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "OrgRoleUser.created_by"`)}
	}
	if _, ok := oruc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrgRoleUser.created_at"`)}
	}
	if _, ok := oruc.mutation.OrgRoleID(); !ok {
		return &ValidationError{Name: "org_role_id", err: errors.New(`ent: missing required field "OrgRoleUser.org_role_id"`)}
	}
	if _, ok := oruc.mutation.OrgUserID(); !ok {
		return &ValidationError{Name: "org_user_id", err: errors.New(`ent: missing required field "OrgRoleUser.org_user_id"`)}
	}
	if _, ok := oruc.mutation.OrgRoleID(); !ok {
		return &ValidationError{Name: "org_role", err: errors.New(`ent: missing required edge "OrgRoleUser.org_role"`)}
	}
	if _, ok := oruc.mutation.OrgUserID(); !ok {
		return &ValidationError{Name: "org_user", err: errors.New(`ent: missing required edge "OrgRoleUser.org_user"`)}
	}
	return nil
}

func (oruc *OrgRoleUserCreate) sqlSave(ctx context.Context) (*OrgRoleUser, error) {
	if err := oruc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oruc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oruc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (oruc *OrgRoleUserCreate) createSpec() (*OrgRoleUser, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgRoleUser{config: oruc.config}
		_spec = sqlgraph.NewCreateSpec(orgroleuser.Table, nil)
	)
	if value, ok := oruc.mutation.CreatedBy(); ok {
		_spec.SetField(orgroleuser.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := oruc.mutation.CreatedAt(); ok {
		_spec.SetField(orgroleuser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oruc.mutation.UpdatedBy(); ok {
		_spec.SetField(orgroleuser.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := oruc.mutation.UpdatedAt(); ok {
		_spec.SetField(orgroleuser.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := oruc.mutation.OrgRoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgroleuser.OrgRoleTable,
			Columns: []string{orgroleuser.OrgRoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orgrole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrgRoleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oruc.mutation.OrgUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orgroleuser.OrgUserTable,
			Columns: []string{orgroleuser.OrgUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orguser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrgUserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrgRoleUserCreateBulk is the builder for creating many OrgRoleUser entities in bulk.
type OrgRoleUserCreateBulk struct {
	config
	builders []*OrgRoleUserCreate
}

// Save creates the OrgRoleUser entities in the database.
func (orucb *OrgRoleUserCreateBulk) Save(ctx context.Context) ([]*OrgRoleUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(orucb.builders))
	nodes := make([]*OrgRoleUser, len(orucb.builders))
	mutators := make([]Mutator, len(orucb.builders))
	for i := range orucb.builders {
		func(i int, root context.Context) {
			builder := orucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgRoleUserMutation)
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
					_, err = mutators[i+1].Mutate(root, orucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, orucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, orucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (orucb *OrgRoleUserCreateBulk) SaveX(ctx context.Context) []*OrgRoleUser {
	v, err := orucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (orucb *OrgRoleUserCreateBulk) Exec(ctx context.Context) error {
	_, err := orucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (orucb *OrgRoleUserCreateBulk) ExecX(ctx context.Context) {
	if err := orucb.Exec(ctx); err != nil {
		panic(err)
	}
}
