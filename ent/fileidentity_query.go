// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/fileidentity"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/predicate"
)

// FileIdentityQuery is the builder for querying FileIdentity entities.
type FileIdentityQuery struct {
	config
	ctx        *QueryContext
	order      []fileidentity.OrderOption
	inters     []Interceptor
	predicates []predicate.FileIdentity
	withSource *FileSourceQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*FileIdentity) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FileIdentityQuery builder.
func (fiq *FileIdentityQuery) Where(ps ...predicate.FileIdentity) *FileIdentityQuery {
	fiq.predicates = append(fiq.predicates, ps...)
	return fiq
}

// Limit the number of records to be returned by this query.
func (fiq *FileIdentityQuery) Limit(limit int) *FileIdentityQuery {
	fiq.ctx.Limit = &limit
	return fiq
}

// Offset to start from.
func (fiq *FileIdentityQuery) Offset(offset int) *FileIdentityQuery {
	fiq.ctx.Offset = &offset
	return fiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fiq *FileIdentityQuery) Unique(unique bool) *FileIdentityQuery {
	fiq.ctx.Unique = &unique
	return fiq
}

// Order specifies how the records should be ordered.
func (fiq *FileIdentityQuery) Order(o ...fileidentity.OrderOption) *FileIdentityQuery {
	fiq.order = append(fiq.order, o...)
	return fiq
}

// QuerySource chains the current query on the "source" edge.
func (fiq *FileIdentityQuery) QuerySource() *FileSourceQuery {
	query := (&FileSourceClient{config: fiq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(fileidentity.Table, fileidentity.FieldID, selector),
			sqlgraph.To(filesource.Table, filesource.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, fileidentity.SourceTable, fileidentity.SourceColumn),
		)
		fromU = sqlgraph.SetNeighbors(fiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FileIdentity entity from the query.
// Returns a *NotFoundError when no FileIdentity was found.
func (fiq *FileIdentityQuery) First(ctx context.Context) (*FileIdentity, error) {
	nodes, err := fiq.Limit(1).All(setContextOp(ctx, fiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fileidentity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fiq *FileIdentityQuery) FirstX(ctx context.Context) *FileIdentity {
	node, err := fiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FileIdentity ID from the query.
// Returns a *NotFoundError when no FileIdentity ID was found.
func (fiq *FileIdentityQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fiq.Limit(1).IDs(setContextOp(ctx, fiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fileidentity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fiq *FileIdentityQuery) FirstIDX(ctx context.Context) int {
	id, err := fiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FileIdentity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FileIdentity entity is found.
// Returns a *NotFoundError when no FileIdentity entities are found.
func (fiq *FileIdentityQuery) Only(ctx context.Context) (*FileIdentity, error) {
	nodes, err := fiq.Limit(2).All(setContextOp(ctx, fiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fileidentity.Label}
	default:
		return nil, &NotSingularError{fileidentity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fiq *FileIdentityQuery) OnlyX(ctx context.Context) *FileIdentity {
	node, err := fiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FileIdentity ID in the query.
// Returns a *NotSingularError when more than one FileIdentity ID is found.
// Returns a *NotFoundError when no entities are found.
func (fiq *FileIdentityQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fiq.Limit(2).IDs(setContextOp(ctx, fiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fileidentity.Label}
	default:
		err = &NotSingularError{fileidentity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fiq *FileIdentityQuery) OnlyIDX(ctx context.Context) int {
	id, err := fiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FileIdentities.
func (fiq *FileIdentityQuery) All(ctx context.Context) ([]*FileIdentity, error) {
	ctx = setContextOp(ctx, fiq.ctx, "All")
	if err := fiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FileIdentity, *FileIdentityQuery]()
	return withInterceptors[[]*FileIdentity](ctx, fiq, qr, fiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fiq *FileIdentityQuery) AllX(ctx context.Context) []*FileIdentity {
	nodes, err := fiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FileIdentity IDs.
func (fiq *FileIdentityQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fiq.ctx.Unique == nil && fiq.path != nil {
		fiq.Unique(true)
	}
	ctx = setContextOp(ctx, fiq.ctx, "IDs")
	if err = fiq.Select(fileidentity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fiq *FileIdentityQuery) IDsX(ctx context.Context) []int {
	ids, err := fiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fiq *FileIdentityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fiq.ctx, "Count")
	if err := fiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fiq, querierCount[*FileIdentityQuery](), fiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fiq *FileIdentityQuery) CountX(ctx context.Context) int {
	count, err := fiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fiq *FileIdentityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fiq.ctx, "Exist")
	switch _, err := fiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fiq *FileIdentityQuery) ExistX(ctx context.Context) bool {
	exist, err := fiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FileIdentityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fiq *FileIdentityQuery) Clone() *FileIdentityQuery {
	if fiq == nil {
		return nil
	}
	return &FileIdentityQuery{
		config:     fiq.config,
		ctx:        fiq.ctx.Clone(),
		order:      append([]fileidentity.OrderOption{}, fiq.order...),
		inters:     append([]Interceptor{}, fiq.inters...),
		predicates: append([]predicate.FileIdentity{}, fiq.predicates...),
		withSource: fiq.withSource.Clone(),
		// clone intermediate query.
		sql:  fiq.sql.Clone(),
		path: fiq.path,
	}
}

// WithSource tells the query-builder to eager-load the nodes that are connected to
// the "source" edge. The optional arguments are used to configure the query builder of the edge.
func (fiq *FileIdentityQuery) WithSource(opts ...func(*FileSourceQuery)) *FileIdentityQuery {
	query := (&FileSourceClient{config: fiq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fiq.withSource = query
	return fiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int `json:"created_by,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FileIdentity.Query().
//		GroupBy(fileidentity.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fiq *FileIdentityQuery) GroupBy(field string, fields ...string) *FileIdentityGroupBy {
	fiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FileIdentityGroupBy{build: fiq}
	grbuild.flds = &fiq.ctx.Fields
	grbuild.label = fileidentity.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int `json:"created_by,omitempty"`
//	}
//
//	client.FileIdentity.Query().
//		Select(fileidentity.FieldCreatedBy).
//		Scan(ctx, &v)
func (fiq *FileIdentityQuery) Select(fields ...string) *FileIdentitySelect {
	fiq.ctx.Fields = append(fiq.ctx.Fields, fields...)
	sbuild := &FileIdentitySelect{FileIdentityQuery: fiq}
	sbuild.label = fileidentity.Label
	sbuild.flds, sbuild.scan = &fiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FileIdentitySelect configured with the given aggregations.
func (fiq *FileIdentityQuery) Aggregate(fns ...AggregateFunc) *FileIdentitySelect {
	return fiq.Select().Aggregate(fns...)
}

func (fiq *FileIdentityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fiq); err != nil {
				return err
			}
		}
	}
	for _, f := range fiq.ctx.Fields {
		if !fileidentity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fiq.path != nil {
		prev, err := fiq.path(ctx)
		if err != nil {
			return err
		}
		fiq.sql = prev
	}
	return nil
}

func (fiq *FileIdentityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FileIdentity, error) {
	var (
		nodes       = []*FileIdentity{}
		_spec       = fiq.querySpec()
		loadedTypes = [1]bool{
			fiq.withSource != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FileIdentity).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FileIdentity{config: fiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(fiq.modifiers) > 0 {
		_spec.Modifiers = fiq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fiq.withSource; query != nil {
		if err := fiq.loadSource(ctx, query, nodes, nil,
			func(n *FileIdentity, e *FileSource) { n.Edges.Source = e }); err != nil {
			return nil, err
		}
	}
	for i := range fiq.loadTotal {
		if err := fiq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fiq *FileIdentityQuery) loadSource(ctx context.Context, query *FileSourceQuery, nodes []*FileIdentity, init func(*FileIdentity), assign func(*FileIdentity, *FileSource)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*FileIdentity)
	for i := range nodes {
		fk := nodes[i].FileSourceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(filesource.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "file_source_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fiq *FileIdentityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fiq.querySpec()
	if len(fiq.modifiers) > 0 {
		_spec.Modifiers = fiq.modifiers
	}
	_spec.Node.Columns = fiq.ctx.Fields
	if len(fiq.ctx.Fields) > 0 {
		_spec.Unique = fiq.ctx.Unique != nil && *fiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fiq.driver, _spec)
}

func (fiq *FileIdentityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(fileidentity.Table, fileidentity.Columns, sqlgraph.NewFieldSpec(fileidentity.FieldID, field.TypeInt))
	_spec.From = fiq.sql
	if unique := fiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fiq.path != nil {
		_spec.Unique = true
	}
	if fields := fiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fileidentity.FieldID)
		for i := range fields {
			if fields[i] != fileidentity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if fiq.withSource != nil {
			_spec.Node.AddColumnOnce(fileidentity.FieldFileSourceID)
		}
	}
	if ps := fiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fiq *FileIdentityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fiq.driver.Dialect())
	t1 := builder.Table(fileidentity.Table)
	columns := fiq.ctx.Fields
	if len(columns) == 0 {
		columns = fileidentity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fiq.sql != nil {
		selector = fiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fiq.ctx.Unique != nil && *fiq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fiq.predicates {
		p(selector)
	}
	for _, p := range fiq.order {
		p(selector)
	}
	if offset := fiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FileIdentityGroupBy is the group-by builder for FileIdentity entities.
type FileIdentityGroupBy struct {
	selector
	build *FileIdentityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (figb *FileIdentityGroupBy) Aggregate(fns ...AggregateFunc) *FileIdentityGroupBy {
	figb.fns = append(figb.fns, fns...)
	return figb
}

// Scan applies the selector query and scans the result into the given value.
func (figb *FileIdentityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, figb.build.ctx, "GroupBy")
	if err := figb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileIdentityQuery, *FileIdentityGroupBy](ctx, figb.build, figb, figb.build.inters, v)
}

func (figb *FileIdentityGroupBy) sqlScan(ctx context.Context, root *FileIdentityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(figb.fns))
	for _, fn := range figb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*figb.flds)+len(figb.fns))
		for _, f := range *figb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*figb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := figb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FileIdentitySelect is the builder for selecting fields of FileIdentity entities.
type FileIdentitySelect struct {
	*FileIdentityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fis *FileIdentitySelect) Aggregate(fns ...AggregateFunc) *FileIdentitySelect {
	fis.fns = append(fis.fns, fns...)
	return fis
}

// Scan applies the selector query and scans the result into the given value.
func (fis *FileIdentitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fis.ctx, "Select")
	if err := fis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileIdentityQuery, *FileIdentitySelect](ctx, fis.FileIdentityQuery, fis, fis.inters, v)
}

func (fis *FileIdentitySelect) sqlScan(ctx context.Context, root *FileIdentityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fis.fns))
	for _, fn := range fis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
