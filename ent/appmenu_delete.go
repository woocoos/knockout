// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/predicate"
)

// AppMenuDelete is the builder for deleting a AppMenu entity.
type AppMenuDelete struct {
	config
	hooks    []Hook
	mutation *AppMenuMutation
}

// Where appends a list predicates to the AppMenuDelete builder.
func (amd *AppMenuDelete) Where(ps ...predicate.AppMenu) *AppMenuDelete {
	amd.mutation.Where(ps...)
	return amd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (amd *AppMenuDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, amd.sqlExec, amd.mutation, amd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (amd *AppMenuDelete) ExecX(ctx context.Context) int {
	n, err := amd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (amd *AppMenuDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(appmenu.Table, sqlgraph.NewFieldSpec(appmenu.FieldID, field.TypeInt))
	if ps := amd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, amd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	amd.mutation.done = true
	return affected, err
}

// AppMenuDeleteOne is the builder for deleting a single AppMenu entity.
type AppMenuDeleteOne struct {
	amd *AppMenuDelete
}

// Where appends a list predicates to the AppMenuDelete builder.
func (amdo *AppMenuDeleteOne) Where(ps ...predicate.AppMenu) *AppMenuDeleteOne {
	amdo.amd.mutation.Where(ps...)
	return amdo
}

// Exec executes the deletion query.
func (amdo *AppMenuDeleteOne) Exec(ctx context.Context) error {
	n, err := amdo.amd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appmenu.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (amdo *AppMenuDeleteOne) ExecX(ctx context.Context) {
	if err := amdo.Exec(ctx); err != nil {
		panic(err)
	}
}
