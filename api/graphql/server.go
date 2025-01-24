package graphql

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler/authz"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout-go/pkg/middleware"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/fileidentity"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgapp"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orgroleuser"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/orguserpreference"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/userpasswordpolicy"
	"github.com/woocoos/knockout/service/resource"
)

type ServerOptions struct {
	portalDB *ent.Client
	casbinDB *casbinent.Client
}

type Server struct {
	ServerOptions
	webSrv   *web.Server
	kosdk    *api.SDK
	resolver *Resolver
}

func NewServer(cnf *conf.AppConfiguration, opts ...ServerOption) *Server {
	s := &Server{
		ServerOptions: ServerOptions{},
	}

	for _, opt := range opts {
		opt(&s.ServerOptions)
	}

	buildCasbin(cnf, s.casbinDB)

	var err error
	s.kosdk, err = api.NewSDK(cnf.Sub("kosdk"))
	if err != nil {
		panic(err)
	}

	rs := resource.NewService(
		resource.WithClient(s.portalDB),
		resource.WithKOSDK(s.kosdk),
		resource.WithCfg(cnf))
	buildPortalHook(s.portalDB, rs)
	s.resolver = NewResolver(WithClient(s.portalDB),
		WithResource(rs))
	s.buildWebEngine(cnf)

	return s
}

func (s *Server) Start(ctx context.Context) error {
	return s.webSrv.Start(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	s.portalDB.Close()
	s.casbinDB.Close()
	return nil
}

func (s *Server) buildWebEngine(cnf *conf.AppConfiguration) {
	s.webSrv = web.New(web.WithConfiguration(cnf.Sub("web")),
		web.WithGracefulStop(),
		gql.RegisterMiddleware(),
		otelweb.RegisterMiddleware(),
		web.WithMiddlewareNewFunc("authz", authz.Middleware),
		middleware.RegisterTenantID(),
		middleware.RegisterTokenSigner(),
	)

	gqlSrv := handler.NewDefaultServer(NewSchema(s.resolver))
	gqlSrv.AroundResponses(middleware.SimplePagination())
	// mutation transaction
	gqlSrv.Use(entgql.Transactioner{TxOpener: s.portalDB})

	if err := gql.RegisterGraphqlServer(s.webSrv, gqlSrv); err != nil {
		panic(err)
	}
}

func buildCasbin(cnf *conf.AppConfiguration, client *casbinent.Client) {
	err := casbin.SetAuthorizer(cnf.Sub("authz"), client)
	if err != nil {
		panic(err)
	}
}

func buildPortalHook(db *ent.Client, ss *resource.Service) {
	hook := resource.NewEntHook(ss)
	db.Org.Intercept(hook.OrgTraverseFunc(org.FieldID))
	// 需要判断parent_id
	db.Org.Use(hook.OrgMutationInAllowOrg())
	db.OrgRole.Intercept(hook.OrgTraverseFunc(orgrole.FieldOrgID))
	db.OrgRole.Use(hook.MutationInAllowOrg(func(ctx context.Context, id int, client *ent.Client) (int, error) {
		return client.OrgRole.Query().Where(orgrole.ID(id)).Select(orgrole.FieldOrgID).Int(ctx)
	}))
	db.OrgPolicy.Intercept(hook.OrgTraverseFunc(orgpolicy.FieldOrgID))
	db.OrgPolicy.Use(hook.MutationInAllowOrg(func(ctx context.Context, id int, client *ent.Client) (int, error) {
		return client.OrgPolicy.Query().Where(orgpolicy.ID(id)).Select(orgpolicy.FieldOrgID).Int(ctx)
	}))
	db.OrgUser.Intercept(hook.OrgTraverseFunc(orguser.FieldOrgID))
	db.OrgUser.Use(hook.MutationInAllowOrg(func(ctx context.Context, id int, client *ent.Client) (int, error) {
		return client.OrgUser.Query().Where(orguser.ID(id)).Select(orguser.FieldOrgID).Int(ctx)
	}))
	db.OrgApp.Intercept(hook.OrgTraverseFunc(orgapp.FieldOrgID))
	db.OrgApp.Use(hook.MutationInAllowOrg(func(ctx context.Context, id int, client *ent.Client) (int, error) {
		return client.OrgApp.Query().Where(orgapp.ID(id)).Select(orgapp.FieldOrgID).Int(ctx)
	}))
	db.App.Intercept(hook.OrgTraverseFunc(app.FieldOwnerOrgID))
	// 需要取owner_org_id处理
	db.App.Use(hook.AppMutationInAllowOrg())
	db.Permission.Intercept(hook.OrgTraverseFunc(permission.FieldOrgID))
	db.Permission.Use(hook.MutationInAllowOrg(func(ctx context.Context, id int, client *ent.Client) (int, error) {
		return client.Permission.Query().Where(permission.ID(id)).Select(permission.FieldOrgID).Int(ctx)
	}))
	db.OrgRoleUser.Intercept(hook.OrgTraverseFunc(orgroleuser.FieldOrgID))
	db.OrgRoleUser.Use(hook.MutationInAllowOrg(func(ctx context.Context, id int, client *ent.Client) (int, error) {
		return client.OrgRoleUser.Query().Where(orgroleuser.ID(id)).Select(orgroleuser.FieldOrgID).Int(ctx)
	}))
	db.OrgUserPreference.Intercept(hook.OrgTraverseFunc(orguserpreference.FieldOrgID))
	db.OrgUserPreference.Use(hook.MutationInAllowOrg(func(ctx context.Context, id int, client *ent.Client) (int, error) {
		return client.OrgUserPreference.Query().Where(orguserpreference.ID(id)).Select(orguserpreference.FieldOrgID).Int(ctx)
	}))
	db.UserPasswordPolicy.Intercept(hook.OrgTraverseFunc(userpasswordpolicy.FieldTenantID))
	db.UserPasswordPolicy.Use(hook.UPPMutationInAllowOrg())
	db.User.Use(hook.UserMutationAllowAll())
	db.UserLoginProfile.Use(hook.UserLoginProfileMutationAllowAll())
	db.UserIdentity.Use(hook.UserIdentityMutationAllowAll())
	db.UserDevice.Use(hook.UserDeviceMutationAllowAll())
	db.OauthClient.Use(hook.OauthClientMutationAllowAll())
	db.FileIdentity.Intercept(hook.OrgTraverseFunc(fileidentity.FieldTenantID))
	db.FileIdentity.Use(hook.FileIdentityMutationInAllowOrg())

}
