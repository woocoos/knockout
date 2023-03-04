// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/predicate"
)

// OrgPolicyDelete is the builder for deleting a OrgPolicy entity.
type OrgPolicyDelete struct {
	config
	hooks    []Hook
	mutation *OrgPolicyMutation
}

// Where appends a list predicates to the OrgPolicyDelete builder.
func (opd *OrgPolicyDelete) Where(ps ...predicate.OrgPolicy) *OrgPolicyDelete {
	opd.mutation.Where(ps...)
	return opd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (opd *OrgPolicyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, OrgPolicyMutation](ctx, opd.sqlExec, opd.mutation, opd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (opd *OrgPolicyDelete) ExecX(ctx context.Context) int {
	n, err := opd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (opd *OrgPolicyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orgpolicy.Table, sqlgraph.NewFieldSpec(orgpolicy.FieldID, field.TypeInt))
	if ps := opd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, opd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	opd.mutation.done = true
	return affected, err
}

// OrgPolicyDeleteOne is the builder for deleting a single OrgPolicy entity.
type OrgPolicyDeleteOne struct {
	opd *OrgPolicyDelete
}

// Where appends a list predicates to the OrgPolicyDelete builder.
func (opdo *OrgPolicyDeleteOne) Where(ps ...predicate.OrgPolicy) *OrgPolicyDeleteOne {
	opdo.opd.mutation.Where(ps...)
	return opdo
}

// Exec executes the deletion query.
func (opdo *OrgPolicyDeleteOne) Exec(ctx context.Context) error {
	n, err := opdo.opd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orgpolicy.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (opdo *OrgPolicyDeleteOne) ExecX(ctx context.Context) {
	if err := opdo.Exec(ctx); err != nil {
		panic(err)
	}
}
