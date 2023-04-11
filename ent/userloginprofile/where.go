// Code generated by ent, DO NOT EDIT.

package userloginprofile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldUserID, v))
}

// LastLoginIP applies equality check predicate on the "last_login_ip" field. It's identical to LastLoginIPEQ.
func LastLoginIP(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldLastLoginIP, v))
}

// LastLoginAt applies equality check predicate on the "last_login_at" field. It's identical to LastLoginAtEQ.
func LastLoginAt(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldLastLoginAt, v))
}

// CanLogin applies equality check predicate on the "can_login" field. It's identical to CanLoginEQ.
func CanLogin(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldCanLogin, v))
}

// PasswordReset applies equality check predicate on the "password_reset" field. It's identical to PasswordResetEQ.
func PasswordReset(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldPasswordReset, v))
}

// VerifyDevice applies equality check predicate on the "verify_device" field. It's identical to VerifyDeviceEQ.
func VerifyDevice(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldVerifyDevice, v))
}

// MfaEnabled applies equality check predicate on the "mfa_enabled" field. It's identical to MfaEnabledEQ.
func MfaEnabled(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldMfaEnabled, v))
}

// MfaSecret applies equality check predicate on the "mfa_secret" field. It's identical to MfaSecretEQ.
func MfaSecret(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldMfaSecret, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotNull(FieldUpdatedAt))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotNull(FieldUserID))
}

// LastLoginIPEQ applies the EQ predicate on the "last_login_ip" field.
func LastLoginIPEQ(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldLastLoginIP, v))
}

// LastLoginIPNEQ applies the NEQ predicate on the "last_login_ip" field.
func LastLoginIPNEQ(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldLastLoginIP, v))
}

// LastLoginIPIn applies the In predicate on the "last_login_ip" field.
func LastLoginIPIn(vs ...string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIn(FieldLastLoginIP, vs...))
}

// LastLoginIPNotIn applies the NotIn predicate on the "last_login_ip" field.
func LastLoginIPNotIn(vs ...string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotIn(FieldLastLoginIP, vs...))
}

// LastLoginIPGT applies the GT predicate on the "last_login_ip" field.
func LastLoginIPGT(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGT(FieldLastLoginIP, v))
}

// LastLoginIPGTE applies the GTE predicate on the "last_login_ip" field.
func LastLoginIPGTE(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGTE(FieldLastLoginIP, v))
}

// LastLoginIPLT applies the LT predicate on the "last_login_ip" field.
func LastLoginIPLT(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLT(FieldLastLoginIP, v))
}

// LastLoginIPLTE applies the LTE predicate on the "last_login_ip" field.
func LastLoginIPLTE(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLTE(FieldLastLoginIP, v))
}

// LastLoginIPContains applies the Contains predicate on the "last_login_ip" field.
func LastLoginIPContains(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldContains(FieldLastLoginIP, v))
}

// LastLoginIPHasPrefix applies the HasPrefix predicate on the "last_login_ip" field.
func LastLoginIPHasPrefix(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldHasPrefix(FieldLastLoginIP, v))
}

// LastLoginIPHasSuffix applies the HasSuffix predicate on the "last_login_ip" field.
func LastLoginIPHasSuffix(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldHasSuffix(FieldLastLoginIP, v))
}

// LastLoginIPIsNil applies the IsNil predicate on the "last_login_ip" field.
func LastLoginIPIsNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIsNull(FieldLastLoginIP))
}

// LastLoginIPNotNil applies the NotNil predicate on the "last_login_ip" field.
func LastLoginIPNotNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotNull(FieldLastLoginIP))
}

// LastLoginIPEqualFold applies the EqualFold predicate on the "last_login_ip" field.
func LastLoginIPEqualFold(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEqualFold(FieldLastLoginIP, v))
}

// LastLoginIPContainsFold applies the ContainsFold predicate on the "last_login_ip" field.
func LastLoginIPContainsFold(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldContainsFold(FieldLastLoginIP, v))
}

// LastLoginAtEQ applies the EQ predicate on the "last_login_at" field.
func LastLoginAtEQ(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldLastLoginAt, v))
}

// LastLoginAtNEQ applies the NEQ predicate on the "last_login_at" field.
func LastLoginAtNEQ(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldLastLoginAt, v))
}

// LastLoginAtIn applies the In predicate on the "last_login_at" field.
func LastLoginAtIn(vs ...time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIn(FieldLastLoginAt, vs...))
}

// LastLoginAtNotIn applies the NotIn predicate on the "last_login_at" field.
func LastLoginAtNotIn(vs ...time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotIn(FieldLastLoginAt, vs...))
}

// LastLoginAtGT applies the GT predicate on the "last_login_at" field.
func LastLoginAtGT(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGT(FieldLastLoginAt, v))
}

// LastLoginAtGTE applies the GTE predicate on the "last_login_at" field.
func LastLoginAtGTE(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGTE(FieldLastLoginAt, v))
}

// LastLoginAtLT applies the LT predicate on the "last_login_at" field.
func LastLoginAtLT(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLT(FieldLastLoginAt, v))
}

// LastLoginAtLTE applies the LTE predicate on the "last_login_at" field.
func LastLoginAtLTE(v time.Time) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLTE(FieldLastLoginAt, v))
}

// LastLoginAtIsNil applies the IsNil predicate on the "last_login_at" field.
func LastLoginAtIsNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIsNull(FieldLastLoginAt))
}

// LastLoginAtNotNil applies the NotNil predicate on the "last_login_at" field.
func LastLoginAtNotNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotNull(FieldLastLoginAt))
}

// CanLoginEQ applies the EQ predicate on the "can_login" field.
func CanLoginEQ(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldCanLogin, v))
}

// CanLoginNEQ applies the NEQ predicate on the "can_login" field.
func CanLoginNEQ(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldCanLogin, v))
}

// CanLoginIsNil applies the IsNil predicate on the "can_login" field.
func CanLoginIsNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIsNull(FieldCanLogin))
}

// CanLoginNotNil applies the NotNil predicate on the "can_login" field.
func CanLoginNotNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotNull(FieldCanLogin))
}

// SetKindEQ applies the EQ predicate on the "set_kind" field.
func SetKindEQ(v SetKind) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldSetKind, v))
}

// SetKindNEQ applies the NEQ predicate on the "set_kind" field.
func SetKindNEQ(v SetKind) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldSetKind, v))
}

// SetKindIn applies the In predicate on the "set_kind" field.
func SetKindIn(vs ...SetKind) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIn(FieldSetKind, vs...))
}

// SetKindNotIn applies the NotIn predicate on the "set_kind" field.
func SetKindNotIn(vs ...SetKind) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotIn(FieldSetKind, vs...))
}

// PasswordResetEQ applies the EQ predicate on the "password_reset" field.
func PasswordResetEQ(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldPasswordReset, v))
}

// PasswordResetNEQ applies the NEQ predicate on the "password_reset" field.
func PasswordResetNEQ(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldPasswordReset, v))
}

// PasswordResetIsNil applies the IsNil predicate on the "password_reset" field.
func PasswordResetIsNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIsNull(FieldPasswordReset))
}

// PasswordResetNotNil applies the NotNil predicate on the "password_reset" field.
func PasswordResetNotNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotNull(FieldPasswordReset))
}

// VerifyDeviceEQ applies the EQ predicate on the "verify_device" field.
func VerifyDeviceEQ(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldVerifyDevice, v))
}

// VerifyDeviceNEQ applies the NEQ predicate on the "verify_device" field.
func VerifyDeviceNEQ(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldVerifyDevice, v))
}

// MfaEnabledEQ applies the EQ predicate on the "mfa_enabled" field.
func MfaEnabledEQ(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldMfaEnabled, v))
}

// MfaEnabledNEQ applies the NEQ predicate on the "mfa_enabled" field.
func MfaEnabledNEQ(v bool) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldMfaEnabled, v))
}

// MfaEnabledIsNil applies the IsNil predicate on the "mfa_enabled" field.
func MfaEnabledIsNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIsNull(FieldMfaEnabled))
}

// MfaEnabledNotNil applies the NotNil predicate on the "mfa_enabled" field.
func MfaEnabledNotNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotNull(FieldMfaEnabled))
}

// MfaSecretEQ applies the EQ predicate on the "mfa_secret" field.
func MfaSecretEQ(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEQ(FieldMfaSecret, v))
}

// MfaSecretNEQ applies the NEQ predicate on the "mfa_secret" field.
func MfaSecretNEQ(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldMfaSecret, v))
}

// MfaSecretIn applies the In predicate on the "mfa_secret" field.
func MfaSecretIn(vs ...string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIn(FieldMfaSecret, vs...))
}

// MfaSecretNotIn applies the NotIn predicate on the "mfa_secret" field.
func MfaSecretNotIn(vs ...string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotIn(FieldMfaSecret, vs...))
}

// MfaSecretGT applies the GT predicate on the "mfa_secret" field.
func MfaSecretGT(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGT(FieldMfaSecret, v))
}

// MfaSecretGTE applies the GTE predicate on the "mfa_secret" field.
func MfaSecretGTE(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldGTE(FieldMfaSecret, v))
}

// MfaSecretLT applies the LT predicate on the "mfa_secret" field.
func MfaSecretLT(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLT(FieldMfaSecret, v))
}

// MfaSecretLTE applies the LTE predicate on the "mfa_secret" field.
func MfaSecretLTE(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldLTE(FieldMfaSecret, v))
}

// MfaSecretContains applies the Contains predicate on the "mfa_secret" field.
func MfaSecretContains(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldContains(FieldMfaSecret, v))
}

// MfaSecretHasPrefix applies the HasPrefix predicate on the "mfa_secret" field.
func MfaSecretHasPrefix(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldHasPrefix(FieldMfaSecret, v))
}

// MfaSecretHasSuffix applies the HasSuffix predicate on the "mfa_secret" field.
func MfaSecretHasSuffix(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldHasSuffix(FieldMfaSecret, v))
}

// MfaSecretIsNil applies the IsNil predicate on the "mfa_secret" field.
func MfaSecretIsNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIsNull(FieldMfaSecret))
}

// MfaSecretNotNil applies the NotNil predicate on the "mfa_secret" field.
func MfaSecretNotNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotNull(FieldMfaSecret))
}

// MfaSecretEqualFold applies the EqualFold predicate on the "mfa_secret" field.
func MfaSecretEqualFold(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldEqualFold(FieldMfaSecret, v))
}

// MfaSecretContainsFold applies the ContainsFold predicate on the "mfa_secret" field.
func MfaSecretContainsFold(v string) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldContainsFold(FieldMfaSecret, v))
}

// MfaStatusEQ applies the EQ predicate on the "mfa_status" field.
func MfaStatusEQ(v typex.SimpleStatus) predicate.UserLoginProfile {
	vc := v
	return predicate.UserLoginProfile(sql.FieldEQ(FieldMfaStatus, vc))
}

// MfaStatusNEQ applies the NEQ predicate on the "mfa_status" field.
func MfaStatusNEQ(v typex.SimpleStatus) predicate.UserLoginProfile {
	vc := v
	return predicate.UserLoginProfile(sql.FieldNEQ(FieldMfaStatus, vc))
}

// MfaStatusIn applies the In predicate on the "mfa_status" field.
func MfaStatusIn(vs ...typex.SimpleStatus) predicate.UserLoginProfile {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserLoginProfile(sql.FieldIn(FieldMfaStatus, v...))
}

// MfaStatusNotIn applies the NotIn predicate on the "mfa_status" field.
func MfaStatusNotIn(vs ...typex.SimpleStatus) predicate.UserLoginProfile {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserLoginProfile(sql.FieldNotIn(FieldMfaStatus, v...))
}

// MfaStatusIsNil applies the IsNil predicate on the "mfa_status" field.
func MfaStatusIsNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldIsNull(FieldMfaStatus))
}

// MfaStatusNotNil applies the NotNil predicate on the "mfa_status" field.
func MfaStatusNotNil() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(sql.FieldNotNull(FieldMfaStatus))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserLoginProfile {
	return predicate.UserLoginProfile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserLoginProfile) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserLoginProfile) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserLoginProfile) predicate.UserLoginProfile {
	return predicate.UserLoginProfile(func(s *sql.Selector) {
		p(s.Not())
	})
}
