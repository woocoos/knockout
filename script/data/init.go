// Package data is the data tools. It is used to initialize the database data or test data.
package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/filesource"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/user"
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
	ds := dataset{
		casbinRules: make([]*casbinent.CasbinRuleCreate, 0),
	}
	drv, err := sql.Open(name, dsn)
	if err != nil {
		panic(err)
	}
	ds.portal = ent.NewClient(ent.Driver(drv))
	ds.casbin = casbinent.NewClient(casbinent.Driver(drv))

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
}

func (*dataset) initOrg(client *ent.Tx) {
	ou := make([]*ent.OrgUserCreate, 0)
	for i := 1; i < 4; i++ {
		// 由于path字段是计算字段，所以这里不需要设置,但需要对org独立保存.
		c := client.Org.Create().SetID(i).SetKind(org.KindOrganization).SetParentID(i - 1).SetStatus(typex.SimpleStatusActive).
			SetCreatedBy(1).SetName("org" + strconv.Itoa(i))
		if i == 1 {
			c.SetKind(org.KindRoot).SetDomain("woocoo.com").SetOwnerID(1)
		}
		if err := c.Exec(context.Background()); err != nil {
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
	ulp := make([]*ent.UserLoginProfileCreate, 0)
	up := make([]*ent.UserPasswordCreate, 0)
	ui := make([]*ent.UserIdentityCreate, 0)
	for i := 1; i < 4; i++ {
		c := client.User.Create().SetID(i).SetUserType(user.UserTypeAccount).SetCreationType(user.CreationTypeManual).
			SetRegisterIP("").SetPrincipalName("user" + strconv.Itoa(i)).SetDisplayName("user" + strconv.Itoa(i)).
			SetStatus(typex.SimpleStatusActive).SetCreatedBy(1).SetEmail("user" + strconv.Itoa(i) + "@localhost")
		if i == 1 {
			c.SetPrincipalName("admin").SetDisplayName("admin").SetEmail("admin@localhost")
		}
		ub = append(ub, c)

		lp := client.UserLoginProfile.Create().SetID(i).SetUserID(i).SetCreatedBy(1).SetSetKind(userloginprofile.SetKindKeep).
			SetCanLogin(true).SetPasswordReset(false).SetMfaSecret("UWZLIIUMPX53NYXB").SetVerifyDevice(false)
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
			SetComments("全部管理权限").SetAutoGrant(true).SetStatus(typex.SimpleStatusActive).SetVersion("V1").
			SetRules([]*types.PolicyRule{
				{
					Effect:    "allow",
					Actions:   []string{"*"},
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

		ps = append(ps, client.Permission.Create().SetID(i).SetOrgID(1).SetOrgPolicyID(i).SetCreatedBy(1).
			SetPrincipalKind(permission.PrincipalKindRole).SetRoleID(i).SetStatus(typex.SimpleStatusActive))

		orus = append(orus, client.OrgRoleUser.Create().SetID(i).SetCreatedBy(1).SetOrgRoleID(i).SetOrgUserID(1).SetUserID(1).SetOrgID(1))

		set.casbinRules = append(set.casbinRules, casbinClient.CasbinRule.Create().SetPtype("g").
			SetV0("1").SetV1(strconv.Itoa(i)).SetV2("1"))
		set.casbinRules = append(set.casbinRules, casbinClient.CasbinRule.Create().SetPtype("p").
			SetV0(strconv.Itoa(i)).SetV1("1").SetV2(ac+":*").SetV3("read").SetV4("allow"))
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
	fs := make([]*ent.FileSourceCreate, 0)
	s1 := client.FileSource.Create().SetID(1).SetKind(filesource.KindLocal).SetComments("本地存储bucket").
		SetEndpoint("http://127.0.0.1:10071").SetBucket("local").SetCreatedBy(1)
	fs = append(fs, s1)
	client.FileSource.CreateBulk(fs...).ExecX(context.Background())
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
				"userPermissions", "userMenus", "userRootOrgs", "node", "login", "appAccess", "/", "orgRoles", "orgGroups", "/user/info",
				"updateUser", "/user/safety", "changePassword", "userApps", "orgUserPreference",
			},
		},
	}).SetName("KOResAccess").SetComments("资源权限管理应用授权，拥有该策略允许登录后台").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, koResAccess)
	// 组织协作-部门管理
	KOResDepartmentRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/org/departments", "organizations",
			},
		},
	}).SetName("KOResDepartmentRead").SetComments("资源权限管理部门管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDepartmentRead)
	KOResDepartmentEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"createOrganization", "updateOrganization",
			},
		},
	}).SetName("KOResDepartmentEdit").SetComments("资源权限管理部门管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDepartmentEdit)
	KOResDepartmentDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteOrganization",
			},
		},
	}).SetName("KOResDepartmentDel").SetComments("资源权限管理部门管理删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDepartmentDel)
	// 组织协作-用户管理
	KOResOrgUsersRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/org/users", "organizations", "orgRecycleUsers", "userGroups", "orgGroups", "userExtendGroupPolicies",
			},
		},
	}).SetName("KOResOrgUsersRead").SetComments("资源权限管理用户管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgUsersRead)
	KOResOrgUsersEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"createOrganizationUser", "moveOrganization", "recoverOrgUser", "updateUser", "updateLoginProfile", "bindUserIdentity",
				"deleteUserIdentity", "sendMFAToUserByEmail", "enableMFA", "createOauthClient", "enableOauthClient", "disableOauthClient",
				"deleteOauthClient", "grant", "revoke", "removeOrganizationUser",
			},
		},
	}).SetName("KOResOrgUsersEdit").SetComments("资源权限管理用户管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgUsersEdit)
	// 组织协作-权限策略
	KOResOrgPoliciesRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/org/policys", "orgPolicyReferences", "orgAppActions",
			},
		},
	}).SetName("KOResOrgPoliciesRead").SetComments("资源权限管理权限策略只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgPoliciesRead)
	KOResOrgPoliciesEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"updateOrganizationPolicy", "createOrganizationPolicy",
			},
		},
	}).SetName("KOResOrgPoliciesEdit").SetComments("资源权限管理权限策略修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgPoliciesEdit)
	KOResOrgPoliciesDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteOrganizationPolicy",
			},
		},
	}).SetName("KOResOrgPoliciesDel").SetComments("资源权限管理权限策略删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgPoliciesDel)
	// 组织协作-用户组
	KOResOrgGroupsRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/org/groups", "orgGroups", "orgRoleUsers",
			},
		},
	}).SetName("KOResOrgGroupsRead").SetComments("资源权限管理用户组只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgGroupsRead)
	KOResOrgGroupsEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"assignRoleUser", "grant", "createRole", "updateRole", "revokeRoleUser",
			},
		},
	}).SetName("KOResOrgGroupsEdit").SetComments("资源权限管理用户组修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgGroupsEdit)
	KOResOrgGroupsDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteRole",
			},
		},
	}).SetName("KOResOrgGroupsDel").SetComments("资源权限管理用户组删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgGroupsDel)
	// 组织协作-角色
	KOResOrgRolesRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/org/roles", "orgRoles", "orgRoleUsers",
			},
		},
	}).SetName("KOResOrgRolesRead").SetComments("资源权限管理角色只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgRolesRead)
	KOResOrgRolesEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"assignRoleUser", "grant", "createRole", "updateRole", "revokeRoleUser",
			},
		},
	}).SetName("KOResOrgRolesEdit").SetComments("资源权限管理角色修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgRolesEdit)
	KOResOrgRolesDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteRole",
			},
		},
	}).SetName("KOResOrgRolesDel").SetComments("资源权限管理角色删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResOrgRolesDel)
	// 系统设置-组织管理
	KOResSystemOrgRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/system/org", "organizations", "orgGroups", "orgRoleUsers", "orgPolicyReferences", "orgAppActions", "apps",
				"orgRecycleUsers", "userGroups", "userExtendGroupPolicies",
			},
		},
	}).SetName("KOResSystemOrgRead").SetComments("资源权限管理组织管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemOrgRead)
	KOResSystemOrgEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{"createRoot", "updateOrganization", "assignRoleUser", "grant", "createRole", "updateRole", "revokeRoleUser",
				"deleteRole", "updateOrganizationPolicy", "createOrganizationPolicy", "deleteOrganizationPolicy", "revokeOrganizationApp",
				"assignOrganizationApp", "deleteOrganization", "createOrganization", "createOrganizationUser", "moveOrganization",
				"recoverOrgUser", "updateUser", "updateLoginProfile", "bindUserIdentity", "deleteUserIdentity", "sendMFAToUserByEmail",
				"enableMFA", "createOauthClient", "enableOauthClient", "disableOauthClient", "deleteOauthClient",
				"revoke"},
		},
	}).SetName("KOResSystemOrgEdit").SetComments("资源权限管理组织管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemOrgEdit)
	// 系统设置-账户管理
	KOResSystemAccountRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/system/account", "users", "userGroups", "orgGroups",
			},
		},
	}).SetName("KOResSystemAccountRead").SetComments("资源权限管理账户管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAccountRead)
	KOResSystemAccountEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"updateUser", "updateLoginProfile", "bindUserIdentity", "deleteUserIdentity", "sendMFAToUserByEmail", "enableMFA",
				"createOauthClient", "enableOauthClient", "disableOauthClient", "deleteOauthClient", "grant", "revoke",
				"createOrganizationAccount", "resetUserPasswordByEmail",
			},
		},
	}).SetName("KOResSystemAccountEdit").SetComments("资源权限管理账户管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAccountEdit)
	KOResSystemAccountDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteUser",
			},
		},
	}).SetName("KOResSystemAccountDel").SetComments("资源权限管理账户管理删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAccountDel)
	// 系统设置-应用管理
	KOResSystemAppRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/system/app", "apps", "appPolicyAssignedToOrgs", "appRoleAssignedToOrgs",
			},
		},
	}).SetName("KOResSystemAppRead").SetComments("资源权限管理应用管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAppRead)
	KOResSystemAppEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"createApp", "updateApp", "createAppActions", "updateAppAction", "deleteAppAction", "createAppPolicy", "updateAppPolicy",
				"deleteAppPolicy", "assignOrganizationAppPolicy", "revokeOrganizationAppPolicy", "createAppMenus", "deleteAppMenu", "updateAppMenu",
				"moveAppMenu", "createAppRole", "updateAppRole", "deleteAppRole", "assignAppRolePolicy", "revokeAppRolePolicy", "revokeOrganizationAppRole",
				"assignOrganizationAppRole", "updateAppRes",
			},
		},
	}).SetName("KOResSystemAppEdit").SetComments("资源权限管理应用管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAppEdit)
	KOResSystemAppDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteApp",
			},
		},
	}).SetName("KOResSystemAppDel").SetComments("资源权限管理应用管理删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResSystemAppDel)
	// 系统设置-文件来源
	KOResFileSourceRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/system/file/source", "fileSources",
			},
		},
	}).SetName("KOResFileSourceRead").SetComments("资源权限管理文件来源只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResFileSourceRead)
	KOResFileSourceEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"createFileSource", "updateFileSource",
			},
		},
	}).SetName("KOResFileSourceEdit").SetComments("资源权限管理文件来源修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResFileSourceEdit)
	KOResFileSourceDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteFileSource",
			},
		},
	}).SetName("KOResFileSourceDel").SetComments("资源权限管理文件来源删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResFileSourceDel)
	// 系统设置-数据字典
	KOResDictRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/dict", "appDicts",
			},
		},
	}).SetName("KOResDictRead").SetComments("资源权限管理文件来源只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDictRead)
	KOResDictEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"updateAppDict", "createAppDict", "updateAppDictItem", "createAppDictItem", "deleteAppDictItem", "moveAppDictItem",
			},
		},
	}).SetName("KOResDictEdit").SetComments("资源权限管理文件来源修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDictEdit)
	KOResDictDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteAppDict",
			},
		},
	}).SetName("KOResDictDel").SetComments("资源权限管理文件来源删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(false)
	aps = append(aps, KOResDictDel)
	// 创建策略
	client.AppPolicy.CreateBulk(aps...).ExecX(context.Background())
}
