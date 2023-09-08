// Code generated by ent, DO NOT EDIT.

package orguserpreference

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/knockout/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldUserID, v))
}

// OrgID applies equality check predicate on the "org_id" field. It's identical to OrgIDEQ.
func OrgID(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldOrgID, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNotNull(FieldUpdatedAt))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNotIn(FieldUserID, vs...))
}

// OrgIDEQ applies the EQ predicate on the "org_id" field.
func OrgIDEQ(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldEQ(FieldOrgID, v))
}

// OrgIDNEQ applies the NEQ predicate on the "org_id" field.
func OrgIDNEQ(v int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNEQ(FieldOrgID, v))
}

// OrgIDIn applies the In predicate on the "org_id" field.
func OrgIDIn(vs ...int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldIn(FieldOrgID, vs...))
}

// OrgIDNotIn applies the NotIn predicate on the "org_id" field.
func OrgIDNotIn(vs ...int) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNotIn(FieldOrgID, vs...))
}

// MenuFavoriteIsNil applies the IsNil predicate on the "menu_favorite" field.
func MenuFavoriteIsNil() predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldIsNull(FieldMenuFavorite))
}

// MenuFavoriteNotNil applies the NotNil predicate on the "menu_favorite" field.
func MenuFavoriteNotNil() predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNotNull(FieldMenuFavorite))
}

// MenuRecentIsNil applies the IsNil predicate on the "menu_recent" field.
func MenuRecentIsNil() predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldIsNull(FieldMenuRecent))
}

// MenuRecentNotNil applies the NotNil predicate on the "menu_recent" field.
func MenuRecentNotNil() predicate.OrgUserPreference {
	return predicate.OrgUserPreference(sql.FieldNotNull(FieldMenuRecent))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.OrgUserPreference {
	return predicate.OrgUserPreference(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrg applies the HasEdge predicate on the "org" edge.
func HasOrg() predicate.OrgUserPreference {
	return predicate.OrgUserPreference(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrgTable, OrgColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrgWith applies the HasEdge predicate on the "org" edge with a given conditions (other predicates).
func HasOrgWith(preds ...predicate.Org) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(func(s *sql.Selector) {
		step := newOrgStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrgUserPreference) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrgUserPreference) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(func(s *sql.Selector) {
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
func Not(p predicate.OrgUserPreference) predicate.OrgUserPreference {
	return predicate.OrgUserPreference(func(s *sql.Selector) {
		p(s.Not())
	})
}
