// Code generated by ent, DO NOT EDIT.

package approlepolicy

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/knockout/ent/predicate"
)

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldUpdatedAt, v))
}

// AppRoleID applies equality check predicate on the "app_role_id" field. It's identical to AppRoleIDEQ.
func AppRoleID(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldAppRoleID, v))
}

// AppPolicyID applies equality check predicate on the "app_policy_id" field. It's identical to AppPolicyIDEQ.
func AppPolicyID(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldAppPolicyID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNotNull(FieldUpdatedAt))
}

// AppRoleIDEQ applies the EQ predicate on the "app_role_id" field.
func AppRoleIDEQ(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldAppRoleID, v))
}

// AppRoleIDNEQ applies the NEQ predicate on the "app_role_id" field.
func AppRoleIDNEQ(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNEQ(FieldAppRoleID, v))
}

// AppRoleIDIn applies the In predicate on the "app_role_id" field.
func AppRoleIDIn(vs ...int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldIn(FieldAppRoleID, vs...))
}

// AppRoleIDNotIn applies the NotIn predicate on the "app_role_id" field.
func AppRoleIDNotIn(vs ...int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNotIn(FieldAppRoleID, vs...))
}

// AppPolicyIDEQ applies the EQ predicate on the "app_policy_id" field.
func AppPolicyIDEQ(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldEQ(FieldAppPolicyID, v))
}

// AppPolicyIDNEQ applies the NEQ predicate on the "app_policy_id" field.
func AppPolicyIDNEQ(v int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNEQ(FieldAppPolicyID, v))
}

// AppPolicyIDIn applies the In predicate on the "app_policy_id" field.
func AppPolicyIDIn(vs ...int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldIn(FieldAppPolicyID, vs...))
}

// AppPolicyIDNotIn applies the NotIn predicate on the "app_policy_id" field.
func AppPolicyIDNotIn(vs ...int) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(sql.FieldNotIn(FieldAppPolicyID, vs...))
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.AppRolePolicy {
	return predicate.AppRolePolicy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, RoleColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.AppRole) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, RoleColumn),
			sqlgraph.To(RoleInverseTable, AppRoleFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPolicy applies the HasEdge predicate on the "policy" edge.
func HasPolicy() predicate.AppRolePolicy {
	return predicate.AppRolePolicy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, PolicyColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, PolicyTable, PolicyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPolicyWith applies the HasEdge predicate on the "policy" edge with a given conditions (other predicates).
func HasPolicyWith(preds ...predicate.AppPolicy) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, PolicyColumn),
			sqlgraph.To(PolicyInverseTable, AppPolicyFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PolicyTable, PolicyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppRolePolicy) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppRolePolicy) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(func(s *sql.Selector) {
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
func Not(p predicate.AppRolePolicy) predicate.AppRolePolicy {
	return predicate.AppRolePolicy(func(s *sql.Selector) {
		p(s.Not())
	})
}
