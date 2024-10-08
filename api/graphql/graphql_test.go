package graphql

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo/pkg/gds"
	"github.com/tsingsun/woocoo/pkg/security"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/appres"
	"github.com/woocoos/knockout/ent/approle"
	"github.com/woocoos/knockout/ent/orgapp"
	"github.com/woocoos/knockout/script/data"
	"github.com/woocoos/knockout/service/resource"
	"github.com/woocoos/knockout/test/testsuite"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/knockout/ent/runtime"
)

// graphqlSuite for graphql
//
// Name with _ will be run after Test* methods. Doing delete operation can be run in those.
type graphqlSuite struct {
	testsuite.BaseSuite
	resolver     *Resolver
	mr           *mutationResolver
	qr           *queryResolver
	casbinClient *casbinent.Client
}

func (s *graphqlSuite) SetupSuite() {
	err := s.BaseSuite.Setup()
	s.Require().NoError(err)
	data.InitBase(s.DriverName, s.DSN)

	s.resolver = NewResolver(WithClient(s.Client), WithResource(&resource.Service{
		Client: s.Client,
	}))

	// 初始化casbin
	drv, err := sql.Open(s.DriverName, s.DSN)
	s.casbinClient = casbinent.NewClient(casbinent.Driver(drv), casbinent.Debug())
	buildCashbin(s.Cnf, s.casbinClient)

	s.mr = &mutationResolver{
		Resolver: s.resolver,
	}
	s.qr = &queryResolver{
		Resolver: s.resolver,
	}
}

func TestGraphqlSuite(t *testing.T) {
	s := &graphqlSuite{
		BaseSuite: testsuite.BaseSuite{
			DSN:        "file:msgcenter?mode=memory&cache=shared&_fk=1",
			DriverName: "sqlite3",
		},
	}
	suite.Run(t, s)
}

func (s *graphqlSuite) TestApp() {
	s.Run("update", func() {
		ap, err := s.Client.App.UpdateOneID(1).SetUpdatedBy(1).SetInput(ent.UpdateAppInput{Scopes: gds.Ptr("a")}).
			Save(context.Background())
		s.Require().NoError(err)
		s.NotNil(ap.Logo)
	})
}

func (s *graphqlSuite) Test_DeleteApp() {
	ctx := context.Background()
	_, err := s.Client.App.Get(ctx, 1)
	s.Require().NoError(err)

	s.Client.OrgApp.Delete().Where(orgapp.AppID(1)).ExecX(ctx)

	_, err = s.mr.DeleteApp(s.NewTestCtx(1, 1), 1)
	s.Require().NoError(err)

	ok := s.Client.AppAction.Query().Where(appaction.AppID(1)).ExistX(ctx)
	s.Require().False(ok)
	ok = s.Client.AppMenu.Query().Where(appmenu.AppID(1)).ExistX(ctx)
	s.Require().False(ok)
	ok = s.Client.AppRes.Query().Where(appres.AppID(1)).ExistX(ctx)
	s.Require().False(ok)
	ok = s.Client.AppRole.Query().Where(approle.AppID(1)).ExistX(ctx)
	s.Require().False(ok)
	ok = s.Client.AppPolicy.Query().Where(apppolicy.AppID(1)).ExistX(ctx)
	s.Require().False(ok)
}

func (s *graphqlSuite) Test_UserPermissions() {
	ctx := security.WithContext(context.Background(), security.NewGenericPrincipalByClaims(jwt.MapClaims{"sub": "1"}))
	ctx = identity.WithTenantID(ctx, 1)

	s.Client.AppAction.Create().SetAppID(1).SetCreatedBy(1).
		SetName("test").SetKind(appaction.KindFunction).SetComments("测试").SetMethod(appaction.MethodRead)

	as, err := s.qr.UserPermissions(ctx, &ent.AppActionWhereInput{
		ID: gds.Ptr(1),
	})
	s.Require().NoError(err)
	acs, err := s.Client.AppAction.Query().Where(appaction.AppID(1)).Count(ctx)
	s.Require().NoError(err)
	s.Equal(len(as), acs)
}
