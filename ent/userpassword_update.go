// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/ent/predicate"
	"github.com/woocoos/knockout/ent/userpassword"
)

// UserPasswordUpdate is the builder for updating UserPassword entities.
type UserPasswordUpdate struct {
	config
	hooks    []Hook
	mutation *UserPasswordMutation
}

// Where appends a list predicates to the UserPasswordUpdate builder.
func (upu *UserPasswordUpdate) Where(ps ...predicate.UserPassword) *UserPasswordUpdate {
	upu.mutation.Where(ps...)
	return upu
}

// SetUpdatedBy sets the "updated_by" field.
func (upu *UserPasswordUpdate) SetUpdatedBy(i int) *UserPasswordUpdate {
	upu.mutation.ResetUpdatedBy()
	upu.mutation.SetUpdatedBy(i)
	return upu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (upu *UserPasswordUpdate) SetNillableUpdatedBy(i *int) *UserPasswordUpdate {
	if i != nil {
		upu.SetUpdatedBy(*i)
	}
	return upu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (upu *UserPasswordUpdate) AddUpdatedBy(i int) *UserPasswordUpdate {
	upu.mutation.AddUpdatedBy(i)
	return upu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (upu *UserPasswordUpdate) ClearUpdatedBy() *UserPasswordUpdate {
	upu.mutation.ClearUpdatedBy()
	return upu
}

// SetUpdatedAt sets the "updated_at" field.
func (upu *UserPasswordUpdate) SetUpdatedAt(t time.Time) *UserPasswordUpdate {
	upu.mutation.SetUpdatedAt(t)
	return upu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (upu *UserPasswordUpdate) SetNillableUpdatedAt(t *time.Time) *UserPasswordUpdate {
	if t != nil {
		upu.SetUpdatedAt(*t)
	}
	return upu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (upu *UserPasswordUpdate) ClearUpdatedAt() *UserPasswordUpdate {
	upu.mutation.ClearUpdatedAt()
	return upu
}

// SetScene sets the "scene" field.
func (upu *UserPasswordUpdate) SetScene(u userpassword.Scene) *UserPasswordUpdate {
	upu.mutation.SetScene(u)
	return upu
}

// SetNillableScene sets the "scene" field if the given value is not nil.
func (upu *UserPasswordUpdate) SetNillableScene(u *userpassword.Scene) *UserPasswordUpdate {
	if u != nil {
		upu.SetScene(*u)
	}
	return upu
}

// SetPassword sets the "password" field.
func (upu *UserPasswordUpdate) SetPassword(s string) *UserPasswordUpdate {
	upu.mutation.SetPassword(s)
	return upu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (upu *UserPasswordUpdate) SetNillablePassword(s *string) *UserPasswordUpdate {
	if s != nil {
		upu.SetPassword(*s)
	}
	return upu
}

// ClearPassword clears the value of the "password" field.
func (upu *UserPasswordUpdate) ClearPassword() *UserPasswordUpdate {
	upu.mutation.ClearPassword()
	return upu
}

// SetSalt sets the "salt" field.
func (upu *UserPasswordUpdate) SetSalt(s string) *UserPasswordUpdate {
	upu.mutation.SetSalt(s)
	return upu
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (upu *UserPasswordUpdate) SetNillableSalt(s *string) *UserPasswordUpdate {
	if s != nil {
		upu.SetSalt(*s)
	}
	return upu
}

// SetStatus sets the "status" field.
func (upu *UserPasswordUpdate) SetStatus(ts typex.SimpleStatus) *UserPasswordUpdate {
	upu.mutation.SetStatus(ts)
	return upu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (upu *UserPasswordUpdate) SetNillableStatus(ts *typex.SimpleStatus) *UserPasswordUpdate {
	if ts != nil {
		upu.SetStatus(*ts)
	}
	return upu
}

// ClearStatus clears the value of the "status" field.
func (upu *UserPasswordUpdate) ClearStatus() *UserPasswordUpdate {
	upu.mutation.ClearStatus()
	return upu
}

// Mutation returns the UserPasswordMutation object of the builder.
func (upu *UserPasswordUpdate) Mutation() *UserPasswordMutation {
	return upu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (upu *UserPasswordUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, upu.sqlSave, upu.mutation, upu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upu *UserPasswordUpdate) SaveX(ctx context.Context) int {
	affected, err := upu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (upu *UserPasswordUpdate) Exec(ctx context.Context) error {
	_, err := upu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upu *UserPasswordUpdate) ExecX(ctx context.Context) {
	if err := upu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upu *UserPasswordUpdate) check() error {
	if v, ok := upu.mutation.Scene(); ok {
		if err := userpassword.SceneValidator(v); err != nil {
			return &ValidationError{Name: "scene", err: fmt.Errorf(`ent: validator failed for field "UserPassword.scene": %w`, err)}
		}
	}
	if v, ok := upu.mutation.Salt(); ok {
		if err := userpassword.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "UserPassword.salt": %w`, err)}
		}
	}
	if v, ok := upu.mutation.Status(); ok {
		if err := userpassword.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserPassword.status": %w`, err)}
		}
	}
	return nil
}

func (upu *UserPasswordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := upu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userpassword.Table, userpassword.Columns, sqlgraph.NewFieldSpec(userpassword.FieldID, field.TypeInt))
	if ps := upu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upu.mutation.UpdatedBy(); ok {
		_spec.SetField(userpassword.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := upu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(userpassword.FieldUpdatedBy, field.TypeInt, value)
	}
	if upu.mutation.UpdatedByCleared() {
		_spec.ClearField(userpassword.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := upu.mutation.UpdatedAt(); ok {
		_spec.SetField(userpassword.FieldUpdatedAt, field.TypeTime, value)
	}
	if upu.mutation.UpdatedAtCleared() {
		_spec.ClearField(userpassword.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := upu.mutation.Scene(); ok {
		_spec.SetField(userpassword.FieldScene, field.TypeEnum, value)
	}
	if value, ok := upu.mutation.Password(); ok {
		_spec.SetField(userpassword.FieldPassword, field.TypeString, value)
	}
	if upu.mutation.PasswordCleared() {
		_spec.ClearField(userpassword.FieldPassword, field.TypeString)
	}
	if value, ok := upu.mutation.Salt(); ok {
		_spec.SetField(userpassword.FieldSalt, field.TypeString, value)
	}
	if value, ok := upu.mutation.Status(); ok {
		_spec.SetField(userpassword.FieldStatus, field.TypeEnum, value)
	}
	if upu.mutation.StatusCleared() {
		_spec.ClearField(userpassword.FieldStatus, field.TypeEnum)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, upu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userpassword.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	upu.mutation.done = true
	return n, nil
}

// UserPasswordUpdateOne is the builder for updating a single UserPassword entity.
type UserPasswordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserPasswordMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (upuo *UserPasswordUpdateOne) SetUpdatedBy(i int) *UserPasswordUpdateOne {
	upuo.mutation.ResetUpdatedBy()
	upuo.mutation.SetUpdatedBy(i)
	return upuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (upuo *UserPasswordUpdateOne) SetNillableUpdatedBy(i *int) *UserPasswordUpdateOne {
	if i != nil {
		upuo.SetUpdatedBy(*i)
	}
	return upuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (upuo *UserPasswordUpdateOne) AddUpdatedBy(i int) *UserPasswordUpdateOne {
	upuo.mutation.AddUpdatedBy(i)
	return upuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (upuo *UserPasswordUpdateOne) ClearUpdatedBy() *UserPasswordUpdateOne {
	upuo.mutation.ClearUpdatedBy()
	return upuo
}

// SetUpdatedAt sets the "updated_at" field.
func (upuo *UserPasswordUpdateOne) SetUpdatedAt(t time.Time) *UserPasswordUpdateOne {
	upuo.mutation.SetUpdatedAt(t)
	return upuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (upuo *UserPasswordUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserPasswordUpdateOne {
	if t != nil {
		upuo.SetUpdatedAt(*t)
	}
	return upuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (upuo *UserPasswordUpdateOne) ClearUpdatedAt() *UserPasswordUpdateOne {
	upuo.mutation.ClearUpdatedAt()
	return upuo
}

// SetScene sets the "scene" field.
func (upuo *UserPasswordUpdateOne) SetScene(u userpassword.Scene) *UserPasswordUpdateOne {
	upuo.mutation.SetScene(u)
	return upuo
}

// SetNillableScene sets the "scene" field if the given value is not nil.
func (upuo *UserPasswordUpdateOne) SetNillableScene(u *userpassword.Scene) *UserPasswordUpdateOne {
	if u != nil {
		upuo.SetScene(*u)
	}
	return upuo
}

// SetPassword sets the "password" field.
func (upuo *UserPasswordUpdateOne) SetPassword(s string) *UserPasswordUpdateOne {
	upuo.mutation.SetPassword(s)
	return upuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (upuo *UserPasswordUpdateOne) SetNillablePassword(s *string) *UserPasswordUpdateOne {
	if s != nil {
		upuo.SetPassword(*s)
	}
	return upuo
}

// ClearPassword clears the value of the "password" field.
func (upuo *UserPasswordUpdateOne) ClearPassword() *UserPasswordUpdateOne {
	upuo.mutation.ClearPassword()
	return upuo
}

// SetSalt sets the "salt" field.
func (upuo *UserPasswordUpdateOne) SetSalt(s string) *UserPasswordUpdateOne {
	upuo.mutation.SetSalt(s)
	return upuo
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (upuo *UserPasswordUpdateOne) SetNillableSalt(s *string) *UserPasswordUpdateOne {
	if s != nil {
		upuo.SetSalt(*s)
	}
	return upuo
}

// SetStatus sets the "status" field.
func (upuo *UserPasswordUpdateOne) SetStatus(ts typex.SimpleStatus) *UserPasswordUpdateOne {
	upuo.mutation.SetStatus(ts)
	return upuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (upuo *UserPasswordUpdateOne) SetNillableStatus(ts *typex.SimpleStatus) *UserPasswordUpdateOne {
	if ts != nil {
		upuo.SetStatus(*ts)
	}
	return upuo
}

// ClearStatus clears the value of the "status" field.
func (upuo *UserPasswordUpdateOne) ClearStatus() *UserPasswordUpdateOne {
	upuo.mutation.ClearStatus()
	return upuo
}

// Mutation returns the UserPasswordMutation object of the builder.
func (upuo *UserPasswordUpdateOne) Mutation() *UserPasswordMutation {
	return upuo.mutation
}

// Where appends a list predicates to the UserPasswordUpdate builder.
func (upuo *UserPasswordUpdateOne) Where(ps ...predicate.UserPassword) *UserPasswordUpdateOne {
	upuo.mutation.Where(ps...)
	return upuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (upuo *UserPasswordUpdateOne) Select(field string, fields ...string) *UserPasswordUpdateOne {
	upuo.fields = append([]string{field}, fields...)
	return upuo
}

// Save executes the query and returns the updated UserPassword entity.
func (upuo *UserPasswordUpdateOne) Save(ctx context.Context) (*UserPassword, error) {
	return withHooks(ctx, upuo.sqlSave, upuo.mutation, upuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (upuo *UserPasswordUpdateOne) SaveX(ctx context.Context) *UserPassword {
	node, err := upuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (upuo *UserPasswordUpdateOne) Exec(ctx context.Context) error {
	_, err := upuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upuo *UserPasswordUpdateOne) ExecX(ctx context.Context) {
	if err := upuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upuo *UserPasswordUpdateOne) check() error {
	if v, ok := upuo.mutation.Scene(); ok {
		if err := userpassword.SceneValidator(v); err != nil {
			return &ValidationError{Name: "scene", err: fmt.Errorf(`ent: validator failed for field "UserPassword.scene": %w`, err)}
		}
	}
	if v, ok := upuo.mutation.Salt(); ok {
		if err := userpassword.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "UserPassword.salt": %w`, err)}
		}
	}
	if v, ok := upuo.mutation.Status(); ok {
		if err := userpassword.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserPassword.status": %w`, err)}
		}
	}
	return nil
}

func (upuo *UserPasswordUpdateOne) sqlSave(ctx context.Context) (_node *UserPassword, err error) {
	if err := upuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userpassword.Table, userpassword.Columns, sqlgraph.NewFieldSpec(userpassword.FieldID, field.TypeInt))
	id, ok := upuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserPassword.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := upuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userpassword.FieldID)
		for _, f := range fields {
			if !userpassword.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userpassword.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := upuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := upuo.mutation.UpdatedBy(); ok {
		_spec.SetField(userpassword.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := upuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(userpassword.FieldUpdatedBy, field.TypeInt, value)
	}
	if upuo.mutation.UpdatedByCleared() {
		_spec.ClearField(userpassword.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := upuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userpassword.FieldUpdatedAt, field.TypeTime, value)
	}
	if upuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(userpassword.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := upuo.mutation.Scene(); ok {
		_spec.SetField(userpassword.FieldScene, field.TypeEnum, value)
	}
	if value, ok := upuo.mutation.Password(); ok {
		_spec.SetField(userpassword.FieldPassword, field.TypeString, value)
	}
	if upuo.mutation.PasswordCleared() {
		_spec.ClearField(userpassword.FieldPassword, field.TypeString)
	}
	if value, ok := upuo.mutation.Salt(); ok {
		_spec.SetField(userpassword.FieldSalt, field.TypeString, value)
	}
	if value, ok := upuo.mutation.Status(); ok {
		_spec.SetField(userpassword.FieldStatus, field.TypeEnum, value)
	}
	if upuo.mutation.StatusCleared() {
		_spec.ClearField(userpassword.FieldStatus, field.TypeEnum)
	}
	_node = &UserPassword{config: upuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, upuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userpassword.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	upuo.mutation.done = true
	return _node, nil
}
