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
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/predicate"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/userdevice"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	ctx                  *QueryContext
	order                []user.OrderOption
	inters               []Interceptor
	predicates           []predicate.User
	withIdentities       *UserIdentityQuery
	withLoginProfile     *UserLoginProfileQuery
	withPasswords        *UserPasswordQuery
	withDevices          *UserDeviceQuery
	withOrgs             *OrgQuery
	withPermissions      *PermissionQuery
	withOrgUser          *OrgUserQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*User) error
	withNamedIdentities  map[string]*UserIdentityQuery
	withNamedPasswords   map[string]*UserPasswordQuery
	withNamedDevices     map[string]*UserDeviceQuery
	withNamedOrgs        map[string]*OrgQuery
	withNamedPermissions map[string]*PermissionQuery
	withNamedOrgUser     map[string]*OrgUserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UserQuery) Unique(unique bool) *UserQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UserQuery) Order(o ...user.OrderOption) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryIdentities chains the current query on the "identities" edge.
func (uq *UserQuery) QueryIdentities() *UserIdentityQuery {
	query := (&UserIdentityClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(useridentity.Table, useridentity.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.IdentitiesTable, user.IdentitiesColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLoginProfile chains the current query on the "login_profile" edge.
func (uq *UserQuery) QueryLoginProfile() *UserLoginProfileQuery {
	query := (&UserLoginProfileClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(userloginprofile.Table, userloginprofile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.LoginProfileTable, user.LoginProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPasswords chains the current query on the "passwords" edge.
func (uq *UserQuery) QueryPasswords() *UserPasswordQuery {
	query := (&UserPasswordClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(userpassword.Table, userpassword.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PasswordsTable, user.PasswordsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDevices chains the current query on the "devices" edge.
func (uq *UserQuery) QueryDevices() *UserDeviceQuery {
	query := (&UserDeviceClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(userdevice.Table, userdevice.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.DevicesTable, user.DevicesColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrgs chains the current query on the "orgs" edge.
func (uq *UserQuery) QueryOrgs() *OrgQuery {
	query := (&OrgClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(org.Table, org.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.OrgsTable, user.OrgsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPermissions chains the current query on the "permissions" edge.
func (uq *UserQuery) QueryPermissions() *PermissionQuery {
	query := (&PermissionClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.PermissionsTable, user.PermissionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrgUser chains the current query on the "org_user" edge.
func (uq *UserQuery) QueryOrgUser() *OrgUserQuery {
	query := (&OrgUserClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(orguser.Table, orguser.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.OrgUserTable, user.OrgUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstIDX(ctx context.Context) int {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyIDX(ctx context.Context) int {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	ctx = setContextOp(ctx, uq.ctx, "All")
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*User, *UserQuery]()
	return withInterceptors[[]*User](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (uq *UserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, "IDs")
	if err = uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []int {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, "Count")
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UserQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, "Exist")
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	if uq == nil {
		return nil
	}
	return &UserQuery{
		config:           uq.config,
		ctx:              uq.ctx.Clone(),
		order:            append([]user.OrderOption{}, uq.order...),
		inters:           append([]Interceptor{}, uq.inters...),
		predicates:       append([]predicate.User{}, uq.predicates...),
		withIdentities:   uq.withIdentities.Clone(),
		withLoginProfile: uq.withLoginProfile.Clone(),
		withPasswords:    uq.withPasswords.Clone(),
		withDevices:      uq.withDevices.Clone(),
		withOrgs:         uq.withOrgs.Clone(),
		withPermissions:  uq.withPermissions.Clone(),
		withOrgUser:      uq.withOrgUser.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithIdentities tells the query-builder to eager-load the nodes that are connected to
// the "identities" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithIdentities(opts ...func(*UserIdentityQuery)) *UserQuery {
	query := (&UserIdentityClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withIdentities = query
	return uq
}

// WithLoginProfile tells the query-builder to eager-load the nodes that are connected to
// the "login_profile" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithLoginProfile(opts ...func(*UserLoginProfileQuery)) *UserQuery {
	query := (&UserLoginProfileClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withLoginProfile = query
	return uq
}

// WithPasswords tells the query-builder to eager-load the nodes that are connected to
// the "passwords" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithPasswords(opts ...func(*UserPasswordQuery)) *UserQuery {
	query := (&UserPasswordClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withPasswords = query
	return uq
}

// WithDevices tells the query-builder to eager-load the nodes that are connected to
// the "devices" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithDevices(opts ...func(*UserDeviceQuery)) *UserQuery {
	query := (&UserDeviceClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withDevices = query
	return uq
}

// WithOrgs tells the query-builder to eager-load the nodes that are connected to
// the "orgs" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithOrgs(opts ...func(*OrgQuery)) *UserQuery {
	query := (&OrgClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withOrgs = query
	return uq
}

// WithPermissions tells the query-builder to eager-load the nodes that are connected to
// the "permissions" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithPermissions(opts ...func(*PermissionQuery)) *UserQuery {
	query := (&PermissionClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withPermissions = query
	return uq
}

// WithOrgUser tells the query-builder to eager-load the nodes that are connected to
// the "org_user" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithOrgUser(opts ...func(*OrgUserQuery)) *UserQuery {
	query := (&OrgUserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withOrgUser = query
	return uq
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
//	client.User.Query().
//		GroupBy(user.FieldCreatedBy).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = user.Label
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
//	client.User.Query().
//		Select(user.FieldCreatedBy).
//		Scan(ctx, &v)
func (uq *UserQuery) Select(fields ...string) *UserSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UserSelect{UserQuery: uq}
	sbuild.label = user.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (uq *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !user.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	var (
		nodes       = []*User{}
		_spec       = uq.querySpec()
		loadedTypes = [7]bool{
			uq.withIdentities != nil,
			uq.withLoginProfile != nil,
			uq.withPasswords != nil,
			uq.withDevices != nil,
			uq.withOrgs != nil,
			uq.withPermissions != nil,
			uq.withOrgUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withIdentities; query != nil {
		if err := uq.loadIdentities(ctx, query, nodes,
			func(n *User) { n.Edges.Identities = []*UserIdentity{} },
			func(n *User, e *UserIdentity) { n.Edges.Identities = append(n.Edges.Identities, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withLoginProfile; query != nil {
		if err := uq.loadLoginProfile(ctx, query, nodes, nil,
			func(n *User, e *UserLoginProfile) { n.Edges.LoginProfile = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withPasswords; query != nil {
		if err := uq.loadPasswords(ctx, query, nodes,
			func(n *User) { n.Edges.Passwords = []*UserPassword{} },
			func(n *User, e *UserPassword) { n.Edges.Passwords = append(n.Edges.Passwords, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withDevices; query != nil {
		if err := uq.loadDevices(ctx, query, nodes,
			func(n *User) { n.Edges.Devices = []*UserDevice{} },
			func(n *User, e *UserDevice) { n.Edges.Devices = append(n.Edges.Devices, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withOrgs; query != nil {
		if err := uq.loadOrgs(ctx, query, nodes,
			func(n *User) { n.Edges.Orgs = []*Org{} },
			func(n *User, e *Org) { n.Edges.Orgs = append(n.Edges.Orgs, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withPermissions; query != nil {
		if err := uq.loadPermissions(ctx, query, nodes,
			func(n *User) { n.Edges.Permissions = []*Permission{} },
			func(n *User, e *Permission) { n.Edges.Permissions = append(n.Edges.Permissions, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withOrgUser; query != nil {
		if err := uq.loadOrgUser(ctx, query, nodes,
			func(n *User) { n.Edges.OrgUser = []*OrgUser{} },
			func(n *User, e *OrgUser) { n.Edges.OrgUser = append(n.Edges.OrgUser, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedIdentities {
		if err := uq.loadIdentities(ctx, query, nodes,
			func(n *User) { n.appendNamedIdentities(name) },
			func(n *User, e *UserIdentity) { n.appendNamedIdentities(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedPasswords {
		if err := uq.loadPasswords(ctx, query, nodes,
			func(n *User) { n.appendNamedPasswords(name) },
			func(n *User, e *UserPassword) { n.appendNamedPasswords(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedDevices {
		if err := uq.loadDevices(ctx, query, nodes,
			func(n *User) { n.appendNamedDevices(name) },
			func(n *User, e *UserDevice) { n.appendNamedDevices(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedOrgs {
		if err := uq.loadOrgs(ctx, query, nodes,
			func(n *User) { n.appendNamedOrgs(name) },
			func(n *User, e *Org) { n.appendNamedOrgs(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedPermissions {
		if err := uq.loadPermissions(ctx, query, nodes,
			func(n *User) { n.appendNamedPermissions(name) },
			func(n *User, e *Permission) { n.appendNamedPermissions(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range uq.withNamedOrgUser {
		if err := uq.loadOrgUser(ctx, query, nodes,
			func(n *User) { n.appendNamedOrgUser(name) },
			func(n *User, e *OrgUser) { n.appendNamedOrgUser(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range uq.loadTotal {
		if err := uq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UserQuery) loadIdentities(ctx context.Context, query *UserIdentityQuery, nodes []*User, init func(*User), assign func(*User, *UserIdentity)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(useridentity.FieldUserID)
	}
	query.Where(predicate.UserIdentity(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.IdentitiesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadLoginProfile(ctx context.Context, query *UserLoginProfileQuery, nodes []*User, init func(*User), assign func(*User, *UserLoginProfile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(userloginprofile.FieldUserID)
	}
	query.Where(predicate.UserLoginProfile(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.LoginProfileColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadPasswords(ctx context.Context, query *UserPasswordQuery, nodes []*User, init func(*User), assign func(*User, *UserPassword)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(userpassword.FieldUserID)
	}
	query.Where(predicate.UserPassword(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.PasswordsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadDevices(ctx context.Context, query *UserDeviceQuery, nodes []*User, init func(*User), assign func(*User, *UserDevice)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(userdevice.FieldUserID)
	}
	query.Where(predicate.UserDevice(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.DevicesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadOrgs(ctx context.Context, query *OrgQuery, nodes []*User, init func(*User), assign func(*User, *Org)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*User)
	nids := make(map[int]map[*User]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(user.OrgsTable)
		s.Join(joinT).On(s.C(org.FieldID), joinT.C(user.OrgsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(user.OrgsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(user.OrgsPrimaryKey[1]))
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
					nids[inValue] = map[*User]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Org](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "orgs" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (uq *UserQuery) loadPermissions(ctx context.Context, query *PermissionQuery, nodes []*User, init func(*User), assign func(*User, *Permission)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(permission.FieldUserID)
	}
	query.Where(predicate.Permission(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.PermissionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadOrgUser(ctx context.Context, query *OrgUserQuery, nodes []*User, init func(*User), assign func(*User, *OrgUser)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(orguser.FieldUserID)
	}
	query.Where(predicate.OrgUser(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.OrgUserColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uq *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	if len(uq.modifiers) > 0 {
		_spec.Modifiers = uq.modifiers
	}
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for i := range fields {
			if fields[i] != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(user.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = user.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedIdentities tells the query-builder to eager-load the nodes that are connected to the "identities"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedIdentities(name string, opts ...func(*UserIdentityQuery)) *UserQuery {
	query := (&UserIdentityClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedIdentities == nil {
		uq.withNamedIdentities = make(map[string]*UserIdentityQuery)
	}
	uq.withNamedIdentities[name] = query
	return uq
}

// WithNamedPasswords tells the query-builder to eager-load the nodes that are connected to the "passwords"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedPasswords(name string, opts ...func(*UserPasswordQuery)) *UserQuery {
	query := (&UserPasswordClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedPasswords == nil {
		uq.withNamedPasswords = make(map[string]*UserPasswordQuery)
	}
	uq.withNamedPasswords[name] = query
	return uq
}

// WithNamedDevices tells the query-builder to eager-load the nodes that are connected to the "devices"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedDevices(name string, opts ...func(*UserDeviceQuery)) *UserQuery {
	query := (&UserDeviceClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedDevices == nil {
		uq.withNamedDevices = make(map[string]*UserDeviceQuery)
	}
	uq.withNamedDevices[name] = query
	return uq
}

// WithNamedOrgs tells the query-builder to eager-load the nodes that are connected to the "orgs"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedOrgs(name string, opts ...func(*OrgQuery)) *UserQuery {
	query := (&OrgClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedOrgs == nil {
		uq.withNamedOrgs = make(map[string]*OrgQuery)
	}
	uq.withNamedOrgs[name] = query
	return uq
}

// WithNamedPermissions tells the query-builder to eager-load the nodes that are connected to the "permissions"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedPermissions(name string, opts ...func(*PermissionQuery)) *UserQuery {
	query := (&PermissionClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedPermissions == nil {
		uq.withNamedPermissions = make(map[string]*PermissionQuery)
	}
	uq.withNamedPermissions[name] = query
	return uq
}

// WithNamedOrgUser tells the query-builder to eager-load the nodes that are connected to the "org_user"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithNamedOrgUser(name string, opts ...func(*OrgUserQuery)) *UserQuery {
	query := (&OrgUserClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if uq.withNamedOrgUser == nil {
		uq.withNamedOrgUser = make(map[string]*OrgUserQuery)
	}
	uq.withNamedOrgUser[name] = query
	return uq
}

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	selector
	build *UserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, "GroupBy")
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UserGroupBy) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, "Select")
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserSelect](ctx, us.UserQuery, us, us.inters, v)
}

func (us *UserSelect) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
