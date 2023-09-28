// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *App) Menus(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *AppMenuOrder, where *AppMenuWhereInput,
) (*AppMenuConnection, error) {
	opts := []AppMenuPaginateOption{
		WithAppMenuOrder(orderBy),
		WithAppMenuFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := a.Edges.totalCount[0][alias]
	if nodes, err := a.NamedMenus(alias); err == nil || hasTotalCount {
		pager, err := newAppMenuPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &AppMenuConnection{Edges: []*AppMenuEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return a.QueryMenus().Paginate(ctx, after, first, before, last, opts...)
}

func (a *App) Actions(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *AppActionOrder, where *AppActionWhereInput,
) (*AppActionConnection, error) {
	opts := []AppActionPaginateOption{
		WithAppActionOrder(orderBy),
		WithAppActionFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := a.Edges.totalCount[1][alias]
	if nodes, err := a.NamedActions(alias); err == nil || hasTotalCount {
		pager, err := newAppActionPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &AppActionConnection{Edges: []*AppActionEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return a.QueryActions().Paginate(ctx, after, first, before, last, opts...)
}

func (a *App) Resources(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *AppResOrder, where *AppResWhereInput,
) (*AppResConnection, error) {
	opts := []AppResPaginateOption{
		WithAppResOrder(orderBy),
		WithAppResFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := a.Edges.totalCount[2][alias]
	if nodes, err := a.NamedResources(alias); err == nil || hasTotalCount {
		pager, err := newAppResPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &AppResConnection{Edges: []*AppResEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return a.QueryResources().Paginate(ctx, after, first, before, last, opts...)
}

func (a *App) Roles(ctx context.Context) (result []*AppRole, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.RolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryRoles().All(ctx)
	}
	return result, err
}

func (a *App) Policies(ctx context.Context) (result []*AppPolicy, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = a.NamedPolicies(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = a.Edges.PoliciesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = a.QueryPolicies().All(ctx)
	}
	return result, err
}

func (a *App) Orgs(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *OrgOrder, where *OrgWhereInput,
) (*OrgConnection, error) {
	opts := []OrgPaginateOption{
		WithOrgOrder(orderBy),
		WithOrgFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := a.Edges.totalCount[5][alias]
	if nodes, err := a.NamedOrgs(alias); err == nil || hasTotalCount {
		pager, err := newOrgPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &OrgConnection{Edges: []*OrgEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return a.QueryOrgs().Paginate(ctx, after, first, before, last, opts...)
}

func (a *App) Dicts(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *AppDictOrder, where *AppDictWhereInput,
) (*AppDictConnection, error) {
	opts := []AppDictPaginateOption{
		WithAppDictOrder(orderBy),
		WithAppDictFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := a.Edges.totalCount[6][alias]
	if nodes, err := a.NamedDicts(alias); err == nil || hasTotalCount {
		pager, err := newAppDictPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &AppDictConnection{Edges: []*AppDictEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return a.QueryDicts().Paginate(ctx, after, first, before, last, opts...)
}

func (aa *AppAction) App(ctx context.Context) (*App, error) {
	result, err := aa.Edges.AppOrErr()
	if IsNotLoaded(err) {
		result, err = aa.QueryApp().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (aa *AppAction) Menus(ctx context.Context) (result []*AppMenu, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = aa.NamedMenus(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = aa.Edges.MenusOrErr()
	}
	if IsNotLoaded(err) {
		result, err = aa.QueryMenus().All(ctx)
	}
	return result, err
}

func (ad *AppDict) App(ctx context.Context) (*App, error) {
	result, err := ad.Edges.AppOrErr()
	if IsNotLoaded(err) {
		result, err = ad.QueryApp().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ad *AppDict) Items(ctx context.Context) (result []*AppDictItem, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ad.NamedItems(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ad.Edges.ItemsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ad.QueryItems().All(ctx)
	}
	return result, err
}

func (adi *AppDictItem) Dict(ctx context.Context) (*AppDict, error) {
	result, err := adi.Edges.DictOrErr()
	if IsNotLoaded(err) {
		result, err = adi.QueryDict().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (adi *AppDictItem) Org(ctx context.Context) (*Org, error) {
	result, err := adi.Edges.OrgOrErr()
	if IsNotLoaded(err) {
		result, err = adi.QueryOrg().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (am *AppMenu) App(ctx context.Context) (*App, error) {
	result, err := am.Edges.AppOrErr()
	if IsNotLoaded(err) {
		result, err = am.QueryApp().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (am *AppMenu) Action(ctx context.Context) (*AppAction, error) {
	result, err := am.Edges.ActionOrErr()
	if IsNotLoaded(err) {
		result, err = am.QueryAction().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ap *AppPolicy) App(ctx context.Context) (*App, error) {
	result, err := ap.Edges.AppOrErr()
	if IsNotLoaded(err) {
		result, err = ap.QueryApp().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ap *AppPolicy) Roles(ctx context.Context) (result []*AppRole, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ap.NamedRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ap.Edges.RolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ap.QueryRoles().All(ctx)
	}
	return result, err
}

func (ar *AppRes) App(ctx context.Context) (*App, error) {
	result, err := ar.Edges.AppOrErr()
	if IsNotLoaded(err) {
		result, err = ar.QueryApp().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ar *AppRole) App(ctx context.Context) (*App, error) {
	result, err := ar.Edges.AppOrErr()
	if IsNotLoaded(err) {
		result, err = ar.QueryApp().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ar *AppRole) Policies(ctx context.Context) (result []*AppPolicy, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ar.NamedPolicies(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ar.Edges.PoliciesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ar.QueryPolicies().All(ctx)
	}
	return result, err
}

func (f *File) Source(ctx context.Context) (*FileSource, error) {
	result, err := f.Edges.SourceOrErr()
	if IsNotLoaded(err) {
		result, err = f.QuerySource().Only(ctx)
	}
	return result, err
}

func (fs *FileSource) Files(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *FileOrder, where *FileWhereInput,
) (*FileConnection, error) {
	opts := []FilePaginateOption{
		WithFileOrder(orderBy),
		WithFileFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := fs.Edges.totalCount[0][alias]
	if nodes, err := fs.NamedFiles(alias); err == nil || hasTotalCount {
		pager, err := newFilePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &FileConnection{Edges: []*FileEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return fs.QueryFiles().Paginate(ctx, after, first, before, last, opts...)
}

func (oc *OauthClient) User(ctx context.Context) (*User, error) {
	result, err := oc.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = oc.QueryUser().Only(ctx)
	}
	return result, err
}

func (o *Org) Parent(ctx context.Context) (*Org, error) {
	result, err := o.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = o.QueryParent().Only(ctx)
	}
	return result, err
}

func (o *Org) Children(ctx context.Context) (result []*Org, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = o.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = o.Edges.ChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = o.QueryChildren().All(ctx)
	}
	return result, err
}

func (o *Org) Owner(ctx context.Context) (*User, error) {
	result, err := o.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = o.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (o *Org) Users(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *UserOrder, where *UserWhereInput,
) (*UserConnection, error) {
	opts := []UserPaginateOption{
		WithUserOrder(orderBy),
		WithUserFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := o.Edges.totalCount[3][alias]
	if nodes, err := o.NamedUsers(alias); err == nil || hasTotalCount {
		pager, err := newUserPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &UserConnection{Edges: []*UserEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return o.QueryUsers().Paginate(ctx, after, first, before, last, opts...)
}

func (o *Org) Permissions(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *PermissionOrder, where *PermissionWhereInput,
) (*PermissionConnection, error) {
	opts := []PermissionPaginateOption{
		WithPermissionOrder(orderBy),
		WithPermissionFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := o.Edges.totalCount[4][alias]
	if nodes, err := o.NamedPermissions(alias); err == nil || hasTotalCount {
		pager, err := newPermissionPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &PermissionConnection{Edges: []*PermissionEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return o.QueryPermissions().Paginate(ctx, after, first, before, last, opts...)
}

func (o *Org) Policies(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *OrgPolicyOrder, where *OrgPolicyWhereInput,
) (*OrgPolicyConnection, error) {
	opts := []OrgPolicyPaginateOption{
		WithOrgPolicyOrder(orderBy),
		WithOrgPolicyFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := o.Edges.totalCount[5][alias]
	if nodes, err := o.NamedPolicies(alias); err == nil || hasTotalCount {
		pager, err := newOrgPolicyPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &OrgPolicyConnection{Edges: []*OrgPolicyEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return o.QueryPolicies().Paginate(ctx, after, first, before, last, opts...)
}

func (o *Org) Apps(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *AppOrder, where *AppWhereInput,
) (*AppConnection, error) {
	opts := []AppPaginateOption{
		WithAppOrder(orderBy),
		WithAppFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := o.Edges.totalCount[6][alias]
	if nodes, err := o.NamedApps(alias); err == nil || hasTotalCount {
		pager, err := newAppPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &AppConnection{Edges: []*AppEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return o.QueryApps().Paginate(ctx, after, first, before, last, opts...)
}

func (op *OrgPolicy) Org(ctx context.Context) (*Org, error) {
	result, err := op.Edges.OrgOrErr()
	if IsNotLoaded(err) {
		result, err = op.QueryOrg().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (op *OrgPolicy) Permissions(ctx context.Context) (result []*Permission, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = op.NamedPermissions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = op.Edges.PermissionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = op.QueryPermissions().All(ctx)
	}
	return result, err
}

func (oup *OrgUserPreference) User(ctx context.Context) (*User, error) {
	result, err := oup.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = oup.QueryUser().Only(ctx)
	}
	return result, err
}

func (oup *OrgUserPreference) Org(ctx context.Context) (*Org, error) {
	result, err := oup.Edges.OrgOrErr()
	if IsNotLoaded(err) {
		result, err = oup.QueryOrg().Only(ctx)
	}
	return result, err
}

func (pe *Permission) Org(ctx context.Context) (*Org, error) {
	result, err := pe.Edges.OrgOrErr()
	if IsNotLoaded(err) {
		result, err = pe.QueryOrg().Only(ctx)
	}
	return result, err
}

func (pe *Permission) User(ctx context.Context) (*User, error) {
	result, err := pe.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = pe.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pe *Permission) Role(ctx context.Context) (*OrgRole, error) {
	result, err := pe.Edges.RoleOrErr()
	if IsNotLoaded(err) {
		result, err = pe.QueryRole().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pe *Permission) OrgPolicy(ctx context.Context) (*OrgPolicy, error) {
	result, err := pe.Edges.OrgPolicyOrErr()
	if IsNotLoaded(err) {
		result, err = pe.QueryOrgPolicy().Only(ctx)
	}
	return result, err
}

func (u *User) Identities(ctx context.Context) (result []*UserIdentity, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedIdentities(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.IdentitiesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryIdentities().All(ctx)
	}
	return result, err
}

func (u *User) LoginProfile(ctx context.Context) (*UserLoginProfile, error) {
	result, err := u.Edges.LoginProfileOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryLoginProfile().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Devices(ctx context.Context) (result []*UserDevice, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedDevices(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.DevicesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryDevices().All(ctx)
	}
	return result, err
}

func (u *User) Permissions(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *PermissionOrder, where *PermissionWhereInput,
) (*PermissionConnection, error) {
	opts := []PermissionPaginateOption{
		WithPermissionOrder(orderBy),
		WithPermissionFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[3][alias]
	if nodes, err := u.NamedPermissions(alias); err == nil || hasTotalCount {
		pager, err := newPermissionPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &PermissionConnection{Edges: []*PermissionEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryPermissions().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) OauthClients(ctx context.Context) (result []*OauthClient, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedOauthClients(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.OauthClientsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryOauthClients().All(ctx)
	}
	return result, err
}

func (ud *UserDevice) User(ctx context.Context) (*User, error) {
	result, err := ud.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = ud.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ui *UserIdentity) User(ctx context.Context) (*User, error) {
	result, err := ui.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = ui.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ulp *UserLoginProfile) User(ctx context.Context) (*User, error) {
	result, err := ulp.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = ulp.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (up *UserPassword) User(ctx context.Context) (*User, error) {
	result, err := up.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = up.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}
