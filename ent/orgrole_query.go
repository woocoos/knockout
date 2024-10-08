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
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orgroleuser"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/predicate"
)

// OrgRoleQuery is the builder for querying OrgRole entities.
type OrgRoleQuery struct {
	config
	ctx                  *QueryContext
	order                []orgrole.OrderOption
	inters               []Interceptor
	predicates           []predicate.OrgRole
	withOrg              *OrgQuery
	withOrgUsers         *OrgUserQuery
	withOrgRoleUser      *OrgRoleUserQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*OrgRole) error
	withNamedOrgUsers    map[string]*OrgUserQuery
	withNamedOrgRoleUser map[string]*OrgRoleUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrgRoleQuery builder.
func (orq *OrgRoleQuery) Where(ps ...predicate.OrgRole) *OrgRoleQuery {
	orq.predicates = append(orq.predicates, ps...)
	return orq
}

// Limit the number of records to be returned by this query.
func (orq *OrgRoleQuery) Limit(limit int) *OrgRoleQuery {
	orq.ctx.Limit = &limit
	return orq
}

// Offset to start from.
func (orq *OrgRoleQuery) Offset(offset int) *OrgRoleQuery {
	orq.ctx.Offset = &offset
	return orq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (orq *OrgRoleQuery) Unique(unique bool) *OrgRoleQuery {
	orq.ctx.Unique = &unique
	return orq
}

// Order specifies how the records should be ordered.
func (orq *OrgRoleQuery) Order(o ...orgrole.OrderOption) *OrgRoleQuery {
	orq.order = append(orq.order, o...)
	return orq
}

// QueryOrg chains the current query on the "org" edge.
func (orq *OrgRoleQuery) QueryOrg() *OrgQuery {
	query := (&OrgClient{config: orq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := orq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := orq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgrole.Table, orgrole.FieldID, selector),
			sqlgraph.To(org.Table, org.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orgrole.OrgTable, orgrole.OrgColumn),
		)
		fromU = sqlgraph.SetNeighbors(orq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrgUsers chains the current query on the "org_users" edge.
func (orq *OrgRoleQuery) QueryOrgUsers() *OrgUserQuery {
	query := (&OrgUserClient{config: orq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := orq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := orq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgrole.Table, orgrole.FieldID, selector),
			sqlgraph.To(orguser.Table, orguser.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, orgrole.OrgUsersTable, orgrole.OrgUsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(orq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrgRoleUser chains the current query on the "org_role_user" edge.
func (orq *OrgRoleQuery) QueryOrgRoleUser() *OrgRoleUserQuery {
	query := (&OrgRoleUserClient{config: orq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := orq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := orq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orgrole.Table, orgrole.FieldID, selector),
			sqlgraph.To(orgroleuser.Table, orgroleuser.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, orgrole.OrgRoleUserTable, orgrole.OrgRoleUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(orq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrgRole entity from the query.
// Returns a *NotFoundError when no OrgRole was found.
func (orq *OrgRoleQuery) First(ctx context.Context) (*OrgRole, error) {
	nodes, err := orq.Limit(1).All(setContextOp(ctx, orq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orgrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (orq *OrgRoleQuery) FirstX(ctx context.Context) *OrgRole {
	node, err := orq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrgRole ID from the query.
// Returns a *NotFoundError when no OrgRole ID was found.
func (orq *OrgRoleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = orq.Limit(1).IDs(setContextOp(ctx, orq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orgrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (orq *OrgRoleQuery) FirstIDX(ctx context.Context) int {
	id, err := orq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrgRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrgRole entity is found.
// Returns a *NotFoundError when no OrgRole entities are found.
func (orq *OrgRoleQuery) Only(ctx context.Context) (*OrgRole, error) {
	nodes, err := orq.Limit(2).All(setContextOp(ctx, orq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orgrole.Label}
	default:
		return nil, &NotSingularError{orgrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (orq *OrgRoleQuery) OnlyX(ctx context.Context) *OrgRole {
	node, err := orq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrgRole ID in the query.
// Returns a *NotSingularError when more than one OrgRole ID is found.
// Returns a *NotFoundError when no entities are found.
func (orq *OrgRoleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = orq.Limit(2).IDs(setContextOp(ctx, orq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orgrole.Label}
	default:
		err = &NotSingularError{orgrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (orq *OrgRoleQuery) OnlyIDX(ctx context.Context) int {
	id, err := orq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrgRoles.
func (orq *OrgRoleQuery) All(ctx context.Context) ([]*OrgRole, error) {
	ctx = setContextOp(ctx, orq.ctx, ent.OpQueryAll)
	if err := orq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrgRole, *OrgRoleQuery]()
	return withInterceptors[[]*OrgRole](ctx, orq, qr, orq.inters)
}

// AllX is like All, but panics if an error occurs.
func (orq *OrgRoleQuery) AllX(ctx context.Context) []*OrgRole {
	nodes, err := orq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrgRole IDs.
func (orq *OrgRoleQuery) IDs(ctx context.Context) (ids []int, err error) {
	if orq.ctx.Unique == nil && orq.path != nil {
		orq.Unique(true)
	}
	ctx = setContextOp(ctx, orq.ctx, ent.OpQueryIDs)
	if err = orq.Select(orgrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (orq *OrgRoleQuery) IDsX(ctx context.Context) []int {
	ids, err := orq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (orq *OrgRoleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, orq.ctx, ent.OpQueryCount)
	if err := orq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, orq, querierCount[*OrgRoleQuery](), orq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (orq *OrgRoleQuery) CountX(ctx context.Context) int {
	count, err := orq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (orq *OrgRoleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, orq.ctx, ent.OpQueryExist)
	switch _, err := orq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (orq *OrgRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := orq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrgRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (orq *OrgRoleQuery) Clone() *OrgRoleQuery {
	if orq == nil {
		return nil
	}
	return &OrgRoleQuery{
		config:          orq.config,
		ctx:             orq.ctx.Clone(),
		order:           append([]orgrole.OrderOption{}, orq.order...),
		inters:          append([]Interceptor{}, orq.inters...),
		predicates:      append([]predicate.OrgRole{}, orq.predicates...),
		withOrg:         orq.withOrg.Clone(),
		withOrgUsers:    orq.withOrgUsers.Clone(),
		withOrgRoleUser: orq.withOrgRoleUser.Clone(),
		// clone intermediate query.
		sql:  orq.sql.Clone(),
		path: orq.path,
	}
}

// WithOrg tells the query-builder to eager-load the nodes that are connected to
// the "org" edge. The optional arguments are used to configure the query builder of the edge.
func (orq *OrgRoleQuery) WithOrg(opts ...func(*OrgQuery)) *OrgRoleQuery {
	query := (&OrgClient{config: orq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	orq.withOrg = query
	return orq
}

// WithOrgUsers tells the query-builder to eager-load the nodes that are connected to
// the "org_users" edge. The optional arguments are used to configure the query builder of the edge.
func (orq *OrgRoleQuery) WithOrgUsers(opts ...func(*OrgUserQuery)) *OrgRoleQuery {
	query := (&OrgUserClient{config: orq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	orq.withOrgUsers = query
	return orq
}

// WithOrgRoleUser tells the query-builder to eager-load the nodes that are connected to
// the "org_role_user" edge. The optional arguments are used to configure the query builder of the edge.
func (orq *OrgRoleQuery) WithOrgRoleUser(opts ...func(*OrgRoleUserQuery)) *OrgRoleQuery {
	query := (&OrgRoleUserClient{config: orq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	orq.withOrgRoleUser = query
	return orq
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
//	client.OrgRole.Query().
//		GroupBy(orgrole.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (orq *OrgRoleQuery) GroupBy(field string, fields ...string) *OrgRoleGroupBy {
	orq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrgRoleGroupBy{build: orq}
	grbuild.flds = &orq.ctx.Fields
	grbuild.label = orgrole.Label
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
//	client.OrgRole.Query().
//		Select(orgrole.FieldCreatedBy).
//		Scan(ctx, &v)
func (orq *OrgRoleQuery) Select(fields ...string) *OrgRoleSelect {
	orq.ctx.Fields = append(orq.ctx.Fields, fields...)
	sbuild := &OrgRoleSelect{OrgRoleQuery: orq}
	sbuild.label = orgrole.Label
	sbuild.flds, sbuild.scan = &orq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrgRoleSelect configured with the given aggregations.
func (orq *OrgRoleQuery) Aggregate(fns ...AggregateFunc) *OrgRoleSelect {
	return orq.Select().Aggregate(fns...)
}

func (orq *OrgRoleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range orq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, orq); err != nil {
				return err
			}
		}
	}
	for _, f := range orq.ctx.Fields {
		if !orgrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if orq.path != nil {
		prev, err := orq.path(ctx)
		if err != nil {
			return err
		}
		orq.sql = prev
	}
	return nil
}

func (orq *OrgRoleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrgRole, error) {
	var (
		nodes       = []*OrgRole{}
		_spec       = orq.querySpec()
		loadedTypes = [3]bool{
			orq.withOrg != nil,
			orq.withOrgUsers != nil,
			orq.withOrgRoleUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrgRole).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrgRole{config: orq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(orq.modifiers) > 0 {
		_spec.Modifiers = orq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, orq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := orq.withOrg; query != nil {
		if err := orq.loadOrg(ctx, query, nodes, nil,
			func(n *OrgRole, e *Org) { n.Edges.Org = e }); err != nil {
			return nil, err
		}
	}
	if query := orq.withOrgUsers; query != nil {
		if err := orq.loadOrgUsers(ctx, query, nodes,
			func(n *OrgRole) { n.Edges.OrgUsers = []*OrgUser{} },
			func(n *OrgRole, e *OrgUser) { n.Edges.OrgUsers = append(n.Edges.OrgUsers, e) }); err != nil {
			return nil, err
		}
	}
	if query := orq.withOrgRoleUser; query != nil {
		if err := orq.loadOrgRoleUser(ctx, query, nodes,
			func(n *OrgRole) { n.Edges.OrgRoleUser = []*OrgRoleUser{} },
			func(n *OrgRole, e *OrgRoleUser) { n.Edges.OrgRoleUser = append(n.Edges.OrgRoleUser, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range orq.withNamedOrgUsers {
		if err := orq.loadOrgUsers(ctx, query, nodes,
			func(n *OrgRole) { n.appendNamedOrgUsers(name) },
			func(n *OrgRole, e *OrgUser) { n.appendNamedOrgUsers(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range orq.withNamedOrgRoleUser {
		if err := orq.loadOrgRoleUser(ctx, query, nodes,
			func(n *OrgRole) { n.appendNamedOrgRoleUser(name) },
			func(n *OrgRole, e *OrgRoleUser) { n.appendNamedOrgRoleUser(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range orq.loadTotal {
		if err := orq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (orq *OrgRoleQuery) loadOrg(ctx context.Context, query *OrgQuery, nodes []*OrgRole, init func(*OrgRole), assign func(*OrgRole, *Org)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*OrgRole)
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
func (orq *OrgRoleQuery) loadOrgUsers(ctx context.Context, query *OrgUserQuery, nodes []*OrgRole, init func(*OrgRole), assign func(*OrgRole, *OrgUser)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*OrgRole)
	nids := make(map[int]map[*OrgRole]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(orgrole.OrgUsersTable)
		s.Join(joinT).On(s.C(orguser.FieldID), joinT.C(orgrole.OrgUsersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(orgrole.OrgUsersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(orgrole.OrgUsersPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*OrgRole]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*OrgUser](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "org_users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (orq *OrgRoleQuery) loadOrgRoleUser(ctx context.Context, query *OrgRoleUserQuery, nodes []*OrgRole, init func(*OrgRole), assign func(*OrgRole, *OrgRoleUser)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*OrgRole)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(orgroleuser.FieldOrgRoleID)
	}
	query.Where(predicate.OrgRoleUser(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(orgrole.OrgRoleUserColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrgRoleID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "org_role_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (orq *OrgRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := orq.querySpec()
	if len(orq.modifiers) > 0 {
		_spec.Modifiers = orq.modifiers
	}
	_spec.Node.Columns = orq.ctx.Fields
	if len(orq.ctx.Fields) > 0 {
		_spec.Unique = orq.ctx.Unique != nil && *orq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, orq.driver, _spec)
}

func (orq *OrgRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orgrole.Table, orgrole.Columns, sqlgraph.NewFieldSpec(orgrole.FieldID, field.TypeInt))
	_spec.From = orq.sql
	if unique := orq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if orq.path != nil {
		_spec.Unique = true
	}
	if fields := orq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgrole.FieldID)
		for i := range fields {
			if fields[i] != orgrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if orq.withOrg != nil {
			_spec.Node.AddColumnOnce(orgrole.FieldOrgID)
		}
	}
	if ps := orq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := orq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := orq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := orq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (orq *OrgRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(orq.driver.Dialect())
	t1 := builder.Table(orgrole.Table)
	columns := orq.ctx.Fields
	if len(columns) == 0 {
		columns = orgrole.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if orq.sql != nil {
		selector = orq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if orq.ctx.Unique != nil && *orq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range orq.predicates {
		p(selector)
	}
	for _, p := range orq.order {
		p(selector)
	}
	if offset := orq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := orq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedOrgUsers tells the query-builder to eager-load the nodes that are connected to the "org_users"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (orq *OrgRoleQuery) WithNamedOrgUsers(name string, opts ...func(*OrgUserQuery)) *OrgRoleQuery {
	query := (&OrgUserClient{config: orq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if orq.withNamedOrgUsers == nil {
		orq.withNamedOrgUsers = make(map[string]*OrgUserQuery)
	}
	orq.withNamedOrgUsers[name] = query
	return orq
}

// WithNamedOrgRoleUser tells the query-builder to eager-load the nodes that are connected to the "org_role_user"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (orq *OrgRoleQuery) WithNamedOrgRoleUser(name string, opts ...func(*OrgRoleUserQuery)) *OrgRoleQuery {
	query := (&OrgRoleUserClient{config: orq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if orq.withNamedOrgRoleUser == nil {
		orq.withNamedOrgRoleUser = make(map[string]*OrgRoleUserQuery)
	}
	orq.withNamedOrgRoleUser[name] = query
	return orq
}

// OrgRoleGroupBy is the group-by builder for OrgRole entities.
type OrgRoleGroupBy struct {
	selector
	build *OrgRoleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (orgb *OrgRoleGroupBy) Aggregate(fns ...AggregateFunc) *OrgRoleGroupBy {
	orgb.fns = append(orgb.fns, fns...)
	return orgb
}

// Scan applies the selector query and scans the result into the given value.
func (orgb *OrgRoleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, orgb.build.ctx, ent.OpQueryGroupBy)
	if err := orgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgRoleQuery, *OrgRoleGroupBy](ctx, orgb.build, orgb, orgb.build.inters, v)
}

func (orgb *OrgRoleGroupBy) sqlScan(ctx context.Context, root *OrgRoleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(orgb.fns))
	for _, fn := range orgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*orgb.flds)+len(orgb.fns))
		for _, f := range *orgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*orgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := orgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrgRoleSelect is the builder for selecting fields of OrgRole entities.
type OrgRoleSelect struct {
	*OrgRoleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ors *OrgRoleSelect) Aggregate(fns ...AggregateFunc) *OrgRoleSelect {
	ors.fns = append(ors.fns, fns...)
	return ors
}

// Scan applies the selector query and scans the result into the given value.
func (ors *OrgRoleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ors.ctx, ent.OpQuerySelect)
	if err := ors.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgRoleQuery, *OrgRoleSelect](ctx, ors.OrgRoleQuery, ors, ors.inters, v)
}

func (ors *OrgRoleSelect) sqlScan(ctx context.Context, root *OrgRoleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ors.fns))
	for _, fn := range ors.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ors.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ors.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
