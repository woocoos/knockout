// Code generated by ent, DO NOT EDIT.

package organizationrole

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/knockout/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrgID applies equality check predicate on the "org_id" field. It's identical to OrgIDEQ.
func OrgID(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldOrgID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldName, v))
}

// AppRoleID applies equality check predicate on the "app_role_id" field. It's identical to AppRoleIDEQ.
func AppRoleID(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldAppRoleID, v))
}

// Comments applies equality check predicate on the "comments" field. It's identical to CommentsEQ.
func Comments(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldComments, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotNull(FieldUpdatedAt))
}

// OrgIDEQ applies the EQ predicate on the "org_id" field.
func OrgIDEQ(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldOrgID, v))
}

// OrgIDNEQ applies the NEQ predicate on the "org_id" field.
func OrgIDNEQ(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNEQ(FieldOrgID, v))
}

// OrgIDIn applies the In predicate on the "org_id" field.
func OrgIDIn(vs ...int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIn(FieldOrgID, vs...))
}

// OrgIDNotIn applies the NotIn predicate on the "org_id" field.
func OrgIDNotIn(vs ...int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotIn(FieldOrgID, vs...))
}

// KindEQ applies the EQ predicate on the "kind" field.
func KindEQ(v Kind) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldKind, v))
}

// KindNEQ applies the NEQ predicate on the "kind" field.
func KindNEQ(v Kind) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNEQ(FieldKind, v))
}

// KindIn applies the In predicate on the "kind" field.
func KindIn(vs ...Kind) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIn(FieldKind, vs...))
}

// KindNotIn applies the NotIn predicate on the "kind" field.
func KindNotIn(vs ...Kind) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotIn(FieldKind, vs...))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldContainsFold(FieldName, v))
}

// AppRoleIDEQ applies the EQ predicate on the "app_role_id" field.
func AppRoleIDEQ(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldAppRoleID, v))
}

// AppRoleIDNEQ applies the NEQ predicate on the "app_role_id" field.
func AppRoleIDNEQ(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNEQ(FieldAppRoleID, v))
}

// AppRoleIDIn applies the In predicate on the "app_role_id" field.
func AppRoleIDIn(vs ...int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIn(FieldAppRoleID, vs...))
}

// AppRoleIDNotIn applies the NotIn predicate on the "app_role_id" field.
func AppRoleIDNotIn(vs ...int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotIn(FieldAppRoleID, vs...))
}

// AppRoleIDGT applies the GT predicate on the "app_role_id" field.
func AppRoleIDGT(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGT(FieldAppRoleID, v))
}

// AppRoleIDGTE applies the GTE predicate on the "app_role_id" field.
func AppRoleIDGTE(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGTE(FieldAppRoleID, v))
}

// AppRoleIDLT applies the LT predicate on the "app_role_id" field.
func AppRoleIDLT(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLT(FieldAppRoleID, v))
}

// AppRoleIDLTE applies the LTE predicate on the "app_role_id" field.
func AppRoleIDLTE(v int) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLTE(FieldAppRoleID, v))
}

// AppRoleIDIsNil applies the IsNil predicate on the "app_role_id" field.
func AppRoleIDIsNil() predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIsNull(FieldAppRoleID))
}

// AppRoleIDNotNil applies the NotNil predicate on the "app_role_id" field.
func AppRoleIDNotNil() predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotNull(FieldAppRoleID))
}

// CommentsEQ applies the EQ predicate on the "comments" field.
func CommentsEQ(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEQ(FieldComments, v))
}

// CommentsNEQ applies the NEQ predicate on the "comments" field.
func CommentsNEQ(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNEQ(FieldComments, v))
}

// CommentsIn applies the In predicate on the "comments" field.
func CommentsIn(vs ...string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIn(FieldComments, vs...))
}

// CommentsNotIn applies the NotIn predicate on the "comments" field.
func CommentsNotIn(vs ...string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotIn(FieldComments, vs...))
}

// CommentsGT applies the GT predicate on the "comments" field.
func CommentsGT(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGT(FieldComments, v))
}

// CommentsGTE applies the GTE predicate on the "comments" field.
func CommentsGTE(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldGTE(FieldComments, v))
}

// CommentsLT applies the LT predicate on the "comments" field.
func CommentsLT(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLT(FieldComments, v))
}

// CommentsLTE applies the LTE predicate on the "comments" field.
func CommentsLTE(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldLTE(FieldComments, v))
}

// CommentsContains applies the Contains predicate on the "comments" field.
func CommentsContains(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldContains(FieldComments, v))
}

// CommentsHasPrefix applies the HasPrefix predicate on the "comments" field.
func CommentsHasPrefix(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldHasPrefix(FieldComments, v))
}

// CommentsHasSuffix applies the HasSuffix predicate on the "comments" field.
func CommentsHasSuffix(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldHasSuffix(FieldComments, v))
}

// CommentsIsNil applies the IsNil predicate on the "comments" field.
func CommentsIsNil() predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldIsNull(FieldComments))
}

// CommentsNotNil applies the NotNil predicate on the "comments" field.
func CommentsNotNil() predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldNotNull(FieldComments))
}

// CommentsEqualFold applies the EqualFold predicate on the "comments" field.
func CommentsEqualFold(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldEqualFold(FieldComments, v))
}

// CommentsContainsFold applies the ContainsFold predicate on the "comments" field.
func CommentsContainsFold(v string) predicate.OrganizationRole {
	return predicate.OrganizationRole(sql.FieldContainsFold(FieldComments, v))
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.OrganizationRole {
	return predicate.OrganizationRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.OrganizationRole {
	return predicate.OrganizationRole(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OrganizationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrganizationRole) predicate.OrganizationRole {
	return predicate.OrganizationRole(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrganizationRole) predicate.OrganizationRole {
	return predicate.OrganizationRole(func(s *sql.Selector) {
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
func Not(p predicate.OrganizationRole) predicate.OrganizationRole {
	return predicate.OrganizationRole(func(s *sql.Selector) {
		p(s.Not())
	})
}
