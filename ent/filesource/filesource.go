// Code generated by ent, DO NOT EDIT.

package filesource

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the filesource type in the database.
	Label = "file_source"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldComments holds the string denoting the comments field in the database.
	FieldComments = "comments"
	// FieldEndpoint holds the string denoting the endpoint field in the database.
	FieldEndpoint = "endpoint"
	// FieldEndpointImmutable holds the string denoting the endpoint_immutable field in the database.
	FieldEndpointImmutable = "endpoint_immutable"
	// FieldStsEndpoint holds the string denoting the sts_endpoint field in the database.
	FieldStsEndpoint = "sts_endpoint"
	// FieldRegion holds the string denoting the region field in the database.
	FieldRegion = "region"
	// FieldBucket holds the string denoting the bucket field in the database.
	FieldBucket = "bucket"
	// FieldBucketURL holds the string denoting the bucket_url field in the database.
	FieldBucketURL = "bucket_url"
	// EdgeIdentities holds the string denoting the identities edge name in mutations.
	EdgeIdentities = "identities"
	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// Table holds the table name of the filesource in the database.
	Table = "file_source"
	// IdentitiesTable is the table that holds the identities relation/edge.
	IdentitiesTable = "file_identity"
	// IdentitiesInverseTable is the table name for the FileIdentity entity.
	// It exists in this package in order to avoid circular dependency with the "fileidentity" package.
	IdentitiesInverseTable = "file_identity"
	// IdentitiesColumn is the table column denoting the identities relation/edge.
	IdentitiesColumn = "file_source_id"
	// FilesTable is the table that holds the files relation/edge.
	FilesTable = "file"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "file"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "source_id"
)

// Columns holds all SQL columns for filesource fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldCreatedAt,
	FieldUpdatedBy,
	FieldUpdatedAt,
	FieldKind,
	FieldComments,
	FieldEndpoint,
	FieldEndpointImmutable,
	FieldStsEndpoint,
	FieldRegion,
	FieldBucket,
	FieldBucketURL,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/woocoos/knockout/ent/runtime"
var (
	Hooks [2]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// EndpointValidator is a validator for the "endpoint" field. It is called by the builders before save.
	EndpointValidator func(string) error
	// DefaultEndpointImmutable holds the default value on creation for the "endpoint_immutable" field.
	DefaultEndpointImmutable bool
	// StsEndpointValidator is a validator for the "sts_endpoint" field. It is called by the builders before save.
	StsEndpointValidator func(string) error
	// RegionValidator is a validator for the "region" field. It is called by the builders before save.
	RegionValidator func(string) error
	// BucketValidator is a validator for the "bucket" field. It is called by the builders before save.
	BucketValidator func(string) error
	// BucketURLValidator is a validator for the "bucket_url" field. It is called by the builders before save.
	BucketURLValidator func(string) error
)

// Kind defines the type for the "kind" enum field.
type Kind string

// Kind values.
const (
	KindLocal  Kind = "local"
	KindMinio  Kind = "minio"
	KindAliOSS Kind = "aliOSS"
	KindAwsS3  Kind = "awsS3"
)

func (k Kind) String() string {
	return string(k)
}

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindLocal, KindMinio, KindAliOSS, KindAwsS3:
		return nil
	default:
		return fmt.Errorf("filesource: invalid enum value for kind field: %q", k)
	}
}

// OrderOption defines the ordering options for the FileSource queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByKind orders the results by the kind field.
func ByKind(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKind, opts...).ToFunc()
}

// ByComments orders the results by the comments field.
func ByComments(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComments, opts...).ToFunc()
}

// ByEndpoint orders the results by the endpoint field.
func ByEndpoint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndpoint, opts...).ToFunc()
}

// ByEndpointImmutable orders the results by the endpoint_immutable field.
func ByEndpointImmutable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndpointImmutable, opts...).ToFunc()
}

// ByStsEndpoint orders the results by the sts_endpoint field.
func ByStsEndpoint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStsEndpoint, opts...).ToFunc()
}

// ByRegion orders the results by the region field.
func ByRegion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegion, opts...).ToFunc()
}

// ByBucket orders the results by the bucket field.
func ByBucket(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBucket, opts...).ToFunc()
}

// ByBucketURL orders the results by the bucket_url field.
func ByBucketURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBucketURL, opts...).ToFunc()
}

// ByIdentitiesCount orders the results by identities count.
func ByIdentitiesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIdentitiesStep(), opts...)
	}
}

// ByIdentities orders the results by identities terms.
func ByIdentities(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIdentitiesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFilesCount orders the results by files count.
func ByFilesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFilesStep(), opts...)
	}
}

// ByFiles orders the results by files terms.
func ByFiles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFilesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newIdentitiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IdentitiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IdentitiesTable, IdentitiesColumn),
	)
}
func newFilesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FilesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Kind) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Kind) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Kind(str)
	if err := KindValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Kind", str)
	}
	return nil
}
