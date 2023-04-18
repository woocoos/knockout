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
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/predicate"
	"github.com/woocoos/knockout/ent/user"
)

// OrgUserQuery is the builder for querying OrgUser entities.
type OrgUserQuery struct {
	config
	ctx               *QueryContext
	order             []orguser.OrderOption
	inters            []Interceptor
	predicates        []predicate.OrgUser
	withOrg           *OrgQuery
	withUser          *UserQuery
	withOrgRoles      *OrgRoleQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*OrgUser) error
	withNamedOrgRoles map[string]*OrgRoleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrgUserQuery builder.
func (ouq *OrgUserQuery) Where(ps ...predicate.OrgUser) *OrgUserQuery {
	ouq.predicates = append(ouq.predicates, ps...)
	return ouq
}

// Limit the number of records to be returned by this query.
func (ouq *OrgUserQuery) Limit(limit int) *OrgUserQuery {
	ouq.ctx.Limit = &limit
	return ouq
}

// Offset to start from.
func (ouq *OrgUserQuery) Offset(offset int) *OrgUserQuery {
	ouq.ctx.Offset = &offset
	return ouq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ouq *OrgUserQuery) Unique(unique bool) *OrgUserQuery {
	ouq.ctx.Unique = &unique
	return ouq
}

// Order specifies how the records should be ordered.
func (ouq *OrgUserQuery) Order(o ...orguser.OrderOption) *OrgUserQuery {
	ouq.order = append(ouq.order, o...)
	return ouq
}

// QueryOrg chains the current query on the "org" edge.
func (ouq *OrgUserQuery) QueryOrg() *OrgQuery {
	query := (&OrgClient{config: ouq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ouq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ouq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orguser.Table, orguser.FieldID, selector),
			sqlgraph.To(org.Table, org.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orguser.OrgTable, orguser.OrgColumn),
		)
		fromU = sqlgraph.SetNeighbors(ouq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (ouq *OrgUserQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ouq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ouq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ouq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orguser.Table, orguser.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, orguser.UserTable, orguser.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ouq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrgRoles chains the current query on the "org_roles" edge.
func (ouq *OrgUserQuery) QueryOrgRoles() *OrgRoleQuery {
	query := (&OrgRoleClient{config: ouq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ouq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ouq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orguser.Table, orguser.FieldID, selector),
			sqlgraph.To(orgrole.Table, orgrole.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, orguser.OrgRolesTable, orguser.OrgRolesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ouq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrgUser entity from the query.
// Returns a *NotFoundError when no OrgUser was found.
func (ouq *OrgUserQuery) First(ctx context.Context) (*OrgUser, error) {
	nodes, err := ouq.Limit(1).All(setContextOp(ctx, ouq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orguser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ouq *OrgUserQuery) FirstX(ctx context.Context) *OrgUser {
	node, err := ouq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrgUser ID from the query.
// Returns a *NotFoundError when no OrgUser ID was found.
func (ouq *OrgUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ouq.Limit(1).IDs(setContextOp(ctx, ouq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orguser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ouq *OrgUserQuery) FirstIDX(ctx context.Context) int {
	id, err := ouq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrgUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrgUser entity is found.
// Returns a *NotFoundError when no OrgUser entities are found.
func (ouq *OrgUserQuery) Only(ctx context.Context) (*OrgUser, error) {
	nodes, err := ouq.Limit(2).All(setContextOp(ctx, ouq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orguser.Label}
	default:
		return nil, &NotSingularError{orguser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ouq *OrgUserQuery) OnlyX(ctx context.Context) *OrgUser {
	node, err := ouq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrgUser ID in the query.
// Returns a *NotSingularError when more than one OrgUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (ouq *OrgUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ouq.Limit(2).IDs(setContextOp(ctx, ouq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orguser.Label}
	default:
		err = &NotSingularError{orguser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ouq *OrgUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := ouq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrgUsers.
func (ouq *OrgUserQuery) All(ctx context.Context) ([]*OrgUser, error) {
	ctx = setContextOp(ctx, ouq.ctx, "All")
	if err := ouq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrgUser, *OrgUserQuery]()
	return withInterceptors[[]*OrgUser](ctx, ouq, qr, ouq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ouq *OrgUserQuery) AllX(ctx context.Context) []*OrgUser {
	nodes, err := ouq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrgUser IDs.
func (ouq *OrgUserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ouq.ctx.Unique == nil && ouq.path != nil {
		ouq.Unique(true)
	}
	ctx = setContextOp(ctx, ouq.ctx, "IDs")
	if err = ouq.Select(orguser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ouq *OrgUserQuery) IDsX(ctx context.Context) []int {
	ids, err := ouq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ouq *OrgUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ouq.ctx, "Count")
	if err := ouq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ouq, querierCount[*OrgUserQuery](), ouq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ouq *OrgUserQuery) CountX(ctx context.Context) int {
	count, err := ouq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ouq *OrgUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ouq.ctx, "Exist")
	switch _, err := ouq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ouq *OrgUserQuery) ExistX(ctx context.Context) bool {
	exist, err := ouq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrgUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ouq *OrgUserQuery) Clone() *OrgUserQuery {
	if ouq == nil {
		return nil
	}
	return &OrgUserQuery{
		config:       ouq.config,
		ctx:          ouq.ctx.Clone(),
		order:        append([]orguser.OrderOption{}, ouq.order...),
		inters:       append([]Interceptor{}, ouq.inters...),
		predicates:   append([]predicate.OrgUser{}, ouq.predicates...),
		withOrg:      ouq.withOrg.Clone(),
		withUser:     ouq.withUser.Clone(),
		withOrgRoles: ouq.withOrgRoles.Clone(),
		// clone intermediate query.
		sql:  ouq.sql.Clone(),
		path: ouq.path,
	}
}

// WithOrg tells the query-builder to eager-load the nodes that are connected to
// the "org" edge. The optional arguments are used to configure the query builder of the edge.
func (ouq *OrgUserQuery) WithOrg(opts ...func(*OrgQuery)) *OrgUserQuery {
	query := (&OrgClient{config: ouq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ouq.withOrg = query
	return ouq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ouq *OrgUserQuery) WithUser(opts ...func(*UserQuery)) *OrgUserQuery {
	query := (&UserClient{config: ouq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ouq.withUser = query
	return ouq
}

// WithOrgRoles tells the query-builder to eager-load the nodes that are connected to
// the "org_roles" edge. The optional arguments are used to configure the query builder of the edge.
func (ouq *OrgUserQuery) WithOrgRoles(opts ...func(*OrgRoleQuery)) *OrgUserQuery {
	query := (&OrgRoleClient{config: ouq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ouq.withOrgRoles = query
	return ouq
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
//	client.OrgUser.Query().
//		GroupBy(orguser.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ouq *OrgUserQuery) GroupBy(field string, fields ...string) *OrgUserGroupBy {
	ouq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrgUserGroupBy{build: ouq}
	grbuild.flds = &ouq.ctx.Fields
	grbuild.label = orguser.Label
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
//	client.OrgUser.Query().
//		Select(orguser.FieldCreatedBy).
//		Scan(ctx, &v)
func (ouq *OrgUserQuery) Select(fields ...string) *OrgUserSelect {
	ouq.ctx.Fields = append(ouq.ctx.Fields, fields...)
	sbuild := &OrgUserSelect{OrgUserQuery: ouq}
	sbuild.label = orguser.Label
	sbuild.flds, sbuild.scan = &ouq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrgUserSelect configured with the given aggregations.
func (ouq *OrgUserQuery) Aggregate(fns ...AggregateFunc) *OrgUserSelect {
	return ouq.Select().Aggregate(fns...)
}

func (ouq *OrgUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ouq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ouq); err != nil {
				return err
			}
		}
	}
	for _, f := range ouq.ctx.Fields {
		if !orguser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ouq.path != nil {
		prev, err := ouq.path(ctx)
		if err != nil {
			return err
		}
		ouq.sql = prev
	}
	return nil
}

func (ouq *OrgUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrgUser, error) {
	var (
		nodes       = []*OrgUser{}
		_spec       = ouq.querySpec()
		loadedTypes = [3]bool{
			ouq.withOrg != nil,
			ouq.withUser != nil,
			ouq.withOrgRoles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrgUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrgUser{config: ouq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(ouq.modifiers) > 0 {
		_spec.Modifiers = ouq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ouq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ouq.withOrg; query != nil {
		if err := ouq.loadOrg(ctx, query, nodes, nil,
			func(n *OrgUser, e *Org) { n.Edges.Org = e }); err != nil {
			return nil, err
		}
	}
	if query := ouq.withUser; query != nil {
		if err := ouq.loadUser(ctx, query, nodes, nil,
			func(n *OrgUser, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := ouq.withOrgRoles; query != nil {
		if err := ouq.loadOrgRoles(ctx, query, nodes,
			func(n *OrgUser) { n.Edges.OrgRoles = []*OrgRole{} },
			func(n *OrgUser, e *OrgRole) { n.Edges.OrgRoles = append(n.Edges.OrgRoles, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range ouq.withNamedOrgRoles {
		if err := ouq.loadOrgRoles(ctx, query, nodes,
			func(n *OrgUser) { n.appendNamedOrgRoles(name) },
			func(n *OrgUser, e *OrgRole) { n.appendNamedOrgRoles(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range ouq.loadTotal {
		if err := ouq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ouq *OrgUserQuery) loadOrg(ctx context.Context, query *OrgQuery, nodes []*OrgUser, init func(*OrgUser), assign func(*OrgUser, *Org)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*OrgUser)
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
func (ouq *OrgUserQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*OrgUser, init func(*OrgUser), assign func(*OrgUser, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*OrgUser)
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
func (ouq *OrgUserQuery) loadOrgRoles(ctx context.Context, query *OrgRoleQuery, nodes []*OrgUser, init func(*OrgUser), assign func(*OrgUser, *OrgRole)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*OrgUser)
	nids := make(map[int]map[*OrgUser]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(orguser.OrgRolesTable)
		s.Join(joinT).On(s.C(orgrole.FieldID), joinT.C(orguser.OrgRolesPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(orguser.OrgRolesPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(orguser.OrgRolesPrimaryKey[1]))
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
					nids[inValue] = map[*OrgUser]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*OrgRole](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "org_roles" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (ouq *OrgUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ouq.querySpec()
	if len(ouq.modifiers) > 0 {
		_spec.Modifiers = ouq.modifiers
	}
	_spec.Node.Columns = ouq.ctx.Fields
	if len(ouq.ctx.Fields) > 0 {
		_spec.Unique = ouq.ctx.Unique != nil && *ouq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ouq.driver, _spec)
}

func (ouq *OrgUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orguser.Table, orguser.Columns, sqlgraph.NewFieldSpec(orguser.FieldID, field.TypeInt))
	_spec.From = ouq.sql
	if unique := ouq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ouq.path != nil {
		_spec.Unique = true
	}
	if fields := ouq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orguser.FieldID)
		for i := range fields {
			if fields[i] != orguser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ouq.withOrg != nil {
			_spec.Node.AddColumnOnce(orguser.FieldOrgID)
		}
		if ouq.withUser != nil {
			_spec.Node.AddColumnOnce(orguser.FieldUserID)
		}
	}
	if ps := ouq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ouq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ouq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ouq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ouq *OrgUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ouq.driver.Dialect())
	t1 := builder.Table(orguser.Table)
	columns := ouq.ctx.Fields
	if len(columns) == 0 {
		columns = orguser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ouq.sql != nil {
		selector = ouq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ouq.ctx.Unique != nil && *ouq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ouq.predicates {
		p(selector)
	}
	for _, p := range ouq.order {
		p(selector)
	}
	if offset := ouq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ouq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedOrgRoles tells the query-builder to eager-load the nodes that are connected to the "org_roles"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (ouq *OrgUserQuery) WithNamedOrgRoles(name string, opts ...func(*OrgRoleQuery)) *OrgUserQuery {
	query := (&OrgRoleClient{config: ouq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if ouq.withNamedOrgRoles == nil {
		ouq.withNamedOrgRoles = make(map[string]*OrgRoleQuery)
	}
	ouq.withNamedOrgRoles[name] = query
	return ouq
}

// OrgUserGroupBy is the group-by builder for OrgUser entities.
type OrgUserGroupBy struct {
	selector
	build *OrgUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ougb *OrgUserGroupBy) Aggregate(fns ...AggregateFunc) *OrgUserGroupBy {
	ougb.fns = append(ougb.fns, fns...)
	return ougb
}

// Scan applies the selector query and scans the result into the given value.
func (ougb *OrgUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ougb.build.ctx, "GroupBy")
	if err := ougb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgUserQuery, *OrgUserGroupBy](ctx, ougb.build, ougb, ougb.build.inters, v)
}

func (ougb *OrgUserGroupBy) sqlScan(ctx context.Context, root *OrgUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ougb.fns))
	for _, fn := range ougb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ougb.flds)+len(ougb.fns))
		for _, f := range *ougb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ougb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ougb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrgUserSelect is the builder for selecting fields of OrgUser entities.
type OrgUserSelect struct {
	*OrgUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ous *OrgUserSelect) Aggregate(fns ...AggregateFunc) *OrgUserSelect {
	ous.fns = append(ous.fns, fns...)
	return ous
}

// Scan applies the selector query and scans the result into the given value.
func (ous *OrgUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ous.ctx, "Select")
	if err := ous.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrgUserQuery, *OrgUserSelect](ctx, ous.OrgUserQuery, ous, ous.inters, v)
}

func (ous *OrgUserSelect) sqlScan(ctx context.Context, root *OrgUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ous.fns))
	for _, fn := range ous.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ous.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ous.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
