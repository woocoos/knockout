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
)

// OrgRoleCreate is the builder for creating a OrgRole entity.
type OrgRoleCreate struct {
	config
	mutation *OrgRoleMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (orc *OrgRoleCreate) SetCreatedBy(i int) *OrgRoleCreate {
	orc.mutation.SetCreatedBy(i)
	return orc
}

// SetCreatedAt sets the "created_at" field.
func (orc *OrgRoleCreate) SetCreatedAt(t time.Time) *OrgRoleCreate {
	orc.mutation.SetCreatedAt(t)
	return orc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (orc *OrgRoleCreate) SetNillableCreatedAt(t *time.Time) *OrgRoleCreate {
	if t != nil {
		orc.SetCreatedAt(*t)
	}
	return orc
}

// SetUpdatedBy sets the "updated_by" field.
func (orc *OrgRoleCreate) SetUpdatedBy(i int) *OrgRoleCreate {
	orc.mutation.SetUpdatedBy(i)
	return orc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (orc *OrgRoleCreate) SetNillableUpdatedBy(i *int) *OrgRoleCreate {
	if i != nil {
		orc.SetUpdatedBy(*i)
	}
	return orc
}

// SetUpdatedAt sets the "updated_at" field.
func (orc *OrgRoleCreate) SetUpdatedAt(t time.Time) *OrgRoleCreate {
	orc.mutation.SetUpdatedAt(t)
	return orc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (orc *OrgRoleCreate) SetNillableUpdatedAt(t *time.Time) *OrgRoleCreate {
	if t != nil {
		orc.SetUpdatedAt(*t)
	}
	return orc
}

// SetOrgID sets the "org_id" field.
func (orc *OrgRoleCreate) SetOrgID(i int) *OrgRoleCreate {
	orc.mutation.SetOrgID(i)
	return orc
}

// SetKind sets the "kind" field.
func (orc *OrgRoleCreate) SetKind(o orgrole.Kind) *OrgRoleCreate {
	orc.mutation.SetKind(o)
	return orc
}

// SetName sets the "name" field.
func (orc *OrgRoleCreate) SetName(s string) *OrgRoleCreate {
	orc.mutation.SetName(s)
	return orc
}

// SetAppRoleID sets the "app_role_id" field.
func (orc *OrgRoleCreate) SetAppRoleID(i int) *OrgRoleCreate {
	orc.mutation.SetAppRoleID(i)
	return orc
}

// SetNillableAppRoleID sets the "app_role_id" field if the given value is not nil.
func (orc *OrgRoleCreate) SetNillableAppRoleID(i *int) *OrgRoleCreate {
	if i != nil {
		orc.SetAppRoleID(*i)
	}
	return orc
}

// SetComments sets the "comments" field.
func (orc *OrgRoleCreate) SetComments(s string) *OrgRoleCreate {
	orc.mutation.SetComments(s)
	return orc
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (orc *OrgRoleCreate) SetNillableComments(s *string) *OrgRoleCreate {
	if s != nil {
		orc.SetComments(*s)
	}
	return orc
}

// SetOrg sets the "org" edge to the Org entity.
func (orc *OrgRoleCreate) SetOrg(o *Org) *OrgRoleCreate {
	return orc.SetOrgID(o.ID)
}

// AddOrgUserIDs adds the "org_users" edge to the OrgUser entity by IDs.
func (orc *OrgRoleCreate) AddOrgUserIDs(ids ...int) *OrgRoleCreate {
	orc.mutation.AddOrgUserIDs(ids...)
	return orc
}

// AddOrgUsers adds the "org_users" edges to the OrgUser entity.
func (orc *OrgRoleCreate) AddOrgUsers(o ...*OrgUser) *OrgRoleCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return orc.AddOrgUserIDs(ids...)
}

// Mutation returns the OrgRoleMutation object of the builder.
func (orc *OrgRoleCreate) Mutation() *OrgRoleMutation {
	return orc.mutation
}

// Save creates the OrgRole in the database.
func (orc *OrgRoleCreate) Save(ctx context.Context) (*OrgRole, error) {
	if err := orc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*OrgRole, OrgRoleMutation](ctx, orc.sqlSave, orc.mutation, orc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (orc *OrgRoleCreate) SaveX(ctx context.Context) *OrgRole {
	v, err := orc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (orc *OrgRoleCreate) Exec(ctx context.Context) error {
	_, err := orc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (orc *OrgRoleCreate) ExecX(ctx context.Context) {
	if err := orc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (orc *OrgRoleCreate) defaults() error {
	if _, ok := orc.mutation.CreatedAt(); !ok {
		if orgrole.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized orgrole.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := orgrole.DefaultCreatedAt()
		orc.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (orc *OrgRoleCreate) check() error {
	if _, ok := orc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "OrgRole.created_by"`)}
	}
	if _, ok := orc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrgRole.created_at"`)}
	}
	if _, ok := orc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "OrgRole.org_id"`)}
	}
	if _, ok := orc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "OrgRole.kind"`)}
	}
	if v, ok := orc.mutation.Kind(); ok {
		if err := orgrole.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "OrgRole.kind": %w`, err)}
		}
	}
	if _, ok := orc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "OrgRole.name"`)}
	}
	if _, ok := orc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org", err: errors.New(`ent: missing required edge "OrgRole.org"`)}
	}
	return nil
}

func (orc *OrgRoleCreate) sqlSave(ctx context.Context) (*OrgRole, error) {
	if err := orc.check(); err != nil {
		return nil, err
	}
	_node, _spec := orc.createSpec()
	if err := sqlgraph.CreateNode(ctx, orc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	orc.mutation.id = &_node.ID
	orc.mutation.done = true
	return _node, nil
}

func (orc *OrgRoleCreate) createSpec() (*OrgRole, *sqlgraph.CreateSpec) {
	var (
		_node = &OrgRole{config: orc.config}
		_spec = sqlgraph.NewCreateSpec(orgrole.Table, sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt))
	)
	if value, ok := orc.mutation.CreatedBy(); ok {
		_spec.SetField(orgrole.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := orc.mutation.CreatedAt(); ok {
		_spec.SetField(orgrole.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := orc.mutation.UpdatedBy(); ok {
		_spec.SetField(orgrole.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := orc.mutation.UpdatedAt(); ok {
		_spec.SetField(orgrole.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := orc.mutation.Kind(); ok {
		_spec.SetField(orgrole.FieldKind, field.TypeEnum, value)
		_node.Kind = value
	}
	if value, ok := orc.mutation.Name(); ok {
		_spec.SetField(orgrole.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := orc.mutation.AppRoleID(); ok {
		_spec.SetField(orgrole.FieldAppRoleID, field.TypeInt, value)
		_node.AppRoleID = value
	}
	if value, ok := orc.mutation.Comments(); ok {
		_spec.SetField(orgrole.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if nodes := orc.mutation.OrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orgrole.OrgTable,
			Columns: []string{orgrole.OrgColumn},
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
	if nodes := orc.mutation.OrgUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   orgrole.OrgUsersTable,
			Columns: orgrole.OrgUsersPrimaryKey,
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
		createE := &OrgRoleUserCreate{config: orc.config, mutation: newOrgRoleUserMutation(orc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrgRoleCreateBulk is the builder for creating many OrgRole entities in bulk.
type OrgRoleCreateBulk struct {
	config
	builders []*OrgRoleCreate
}

// Save creates the OrgRole entities in the database.
func (orcb *OrgRoleCreateBulk) Save(ctx context.Context) ([]*OrgRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(orcb.builders))
	nodes := make([]*OrgRole, len(orcb.builders))
	mutators := make([]Mutator, len(orcb.builders))
	for i := range orcb.builders {
		func(i int, root context.Context) {
			builder := orcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrgRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, orcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, orcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, orcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (orcb *OrgRoleCreateBulk) SaveX(ctx context.Context) []*OrgRole {
	v, err := orcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (orcb *OrgRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := orcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (orcb *OrgRoleCreateBulk) ExecX(ctx context.Context) {
	if err := orcb.Exec(ctx); err != nil {
		panic(err)
	}
}
