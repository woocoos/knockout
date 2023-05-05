// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/orgroleuser"
	"github.com/woocoos/knockout/ent/predicate"
)

// OrgRoleUserDelete is the builder for deleting a OrgRoleUser entity.
type OrgRoleUserDelete struct {
	config
	hooks    []Hook
	mutation *OrgRoleUserMutation
}

// Where appends a list predicates to the OrgRoleUserDelete builder.
func (orud *OrgRoleUserDelete) Where(ps ...predicate.OrgRoleUser) *OrgRoleUserDelete {
	orud.mutation.Where(ps...)
	return orud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (orud *OrgRoleUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, orud.sqlExec, orud.mutation, orud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (orud *OrgRoleUserDelete) ExecX(ctx context.Context) int {
	n, err := orud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (orud *OrgRoleUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orgroleuser.Table, sqlgraph.NewFieldSpec(orgroleuser.FieldID, field.TypeInt))
	if ps := orud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, orud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	orud.mutation.done = true
	return affected, err
}

// OrgRoleUserDeleteOne is the builder for deleting a single OrgRoleUser entity.
type OrgRoleUserDeleteOne struct {
	orud *OrgRoleUserDelete
}

// Where appends a list predicates to the OrgRoleUserDelete builder.
func (orudo *OrgRoleUserDeleteOne) Where(ps ...predicate.OrgRoleUser) *OrgRoleUserDeleteOne {
	orudo.orud.mutation.Where(ps...)
	return orudo
}

// Exec executes the deletion query.
func (orudo *OrgRoleUserDeleteOne) Exec(ctx context.Context) error {
	n, err := orudo.orud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orgroleuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (orudo *OrgRoleUserDeleteOne) ExecX(ctx context.Context) {
	if err := orudo.Exec(ctx); err != nil {
		panic(err)
	}
}
