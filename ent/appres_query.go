// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appres"
	"github.com/woocoos/knockout/ent/predicate"
)

// AppResQuery is the builder for querying AppRes entities.
type AppResQuery struct {
	config
	ctx        *QueryContext
	order      []appres.OrderOption
	inters     []Interceptor
	predicates []predicate.AppRes
	withApp    *AppQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*AppRes) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppResQuery builder.
func (arq *AppResQuery) Where(ps ...predicate.AppRes) *AppResQuery {
	arq.predicates = append(arq.predicates, ps...)
	return arq
}

// Limit the number of records to be returned by this query.
func (arq *AppResQuery) Limit(limit int) *AppResQuery {
	arq.ctx.Limit = &limit
	return arq
}

// Offset to start from.
func (arq *AppResQuery) Offset(offset int) *AppResQuery {
	arq.ctx.Offset = &offset
	return arq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (arq *AppResQuery) Unique(unique bool) *AppResQuery {
	arq.ctx.Unique = &unique
	return arq
}

// Order specifies how the records should be ordered.
func (arq *AppResQuery) Order(o ...appres.OrderOption) *AppResQuery {
	arq.order = append(arq.order, o...)
	return arq
}

// QueryApp chains the current query on the "app" edge.
func (arq *AppResQuery) QueryApp() *AppQuery {
	query := (&AppClient{config: arq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := arq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := arq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appres.Table, appres.FieldID, selector),
			sqlgraph.To(app.Table, app.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, appres.AppTable, appres.AppColumn),
		)
		fromU = sqlgraph.SetNeighbors(arq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AppRes entity from the query.
// Returns a *NotFoundError when no AppRes was found.
func (arq *AppResQuery) First(ctx context.Context) (*AppRes, error) {
	nodes, err := arq.Limit(1).All(setContextOp(ctx, arq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appres.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (arq *AppResQuery) FirstX(ctx context.Context) *AppRes {
	node, err := arq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppRes ID from the query.
// Returns a *NotFoundError when no AppRes ID was found.
func (arq *AppResQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(1).IDs(setContextOp(ctx, arq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appres.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (arq *AppResQuery) FirstIDX(ctx context.Context) int {
	id, err := arq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppRes entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppRes entity is found.
// Returns a *NotFoundError when no AppRes entities are found.
func (arq *AppResQuery) Only(ctx context.Context) (*AppRes, error) {
	nodes, err := arq.Limit(2).All(setContextOp(ctx, arq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appres.Label}
	default:
		return nil, &NotSingularError{appres.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (arq *AppResQuery) OnlyX(ctx context.Context) *AppRes {
	node, err := arq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppRes ID in the query.
// Returns a *NotSingularError when more than one AppRes ID is found.
// Returns a *NotFoundError when no entities are found.
func (arq *AppResQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = arq.Limit(2).IDs(setContextOp(ctx, arq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appres.Label}
	default:
		err = &NotSingularError{appres.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (arq *AppResQuery) OnlyIDX(ctx context.Context) int {
	id, err := arq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppResSlice.
func (arq *AppResQuery) All(ctx context.Context) ([]*AppRes, error) {
	ctx = setContextOp(ctx, arq.ctx, "All")
	if err := arq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppRes, *AppResQuery]()
	return withInterceptors[[]*AppRes](ctx, arq, qr, arq.inters)
}

// AllX is like All, but panics if an error occurs.
func (arq *AppResQuery) AllX(ctx context.Context) []*AppRes {
	nodes, err := arq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppRes IDs.
func (arq *AppResQuery) IDs(ctx context.Context) (ids []int, err error) {
	if arq.ctx.Unique == nil && arq.path != nil {
		arq.Unique(true)
	}
	ctx = setContextOp(ctx, arq.ctx, "IDs")
	if err = arq.Select(appres.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (arq *AppResQuery) IDsX(ctx context.Context) []int {
	ids, err := arq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (arq *AppResQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, arq.ctx, "Count")
	if err := arq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, arq, querierCount[*AppResQuery](), arq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (arq *AppResQuery) CountX(ctx context.Context) int {
	count, err := arq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (arq *AppResQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, arq.ctx, "Exist")
	switch _, err := arq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (arq *AppResQuery) ExistX(ctx context.Context) bool {
	exist, err := arq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppResQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (arq *AppResQuery) Clone() *AppResQuery {
	if arq == nil {
		return nil
	}
	return &AppResQuery{
		config:     arq.config,
		ctx:        arq.ctx.Clone(),
		order:      append([]appres.OrderOption{}, arq.order...),
		inters:     append([]Interceptor{}, arq.inters...),
		predicates: append([]predicate.AppRes{}, arq.predicates...),
		withApp:    arq.withApp.Clone(),
		// clone intermediate query.
		sql:  arq.sql.Clone(),
		path: arq.path,
	}
}

// WithApp tells the query-builder to eager-load the nodes that are connected to
// the "app" edge. The optional arguments are used to configure the query builder of the edge.
func (arq *AppResQuery) WithApp(opts ...func(*AppQuery)) *AppResQuery {
	query := (&AppClient{config: arq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	arq.withApp = query
	return arq
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
//	client.AppRes.Query().
//		GroupBy(appres.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (arq *AppResQuery) GroupBy(field string, fields ...string) *AppResGroupBy {
	arq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppResGroupBy{build: arq}
	grbuild.flds = &arq.ctx.Fields
	grbuild.label = appres.Label
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
//	client.AppRes.Query().
//		Select(appres.FieldCreatedBy).
//		Scan(ctx, &v)
func (arq *AppResQuery) Select(fields ...string) *AppResSelect {
	arq.ctx.Fields = append(arq.ctx.Fields, fields...)
	sbuild := &AppResSelect{AppResQuery: arq}
	sbuild.label = appres.Label
	sbuild.flds, sbuild.scan = &arq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppResSelect configured with the given aggregations.
func (arq *AppResQuery) Aggregate(fns ...AggregateFunc) *AppResSelect {
	return arq.Select().Aggregate(fns...)
}

func (arq *AppResQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range arq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, arq); err != nil {
				return err
			}
		}
	}
	for _, f := range arq.ctx.Fields {
		if !appres.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if arq.path != nil {
		prev, err := arq.path(ctx)
		if err != nil {
			return err
		}
		arq.sql = prev
	}
	return nil
}

func (arq *AppResQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppRes, error) {
	var (
		nodes       = []*AppRes{}
		_spec       = arq.querySpec()
		loadedTypes = [1]bool{
			arq.withApp != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppRes).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppRes{config: arq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(arq.modifiers) > 0 {
		_spec.Modifiers = arq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, arq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := arq.withApp; query != nil {
		if err := arq.loadApp(ctx, query, nodes, nil,
			func(n *AppRes, e *App) { n.Edges.App = e }); err != nil {
			return nil, err
		}
	}
	for i := range arq.loadTotal {
		if err := arq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (arq *AppResQuery) loadApp(ctx context.Context, query *AppQuery, nodes []*AppRes, init func(*AppRes), assign func(*AppRes, *App)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AppRes)
	for i := range nodes {
		fk := nodes[i].AppID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(app.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "app_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (arq *AppResQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := arq.querySpec()
	if len(arq.modifiers) > 0 {
		_spec.Modifiers = arq.modifiers
	}
	_spec.Node.Columns = arq.ctx.Fields
	if len(arq.ctx.Fields) > 0 {
		_spec.Unique = arq.ctx.Unique != nil && *arq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, arq.driver, _spec)
}

func (arq *AppResQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appres.Table, appres.Columns, sqlgraph.NewFieldSpec(appres.FieldID, field.TypeInt))
	_spec.From = arq.sql
	if unique := arq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if arq.path != nil {
		_spec.Unique = true
	}
	if fields := arq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appres.FieldID)
		for i := range fields {
			if fields[i] != appres.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if arq.withApp != nil {
			_spec.Node.AddColumnOnce(appres.FieldAppID)
		}
	}
	if ps := arq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := arq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := arq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := arq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (arq *AppResQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(arq.driver.Dialect())
	t1 := builder.Table(appres.Table)
	columns := arq.ctx.Fields
	if len(columns) == 0 {
		columns = appres.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if arq.sql != nil {
		selector = arq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if arq.ctx.Unique != nil && *arq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range arq.predicates {
		p(selector)
	}
	for _, p := range arq.order {
		p(selector)
	}
	if offset := arq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := arq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppResGroupBy is the group-by builder for AppRes entities.
type AppResGroupBy struct {
	selector
	build *AppResQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (argb *AppResGroupBy) Aggregate(fns ...AggregateFunc) *AppResGroupBy {
	argb.fns = append(argb.fns, fns...)
	return argb
}

// Scan applies the selector query and scans the result into the given value.
func (argb *AppResGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, argb.build.ctx, "GroupBy")
	if err := argb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppResQuery, *AppResGroupBy](ctx, argb.build, argb, argb.build.inters, v)
}

func (argb *AppResGroupBy) sqlScan(ctx context.Context, root *AppResQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(argb.fns))
	for _, fn := range argb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*argb.flds)+len(argb.fns))
		for _, f := range *argb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*argb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := argb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppResSelect is the builder for selecting fields of AppRes entities.
type AppResSelect struct {
	*AppResQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ars *AppResSelect) Aggregate(fns ...AggregateFunc) *AppResSelect {
	ars.fns = append(ars.fns, fns...)
	return ars
}

// Scan applies the selector query and scans the result into the given value.
func (ars *AppResSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ars.ctx, "Select")
	if err := ars.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppResQuery, *AppResSelect](ctx, ars.AppResQuery, ars, ars.inters, v)
}

func (ars *AppResSelect) sqlScan(ctx context.Context, root *AppResQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ars.fns))
	for _, fn := range ars.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ars.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ars.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
