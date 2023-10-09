//go:build ignore

package main

import (
	"context"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	_ "github.com/woocoos/knockout/ent/runtime"
	"log"
)

// receive two arguments: the migration name and the database dsn.
var (
	dsn  = flag.String("dsn", "deo:deo135@$^@tcp(192.168.0.18:3306)/portal?parseTime=true&loc=Local", "")
	name = flag.String("name", "mysql", "driver name")
)

func main() {
	flag.Parse()
	client, err := ent.Open(*name, *dsn, ent.Debug())
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	tx, err := client.Tx(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()
	initResourcePolicy(tx)
}

func initResourcePolicy(client *ent.Tx) {
	createBy := 0
	version := "1"
	ap, err := client.App.Query().Where(app.Code("resource")).Only(context.Background())
	if err != nil {
		panic(err)
	}
	aps := make([]*ent.AppPolicyCreate, 0)
	// 基础策略
	KOResBase := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"userPermissions", "userMenus", "userRootOrgs", "node",
			},
		},
	}).SetName("KOResBase").SetComments("资源权限管理基础权限，后台用户必要策略").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResBase)
	// 工作台
	KOResConsole := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/", "orgRoles", "orgGroups",
			},
		},
	}).SetName("KOResConsole").SetComments("资源权限管理工作台只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResConsole)
	// 个人中心-基本信息
	KOResUserInfoRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/user/info",
			},
		},
	}).SetName("KOResUserInfoRead").SetComments("资源权限管理基本信息只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResUserInfoRead)
	KOResUserInfoEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"updateUser",
			},
		},
	}).SetName("KOResUserInfoEdit").SetComments("资源权限管理基本信息修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResUserInfoEdit)
	// 个人中心-安全设置
	KOResUserSafetyRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/user/safety",
			},
		},
	}).SetName("KOResUserSafetyRead").SetComments("资源权限管理安全设置只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResUserSafetyRead)
	KOResUserSafetyEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			//Actions: []string{
			//	"changePassword", "/mfa/bind-prepare", "/mfa/bind", "/mfa/unbind",
			//},
			Actions: []string{
				"changePassword",
			},
		},
	}).SetName("KOResUserSafetyEdit").SetComments("资源权限管理安全设置修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResUserSafetyEdit)
	// 组织协作-部门管理
	KOResDepartmentRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/org/departments", "organizations",
			},
		},
	}).SetName("KOResDepartmentRead").SetComments("资源权限管理部门管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResDepartmentRead)
	KOResDepartmentEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"createOrganization", "updateOrganization",
			},
		},
	}).SetName("KOResDepartmentEdit").SetComments("资源权限管理部门管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResDepartmentEdit)
	KOResDepartmentDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteOrganization",
			},
		},
	}).SetName("KOResDepartmentDel").SetComments("资源权限管理部门管理删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResDepartmentDel)
	// 组织协作-用户管理
	KOResOrgUsersRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/org/users", "organizations", "orgRecycleUsers", "userGroups", "orgGroups", "userExtendGroupPolicies",
			},
		},
	}).SetName("KOResOrgUsersRead").SetComments("资源权限管理用户管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
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
	}).SetName("KOResOrgUsersEdit").SetComments("资源权限管理用户管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResOrgUsersEdit)
	// 组织协作-权限策略
	KOResOrgPoliciesRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/org/policys", "orgPolicyReferences", "orgAppActions",
			},
		},
	}).SetName("KOResOrgPoliciesRead").SetComments("资源权限管理权限策略只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResOrgPoliciesRead)
	KOResOrgPoliciesEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"updateOrganizationPolicy", "createOrganizationPolicy",
			},
		},
	}).SetName("KOResOrgPoliciesEdit").SetComments("资源权限管理权限策略修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResOrgPoliciesEdit)
	KOResOrgPoliciesDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteOrganizationPolicy",
			},
		},
	}).SetName("KOResOrgPoliciesDel").SetComments("资源权限管理权限策略删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResOrgPoliciesDel)
	// 组织协作-用户组
	KOResOrgGroupsRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/org/groups", "orgGroups", "orgRoleUsers",
			},
		},
	}).SetName("KOResOrgGroupsRead").SetComments("资源权限管理用户组只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResOrgGroupsRead)
	KOResOrgGroupsEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"assignRoleUser", "grant", "createRole", "updateRole", "revokeRoleUser",
			},
		},
	}).SetName("KOResOrgGroupsEdit").SetComments("资源权限管理用户组修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResOrgGroupsEdit)
	KOResOrgGroupsDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteRole",
			},
		},
	}).SetName("KOResOrgGroupsDel").SetComments("资源权限管理用户组删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResOrgGroupsDel)
	// 组织协作-角色
	KOResOrgRolesRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/org/roles", "orgRoles", "orgRoleUsers",
			},
		},
	}).SetName("KOResOrgRolesRead").SetComments("资源权限管理角色只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResOrgRolesRead)
	KOResOrgRolesEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"assignRoleUser", "grant", "createRole", "updateRole", "revokeRoleUser",
			},
		},
	}).SetName("KOResOrgRolesEdit").SetComments("资源权限管理角色修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResOrgRolesEdit)
	KOResOrgRolesDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteRole",
			},
		},
	}).SetName("KOResOrgRolesDel").SetComments("资源权限管理角色删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
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
	}).SetName("KOResSystemOrgRead").SetComments("资源权限管理组织管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
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
	}).SetName("KOResSystemOrgEdit").SetComments("资源权限管理组织管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResSystemOrgEdit)
	// 系统设置-账户管理
	KOResSystemAccountRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/system/account", "users", "userGroups", "orgGroups",
			},
		},
	}).SetName("KOResSystemAccountRead").SetComments("资源权限管理账户管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
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
	}).SetName("KOResSystemAccountEdit").SetComments("资源权限管理账户管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResSystemAccountEdit)
	KOResSystemAccountDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteUser",
			},
		},
	}).SetName("KOResSystemAccountDel").SetComments("资源权限管理账户管理删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResSystemAccountDel)
	// 系统设置-应用管理
	KOResSystemAppRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/system/app", "apps", "appPolicyAssignedToOrgs", "appRoleAssignedToOrgs",
			},
		},
	}).SetName("KOResSystemAppRead").SetComments("资源权限管理应用管理只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
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
	}).SetName("KOResSystemAppEdit").SetComments("资源权限管理应用管理修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResSystemAppEdit)
	KOResSystemAppDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteApp",
			},
		},
	}).SetName("KOResSystemAppDel").SetComments("资源权限管理应用管理删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResSystemAppDel)
	// 系统设置-文件来源
	KOResFileSourceRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/system/file/source", "fileSources",
			},
		},
	}).SetName("KOResFileSourceRead").SetComments("资源权限管理文件来源只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResFileSourceRead)
	KOResFileSourceEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"createFileSource", "updateFileSource",
			},
		},
	}).SetName("KOResFileSourceEdit").SetComments("资源权限管理文件来源修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResFileSourceEdit)
	KOResFileSourceDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteFileSource",
			},
		},
	}).SetName("KOResFileSourceDel").SetComments("资源权限管理文件来源删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResFileSourceDel)
	// 系统设置-数据字典
	KOResDictRead := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"/dict", "appDicts",
			},
		},
	}).SetName("KOResDictRead").SetComments("资源权限管理文件来源只读").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResDictRead)
	KOResDictEdit := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"updateAppDict", "createAppDict", "updateAppDictItem", "createAppDictItem", "deleteAppDictItem", "moveAppDictItem",
			},
		},
	}).SetName("KOResDictEdit").SetComments("资源权限管理文件来源修改").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResDictEdit)
	KOResDictDel := client.AppPolicy.Create().SetCreatedBy(createBy).SetVersion(version).SetRules([]*types.PolicyRule{
		{
			Effect: types.PolicyEffectAllow,
			Actions: []string{
				"deleteAppDict",
			},
		},
	}).SetName("KOResDictDel").SetComments("资源权限管理文件来源删除").SetAppID(ap.ID).SetStatus(typex.SimpleStatusActive).SetAutoGrant(true)
	aps = append(aps, KOResDictDel)
	// 创建策略
	err = client.AppPolicy.CreateBulk(aps...).Exec(context.Background())
	if err != nil {
		panic(err)
	}
}
