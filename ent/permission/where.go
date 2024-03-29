// Code generated by ent, DO NOT EDIT.

package permission

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrgID applies equality check predicate on the "org_id" field. It's identical to OrgIDEQ.
func OrgID(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldOrgID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUserID, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldRoleID, v))
}

// OrgPolicyID applies equality check predicate on the "org_policy_id" field. It's identical to OrgPolicyIDEQ.
func OrgPolicyID(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldOrgPolicyID, v))
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldStartAt, v))
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldEndAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldUpdatedAt))
}

// OrgIDEQ applies the EQ predicate on the "org_id" field.
func OrgIDEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldOrgID, v))
}

// OrgIDNEQ applies the NEQ predicate on the "org_id" field.
func OrgIDNEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldOrgID, v))
}

// OrgIDIn applies the In predicate on the "org_id" field.
func OrgIDIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldOrgID, vs...))
}

// OrgIDNotIn applies the NotIn predicate on the "org_id" field.
func OrgIDNotIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldOrgID, vs...))
}

// PrincipalKindEQ applies the EQ predicate on the "principal_kind" field.
func PrincipalKindEQ(v PrincipalKind) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldPrincipalKind, v))
}

// PrincipalKindNEQ applies the NEQ predicate on the "principal_kind" field.
func PrincipalKindNEQ(v PrincipalKind) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldPrincipalKind, v))
}

// PrincipalKindIn applies the In predicate on the "principal_kind" field.
func PrincipalKindIn(vs ...PrincipalKind) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldPrincipalKind, vs...))
}

// PrincipalKindNotIn applies the NotIn predicate on the "principal_kind" field.
func PrincipalKindNotIn(vs ...PrincipalKind) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldPrincipalKind, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldUserID))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDIsNil applies the IsNil predicate on the "role_id" field.
func RoleIDIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldRoleID))
}

// RoleIDNotNil applies the NotNil predicate on the "role_id" field.
func RoleIDNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldRoleID))
}

// OrgPolicyIDEQ applies the EQ predicate on the "org_policy_id" field.
func OrgPolicyIDEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldOrgPolicyID, v))
}

// OrgPolicyIDNEQ applies the NEQ predicate on the "org_policy_id" field.
func OrgPolicyIDNEQ(v int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldOrgPolicyID, v))
}

// OrgPolicyIDIn applies the In predicate on the "org_policy_id" field.
func OrgPolicyIDIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldOrgPolicyID, vs...))
}

// OrgPolicyIDNotIn applies the NotIn predicate on the "org_policy_id" field.
func OrgPolicyIDNotIn(vs ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldOrgPolicyID, vs...))
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldStartAt, v))
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldStartAt, v))
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldStartAt, vs...))
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldStartAt, vs...))
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldStartAt, v))
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldStartAt, v))
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldStartAt, v))
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldStartAt, v))
}

// StartAtIsNil applies the IsNil predicate on the "start_at" field.
func StartAtIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldStartAt))
}

// StartAtNotNil applies the NotNil predicate on the "start_at" field.
func StartAtNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldStartAt))
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldEndAt, v))
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldEndAt, v))
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldEndAt, vs...))
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldEndAt, vs...))
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldEndAt, v))
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldEndAt, v))
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldEndAt, v))
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldEndAt, v))
}

// EndAtIsNil applies the IsNil predicate on the "end_at" field.
func EndAtIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldEndAt))
}

// EndAtNotNil applies the NotNil predicate on the "end_at" field.
func EndAtNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldEndAt))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v typex.SimpleStatus) predicate.Permission {
	vc := v
	return predicate.Permission(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v typex.SimpleStatus) predicate.Permission {
	vc := v
	return predicate.Permission(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...typex.SimpleStatus) predicate.Permission {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...typex.SimpleStatus) predicate.Permission {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Permission(sql.FieldNotIn(FieldStatus, v...))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Permission {
	return predicate.Permission(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Permission {
	return predicate.Permission(sql.FieldNotNull(FieldStatus))
}

// HasOrg applies the HasEdge predicate on the "org" edge.
func HasOrg() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrgTable, OrgColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrgWith applies the HasEdge predicate on the "org" edge with a given conditions (other predicates).
func HasOrgWith(preds ...predicate.Org) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newOrgStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.OrgRole) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrgPolicy applies the HasEdge predicate on the "org_policy" edge.
func HasOrgPolicy() predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrgPolicyTable, OrgPolicyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrgPolicyWith applies the HasEdge predicate on the "org_policy" edge with a given conditions (other predicates).
func HasOrgPolicyWith(preds ...predicate.OrgPolicy) predicate.Permission {
	return predicate.Permission(func(s *sql.Selector) {
		step := newOrgPolicyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.NotPredicates(p))
}
