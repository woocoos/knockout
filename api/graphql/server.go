package graphql

import (
	"context"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/tsingsun/woocoo/contrib/gql"
	"github.com/tsingsun/woocoo/contrib/telemetry/otelweb"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/store/redisx"
	"github.com/tsingsun/woocoo/web"
	"github.com/tsingsun/woocoo/web/handler/authz"
	entadapter "github.com/woocoos/casbin-ent-adapter"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout-go/pkg/middleware"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/fileidentity"
	"github.com/woocoos/knockout/ent/oauthclient"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgapp"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orgroleuser"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/orguserpreference"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/userdevice"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpasswordpolicy"
	"github.com/woocoos/knockout/security"
	"github.com/woocoos/knockout/service/resource"
)

type ServerOptions struct {
	portalDB    *ent.Client
	casbinDB    *casbinent.Client
	kosdk       *api.SDK
	redisClient *redisx.Client
}

type Server struct {
	ServerOptions
	webSrv   *web.Server
	resolver *Resolver
}

func NewServer(cnf *conf.AppConfiguration, opts ...ServerOption) *Server {
	s := &Server{
		ServerOptions: ServerOptions{},
	}

	for _, opt := range opts {
		opt(&s.ServerOptions)
	}

	s.buildRedis(cnf)
	buildCasbin(cnf, s.casbinDB)

	rs := resource.NewService(
		resource.WithClient(s.portalDB),
		resource.WithRedis(s.redisClient),
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
	adapter, err := entadapter.NewAdapterWithClient(client)
	if err != nil {
		panic(err)
	}
	err = casbin.SetAuthorizer(cnf.Sub("authz"), casbin.WithAdapter(adapter))
	if err != nil {
		panic(err)
	}
}

func buildPortalHook(db *ent.Client, ss *resource.Service) {
	hook := security.NewEntHook(db)
	db.Org.Intercept(hook.OrgTraverseFunc(org.FieldID))
	// 需要判断parent_id
	db.Org.Use(hook.OrgMutationInAllowOrg(security.AllOp, org.FieldID))
	db.OrgRole.Intercept(hook.OrgTraverseFunc(orgrole.FieldOrgID))
	db.OrgRole.Use(hook.OrgMutationInAllowOrg(security.AllOp, orgrole.FieldOrgID))
	db.OrgPolicy.Intercept(hook.OrgTraverseFunc(orgpolicy.FieldOrgID))
	db.OrgRole.Use(hook.OrgMutationInAllowOrg(security.AllOp, orgrole.FieldOrgID))
	db.OrgUser.Intercept(hook.OrgTraverseFunc(orguser.FieldOrgID))
	db.OrgUser.Use(hook.OrgMutationInAllowOrg(security.AllOp, orguser.FieldOrgID))
	db.OrgApp.Intercept(hook.OrgTraverseFunc(orgapp.FieldOrgID))
	db.OrgApp.Use(hook.OrgMutationInAllowOrg(security.AllOp, orgapp.FieldOrgID))
	db.App.Intercept(hook.OrgTraverseFunc(app.FieldOwnerOrgID))
	db.App.Use(hook.OrgMutationInAllowOrg(security.AllOp, app.FieldOwnerOrgID))
	db.Permission.Intercept(hook.OrgTraverseFunc(permission.FieldOrgID))
	db.Permission.Use(hook.OrgMutationInAllowOrg(security.AllOp, permission.FieldOrgID))
	db.OrgRoleUser.Intercept(hook.OrgTraverseFunc(orgroleuser.FieldOrgID))
	db.OrgRoleUser.Use(hook.OrgMutationInAllowOrg(security.AllOp, orgroleuser.FieldOrgID))
	db.OrgUserPreference.Intercept(hook.OrgTraverseFunc(orguserpreference.FieldOrgID))
	db.OrgUserPreference.Use(hook.OrgMutationInAllowOrg(security.AllOp, orguserpreference.FieldOrgID))
	db.UserPasswordPolicy.Intercept(hook.OrgTraverseFunc(userpasswordpolicy.FieldTenantID))
	db.UserPasswordPolicy.Use(hook.OrgMutationInAllowOrg(security.AllOp, userpasswordpolicy.FieldTenantID))
	db.User.Use(hook.UserMutationAllow(security.AllOp, user.FieldID))
	db.UserLoginProfile.Use(hook.UserMutationAllow(security.AllOp, userloginprofile.FieldUserID))
	db.UserIdentity.Use(hook.UserMutationAllow(security.AllOp, useridentity.FieldUserID))
	db.UserDevice.Use(hook.UserMutationAllow(security.AllOp, userdevice.FieldUserID))
	db.OauthClient.Use(hook.UserMutationAllow(security.AllOp, oauthclient.FieldUserID))
	db.FileIdentity.Intercept(hook.OrgTraverseFunc(fileidentity.FieldTenantID))
	db.FileIdentity.Use(hook.OrgMutationInAllowOrg(security.AllOp, fileidentity.FieldTenantID))
}

func (s *Server) buildRedis(cnf *conf.AppConfiguration) {
	if cnf.IsSet("store.redis") {
		cli, err := redisx.NewClient(cnf.Sub("store.redis"))
		if err != nil {
			panic(err)
		}
		s.redisClient = cli
	}
}
