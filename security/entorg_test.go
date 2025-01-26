package security

import (
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgapp"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orguserpreference"
	"github.com/woocoos/knockout/test/testsuite"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/knockout/ent/runtime"
)

type testSuite struct {
	testsuite.BaseSuite
	authorizer *casbin.Authorizer
	entHook    *EntHook
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
	t.authorizer = security.DefaultAuthorizer.(*casbin.Authorizer)
	t.entHook = NewEntHook(t.Client)
	t.initData()
	_, err = t.authorizer.Enforcer.AddRoleForUserInDomain("1", "r_1", "1")
	t.Require().NoError(err)
}

func (t *testSuite) initData() {
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
}

func (t *testSuite) TestOrgHook() {
	ctx := testsuite.NewTestCtx(1, 1, t.Client)
	client := t.Client.Debug()

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
		err = client.Org.UpdateOneID(11).SetUpdatedBy(2).Exec(ctx)
		t.NoError(err)
		err = client.Org.UpdateOneID(22).SetUpdatedBy(2).Exec(ctx)
		t.ErrorIs(err, ErrTenantIDNotAllow, "cannot update org 22")
		err = client.Org.DeleteOneID(22).Exec(ctx)
		t.True(ent.IsNotFound(err), "cannot delete org 22")
		ctx2 := testsuite.NewTestCtx(1, 2, t.Client)
		cc, err = client.Org.Delete().Where(org.IDIn(1, 11, 2, 22)).Exec(ctx2)
		t.Require().NoError(err)
		t.Equal(2, cc, "only delete org 2")
		err = client.Org.DeleteOneID(11).Exec(ctx2)
		t.True(ent.IsNotFound(err), "cannot delete org 11")

		client.Org.GetX(ctx, 1)
		err = client.Org.Create().SetName("new").SetParentID(11).SetKind(org.KindOrganization).SetCreatedBy(1).
			Exec(ctx)
		t.Require().NoError(err)
		err = client.Org.Create().SetName("new").SetParentID(22).SetKind(org.KindOrganization).SetCreatedBy(1).
			Exec(ctx)
		t.Require().ErrorIs(err, ErrTenantIDNotAllow)
	})
	t.Run("ref", func() {
		client.OrgApp.CreateBulk(
			client.OrgApp.Create().SetID(1).SetOrgID(1).SetAppID(1).SetCreatedBy(1),
			client.OrgApp.Create().SetID(2).SetOrgID(2).SetAppID(2).SetCreatedBy(1),
			client.OrgApp.Create().SetID(3).SetOrgID(3).SetAppID(3).SetCreatedBy(1),
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
		err = client.OrgApp.UpdateOneID(1).SetUpdatedBy(2).Exec(ctx)
		t.Require().NoError(err)
		err = client.OrgApp.UpdateOneID(2).SetUpdatedBy(2).Exec(ctx)
		t.Require().ErrorIs(err, ErrTenantIDNotAllow)
		cc, err = client.OrgApp.Delete().Where(orgapp.AppIDIn(1, 2)).Exec(ctx)
		t.Require().NoError(err)
		t.Equal(1, cc)
		err = client.OrgApp.DeleteOneID(2).Exec(ctx)
		t.Require().True(ent.IsNotFound(err), "cannot delete orgapp 2")
		err = client.OrgApp.Create().SetOrgID(1).SetAppID(3).SetCreatedBy(1).Exec(ctx)
		t.Require().NoError(err)
		err = client.OrgApp.Create().SetOrgID(2).SetAppID(3).SetCreatedBy(1).Exec(ctx)
		t.Require().Error(err)
		ctx2 := testsuite.NewTestCtx(2, 22, t.Client)
		_, err = client.OrgApp.UpdateOneID(3).SetUpdatedBy(2).Save(ctx2)
		t.Require().Error(err)
	})
}

func (t *testSuite) TestUser() {
	ctx := testsuite.NewTestCtx(1, 1, t.Client)
	client := t.Client.Debug()
	client.OrgUserPreference.CreateBulk(
		client.OrgUserPreference.Create().SetID(1).SetOrgID(1).SetUserID(1).SetMenuFavorite([]int{1, 2}).SetCreatedBy(1),
		client.OrgUserPreference.Create().SetID(2).SetOrgID(1).SetUserID(2).SetMenuFavorite([]int{1, 2}).SetCreatedBy(1),
	).ExecX(ctx)

	client.OrgUserPreference.Use(t.entHook.UserMutationAllow(
		AllOp,
		orguserpreference.FieldUserID))

	_, err := t.authorizer.Enforcer.AddPolicy("r_1", "1", userOtherRes, "read", "allow")
	t.Require().NoError(err)
	err = client.OrgUserPreference.Create().SetID(3).SetOrgID(1).SetUserID(3).SetMenuFavorite([]int{1, 2}).
		SetCreatedBy(1).Exec(ctx)
	t.Require().NoError(err)
	_, err = client.OrgUserPreference.UpdateOneID(3).SetUpdatedBy(2).Save(ctx)
	t.Require().NoError(err)
	ctx2 := testsuite.NewTestCtx(2, 1, t.Client)
	_, err = client.OrgUserPreference.UpdateOneID(3).SetUpdatedBy(2).Save(ctx2)
	t.Require().ErrorIs(err, ErrMutationOtherUserNotAllow)
	err = client.OrgUserPreference.DeleteOneID(3).Exec(ctx2)
	t.Require().True(ent.IsNotFound(err))

	cc, err := client.OrgUserPreference.Delete().Where(orguserpreference.IDIn(1, 2, 3)).Exec(ctx2)
	t.Require().NoError(err)
	t.Equal(1, cc, "删除之前创建的")

	cc, err = client.OrgUserPreference.Delete().Where(orguserpreference.IDIn(1, 2, 3)).Exec(ctx)
	t.Require().NoError(err)
	t.Equal(2, cc, "删除现存的所有的")
}
