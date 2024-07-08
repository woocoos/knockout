// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"github.com/woocoos/entcache"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appdict"
	"github.com/woocoos/knockout/ent/appdictitem"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/appres"
	"github.com/woocoos/knockout/ent/approle"
	"github.com/woocoos/knockout/ent/file"
	"github.com/woocoos/knockout/ent/fileidentity"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/oauthclient"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orguserpreference"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/userdevice"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

var appImplementors = []string{"App", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*App) IsNode() {}

var appactionImplementors = []string{"AppAction", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*AppAction) IsNode() {}

var appdictImplementors = []string{"AppDict", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*AppDict) IsNode() {}

var appdictitemImplementors = []string{"AppDictItem", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*AppDictItem) IsNode() {}

var appmenuImplementors = []string{"AppMenu", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*AppMenu) IsNode() {}

var apppolicyImplementors = []string{"AppPolicy", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*AppPolicy) IsNode() {}

var appresImplementors = []string{"AppRes", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*AppRes) IsNode() {}

var approleImplementors = []string{"AppRole", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*AppRole) IsNode() {}

var fileImplementors = []string{"File", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*File) IsNode() {}

var fileidentityImplementors = []string{"FileIdentity", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*FileIdentity) IsNode() {}

var filesourceImplementors = []string{"FileSource", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*FileSource) IsNode() {}

var oauthclientImplementors = []string{"OauthClient", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*OauthClient) IsNode() {}

var orgImplementors = []string{"Org", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Org) IsNode() {}

var orgpolicyImplementors = []string{"OrgPolicy", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*OrgPolicy) IsNode() {}

var orgroleImplementors = []string{"OrgRole", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*OrgRole) IsNode() {}

var orguserpreferenceImplementors = []string{"OrgUserPreference", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*OrgUserPreference) IsNode() {}

var permissionImplementors = []string{"Permission", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Permission) IsNode() {}

var userImplementors = []string{"User", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*User) IsNode() {}

var userdeviceImplementors = []string{"UserDevice", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*UserDevice) IsNode() {}

var useridentityImplementors = []string{"UserIdentity", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*UserIdentity) IsNode() {}

var userloginprofileImplementors = []string{"UserLoginProfile", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*UserLoginProfile) IsNode() {}

var userpasswordImplementors = []string{"UserPassword", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*UserPassword) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case app.Table:
		query := c.App.Query().
			Where(app.ID(id))
		query, err := query.CollectFields(ctx, appImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "App", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case appaction.Table:
		query := c.AppAction.Query().
			Where(appaction.ID(id))
		query, err := query.CollectFields(ctx, appactionImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "AppAction", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case appdict.Table:
		query := c.AppDict.Query().
			Where(appdict.ID(id))
		query, err := query.CollectFields(ctx, appdictImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "AppDict", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case appdictitem.Table:
		query := c.AppDictItem.Query().
			Where(appdictitem.ID(id))
		query, err := query.CollectFields(ctx, appdictitemImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "AppDictItem", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case appmenu.Table:
		query := c.AppMenu.Query().
			Where(appmenu.ID(id))
		query, err := query.CollectFields(ctx, appmenuImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "AppMenu", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case apppolicy.Table:
		query := c.AppPolicy.Query().
			Where(apppolicy.ID(id))
		query, err := query.CollectFields(ctx, apppolicyImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "AppPolicy", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case appres.Table:
		query := c.AppRes.Query().
			Where(appres.ID(id))
		query, err := query.CollectFields(ctx, appresImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "AppRes", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case approle.Table:
		query := c.AppRole.Query().
			Where(approle.ID(id))
		query, err := query.CollectFields(ctx, approleImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "AppRole", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case file.Table:
		query := c.File.Query().
			Where(file.ID(id))
		query, err := query.CollectFields(ctx, fileImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "File", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case fileidentity.Table:
		query := c.FileIdentity.Query().
			Where(fileidentity.ID(id))
		query, err := query.CollectFields(ctx, fileidentityImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "FileIdentity", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case filesource.Table:
		query := c.FileSource.Query().
			Where(filesource.ID(id))
		query, err := query.CollectFields(ctx, filesourceImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "FileSource", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case oauthclient.Table:
		query := c.OauthClient.Query().
			Where(oauthclient.ID(id))
		query, err := query.CollectFields(ctx, oauthclientImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "OauthClient", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case org.Table:
		query := c.Org.Query().
			Where(org.ID(id))
		query, err := query.CollectFields(ctx, orgImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "Org", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case orgpolicy.Table:
		query := c.OrgPolicy.Query().
			Where(orgpolicy.ID(id))
		query, err := query.CollectFields(ctx, orgpolicyImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "OrgPolicy", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case orgrole.Table:
		query := c.OrgRole.Query().
			Where(orgrole.ID(id))
		query, err := query.CollectFields(ctx, orgroleImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "OrgRole", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case orguserpreference.Table:
		query := c.OrgUserPreference.Query().
			Where(orguserpreference.ID(id))
		query, err := query.CollectFields(ctx, orguserpreferenceImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "OrgUserPreference", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case permission.Table:
		query := c.Permission.Query().
			Where(permission.ID(id))
		query, err := query.CollectFields(ctx, permissionImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "Permission", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		query, err := query.CollectFields(ctx, userImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "User", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case userdevice.Table:
		query := c.UserDevice.Query().
			Where(userdevice.ID(id))
		query, err := query.CollectFields(ctx, userdeviceImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "UserDevice", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case useridentity.Table:
		query := c.UserIdentity.Query().
			Where(useridentity.ID(id))
		query, err := query.CollectFields(ctx, useridentityImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "UserIdentity", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case userloginprofile.Table:
		query := c.UserLoginProfile.Query().
			Where(userloginprofile.ID(id))
		query, err := query.CollectFields(ctx, userloginprofileImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "UserLoginProfile", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	case userpassword.Table:
		query := c.UserPassword.Query().
			Where(userpassword.ID(id))
		query, err := query.CollectFields(ctx, userpasswordImplementors...)
		if err != nil {
			return nil, err
		}
		n, err := query.Only(entcache.WithRefEntryKey(ctx, "UserPassword", id))
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case app.Table:
		query := c.App.Query().
			Where(app.IDIn(ids...))
		query, err := query.CollectFields(ctx, appImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case appaction.Table:
		query := c.AppAction.Query().
			Where(appaction.IDIn(ids...))
		query, err := query.CollectFields(ctx, appactionImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case appdict.Table:
		query := c.AppDict.Query().
			Where(appdict.IDIn(ids...))
		query, err := query.CollectFields(ctx, appdictImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case appdictitem.Table:
		query := c.AppDictItem.Query().
			Where(appdictitem.IDIn(ids...))
		query, err := query.CollectFields(ctx, appdictitemImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case appmenu.Table:
		query := c.AppMenu.Query().
			Where(appmenu.IDIn(ids...))
		query, err := query.CollectFields(ctx, appmenuImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case apppolicy.Table:
		query := c.AppPolicy.Query().
			Where(apppolicy.IDIn(ids...))
		query, err := query.CollectFields(ctx, apppolicyImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case appres.Table:
		query := c.AppRes.Query().
			Where(appres.IDIn(ids...))
		query, err := query.CollectFields(ctx, appresImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case approle.Table:
		query := c.AppRole.Query().
			Where(approle.IDIn(ids...))
		query, err := query.CollectFields(ctx, approleImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case file.Table:
		query := c.File.Query().
			Where(file.IDIn(ids...))
		query, err := query.CollectFields(ctx, fileImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case fileidentity.Table:
		query := c.FileIdentity.Query().
			Where(fileidentity.IDIn(ids...))
		query, err := query.CollectFields(ctx, fileidentityImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case filesource.Table:
		query := c.FileSource.Query().
			Where(filesource.IDIn(ids...))
		query, err := query.CollectFields(ctx, filesourceImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case oauthclient.Table:
		query := c.OauthClient.Query().
			Where(oauthclient.IDIn(ids...))
		query, err := query.CollectFields(ctx, oauthclientImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case org.Table:
		query := c.Org.Query().
			Where(org.IDIn(ids...))
		query, err := query.CollectFields(ctx, orgImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case orgpolicy.Table:
		query := c.OrgPolicy.Query().
			Where(orgpolicy.IDIn(ids...))
		query, err := query.CollectFields(ctx, orgpolicyImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case orgrole.Table:
		query := c.OrgRole.Query().
			Where(orgrole.IDIn(ids...))
		query, err := query.CollectFields(ctx, orgroleImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case orguserpreference.Table:
		query := c.OrgUserPreference.Query().
			Where(orguserpreference.IDIn(ids...))
		query, err := query.CollectFields(ctx, orguserpreferenceImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case permission.Table:
		query := c.Permission.Query().
			Where(permission.IDIn(ids...))
		query, err := query.CollectFields(ctx, permissionImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, userImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case userdevice.Table:
		query := c.UserDevice.Query().
			Where(userdevice.IDIn(ids...))
		query, err := query.CollectFields(ctx, userdeviceImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case useridentity.Table:
		query := c.UserIdentity.Query().
			Where(useridentity.IDIn(ids...))
		query, err := query.CollectFields(ctx, useridentityImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case userloginprofile.Table:
		query := c.UserLoginProfile.Query().
			Where(userloginprofile.IDIn(ids...))
		query, err := query.CollectFields(ctx, userloginprofileImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case userpassword.Table:
		query := c.UserPassword.Query().
			Where(userpassword.IDIn(ids...))
		query, err := query.CollectFields(ctx, userpasswordImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
