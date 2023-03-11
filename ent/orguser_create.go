// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/user"
)

// OrgUserCreate is the builder for creating a OrgUser entity.
type OrgUserCreate struct {
	config
	mutation *OrgUserMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (ouc *OrgUserCreate) SetCreatedBy(i int) *OrgUserCreate {
	ouc.mutation.SetCreatedBy(i)
	return ouc
}

// SetCreatedAt sets the "created_at" field.
func (ouc *OrgUserCreate) SetCreatedAt(t time.Time) *OrgUserCreate {
	ouc.mutation.SetCreatedAt(t)
	return ouc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouc *OrgUserCreate) SetNillableCreatedAt(t *time.Time) *OrgUserCreate {
	if t != nil {
		ouc.SetCreatedAt(*t)
	}
	return ouc
}

// SetUpdatedBy sets the "updated_by" field.
func (ouc *OrgUserCreate) SetUpdatedBy(i int) *OrgUserCreate {
	ouc.mutation.SetUpdatedBy(i)
	return ouc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ouc *OrgUserCreate) SetNillableUpdatedBy(i *int) *OrgUserCreate {
	if i != nil {
		ouc.SetUpdatedBy(*i)
	}
	return ouc
}

// SetUpdatedAt sets the "updated_at" field.
func (ouc *OrgUserCreate) SetUpdatedAt(t time.Time) *OrgUserCreate {
	ouc.mutation.SetUpdatedAt(t)
	return ouc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ouc *OrgUserCreate) SetNillableUpdatedAt(t *time.Time) *OrgUserCreate {
	if t != nil {
		ouc.SetUpdatedAt(*t)
	}
	return ouc
}

// SetOrgID sets the "org_id" field.
func (ouc *OrgUserCreate) SetOrgID(i int) *OrgUserCreate {
	ouc.mutation.SetOrgID(i)
	return ouc
}

// SetUserID sets the "user_id" field.
func (ouc *OrgUserCreate) SetUserID(i int) *OrgUserCreate {
	ouc.mutation.SetUserID(i)
	return ouc
}

// SetDisplayName sets the "display_name" field.
func (ouc *OrgUserCreate) SetDisplayName(s string) *OrgUserCreate {
	ouc.mutation.SetDisplayName(s)
	return ouc
}

// SetID sets the "id" field.
func (ouc *OrgUserCreate) SetID(i int) *OrgUserCreate {
	ouc.mutation.SetID(i)
	return ouc
}

// SetOrg sets the "org" edge to the Org entity.
func (ouc *OrgUserCreate) SetOrg(o *Org) *OrgUserCreate {
	return ouc.SetOrgID(o.ID)
}

// SetUser sets the "user" edge to the User entity.
func (ouc *OrgUserCreate) SetUser(u *User) *OrgUserCreate {
	return ouc.SetUserID(u.ID)
}

// AddOrgRoleIDs adds the "org_roles" edge to the OrgRole entity by IDs.
func (ouc *OrgUserCreate) AddOrgRoleIDs(ids ...int) *OrgUserCreate {
	ouc.mutation.AddOrgRoleIDs(ids...)
	return ouc
}

// AddOrgRoles adds the "org_roles" edges to the OrgRole entity.
func (ouc *OrgUserCreate) AddOrgRoles(o ...*OrgRole) *OrgUserCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouc.AddOrgRoleIDs(ids...)
}

// Mutation returns the OrgUserMutation object of the builder.
func (ouc *OrgUserCreate) Mutation() *OrgUserMutation {
	return ouc.mutation
}

// Save creates the OrgUser in the database.
func (ouc *OrgUserCreate) Save(ctx context.Context) (*OrgUser, error) {
	if err := ouc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*OrgUser, OrgUserMutation](ctx, ouc.sqlSave, ouc.mutation, ouc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ouc *OrgUserCreate) SaveX(ctx context.Context) *OrgUser {
	v, err := ouc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ouc *OrgUserCreate) Exec(ctx context.Context) error {
	_, err := ouc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouc *OrgUserCreate) ExecX(ctx context.Context) {
	if err := ouc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouc *OrgUserCreate) defaults() error {
	if _, ok := ouc.mutation.CreatedAt(); !ok {
		if orguser.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized orguser.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := orguser.DefaultCreatedAt()
		ouc.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ouc *OrgUserCreate) check() error {
	if _, ok := ouc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "OrgUser.created_by"`)}
	}
	if _, ok := ouc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrgUser.created_at"`)}
	}
	if _, ok := ouc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "OrgUser.org_id"`)}
	}
	if _, ok := ouc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "OrgUser.user_id"`)}
	}
	if _, ok := ouc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "OrgUser.display_name"`)}
	}
	if _, ok := ouc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org", err: errors.New(`ent: missing required edge "OrgUser.org"`)}
	}
	if _, ok := ouc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "OrgUser.user"`)}
	}
	return nil
}

func (ouc *OrgUserCreate) sqlSave(ctx context.Context) (*OrgUser, error) {
	if err := ouc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ouc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ouc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ouc.mutation.id = &_node.ID
	ouc.mutation.done = true
	return _node, nil
}

func (ouc *OrgUserCreate) createSpec() (*OrgUser, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgUser{config: ouc.config}
		_spec = sqlgraph.NewCreateSpec(orguser.Table, sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt))
	)
	if id, ok := ouc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ouc.mutation.CreatedBy(); ok {
		_spec.SetField(orguser.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := ouc.mutation.CreatedAt(); ok {
		_spec.SetField(orguser.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ouc.mutation.UpdatedBy(); ok {
		_spec.SetField(orguser.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := ouc.mutation.UpdatedAt(); ok {
		_spec.SetField(orguser.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ouc.mutation.DisplayName(); ok {
		_spec.SetField(orguser.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if nodes := ouc.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguser.OrgTable,
			Columns: []string{orguser.OrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: org.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrgID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ouc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orguser.UserTable,
			Columns: []string{orguser.UserColumn},
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
	if nodes := ouc.mutation.OrgRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   orguser.OrgRolesTable,
			Columns: orguser.OrgRolesPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrgUserCreateBulk is the builder for creating many OrgUser entities in bulk.
type OrgUserCreateBulk struct {
	config
	builders []*OrgUserCreate
}

// Save creates the OrgUser entities in the database.
func (oucb *OrgUserCreateBulk) Save(ctx context.Context) ([]*OrgUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oucb.builders))
	nodes := make([]*OrgUser, len(oucb.builders))
	mutators := make([]Mutator, len(oucb.builders))
	for i := range oucb.builders {
		func(i int, root context.Context) {
			builder := oucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgUserMutation)
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
					_, err = mutators[i+1].Mutate(root, oucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oucb *OrgUserCreateBulk) SaveX(ctx context.Context) []*OrgUser {
	v, err := oucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oucb *OrgUserCreateBulk) Exec(ctx context.Context) error {
	_, err := oucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oucb *OrgUserCreateBulk) ExecX(ctx context.Context) {
	if err := oucb.Exec(ctx); err != nil {
		panic(err)
	}
}
