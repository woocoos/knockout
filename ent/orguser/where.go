// Code generated by ent, DO NOT EDIT.

package orguser

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/knockout/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrgID applies equality check predicate on the "org_id" field. It's identical to OrgIDEQ.
func OrgID(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldOrgID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldUserID, v))
}

// JoinedAt applies equality check predicate on the "joined_at" field. It's identical to JoinedAtEQ.
func JoinedAt(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldJoinedAt, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldDisplayName, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.OrgUser {
	return predicate.OrgUser(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.OrgUser {
	return predicate.OrgUser(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNotNull(FieldUpdatedAt))
}

// OrgIDEQ applies the EQ predicate on the "org_id" field.
func OrgIDEQ(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldOrgID, v))
}

// OrgIDNEQ applies the NEQ predicate on the "org_id" field.
func OrgIDNEQ(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNEQ(FieldOrgID, v))
}

// OrgIDIn applies the In predicate on the "org_id" field.
func OrgIDIn(vs ...int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldIn(FieldOrgID, vs...))
}

// OrgIDNotIn applies the NotIn predicate on the "org_id" field.
func OrgIDNotIn(vs ...int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNotIn(FieldOrgID, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNotIn(FieldUserID, vs...))
}

// JoinedAtEQ applies the EQ predicate on the "joined_at" field.
func JoinedAtEQ(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldJoinedAt, v))
}

// JoinedAtNEQ applies the NEQ predicate on the "joined_at" field.
func JoinedAtNEQ(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNEQ(FieldJoinedAt, v))
}

// JoinedAtIn applies the In predicate on the "joined_at" field.
func JoinedAtIn(vs ...time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldIn(FieldJoinedAt, vs...))
}

// JoinedAtNotIn applies the NotIn predicate on the "joined_at" field.
func JoinedAtNotIn(vs ...time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNotIn(FieldJoinedAt, vs...))
}

// JoinedAtGT applies the GT predicate on the "joined_at" field.
func JoinedAtGT(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGT(FieldJoinedAt, v))
}

// JoinedAtGTE applies the GTE predicate on the "joined_at" field.
func JoinedAtGTE(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGTE(FieldJoinedAt, v))
}

// JoinedAtLT applies the LT predicate on the "joined_at" field.
func JoinedAtLT(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLT(FieldJoinedAt, v))
}

// JoinedAtLTE applies the LTE predicate on the "joined_at" field.
func JoinedAtLTE(v time.Time) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLTE(FieldJoinedAt, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.OrgUser {
	return predicate.OrgUser(sql.FieldContainsFold(FieldDisplayName, v))
}

// HasOrg applies the HasEdge predicate on the "org" edge.
func HasOrg() predicate.OrgUser {
	return predicate.OrgUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrgTable, OrgColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrgWith applies the HasEdge predicate on the "org" edge with a given conditions (other predicates).
func HasOrgWith(preds ...predicate.Org) predicate.OrgUser {
	return predicate.OrgUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrgInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrgTable, OrgColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.OrgUser {
	return predicate.OrgUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.OrgUser {
	return predicate.OrgUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrgRoles applies the HasEdge predicate on the "org_roles" edge.
func HasOrgRoles() predicate.OrgUser {
	return predicate.OrgUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OrgRolesTable, OrgRolesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrgRolesWith applies the HasEdge predicate on the "org_roles" edge with a given conditions (other predicates).
func HasOrgRolesWith(preds ...predicate.OrgRole) predicate.OrgUser {
	return predicate.OrgUser(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrgRolesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OrgRolesTable, OrgRolesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrgUser) predicate.OrgUser {
	return predicate.OrgUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrgUser) predicate.OrgUser {
	return predicate.OrgUser(func(s *sql.Selector) {
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
func Not(p predicate.OrgUser) predicate.OrgUser {
	return predicate.OrgUser(func(s *sql.Selector) {
		p(s.Not())
	})
}
