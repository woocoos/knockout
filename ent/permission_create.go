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
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/user"
)

// PermissionCreate is the builder for creating a Permission entity.
type PermissionCreate struct {
	config
	mutation *PermissionMutation
	hooks    []Hook
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
	return withHooks[*Permission, PermissionMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
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
	if value, ok := pc.mutation.RoleID(); ok {
		_spec.SetField(permission.FieldRoleID, field.TypeInt, value)
		_node.RoleID = value
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

// PermissionCreateBulk is the builder for creating many Permission entities in bulk.
type PermissionCreateBulk struct {
	config
	builders []*PermissionCreate
}

// Save creates the Permission entities in the database.
func (pcb *PermissionCreateBulk) Save(ctx context.Context) ([]*Permission, error) {
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
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
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
