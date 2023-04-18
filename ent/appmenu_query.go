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
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/predicate"
)

// AppMenuQuery is the builder for querying AppMenu entities.
type AppMenuQuery struct {
	config
	ctx        *QueryContext
	order      []appmenu.OrderOption
	inters     []Interceptor
	predicates []predicate.AppMenu
	withApp    *AppQuery
	withAction *AppActionQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*AppMenu) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppMenuQuery builder.
func (amq *AppMenuQuery) Where(ps ...predicate.AppMenu) *AppMenuQuery {
	amq.predicates = append(amq.predicates, ps...)
	return amq
}

// Limit the number of records to be returned by this query.
func (amq *AppMenuQuery) Limit(limit int) *AppMenuQuery {
	amq.ctx.Limit = &limit
	return amq
}

// Offset to start from.
func (amq *AppMenuQuery) Offset(offset int) *AppMenuQuery {
	amq.ctx.Offset = &offset
	return amq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (amq *AppMenuQuery) Unique(unique bool) *AppMenuQuery {
	amq.ctx.Unique = &unique
	return amq
}

// Order specifies how the records should be ordered.
func (amq *AppMenuQuery) Order(o ...appmenu.OrderOption) *AppMenuQuery {
	amq.order = append(amq.order, o...)
	return amq
}

// QueryApp chains the current query on the "app" edge.
func (amq *AppMenuQuery) QueryApp() *AppQuery {
	query := (&AppClient{config: amq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := amq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := amq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appmenu.Table, appmenu.FieldID, selector),
			sqlgraph.To(app.Table, app.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, appmenu.AppTable, appmenu.AppColumn),
		)
		fromU = sqlgraph.SetNeighbors(amq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAction chains the current query on the "action" edge.
func (amq *AppMenuQuery) QueryAction() *AppActionQuery {
	query := (&AppActionClient{config: amq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := amq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := amq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appmenu.Table, appmenu.FieldID, selector),
			sqlgraph.To(appaction.Table, appaction.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, appmenu.ActionTable, appmenu.ActionColumn),
		)
		fromU = sqlgraph.SetNeighbors(amq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AppMenu entity from the query.
// Returns a *NotFoundError when no AppMenu was found.
func (amq *AppMenuQuery) First(ctx context.Context) (*AppMenu, error) {
	nodes, err := amq.Limit(1).All(setContextOp(ctx, amq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appmenu.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (amq *AppMenuQuery) FirstX(ctx context.Context) *AppMenu {
	node, err := amq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppMenu ID from the query.
// Returns a *NotFoundError when no AppMenu ID was found.
func (amq *AppMenuQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = amq.Limit(1).IDs(setContextOp(ctx, amq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appmenu.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (amq *AppMenuQuery) FirstIDX(ctx context.Context) int {
	id, err := amq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppMenu entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppMenu entity is found.
// Returns a *NotFoundError when no AppMenu entities are found.
func (amq *AppMenuQuery) Only(ctx context.Context) (*AppMenu, error) {
	nodes, err := amq.Limit(2).All(setContextOp(ctx, amq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appmenu.Label}
	default:
		return nil, &NotSingularError{appmenu.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (amq *AppMenuQuery) OnlyX(ctx context.Context) *AppMenu {
	node, err := amq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppMenu ID in the query.
// Returns a *NotSingularError when more than one AppMenu ID is found.
// Returns a *NotFoundError when no entities are found.
func (amq *AppMenuQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = amq.Limit(2).IDs(setContextOp(ctx, amq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appmenu.Label}
	default:
		err = &NotSingularError{appmenu.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (amq *AppMenuQuery) OnlyIDX(ctx context.Context) int {
	id, err := amq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppMenus.
func (amq *AppMenuQuery) All(ctx context.Context) ([]*AppMenu, error) {
	ctx = setContextOp(ctx, amq.ctx, "All")
	if err := amq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppMenu, *AppMenuQuery]()
	return withInterceptors[[]*AppMenu](ctx, amq, qr, amq.inters)
}

// AllX is like All, but panics if an error occurs.
func (amq *AppMenuQuery) AllX(ctx context.Context) []*AppMenu {
	nodes, err := amq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppMenu IDs.
func (amq *AppMenuQuery) IDs(ctx context.Context) (ids []int, err error) {
	if amq.ctx.Unique == nil && amq.path != nil {
		amq.Unique(true)
	}
	ctx = setContextOp(ctx, amq.ctx, "IDs")
	if err = amq.Select(appmenu.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (amq *AppMenuQuery) IDsX(ctx context.Context) []int {
	ids, err := amq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (amq *AppMenuQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, amq.ctx, "Count")
	if err := amq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, amq, querierCount[*AppMenuQuery](), amq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (amq *AppMenuQuery) CountX(ctx context.Context) int {
	count, err := amq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (amq *AppMenuQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, amq.ctx, "Exist")
	switch _, err := amq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (amq *AppMenuQuery) ExistX(ctx context.Context) bool {
	exist, err := amq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppMenuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (amq *AppMenuQuery) Clone() *AppMenuQuery {
	if amq == nil {
		return nil
	}
	return &AppMenuQuery{
		config:     amq.config,
		ctx:        amq.ctx.Clone(),
		order:      append([]appmenu.OrderOption{}, amq.order...),
		inters:     append([]Interceptor{}, amq.inters...),
		predicates: append([]predicate.AppMenu{}, amq.predicates...),
		withApp:    amq.withApp.Clone(),
		withAction: amq.withAction.Clone(),
		// clone intermediate query.
		sql:  amq.sql.Clone(),
		path: amq.path,
	}
}

// WithApp tells the query-builder to eager-load the nodes that are connected to
// the "app" edge. The optional arguments are used to configure the query builder of the edge.
func (amq *AppMenuQuery) WithApp(opts ...func(*AppQuery)) *AppMenuQuery {
	query := (&AppClient{config: amq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	amq.withApp = query
	return amq
}

// WithAction tells the query-builder to eager-load the nodes that are connected to
// the "action" edge. The optional arguments are used to configure the query builder of the edge.
func (amq *AppMenuQuery) WithAction(opts ...func(*AppActionQuery)) *AppMenuQuery {
	query := (&AppActionClient{config: amq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	amq.withAction = query
	return amq
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
//	client.AppMenu.Query().
//		GroupBy(appmenu.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (amq *AppMenuQuery) GroupBy(field string, fields ...string) *AppMenuGroupBy {
	amq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppMenuGroupBy{build: amq}
	grbuild.flds = &amq.ctx.Fields
	grbuild.label = appmenu.Label
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
//	client.AppMenu.Query().
//		Select(appmenu.FieldCreatedBy).
//		Scan(ctx, &v)
func (amq *AppMenuQuery) Select(fields ...string) *AppMenuSelect {
	amq.ctx.Fields = append(amq.ctx.Fields, fields...)
	sbuild := &AppMenuSelect{AppMenuQuery: amq}
	sbuild.label = appmenu.Label
	sbuild.flds, sbuild.scan = &amq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppMenuSelect configured with the given aggregations.
func (amq *AppMenuQuery) Aggregate(fns ...AggregateFunc) *AppMenuSelect {
	return amq.Select().Aggregate(fns...)
}

func (amq *AppMenuQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range amq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, amq); err != nil {
				return err
			}
		}
	}
	for _, f := range amq.ctx.Fields {
		if !appmenu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if amq.path != nil {
		prev, err := amq.path(ctx)
		if err != nil {
			return err
		}
		amq.sql = prev
	}
	return nil
}

func (amq *AppMenuQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppMenu, error) {
	var (
		nodes       = []*AppMenu{}
		_spec       = amq.querySpec()
		loadedTypes = [2]bool{
			amq.withApp != nil,
			amq.withAction != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppMenu).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppMenu{config: amq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(amq.modifiers) > 0 {
		_spec.Modifiers = amq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, amq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := amq.withApp; query != nil {
		if err := amq.loadApp(ctx, query, nodes, nil,
			func(n *AppMenu, e *App) { n.Edges.App = e }); err != nil {
			return nil, err
		}
	}
	if query := amq.withAction; query != nil {
		if err := amq.loadAction(ctx, query, nodes, nil,
			func(n *AppMenu, e *AppAction) { n.Edges.Action = e }); err != nil {
			return nil, err
		}
	}
	for i := range amq.loadTotal {
		if err := amq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (amq *AppMenuQuery) loadApp(ctx context.Context, query *AppQuery, nodes []*AppMenu, init func(*AppMenu), assign func(*AppMenu, *App)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AppMenu)
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
func (amq *AppMenuQuery) loadAction(ctx context.Context, query *AppActionQuery, nodes []*AppMenu, init func(*AppMenu), assign func(*AppMenu, *AppAction)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AppMenu)
	for i := range nodes {
		if nodes[i].ActionID == nil {
			continue
		}
		fk := *nodes[i].ActionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(appaction.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "action_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (amq *AppMenuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := amq.querySpec()
	if len(amq.modifiers) > 0 {
		_spec.Modifiers = amq.modifiers
	}
	_spec.Node.Columns = amq.ctx.Fields
	if len(amq.ctx.Fields) > 0 {
		_spec.Unique = amq.ctx.Unique != nil && *amq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, amq.driver, _spec)
}

func (amq *AppMenuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appmenu.Table, appmenu.Columns, sqlgraph.NewFieldSpec(appmenu.FieldID, field.TypeInt))
	_spec.From = amq.sql
	if unique := amq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if amq.path != nil {
		_spec.Unique = true
	}
	if fields := amq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appmenu.FieldID)
		for i := range fields {
			if fields[i] != appmenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if amq.withApp != nil {
			_spec.Node.AddColumnOnce(appmenu.FieldAppID)
		}
		if amq.withAction != nil {
			_spec.Node.AddColumnOnce(appmenu.FieldActionID)
		}
	}
	if ps := amq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := amq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := amq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := amq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (amq *AppMenuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(amq.driver.Dialect())
	t1 := builder.Table(appmenu.Table)
	columns := amq.ctx.Fields
	if len(columns) == 0 {
		columns = appmenu.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if amq.sql != nil {
		selector = amq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if amq.ctx.Unique != nil && *amq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range amq.predicates {
		p(selector)
	}
	for _, p := range amq.order {
		p(selector)
	}
	if offset := amq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := amq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppMenuGroupBy is the group-by builder for AppMenu entities.
type AppMenuGroupBy struct {
	selector
	build *AppMenuQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (amgb *AppMenuGroupBy) Aggregate(fns ...AggregateFunc) *AppMenuGroupBy {
	amgb.fns = append(amgb.fns, fns...)
	return amgb
}

// Scan applies the selector query and scans the result into the given value.
func (amgb *AppMenuGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, amgb.build.ctx, "GroupBy")
	if err := amgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppMenuQuery, *AppMenuGroupBy](ctx, amgb.build, amgb, amgb.build.inters, v)
}

func (amgb *AppMenuGroupBy) sqlScan(ctx context.Context, root *AppMenuQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(amgb.fns))
	for _, fn := range amgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*amgb.flds)+len(amgb.fns))
		for _, f := range *amgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*amgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := amgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppMenuSelect is the builder for selecting fields of AppMenu entities.
type AppMenuSelect struct {
	*AppMenuQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ams *AppMenuSelect) Aggregate(fns ...AggregateFunc) *AppMenuSelect {
	ams.fns = append(ams.fns, fns...)
	return ams
}

// Scan applies the selector query and scans the result into the given value.
func (ams *AppMenuSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ams.ctx, "Select")
	if err := ams.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppMenuQuery, *AppMenuSelect](ctx, ams.AppMenuQuery, ams, ams.inters, v)
}

func (ams *AppMenuSelect) sqlScan(ctx context.Context, root *AppMenuQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ams.fns))
	for _, fn := range ams.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ams.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ams.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
