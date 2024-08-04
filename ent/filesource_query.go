// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/file"
	"github.com/woocoos/knockout/ent/fileidentity"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/predicate"
)

// FileSourceQuery is the builder for querying FileSource entities.
type FileSourceQuery struct {
	config
	ctx                 *QueryContext
	order               []filesource.OrderOption
	inters              []Interceptor
	predicates          []predicate.FileSource
	withIdentities      *FileIdentityQuery
	withFiles           *FileQuery
	modifiers           []func(*sql.Selector)
	loadTotal           []func(context.Context, []*FileSource) error
	withNamedIdentities map[string]*FileIdentityQuery
	withNamedFiles      map[string]*FileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FileSourceQuery builder.
func (fsq *FileSourceQuery) Where(ps ...predicate.FileSource) *FileSourceQuery {
	fsq.predicates = append(fsq.predicates, ps...)
	return fsq
}

// Limit the number of records to be returned by this query.
func (fsq *FileSourceQuery) Limit(limit int) *FileSourceQuery {
	fsq.ctx.Limit = &limit
	return fsq
}

// Offset to start from.
func (fsq *FileSourceQuery) Offset(offset int) *FileSourceQuery {
	fsq.ctx.Offset = &offset
	return fsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fsq *FileSourceQuery) Unique(unique bool) *FileSourceQuery {
	fsq.ctx.Unique = &unique
	return fsq
}

// Order specifies how the records should be ordered.
func (fsq *FileSourceQuery) Order(o ...filesource.OrderOption) *FileSourceQuery {
	fsq.order = append(fsq.order, o...)
	return fsq
}

// QueryIdentities chains the current query on the "identities" edge.
func (fsq *FileSourceQuery) QueryIdentities() *FileIdentityQuery {
	query := (&FileIdentityClient{config: fsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(filesource.Table, filesource.FieldID, selector),
			sqlgraph.To(fileidentity.Table, fileidentity.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, filesource.IdentitiesTable, filesource.IdentitiesColumn),
		)
		fromU = sqlgraph.SetNeighbors(fsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFiles chains the current query on the "files" edge.
func (fsq *FileSourceQuery) QueryFiles() *FileQuery {
	query := (&FileClient{config: fsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(filesource.Table, filesource.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, filesource.FilesTable, filesource.FilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(fsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FileSource entity from the query.
// Returns a *NotFoundError when no FileSource was found.
func (fsq *FileSourceQuery) First(ctx context.Context) (*FileSource, error) {
	nodes, err := fsq.Limit(1).All(setContextOp(ctx, fsq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{filesource.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fsq *FileSourceQuery) FirstX(ctx context.Context) *FileSource {
	node, err := fsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FileSource ID from the query.
// Returns a *NotFoundError when no FileSource ID was found.
func (fsq *FileSourceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fsq.Limit(1).IDs(setContextOp(ctx, fsq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{filesource.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fsq *FileSourceQuery) FirstIDX(ctx context.Context) int {
	id, err := fsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FileSource entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FileSource entity is found.
// Returns a *NotFoundError when no FileSource entities are found.
func (fsq *FileSourceQuery) Only(ctx context.Context) (*FileSource, error) {
	nodes, err := fsq.Limit(2).All(setContextOp(ctx, fsq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{filesource.Label}
	default:
		return nil, &NotSingularError{filesource.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fsq *FileSourceQuery) OnlyX(ctx context.Context) *FileSource {
	node, err := fsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FileSource ID in the query.
// Returns a *NotSingularError when more than one FileSource ID is found.
// Returns a *NotFoundError when no entities are found.
func (fsq *FileSourceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fsq.Limit(2).IDs(setContextOp(ctx, fsq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{filesource.Label}
	default:
		err = &NotSingularError{filesource.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fsq *FileSourceQuery) OnlyIDX(ctx context.Context) int {
	id, err := fsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FileSources.
func (fsq *FileSourceQuery) All(ctx context.Context) ([]*FileSource, error) {
	ctx = setContextOp(ctx, fsq.ctx, ent.OpQueryAll)
	if err := fsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FileSource, *FileSourceQuery]()
	return withInterceptors[[]*FileSource](ctx, fsq, qr, fsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fsq *FileSourceQuery) AllX(ctx context.Context) []*FileSource {
	nodes, err := fsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FileSource IDs.
func (fsq *FileSourceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fsq.ctx.Unique == nil && fsq.path != nil {
		fsq.Unique(true)
	}
	ctx = setContextOp(ctx, fsq.ctx, ent.OpQueryIDs)
	if err = fsq.Select(filesource.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fsq *FileSourceQuery) IDsX(ctx context.Context) []int {
	ids, err := fsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fsq *FileSourceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fsq.ctx, ent.OpQueryCount)
	if err := fsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fsq, querierCount[*FileSourceQuery](), fsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fsq *FileSourceQuery) CountX(ctx context.Context) int {
	count, err := fsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fsq *FileSourceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fsq.ctx, ent.OpQueryExist)
	switch _, err := fsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fsq *FileSourceQuery) ExistX(ctx context.Context) bool {
	exist, err := fsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FileSourceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fsq *FileSourceQuery) Clone() *FileSourceQuery {
	if fsq == nil {
		return nil
	}
	return &FileSourceQuery{
		config:         fsq.config,
		ctx:            fsq.ctx.Clone(),
		order:          append([]filesource.OrderOption{}, fsq.order...),
		inters:         append([]Interceptor{}, fsq.inters...),
		predicates:     append([]predicate.FileSource{}, fsq.predicates...),
		withIdentities: fsq.withIdentities.Clone(),
		withFiles:      fsq.withFiles.Clone(),
		// clone intermediate query.
		sql:  fsq.sql.Clone(),
		path: fsq.path,
	}
}

// WithIdentities tells the query-builder to eager-load the nodes that are connected to
// the "identities" edge. The optional arguments are used to configure the query builder of the edge.
func (fsq *FileSourceQuery) WithIdentities(opts ...func(*FileIdentityQuery)) *FileSourceQuery {
	query := (&FileIdentityClient{config: fsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fsq.withIdentities = query
	return fsq
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (fsq *FileSourceQuery) WithFiles(opts ...func(*FileQuery)) *FileSourceQuery {
	query := (&FileClient{config: fsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fsq.withFiles = query
	return fsq
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
//	client.FileSource.Query().
//		GroupBy(filesource.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fsq *FileSourceQuery) GroupBy(field string, fields ...string) *FileSourceGroupBy {
	fsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FileSourceGroupBy{build: fsq}
	grbuild.flds = &fsq.ctx.Fields
	grbuild.label = filesource.Label
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
//	client.FileSource.Query().
//		Select(filesource.FieldCreatedBy).
//		Scan(ctx, &v)
func (fsq *FileSourceQuery) Select(fields ...string) *FileSourceSelect {
	fsq.ctx.Fields = append(fsq.ctx.Fields, fields...)
	sbuild := &FileSourceSelect{FileSourceQuery: fsq}
	sbuild.label = filesource.Label
	sbuild.flds, sbuild.scan = &fsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FileSourceSelect configured with the given aggregations.
func (fsq *FileSourceQuery) Aggregate(fns ...AggregateFunc) *FileSourceSelect {
	return fsq.Select().Aggregate(fns...)
}

func (fsq *FileSourceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fsq); err != nil {
				return err
			}
		}
	}
	for _, f := range fsq.ctx.Fields {
		if !filesource.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fsq.path != nil {
		prev, err := fsq.path(ctx)
		if err != nil {
			return err
		}
		fsq.sql = prev
	}
	return nil
}

func (fsq *FileSourceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FileSource, error) {
	var (
		nodes       = []*FileSource{}
		_spec       = fsq.querySpec()
		loadedTypes = [2]bool{
			fsq.withIdentities != nil,
			fsq.withFiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FileSource).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FileSource{config: fsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(fsq.modifiers) > 0 {
		_spec.Modifiers = fsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fsq.withIdentities; query != nil {
		if err := fsq.loadIdentities(ctx, query, nodes,
			func(n *FileSource) { n.Edges.Identities = []*FileIdentity{} },
			func(n *FileSource, e *FileIdentity) { n.Edges.Identities = append(n.Edges.Identities, e) }); err != nil {
			return nil, err
		}
	}
	if query := fsq.withFiles; query != nil {
		if err := fsq.loadFiles(ctx, query, nodes,
			func(n *FileSource) { n.Edges.Files = []*File{} },
			func(n *FileSource, e *File) { n.Edges.Files = append(n.Edges.Files, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range fsq.withNamedIdentities {
		if err := fsq.loadIdentities(ctx, query, nodes,
			func(n *FileSource) { n.appendNamedIdentities(name) },
			func(n *FileSource, e *FileIdentity) { n.appendNamedIdentities(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range fsq.withNamedFiles {
		if err := fsq.loadFiles(ctx, query, nodes,
			func(n *FileSource) { n.appendNamedFiles(name) },
			func(n *FileSource, e *File) { n.appendNamedFiles(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range fsq.loadTotal {
		if err := fsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fsq *FileSourceQuery) loadIdentities(ctx context.Context, query *FileIdentityQuery, nodes []*FileSource, init func(*FileSource), assign func(*FileSource, *FileIdentity)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*FileSource)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(fileidentity.FieldFileSourceID)
	}
	query.Where(predicate.FileIdentity(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(filesource.IdentitiesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.FileSourceID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "file_source_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (fsq *FileSourceQuery) loadFiles(ctx context.Context, query *FileQuery, nodes []*FileSource, init func(*FileSource), assign func(*FileSource, *File)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*FileSource)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(file.FieldSourceID)
	}
	query.Where(predicate.File(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(filesource.FilesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.SourceID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "source_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fsq *FileSourceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fsq.querySpec()
	if len(fsq.modifiers) > 0 {
		_spec.Modifiers = fsq.modifiers
	}
	_spec.Node.Columns = fsq.ctx.Fields
	if len(fsq.ctx.Fields) > 0 {
		_spec.Unique = fsq.ctx.Unique != nil && *fsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fsq.driver, _spec)
}

func (fsq *FileSourceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(filesource.Table, filesource.Columns, sqlgraph.NewFieldSpec(filesource.FieldID, field.TypeInt))
	_spec.From = fsq.sql
	if unique := fsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fsq.path != nil {
		_spec.Unique = true
	}
	if fields := fsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filesource.FieldID)
		for i := range fields {
			if fields[i] != filesource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fsq *FileSourceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fsq.driver.Dialect())
	t1 := builder.Table(filesource.Table)
	columns := fsq.ctx.Fields
	if len(columns) == 0 {
		columns = filesource.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fsq.sql != nil {
		selector = fsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fsq.ctx.Unique != nil && *fsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fsq.predicates {
		p(selector)
	}
	for _, p := range fsq.order {
		p(selector)
	}
	if offset := fsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedIdentities tells the query-builder to eager-load the nodes that are connected to the "identities"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (fsq *FileSourceQuery) WithNamedIdentities(name string, opts ...func(*FileIdentityQuery)) *FileSourceQuery {
	query := (&FileIdentityClient{config: fsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if fsq.withNamedIdentities == nil {
		fsq.withNamedIdentities = make(map[string]*FileIdentityQuery)
	}
	fsq.withNamedIdentities[name] = query
	return fsq
}

// WithNamedFiles tells the query-builder to eager-load the nodes that are connected to the "files"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (fsq *FileSourceQuery) WithNamedFiles(name string, opts ...func(*FileQuery)) *FileSourceQuery {
	query := (&FileClient{config: fsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if fsq.withNamedFiles == nil {
		fsq.withNamedFiles = make(map[string]*FileQuery)
	}
	fsq.withNamedFiles[name] = query
	return fsq
}

// FileSourceGroupBy is the group-by builder for FileSource entities.
type FileSourceGroupBy struct {
	selector
	build *FileSourceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fsgb *FileSourceGroupBy) Aggregate(fns ...AggregateFunc) *FileSourceGroupBy {
	fsgb.fns = append(fsgb.fns, fns...)
	return fsgb
}

// Scan applies the selector query and scans the result into the given value.
func (fsgb *FileSourceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fsgb.build.ctx, ent.OpQueryGroupBy)
	if err := fsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileSourceQuery, *FileSourceGroupBy](ctx, fsgb.build, fsgb, fsgb.build.inters, v)
}

func (fsgb *FileSourceGroupBy) sqlScan(ctx context.Context, root *FileSourceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fsgb.fns))
	for _, fn := range fsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fsgb.flds)+len(fsgb.fns))
		for _, f := range *fsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FileSourceSelect is the builder for selecting fields of FileSource entities.
type FileSourceSelect struct {
	*FileSourceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fss *FileSourceSelect) Aggregate(fns ...AggregateFunc) *FileSourceSelect {
	fss.fns = append(fss.fns, fns...)
	return fss
}

// Scan applies the selector query and scans the result into the given value.
func (fss *FileSourceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fss.ctx, ent.OpQuerySelect)
	if err := fss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileSourceQuery, *FileSourceSelect](ctx, fss.FileSourceQuery, fss, fss.inters, v)
}

func (fss *FileSourceSelect) sqlScan(ctx context.Context, root *FileSourceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fss.fns))
	for _, fn := range fss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
