package graphql

import (
	"context"
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

	t.Client.Quota.Create().SetID(1).SetOrgID(1).SetQuotaItem(quotaItems[0]).SetLimit(100).SetStartAt(time.Now()).SetCreatedBy(1).SaveX(ctx)
	t.Client.Quota.Create().SetID(2).SetOrgID(1).SetQuotaItem(quotaItems[1]).SetLimit(1000).SetStartAt(time.Now().AddDate(0, 0, -1)).
		SetEndAt(time.Now().AddDate(0, 1, 0)).SetCreatedBy(1).SaveX(ctx)
	t.Client.Quota.Create().SetID(3).SetOrgID(1).SetQuotaItem(quotaItems[2]).SetLimit(1).SetStartAt(time.Now()).SetCreatedBy(1).SaveX(ctx)
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
                            org {
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
						StartAt   time.Time
						EndAt     *time.Time
						QuotaItem struct {
							Code string
							Name string
						}
						Org struct {
							Name string
						}
					}
				}
			}
		}

		err := t.gqlClient.Post(query, &resp)
		t.Require().NoError(err)
		t.Require().Len(resp.Quotas.Edges, 2)
	})
}
