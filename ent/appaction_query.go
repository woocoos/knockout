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
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/predicate"
)

// AppActionQuery is the builder for querying AppAction entities.
type AppActionQuery struct {
	config
	ctx            *QueryContext
	order          []appaction.OrderOption
	inters         []Interceptor
	predicates     []predicate.AppAction
	withApp        *AppQuery
	withMenus      *AppMenuQuery
	modifiers      []func(*sql.Selector)
	loadTotal      []func(context.Context, []*AppAction) error
	withNamedMenus map[string]*AppMenuQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppActionQuery builder.
func (aaq *AppActionQuery) Where(ps ...predicate.AppAction) *AppActionQuery {
	aaq.predicates = append(aaq.predicates, ps...)
	return aaq
}

// Limit the number of records to be returned by this query.
func (aaq *AppActionQuery) Limit(limit int) *AppActionQuery {
	aaq.ctx.Limit = &limit
	return aaq
}

// Offset to start from.
func (aaq *AppActionQuery) Offset(offset int) *AppActionQuery {
	aaq.ctx.Offset = &offset
	return aaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aaq *AppActionQuery) Unique(unique bool) *AppActionQuery {
	aaq.ctx.Unique = &unique
	return aaq
}

// Order specifies how the records should be ordered.
func (aaq *AppActionQuery) Order(o ...appaction.OrderOption) *AppActionQuery {
	aaq.order = append(aaq.order, o...)
	return aaq
}

// QueryApp chains the current query on the "app" edge.
func (aaq *AppActionQuery) QueryApp() *AppQuery {
	query := (&AppClient{config: aaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appaction.Table, appaction.FieldID, selector),
			sqlgraph.To(app.Table, app.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, appaction.AppTable, appaction.AppColumn),
		)
		fromU = sqlgraph.SetNeighbors(aaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMenus chains the current query on the "menus" edge.
func (aaq *AppActionQuery) QueryMenus() *AppMenuQuery {
	query := (&AppMenuClient{config: aaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(appaction.Table, appaction.FieldID, selector),
			sqlgraph.To(appmenu.Table, appmenu.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, appaction.MenusTable, appaction.MenusColumn),
		)
		fromU = sqlgraph.SetNeighbors(aaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AppAction entity from the query.
// Returns a *NotFoundError when no AppAction was found.
func (aaq *AppActionQuery) First(ctx context.Context) (*AppAction, error) {
	nodes, err := aaq.Limit(1).All(setContextOp(ctx, aaq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appaction.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aaq *AppActionQuery) FirstX(ctx context.Context) *AppAction {
	node, err := aaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppAction ID from the query.
// Returns a *NotFoundError when no AppAction ID was found.
func (aaq *AppActionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aaq.Limit(1).IDs(setContextOp(ctx, aaq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appaction.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aaq *AppActionQuery) FirstIDX(ctx context.Context) int {
	id, err := aaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppAction entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppAction entity is found.
// Returns a *NotFoundError when no AppAction entities are found.
func (aaq *AppActionQuery) Only(ctx context.Context) (*AppAction, error) {
	nodes, err := aaq.Limit(2).All(setContextOp(ctx, aaq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appaction.Label}
	default:
		return nil, &NotSingularError{appaction.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aaq *AppActionQuery) OnlyX(ctx context.Context) *AppAction {
	node, err := aaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppAction ID in the query.
// Returns a *NotSingularError when more than one AppAction ID is found.
// Returns a *NotFoundError when no entities are found.
func (aaq *AppActionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aaq.Limit(2).IDs(setContextOp(ctx, aaq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appaction.Label}
	default:
		err = &NotSingularError{appaction.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aaq *AppActionQuery) OnlyIDX(ctx context.Context) int {
	id, err := aaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppActions.
func (aaq *AppActionQuery) All(ctx context.Context) ([]*AppAction, error) {
	ctx = setContextOp(ctx, aaq.ctx, ent.OpQueryAll)
	if err := aaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppAction, *AppActionQuery]()
	return withInterceptors[[]*AppAction](ctx, aaq, qr, aaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aaq *AppActionQuery) AllX(ctx context.Context) []*AppAction {
	nodes, err := aaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppAction IDs.
func (aaq *AppActionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aaq.ctx.Unique == nil && aaq.path != nil {
		aaq.Unique(true)
	}
	ctx = setContextOp(ctx, aaq.ctx, ent.OpQueryIDs)
	if err = aaq.Select(appaction.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aaq *AppActionQuery) IDsX(ctx context.Context) []int {
	ids, err := aaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aaq *AppActionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aaq.ctx, ent.OpQueryCount)
	if err := aaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aaq, querierCount[*AppActionQuery](), aaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aaq *AppActionQuery) CountX(ctx context.Context) int {
	count, err := aaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aaq *AppActionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aaq.ctx, ent.OpQueryExist)
	switch _, err := aaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aaq *AppActionQuery) ExistX(ctx context.Context) bool {
	exist, err := aaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppActionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aaq *AppActionQuery) Clone() *AppActionQuery {
	if aaq == nil {
		return nil
	}
	return &AppActionQuery{
		config:     aaq.config,
		ctx:        aaq.ctx.Clone(),
		order:      append([]appaction.OrderOption{}, aaq.order...),
		inters:     append([]Interceptor{}, aaq.inters...),
		predicates: append([]predicate.AppAction{}, aaq.predicates...),
		withApp:    aaq.withApp.Clone(),
		withMenus:  aaq.withMenus.Clone(),
		// clone intermediate query.
		sql:  aaq.sql.Clone(),
		path: aaq.path,
	}
}

// WithApp tells the query-builder to eager-load the nodes that are connected to
// the "app" edge. The optional arguments are used to configure the query builder of the edge.
func (aaq *AppActionQuery) WithApp(opts ...func(*AppQuery)) *AppActionQuery {
	query := (&AppClient{config: aaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aaq.withApp = query
	return aaq
}

// WithMenus tells the query-builder to eager-load the nodes that are connected to
// the "menus" edge. The optional arguments are used to configure the query builder of the edge.
func (aaq *AppActionQuery) WithMenus(opts ...func(*AppMenuQuery)) *AppActionQuery {
	query := (&AppMenuClient{config: aaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aaq.withMenus = query
	return aaq
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
//	client.AppAction.Query().
//		GroupBy(appaction.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aaq *AppActionQuery) GroupBy(field string, fields ...string) *AppActionGroupBy {
	aaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppActionGroupBy{build: aaq}
	grbuild.flds = &aaq.ctx.Fields
	grbuild.label = appaction.Label
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
//	client.AppAction.Query().
//		Select(appaction.FieldCreatedBy).
//		Scan(ctx, &v)
func (aaq *AppActionQuery) Select(fields ...string) *AppActionSelect {
	aaq.ctx.Fields = append(aaq.ctx.Fields, fields...)
	sbuild := &AppActionSelect{AppActionQuery: aaq}
	sbuild.label = appaction.Label
	sbuild.flds, sbuild.scan = &aaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppActionSelect configured with the given aggregations.
func (aaq *AppActionQuery) Aggregate(fns ...AggregateFunc) *AppActionSelect {
	return aaq.Select().Aggregate(fns...)
}

func (aaq *AppActionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aaq); err != nil {
				return err
			}
		}
	}
	for _, f := range aaq.ctx.Fields {
		if !appaction.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aaq.path != nil {
		prev, err := aaq.path(ctx)
		if err != nil {
			return err
		}
		aaq.sql = prev
	}
	return nil
}

func (aaq *AppActionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppAction, error) {
	var (
		nodes       = []*AppAction{}
		_spec       = aaq.querySpec()
		loadedTypes = [2]bool{
			aaq.withApp != nil,
			aaq.withMenus != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppAction).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppAction{config: aaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aaq.modifiers) > 0 {
		_spec.Modifiers = aaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aaq.withApp; query != nil {
		if err := aaq.loadApp(ctx, query, nodes, nil,
			func(n *AppAction, e *App) { n.Edges.App = e }); err != nil {
			return nil, err
		}
	}
	if query := aaq.withMenus; query != nil {
		if err := aaq.loadMenus(ctx, query, nodes,
			func(n *AppAction) { n.Edges.Menus = []*AppMenu{} },
			func(n *AppAction, e *AppMenu) { n.Edges.Menus = append(n.Edges.Menus, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aaq.withNamedMenus {
		if err := aaq.loadMenus(ctx, query, nodes,
			func(n *AppAction) { n.appendNamedMenus(name) },
			func(n *AppAction, e *AppMenu) { n.appendNamedMenus(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range aaq.loadTotal {
		if err := aaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aaq *AppActionQuery) loadApp(ctx context.Context, query *AppQuery, nodes []*AppAction, init func(*AppAction), assign func(*AppAction, *App)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*AppAction)
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
func (aaq *AppActionQuery) loadMenus(ctx context.Context, query *AppMenuQuery, nodes []*AppAction, init func(*AppAction), assign func(*AppAction, *AppMenu)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*AppAction)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(appmenu.FieldActionID)
	}
	query.Where(predicate.AppMenu(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(appaction.MenusColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ActionID
		if fk == nil {
			return fmt.Errorf(`foreign-key "action_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "action_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aaq *AppActionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aaq.querySpec()
	if len(aaq.modifiers) > 0 {
		_spec.Modifiers = aaq.modifiers
	}
	_spec.Node.Columns = aaq.ctx.Fields
	if len(aaq.ctx.Fields) > 0 {
		_spec.Unique = aaq.ctx.Unique != nil && *aaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aaq.driver, _spec)
}

func (aaq *AppActionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appaction.Table, appaction.Columns, sqlgraph.NewFieldSpec(appaction.FieldID, field.TypeInt))
	_spec.From = aaq.sql
	if unique := aaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aaq.path != nil {
		_spec.Unique = true
	}
	if fields := aaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appaction.FieldID)
		for i := range fields {
			if fields[i] != appaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if aaq.withApp != nil {
			_spec.Node.AddColumnOnce(appaction.FieldAppID)
		}
	}
	if ps := aaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aaq *AppActionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aaq.driver.Dialect())
	t1 := builder.Table(appaction.Table)
	columns := aaq.ctx.Fields
	if len(columns) == 0 {
		columns = appaction.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aaq.sql != nil {
		selector = aaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aaq.ctx.Unique != nil && *aaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aaq.predicates {
		p(selector)
	}
	for _, p := range aaq.order {
		p(selector)
	}
	if offset := aaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedMenus tells the query-builder to eager-load the nodes that are connected to the "menus"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aaq *AppActionQuery) WithNamedMenus(name string, opts ...func(*AppMenuQuery)) *AppActionQuery {
	query := (&AppMenuClient{config: aaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if aaq.withNamedMenus == nil {
		aaq.withNamedMenus = make(map[string]*AppMenuQuery)
	}
	aaq.withNamedMenus[name] = query
	return aaq
}

// AppActionGroupBy is the group-by builder for AppAction entities.
type AppActionGroupBy struct {
	selector
	build *AppActionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aagb *AppActionGroupBy) Aggregate(fns ...AggregateFunc) *AppActionGroupBy {
	aagb.fns = append(aagb.fns, fns...)
	return aagb
}

// Scan applies the selector query and scans the result into the given value.
func (aagb *AppActionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aagb.build.ctx, ent.OpQueryGroupBy)
	if err := aagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppActionQuery, *AppActionGroupBy](ctx, aagb.build, aagb, aagb.build.inters, v)
}

func (aagb *AppActionGroupBy) sqlScan(ctx context.Context, root *AppActionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aagb.fns))
	for _, fn := range aagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aagb.flds)+len(aagb.fns))
		for _, f := range *aagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppActionSelect is the builder for selecting fields of AppAction entities.
type AppActionSelect struct {
	*AppActionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aas *AppActionSelect) Aggregate(fns ...AggregateFunc) *AppActionSelect {
	aas.fns = append(aas.fns, fns...)
	return aas
}

// Scan applies the selector query and scans the result into the given value.
func (aas *AppActionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aas.ctx, ent.OpQuerySelect)
	if err := aas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppActionQuery, *AppActionSelect](ctx, aas.AppActionQuery, aas, aas.inters, v)
}

func (aas *AppActionSelect) sqlScan(ctx context.Context, root *AppActionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aas.fns))
	for _, fn := range aas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
