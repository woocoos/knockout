package testinit

import (
	"context"
	"testing"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	schemahook "github.com/woocoos/knockout/codegen/entgen/hook"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/country"
	"github.com/woocoos/knockout/ent/migrate"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/region"
	"github.com/woocoos/knockout/ent/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/woocoos/knockout/ent/runtime"
)

// setupSqliteDB 创建SQLite内存数据库并初始化schema, 返回可用于验证数据的client.
func setupSqliteDB(t *testing.T) *ent.Client {
	t.Helper()
	dsn := "file:initbase_test?mode=memory&cache=shared&_fk=1"

	drv, err := sql.Open("sqlite3", dsn)
	require.NoError(t, err)

	client := ent.NewClient(ent.Driver(drv))
	schemahook.RegisterAllHooks(client)

	err = client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false))
	require.NoError(t, err)

	// 创建casbin schema
	casbinClient := casbinent.NewClient(casbinent.Driver(drv))
	err = casbinClient.Schema.Create(context.Background(),
		schema.WithForeignKeys(false))
	require.NoError(t, err)

	return client
}

func TestInitBase_WithSqlite(t *testing.T) {
	client := setupSqliteDB(t)
	defer client.Close()

	dsn := "file:initbase_test?mode=memory&cache=shared&_fk=1"
	require.NotPanics(t, func() {
		InitBase("sqlite3", dsn)
	})

	ctx := context.Background()

	t.Run("Users", func(t *testing.T) {
		users := client.User.Query().AllX(ctx)
		assert.Len(t, users, 3)

		admin := client.User.GetX(ctx, 1)
		assert.Equal(t, "admin", admin.PrincipalName)
		assert.Equal(t, "admin", admin.DisplayName)
		assert.Equal(t, user.UserTypeAccount, admin.UserType)

		user2 := client.User.GetX(ctx, 2)
		assert.Equal(t, "user2", user2.PrincipalName)
	})

	t.Run("UserAddrs", func(t *testing.T) {
		addrs := client.UserAddr.Query().AllX(ctx)
		assert.Len(t, addrs, 3)

		adminAddr := client.UserAddr.GetX(ctx, 1)
		assert.Equal(t, "admin@localhost", adminAddr.Email)

		user2Addr := client.UserAddr.GetX(ctx, 2)
		assert.Equal(t, "user2@localhost", user2Addr.Email)
	})

	t.Run("UserLoginProfiles", func(t *testing.T) {
		profiles := client.UserLoginProfile.Query().AllX(ctx)
		assert.Len(t, profiles, 3)

		p1 := client.UserLoginProfile.GetX(ctx, 1)
		assert.True(t, p1.CanLogin)
		assert.Equal(t, "UWZLIIUMPX53NYXB", p1.MfaSecret)
	})

	t.Run("UserPasswords", func(t *testing.T) {
		passwords := client.UserPassword.Query().AllX(ctx)
		assert.Len(t, passwords, 3)
	})

	t.Run("UserIdentities", func(t *testing.T) {
		identities := client.UserIdentity.Query().AllX(ctx)
		assert.Len(t, identities, 3)

		adminId := client.UserIdentity.GetX(ctx, 1)
		assert.Equal(t, "admin", adminId.Code)

		user2Id := client.UserIdentity.GetX(ctx, 2)
		assert.Equal(t, "user2", user2Id.Code)
	})

	t.Run("App", func(t *testing.T) {
		apps := client.App.Query().AllX(ctx)
		assert.Len(t, apps, 1)

		ap := client.App.GetX(ctx, 1)
		assert.Equal(t, "resource", ap.Code)
		assert.Equal(t, "资源权限管理", ap.Name)
	})

	t.Run("AppActions", func(t *testing.T) {
		actions := client.AppAction.Query().AllX(ctx)
		assert.Len(t, actions, 103) // 1(login) + 102(InitResourcePolicy所需)

		action := client.AppAction.GetX(ctx, 1)
		assert.Equal(t, "login", action.Name)
		assert.Equal(t, 1, action.AppID)
	})

	t.Run("AppRoles", func(t *testing.T) {
		roles := client.AppRole.Query().AllX(ctx)
		assert.Len(t, roles, 1)

		role := client.AppRole.GetX(ctx, 1)
		assert.Equal(t, "管理员", role.Name)
		assert.True(t, role.AutoGrant)
	})

	t.Run("AppPolicies", func(t *testing.T) {
		policies := client.AppPolicy.Query().AllX(ctx)
		assert.Len(t, policies, 1)

		policy := client.AppPolicy.GetX(ctx, 1)
		assert.Equal(t, "全部管理权限", policy.Name)
		assert.Len(t, policy.Rules, 1)
		assert.Equal(t, types.PolicyEffectAllow, policy.Rules[0].Effect)
	})

	t.Run("AppMenus", func(t *testing.T) {
		menus := client.AppMenu.Query().AllX(ctx)
		assert.Len(t, menus, 1)

		menu := client.AppMenu.GetX(ctx, 1)
		assert.Equal(t, "应用入口", menu.Name)
	})

	t.Run("Orgs", func(t *testing.T) {
		orgs := client.Org.Query().AllX(ctx)
		assert.Len(t, orgs, 3)

		org1 := client.Org.GetX(ctx, 1)
		assert.Equal(t, org.KindRoot, org1.Kind)
		assert.Equal(t, "woocoo.com", org1.Domain)
		require.NotNil(t, org1.OwnerID)
		assert.Equal(t, 1, *org1.OwnerID)
		assert.NotEmpty(t, org1.Path, "path应由hook自动计算")

		org2 := client.Org.GetX(ctx, 2)
		assert.Equal(t, org.KindOrganization, org2.Kind)
		assert.NotEmpty(t, org2.Path, "path应由hook自动计算")

		org3 := client.Org.GetX(ctx, 3)
		assert.NotEmpty(t, org3.Path)
	})

	t.Run("OrgUsers", func(t *testing.T) {
		orgUsers := client.OrgUser.Query().AllX(ctx)
		// i=1: org1-user1
		// i=2: org1-user2, org2-user2
		// i=3: org1-user3, org3-user3
		assert.Len(t, orgUsers, 5)

		// org1应有3个成员
		count := client.OrgUser.Query().Where(
			func(s *sql.Selector) {
				s.Where(sql.EQ("org_id", 1))
			},
		).CountX(ctx)
		assert.Equal(t, 3, count)
	})

	t.Run("OrgApps", func(t *testing.T) {
		orgApps := client.OrgApp.Query().AllX(ctx)
		assert.Len(t, orgApps, 1)

		oa := client.OrgApp.GetX(ctx, 1)
		assert.Equal(t, 1, oa.OrgID)
		assert.Equal(t, 1, oa.AppID)
	})

	t.Run("OrgRoles", func(t *testing.T) {
		roles := client.OrgRole.Query().AllX(ctx)
		assert.Len(t, roles, 2)

		role1 := client.OrgRole.GetX(ctx, 1)
		assert.Equal(t, "管理员", role1.Name)
		assert.Equal(t, 1, role1.AppRoleID)

		role2 := client.OrgRole.GetX(ctx, 2)
		assert.Equal(t, "administrators", role2.Name)
	})

	t.Run("Permissions", func(t *testing.T) {
		perms := client.Permission.Query().AllX(ctx)
		assert.Len(t, perms, 1)
	})

	t.Run("FileSources", func(t *testing.T) {
		sources := client.FileSource.Query().AllX(ctx)
		assert.Len(t, sources, 1)

		fs := client.FileSource.GetX(ctx, 1)
		assert.Equal(t, "woocootest", fs.Bucket)
	})

	t.Run("FileIdentities", func(t *testing.T) {
		ids := client.FileIdentity.Query().AllX(ctx)
		assert.Len(t, ids, 1)
	})

	t.Run("OauthClients", func(t *testing.T) {
		clients := client.OauthClient.Query().AllX(ctx)
		assert.Len(t, clients, 1)

		oc := client.OauthClient.GetX(ctx, 1)
		assert.Equal(t, "206734260394752", oc.ClientID)
	})

	t.Run("Countries", func(t *testing.T) {
		countries := client.Country.Query().AllX(ctx)
		assert.Len(t, countries, 2)

		cn := client.Country.Query().Where(country.Code("86")).OnlyX(ctx)
		assert.Equal(t, "中国", cn.Name)

		hk := client.Country.Query().Where(country.Code("852")).OnlyX(ctx)
		assert.Equal(t, "中国香港", hk.Name)
	})

	t.Run("Regions", func(t *testing.T) {
		regions := client.Region.Query().AllX(ctx)
		assert.Len(t, regions, 8)

		// 验证中国的region层级
		cnRegions := client.Region.Query().Where(region.CountryID(1)).AllX(ctx)
		assert.Len(t, cnRegions, 4)

		// 验证香港的region层级
		hkRegions := client.Region.Query().Where(region.CountryID(2)).AllX(ctx)
		assert.Len(t, hkRegions, 4)
	})

	t.Run("AppDicts", func(t *testing.T) {
		dicts := client.AppDict.Query().AllX(ctx)
		assert.Len(t, dicts, 1)

		d := client.AppDict.GetX(ctx, 1)
		assert.Equal(t, "DLSH", d.Code)
	})

	t.Run("AppDictItems", func(t *testing.T) {
		items := client.AppDictItem.Query().AllX(ctx)
		assert.Len(t, items, 4)
	})

	t.Run("CasbinRules", func(t *testing.T) {
		// InitBase通过独立的casbin连接创建规则, 需要重新连接验证
		dsn := "file:initbase_test?mode=memory&cache=shared&_fk=1"
		drv, err := sql.Open("sqlite3", dsn)
		require.NoError(t, err)
		casbinClient := casbinent.NewClient(casbinent.Driver(drv))
		defer casbinClient.Close()

		rules := casbinClient.CasbinRule.Query().AllX(ctx)
		assert.Len(t, rules, 2)

		// 验证g规则(角色分配)
		gRule := casbinClient.CasbinRule.Query().Where(
			func(s *sql.Selector) {
				s.Where(sql.EQ("ptype", "g"))
			},
		).OnlyX(ctx)
		assert.Equal(t, "1", gRule.V0)
		assert.Equal(t, "1", gRule.V1)

		// 验证p规则(权限策略)
		pRule := casbinClient.CasbinRule.Query().Where(
			func(s *sql.Selector) {
				s.Where(sql.EQ("ptype", "p"))
			},
		).OnlyX(ctx)
		assert.Equal(t, "1", pRule.V0)
		assert.Equal(t, "resource:*", pRule.V2)
		assert.Equal(t, "allow", pRule.V4)
	})
}

func TestInitResourcePolicy_WithSqlite(t *testing.T) {
	dsn := "file:initpolicy_test?mode=memory&cache=shared&_fk=1"

	// 创建schema
	drv, err := sql.Open("sqlite3", dsn)
	require.NoError(t, err)
	client := ent.NewClient(ent.Driver(drv))
	defer client.Close()
	schemahook.RegisterAllHooks(client)

	err = client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false))
	require.NoError(t, err)

	casbinClient := casbinent.NewClient(casbinent.Driver(drv))
	err = casbinClient.Schema.Create(context.Background(),
		schema.WithForeignKeys(false))
	require.NoError(t, err)

	require.NotPanics(t, func() {
		InitBase("sqlite3", dsn)
	})

	ctx := context.Background()
	tx, err := client.Tx(ctx)
	require.NoError(t, err)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// InitResourcePolicy创建细粒度策略, 需要对应的AppAction存在
	// 当前initApp只创建了"login"一个action, 而InitResourcePolicy引用了大量不存在的action
	// 这会导致AppPolicyRulesHook验证失败
	InitResourcePolicy(tx)

	err = tx.Commit()
	require.NoError(t, err)

	// 验证策略数量: initApp创建1个 + InitResourcePolicy创建多个
	policies := client.AppPolicy.Query().AllX(ctx)
	assert.Greater(t, len(policies), 1, "InitResourcePolicy应创建额外的策略")
}
