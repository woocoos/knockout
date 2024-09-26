// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/predicate"
	"github.com/woocoos/knockout/ent/region"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/useraddr"
)

// UserAddrQuery is the builder for querying UserAddr entities.
type UserAddrQuery struct {
	config
	ctx        *QueryContext
	order      []useraddr.OrderOption
	inters     []Interceptor
	predicates []predicate.UserAddr
	withUser   *UserQuery
	withRegion *RegionQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*UserAddr) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserAddrQuery builder.
func (uaq *UserAddrQuery) Where(ps ...predicate.UserAddr) *UserAddrQuery {
	uaq.predicates = append(uaq.predicates, ps...)
	return uaq
}

// Limit the number of records to be returned by this query.
func (uaq *UserAddrQuery) Limit(limit int) *UserAddrQuery {
	uaq.ctx.Limit = &limit
	return uaq
}

// Offset to start from.
func (uaq *UserAddrQuery) Offset(offset int) *UserAddrQuery {
	uaq.ctx.Offset = &offset
	return uaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uaq *UserAddrQuery) Unique(unique bool) *UserAddrQuery {
	uaq.ctx.Unique = &unique
	return uaq
}

// Order specifies how the records should be ordered.
func (uaq *UserAddrQuery) Order(o ...useraddr.OrderOption) *UserAddrQuery {
	uaq.order = append(uaq.order, o...)
	return uaq
}

// QueryUser chains the current query on the "user" edge.
func (uaq *UserAddrQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useraddr.Table, useraddr.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, useraddr.UserTable, useraddr.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRegion chains the current query on the "region" edge.
func (uaq *UserAddrQuery) QueryRegion() *RegionQuery {
	query := (&RegionClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useraddr.Table, useraddr.FieldID, selector),
			sqlgraph.To(region.Table, region.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, useraddr.RegionTable, useraddr.RegionColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserAddr entity from the query.
// Returns a *NotFoundError when no UserAddr was found.
func (uaq *UserAddrQuery) First(ctx context.Context) (*UserAddr, error) {
	nodes, err := uaq.Limit(1).All(setContextOp(ctx, uaq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{useraddr.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uaq *UserAddrQuery) FirstX(ctx context.Context) *UserAddr {
	node, err := uaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserAddr ID from the query.
// Returns a *NotFoundError when no UserAddr ID was found.
func (uaq *UserAddrQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(1).IDs(setContextOp(ctx, uaq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{useraddr.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uaq *UserAddrQuery) FirstIDX(ctx context.Context) int {
	id, err := uaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserAddr entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserAddr entity is found.
// Returns a *NotFoundError when no UserAddr entities are found.
func (uaq *UserAddrQuery) Only(ctx context.Context) (*UserAddr, error) {
	nodes, err := uaq.Limit(2).All(setContextOp(ctx, uaq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{useraddr.Label}
	default:
		return nil, &NotSingularError{useraddr.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uaq *UserAddrQuery) OnlyX(ctx context.Context) *UserAddr {
	node, err := uaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserAddr ID in the query.
// Returns a *NotSingularError when more than one UserAddr ID is found.
// Returns a *NotFoundError when no entities are found.
func (uaq *UserAddrQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(2).IDs(setContextOp(ctx, uaq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{useraddr.Label}
	default:
		err = &NotSingularError{useraddr.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uaq *UserAddrQuery) OnlyIDX(ctx context.Context) int {
	id, err := uaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserAddrs.
func (uaq *UserAddrQuery) All(ctx context.Context) ([]*UserAddr, error) {
	ctx = setContextOp(ctx, uaq.ctx, ent.OpQueryAll)
	if err := uaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserAddr, *UserAddrQuery]()
	return withInterceptors[[]*UserAddr](ctx, uaq, qr, uaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uaq *UserAddrQuery) AllX(ctx context.Context) []*UserAddr {
	nodes, err := uaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserAddr IDs.
func (uaq *UserAddrQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uaq.ctx.Unique == nil && uaq.path != nil {
		uaq.Unique(true)
	}
	ctx = setContextOp(ctx, uaq.ctx, ent.OpQueryIDs)
	if err = uaq.Select(useraddr.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uaq *UserAddrQuery) IDsX(ctx context.Context) []int {
	ids, err := uaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uaq *UserAddrQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uaq.ctx, ent.OpQueryCount)
	if err := uaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uaq, querierCount[*UserAddrQuery](), uaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uaq *UserAddrQuery) CountX(ctx context.Context) int {
	count, err := uaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uaq *UserAddrQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uaq.ctx, ent.OpQueryExist)
	switch _, err := uaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uaq *UserAddrQuery) ExistX(ctx context.Context) bool {
	exist, err := uaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserAddrQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uaq *UserAddrQuery) Clone() *UserAddrQuery {
	if uaq == nil {
		return nil
	}
	return &UserAddrQuery{
		config:     uaq.config,
		ctx:        uaq.ctx.Clone(),
		order:      append([]useraddr.OrderOption{}, uaq.order...),
		inters:     append([]Interceptor{}, uaq.inters...),
		predicates: append([]predicate.UserAddr{}, uaq.predicates...),
		withUser:   uaq.withUser.Clone(),
		withRegion: uaq.withRegion.Clone(),
		// clone intermediate query.
		sql:  uaq.sql.Clone(),
		path: uaq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserAddrQuery) WithUser(opts ...func(*UserQuery)) *UserAddrQuery {
	query := (&UserClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withUser = query
	return uaq
}

// WithRegion tells the query-builder to eager-load the nodes that are connected to
// the "region" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserAddrQuery) WithRegion(opts ...func(*RegionQuery)) *UserAddrQuery {
	query := (&RegionClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withRegion = query
	return uaq
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
//	client.UserAddr.Query().
//		GroupBy(useraddr.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uaq *UserAddrQuery) GroupBy(field string, fields ...string) *UserAddrGroupBy {
	uaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserAddrGroupBy{build: uaq}
	grbuild.flds = &uaq.ctx.Fields
	grbuild.label = useraddr.Label
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
//	client.UserAddr.Query().
//		Select(useraddr.FieldCreatedBy).
//		Scan(ctx, &v)
func (uaq *UserAddrQuery) Select(fields ...string) *UserAddrSelect {
	uaq.ctx.Fields = append(uaq.ctx.Fields, fields...)
	sbuild := &UserAddrSelect{UserAddrQuery: uaq}
	sbuild.label = useraddr.Label
	sbuild.flds, sbuild.scan = &uaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserAddrSelect configured with the given aggregations.
func (uaq *UserAddrQuery) Aggregate(fns ...AggregateFunc) *UserAddrSelect {
	return uaq.Select().Aggregate(fns...)
}

func (uaq *UserAddrQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uaq); err != nil {
				return err
			}
		}
	}
	for _, f := range uaq.ctx.Fields {
		if !useraddr.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uaq.path != nil {
		prev, err := uaq.path(ctx)
		if err != nil {
			return err
		}
		uaq.sql = prev
	}
	return nil
}

func (uaq *UserAddrQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserAddr, error) {
	var (
		nodes       = []*UserAddr{}
		_spec       = uaq.querySpec()
		loadedTypes = [2]bool{
			uaq.withUser != nil,
			uaq.withRegion != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserAddr).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserAddr{config: uaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(uaq.modifiers) > 0 {
		_spec.Modifiers = uaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uaq.withUser; query != nil {
		if err := uaq.loadUser(ctx, query, nodes, nil,
			func(n *UserAddr, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := uaq.withRegion; query != nil {
		if err := uaq.loadRegion(ctx, query, nodes, nil,
			func(n *UserAddr, e *Region) { n.Edges.Region = e }); err != nil {
			return nil, err
		}
	}
	for i := range uaq.loadTotal {
		if err := uaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uaq *UserAddrQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserAddr, init func(*UserAddr), assign func(*UserAddr, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserAddr)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (uaq *UserAddrQuery) loadRegion(ctx context.Context, query *RegionQuery, nodes []*UserAddr, init func(*UserAddr), assign func(*UserAddr, *Region)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserAddr)
	for i := range nodes {
		if nodes[i].RegionID == nil {
			continue
		}
		fk := *nodes[i].RegionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(region.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "region_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (uaq *UserAddrQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uaq.querySpec()
	if len(uaq.modifiers) > 0 {
		_spec.Modifiers = uaq.modifiers
	}
	_spec.Node.Columns = uaq.ctx.Fields
	if len(uaq.ctx.Fields) > 0 {
		_spec.Unique = uaq.ctx.Unique != nil && *uaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uaq.driver, _spec)
}

func (uaq *UserAddrQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(useraddr.Table, useraddr.Columns, sqlgraph.NewFieldSpec(useraddr.FieldID, field.TypeInt))
	_spec.From = uaq.sql
	if unique := uaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uaq.path != nil {
		_spec.Unique = true
	}
	if fields := uaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, useraddr.FieldID)
		for i := range fields {
			if fields[i] != useraddr.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if uaq.withUser != nil {
			_spec.Node.AddColumnOnce(useraddr.FieldUserID)
		}
		if uaq.withRegion != nil {
			_spec.Node.AddColumnOnce(useraddr.FieldRegionID)
		}
	}
	if ps := uaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uaq *UserAddrQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uaq.driver.Dialect())
	t1 := builder.Table(useraddr.Table)
	columns := uaq.ctx.Fields
	if len(columns) == 0 {
		columns = useraddr.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uaq.sql != nil {
		selector = uaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uaq.ctx.Unique != nil && *uaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uaq.predicates {
		p(selector)
	}
	for _, p := range uaq.order {
		p(selector)
	}
	if offset := uaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserAddrGroupBy is the group-by builder for UserAddr entities.
type UserAddrGroupBy struct {
	selector
	build *UserAddrQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uagb *UserAddrGroupBy) Aggregate(fns ...AggregateFunc) *UserAddrGroupBy {
	uagb.fns = append(uagb.fns, fns...)
	return uagb
}

// Scan applies the selector query and scans the result into the given value.
func (uagb *UserAddrGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uagb.build.ctx, ent.OpQueryGroupBy)
	if err := uagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserAddrQuery, *UserAddrGroupBy](ctx, uagb.build, uagb, uagb.build.inters, v)
}

func (uagb *UserAddrGroupBy) sqlScan(ctx context.Context, root *UserAddrQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uagb.fns))
	for _, fn := range uagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uagb.flds)+len(uagb.fns))
		for _, f := range *uagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserAddrSelect is the builder for selecting fields of UserAddr entities.
type UserAddrSelect struct {
	*UserAddrQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uas *UserAddrSelect) Aggregate(fns ...AggregateFunc) *UserAddrSelect {
	uas.fns = append(uas.fns, fns...)
	return uas
}

// Scan applies the selector query and scans the result into the given value.
func (uas *UserAddrSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uas.ctx, ent.OpQuerySelect)
	if err := uas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserAddrQuery, *UserAddrSelect](ctx, uas.UserAddrQuery, uas, uas.inters, v)
}

func (uas *UserAddrSelect) sqlScan(ctx context.Context, root *UserAddrQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uas.fns))
	for _, fn := range uas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
