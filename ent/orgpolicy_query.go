// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/predicate"
)

// OrgPolicyQuery is the builder for querying OrgPolicy entities.
type OrgPolicyQuery struct {
	config
	ctx                  *QueryContext
	order                []orgpolicy.Order
	inters               []Interceptor
	predicates           []predicate.OrgPolicy
	withOrg              *OrgQuery
	withPermissions      *PermissionQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*OrgPolicy) error
	withNamedPermissions map[string]*PermissionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrgPolicyQuery builder.
func (opq *OrgPolicyQuery) Where(ps ...predicate.OrgPolicy) *OrgPolicyQuery {
	opq.predicates = append(opq.predicates, ps...)
	return opq
}

// Limit the number of records to be returned by this query.
func (opq *OrgPolicyQuery) Limit(limit int) *OrgPolicyQuery {
	opq.ctx.Limit = &limit
	return opq
}

// Offset to start from.
func (opq *OrgPolicyQuery) Offset(offset int) *OrgPolicyQuery {
	opq.ctx.Offset = &offset
	return opq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (opq *OrgPolicyQuery) Unique(unique bool) *OrgPolicyQuery {
	opq.ctx.Unique = &unique
	return opq
}

// Order specifies how the records should be ordered.
func (opq *OrgPolicyQuery) Order(o ...orgpolicy.Order) *OrgPolicyQuery {
	opq.order = append(opq.order, o...)
	return opq
}

// QueryOrg chains the current query on the "org" edge.
func (opq *OrgPolicyQuery) QueryOrg() *OrgQuery {
	query := (&OrgClient{config: opq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := opq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := opq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgpolicy.Table, orgpolicy.FieldID, selector),
			sqlgraph.To(org.Table, org.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orgpolicy.OrgTable, orgpolicy.OrgColumn),
		)
		fromU = sqlgraph.SetNeighbors(opq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPermissions chains the current query on the "permissions" edge.
func (opq *OrgPolicyQuery) QueryPermissions() *PermissionQuery {
	query := (&PermissionClient{config: opq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := opq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := opq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgpolicy.Table, orgpolicy.FieldID, selector),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, orgpolicy.PermissionsTable, orgpolicy.PermissionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(opq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrgPolicy entity from the query.
// Returns a *NotFoundError when no OrgPolicy was found.
func (opq *OrgPolicyQuery) First(ctx context.Context) (*OrgPolicy, error) {
	nodes, err := opq.Limit(1).All(setContextOp(ctx, opq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orgpolicy.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (opq *OrgPolicyQuery) FirstX(ctx context.Context) *OrgPolicy {
	node, err := opq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrgPolicy ID from the query.
// Returns a *NotFoundError when no OrgPolicy ID was found.
func (opq *OrgPolicyQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = opq.Limit(1).IDs(setContextOp(ctx, opq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orgpolicy.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (opq *OrgPolicyQuery) FirstIDX(ctx context.Context) int {
	id, err := opq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrgPolicy entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrgPolicy entity is found.
// Returns a *NotFoundError when no OrgPolicy entities are found.
func (opq *OrgPolicyQuery) Only(ctx context.Context) (*OrgPolicy, error) {
	nodes, err := opq.Limit(2).All(setContextOp(ctx, opq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orgpolicy.Label}
	default:
		return nil, &NotSingularError{orgpolicy.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (opq *OrgPolicyQuery) OnlyX(ctx context.Context) *OrgPolicy {
	node, err := opq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrgPolicy ID in the query.
// Returns a *NotSingularError when more than one OrgPolicy ID is found.
// Returns a *NotFoundError when no entities are found.
func (opq *OrgPolicyQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = opq.Limit(2).IDs(setContextOp(ctx, opq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orgpolicy.Label}
	default:
		err = &NotSingularError{orgpolicy.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (opq *OrgPolicyQuery) OnlyIDX(ctx context.Context) int {
	id, err := opq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrgPolicies.
func (opq *OrgPolicyQuery) All(ctx context.Context) ([]*OrgPolicy, error) {
	ctx = setContextOp(ctx, opq.ctx, "All")
	if err := opq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrgPolicy, *OrgPolicyQuery]()
	return withInterceptors[[]*OrgPolicy](ctx, opq, qr, opq.inters)
}

// AllX is like All, but panics if an error occurs.
func (opq *OrgPolicyQuery) AllX(ctx context.Context) []*OrgPolicy {
	nodes, err := opq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrgPolicy IDs.
func (opq *OrgPolicyQuery) IDs(ctx context.Context) (ids []int, err error) {
	if opq.ctx.Unique == nil && opq.path != nil {
		opq.Unique(true)
	}
	ctx = setContextOp(ctx, opq.ctx, "IDs")
	if err = opq.Select(orgpolicy.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (opq *OrgPolicyQuery) IDsX(ctx context.Context) []int {
	ids, err := opq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (opq *OrgPolicyQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, opq.ctx, "Count")
	if err := opq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, opq, querierCount[*OrgPolicyQuery](), opq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (opq *OrgPolicyQuery) CountX(ctx context.Context) int {
	count, err := opq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (opq *OrgPolicyQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, opq.ctx, "Exist")
	switch _, err := opq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (opq *OrgPolicyQuery) ExistX(ctx context.Context) bool {
	exist, err := opq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrgPolicyQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (opq *OrgPolicyQuery) Clone() *OrgPolicyQuery {
	if opq == nil {
		return nil
	}
	return &OrgPolicyQuery{
		config:          opq.config,
		ctx:             opq.ctx.Clone(),
		order:           append([]orgpolicy.Order{}, opq.order...),
		inters:          append([]Interceptor{}, opq.inters...),
		predicates:      append([]predicate.OrgPolicy{}, opq.predicates...),
		withOrg:         opq.withOrg.Clone(),
		withPermissions: opq.withPermissions.Clone(),
		// clone intermediate query.
		sql:  opq.sql.Clone(),
		path: opq.path,
	}
}

// WithOrg tells the query-builder to eager-load the nodes that are connected to
// the "org" edge. The optional arguments are used to configure the query builder of the edge.
func (opq *OrgPolicyQuery) WithOrg(opts ...func(*OrgQuery)) *OrgPolicyQuery {
	query := (&OrgClient{config: opq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	opq.withOrg = query
	return opq
}

// WithPermissions tells the query-builder to eager-load the nodes that are connected to
// the "permissions" edge. The optional arguments are used to configure the query builder of the edge.
func (opq *OrgPolicyQuery) WithPermissions(opts ...func(*PermissionQuery)) *OrgPolicyQuery {
	query := (&PermissionClient{config: opq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	opq.withPermissions = query
	return opq
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
//	client.OrgPolicy.Query().
//		GroupBy(orgpolicy.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (opq *OrgPolicyQuery) GroupBy(field string, fields ...string) *OrgPolicyGroupBy {
	opq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrgPolicyGroupBy{build: opq}
	grbuild.flds = &opq.ctx.Fields
	grbuild.label = orgpolicy.Label
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
//	client.OrgPolicy.Query().
//		Select(orgpolicy.FieldCreatedBy).
//		Scan(ctx, &v)
func (opq *OrgPolicyQuery) Select(fields ...string) *OrgPolicySelect {
	opq.ctx.Fields = append(opq.ctx.Fields, fields...)
	sbuild := &OrgPolicySelect{OrgPolicyQuery: opq}
	sbuild.label = orgpolicy.Label
	sbuild.flds, sbuild.scan = &opq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrgPolicySelect configured with the given aggregations.
func (opq *OrgPolicyQuery) Aggregate(fns ...AggregateFunc) *OrgPolicySelect {
	return opq.Select().Aggregate(fns...)
}

func (opq *OrgPolicyQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range opq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, opq); err != nil {
				return err
			}
		}
	}
	for _, f := range opq.ctx.Fields {
		if !orgpolicy.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if opq.path != nil {
		prev, err := opq.path(ctx)
		if err != nil {
			return err
		}
		opq.sql = prev
	}
	return nil
}

func (opq *OrgPolicyQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrgPolicy, error) {
	var (
		nodes       = []*OrgPolicy{}
		_spec       = opq.querySpec()
		loadedTypes = [2]bool{
			opq.withOrg != nil,
			opq.withPermissions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrgPolicy).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrgPolicy{config: opq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(opq.modifiers) > 0 {
		_spec.Modifiers = opq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, opq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := opq.withOrg; query != nil {
		if err := opq.loadOrg(ctx, query, nodes, nil,
			func(n *OrgPolicy, e *Org) { n.Edges.Org = e }); err != nil {
			return nil, err
		}
	}
	if query := opq.withPermissions; query != nil {
		if err := opq.loadPermissions(ctx, query, nodes,
			func(n *OrgPolicy) { n.Edges.Permissions = []*Permission{} },
			func(n *OrgPolicy, e *Permission) { n.Edges.Permissions = append(n.Edges.Permissions, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range opq.withNamedPermissions {
		if err := opq.loadPermissions(ctx, query, nodes,
			func(n *OrgPolicy) { n.appendNamedPermissions(name) },
			func(n *OrgPolicy, e *Permission) { n.appendNamedPermissions(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range opq.loadTotal {
		if err := opq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (opq *OrgPolicyQuery) loadOrg(ctx context.Context, query *OrgQuery, nodes []*OrgPolicy, init func(*OrgPolicy), assign func(*OrgPolicy, *Org)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*OrgPolicy)
	for i := range nodes {
		fk := nodes[i].OrgID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(org.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "org_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (opq *OrgPolicyQuery) loadPermissions(ctx context.Context, query *PermissionQuery, nodes []*OrgPolicy, init func(*OrgPolicy), assign func(*OrgPolicy, *Permission)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*OrgPolicy)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.InValues(orgpolicy.PermissionsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrgPolicyID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "org_policy_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (opq *OrgPolicyQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := opq.querySpec()
	if len(opq.modifiers) > 0 {
		_spec.Modifiers = opq.modifiers
	}
	_spec.Node.Columns = opq.ctx.Fields
	if len(opq.ctx.Fields) > 0 {
		_spec.Unique = opq.ctx.Unique != nil && *opq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, opq.driver, _spec)
}

func (opq *OrgPolicyQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orgpolicy.Table, orgpolicy.Columns, sqlgraph.NewFieldSpec(orgpolicy.FieldID, field.TypeInt))
	_spec.From = opq.sql
	if unique := opq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if opq.path != nil {
		_spec.Unique = true
	}
	if fields := opq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgpolicy.FieldID)
		for i := range fields {
			if fields[i] != orgpolicy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if opq.withOrg != nil {
			_spec.Node.AddColumnOnce(orgpolicy.FieldOrgID)
		}
	}
	if ps := opq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := opq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := opq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := opq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (opq *OrgPolicyQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(opq.driver.Dialect())
	t1 := builder.Table(orgpolicy.Table)
	columns := opq.ctx.Fields
	if len(columns) == 0 {
		columns = orgpolicy.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if opq.sql != nil {
		selector = opq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if opq.ctx.Unique != nil && *opq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range opq.predicates {
		p(selector)
	}
	for _, p := range opq.order {
		p(selector)
	}
	if offset := opq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := opq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedPermissions tells the query-builder to eager-load the nodes that are connected to the "permissions"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (opq *OrgPolicyQuery) WithNamedPermissions(name string, opts ...func(*PermissionQuery)) *OrgPolicyQuery {
	query := (&PermissionClient{config: opq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if opq.withNamedPermissions == nil {
		opq.withNamedPermissions = make(map[string]*PermissionQuery)
	}
	opq.withNamedPermissions[name] = query
	return opq
}

// OrgPolicyGroupBy is the group-by builder for OrgPolicy entities.
type OrgPolicyGroupBy struct {
	selector
	build *OrgPolicyQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (opgb *OrgPolicyGroupBy) Aggregate(fns ...AggregateFunc) *OrgPolicyGroupBy {
	opgb.fns = append(opgb.fns, fns...)
	return opgb
}

// Scan applies the selector query and scans the result into the given value.
func (opgb *OrgPolicyGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, opgb.build.ctx, "GroupBy")
	if err := opgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgPolicyQuery, *OrgPolicyGroupBy](ctx, opgb.build, opgb, opgb.build.inters, v)
}

func (opgb *OrgPolicyGroupBy) sqlScan(ctx context.Context, root *OrgPolicyQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(opgb.fns))
	for _, fn := range opgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*opgb.flds)+len(opgb.fns))
		for _, f := range *opgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*opgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := opgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrgPolicySelect is the builder for selecting fields of OrgPolicy entities.
type OrgPolicySelect struct {
	*OrgPolicyQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ops *OrgPolicySelect) Aggregate(fns ...AggregateFunc) *OrgPolicySelect {
	ops.fns = append(ops.fns, fns...)
	return ops
}

// Scan applies the selector query and scans the result into the given value.
func (ops *OrgPolicySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ops.ctx, "Select")
	if err := ops.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgPolicyQuery, *OrgPolicySelect](ctx, ops.OrgPolicyQuery, ops, ops.inters, v)
}

func (ops *OrgPolicySelect) sqlScan(ctx context.Context, root *OrgPolicyQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ops.fns))
	for _, fn := range ops.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ops.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ops.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
