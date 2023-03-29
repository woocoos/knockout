// go:build ignore
package main

import (
	"context"
	"flag"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/org"
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
	initUser(tx)
	initOrg(tx)
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
			SetStatus(typex.SimpleStatusActive).SetCreatedBy(1)
		if i == 1 {
			c.SetPrincipalName("admin").SetDisplayName("admin")
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
		ui = append(ui, id)
	}
	err := client.User.CreateBulk(ub...).Exec(context.Background())
	if err != nil {
		panic(err)
	}
	err = client.UserLoginProfile.CreateBulk(ulp...).Exec(context.Background())
	if err != nil {
		panic(err)
	}
	err = client.UserPassword.CreateBulk(up...).Exec(context.Background())
	if err != nil {
		panic(err)
	}
	err = client.UserIdentity.CreateBulk(ui...).Exec(context.Background())
	if err != nil {
		panic(err)
	}
}
