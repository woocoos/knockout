// Package testinit provides test initialization data for knockout.
package testinit

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tsingsun/woocoo/pkg/security"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout/codegen/entgen/hook"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/useraddr"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
	"strconv"
)

type dataset struct {
	portal *ent.Client
	casbin *casbinent.Client

	casbinRules []*casbinent.CasbinRuleCreate
}

// InitBase init base data: app, file source, oauth client, user, org and their relationship.
//
// in sqlite,
func InitBase(name, dsn string) {
	drv, err := sql.Open(name, dsn)
	if err != nil {
		panic(err)
	}
	portal := ent.NewClient(ent.Driver(drv))
	hook.RegisterAllHooks(portal)
	casbin := casbinent.NewClient(casbinent.Driver(drv))
	InitBaseWithClients(portal, casbin)
}

// InitBaseWithClients 使用已有的 client 初始化基础数据.
// 适用于需要共享数据库连接的场景(如 SQLite 内存数据库).
func InitBaseWithClients(portal *ent.Client, casbin *casbinent.Client) {
	ds := dataset{
		portal:      portal,
		casbin:      casbin,
		casbinRules: make([]*casbinent.CasbinRuleCreate, 0),
	}

	ctx := context.Background()
	tx, err := ds.portal.Tx(ctx)
	if err != nil {
		panic(err)
	}
	casbinTx, err := ds.casbin.Tx(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			casbinTx.Rollback()
			panic(err)
		} else {
			tx.Commit()
			if len(ds.casbinRules) > 0 {
				casbinTx.CasbinRule.CreateBulk(ds.casbinRules...).ExecX(context.Background())
			}
			casbinTx.Commit()
		}
	}()
	// 初始化应用
	ds.initApp(tx, casbinTx)
	// 初始化文件存储
	ds.initFileSource(tx)
	// 初始化OauthClient
	ds.initOauthClient(tx)
	// 初始化用户
	ds.initUser(tx)
	// 初始化组织
	ds.initOrg(tx)
	// 初始化国家地区
	ds.initRegion(tx)
	// 初始化字典数据
	ds.initAppDict(tx)
}

func (*dataset) initOrg(client *ent.Tx) {
	ou := make([]*ent.OrgUserCreate, 0)
	for i := 1; i < 4; i++ {
		// 由于path字段是计算字段，所以这里不需要设置,但需要对org独立保存.
		c := client.Org.Create().SetID(i).SetKind(org.KindOrganization).SetParentID(i - 1).SetStatus(typex.SimpleStatusActive).
			SetCreatedBy(1).SetUpdatedBy(1).SetName("org" + strconv.Itoa(i))
		if i == 1 {
			c.SetKind(org.KindRoot).SetDomain("woocoo.com").SetOwnerID(1)
		}
		ctx := security.WithContext(context.Background(), security.NewGenericPrincipalByClaims(jwt.MapClaims{"sub": "1"}))
		if err := c.Exec(ctx); err != nil {
			panic(err)
		}

	}
	for i := 1; i < 4; i++ {
		if i != 1 { // 1为根目录,所有用户需要加入根目录
			u := client.OrgUser.Create().SetOrgID(1).SetUserID(i).SetCreatedBy(1).SetDisplayName("user" + strconv.Itoa(i))
			ou = append(ou, u)
		}
		u := client.OrgUser.Create().SetOrgID(i).SetUserID(i).SetCreatedBy(1).SetDisplayName("user" + strconv.Itoa(i))
		ou = append(ou, u)
	}
	err := client.OrgUser.CreateBulk(ou...).Exec(context.Background())
	if err != nil {
		panic(err)
	}
}

func (*dataset) initUser(client *ent.Tx) {
	ub := make([]*ent.UserCreate, 0)
	ua := make([]*ent.UserAddrCreate, 0)
	ulp := make([]*ent.UserLoginProfileCreate, 0)
	up := make([]*ent.UserPasswordCreate, 0)
	ui := make([]*ent.UserIdentityCreate, 0)
	for i := 1; i < 4; i++ {
		c := client.User.Create().SetID(i).SetUserType(user.UserTypeAccount).SetCreationType(user.CreationTypeManual).
			SetRegisterIP("").SetPrincipalName("user" + strconv.Itoa(i)).SetDisplayName("user" + strconv.Itoa(i)).
			SetStatus(types.UserStatusActive).SetCreatedBy(1)
		if i == 1 {
			c.SetPrincipalName("admin").SetDisplayName("admin")
		}
		ub = append(ub, c)

		a := client.UserAddr.Create().SetID(i).SetUserID(i).SetCreatedBy(1).SetAddrType(useraddr.AddrTypeContact).SetEmail("user" + strconv.Itoa(i) + "@localhost")
		if i == 1 {
			a.SetEmail("admin@localhost")
		}
		ua = append(ua, a)

		lp := client.UserLoginProfile.Create().SetID(i).SetUserID(i).SetCreatedBy(1).SetSetKind(userloginprofile.SetKindKeep).
			SetCanLogin(true).SetPasswordReset(false).SetMfaSecret("UWZLIIUMPX53NYXB").SetVerifyDevice(true)
		ulp = append(ulp, lp)

		p := client.UserPassword.Create().SetID(i).SetUserID(i).SetCreatedBy(1).SetScene(userpassword.SceneLogin).
			SetStatus(typex.SimpleStatusActive).SetPassword("123456").SetSalt("123456").SetPassword("9b1063951d443cfac15cc879efb4054f4f4fd599e1b1a9aee67b0301e19e40fe")
		up = append(up, p)

		id := client.UserIdentity.Create().SetID(i).SetUserID(i).SetCreatedBy(1).SetKind(useridentity.KindName).
			SetCode("user" + strconv.Itoa(i))
		if i == 1 {
			id.SetCode("admin")
		}
		ui = append(ui, id)
	}
	client.User.CreateBulk(ub...).ExecX(context.Background())
	client.UserAddr.CreateBulk(ua...).ExecX(context.Background())
	client.UserLoginProfile.CreateBulk(ulp...).ExecX(context.Background())
	client.UserPassword.CreateBulk(up...).ExecX(context.Background())
	client.UserIdentity.CreateBulk(ui...).ExecX(context.Background())
}

func (set *dataset) initApp(client *ent.Tx, casbinClient *casbinent.Tx) {
	apps := make([]*ent.AppCreate, 0)
	ars := make([]*ent.AppRoleCreate, 0)
	ras := make([]*ent.AppActionCreate, 0)
	aps := make([]*ent.AppPolicyCreate, 0)
	rps := make([]*ent.AppRolePolicyCreate, 0)
	ams := make([]*ent.AppMenuCreate, 0)
	oas := make([]*ent.OrgAppCreate, 0)
	ops := make([]*ent.OrgPolicyCreate, 0)
	ors := make([]*ent.OrgRoleCreate, 0)
	ps := make([]*ent.PermissionCreate, 0)
	orus := make([]*ent.OrgRoleUserCreate, 0)
	for i := 1; i < 2; i++ {
		ac := "resource"
		a := client.App.Create().SetID(i).SetName("资源权限管理").SetCode(ac).SetKind(app.KindWeb).
			SetComments("资源权限管理是管理组织目录中的应用,组织,人员以及授权信息").SetStatus(typex.SimpleStatusActive).
			SetCreatedBy(1).SetOwnerOrgID(1)
		apps = append(apps, a)

		ras = append(ras, client.AppAction.Create().SetID(i).SetAppID(i).SetCreatedBy(1).
			SetName("login").SetKind(appaction.KindFunction).SetComments("登陆授权").SetMethod(appaction.MethodRead),
		)

		ars = append(ars, client.AppRole.Create().SetID(i).SetAppID(i).SetCreatedBy(1).SetName("管理员").
			SetComments("管理员角色").SetAutoGrant(true).SetEditable(true),
		)
		aps = append(aps, client.AppPolicy.Create().SetID(i).SetAppID(i).SetCreatedBy(1).SetName("全部管理权限").
			SetComments("全部管理权限").SetAutoGrant(true).SetKind(apppolicy.KindApp).SetStatus(typex.SimpleStatusActive).SetVersion("V1").
			SetRules([]*types.PolicyRule{
				{
					Effect:    "allow",
					Actions:   []string{ac + ":*"},
					Resources: []string{},
				},
			}),
		)
		ams = append(ams, client.AppMenu.Create().SetID(i).SetAppID(i).SetCreatedBy(1).SetActionID(i).SetParentID(0).
			SetName("应用入口").SetKind(appmenu.KindMenu),
		)
		rps = append(rps, client.AppRolePolicy.Create().SetID(i).SetAppID(i).SetRoleID(i).SetPolicyID(i).SetCreatedBy(1))

		oas = append(oas, client.OrgApp.Create().SetID(i).SetOrgID(1).SetAppID(i).SetCreatedBy(1))

		ops = append(ops, client.OrgPolicy.Create().SetID(i).SetOrgID(1).SetAppID(i).SetAppPolicyID(i).
			SetCreatedBy(1).SetName("全部管理权限").SetRules([]*types.PolicyRule{
			{
				Effect:    "allow",
				Actions:   []string{ac + ":*"},
				Resources: []string{},
			},
		}))

		ors = append(ors, client.OrgRole.Create().SetID(i).SetOrgID(1).SetAppRoleID(i).SetName("管理员").
			SetCreatedBy(1).SetKind(orgrole.KindRole))
		ors = append(ors, client.OrgRole.Create().SetID(2).SetOrgID(1).SetName("administrators").
			SetCreatedBy(1).SetKind(orgrole.KindRole))

		ps = append(ps, client.Permission.Create().SetID(i).SetOrgID(1).SetOrgPolicyID(i).SetCreatedBy(1).
			SetPrincipalKind(permission.PrincipalKindRole).SetRoleID(i).SetStatus(typex.SimpleStatusActive))

		orus = append(orus, client.OrgRoleUser.Create().SetID(i).SetCreatedBy(1).SetOrgRoleID(i).SetOrgUserID(1).SetUserID(1).SetOrgID(1))

		set.casbinRules = append(set.casbinRules, casbinClient.CasbinRule.Create().SetPtype("g").
			SetV0("1").SetV1(strconv.Itoa(i)).SetV2("1"))
		set.casbinRules = append(set.casbinRules, casbinClient.CasbinRule.Create().SetPtype("p").
			SetV0(strconv.Itoa(i)).SetV1("1").SetV2(ac+":*").SetV3("read").SetV4("allow"))
	}
	// InitResourcePolicy所需的全部action (login已在循环中创建, ID=1)
	type actionDef struct {
		name   string
		kind   appaction.Kind
		method appaction.Method
	}
	resourceActions := []actionDef{
		// 路由类action
		{"/", appaction.KindRoute, appaction.MethodRead},
		{"/dict", appaction.KindRoute, appaction.MethodRead},
		{"/org/departments", appaction.KindRoute, appaction.MethodRead},
		{"/org/groups", appaction.KindRoute, appaction.MethodRead},
		{"/org/policys", appaction.KindRoute, appaction.MethodRead},
		{"/org/roles", appaction.KindRoute, appaction.MethodRead},
		{"/org/users", appaction.KindRoute, appaction.MethodRead},
		{"/system/account", appaction.KindRoute, appaction.MethodRead},
		{"/system/app", appaction.KindRoute, appaction.MethodRead},
		{"/system/file/source", appaction.KindRoute, appaction.MethodRead},
		{"/system/org", appaction.KindRoute, appaction.MethodRead},
		{"/user/info", appaction.KindRoute, appaction.MethodRead},
		{"/user/safety", appaction.KindRoute, appaction.MethodRead},
		// GraphQL查询类action (列表)
		{"organizations", appaction.KindGraphql, appaction.MethodList},
		{"users", appaction.KindGraphql, appaction.MethodList},
		{"apps", appaction.KindGraphql, appaction.MethodList},
		{"fileSources", appaction.KindGraphql, appaction.MethodList},
		{"appDicts", appaction.KindGraphql, appaction.MethodList},
		// GraphQL查询类action (单条/自定义)
		{"userPermissions", appaction.KindGraphql, appaction.MethodRead},
		{"userMenus", appaction.KindGraphql, appaction.MethodRead},
		{"userRootOrgs", appaction.KindGraphql, appaction.MethodRead},
		{"node", appaction.KindGraphql, appaction.MethodRead},
		{"appAccess", appaction.KindGraphql, appaction.MethodRead},
		{"orgRoles", appaction.KindGraphql, appaction.MethodRead},
		{"orgGroups", appaction.KindGraphql, appaction.MethodRead},
		{"orgUserPreference", appaction.KindGraphql, appaction.MethodRead},
		{"orgRecycleUsers", appaction.KindGraphql, appaction.MethodRead},
		{"userGroups", appaction.KindGraphql, appaction.MethodRead},
		{"userExtendGroupPolicies", appaction.KindGraphql, appaction.MethodRead},
		{"orgPolicyReferences", appaction.KindGraphql, appaction.MethodRead},
		{"orgAppActions", appaction.KindGraphql, appaction.MethodRead},
		{"orgRoleUsers", appaction.KindGraphql, appaction.MethodRead},
		{"appPolicyAssignedToOrgs", appaction.KindGraphql, appaction.MethodRead},
		{"appRoleAssignedToOrgs", appaction.KindGraphql, appaction.MethodRead},
		{"userApps", appaction.KindGraphql, appaction.MethodRead},
		// GraphQL变更类action
		{"changePassword", appaction.KindGraphql, appaction.MethodWrite},
		{"createOrganization", appaction.KindGraphql, appaction.MethodWrite},
		{"updateOrganization", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteOrganization", appaction.KindGraphql, appaction.MethodWrite},
		{"createOrganizationUser", appaction.KindGraphql, appaction.MethodWrite},
		{"moveOrganization", appaction.KindGraphql, appaction.MethodWrite},
		{"recoverOrgUser", appaction.KindGraphql, appaction.MethodWrite},
		{"updateUser", appaction.KindGraphql, appaction.MethodWrite},
		{"updateLoginProfile", appaction.KindGraphql, appaction.MethodWrite},
		{"bindUserIdentity", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteUserIdentity", appaction.KindGraphql, appaction.MethodWrite},
		{"sendMFAToUserByEmail", appaction.KindGraphql, appaction.MethodWrite},
		{"enableMFA", appaction.KindGraphql, appaction.MethodWrite},
		{"createOauthClient", appaction.KindGraphql, appaction.MethodWrite},
		{"enableOauthClient", appaction.KindGraphql, appaction.MethodWrite},
		{"disableOauthClient", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteOauthClient", appaction.KindGraphql, appaction.MethodWrite},
		{"grant", appaction.KindGraphql, appaction.MethodWrite},
		{"revoke", appaction.KindGraphql, appaction.MethodWrite},
		{"removeOrganizationUser", appaction.KindGraphql, appaction.MethodWrite},
		{"updateOrganizationPolicy", appaction.KindGraphql, appaction.MethodWrite},
		{"createOrganizationPolicy", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteOrganizationPolicy", appaction.KindGraphql, appaction.MethodWrite},
		{"assignRoleUser", appaction.KindGraphql, appaction.MethodWrite},
		{"createRole", appaction.KindGraphql, appaction.MethodWrite},
		{"updateRole", appaction.KindGraphql, appaction.MethodWrite},
		{"revokeRoleUser", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteRole", appaction.KindGraphql, appaction.MethodWrite},
		{"createRoot", appaction.KindGraphql, appaction.MethodWrite},
		{"revokeOrganizationApp", appaction.KindGraphql, appaction.MethodWrite},
		{"assignOrganizationApp", appaction.KindGraphql, appaction.MethodWrite},
		{"createOrganizationAccount", appaction.KindGraphql, appaction.MethodWrite},
		{"resetUserPasswordByEmail", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteUser", appaction.KindGraphql, appaction.MethodWrite},
		{"createApp", appaction.KindGraphql, appaction.MethodWrite},
		{"updateApp", appaction.KindGraphql, appaction.MethodWrite},
		{"createAppActions", appaction.KindGraphql, appaction.MethodWrite},
		{"updateAppAction", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteAppAction", appaction.KindGraphql, appaction.MethodWrite},
		{"createAppPolicy", appaction.KindGraphql, appaction.MethodWrite},
		{"updateAppPolicy", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteAppPolicy", appaction.KindGraphql, appaction.MethodWrite},
		{"assignOrganizationAppPolicy", appaction.KindGraphql, appaction.MethodWrite},
		{"revokeOrganizationAppPolicy", appaction.KindGraphql, appaction.MethodWrite},
		{"createAppMenus", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteAppMenu", appaction.KindGraphql, appaction.MethodWrite},
		{"updateAppMenu", appaction.KindGraphql, appaction.MethodWrite},
		{"moveAppMenu", appaction.KindGraphql, appaction.MethodWrite},
		{"createAppRole", appaction.KindGraphql, appaction.MethodWrite},
		{"updateAppRole", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteAppRole", appaction.KindGraphql, appaction.MethodWrite},
		{"assignAppRolePolicy", appaction.KindGraphql, appaction.MethodWrite},
		{"revokeAppRolePolicy", appaction.KindGraphql, appaction.MethodWrite},
		{"revokeOrganizationAppRole", appaction.KindGraphql, appaction.MethodWrite},
		{"assignOrganizationAppRole", appaction.KindGraphql, appaction.MethodWrite},
		{"updateAppRes", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteApp", appaction.KindGraphql, appaction.MethodWrite},
		{"createFileSource", appaction.KindGraphql, appaction.MethodWrite},
		{"updateFileSource", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteFileSource", appaction.KindGraphql, appaction.MethodWrite},
		{"updateAppDict", appaction.KindGraphql, appaction.MethodWrite},
		{"createAppDict", appaction.KindGraphql, appaction.MethodWrite},
		{"updateAppDictItem", appaction.KindGraphql, appaction.MethodWrite},
		{"createAppDictItem", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteAppDictItem", appaction.KindGraphql, appaction.MethodWrite},
		{"moveAppDictItem", appaction.KindGraphql, appaction.MethodWrite},
		{"deleteAppDict", appaction.KindGraphql, appaction.MethodWrite},
	}
	for idx, a := range resourceActions {
		ras = append(ras, client.AppAction.Create().
			SetID(idx + 2). // ID 1 是 login
			SetAppID(1).
			SetCreatedBy(1).
			SetName(a.name).
			SetKind(a.kind).
			SetMethod(a.method))
	}
	client.App.CreateBulk(apps...).ExecX(context.Background())
	client.AppAction.CreateBulk(ras...).ExecX(context.Background())
	client.AppRole.CreateBulk(ars...).ExecX(context.Background())
	client.AppPolicy.CreateBulk(aps...).ExecX(context.Background())
	client.AppRolePolicy.CreateBulk(rps...).ExecX(context.Background())
	client.AppMenu.CreateBulk(ams...).ExecX(context.Background())
	client.OrgApp.CreateBulk(oas...).ExecX(context.Background())
	client.OrgPolicy.CreateBulk(ops...).ExecX(context.Background())
	client.OrgRole.CreateBulk(ors...).ExecX(context.Background())
	client.Permission.CreateBulk(ps...).ExecX(context.Background())
	client.OrgRoleUser.CreateBulk(orus...).ExecX(context.Background())
}

func (*dataset) initFileSource(client *ent.Tx) {
	tenantID := 1
	fs := make([]*ent.FileSourceCreate, 0)
	s1 := client.FileSource.Create().SetID(1).SetKind(filesource.KindMinio).SetComments("本地存储bucket").
		SetEndpoint("http://192.168.0.17:32650").SetBucket("woocootest").SetRegion("minio").SetStsEndpoint("http://192.168.0.17:32650").
		SetBucketURL("http://192.168.0.17:32650/woocootest").SetCreatedBy(1)
	fs = append(fs, s1)
	client.FileSource.CreateBulk(fs...).ExecX(context.Background())

	fi := make([]*ent.FileIdentityCreate, 0)
	s2 := client.FileIdentity.Create().SetID(1).SetCreatedBy(1).SetFileSourceID(1).SetAccessKeyID("test").SetAccessKeySecret("test1234").
		SetIsDefault(true).SetDurationSeconds(3600).SetPolicy("").SetRoleArn("arn:aws:s3:::*").SetTenantID(tenantID)
	fi = append(fi, s2)
	client.FileIdentity.CreateBulk(fi...).ExecX(identity.WithTenantID(context.Background(), tenantID))
}

func (*dataset) initOauthClient(client *ent.Tx) {
	oc := make([]*ent.OauthClientCreate, 0)
	s1 := client.OauthClient.Create().SetID(1).SetName("系统").SetClientID("206734260394752").SetClientSecret("T2UlqISVFq4DR9InXamj3l74iWdu3Tyr").
		SetGrantTypes("client_credentials").SetStatus(typex.SimpleStatusActive).SetUserID(1).SetCreatedBy(1)
	oc = append(oc, s1)
	client.OauthClient.CreateBulk(oc...).ExecX(context.Background())
}

// InitResourcePolicy init resource policy.
func InitResourcePolicy(client *ent.Tx) {
	createBy := 0
	version := "1"
	ap, err := client.App.Query().Where(app.Code("resource")).Only(context.Background())
	if err != nil {
		panic(err)
	}
	aps := make([]*ent.AppPolicyCreate, 0)
	// 应用授权
	koResAccess := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:userPermissions", "resource:userMenus", "resource:userRootOrgs", "resource:node", "resource:login", "resource:appAccess", "resource:/", "resource:orgRoles", "resource:orgGroups", "resource:/user/info",
				"resource:updateUser", "resource:/user/safety", "resource:changePassword", "resource:userApps", "resource:orgUserPreference",
			},
		},
	}).SetName("KOResAccess").SetComments("资源权限管理应用授权，拥有该策略允许登录后台").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, koResAccess)
	// 组织协作-部门管理
	KOResDepartmentRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:/org/departments", "resource:organizations",
			},
		},
	}).SetName("KOResDepartmentRead").SetComments("资源权限管理部门管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDepartmentRead)
	KOResDepartmentEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:createOrganization", "resource:updateOrganization",
			},
		},
	}).SetName("KOResDepartmentEdit").SetComments("资源权限管理部门管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDepartmentEdit)
	KOResDepartmentDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:deleteOrganization",
			},
		},
	}).SetName("KOResDepartmentDel").SetComments("资源权限管理部门管理删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDepartmentDel)
	// 组织协作-用户管理
	KOResOrgUsersRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:/org/users", "resource:organizations", "resource:orgRecycleUsers", "resource:userGroups", "resource:orgGroups", "resource:userExtendGroupPolicies",
			},
		},
	}).SetName("KOResOrgUsersRead").SetComments("资源权限管理用户管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgUsersRead)
	KOResOrgUsersEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:createOrganizationUser", "resource:moveOrganization", "resource:recoverOrgUser", "resource:updateUser", "resource:updateLoginProfile", "resource:bindUserIdentity",
				"resource:deleteUserIdentity", "resource:sendMFAToUserByEmail", "resource:enableMFA", "resource:createOauthClient", "resource:enableOauthClient", "resource:disableOauthClient",
				"resource:deleteOauthClient", "resource:grant", "resource:revoke", "resource:removeOrganizationUser",
			},
		},
	}).SetName("KOResOrgUsersEdit").SetComments("资源权限管理用户管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgUsersEdit)
	// 组织协作-权限策略
	KOResOrgPoliciesRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:/org/policys", "resource:orgPolicyReferences", "resource:orgAppActions",
			},
		},
	}).SetName("KOResOrgPoliciesRead").SetComments("资源权限管理权限策略只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgPoliciesRead)
	KOResOrgPoliciesEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:updateOrganizationPolicy", "resource:createOrganizationPolicy",
			},
		},
	}).SetName("KOResOrgPoliciesEdit").SetComments("资源权限管理权限策略修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgPoliciesEdit)
	KOResOrgPoliciesDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:deleteOrganizationPolicy",
			},
		},
	}).SetName("KOResOrgPoliciesDel").SetComments("资源权限管理权限策略删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgPoliciesDel)
	// 组织协作-用户组
	KOResOrgGroupsRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:/org/groups", "resource:orgGroups", "resource:orgRoleUsers",
			},
		},
	}).SetName("KOResOrgGroupsRead").SetComments("资源权限管理用户组只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgGroupsRead)
	KOResOrgGroupsEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:assignRoleUser", "resource:grant", "resource:createRole", "resource:updateRole", "resource:revokeRoleUser",
			},
		},
	}).SetName("KOResOrgGroupsEdit").SetComments("资源权限管理用户组修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgGroupsEdit)
	KOResOrgGroupsDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:deleteRole",
			},
		},
	}).SetName("KOResOrgGroupsDel").SetComments("资源权限管理用户组删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgGroupsDel)
	// 组织协作-角色
	KOResOrgRolesRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:/org/roles", "resource:orgRoles", "resource:orgRoleUsers",
			},
		},
	}).SetName("KOResOrgRolesRead").SetComments("资源权限管理角色只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgRolesRead)
	KOResOrgRolesEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:assignRoleUser", "resource:grant", "resource:createRole", "resource:updateRole", "resource:revokeRoleUser",
			},
		},
	}).SetName("KOResOrgRolesEdit").SetComments("资源权限管理角色修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgRolesEdit)
	KOResOrgRolesDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:deleteRole",
			},
		},
	}).SetName("KOResOrgRolesDel").SetComments("资源权限管理角色删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgRolesDel)
	// 系统设置-组织管理
	KOResSystemOrgRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:/system/org", "resource:organizations", "resource:orgGroups", "resource:orgRoleUsers", "resource:orgPolicyReferences", "resource:orgAppActions", "resource:apps",
				"resource:orgRecycleUsers", "resource:userGroups", "resource:userExtendGroupPolicies",
			},
		},
	}).SetName("KOResSystemOrgRead").SetComments("资源权限管理组织管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemOrgRead)
	KOResSystemOrgEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{"resource:createRoot", "resource:updateOrganization", "resource:assignRoleUser", "resource:grant", "resource:createRole", "resource:updateRole", "resource:revokeRoleUser",
				"resource:deleteRole", "resource:updateOrganizationPolicy", "resource:createOrganizationPolicy", "resource:deleteOrganizationPolicy", "resource:revokeOrganizationApp",
				"resource:assignOrganizationApp", "resource:deleteOrganization", "resource:createOrganization", "resource:createOrganizationUser", "resource:moveOrganization",
				"resource:recoverOrgUser", "resource:updateUser", "resource:updateLoginProfile", "resource:bindUserIdentity", "resource:deleteUserIdentity", "resource:sendMFAToUserByEmail",
				"resource:enableMFA", "resource:createOauthClient", "resource:enableOauthClient", "resource:disableOauthClient", "resource:deleteOauthClient",
				"resource:revoke"},
		},
	}).SetName("KOResSystemOrgEdit").SetComments("资源权限管理组织管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemOrgEdit)
	// 系统设置-账户管理
	KOResSystemAccountRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:/system/account", "resource:users", "resource:userGroups", "resource:orgGroups",
			},
		},
	}).SetName("KOResSystemAccountRead").SetComments("资源权限管理账户管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAccountRead)
	KOResSystemAccountEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:updateUser", "resource:updateLoginProfile", "resource:bindUserIdentity", "resource:deleteUserIdentity", "resource:sendMFAToUserByEmail", "resource:enableMFA",
				"resource:createOauthClient", "resource:enableOauthClient", "resource:disableOauthClient", "resource:deleteOauthClient", "resource:grant", "resource:revoke",
				"resource:createOrganizationAccount", "resource:resetUserPasswordByEmail",
			},
		},
	}).SetName("KOResSystemAccountEdit").SetComments("资源权限管理账户管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAccountEdit)
	KOResSystemAccountDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:deleteUser",
			},
		},
	}).SetName("KOResSystemAccountDel").SetComments("资源权限管理账户管理删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAccountDel)
	// 系统设置-应用管理
	KOResSystemAppRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:/system/app", "resource:apps", "resource:appPolicyAssignedToOrgs", "resource:appRoleAssignedToOrgs",
			},
		},
	}).SetName("KOResSystemAppRead").SetComments("资源权限管理应用管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAppRead)
	KOResSystemAppEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:createApp", "resource:updateApp", "resource:createAppActions", "resource:updateAppAction", "resource:deleteAppAction", "resource:createAppPolicy", "resource:updateAppPolicy",
				"resource:deleteAppPolicy", "resource:assignOrganizationAppPolicy", "resource:revokeOrganizationAppPolicy", "resource:createAppMenus", "resource:deleteAppMenu", "resource:updateAppMenu",
				"resource:moveAppMenu", "resource:createAppRole", "resource:updateAppRole", "resource:deleteAppRole", "resource:assignAppRolePolicy", "resource:revokeAppRolePolicy", "resource:revokeOrganizationAppRole",
				"resource:assignOrganizationAppRole", "resource:updateAppRes",
			},
		},
	}).SetName("KOResSystemAppEdit").SetComments("资源权限管理应用管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAppEdit)
	KOResSystemAppDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:deleteApp",
			},
		},
	}).SetName("KOResSystemAppDel").SetComments("资源权限管理应用管理删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAppDel)
	// 系统设置-文件来源
	KOResFileSourceRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:/system/file/source", "resource:fileSources",
			},
		},
	}).SetName("KOResFileSourceRead").SetComments("资源权限管理文件来源只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResFileSourceRead)
	KOResFileSourceEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:createFileSource", "resource:updateFileSource",
			},
		},
	}).SetName("KOResFileSourceEdit").SetComments("资源权限管理文件来源修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResFileSourceEdit)
	KOResFileSourceDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:deleteFileSource",
			},
		},
	}).SetName("KOResFileSourceDel").SetComments("资源权限管理文件来源删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResFileSourceDel)
	// 系统设置-数据字典
	KOResDictRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:/dict", "resource:appDicts",
			},
		},
	}).SetName("KOResDictRead").SetComments("资源权限管理文件来源只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDictRead)
	KOResDictEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:updateAppDict", "resource:createAppDict", "resource:updateAppDictItem", "resource:createAppDictItem", "resource:deleteAppDictItem", "resource:moveAppDictItem",
			},
		},
	}).SetName("KOResDictEdit").SetComments("资源权限管理文件来源修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDictEdit)
	KOResDictDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"resource:deleteAppDict",
			},
		},
	}).SetName("KOResDictDel").SetComments("资源权限管理文件来源删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDictDel)
	// 创建策略
	client.AppPolicy.CreateBulk(aps...).ExecX(context.Background())
}

func (*dataset) initRegion(client *ent.Tx) {
	// country
	cnty := make([]*ent.CountryCreate, 0)
	c1 := client.Country.Create().SetID(1).SetCreatedBy(1).SetName("中国").SetCode("86").SetNameEn("China")
	c2 := client.Country.Create().SetID(2).SetCreatedBy(1).SetName("中国香港").SetCode("852").SetNameEn("HongKong")
	cnty = append(cnty, c1, c2)
	client.Country.CreateBulk(cnty...).ExecX(context.Background())
	// region
	regs := make([]*ent.RegionCreate, 0)
	r1 := client.Region.Create().SetID(1).SetCreatedBy(1).SetName("北京").SetCountryID(1)
	r2 := client.Region.Create().SetID(2).SetCreatedBy(1).SetName("北京市").SetParentID(1).SetCountryID(1)
	r3 := client.Region.Create().SetID(3).SetCreatedBy(1).SetName("东城区").SetParentID(2).SetCountryID(1)
	r4 := client.Region.Create().SetID(4).SetCreatedBy(1).SetName("西城区").SetParentID(2).SetCountryID(1)
	regs = append(regs, r1, r2, r3, r4)
	r5 := client.Region.Create().SetID(5).SetCreatedBy(1).SetName("香港特别行政区").SetCountryID(2)
	r6 := client.Region.Create().SetID(6).SetCreatedBy(1).SetName("九龙").SetParentID(5).SetCountryID(2)
	r7 := client.Region.Create().SetID(7).SetCreatedBy(1).SetName("香港岛").SetParentID(5).SetCountryID(2)
	r8 := client.Region.Create().SetID(8).SetCreatedBy(1).SetName("新界").SetParentID(5).SetCountryID(2)
	regs = append(regs, r5, r6, r7, r8)
	client.Region.CreateBulk(regs...).ExecX(context.Background())
}

func (*dataset) initAppDict(client *ent.Tx) {
	ctx := identity.WithTenantID(context.Background(), 1)
	// appdict
	dicts := make([]*ent.AppDictCreate, 0)
	d1 := client.AppDict.Create().SetID(1).SetCreatedBy(1).SetAppID(1).SetName("地理时区").SetCode("DLSH")
	dicts = append(dicts, d1)
	client.AppDict.CreateBulk(dicts...).ExecX(ctx)
	// appdictitems
	items := make([]*ent.AppDictItemCreate, 0)
	i1 := client.AppDictItem.Create().SetID(1).SetCreatedBy(1).SetRefCode("resource:DLSH").SetOrgID(1).SetDictID(1).SetName("Asia/Hong_Kong").SetCode("Asia/Hong_Kong").SetStatus(typex.SimpleStatusActive)
	i2 := client.AppDictItem.Create().SetID(2).SetCreatedBy(1).SetRefCode("resource:DLSH").SetDictID(1).SetName("Asia/Shanghai").SetCode("Asia/Shanghai").SetStatus(typex.SimpleStatusActive)
	i3 := client.AppDictItem.Create().SetID(3).SetCreatedBy(1).SetRefCode("resource:DLSH").SetOrgID(1).SetDictID(1).SetName("上海时区").SetCode("Asia/Shanghai").SetStatus(typex.SimpleStatusActive)
	i4 := client.AppDictItem.Create().SetID(4).SetCreatedBy(1).SetRefCode("resource:DLSH").SetOrgID(2).SetDictID(1).SetName("上海时区").SetCode("Asia/Shanghai").SetStatus(typex.SimpleStatusActive)
	items = append(items, i1, i2, i3, i4)
	client.AppDictItem.CreateBulk(items...).ExecX(ctx)
}
