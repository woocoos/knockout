// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/approle"
	"github.com/woocoos/knockout/ent/predicate"
)

// AppRoleDelete is the builder for deleting a AppRole entity.
type AppRoleDelete struct {
	config
	hooks    []Hook
	mutation *AppRoleMutation
}

// Where appends a list predicates to the AppRoleDelete builder.
func (ard *AppRoleDelete) Where(ps ...predicate.AppRole) *AppRoleDelete {
	ard.mutation.Where(ps...)
	return ard
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ard *AppRoleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ard.sqlExec, ard.mutation, ard.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ard *AppRoleDelete) ExecX(ctx context.Context) int {
	n, err := ard.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ard *AppRoleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(approle.Table, sqlgraph.NewFieldSpec(approle.FieldID, field.TypeInt))
	if ps := ard.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ard.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ard.mutation.done = true
	return affected, err
}

// AppRoleDeleteOne is the builder for deleting a single AppRole entity.
type AppRoleDeleteOne struct {
	ard *AppRoleDelete
}

// Where appends a list predicates to the AppRoleDelete builder.
func (ardo *AppRoleDeleteOne) Where(ps ...predicate.AppRole) *AppRoleDeleteOne {
	ardo.ard.mutation.Where(ps...)
	return ardo
}

// Exec executes the deletion query.
func (ardo *AppRoleDeleteOne) Exec(ctx context.Context) error {
	n, err := ardo.ard.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{approle.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ardo *AppRoleDeleteOne) ExecX(ctx context.Context) {
	if err := ardo.Exec(ctx); err != nil {
		panic(err)
	}
}
