package graphql

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	"github.com/woocoos/knockout/api/graphql/model"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/service/resource"
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
	sec "github.com/woocoos/knockout/security"
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
	kosdk, err := api.NewSDK(t.Cnf.Sub("kosdk"))
	if err != nil {
		panic(err)
	}
	//redisClient ,err := gredis.NewClient(t.Cnf.Sub("redis"))
	t.server = NewServer(t.Cnf, WithPortalDB(t.CacheClient), WithCasbinDB(t.AuthDbClient), WithKOSdk(kosdk))
	t.mr = &mutationResolver{
		Resolver: t.server.resolver,
	}
	t.qr = &queryResolver{
		Resolver: t.server.resolver,
	}
	err = fmterr.InitErrorHandler(t.Cnf.Sub("errors"))
	t.Require().NoError(err)
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

func (t *graphqlSuite) TestOrganization() {
	t.Run("query", func() {
		ctx := t.NewTestCtx(1, 1)
		orgs, err := t.qr.Organizations(ctx, nil, nil, nil, nil, nil, nil)
		t.Require().NoError(err)
		t.Require().NotNil(orgs)
	})
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

func (t *graphqlSuite) TestGqlTypeQuery() {
	const query = `
query AppKind{
  __type(name: "AppKind"){
    name,
    enumValues{
      name, 
      description
    }
  }
}
`
	var resp map[string]any
	err := t.gqlClient.Post(query, &resp)
	t.Require().NoError(err)
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

	t.Client.Quota.Create().SetID(1).SetTenantID(1).SetQuotaItem(quotaItems[0]).SetLimit(100).SetStartAt(time.Now()).SetCreatedBy(1).SaveX(ctx)
	t.Client.Quota.Create().SetID(2).SetTenantID(1).SetQuotaItem(quotaItems[1]).SetLimit(1000).SetStartAt(time.Now().AddDate(0, 0, -1)).
		SetEndAt(time.Now().AddDate(0, 1, 0)).SetCreatedBy(1).SaveX(ctx)
	t.Client.Quota.Create().SetID(3).SetTenantID(1).SetQuotaItem(quotaItems[2]).SetLimit(1).SetStartAt(time.Now()).SetCreatedBy(1).SaveX(ctx)
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
	t.Run("create quota not tenant_id and org_id", func() {
		_, err := t.mr.client.Quota.Create().SetID(10).SetStartAt(time.Now()).SetCreatedBy(1).Save(ctx)
		t.Require().Error(err)
	})
}

func (t *graphqlSuite) TestAppPolicyRulesCache() {
	ctx := context.Background()
	t.Client.AppPolicy.Create()
	ac := "resource"
	appID := 1
	aas := []string{"userPermissions", "userMenus", "userRootOrgs"}
	// 创建action
	aaCreates := make([]*ent.AppActionCreate, 0)
	for i, a := range aas {
		create := t.Client.AppAction.Create().SetAppID(appID).SetCreatedBy(1).
			SetName(a).SetKind(appaction.KindFunction).SetComments("策略描述" + strconv.Itoa(i)).SetMethod(appaction.MethodRead)
		aaCreates = append(aaCreates, create)
	}
	err := t.Client.AppAction.CreateBulk(aaCreates...).Exec(ctx)
	t.Require().NoError(err)
	// 创建appPolicy
	ap, err := t.Client.AppPolicy.Create().SetCreatedBy(1).SetVersion("1").SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				ac + ":userPermissions",
			},
		},
	}).SetName("KOResAccess").SetComments("资源权限管理应用授权").SetAppID(appID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true).Save(ctx)
	t.Require().NoError(err)
	// base64加密
	id := fmt.Sprintf("app_policy:%d", ap.ID)
	gid := base64.StdEncoding.EncodeToString([]byte(id))
	var nodeQuery = `
            query appPolicyInfo {
			  node(id: "` + gid + `") {
				... on AppPolicy {
				  comments
				  rules {
					effect
					actions
				  }
				}
			  }
			}
        `
	var nodeResp struct {
		Node struct {
			Comments string
			Rules    []struct {
				Effect  string
				Actions []string
			}
		}
	}
	t.Run("query appPolicy", func() {
		err := t.gqlClient.Post(nodeQuery, &nodeResp)
		t.Require().NoError(err)
		t.Require().Len(nodeResp.Node.Rules[0].Actions, 1)
	})
	t.Run("query appPolicy after update", func() {
		const query = `
            mutation UpdateAppPolicy ($policyID: ID!, $input: UpdateAppPolicyInput!) {
			  updateAppPolicy(policyID: $policyID, input: $input) {
				comments
				rules {
				  effect
				  actions
				}
			  }
			}
        `
		variables := map[string]interface{}{
			"input": map[string]interface{}{
				"rules": []interface{}{
					map[string]interface{}{
						"effect": "allow",
						"actions": []string{
							ac + ":userPermissions",
							ac + ":userMenus",
							ac + ":userRootOrgs",
						},
					},
				},
			},
			"policyID": strconv.Itoa(ap.ID),
		}
		var resp struct {
			UpdateAppPolicy struct {
				Comments string
				Rules    []struct {
					Effect  string
					Actions []string
				}
			}
		}

		err = t.gqlClient.Post(query, &resp, client.Var("input", variables["input"]), client.Var("policyID", variables["policyID"]))
		t.Require().NoError(err)
		t.Require().Len(resp.UpdateAppPolicy.Rules[0].Actions, 3)
		time.Sleep(time.Second * 5)
		err := t.gqlClient.Post(nodeQuery, &nodeResp)
		t.Require().NoError(err)
		t.Require().Len(nodeResp.Node.Rules[0].Actions, 3)
	})
}

// TestOrgRoleHook 测试组织角色Hook
func (t *graphqlSuite) TestOrgRoleHook() {
	uid := 1
	loginTid := 1
	otherTid := 100
	rootRoleID := 2
	otherTidRoleID := 100
	loginTidRoleID := 99
	ctx := testsuite.NewTestCtx(1, 1, t.CacheClient)
	// 创建其他根组织
	err := t.Client.Org.Create().SetID(otherTid).SetKind(org.KindOrganization).SetParentID(0).SetStatus(typex.SimpleStatusActive).
		SetCreatedBy(1).SetUpdatedBy(1).SetName("org" + strconv.Itoa(otherTid)).Exec(ctx)
	// 登录组织创建角色
	err = t.Client.OrgRole.Create().SetID(loginTidRoleID).SetOrgID(loginTid).SetName("测试角色99").
		SetCreatedBy(uid).SetKind(orgrole.KindRole).Exec(ctx)
	t.Require().NoError(err)
	// 其他根组织创建角色
	err = t.Client.OrgRole.Create().SetID(otherTidRoleID).SetOrgID(otherTid).SetName("测试角色100").
		SetCreatedBy(uid).SetKind(orgrole.KindRole).Exec(ctx)
	t.Require().NoError(err)

	t.Run("OrgTraverseFunc：query orgRoles with not assign administrators role", func() {
		first := 20
		kind := orgrole.KindRole
		ors, err := t.qr.OrgRoles(ctx, nil, &first, nil, nil, nil, &ent.OrgRoleWhereInput{
			Kind:  &kind,
			OrgID: &loginTid,
		})
		t.Require().NoError(err)
		// 当前登录组织，应该返回数据
		t.Require().Greater(len(ors.Edges), 0)
		ors2, err := t.qr.OrgRoles(ctx, nil, &first, nil, nil, nil, &ent.OrgRoleWhereInput{
			Kind:  &kind,
			OrgID: &otherTid,
		})
		t.Require().NoError(err)
		// 其他组织，没授权administrators，无法访问角色
		t.Require().Len(ors2.Edges, 0)
	})
	t.Run("MutationInAllowOrg：create orgRole with not assign administrators role", func() {
		// 未授权administrators角色，loginTid能创建角色，otherTid抛异常无法创建
		_, err = t.mr.CreateRole(ctx, ent.CreateOrgRoleInput{
			Kind:  orgrole.KindRole,
			Name:  "测试角色102",
			OrgID: &loginTid,
		})
		t.Require().NoError(err)
		_, err = t.mr.CreateRole(ctx, ent.CreateOrgRoleInput{
			Kind:  orgrole.KindRole,
			Name:  "测试角色103",
			OrgID: &otherTid,
		})
		t.Require().ErrorIs(err, sec.ErrTenantIDNotAllow)
	})
	t.Run("MutationInAllowOrg：update orgRole with not assign administrators role", func() {
		// 未授权administrators角色，loginTid能更新角色与otherTid不能更新角色
		comments := "测试更新"
		_, err = t.mr.UpdateRole(ctx, loginTidRoleID, ent.UpdateOrgRoleInput{
			Comments: &comments,
		})
		t.Require().NoError(err)
		_, err = t.mr.UpdateRole(ctx, otherTidRoleID, ent.UpdateOrgRoleInput{
			Comments: &comments,
		})
		t.Require().ErrorIs(err, sec.ErrTenantIDNotAllow)
	})
	t.Run("MutationInAllowOrg：delete orgRole with not assign administrators role", func() {
		// 未授权administrators角色，loginTidRoleID允许删除，otherTidRoleID不允许删除
		_, err = t.mr.DeleteRole(ctx, loginTidRoleID)
		t.Require().NoError(err)
		_, err = t.mr.DeleteRole(ctx, otherTidRoleID)
		t.Require().True(ent.IsNotFound(err))
		// 恢复登录组织角色
		err = t.Client.OrgRole.Create().SetID(loginTidRoleID).SetOrgID(loginTid).SetName("测试角色99").
			SetCreatedBy(uid).SetKind(orgrole.KindRole).Exec(ctx)
		t.Require().NoError(err)
	})

	// 给root角色添加用户
	has, err := t.mr.AssignRoleUser(ctx, model.AssignRoleUserInput{
		OrgRoleID: rootRoleID,
		UserID:    uid,
	})
	t.Require().NoError(err)
	t.Require().True(has)

	t.Run("OrgTraverseFunc：query orgRoles with has assign administrators role", func() {
		first := 20
		kind := orgrole.KindRole
		ors, err := t.qr.OrgRoles(ctx, nil, &first, nil, nil, nil, &ent.OrgRoleWhereInput{
			Kind:  &kind,
			OrgID: &loginTid,
		})
		t.Require().NoError(err)
		// 当前登录组织，返回数据
		t.Require().Greater(len(ors.Edges), 0)
		ors2, err := t.qr.OrgRoles(ctx, nil, &first, nil, nil, nil, &ent.OrgRoleWhereInput{
			Kind:  &kind,
			OrgID: &otherTid,
		})
		t.Require().NoError(err)
		// 其他组织，授权了administrators，返回角色
		t.Require().Greater(len(ors2.Edges), 0)
	})
	t.Run("MutationInAllowOrg：create orgRole with has assign administrators role", func() {
		// 授权administrators角色，loginTid与otherTid都能创建角色
		_, err = t.mr.CreateRole(ctx, ent.CreateOrgRoleInput{
			Kind:  orgrole.KindRole,
			Name:  "测试角色104",
			OrgID: &loginTid,
		})
		t.Require().NoError(err)
		_, err = t.mr.CreateRole(ctx, ent.CreateOrgRoleInput{
			Kind:  orgrole.KindRole,
			Name:  "测试角色105",
			OrgID: &otherTid,
		})
		t.Require().NoError(err)
	})
	t.Run("MutationInAllowOrg：update orgRole with has assign administrators role", func() {
		// 授权administrators角色，loginTid与otherTid都能更新角色
		comments := "测试更新"
		_, err = t.mr.UpdateRole(ctx, loginTidRoleID, ent.UpdateOrgRoleInput{
			Comments: &comments,
		})
		t.Require().NoError(err)
		_, err = t.mr.UpdateRole(ctx, otherTidRoleID, ent.UpdateOrgRoleInput{
			Comments: &comments,
		})
		t.Require().NoError(err)
	})
	t.Run("MutationInAllowOrg：delete orgRole with has assign administrators role", func() {
		// 授权administrators角色，otherTidRoleID允许被删除
		_, err = t.mr.DeleteRole(ctx, loginTidRoleID)
		t.Require().NoError(err)
		_, err = t.mr.DeleteRole(ctx, otherTidRoleID)
		t.Require().NoError(err)
	})
}

func (t *graphqlSuite) TestFileIdentity() {
	ctx := testsuite.NewTestCtx(1, 1, t.CacheClient)
	fis, err := t.qr.OrgFileIdentities(ctx)
	t.Require().NoError(err)
	t.Equal(1, len(fis))
	t.Equal(1, fis[0].TenantID)
	// 测试子组织2没配置文件来源，取上级组织配置
	ctx2 := testsuite.NewTestCtx(1, 2, t.CacheClient)
	fis, err = t.qr.OrgFileIdentities(ctx2)
	t.Require().NoError(err)
	t.Equal(1, len(fis))
	t.Equal(1, fis[0].TenantID)
}

// TestSyncAppRoleToOrg 测试应用角色同步到组织，由于数据问题，子用例无法单独跑
func (t *graphqlSuite) TestSyncAppRoleToOrg() {
	ctx := t.NewTestCtx(1, 1)
	appID := 1
	orgID := 1
	appCode := "resource"
	// 应用权限
	ras := make([]*ent.AppActionCreate, 0)
	for i := 0; i < 9; i++ {
		ras = append(ras, t.Client.AppAction.Create().SetAppID(appID).SetCreatedBy(1).
			SetName(fmt.Sprintf("testAction%d", i)).SetKind(appaction.KindGraphql).SetComments("登陆授权").SetMethod(appaction.MethodRead),
		)
	}
	t.Client.AppAction.CreateBulk(ras...).ExecX(ctx)
	// 应用策略1
	ap1, err := t.Client.AppPolicy.Create().SetCreatedBy(1).SetVersion("1").SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				appCode + ":testAction1",
				appCode + ":testAction2",
				appCode + ":testAction3",
			},
		},
	}).SetName("ResourceTest1").SetAppID(appID).SetStatus(typex.SimpleStatusActive).Save(ctx)
	t.NoError(err)
	// 应用策略2
	ap2, err := t.Client.AppPolicy.Create().SetCreatedBy(1).SetVersion("1").SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				appCode + ":testAction3",
				appCode + ":testAction4",
			},
		},
	}).SetName("ResourceTest2").SetAppID(appID).SetStatus(typex.SimpleStatusActive).Save(ctx)
	t.NoError(err)
	// 应用角色
	ar, err := t.Client.AppRole.Create().SetAppID(appID).SetCreatedBy(1).SetName("测试角色同步").
		SetComments("管理员角色").SetAutoGrant(true).SetEditable(true).Save(ctx)
	t.NoError(err)
	// 策略1给应用角色
	_, err = t.mr.AssignAppRolePolicy(ctx, appID, ar.ID, []int{ap1.ID})
	t.NoError(err)
	// 角色授权给组织
	_, err = t.mr.AssignOrganizationAppRole(ctx, orgID, ar.ID)
	t.NoError(err)
	// 查询授权的组织角色
	or, err := t.Client.OrgRole.Query().Where(orgrole.AppRoleID(ar.ID), orgrole.OrgID(orgID), orgrole.KindEQ(orgrole.KindRole)).Only(ctx)
	t.NoError(err)
	t.Run("应用角色添加策略2，并同步到组织", func() {
		// 应用角色添加策略2
		_, err = t.mr.AssignAppRolePolicy(ctx, appID, ar.ID, []int{ap2.ID})
		t.NoError(err)
		// 查询现有授权的ap1
		ps, err := t.Client.Permission.Query().Where(
			permission.PrincipalKindEQ(permission.PrincipalKindRole),
			permission.OrgID(orgID),
			permission.RoleID(or.ID),
		).WithOrgPolicy().All(ctx)
		t.NoError(err)
		if len(ps) != 1 || *ps[0].Edges.OrgPolicy.AppPolicyID != ap1.ID {
			t.Fail("策略ap1未正确授权到组织角色")
		}
		// 应用角色同步到组织
		_, err = t.mr.SyncAppRoleToOrg(ctx, orgID, ar.ID)
		t.NoError(err)
		// 查询应用角色授权了ap1、ap2
		ps, err = t.Client.Permission.Query().Where(
			permission.PrincipalKindEQ(permission.PrincipalKindRole),
			permission.OrgID(orgID),
			permission.RoleID(or.ID),
		).All(ctx)
		t.NoError(err)
		if len(ps) != 2 {
			t.Fail("策略ap1、ap2未正确授权到组织角色")
		}
	})
	var group *ent.OrgRole
	t.Run("组织角色策略授权给用户组", func() {
		// 创建用户组
		group, err = t.Client.OrgRole.Create().SetOrgID(1).SetName("测试用户组1").
			SetCreatedBy(1).SetKind(orgrole.KindGroup).Save(ctx)
		t.NoError(err)
		//
		ps, err := t.Client.Permission.Query().Where(
			permission.PrincipalKindEQ(permission.PrincipalKindRole),
			permission.OrgID(orgID),
			permission.RoleID(or.ID),
		).All(ctx)
		t.NoError(err)
		// 给用户组授权策略ap1、ap2
		_, err = t.mr.Grant(ctx, ent.CreatePermissionInput{
			PrincipalKind: permission.PrincipalKindRole,
			OrgID:         orgID,
			RoleID:        &group.ID,
			OrgPolicyID:   ps[0].OrgPolicyID,
		})
		t.NoError(err)
		_, err = t.mr.Grant(ctx, ent.CreatePermissionInput{
			PrincipalKind: permission.PrincipalKindRole,
			OrgID:         orgID,
			RoleID:        &group.ID,
			OrgPolicyID:   ps[1].OrgPolicyID,
		})
		t.NoError(err)
	})
	t.Run("应用角色移除策略1，并同步到组织", func() {
		// 应用角色移除策略1
		_, err = t.mr.RevokeAppRolePolicy(ctx, appID, ar.ID, []int{ap1.ID})
		t.NoError(err)
		// 应用角色同步到组织
		_, err = t.mr.SyncAppRoleToOrg(ctx, orgID, ar.ID)
		// 查询组织角色现有授权的ap2
		ps, err := t.Client.Permission.Query().Where(
			permission.PrincipalKindEQ(permission.PrincipalKindRole),
			permission.OrgID(orgID),
			permission.RoleID(or.ID),
		).WithOrgPolicy().All(ctx)
		t.NoError(err)
		if len(ps) != 1 || *ps[0].Edges.OrgPolicy.AppPolicyID != ap2.ID {
			t.Fail("组织角色策略同步失败")
			fmt.Println(len(ps))
		}
		// 查询用户组授权策略ap2
		ps, err = t.Client.Permission.Query().Where(
			permission.PrincipalKindEQ(permission.PrincipalKindRole),
			permission.OrgID(orgID),
			permission.RoleID(group.ID),
		).WithOrgPolicy().All(ctx)
		t.NoError(err)
		if len(ps) != 1 || *ps[0].Edges.OrgPolicy.AppPolicyID != ap2.ID {
			t.Fail("组织用户组同步策略失败")
			fmt.Println(len(ps))
		}
	})
}

func (t *graphqlSuite) TestAppDictByRefCode() {
	const query = `
query appDictByRefCode{
  appDictByRefCode(refCodes: ["resource:DLSH"]){
    id,code,name,items{
		id,code,name,refCode,orgID
	}
  }
}
`
	var resp struct {
		AppDictByRefCode []struct {
			ID    string
			Code  string
			Name  string
			Items []struct {
				ID      string
				Code    string
				OrgID   string
				Name    string
				RefCode string
			}
		}
	}
	err := t.gqlClient.Post(query, &resp)
	t.Require().NoError(err)
	t.Equal(4, len(resp.AppDictByRefCode[0].Items))
}

func (t *graphqlSuite) TestAppDictItemByRefCode() {
	const query = `
query appDictItemByRefCode{
  appDictItemByRefCode(refCode: "resource:DLSH"){
    id,code,name,refCode
  }
}
`
	var resp struct {
		AppDictItemByRefCode []struct {
			ID      string
			Code    string
			Name    string
			RefCode string
		}
	}
	err := t.gqlClient.Post(query, &resp)
	t.Require().NoError(err)
	t.Equal(2, len(resp.AppDictItemByRefCode))
}

// 修改密码移除其他登录的token
func (t *graphqlSuite) TestChangePassword() {
	// redis设置值
	//token:1:9af4ea59-7a31-4658-8a2e-4b0d4f849cda
	_ = t.Redis.Set("token:1:9af4ea59-7a31-4658-8a2e-4b0d4f849cda", "1")
	_ = t.Redis.Set("token:1:9af4ea59-7a31-4658-8a2e-4b0d4f849baa", "2")
	const query = `
mutation changePassword($oldPwd: String!,$newPwd: String!){
  changePassword(oldPwd: $oldPwd, newPwd: $newPwd)
}
`
	variables := map[string]interface{}{
		"oldPwd": resource.SHA256("123456"),
		"newPwd": resource.SHA256("1234567"),
	}
	var resp struct {
		ChangePassword bool
	}
	err := t.gqlClient.Post(query, &resp, client.Var("oldPwd", variables["oldPwd"]), client.Var("newPwd", variables["newPwd"]))
	t.Require().NoError(err)
	t.Equal(true, resp.ChangePassword)
	t.Equal(true, t.Redis.Exists("token:1:9af4ea59-7a31-4658-8a2e-4b0d4f849cda"))
	t.Equal(false, t.Redis.Exists("token:1:9af4ea59-7a31-4658-8a2e-4b0d4f849baa"))
}
