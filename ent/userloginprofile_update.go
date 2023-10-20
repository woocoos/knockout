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
	"github.com/woocoos/knockout/ent/userloginprofile"
)

// UserLoginProfileUpdate is the builder for updating UserLoginProfile entities.
type UserLoginProfileUpdate struct {
	config
	hooks    []Hook
	mutation *UserLoginProfileMutation
}

// Where appends a list predicates to the UserLoginProfileUpdate builder.
func (ulpu *UserLoginProfileUpdate) Where(ps ...predicate.UserLoginProfile) *UserLoginProfileUpdate {
	ulpu.mutation.Where(ps...)
	return ulpu
}

// SetUpdatedBy sets the "updated_by" field.
func (ulpu *UserLoginProfileUpdate) SetUpdatedBy(i int) *UserLoginProfileUpdate {
	ulpu.mutation.ResetUpdatedBy()
	ulpu.mutation.SetUpdatedBy(i)
	return ulpu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ulpu *UserLoginProfileUpdate) SetNillableUpdatedBy(i *int) *UserLoginProfileUpdate {
	if i != nil {
		ulpu.SetUpdatedBy(*i)
	}
	return ulpu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ulpu *UserLoginProfileUpdate) AddUpdatedBy(i int) *UserLoginProfileUpdate {
	ulpu.mutation.AddUpdatedBy(i)
	return ulpu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ulpu *UserLoginProfileUpdate) ClearUpdatedBy() *UserLoginProfileUpdate {
	ulpu.mutation.ClearUpdatedBy()
	return ulpu
}

// SetUpdatedAt sets the "updated_at" field.
func (ulpu *UserLoginProfileUpdate) SetUpdatedAt(t time.Time) *UserLoginProfileUpdate {
	ulpu.mutation.SetUpdatedAt(t)
	return ulpu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ulpu *UserLoginProfileUpdate) SetNillableUpdatedAt(t *time.Time) *UserLoginProfileUpdate {
	if t != nil {
		ulpu.SetUpdatedAt(*t)
	}
	return ulpu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ulpu *UserLoginProfileUpdate) ClearUpdatedAt() *UserLoginProfileUpdate {
	ulpu.mutation.ClearUpdatedAt()
	return ulpu
}

// SetLastLoginIP sets the "last_login_ip" field.
func (ulpu *UserLoginProfileUpdate) SetLastLoginIP(s string) *UserLoginProfileUpdate {
	ulpu.mutation.SetLastLoginIP(s)
	return ulpu
}

// SetNillableLastLoginIP sets the "last_login_ip" field if the given value is not nil.
func (ulpu *UserLoginProfileUpdate) SetNillableLastLoginIP(s *string) *UserLoginProfileUpdate {
	if s != nil {
		ulpu.SetLastLoginIP(*s)
	}
	return ulpu
}

// ClearLastLoginIP clears the value of the "last_login_ip" field.
func (ulpu *UserLoginProfileUpdate) ClearLastLoginIP() *UserLoginProfileUpdate {
	ulpu.mutation.ClearLastLoginIP()
	return ulpu
}

// SetLastLoginAt sets the "last_login_at" field.
func (ulpu *UserLoginProfileUpdate) SetLastLoginAt(t time.Time) *UserLoginProfileUpdate {
	ulpu.mutation.SetLastLoginAt(t)
	return ulpu
}

// SetNillableLastLoginAt sets the "last_login_at" field if the given value is not nil.
func (ulpu *UserLoginProfileUpdate) SetNillableLastLoginAt(t *time.Time) *UserLoginProfileUpdate {
	if t != nil {
		ulpu.SetLastLoginAt(*t)
	}
	return ulpu
}

// ClearLastLoginAt clears the value of the "last_login_at" field.
func (ulpu *UserLoginProfileUpdate) ClearLastLoginAt() *UserLoginProfileUpdate {
	ulpu.mutation.ClearLastLoginAt()
	return ulpu
}

// SetCanLogin sets the "can_login" field.
func (ulpu *UserLoginProfileUpdate) SetCanLogin(b bool) *UserLoginProfileUpdate {
	ulpu.mutation.SetCanLogin(b)
	return ulpu
}

// SetNillableCanLogin sets the "can_login" field if the given value is not nil.
func (ulpu *UserLoginProfileUpdate) SetNillableCanLogin(b *bool) *UserLoginProfileUpdate {
	if b != nil {
		ulpu.SetCanLogin(*b)
	}
	return ulpu
}

// ClearCanLogin clears the value of the "can_login" field.
func (ulpu *UserLoginProfileUpdate) ClearCanLogin() *UserLoginProfileUpdate {
	ulpu.mutation.ClearCanLogin()
	return ulpu
}

// SetSetKind sets the "set_kind" field.
func (ulpu *UserLoginProfileUpdate) SetSetKind(uk userloginprofile.SetKind) *UserLoginProfileUpdate {
	ulpu.mutation.SetSetKind(uk)
	return ulpu
}

// SetPasswordReset sets the "password_reset" field.
func (ulpu *UserLoginProfileUpdate) SetPasswordReset(b bool) *UserLoginProfileUpdate {
	ulpu.mutation.SetPasswordReset(b)
	return ulpu
}

// SetNillablePasswordReset sets the "password_reset" field if the given value is not nil.
func (ulpu *UserLoginProfileUpdate) SetNillablePasswordReset(b *bool) *UserLoginProfileUpdate {
	if b != nil {
		ulpu.SetPasswordReset(*b)
	}
	return ulpu
}

// ClearPasswordReset clears the value of the "password_reset" field.
func (ulpu *UserLoginProfileUpdate) ClearPasswordReset() *UserLoginProfileUpdate {
	ulpu.mutation.ClearPasswordReset()
	return ulpu
}

// SetVerifyDevice sets the "verify_device" field.
func (ulpu *UserLoginProfileUpdate) SetVerifyDevice(b bool) *UserLoginProfileUpdate {
	ulpu.mutation.SetVerifyDevice(b)
	return ulpu
}

// SetMfaEnabled sets the "mfa_enabled" field.
func (ulpu *UserLoginProfileUpdate) SetMfaEnabled(b bool) *UserLoginProfileUpdate {
	ulpu.mutation.SetMfaEnabled(b)
	return ulpu
}

// SetNillableMfaEnabled sets the "mfa_enabled" field if the given value is not nil.
func (ulpu *UserLoginProfileUpdate) SetNillableMfaEnabled(b *bool) *UserLoginProfileUpdate {
	if b != nil {
		ulpu.SetMfaEnabled(*b)
	}
	return ulpu
}

// ClearMfaEnabled clears the value of the "mfa_enabled" field.
func (ulpu *UserLoginProfileUpdate) ClearMfaEnabled() *UserLoginProfileUpdate {
	ulpu.mutation.ClearMfaEnabled()
	return ulpu
}

// SetMfaSecret sets the "mfa_secret" field.
func (ulpu *UserLoginProfileUpdate) SetMfaSecret(s string) *UserLoginProfileUpdate {
	ulpu.mutation.SetMfaSecret(s)
	return ulpu
}

// SetNillableMfaSecret sets the "mfa_secret" field if the given value is not nil.
func (ulpu *UserLoginProfileUpdate) SetNillableMfaSecret(s *string) *UserLoginProfileUpdate {
	if s != nil {
		ulpu.SetMfaSecret(*s)
	}
	return ulpu
}

// ClearMfaSecret clears the value of the "mfa_secret" field.
func (ulpu *UserLoginProfileUpdate) ClearMfaSecret() *UserLoginProfileUpdate {
	ulpu.mutation.ClearMfaSecret()
	return ulpu
}

// SetMfaStatus sets the "mfa_status" field.
func (ulpu *UserLoginProfileUpdate) SetMfaStatus(ts typex.SimpleStatus) *UserLoginProfileUpdate {
	ulpu.mutation.SetMfaStatus(ts)
	return ulpu
}

// SetNillableMfaStatus sets the "mfa_status" field if the given value is not nil.
func (ulpu *UserLoginProfileUpdate) SetNillableMfaStatus(ts *typex.SimpleStatus) *UserLoginProfileUpdate {
	if ts != nil {
		ulpu.SetMfaStatus(*ts)
	}
	return ulpu
}

// ClearMfaStatus clears the value of the "mfa_status" field.
func (ulpu *UserLoginProfileUpdate) ClearMfaStatus() *UserLoginProfileUpdate {
	ulpu.mutation.ClearMfaStatus()
	return ulpu
}

// Mutation returns the UserLoginProfileMutation object of the builder.
func (ulpu *UserLoginProfileUpdate) Mutation() *UserLoginProfileMutation {
	return ulpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ulpu *UserLoginProfileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ulpu.sqlSave, ulpu.mutation, ulpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ulpu *UserLoginProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := ulpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ulpu *UserLoginProfileUpdate) Exec(ctx context.Context) error {
	_, err := ulpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulpu *UserLoginProfileUpdate) ExecX(ctx context.Context) {
	if err := ulpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ulpu *UserLoginProfileUpdate) check() error {
	if v, ok := ulpu.mutation.SetKind(); ok {
		if err := userloginprofile.SetKindValidator(v); err != nil {
			return &ValidationError{Name: "set_kind", err: fmt.Errorf(`ent: validator failed for field "UserLoginProfile.set_kind": %w`, err)}
		}
	}
	if v, ok := ulpu.mutation.MfaSecret(); ok {
		if err := userloginprofile.MfaSecretValidator(v); err != nil {
			return &ValidationError{Name: "mfa_secret", err: fmt.Errorf(`ent: validator failed for field "UserLoginProfile.mfa_secret": %w`, err)}
		}
	}
	if v, ok := ulpu.mutation.MfaStatus(); ok {
		if err := userloginprofile.MfaStatusValidator(v); err != nil {
			return &ValidationError{Name: "mfa_status", err: fmt.Errorf(`ent: validator failed for field "UserLoginProfile.mfa_status": %w`, err)}
		}
	}
	return nil
}

func (ulpu *UserLoginProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ulpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userloginprofile.Table, userloginprofile.Columns, sqlgraph.NewFieldSpec(userloginprofile.FieldID, field.TypeInt))
	if ps := ulpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ulpu.mutation.UpdatedBy(); ok {
		_spec.SetField(userloginprofile.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ulpu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(userloginprofile.FieldUpdatedBy, field.TypeInt, value)
	}
	if ulpu.mutation.UpdatedByCleared() {
		_spec.ClearField(userloginprofile.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := ulpu.mutation.UpdatedAt(); ok {
		_spec.SetField(userloginprofile.FieldUpdatedAt, field.TypeTime, value)
	}
	if ulpu.mutation.UpdatedAtCleared() {
		_spec.ClearField(userloginprofile.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ulpu.mutation.LastLoginIP(); ok {
		_spec.SetField(userloginprofile.FieldLastLoginIP, field.TypeString, value)
	}
	if ulpu.mutation.LastLoginIPCleared() {
		_spec.ClearField(userloginprofile.FieldLastLoginIP, field.TypeString)
	}
	if value, ok := ulpu.mutation.LastLoginAt(); ok {
		_spec.SetField(userloginprofile.FieldLastLoginAt, field.TypeTime, value)
	}
	if ulpu.mutation.LastLoginAtCleared() {
		_spec.ClearField(userloginprofile.FieldLastLoginAt, field.TypeTime)
	}
	if value, ok := ulpu.mutation.CanLogin(); ok {
		_spec.SetField(userloginprofile.FieldCanLogin, field.TypeBool, value)
	}
	if ulpu.mutation.CanLoginCleared() {
		_spec.ClearField(userloginprofile.FieldCanLogin, field.TypeBool)
	}
	if value, ok := ulpu.mutation.SetKind(); ok {
		_spec.SetField(userloginprofile.FieldSetKind, field.TypeEnum, value)
	}
	if value, ok := ulpu.mutation.PasswordReset(); ok {
		_spec.SetField(userloginprofile.FieldPasswordReset, field.TypeBool, value)
	}
	if ulpu.mutation.PasswordResetCleared() {
		_spec.ClearField(userloginprofile.FieldPasswordReset, field.TypeBool)
	}
	if value, ok := ulpu.mutation.VerifyDevice(); ok {
		_spec.SetField(userloginprofile.FieldVerifyDevice, field.TypeBool, value)
	}
	if value, ok := ulpu.mutation.MfaEnabled(); ok {
		_spec.SetField(userloginprofile.FieldMfaEnabled, field.TypeBool, value)
	}
	if ulpu.mutation.MfaEnabledCleared() {
		_spec.ClearField(userloginprofile.FieldMfaEnabled, field.TypeBool)
	}
	if value, ok := ulpu.mutation.MfaSecret(); ok {
		_spec.SetField(userloginprofile.FieldMfaSecret, field.TypeString, value)
	}
	if ulpu.mutation.MfaSecretCleared() {
		_spec.ClearField(userloginprofile.FieldMfaSecret, field.TypeString)
	}
	if value, ok := ulpu.mutation.MfaStatus(); ok {
		_spec.SetField(userloginprofile.FieldMfaStatus, field.TypeEnum, value)
	}
	if ulpu.mutation.MfaStatusCleared() {
		_spec.ClearField(userloginprofile.FieldMfaStatus, field.TypeEnum)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ulpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userloginprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ulpu.mutation.done = true
	return n, nil
}

// UserLoginProfileUpdateOne is the builder for updating a single UserLoginProfile entity.
type UserLoginProfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserLoginProfileMutation
}

// SetUpdatedBy sets the "updated_by" field.
func (ulpuo *UserLoginProfileUpdateOne) SetUpdatedBy(i int) *UserLoginProfileUpdateOne {
	ulpuo.mutation.ResetUpdatedBy()
	ulpuo.mutation.SetUpdatedBy(i)
	return ulpuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ulpuo *UserLoginProfileUpdateOne) SetNillableUpdatedBy(i *int) *UserLoginProfileUpdateOne {
	if i != nil {
		ulpuo.SetUpdatedBy(*i)
	}
	return ulpuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ulpuo *UserLoginProfileUpdateOne) AddUpdatedBy(i int) *UserLoginProfileUpdateOne {
	ulpuo.mutation.AddUpdatedBy(i)
	return ulpuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ulpuo *UserLoginProfileUpdateOne) ClearUpdatedBy() *UserLoginProfileUpdateOne {
	ulpuo.mutation.ClearUpdatedBy()
	return ulpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ulpuo *UserLoginProfileUpdateOne) SetUpdatedAt(t time.Time) *UserLoginProfileUpdateOne {
	ulpuo.mutation.SetUpdatedAt(t)
	return ulpuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ulpuo *UserLoginProfileUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserLoginProfileUpdateOne {
	if t != nil {
		ulpuo.SetUpdatedAt(*t)
	}
	return ulpuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ulpuo *UserLoginProfileUpdateOne) ClearUpdatedAt() *UserLoginProfileUpdateOne {
	ulpuo.mutation.ClearUpdatedAt()
	return ulpuo
}

// SetLastLoginIP sets the "last_login_ip" field.
func (ulpuo *UserLoginProfileUpdateOne) SetLastLoginIP(s string) *UserLoginProfileUpdateOne {
	ulpuo.mutation.SetLastLoginIP(s)
	return ulpuo
}

// SetNillableLastLoginIP sets the "last_login_ip" field if the given value is not nil.
func (ulpuo *UserLoginProfileUpdateOne) SetNillableLastLoginIP(s *string) *UserLoginProfileUpdateOne {
	if s != nil {
		ulpuo.SetLastLoginIP(*s)
	}
	return ulpuo
}

// ClearLastLoginIP clears the value of the "last_login_ip" field.
func (ulpuo *UserLoginProfileUpdateOne) ClearLastLoginIP() *UserLoginProfileUpdateOne {
	ulpuo.mutation.ClearLastLoginIP()
	return ulpuo
}

// SetLastLoginAt sets the "last_login_at" field.
func (ulpuo *UserLoginProfileUpdateOne) SetLastLoginAt(t time.Time) *UserLoginProfileUpdateOne {
	ulpuo.mutation.SetLastLoginAt(t)
	return ulpuo
}

// SetNillableLastLoginAt sets the "last_login_at" field if the given value is not nil.
func (ulpuo *UserLoginProfileUpdateOne) SetNillableLastLoginAt(t *time.Time) *UserLoginProfileUpdateOne {
	if t != nil {
		ulpuo.SetLastLoginAt(*t)
	}
	return ulpuo
}

// ClearLastLoginAt clears the value of the "last_login_at" field.
func (ulpuo *UserLoginProfileUpdateOne) ClearLastLoginAt() *UserLoginProfileUpdateOne {
	ulpuo.mutation.ClearLastLoginAt()
	return ulpuo
}

// SetCanLogin sets the "can_login" field.
func (ulpuo *UserLoginProfileUpdateOne) SetCanLogin(b bool) *UserLoginProfileUpdateOne {
	ulpuo.mutation.SetCanLogin(b)
	return ulpuo
}

// SetNillableCanLogin sets the "can_login" field if the given value is not nil.
func (ulpuo *UserLoginProfileUpdateOne) SetNillableCanLogin(b *bool) *UserLoginProfileUpdateOne {
	if b != nil {
		ulpuo.SetCanLogin(*b)
	}
	return ulpuo
}

// ClearCanLogin clears the value of the "can_login" field.
func (ulpuo *UserLoginProfileUpdateOne) ClearCanLogin() *UserLoginProfileUpdateOne {
	ulpuo.mutation.ClearCanLogin()
	return ulpuo
}

// SetSetKind sets the "set_kind" field.
func (ulpuo *UserLoginProfileUpdateOne) SetSetKind(uk userloginprofile.SetKind) *UserLoginProfileUpdateOne {
	ulpuo.mutation.SetSetKind(uk)
	return ulpuo
}

// SetPasswordReset sets the "password_reset" field.
func (ulpuo *UserLoginProfileUpdateOne) SetPasswordReset(b bool) *UserLoginProfileUpdateOne {
	ulpuo.mutation.SetPasswordReset(b)
	return ulpuo
}

// SetNillablePasswordReset sets the "password_reset" field if the given value is not nil.
func (ulpuo *UserLoginProfileUpdateOne) SetNillablePasswordReset(b *bool) *UserLoginProfileUpdateOne {
	if b != nil {
		ulpuo.SetPasswordReset(*b)
	}
	return ulpuo
}

// ClearPasswordReset clears the value of the "password_reset" field.
func (ulpuo *UserLoginProfileUpdateOne) ClearPasswordReset() *UserLoginProfileUpdateOne {
	ulpuo.mutation.ClearPasswordReset()
	return ulpuo
}

// SetVerifyDevice sets the "verify_device" field.
func (ulpuo *UserLoginProfileUpdateOne) SetVerifyDevice(b bool) *UserLoginProfileUpdateOne {
	ulpuo.mutation.SetVerifyDevice(b)
	return ulpuo
}

// SetMfaEnabled sets the "mfa_enabled" field.
func (ulpuo *UserLoginProfileUpdateOne) SetMfaEnabled(b bool) *UserLoginProfileUpdateOne {
	ulpuo.mutation.SetMfaEnabled(b)
	return ulpuo
}

// SetNillableMfaEnabled sets the "mfa_enabled" field if the given value is not nil.
func (ulpuo *UserLoginProfileUpdateOne) SetNillableMfaEnabled(b *bool) *UserLoginProfileUpdateOne {
	if b != nil {
		ulpuo.SetMfaEnabled(*b)
	}
	return ulpuo
}

// ClearMfaEnabled clears the value of the "mfa_enabled" field.
func (ulpuo *UserLoginProfileUpdateOne) ClearMfaEnabled() *UserLoginProfileUpdateOne {
	ulpuo.mutation.ClearMfaEnabled()
	return ulpuo
}

// SetMfaSecret sets the "mfa_secret" field.
func (ulpuo *UserLoginProfileUpdateOne) SetMfaSecret(s string) *UserLoginProfileUpdateOne {
	ulpuo.mutation.SetMfaSecret(s)
	return ulpuo
}

// SetNillableMfaSecret sets the "mfa_secret" field if the given value is not nil.
func (ulpuo *UserLoginProfileUpdateOne) SetNillableMfaSecret(s *string) *UserLoginProfileUpdateOne {
	if s != nil {
		ulpuo.SetMfaSecret(*s)
	}
	return ulpuo
}

// ClearMfaSecret clears the value of the "mfa_secret" field.
func (ulpuo *UserLoginProfileUpdateOne) ClearMfaSecret() *UserLoginProfileUpdateOne {
	ulpuo.mutation.ClearMfaSecret()
	return ulpuo
}

// SetMfaStatus sets the "mfa_status" field.
func (ulpuo *UserLoginProfileUpdateOne) SetMfaStatus(ts typex.SimpleStatus) *UserLoginProfileUpdateOne {
	ulpuo.mutation.SetMfaStatus(ts)
	return ulpuo
}

// SetNillableMfaStatus sets the "mfa_status" field if the given value is not nil.
func (ulpuo *UserLoginProfileUpdateOne) SetNillableMfaStatus(ts *typex.SimpleStatus) *UserLoginProfileUpdateOne {
	if ts != nil {
		ulpuo.SetMfaStatus(*ts)
	}
	return ulpuo
}

// ClearMfaStatus clears the value of the "mfa_status" field.
func (ulpuo *UserLoginProfileUpdateOne) ClearMfaStatus() *UserLoginProfileUpdateOne {
	ulpuo.mutation.ClearMfaStatus()
	return ulpuo
}

// Mutation returns the UserLoginProfileMutation object of the builder.
func (ulpuo *UserLoginProfileUpdateOne) Mutation() *UserLoginProfileMutation {
	return ulpuo.mutation
}

// Where appends a list predicates to the UserLoginProfileUpdate builder.
func (ulpuo *UserLoginProfileUpdateOne) Where(ps ...predicate.UserLoginProfile) *UserLoginProfileUpdateOne {
	ulpuo.mutation.Where(ps...)
	return ulpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ulpuo *UserLoginProfileUpdateOne) Select(field string, fields ...string) *UserLoginProfileUpdateOne {
	ulpuo.fields = append([]string{field}, fields...)
	return ulpuo
}

// Save executes the query and returns the updated UserLoginProfile entity.
func (ulpuo *UserLoginProfileUpdateOne) Save(ctx context.Context) (*UserLoginProfile, error) {
	return withHooks(ctx, ulpuo.sqlSave, ulpuo.mutation, ulpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ulpuo *UserLoginProfileUpdateOne) SaveX(ctx context.Context) *UserLoginProfile {
	node, err := ulpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ulpuo *UserLoginProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := ulpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ulpuo *UserLoginProfileUpdateOne) ExecX(ctx context.Context) {
	if err := ulpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ulpuo *UserLoginProfileUpdateOne) check() error {
	if v, ok := ulpuo.mutation.SetKind(); ok {
		if err := userloginprofile.SetKindValidator(v); err != nil {
			return &ValidationError{Name: "set_kind", err: fmt.Errorf(`ent: validator failed for field "UserLoginProfile.set_kind": %w`, err)}
		}
	}
	if v, ok := ulpuo.mutation.MfaSecret(); ok {
		if err := userloginprofile.MfaSecretValidator(v); err != nil {
			return &ValidationError{Name: "mfa_secret", err: fmt.Errorf(`ent: validator failed for field "UserLoginProfile.mfa_secret": %w`, err)}
		}
	}
	if v, ok := ulpuo.mutation.MfaStatus(); ok {
		if err := userloginprofile.MfaStatusValidator(v); err != nil {
			return &ValidationError{Name: "mfa_status", err: fmt.Errorf(`ent: validator failed for field "UserLoginProfile.mfa_status": %w`, err)}
		}
	}
	return nil
}

func (ulpuo *UserLoginProfileUpdateOne) sqlSave(ctx context.Context) (_node *UserLoginProfile, err error) {
	if err := ulpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userloginprofile.Table, userloginprofile.Columns, sqlgraph.NewFieldSpec(userloginprofile.FieldID, field.TypeInt))
	id, ok := ulpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserLoginProfile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ulpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userloginprofile.FieldID)
		for _, f := range fields {
			if !userloginprofile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userloginprofile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ulpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ulpuo.mutation.UpdatedBy(); ok {
		_spec.SetField(userloginprofile.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := ulpuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(userloginprofile.FieldUpdatedBy, field.TypeInt, value)
	}
	if ulpuo.mutation.UpdatedByCleared() {
		_spec.ClearField(userloginprofile.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := ulpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userloginprofile.FieldUpdatedAt, field.TypeTime, value)
	}
	if ulpuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(userloginprofile.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := ulpuo.mutation.LastLoginIP(); ok {
		_spec.SetField(userloginprofile.FieldLastLoginIP, field.TypeString, value)
	}
	if ulpuo.mutation.LastLoginIPCleared() {
		_spec.ClearField(userloginprofile.FieldLastLoginIP, field.TypeString)
	}
	if value, ok := ulpuo.mutation.LastLoginAt(); ok {
		_spec.SetField(userloginprofile.FieldLastLoginAt, field.TypeTime, value)
	}
	if ulpuo.mutation.LastLoginAtCleared() {
		_spec.ClearField(userloginprofile.FieldLastLoginAt, field.TypeTime)
	}
	if value, ok := ulpuo.mutation.CanLogin(); ok {
		_spec.SetField(userloginprofile.FieldCanLogin, field.TypeBool, value)
	}
	if ulpuo.mutation.CanLoginCleared() {
		_spec.ClearField(userloginprofile.FieldCanLogin, field.TypeBool)
	}
	if value, ok := ulpuo.mutation.SetKind(); ok {
		_spec.SetField(userloginprofile.FieldSetKind, field.TypeEnum, value)
	}
	if value, ok := ulpuo.mutation.PasswordReset(); ok {
		_spec.SetField(userloginprofile.FieldPasswordReset, field.TypeBool, value)
	}
	if ulpuo.mutation.PasswordResetCleared() {
		_spec.ClearField(userloginprofile.FieldPasswordReset, field.TypeBool)
	}
	if value, ok := ulpuo.mutation.VerifyDevice(); ok {
		_spec.SetField(userloginprofile.FieldVerifyDevice, field.TypeBool, value)
	}
	if value, ok := ulpuo.mutation.MfaEnabled(); ok {
		_spec.SetField(userloginprofile.FieldMfaEnabled, field.TypeBool, value)
	}
	if ulpuo.mutation.MfaEnabledCleared() {
		_spec.ClearField(userloginprofile.FieldMfaEnabled, field.TypeBool)
	}
	if value, ok := ulpuo.mutation.MfaSecret(); ok {
		_spec.SetField(userloginprofile.FieldMfaSecret, field.TypeString, value)
	}
	if ulpuo.mutation.MfaSecretCleared() {
		_spec.ClearField(userloginprofile.FieldMfaSecret, field.TypeString)
	}
	if value, ok := ulpuo.mutation.MfaStatus(); ok {
		_spec.SetField(userloginprofile.FieldMfaStatus, field.TypeEnum, value)
	}
	if ulpuo.mutation.MfaStatusCleared() {
		_spec.ClearField(userloginprofile.FieldMfaStatus, field.TypeEnum)
	}
	_node = &UserLoginProfile{config: ulpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ulpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userloginprofile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ulpuo.mutation.done = true
	return _node, nil
}
