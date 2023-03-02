// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/approle"
)

// AppRoleCreate is the builder for creating a AppRole entity.
type AppRoleCreate struct {
	config
	mutation *AppRoleMutation
	hooks    []Hook
}

// SetCreatedBy sets the "created_by" field.
func (arc *AppRoleCreate) SetCreatedBy(i int) *AppRoleCreate {
	arc.mutation.SetCreatedBy(i)
	return arc
}

// SetCreatedAt sets the "created_at" field.
func (arc *AppRoleCreate) SetCreatedAt(t time.Time) *AppRoleCreate {
	arc.mutation.SetCreatedAt(t)
	return arc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableCreatedAt(t *time.Time) *AppRoleCreate {
	if t != nil {
		arc.SetCreatedAt(*t)
	}
	return arc
}

// SetUpdatedBy sets the "updated_by" field.
func (arc *AppRoleCreate) SetUpdatedBy(i int) *AppRoleCreate {
	arc.mutation.SetUpdatedBy(i)
	return arc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableUpdatedBy(i *int) *AppRoleCreate {
	if i != nil {
		arc.SetUpdatedBy(*i)
	}
	return arc
}

// SetUpdatedAt sets the "updated_at" field.
func (arc *AppRoleCreate) SetUpdatedAt(t time.Time) *AppRoleCreate {
	arc.mutation.SetUpdatedAt(t)
	return arc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableUpdatedAt(t *time.Time) *AppRoleCreate {
	if t != nil {
		arc.SetUpdatedAt(*t)
	}
	return arc
}

// SetAppID sets the "app_id" field.
func (arc *AppRoleCreate) SetAppID(i int) *AppRoleCreate {
	arc.mutation.SetAppID(i)
	return arc
}

// SetName sets the "name" field.
func (arc *AppRoleCreate) SetName(s string) *AppRoleCreate {
	arc.mutation.SetName(s)
	return arc
}

// SetComments sets the "comments" field.
func (arc *AppRoleCreate) SetComments(s string) *AppRoleCreate {
	arc.mutation.SetComments(s)
	return arc
}

// SetNillableComments sets the "comments" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableComments(s *string) *AppRoleCreate {
	if s != nil {
		arc.SetComments(*s)
	}
	return arc
}

// SetAutoGrant sets the "auto_grant" field.
func (arc *AppRoleCreate) SetAutoGrant(b bool) *AppRoleCreate {
	arc.mutation.SetAutoGrant(b)
	return arc
}

// SetNillableAutoGrant sets the "auto_grant" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableAutoGrant(b *bool) *AppRoleCreate {
	if b != nil {
		arc.SetAutoGrant(*b)
	}
	return arc
}

// SetEditable sets the "editable" field.
func (arc *AppRoleCreate) SetEditable(b bool) *AppRoleCreate {
	arc.mutation.SetEditable(b)
	return arc
}

// SetNillableEditable sets the "editable" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableEditable(b *bool) *AppRoleCreate {
	if b != nil {
		arc.SetEditable(*b)
	}
	return arc
}

// SetID sets the "id" field.
func (arc *AppRoleCreate) SetID(i int) *AppRoleCreate {
	arc.mutation.SetID(i)
	return arc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (arc *AppRoleCreate) SetNillableID(i *int) *AppRoleCreate {
	if i != nil {
		arc.SetID(*i)
	}
	return arc
}

// SetApp sets the "app" edge to the App entity.
func (arc *AppRoleCreate) SetApp(a *App) *AppRoleCreate {
	return arc.SetAppID(a.ID)
}

// AddPolicyIDs adds the "policies" edge to the AppPolicy entity by IDs.
func (arc *AppRoleCreate) AddPolicyIDs(ids ...int) *AppRoleCreate {
	arc.mutation.AddPolicyIDs(ids...)
	return arc
}

// AddPolicies adds the "policies" edges to the AppPolicy entity.
func (arc *AppRoleCreate) AddPolicies(a ...*AppPolicy) *AppRoleCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return arc.AddPolicyIDs(ids...)
}

// Mutation returns the AppRoleMutation object of the builder.
func (arc *AppRoleCreate) Mutation() *AppRoleMutation {
	return arc.mutation
}

// Save creates the AppRole in the database.
func (arc *AppRoleCreate) Save(ctx context.Context) (*AppRole, error) {
	if err := arc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*AppRole, AppRoleMutation](ctx, arc.sqlSave, arc.mutation, arc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (arc *AppRoleCreate) SaveX(ctx context.Context) *AppRole {
	v, err := arc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arc *AppRoleCreate) Exec(ctx context.Context) error {
	_, err := arc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arc *AppRoleCreate) ExecX(ctx context.Context) {
	if err := arc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (arc *AppRoleCreate) defaults() error {
	if _, ok := arc.mutation.CreatedAt(); !ok {
		if approle.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized approle.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := approle.DefaultCreatedAt()
		arc.mutation.SetCreatedAt(v)
	}
	if _, ok := arc.mutation.AutoGrant(); !ok {
		v := approle.DefaultAutoGrant
		arc.mutation.SetAutoGrant(v)
	}
	if _, ok := arc.mutation.Editable(); !ok {
		v := approle.DefaultEditable
		arc.mutation.SetEditable(v)
	}
	if _, ok := arc.mutation.ID(); !ok {
		if approle.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized approle.DefaultID (forgotten import ent/runtime?)")
		}
		v := approle.DefaultID()
		arc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (arc *AppRoleCreate) check() error {
	if _, ok := arc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "AppRole.created_by"`)}
	}
	if _, ok := arc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppRole.created_at"`)}
	}
	if _, ok := arc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppRole.app_id"`)}
	}
	if _, ok := arc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AppRole.name"`)}
	}
	if _, ok := arc.mutation.AutoGrant(); !ok {
		return &ValidationError{Name: "auto_grant", err: errors.New(`ent: missing required field "AppRole.auto_grant"`)}
	}
	if _, ok := arc.mutation.Editable(); !ok {
		return &ValidationError{Name: "editable", err: errors.New(`ent: missing required field "AppRole.editable"`)}
	}
	if _, ok := arc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app", err: errors.New(`ent: missing required edge "AppRole.app"`)}
	}
	return nil
}

func (arc *AppRoleCreate) sqlSave(ctx context.Context) (*AppRole, error) {
	if err := arc.check(); err != nil {
		return nil, err
	}
	_node, _spec := arc.createSpec()
	if err := sqlgraph.CreateNode(ctx, arc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	arc.mutation.id = &_node.ID
	arc.mutation.done = true
	return _node, nil
}

func (arc *AppRoleCreate) createSpec() (*AppRole, *sqlgraph.CreateSpec) {
	var (
		_node = &AppRole{config: arc.config}
		_spec = sqlgraph.NewCreateSpec(approle.Table, sqlgraph.NewFieldSpec(approle.FieldID, field.TypeInt))
	)
	if id, ok := arc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := arc.mutation.CreatedBy(); ok {
		_spec.SetField(approle.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := arc.mutation.CreatedAt(); ok {
		_spec.SetField(approle.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := arc.mutation.UpdatedBy(); ok {
		_spec.SetField(approle.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := arc.mutation.UpdatedAt(); ok {
		_spec.SetField(approle.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := arc.mutation.Name(); ok {
		_spec.SetField(approle.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := arc.mutation.Comments(); ok {
		_spec.SetField(approle.FieldComments, field.TypeString, value)
		_node.Comments = value
	}
	if value, ok := arc.mutation.AutoGrant(); ok {
		_spec.SetField(approle.FieldAutoGrant, field.TypeBool, value)
		_node.AutoGrant = value
	}
	if value, ok := arc.mutation.Editable(); ok {
		_spec.SetField(approle.FieldEditable, field.TypeBool, value)
		_node.Editable = value
	}
	if nodes := arc.mutation.AppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   approle.AppTable,
			Columns: []string{approle.AppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: app.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AppID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := arc.mutation.PoliciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   approle.PoliciesTable,
			Columns: approle.PoliciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: apppolicy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AppRolePolicyCreate{config: arc.config, mutation: newAppRolePolicyMutation(arc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AppRoleCreateBulk is the builder for creating many AppRole entities in bulk.
type AppRoleCreateBulk struct {
	config
	builders []*AppRoleCreate
}

// Save creates the AppRole entities in the database.
func (arcb *AppRoleCreateBulk) Save(ctx context.Context) ([]*AppRole, error) {
	specs := make([]*sqlgraph.CreateSpec, len(arcb.builders))
	nodes := make([]*AppRole, len(arcb.builders))
	mutators := make([]Mutator, len(arcb.builders))
	for i := range arcb.builders {
		func(i int, root context.Context) {
			builder := arcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppRoleMutation)
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
					_, err = mutators[i+1].Mutate(root, arcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, arcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, arcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (arcb *AppRoleCreateBulk) SaveX(ctx context.Context) []*AppRole {
	v, err := arcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (arcb *AppRoleCreateBulk) Exec(ctx context.Context) error {
	_, err := arcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arcb *AppRoleCreateBulk) ExecX(ctx context.Context) {
	if err := arcb.Exec(ctx); err != nil {
		panic(err)
	}
}
