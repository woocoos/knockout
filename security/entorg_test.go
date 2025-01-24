package security

import (
	"github.com/stretchr/testify/suite"
	"github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgapp"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/test/testsuite"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/knockout/ent/runtime"
)

type testSuite struct {
	testsuite.BaseSuite
	entHook *EntHook
}

func TestSuite(t *testing.T) {
	suite.Run(t, &testSuite{
		BaseSuite: testsuite.BaseSuite{
			DriverName: "sqlite3",
			DSN:        "file:security?mode=memory&cache=shared&_fk=1",
		},
	})
}

func (t *testSuite) SetupSuite() {
	t.Require().NoError(t.BaseSuite.Setup())
	err := casbin.SetAuthorizer(t.App.AppConfiguration().Sub("authz"), t.AuthDbClient)
	if err != nil {
		panic(err)
	}
	t.entHook = NewEntHook(t.Client)
}

func (t *testSuite) TestOrgHook() {
	ctx := testsuite.NewTestCtx(1, 1, t.Client)
	client := t.Client.Debug()
	client.Org.CreateBulk(
		client.Org.Create().SetID(1).SetName("org1").SetPath("org1").SetParentID(0).SetCreatedBy(1),
		client.Org.Create().SetID(11).SetName("org11").SetPath("org1/org11").SetParentID(1).SetCreatedBy(1),
		client.Org.Create().SetID(2).SetName("org2").SetPath("org2").SetParentID(0).SetCreatedBy(1),
		client.Org.Create().SetID(22).SetName("org22").SetPath("org2/org22").SetParentID(2).SetCreatedBy(1),
	).ExecX(ctx)
	client.OrgRole.CreateBulk(
		client.OrgRole.Create().SetID(1).SetName("administrators").SetOrgID(1).SetKind(orgrole.KindGroup).SetCreatedBy(1),
	).ExecX(ctx)

	t.Run("org", func() {
		org2 := client.Org.GetX(ctx, 2)

		client.Org.Use(t.entHook.OrgMutationInAllowOrg(
			AllOp,
			org.FieldID))
		client.Org.Intercept(t.entHook.OrgTraverseFunc(org.FieldID))

		orgs, err := client.Org.Query().Where(org.IDIn(1, 11, 2, 22)).All(ctx)
		t.Require().NoError(err)
		t.Equal(2, len(orgs))
		t.Equal("org1", orgs[0].Name)
		t.Equal("org11", orgs[1].Name)

		_, err = org2.Update().SetName("will update fail").Save(ctx)
		t.Require().Error(err)
		cc, err := client.Org.Update().SetUpdatedBy(2).Where(org.IDIn(1, 11, 2, 22)).Save(ctx)
		t.Require().NoError(err)
		t.Equal(2, cc)
		ctx2 := testsuite.NewTestCtx(1, 2, t.Client)
		cc, err = client.Org.Delete().Where(org.IDIn(1, 11, 2, 22)).Exec(ctx2)
		t.Require().NoError(err)
		t.Equal(2, cc, "only delete org 2")
		client.Org.GetX(ctx, 1)
		err = client.Org.Create().SetName("new").SetParentID(11).SetKind(org.KindOrganization).SetCreatedBy(1).
			Exec(ctx)
		t.Require().NoError(err)
		err = client.Org.Create().SetName("new").SetParentID(22).SetKind(org.KindOrganization).SetCreatedBy(1).
			Exec(ctx)
		t.Require().Error(err)
	})
	t.Run("ref", func() {
		client.OrgApp.CreateBulk(
			client.OrgApp.Create().SetID(1).SetOrgID(1).SetAppID(1).SetCreatedBy(1),
			client.OrgApp.Create().SetID(2).SetOrgID(2).SetAppID(2).SetCreatedBy(1),
		).ExecX(ctx)
		client.OrgApp.Use(t.entHook.OrgMutationInAllowOrg(
			AllOp,
			orgapp.FieldOrgID,
		))

		client.OrgApp.Intercept(t.entHook.OrgTraverseFunc(orgapp.FieldOrgID))
		orgapps, err := client.OrgApp.Query().Where(orgapp.OrgIDIn(1, 11, 2, 22)).All(ctx)
		t.Require().NoError(err)
		t.Equal(1, len(orgapps))
		t.Equal(1, orgapps[0].OrgID)
		cc, err := client.OrgApp.Update().SetUpdatedBy(2).Where(orgapp.AppIDIn(1, 2)).Save(ctx)
		t.Require().NoError(err)
		t.Equal(1, cc)
		cc, err = client.OrgApp.Delete().Where(orgapp.AppIDIn(1, 2)).Exec(ctx)
		t.Require().NoError(err)
		t.Equal(1, cc)
		err = client.OrgApp.Create().SetOrgID(1).SetAppID(3).SetCreatedBy(1).Exec(ctx)
		t.Require().NoError(err)
		err = client.OrgApp.Create().SetOrgID(2).SetAppID(3).SetCreatedBy(1).Exec(ctx)
		t.Require().Error(err)
	})
}
