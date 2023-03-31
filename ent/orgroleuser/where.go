// Code generated by ent, DO NOT EDIT.

package orgroleuser

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/knockout/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrgRoleID applies equality check predicate on the "org_role_id" field. It's identical to OrgRoleIDEQ.
func OrgRoleID(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldOrgRoleID, v))
}

// OrgUserID applies equality check predicate on the "org_user_id" field. It's identical to OrgUserIDEQ.
func OrgUserID(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldOrgUserID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotNull(FieldUpdatedAt))
}

// OrgRoleIDEQ applies the EQ predicate on the "org_role_id" field.
func OrgRoleIDEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldOrgRoleID, v))
}

// OrgRoleIDNEQ applies the NEQ predicate on the "org_role_id" field.
func OrgRoleIDNEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldOrgRoleID, v))
}

// OrgRoleIDIn applies the In predicate on the "org_role_id" field.
func OrgRoleIDIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldOrgRoleID, vs...))
}

// OrgRoleIDNotIn applies the NotIn predicate on the "org_role_id" field.
func OrgRoleIDNotIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldOrgRoleID, vs...))
}

// OrgUserIDEQ applies the EQ predicate on the "org_user_id" field.
func OrgUserIDEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldEQ(FieldOrgUserID, v))
}

// OrgUserIDNEQ applies the NEQ predicate on the "org_user_id" field.
func OrgUserIDNEQ(v int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNEQ(FieldOrgUserID, v))
}

// OrgUserIDIn applies the In predicate on the "org_user_id" field.
func OrgUserIDIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldIn(FieldOrgUserID, vs...))
}

// OrgUserIDNotIn applies the NotIn predicate on the "org_user_id" field.
func OrgUserIDNotIn(vs ...int) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(sql.FieldNotIn(FieldOrgUserID, vs...))
}

// HasOrgRole applies the HasEdge predicate on the "org_role" edge.
func HasOrgRole() predicate.OrgRoleUser {
	return predicate.OrgRoleUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrgRoleTable, OrgRoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrgRoleWith applies the HasEdge predicate on the "org_role" edge with a given conditions (other predicates).
func HasOrgRoleWith(preds ...predicate.OrgRole) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrgRoleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrgRoleTable, OrgRoleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrgUser applies the HasEdge predicate on the "org_user" edge.
func HasOrgUser() predicate.OrgRoleUser {
	return predicate.OrgRoleUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrgUserTable, OrgUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrgUserWith applies the HasEdge predicate on the "org_user" edge with a given conditions (other predicates).
func HasOrgUserWith(preds ...predicate.OrgUser) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrgUserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrgUserTable, OrgUserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrgRoleUser) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrgRoleUser) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(func(s *sql.Selector) {
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
func Not(p predicate.OrgRoleUser) predicate.OrgRoleUser {
	return predicate.OrgRoleUser(func(s *sql.Selector) {
		p(s.Not())
	})
}
