package graphql

import (
	"context"
	"github.com/stretchr/testify/suite"
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
	resolver *Resolver
	mr       *mutationResolver
	qr       *queryResolver
}

func (s *graphqlSuite) SetupSuite() {
	err := s.BaseSuite.Setup()
	s.Require().NoError(err)
	data.InitBase(s.DriverName, s.DSN)

	s.resolver = NewResolver(WithClient(s.Client), WithResource(&resource.Service{
		Client:     s.Client,
		HttpClient: nil,
		OASOptions: resource.OASOptions{},
	}))

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
