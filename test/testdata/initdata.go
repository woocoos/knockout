//go:build ignore

package main

import (
	"context"
	"flag"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/entco/schemax/typex"
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
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/knockout/ent/runtime"
)

// receive two arguments: the migration name and the database dsn.
var (
	dsn  = flag.String("dsn", "root:@tcp(localhost:3306)/portal?parseTime=true&loc=Local", "")
	name = flag.String("name", "mysql", "driver name")
)

func main() {
	flag.Parse()
	client, err := ent.Open(*name, *dsn, ent.Debug())
	casbinClient, err := casbinent.Open(*name, *dsn, casbinent.Debug())
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	tx, err := client.Tx(ctx)
	if err != nil {
		log.Fatal(err)
	}
	casbinTx, err := casbinClient.Tx(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			casbinTx.Rollback()
			panic(err)
		} else {
			tx.Commit()
			casbinTx.Commit()
		}
	}()
	initUser(tx)
	initOrg(tx)
	initApp(tx, casbinTx)
	initFileSource(tx)
	initOauthClient(tx)
}

func initOrg(client *ent.Tx) {
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

func initUser(client *ent.Tx) {
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

func initApp(client *ent.Tx, casbinClient *casbinent.Tx) {
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
	casbinRules := make([]*casbinent.CasbinRuleCreate, 0)
	for i := 1; i < 2; i++ {
		ac := "resource"
		a := client.App.Create().SetID(i).SetName("资源权限管理").SetCode(ac).SetKind(app.KindWeb).
			SetComments("资源权限管理是管理组织目录中的应用,组织,人员以及授权信息").SetStatus(typex.SimpleStatusActive).
			SetCreatedBy(1).SetOwnerOrgID(1)
		apps = append(apps, a)

		ras = append(ras, client.AppAction.Create().SetID(i).SetAppID(i).SetCreatedBy(1).
			SetName("entry").SetKind(appaction.KindGraphql).SetComments("应用入口").SetMethod(appaction.MethodList),
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

		casbinRules = append(casbinRules, casbinClient.CasbinRule.Create().SetPtype("g").
			SetV0("1").SetV1(strconv.Itoa(i)).SetV2("1"))
		casbinRules = append(casbinRules, casbinClient.CasbinRule.Create().SetPtype("p").
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
	casbinClient.CasbinRule.CreateBulk(casbinRules...).ExecX(context.Background())
}

func initFileSource(client *ent.Tx) {
	fs := make([]*ent.FileSourceCreate, 0)
	s1 := client.FileSource.Create().SetID(1).SetKind(filesource.KindLocal).SetComments("本地存储bucket").
		SetEndpoint("http://127.0.0.1:10071").SetBucket("local").SetCreatedBy(1)
	fs = append(fs, s1)
	client.FileSource.CreateBulk(fs...).ExecX(context.Background())
}

func initOauthClient(client *ent.Tx) {
	oc := make([]*ent.OauthClientCreate, 0)
	s1 := client.OauthClient.Create().SetID(1).SetName("系统").SetClientID("206734260394752").SetClientSecret("T2UlqISVFq4DR9InXamj3l74iWdu3Tyr").
		SetGrantTypes("client_credentials").SetStatus(typex.SimpleStatusActive).SetUserID(1).SetCreatedBy(1)
	oc = append(oc, s1)
	client.OauthClient.CreateBulk(oc...).ExecX(context.Background())
}
