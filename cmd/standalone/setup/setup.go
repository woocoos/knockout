package setup

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tsingsun/woocoo/pkg/security"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/identity"
	schemahook "github.com/woocoos/knockout/codegen/entgen/hook"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/migrate"
	"github.com/woocoos/knockout/script/data"
)

// Config 系统初始化配置
type Config struct {
	// 数据库驱动名
	DriverName string
	// 数据库连接字符串
	DSN string
	// Redis 地址
	RedisAddr string
	// 管理员登录名
	AdminPrincipalName string
	// 管理员显示名
	AdminDisplayName string
	// 管理员密码(明文)
	AdminPassword string
	// 管理员邮箱
	AdminEmail string
	// 根组织域名
	OrgDomain string
}

// IsInitialized 检查系统是否已完成初始化.
func IsInitialized(name, dsn string) bool {
	drv, err := sql.Open(name, dsn)
	if err != nil {
		return false
	}
	client := ent.NewClient(ent.Driver(drv))
	defer client.Close()
	count, err := client.App.Query().Count(context.Background())
	if err != nil {
		return false
	}
	return count > 0
}

// Do 执行系统初始化: 创建schema -> 初始化基础数据 -> 更新管理员 -> 初始化资源策略.
func Do(cfg Config) error {
	drv, err := sql.Open(cfg.DriverName, cfg.DSN)
	if err != nil {
		return err
	}
	defer drv.Close()

	// 创建 portal schema
	client := ent.NewClient(ent.Driver(drv))
	schemahook.RegisterAllHooks(client)
	err = client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false))
	if err != nil {
		return err
	}

	// 创建 casbin schema
	casbinClient := casbinent.NewClient(casbinent.Driver(drv))
	err = casbinClient.Schema.Create(context.Background(),
		schema.WithForeignKeys(false))
	if err != nil {
		return err
	}

	// 初始化基础数据 (使用已有的 client, 共享数据库连接)
	data.InitBaseWithClients(client, casbinClient)

	// 更新管理员信息
	if err := updateAdmin(drv, cfg); err != nil {
		return err
	}

	// 初始化资源策略
	tx, err := client.Tx(adminCtx(client))
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()
	data.InitResourcePolicy(tx)
	return tx.Commit()
}

// updateAdmin 更新默认管理员的账号信息
func updateAdmin(drv *sql.Driver, cfg Config) error {
	client := ent.NewClient(ent.Driver(drv))
	schemahook.RegisterAllHooks(client)
	// 不关闭 client, 因为 drv 是共享的

	ctx := adminCtx(client)

	if cfg.AdminPrincipalName != "" {
		// 检查是否需要更新(避免 hook 唯一性校验冲突)
		currentUser := client.User.GetX(ctx, 1)
		if currentUser.PrincipalName != cfg.AdminPrincipalName {
			if err := client.User.UpdateOneID(1).
				SetPrincipalName(cfg.AdminPrincipalName).
				SetDisplayName(cfg.AdminDisplayName).
				Exec(ctx); err != nil {
				return err
			}
			currentIdentity := client.UserIdentity.GetX(ctx, 1)
			if currentIdentity.Code != cfg.AdminPrincipalName {
				if err := client.UserIdentity.UpdateOneID(1).
					SetCode(cfg.AdminPrincipalName).
					Exec(ctx); err != nil {
					return err
				}
			}
		}
	}

	if cfg.AdminEmail != "" {
		currentAddr := client.UserAddr.GetX(ctx, 1)
		if currentAddr.Email != cfg.AdminEmail {
			if err := client.UserAddr.UpdateOneID(1).
				SetEmail(cfg.AdminEmail).
				Exec(ctx); err != nil {
				return err
			}
		}
	}

	if cfg.AdminPassword != "" {
		salt := randomStr(5)
		hashPwd := sha256Hash(cfg.AdminPassword + salt)
		if err := client.UserPassword.UpdateOneID(1).
			SetPassword(hashPwd).
			SetSalt(salt).
			SetStatus(typex.SimpleStatusActive).
			Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}

// adminCtx 创建管理员安全上下文
func adminCtx(client *ent.Client) context.Context {
	ctx := ent.NewContext(context.Background(), client)
	ctx = security.WithContext(ctx, security.NewGenericPrincipalByClaims(jwt.MapClaims{"sub": "1"}))
	ctx = identity.WithTenantID(ctx, 1)
	return ctx
}

func sha256Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func randomStr(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	result := make([]byte, n)
	for i := range result {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[idx.Int64()]
	}
	return string(result)
}

// DefaultConfig 返回默认的管理员配置
func DefaultConfig() Config {
	return Config{
		DriverName:         "mysql",
		DSN:                "root:@tcp(localhost:3306)/portal?parseTime=true&loc=Local",
		RedisAddr:          "127.0.0.1:6379",
		AdminPrincipalName: "admin",
		AdminDisplayName:   "admin",
		AdminPassword:      "123456",
		AdminEmail:         "admin@localhost",
		OrgDomain:          "woocoo.com",
	}
}
