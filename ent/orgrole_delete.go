// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/predicate"
)

// OrgRoleDelete is the builder for deleting a OrgRole entity.
type OrgRoleDelete struct {
	config
	hooks    []Hook
	mutation *OrgRoleMutation
}

// Where appends a list predicates to the OrgRoleDelete builder.
func (ord *OrgRoleDelete) Where(ps ...predicate.OrgRole) *OrgRoleDelete {
	ord.mutation.Where(ps...)
	return ord
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ord *OrgRoleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, OrgRoleMutation](ctx, ord.sqlExec, ord.mutation, ord.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ord *OrgRoleDelete) ExecX(ctx context.Context) int {
	n, err := ord.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ord *OrgRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orgrole.Table, sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt))
	if ps := ord.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ord.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ord.mutation.done = true
	return affected, err
}

// OrgRoleDeleteOne is the builder for deleting a single OrgRole entity.
type OrgRoleDeleteOne struct {
	ord *OrgRoleDelete
}

// Where appends a list predicates to the OrgRoleDelete builder.
func (ordo *OrgRoleDeleteOne) Where(ps ...predicate.OrgRole) *OrgRoleDeleteOne {
	ordo.ord.mutation.Where(ps...)
	return ordo
}

// Exec executes the deletion query.
func (ordo *OrgRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := ordo.ord.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orgrole.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ordo *OrgRoleDeleteOne) ExecX(ctx context.Context) {
	if err := ordo.Exec(ctx); err != nil {
		panic(err)
	}
}
