// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/organization"
	"github.com/woocoos/knockout/ent/organizationrole"
)

// OrganizationRoleCreate is the builder for creating a OrganizationRole entity.
type OrganizationRoleCreate struct {
	config
	mutation *OrganizationRoleMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (orc *OrganizationRoleCreate) SetCreatedBy(i int) *OrganizationRoleCreate {
	orc.mutation.SetCreatedBy(i)
	return orc
}

// SetCreatedAt sets the "created_at" field.
func (orc *OrganizationRoleCreate) SetCreatedAt(t time.Time) *OrganizationRoleCreate {
	orc.mutation.SetCreatedAt(t)
	return orc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (orc *OrganizationRoleCreate) SetNillableCreatedAt(t *time.Time) *OrganizationRoleCreate {
	if t != nil {
		orc.SetCreatedAt(*t)
	}
	return orc
}

// SetUpdatedBy sets the "updated_by" field.
func (orc *OrganizationRoleCreate) SetUpdatedBy(i int) *OrganizationRoleCreate {
	orc.mutation.SetUpdatedBy(i)
	return orc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (orc *OrganizationRoleCreate) SetNillableUpdatedBy(i *int) *OrganizationRoleCreate {
	if i != nil {
		orc.SetUpdatedBy(*i)
	}
	return orc
}

// SetUpdatedAt sets the "updated_at" field.
func (orc *OrganizationRoleCreate) SetUpdatedAt(t time.Time) *OrganizationRoleCreate {
	orc.mutation.SetUpdatedAt(t)
	return orc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (orc *OrganizationRoleCreate) SetNillableUpdatedAt(t *time.Time) *OrganizationRoleCreate {
	if t != nil {
		orc.SetUpdatedAt(*t)
	}
	return orc
}

// SetOrgID sets the "org_id" field.
func (orc *OrganizationRoleCreate) SetOrgID(i int) *OrganizationRoleCreate {
	orc.mutation.SetOrgID(i)
	return orc
}

// SetKind sets the "kind" field.
func (orc *OrganizationRoleCreate) SetKind(o organizationrole.Kind) *OrganizationRoleCreate {
	orc.mutation.SetKind(o)
	return orc
}

// SetName sets the "name" field.
func (orc *OrganizationRoleCreate) SetName(s string) *OrganizationRoleCreate {
	orc.mutation.SetName(s)
	return orc
}

// SetAppRoleID sets the "app_role_id" field.
func (orc *OrganizationRoleCreate) SetAppRoleID(i int) *OrganizationRoleCreate {
	orc.mutation.SetAppRoleID(i)
	return orc
}

// SetNillableAppRoleID sets the "app_role_id" field if the given value is not nil.
func (orc *OrganizationRoleCreate) SetNillableAppRoleID(i *int) *OrganizationRoleCreate {
	if i != nil {
		orc.SetAppRoleID(*i)
	}
	return orc
}

// SetComments sets the "comments" field.
func (orc *OrganizationRoleCreate) SetComments(s string) *OrganizationRoleCreate {
	orc.mutation.SetComments(s)
	return orc
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (orc *OrganizationRoleCreate) SetNillableComments(s *string) *OrganizationRoleCreate {
	if s != nil {
		orc.SetComments(*s)
	}
	return orc
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (orc *OrganizationRoleCreate) SetOrganizationID(id int) *OrganizationRoleCreate {
	orc.mutation.SetOrganizationID(id)
	return orc
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (orc *OrganizationRoleCreate) SetOrganization(o *Organization) *OrganizationRoleCreate {
	return orc.SetOrganizationID(o.ID)
}

// Mutation returns the OrganizationRoleMutation object of the builder.
func (orc *OrganizationRoleCreate) Mutation() *OrganizationRoleMutation {
	return orc.mutation
}

// Save creates the OrganizationRole in the database.
func (orc *OrganizationRoleCreate) Save(ctx context.Context) (*OrganizationRole, error) {
	if err := orc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*OrganizationRole, OrganizationRoleMutation](ctx, orc.sqlSave, orc.mutation, orc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (orc *OrganizationRoleCreate) SaveX(ctx context.Context) *OrganizationRole {
	v, err := orc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (orc *OrganizationRoleCreate) Exec(ctx context.Context) error {
	_, err := orc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (orc *OrganizationRoleCreate) ExecX(ctx context.Context) {
	if err := orc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (orc *OrganizationRoleCreate) defaults() error {
	if _, ok := orc.mutation.CreatedAt(); !ok {
		if organizationrole.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized organizationrole.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := organizationrole.DefaultCreatedAt()
		orc.mutation.SetCreatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (orc *OrganizationRoleCreate) check() error {
	if _, ok := orc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "OrganizationRole.created_by"`)}
	}
	if _, ok := orc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrganizationRole.created_at"`)}
	}
	if _, ok := orc.mutation.OrgID(); !ok {
		return &ValidationError{Name: "org_id", err: errors.New(`ent: missing required field "OrganizationRole.org_id"`)}
	}
	if _, ok := orc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "OrganizationRole.kind"`)}
	}
	if v, ok := orc.mutation.Kind(); ok {
		if err := organizationrole.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "OrganizationRole.kind": %w`, err)}
		}
	}
	if _, ok := orc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "OrganizationRole.name"`)}
	}
	if _, ok := orc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`ent: missing required edge "OrganizationRole.organization"`)}
	}
	return nil
}

func (orc *OrganizationRoleCreate) sqlSave(ctx context.Context) (*OrganizationRole, error) {
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

func (orc *OrganizationRoleCreate) createSpec() (*OrganizationRole, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationRole{config: orc.config}
		_spec = sqlgraph.NewCreateSpec(organizationrole.Table, sqlgraph.NewFieldSpec(organizationrole.FieldID, field.TypeInt))
	)
	if value, ok := orc.mutation.CreatedBy(); ok {
		_spec.SetField(organizationrole.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := orc.mutation.CreatedAt(); ok {
		_spec.SetField(organizationrole.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := orc.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationrole.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := orc.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationrole.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := orc.mutation.Kind(); ok {
		_spec.SetField(organizationrole.FieldKind, field.TypeEnum, value)
		_node.Kind = value
	}
	if value, ok := orc.mutation.Name(); ok {
		_spec.SetField(organizationrole.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := orc.mutation.AppRoleID(); ok {
		_spec.SetField(organizationrole.FieldAppRoleID, field.TypeInt, value)
		_node.AppRoleID = value
	}
	if value, ok := orc.mutation.Comments(); ok {
		_spec.SetField(organizationrole.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if nodes := orc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organizationrole.OrganizationTable,
			Columns: []string{organizationrole.OrganizationColumn},
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
		_node.OrgID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrganizationRoleCreateBulk is the builder for creating many OrganizationRole entities in bulk.
type OrganizationRoleCreateBulk struct {
	config
	builders []*OrganizationRoleCreate
}

// Save creates the OrganizationRole entities in the database.
func (orcb *OrganizationRoleCreateBulk) Save(ctx context.Context) ([]*OrganizationRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(orcb.builders))
	nodes := make([]*OrganizationRole, len(orcb.builders))
	mutators := make([]Mutator, len(orcb.builders))
	for i := range orcb.builders {
		func(i int, root context.Context) {
			builder := orcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationRoleMutation)
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
func (orcb *OrganizationRoleCreateBulk) SaveX(ctx context.Context) []*OrganizationRole {
	v, err := orcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (orcb *OrganizationRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := orcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (orcb *OrganizationRoleCreateBulk) ExecX(ctx context.Context) {
	if err := orcb.Exec(ctx); err != nil {
		panic(err)
	}
}
