// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/predicate"
)

// OrgUserDelete is the builder for deleting a OrgUser entity.
type OrgUserDelete struct {
	config
	hooks    []Hook
	mutation *OrgUserMutation
}

// Where appends a list predicates to the OrgUserDelete builder.
func (oud *OrgUserDelete) Where(ps ...predicate.OrgUser) *OrgUserDelete {
	oud.mutation.Where(ps...)
	return oud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oud *OrgUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, oud.sqlExec, oud.mutation, oud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (oud *OrgUserDelete) ExecX(ctx context.Context) int {
	n, err := oud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oud *OrgUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orguser.Table, sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt))
	if ps := oud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, oud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	oud.mutation.done = true
	return affected, err
}

// OrgUserDeleteOne is the builder for deleting a single OrgUser entity.
type OrgUserDeleteOne struct {
	oud *OrgUserDelete
}

// Where appends a list predicates to the OrgUserDelete builder.
func (oudo *OrgUserDeleteOne) Where(ps ...predicate.OrgUser) *OrgUserDeleteOne {
	oudo.oud.mutation.Where(ps...)
	return oudo
}

// Exec executes the deletion query.
func (oudo *OrgUserDeleteOne) Exec(ctx context.Context) error {
	n, err := oudo.oud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orguser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oudo *OrgUserDeleteOne) ExecX(ctx context.Context) {
	if err := oudo.Exec(ctx); err != nil {
		panic(err)
	}
}
