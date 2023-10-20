// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeletedAt, v))
}

// PrincipalName applies equality check predicate on the "principal_name" field. It's identical to PrincipalNameEQ.
func PrincipalName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPrincipalName, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// RegisterIP applies equality check predicate on the "register_ip" field. It's identical to RegisterIPEQ.
func RegisterIP(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRegisterIP, v))
}

// Comments applies equality check predicate on the "comments" field. It's identical to CommentsEQ.
func Comments(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldComments, v))
}

// AvatarFileID applies equality check predicate on the "avatar_file_id" field. It's identical to AvatarFileIDEQ.
func AvatarFileID(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarFileID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldDeletedAt))
}

// PrincipalNameEQ applies the EQ predicate on the "principal_name" field.
func PrincipalNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPrincipalName, v))
}

// PrincipalNameNEQ applies the NEQ predicate on the "principal_name" field.
func PrincipalNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPrincipalName, v))
}

// PrincipalNameIn applies the In predicate on the "principal_name" field.
func PrincipalNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPrincipalName, vs...))
}

// PrincipalNameNotIn applies the NotIn predicate on the "principal_name" field.
func PrincipalNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPrincipalName, vs...))
}

// PrincipalNameGT applies the GT predicate on the "principal_name" field.
func PrincipalNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPrincipalName, v))
}

// PrincipalNameGTE applies the GTE predicate on the "principal_name" field.
func PrincipalNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPrincipalName, v))
}

// PrincipalNameLT applies the LT predicate on the "principal_name" field.
func PrincipalNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPrincipalName, v))
}

// PrincipalNameLTE applies the LTE predicate on the "principal_name" field.
func PrincipalNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPrincipalName, v))
}

// PrincipalNameContains applies the Contains predicate on the "principal_name" field.
func PrincipalNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPrincipalName, v))
}

// PrincipalNameHasPrefix applies the HasPrefix predicate on the "principal_name" field.
func PrincipalNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPrincipalName, v))
}

// PrincipalNameHasSuffix applies the HasSuffix predicate on the "principal_name" field.
func PrincipalNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPrincipalName, v))
}

// PrincipalNameEqualFold applies the EqualFold predicate on the "principal_name" field.
func PrincipalNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPrincipalName, v))
}

// PrincipalNameContainsFold applies the ContainsFold predicate on the "principal_name" field.
func PrincipalNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPrincipalName, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldDisplayName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileIsNil applies the IsNil predicate on the "mobile" field.
func MobileIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldMobile))
}

// MobileNotNil applies the NotNil predicate on the "mobile" field.
func MobileNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldMobile))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldMobile, v))
}

// UserTypeEQ applies the EQ predicate on the "user_type" field.
func UserTypeEQ(v UserType) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUserType, v))
}

// UserTypeNEQ applies the NEQ predicate on the "user_type" field.
func UserTypeNEQ(v UserType) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUserType, v))
}

// UserTypeIn applies the In predicate on the "user_type" field.
func UserTypeIn(vs ...UserType) predicate.User {
	return predicate.User(sql.FieldIn(FieldUserType, vs...))
}

// UserTypeNotIn applies the NotIn predicate on the "user_type" field.
func UserTypeNotIn(vs ...UserType) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUserType, vs...))
}

// CreationTypeEQ applies the EQ predicate on the "creation_type" field.
func CreationTypeEQ(v CreationType) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreationType, v))
}

// CreationTypeNEQ applies the NEQ predicate on the "creation_type" field.
func CreationTypeNEQ(v CreationType) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreationType, v))
}

// CreationTypeIn applies the In predicate on the "creation_type" field.
func CreationTypeIn(vs ...CreationType) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreationType, vs...))
}

// CreationTypeNotIn applies the NotIn predicate on the "creation_type" field.
func CreationTypeNotIn(vs ...CreationType) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreationType, vs...))
}

// RegisterIPEQ applies the EQ predicate on the "register_ip" field.
func RegisterIPEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRegisterIP, v))
}

// RegisterIPNEQ applies the NEQ predicate on the "register_ip" field.
func RegisterIPNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRegisterIP, v))
}

// RegisterIPIn applies the In predicate on the "register_ip" field.
func RegisterIPIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldRegisterIP, vs...))
}

// RegisterIPNotIn applies the NotIn predicate on the "register_ip" field.
func RegisterIPNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRegisterIP, vs...))
}

// RegisterIPGT applies the GT predicate on the "register_ip" field.
func RegisterIPGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldRegisterIP, v))
}

// RegisterIPGTE applies the GTE predicate on the "register_ip" field.
func RegisterIPGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldRegisterIP, v))
}

// RegisterIPLT applies the LT predicate on the "register_ip" field.
func RegisterIPLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldRegisterIP, v))
}

// RegisterIPLTE applies the LTE predicate on the "register_ip" field.
func RegisterIPLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldRegisterIP, v))
}

// RegisterIPContains applies the Contains predicate on the "register_ip" field.
func RegisterIPContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldRegisterIP, v))
}

// RegisterIPHasPrefix applies the HasPrefix predicate on the "register_ip" field.
func RegisterIPHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldRegisterIP, v))
}

// RegisterIPHasSuffix applies the HasSuffix predicate on the "register_ip" field.
func RegisterIPHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldRegisterIP, v))
}

// RegisterIPEqualFold applies the EqualFold predicate on the "register_ip" field.
func RegisterIPEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldRegisterIP, v))
}

// RegisterIPContainsFold applies the ContainsFold predicate on the "register_ip" field.
func RegisterIPContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldRegisterIP, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v typex.SimpleStatus) predicate.User {
	vc := v
	return predicate.User(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v typex.SimpleStatus) predicate.User {
	vc := v
	return predicate.User(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...typex.SimpleStatus) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...typex.SimpleStatus) predicate.User {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(sql.FieldNotIn(FieldStatus, v...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldStatus))
}

// CommentsEQ applies the EQ predicate on the "comments" field.
func CommentsEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldComments, v))
}

// CommentsNEQ applies the NEQ predicate on the "comments" field.
func CommentsNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldComments, v))
}

// CommentsIn applies the In predicate on the "comments" field.
func CommentsIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldComments, vs...))
}

// CommentsNotIn applies the NotIn predicate on the "comments" field.
func CommentsNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldComments, vs...))
}

// CommentsGT applies the GT predicate on the "comments" field.
func CommentsGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldComments, v))
}

// CommentsGTE applies the GTE predicate on the "comments" field.
func CommentsGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldComments, v))
}

// CommentsLT applies the LT predicate on the "comments" field.
func CommentsLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldComments, v))
}

// CommentsLTE applies the LTE predicate on the "comments" field.
func CommentsLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldComments, v))
}

// CommentsContains applies the Contains predicate on the "comments" field.
func CommentsContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldComments, v))
}

// CommentsHasPrefix applies the HasPrefix predicate on the "comments" field.
func CommentsHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldComments, v))
}

// CommentsHasSuffix applies the HasSuffix predicate on the "comments" field.
func CommentsHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldComments, v))
}

// CommentsIsNil applies the IsNil predicate on the "comments" field.
func CommentsIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldComments))
}

// CommentsNotNil applies the NotNil predicate on the "comments" field.
func CommentsNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldComments))
}

// CommentsEqualFold applies the EqualFold predicate on the "comments" field.
func CommentsEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldComments, v))
}

// CommentsContainsFold applies the ContainsFold predicate on the "comments" field.
func CommentsContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldComments, v))
}

// AvatarFileIDEQ applies the EQ predicate on the "avatar_file_id" field.
func AvatarFileIDEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAvatarFileID, v))
}

// AvatarFileIDNEQ applies the NEQ predicate on the "avatar_file_id" field.
func AvatarFileIDNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAvatarFileID, v))
}

// AvatarFileIDIn applies the In predicate on the "avatar_file_id" field.
func AvatarFileIDIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldAvatarFileID, vs...))
}

// AvatarFileIDNotIn applies the NotIn predicate on the "avatar_file_id" field.
func AvatarFileIDNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAvatarFileID, vs...))
}

// AvatarFileIDGT applies the GT predicate on the "avatar_file_id" field.
func AvatarFileIDGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldAvatarFileID, v))
}

// AvatarFileIDGTE applies the GTE predicate on the "avatar_file_id" field.
func AvatarFileIDGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAvatarFileID, v))
}

// AvatarFileIDLT applies the LT predicate on the "avatar_file_id" field.
func AvatarFileIDLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldAvatarFileID, v))
}

// AvatarFileIDLTE applies the LTE predicate on the "avatar_file_id" field.
func AvatarFileIDLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAvatarFileID, v))
}

// AvatarFileIDIsNil applies the IsNil predicate on the "avatar_file_id" field.
func AvatarFileIDIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAvatarFileID))
}

// AvatarFileIDNotNil applies the NotNil predicate on the "avatar_file_id" field.
func AvatarFileIDNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAvatarFileID))
}

// HasIdentities applies the HasEdge predicate on the "identities" edge.
func HasIdentities() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IdentitiesTable, IdentitiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIdentitiesWith applies the HasEdge predicate on the "identities" edge with a given conditions (other predicates).
func HasIdentitiesWith(preds ...predicate.UserIdentity) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newIdentitiesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLoginProfile applies the HasEdge predicate on the "login_profile" edge.
func HasLoginProfile() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, LoginProfileTable, LoginProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLoginProfileWith applies the HasEdge predicate on the "login_profile" edge with a given conditions (other predicates).
func HasLoginProfileWith(preds ...predicate.UserLoginProfile) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newLoginProfileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPasswords applies the HasEdge predicate on the "passwords" edge.
func HasPasswords() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PasswordsTable, PasswordsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPasswordsWith applies the HasEdge predicate on the "passwords" edge with a given conditions (other predicates).
func HasPasswordsWith(preds ...predicate.UserPassword) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPasswordsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDevices applies the HasEdge predicate on the "devices" edge.
func HasDevices() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DevicesTable, DevicesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDevicesWith applies the HasEdge predicate on the "devices" edge with a given conditions (other predicates).
func HasDevicesWith(preds ...predicate.UserDevice) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newDevicesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrgs applies the HasEdge predicate on the "orgs" edge.
func HasOrgs() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OrgsTable, OrgsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrgsWith applies the HasEdge predicate on the "orgs" edge with a given conditions (other predicates).
func HasOrgsWith(preds ...predicate.Org) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newOrgsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPermissions applies the HasEdge predicate on the "permissions" edge.
func HasPermissions() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, PermissionsTable, PermissionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionsWith applies the HasEdge predicate on the "permissions" edge with a given conditions (other predicates).
func HasPermissionsWith(preds ...predicate.Permission) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPermissionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOauthClients applies the HasEdge predicate on the "oauth_clients" edge.
func HasOauthClients() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OauthClientsTable, OauthClientsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOauthClientsWith applies the HasEdge predicate on the "oauth_clients" edge with a given conditions (other predicates).
func HasOauthClientsWith(preds ...predicate.OauthClient) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newOauthClientsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrgUser applies the HasEdge predicate on the "org_user" edge.
func HasOrgUser() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, OrgUserTable, OrgUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrgUserWith applies the HasEdge predicate on the "org_user" edge with a given conditions (other predicates).
func HasOrgUserWith(preds ...predicate.OrgUser) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newOrgUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
