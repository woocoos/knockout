// Code generated by ent, DO NOT EDIT.

package filesource

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/woocoos/knockout/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FileSource {
	return predicate.FileSource(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FileSource {
	return predicate.FileSource(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FileSource {
	return predicate.FileSource(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FileSource {
	return predicate.FileSource(sql.FieldLTE(FieldID, id))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldUpdatedAt, v))
}

// Comments applies equality check predicate on the "comments" field. It's identical to CommentsEQ.
func Comments(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldComments, v))
}

// Endpoint applies equality check predicate on the "endpoint" field. It's identical to EndpointEQ.
func Endpoint(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldEndpoint, v))
}

// StsEndpoint applies equality check predicate on the "sts_endpoint" field. It's identical to StsEndpointEQ.
func StsEndpoint(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldStsEndpoint, v))
}

// Region applies equality check predicate on the "region" field. It's identical to RegionEQ.
func Region(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldRegion, v))
}

// Bucket applies equality check predicate on the "bucket" field. It's identical to BucketEQ.
func Bucket(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldBucket, v))
}

// BucketUrl applies equality check predicate on the "bucketUrl" field. It's identical to BucketUrlEQ.
func BucketUrl(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldBucketUrl, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int) predicate.FileSource {
	return predicate.FileSource(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.FileSource {
	return predicate.FileSource(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.FileSource {
	return predicate.FileSource(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.FileSource {
	return predicate.FileSource(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.FileSource {
	return predicate.FileSource(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.FileSource {
	return predicate.FileSource(sql.FieldNotNull(FieldUpdatedAt))
}

// KindEQ applies the EQ predicate on the "kind" field.
func KindEQ(v Kind) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldKind, v))
}

// KindNEQ applies the NEQ predicate on the "kind" field.
func KindNEQ(v Kind) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldKind, v))
}

// KindIn applies the In predicate on the "kind" field.
func KindIn(vs ...Kind) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldKind, vs...))
}

// KindNotIn applies the NotIn predicate on the "kind" field.
func KindNotIn(vs ...Kind) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldKind, vs...))
}

// CommentsEQ applies the EQ predicate on the "comments" field.
func CommentsEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldComments, v))
}

// CommentsNEQ applies the NEQ predicate on the "comments" field.
func CommentsNEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldComments, v))
}

// CommentsIn applies the In predicate on the "comments" field.
func CommentsIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldComments, vs...))
}

// CommentsNotIn applies the NotIn predicate on the "comments" field.
func CommentsNotIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldComments, vs...))
}

// CommentsGT applies the GT predicate on the "comments" field.
func CommentsGT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGT(FieldComments, v))
}

// CommentsGTE applies the GTE predicate on the "comments" field.
func CommentsGTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGTE(FieldComments, v))
}

// CommentsLT applies the LT predicate on the "comments" field.
func CommentsLT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLT(FieldComments, v))
}

// CommentsLTE applies the LTE predicate on the "comments" field.
func CommentsLTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLTE(FieldComments, v))
}

// CommentsContains applies the Contains predicate on the "comments" field.
func CommentsContains(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContains(FieldComments, v))
}

// CommentsHasPrefix applies the HasPrefix predicate on the "comments" field.
func CommentsHasPrefix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasPrefix(FieldComments, v))
}

// CommentsHasSuffix applies the HasSuffix predicate on the "comments" field.
func CommentsHasSuffix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasSuffix(FieldComments, v))
}

// CommentsIsNil applies the IsNil predicate on the "comments" field.
func CommentsIsNil() predicate.FileSource {
	return predicate.FileSource(sql.FieldIsNull(FieldComments))
}

// CommentsNotNil applies the NotNil predicate on the "comments" field.
func CommentsNotNil() predicate.FileSource {
	return predicate.FileSource(sql.FieldNotNull(FieldComments))
}

// CommentsEqualFold applies the EqualFold predicate on the "comments" field.
func CommentsEqualFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEqualFold(FieldComments, v))
}

// CommentsContainsFold applies the ContainsFold predicate on the "comments" field.
func CommentsContainsFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContainsFold(FieldComments, v))
}

// EndpointEQ applies the EQ predicate on the "endpoint" field.
func EndpointEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldEndpoint, v))
}

// EndpointNEQ applies the NEQ predicate on the "endpoint" field.
func EndpointNEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldEndpoint, v))
}

// EndpointIn applies the In predicate on the "endpoint" field.
func EndpointIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldEndpoint, vs...))
}

// EndpointNotIn applies the NotIn predicate on the "endpoint" field.
func EndpointNotIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldEndpoint, vs...))
}

// EndpointGT applies the GT predicate on the "endpoint" field.
func EndpointGT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGT(FieldEndpoint, v))
}

// EndpointGTE applies the GTE predicate on the "endpoint" field.
func EndpointGTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGTE(FieldEndpoint, v))
}

// EndpointLT applies the LT predicate on the "endpoint" field.
func EndpointLT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLT(FieldEndpoint, v))
}

// EndpointLTE applies the LTE predicate on the "endpoint" field.
func EndpointLTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLTE(FieldEndpoint, v))
}

// EndpointContains applies the Contains predicate on the "endpoint" field.
func EndpointContains(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContains(FieldEndpoint, v))
}

// EndpointHasPrefix applies the HasPrefix predicate on the "endpoint" field.
func EndpointHasPrefix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasPrefix(FieldEndpoint, v))
}

// EndpointHasSuffix applies the HasSuffix predicate on the "endpoint" field.
func EndpointHasSuffix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasSuffix(FieldEndpoint, v))
}

// EndpointEqualFold applies the EqualFold predicate on the "endpoint" field.
func EndpointEqualFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEqualFold(FieldEndpoint, v))
}

// EndpointContainsFold applies the ContainsFold predicate on the "endpoint" field.
func EndpointContainsFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContainsFold(FieldEndpoint, v))
}

// StsEndpointEQ applies the EQ predicate on the "sts_endpoint" field.
func StsEndpointEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldStsEndpoint, v))
}

// StsEndpointNEQ applies the NEQ predicate on the "sts_endpoint" field.
func StsEndpointNEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldStsEndpoint, v))
}

// StsEndpointIn applies the In predicate on the "sts_endpoint" field.
func StsEndpointIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldStsEndpoint, vs...))
}

// StsEndpointNotIn applies the NotIn predicate on the "sts_endpoint" field.
func StsEndpointNotIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldStsEndpoint, vs...))
}

// StsEndpointGT applies the GT predicate on the "sts_endpoint" field.
func StsEndpointGT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGT(FieldStsEndpoint, v))
}

// StsEndpointGTE applies the GTE predicate on the "sts_endpoint" field.
func StsEndpointGTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGTE(FieldStsEndpoint, v))
}

// StsEndpointLT applies the LT predicate on the "sts_endpoint" field.
func StsEndpointLT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLT(FieldStsEndpoint, v))
}

// StsEndpointLTE applies the LTE predicate on the "sts_endpoint" field.
func StsEndpointLTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLTE(FieldStsEndpoint, v))
}

// StsEndpointContains applies the Contains predicate on the "sts_endpoint" field.
func StsEndpointContains(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContains(FieldStsEndpoint, v))
}

// StsEndpointHasPrefix applies the HasPrefix predicate on the "sts_endpoint" field.
func StsEndpointHasPrefix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasPrefix(FieldStsEndpoint, v))
}

// StsEndpointHasSuffix applies the HasSuffix predicate on the "sts_endpoint" field.
func StsEndpointHasSuffix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasSuffix(FieldStsEndpoint, v))
}

// StsEndpointEqualFold applies the EqualFold predicate on the "sts_endpoint" field.
func StsEndpointEqualFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEqualFold(FieldStsEndpoint, v))
}

// StsEndpointContainsFold applies the ContainsFold predicate on the "sts_endpoint" field.
func StsEndpointContainsFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContainsFold(FieldStsEndpoint, v))
}

// RegionEQ applies the EQ predicate on the "region" field.
func RegionEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldRegion, v))
}

// RegionNEQ applies the NEQ predicate on the "region" field.
func RegionNEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldRegion, v))
}

// RegionIn applies the In predicate on the "region" field.
func RegionIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldRegion, vs...))
}

// RegionNotIn applies the NotIn predicate on the "region" field.
func RegionNotIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldRegion, vs...))
}

// RegionGT applies the GT predicate on the "region" field.
func RegionGT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGT(FieldRegion, v))
}

// RegionGTE applies the GTE predicate on the "region" field.
func RegionGTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGTE(FieldRegion, v))
}

// RegionLT applies the LT predicate on the "region" field.
func RegionLT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLT(FieldRegion, v))
}

// RegionLTE applies the LTE predicate on the "region" field.
func RegionLTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLTE(FieldRegion, v))
}

// RegionContains applies the Contains predicate on the "region" field.
func RegionContains(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContains(FieldRegion, v))
}

// RegionHasPrefix applies the HasPrefix predicate on the "region" field.
func RegionHasPrefix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasPrefix(FieldRegion, v))
}

// RegionHasSuffix applies the HasSuffix predicate on the "region" field.
func RegionHasSuffix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasSuffix(FieldRegion, v))
}

// RegionEqualFold applies the EqualFold predicate on the "region" field.
func RegionEqualFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEqualFold(FieldRegion, v))
}

// RegionContainsFold applies the ContainsFold predicate on the "region" field.
func RegionContainsFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContainsFold(FieldRegion, v))
}

// BucketEQ applies the EQ predicate on the "bucket" field.
func BucketEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldBucket, v))
}

// BucketNEQ applies the NEQ predicate on the "bucket" field.
func BucketNEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldBucket, v))
}

// BucketIn applies the In predicate on the "bucket" field.
func BucketIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldBucket, vs...))
}

// BucketNotIn applies the NotIn predicate on the "bucket" field.
func BucketNotIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldBucket, vs...))
}

// BucketGT applies the GT predicate on the "bucket" field.
func BucketGT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGT(FieldBucket, v))
}

// BucketGTE applies the GTE predicate on the "bucket" field.
func BucketGTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGTE(FieldBucket, v))
}

// BucketLT applies the LT predicate on the "bucket" field.
func BucketLT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLT(FieldBucket, v))
}

// BucketLTE applies the LTE predicate on the "bucket" field.
func BucketLTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLTE(FieldBucket, v))
}

// BucketContains applies the Contains predicate on the "bucket" field.
func BucketContains(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContains(FieldBucket, v))
}

// BucketHasPrefix applies the HasPrefix predicate on the "bucket" field.
func BucketHasPrefix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasPrefix(FieldBucket, v))
}

// BucketHasSuffix applies the HasSuffix predicate on the "bucket" field.
func BucketHasSuffix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasSuffix(FieldBucket, v))
}

// BucketEqualFold applies the EqualFold predicate on the "bucket" field.
func BucketEqualFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEqualFold(FieldBucket, v))
}

// BucketContainsFold applies the ContainsFold predicate on the "bucket" field.
func BucketContainsFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContainsFold(FieldBucket, v))
}

// BucketUrlEQ applies the EQ predicate on the "bucketUrl" field.
func BucketUrlEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEQ(FieldBucketUrl, v))
}

// BucketUrlNEQ applies the NEQ predicate on the "bucketUrl" field.
func BucketUrlNEQ(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNEQ(FieldBucketUrl, v))
}

// BucketUrlIn applies the In predicate on the "bucketUrl" field.
func BucketUrlIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldIn(FieldBucketUrl, vs...))
}

// BucketUrlNotIn applies the NotIn predicate on the "bucketUrl" field.
func BucketUrlNotIn(vs ...string) predicate.FileSource {
	return predicate.FileSource(sql.FieldNotIn(FieldBucketUrl, vs...))
}

// BucketUrlGT applies the GT predicate on the "bucketUrl" field.
func BucketUrlGT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGT(FieldBucketUrl, v))
}

// BucketUrlGTE applies the GTE predicate on the "bucketUrl" field.
func BucketUrlGTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldGTE(FieldBucketUrl, v))
}

// BucketUrlLT applies the LT predicate on the "bucketUrl" field.
func BucketUrlLT(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLT(FieldBucketUrl, v))
}

// BucketUrlLTE applies the LTE predicate on the "bucketUrl" field.
func BucketUrlLTE(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldLTE(FieldBucketUrl, v))
}

// BucketUrlContains applies the Contains predicate on the "bucketUrl" field.
func BucketUrlContains(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContains(FieldBucketUrl, v))
}

// BucketUrlHasPrefix applies the HasPrefix predicate on the "bucketUrl" field.
func BucketUrlHasPrefix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasPrefix(FieldBucketUrl, v))
}

// BucketUrlHasSuffix applies the HasSuffix predicate on the "bucketUrl" field.
func BucketUrlHasSuffix(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldHasSuffix(FieldBucketUrl, v))
}

// BucketUrlIsNil applies the IsNil predicate on the "bucketUrl" field.
func BucketUrlIsNil() predicate.FileSource {
	return predicate.FileSource(sql.FieldIsNull(FieldBucketUrl))
}

// BucketUrlNotNil applies the NotNil predicate on the "bucketUrl" field.
func BucketUrlNotNil() predicate.FileSource {
	return predicate.FileSource(sql.FieldNotNull(FieldBucketUrl))
}

// BucketUrlEqualFold applies the EqualFold predicate on the "bucketUrl" field.
func BucketUrlEqualFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldEqualFold(FieldBucketUrl, v))
}

// BucketUrlContainsFold applies the ContainsFold predicate on the "bucketUrl" field.
func BucketUrlContainsFold(v string) predicate.FileSource {
	return predicate.FileSource(sql.FieldContainsFold(FieldBucketUrl, v))
}

// HasIdentities applies the HasEdge predicate on the "identities" edge.
func HasIdentities() predicate.FileSource {
	return predicate.FileSource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IdentitiesTable, IdentitiesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIdentitiesWith applies the HasEdge predicate on the "identities" edge with a given conditions (other predicates).
func HasIdentitiesWith(preds ...predicate.FileIdentity) predicate.FileSource {
	return predicate.FileSource(func(s *sql.Selector) {
		step := newIdentitiesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.FileSource {
	return predicate.FileSource(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.FileSource {
	return predicate.FileSource(func(s *sql.Selector) {
		step := newFilesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FileSource) predicate.FileSource {
	return predicate.FileSource(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FileSource) predicate.FileSource {
	return predicate.FileSource(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FileSource) predicate.FileSource {
	return predicate.FileSource(sql.NotPredicates(p))
}
