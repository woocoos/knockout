package graphql

import (
	"context"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/apppolicyview"
	"github.com/woocoos/knockout/ent/orgrole"
	"strconv"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo/pkg/gds"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/appres"
	"github.com/woocoos/knockout/ent/approle"
	"github.com/woocoos/knockout/ent/orgapp"
	_ "github.com/woocoos/knockout/ent/runtime"
	"github.com/woocoos/knockout/script/data"
	"github.com/woocoos/knockout/test/testsuite"
)

// graphqlSuite for graphql
//
// Name with _ will be run after Test* methods. Doing delete operation can be run in those.
type graphqlSuite struct {
	testsuite.BaseSuite
	mr        *mutationResolver
	qr        *queryResolver
	server    *Server
	gqlClient *client.Client
}

func (t *graphqlSuite) SetupSuite() {
	err := t.BaseSuite.Setup()
	t.Require().NoError(err)
	data.InitBase(t.DriverName, t.DSN)

	buildCashbin(t.Cnf, t.AuthDbClient)

	t.server = &Server{
		casbinClient: t.AuthDbClient,
		portalClient: t.Client,
	}
	t.server.buildWebEngine(t.Cnf)
	t.mr = &mutationResolver{
		Resolver: t.server.resolver,
	}
	t.qr = &queryResolver{
		Resolver: t.server.resolver,
	}
	t.gqlClient = client.New(t.server.webSrv.Router(), func(bd *client.Request) {
		bd.HTTP.URL.Path = "/graphql/query"
		bd.HTTP.Header.Set("Authorization", "Bearer "+t.BearToken())
		bd.HTTP.Header.Set("X-Tenant-ID", "1")
	})
}

func TestGraphqlSuite(t *testing.T) {
	s := &graphqlSuite{
		BaseSuite: testsuite.BaseSuite{
			DSN:        "file:graphql?mode=memory&cache=shared&_fk=1",
			DriverName: "sqlite3",
		},
	}
	suite.Run(t, s)
}

func (t *graphqlSuite) TestApp() {
	t.Run("update", func() {
		ap, err := t.Client.App.UpdateOneID(1).SetUpdatedBy(1).SetInput(ent.UpdateAppInput{Scopes: gds.Ptr("a")}).
			Save(context.Background())
		t.Require().NoError(err)
		t.NotNil(ap.Logo)
	})
}

func (t *graphqlSuite) Test_DeleteApp() {
	ctx := context.Background()
	_, err := t.Client.App.Get(ctx, 1)
	t.Require().NoError(err)

	t.Client.OrgApp.Delete().Where(orgapp.AppID(1)).ExecX(ctx)

	_, err = t.mr.DeleteApp(t.NewTestCtx(1, 1), 1)
	t.Require().NoError(err)

	ok := t.Client.AppAction.Query().Where(appaction.AppID(1)).ExistX(ctx)
	t.Require().False(ok)
	ok = t.Client.AppMenu.Query().Where(appmenu.AppID(1)).ExistX(ctx)
	t.Require().False(ok)
	ok = t.Client.AppRes.Query().Where(appres.AppID(1)).ExistX(ctx)
	t.Require().False(ok)
	ok = t.Client.AppRole.Query().Where(approle.AppID(1)).ExistX(ctx)
	t.Require().False(ok)
	ok = t.Client.AppPolicy.Query().Where(apppolicy.AppID(1)).ExistX(ctx)
	t.Require().False(ok)
}

func (t *graphqlSuite) Test_UserPermissions() {
	ctx := security.WithContext(context.Background(), security.NewGenericPrincipalByClaims(jwt.MapClaims{"sub": "1"}))
	ctx = identity.WithTenantID(ctx, 1)

	t.Client.AppAction.Create().SetAppID(1).SetCreatedBy(1).
		SetName("test").SetKind(appaction.KindFunction).SetComments("测试").SetMethod(appaction.MethodRead)

	as, err := t.qr.UserPermissions(ctx, &ent.AppActionWhereInput{
		ID: gds.Ptr(1),
	})
	t.Require().NoError(err)
	acs, err := t.Client.AppAction.Query().Where(appaction.AppID(1)).Count(ctx)
	t.Require().NoError(err)
	t.Equal(len(as), acs)
}

func (t *graphqlSuite) TestQuota() {
	ctx := context.Background()
	quotaItems := []*ent.QuotaItem{
		t.Client.QuotaItem.Create().SetID(1).SetCode("users").SetName("用户数量限制").SetResourceType("number").SetUnit("个").SetActive(true).
			SetCreatedBy(1).SaveX(ctx),

		t.Client.QuotaItem.Create().SetID(2).SetCode("storage").SetName("存储空间").SetResourceType("storage").SetUnit("GB").SetActive(true).
			SetCreatedBy(1).SaveX(ctx),
		t.Client.QuotaItem.Create().SetID(3).SetCode("api").SetName("API调用次数").SetResourceType("number").SetUnit("次").SetActive(true).
			SetCreatedBy(1).SaveX(ctx),
	}

	t.Client.Quota.Create().SetID(1).SetTenantID(1).SetUserID(0).SetQuotaItem(quotaItems[0]).SetLimit(100).SetStartAt(time.Now()).SetCreatedBy(1).SaveX(ctx)
	t.Client.Quota.Create().SetID(2).SetTenantID(1).SetUserID(0).SetQuotaItem(quotaItems[1]).SetLimit(1000).SetStartAt(time.Now().AddDate(0, 0, -1)).
		SetEndAt(time.Now().AddDate(0, 1, 0)).SetCreatedBy(1).SaveX(ctx)
	t.Client.Quota.Create().SetID(3).SetTenantID(1).SetUserID(0).SetQuotaItem(quotaItems[2]).SetLimit(1).SetStartAt(time.Now()).SetCreatedBy(1).SaveX(ctx)
	t.Run("query quota items", func() {
		const query = `
            query {
                quotaItems(first: 10) {
                    edges {
                        node {
                            id
                            code
                            name
                            resourceType
                            unit
                            active
                        }
                    }
                }
            }
        `

		var resp struct {
			QuotaItems struct {
				Edges []struct {
					Node struct {
						ID           string
						Code         string
						Name         string
						ResourceType string
						Unit         string
						Active       bool
					}
				}
			}
		}
		err := t.gqlClient.Post(query, &resp)
		t.Require().NoError(err)
		t.Require().Len(resp.QuotaItems.Edges, 3)
		t.Equal("users", resp.QuotaItems.Edges[0].Node.Code)
	})
	t.Run("query quotas", func() {
		const query = `
            query {
                quotas(first: 10) {
                    edges {
                        node {
                            id
                            limit
                            used
                            startAt
                            endAt
                            quotaItem {
                                code
                                name
                            }
                        }
                    }
                }
            }
        `

		var resp struct {
			Quotas struct {
				Edges []struct {
					Node struct {
						ID        string
						Limit     int64
						Used      int64
						StartAt   string
						EndAt     string
						QuotaItem struct {
							Code string
							Name string
						}
					}
				}
			}
		}

		err := t.gqlClient.Post(query, &resp)
		t.Require().NoError(err)
		t.Require().Len(resp.Quotas.Edges, 3)
	})

	t.Run("create quota item", func() {
		const mutation = `
            mutation CreateQuotaItem($input: CreateQuotaItemInput!) {
                createQuotaItem(input: $input) {
                    id
                    code
                    name
                    resourceType
                    unit
                    active
                }
            }
        `

		variables := map[string]interface{}{
			"input": map[string]interface{}{
				"code":         "api-calls",
				"name":         "API调用次数",
				"resourceType": "number",
				"unit":         "次/天",
				"active":       true,
			},
		}

		var resp struct {
			CreateQuotaItem struct {
				ID           string
				Code         string
				Name         string
				ResourceType string
				Unit         string
				Active       bool
			}
		}

		err := t.gqlClient.Post(mutation, &resp, client.Var("input", variables["input"]))
		t.Require().NoError(err)
		t.Require().Equal("api-calls", resp.CreateQuotaItem.Code)
	})
}

// Test_AppPolicyView 测试应用策略视图
func (t *graphqlSuite) Test_AppPolicyView() {
	userID := 1
	tenantID := 1
	appID := 10
	ctx := security.WithContext(context.Background(), security.NewGenericPrincipalByClaims(jwt.MapClaims{"sub": strconv.Itoa(userID)}))
	ctx = identity.WithTenantID(ctx, tenantID)

	// 创建应用
	ac := "resource1"
	t.Client.App.Create().SetID(appID).SetName("资源权限管理1").SetCode(ac).SetKind(app.KindWeb).
		SetComments("资源权限管理是管理组织目录中的应用,组织,人员以及授权信息").SetStatus(typex.SimpleStatusActive).
		SetCreatedBy(userID).SetOwnerOrgID(userID).ExecX(ctx)
	// 创建action
	aaCreates := make([]*ent.AppActionCreate, 0)
	for i := 5; i < 10; i++ {
		aaCreate := t.Client.AppAction.Create().SetID(i).SetCreatedBy(userID).
			SetAppID(appID).SetName("action" + strconv.Itoa(i)).SetComments("权限" + strconv.Itoa(i)).SetKind(appaction.KindGraphql).SetMethod(appaction.MethodRead)
		aaCreates = append(aaCreates, aaCreate)
	}
	t.Client.AppAction.CreateBulk(aaCreates...).ExecX(ctx)
	// 创建应用权限策略、策略视图
	apCreates := make([]*ent.AppPolicyCreate, 0)
	apvCreates := make([]*ent.AppPolicyViewCreate, 0)
	for i := 5; i < 10; i++ {
		// 权限策略
		apCreate := t.Client.AppPolicy.Create().SetID(i).SetKind(apppolicy.KindView).SetCreatedBy(userID).SetName("权限策略" + strconv.Itoa(i)).SetAppID(appID).SetStatus(typex.SimpleStatusActive).
			SetRules([]*types.PolicyRule{
				{
					Effect: types.PolicyEffectAllow,
					Actions: []string{
						"action" + strconv.Itoa(i),
					},
				},
			})
		apCreates = append(apCreates, apCreate)

		// 视图目录
		apvParentCreate := t.Client.AppPolicyView.Create().SetID(i + 10).SetAppID(appID).SetKind(apppolicyview.KindDir).SetName("视图目录" + strconv.Itoa(i)).SetParentID(0)
		// 视图策略
		apvCreate := t.Client.AppPolicyView.Create().SetID(i).SetAppID(appID).SetAppPolicyID(i).SetKind(apppolicyview.KindPolicy).SetName("视图策略" + strconv.Itoa(i)).SetParentID(i + 10)
		apvCreates = append(apvCreates, apvCreate, apvParentCreate)
	}
	t.Client.AppPolicy.CreateBulk(apCreates...).ExecX(ctx)
	t.Client.AppPolicyView.CreateBulk(apvCreates...).ExecX(ctx)
	// 创建应用角色
	arCreates := make([]*ent.AppRoleCreate, 0)
	for i := 5; i < 10; i++ {
		arCreate := t.Client.AppRole.Create().SetID(i).SetCreatedBy(userID).SetAppID(appID).SetName("角色" + strconv.Itoa(i)).SetComments("角色" + strconv.Itoa(i)).SetAutoGrant(true).SetEditable(true)
		arCreates = append(arCreates, arCreate)
	}
	ctx = ent.NewContext(ctx, t.Client)
	t.Client.AppRole.CreateBulk(arCreates...).ExecX(ctx)

	// 角色授权策略视图
	assignRoleID := 5
	assignAppPolicyIDs := []int{5, 6, 7}
	revokeAppPolicyIDs := []int{7}
	has, err := t.mr.AssignAppRolePolicyView(ctx, appID, assignRoleID, assignAppPolicyIDs, nil)
	t.Require().NoError(err)
	t.Require().True(has)
	// 获取角色策略视图验证
	appRoleViews, err := t.qr.AppPolicyView(ctx, ac)
	t.Require().NoError(err)
	apvResolver := &appPolicyViewResolver{
		Resolver: t.server.resolver,
	}
	// 验证角色5授权的应用视图
	appRoleAssignedNums := 0
	for _, appRoleView := range appRoleViews {
		has, err = apvResolver.AppRoleAssigned(ctx, appRoleView, assignRoleID)
		t.Require().NoError(err)
		if has {
			appRoleAssignedNums += 1
		}
	}
	if appRoleAssignedNums != len(assignAppPolicyIDs) {
		t.Require().Fail("角色授权策略视图失败")
	}
	// 取消授权策略视图7
	has, err = t.mr.AssignAppRolePolicyView(ctx, appID, assignRoleID, nil, revokeAppPolicyIDs)
	t.Require().NoError(err)
	t.Require().True(has)
	hasPolicy7 := false
	for _, appRoleView := range appRoleViews {
		has, err = apvResolver.AppRoleAssigned(ctx, appRoleView, assignRoleID)
		t.Require().NoError(err)
		if has {
			if appRoleView.PolicyID == 7 {
				hasPolicy7 = true
				break
			}
		}
	}
	if hasPolicy7 {
		t.Require().Fail("取消授权策略视图7失败")
	}

	// 应用授权组织
	has, err = t.mr.AssignOrganizationApp(ctx, tenantID, appID)
	t.Require().NoError(err)
	t.Require().True(has)

	// 获取组织权限策略视图
	orgPolicyViews, err := t.qr.OrgPolicyView(ctx, ac)
	t.Require().NoError(err)
	orgPolicyViewPolicies := 0
	// 组织策略
	orgPolicys := make([]*ent.OrgPolicy, 0)
	for _, opv := range orgPolicyViews {
		if opv.Kind == apppolicyview.KindPolicy {
			op, err := apvResolver.OrgPolicy(ctx, opv)
			t.Require().NoError(err)
			orgPolicys = append(orgPolicys, op)
		}
		for _, pid := range assignAppPolicyIDs {
			if opv.PolicyID == pid {
				orgPolicyViewPolicies += 1
				break
			}
		}
	}
	// 由于取消了一个授权，所以判断的时候+1
	if orgPolicyViewPolicies+1 != len(assignAppPolicyIDs) {
		t.Require().Fail("策略视图授权组织管理员错误")
	}

	// 给组织用户授权策略视图
	assignUserID := 2
	has, err = t.mr.AssignUserPolicyView(ctx, tenantID, assignUserID, []int{orgPolicys[0].ID, orgPolicys[1].ID}, nil)
	t.Require().NoError(err)
	t.Require().True(has)
	// 用户权限策略视图判断是否选中
	has, err = apvResolver.OrgUserAssigned(ctx, orgPolicyViews[0], assignUserID)
	t.Require().NoError(err)
	t.Require().True(has)
	has, err = apvResolver.OrgUserAssigned(ctx, orgPolicyViews[1], assignUserID)
	t.Require().NoError(err)
	t.Require().True(has)
	// 取消授权用户策略视图
	has, err = t.mr.AssignUserPolicyView(ctx, tenantID, assignUserID, nil, []int{orgPolicys[0].ID})
	t.Require().NoError(err)
	t.Require().True(has)
	// 判断orgPolicys[0]是否解除授权
	has, err = apvResolver.OrgUserAssigned(ctx, orgPolicyViews[0], assignUserID)
	t.Require().NoError(err)
	t.Require().False(has)

	// 创建用户组
	orgRole, err := t.mr.CreateRole(ctx, ent.CreateOrgRoleInput{
		Kind:  orgrole.KindGroup,
		OrgID: &tenantID,
		Name:  "用户组1",
	})
	t.Require().NoError(err)
	// 给用户组授权策略视图
	has, err = t.mr.AssignOrgRolePolicyView(ctx, tenantID, orgRole.ID, []int{orgPolicys[1].ID}, nil)
	t.Require().NoError(err)
	t.Require().True(has)
	// 用户组/角色权限策略视图判断是否选中
	has, _ = apvResolver.OrgRoleAssigned(ctx, orgPolicyViews[0], orgRole.ID)
	t.Require().False(has)
	has, err = apvResolver.OrgRoleAssigned(ctx, orgPolicyViews[1], orgRole.ID)
	t.Require().NoError(err)
	t.Require().True(has)
	// 取消授权用户组策略视图
	has, err = t.mr.AssignOrgRolePolicyView(ctx, tenantID, orgRole.ID, nil, []int{orgPolicys[1].ID})
	t.Require().NoError(err)
	t.Require().True(has)
	// 判断orgPolicys[1]是否解除授权
	has, err = apvResolver.OrgRoleAssigned(ctx, orgPolicyViews[1], orgRole.ID)
	t.Require().NoError(err)
	t.Require().False(has)
}
