package setup

import (
	"context"
	"os"
	"testing"

	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	schemahook "github.com/woocoos/knockout/codegen/entgen/hook"
	"github.com/woocoos/knockout/ent"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/woocoos/knockout/ent/runtime"
)

func TestSetup_WithSqlite(t *testing.T) {
	// 使用文件型 SQLite, 避免内存数据库连接关闭后数据丢失
	dbFile := "setup_test.db"
	defer os.Remove(dbFile)
	dsn := dbFile + "?_fk=1"

	assert.False(t, IsInitialized("sqlite3", dsn))

	cfg := Config{
		DriverName:         "sqlite3",
		DSN:                dsn,
		AdminPrincipalName: "myadmin",
		AdminDisplayName:   "My Admin",
		AdminPassword:      "securepass123",
		AdminEmail:         "myadmin@example.com",
	}
	require.NotPanics(t, func() {
		err := Do(cfg)
		require.NoError(t, err)
	})

	assert.True(t, IsInitialized("sqlite3", dsn))

	// 验证自定义管理员信息
	drv, err := sql.Open("sqlite3", dsn)
	require.NoError(t, err)
	client := ent.NewClient(ent.Driver(drv))
	defer client.Close()
	schemahook.RegisterAllHooks(client)

	ctx := context.Background()

	admin := client.User.GetX(ctx, 1)
	assert.Equal(t, "myadmin", admin.PrincipalName)
	assert.Equal(t, "My Admin", admin.DisplayName)

	adminAddr := client.UserAddr.GetX(ctx, 1)
	assert.Equal(t, "myadmin@example.com", adminAddr.Email)

	adminId := client.UserIdentity.GetX(ctx, 1)
	assert.Equal(t, "myadmin", adminId.Code)

	adminPwd := client.UserPassword.GetX(ctx, 1)
	assert.NotEmpty(t, adminPwd.Salt)
	assert.NotEmpty(t, adminPwd.Password)
	assert.Len(t, adminPwd.Password, 64)

	assert.Greater(t, client.App.Query().CountX(ctx), 0)
	assert.Greater(t, client.Org.Query().CountX(ctx), 0)
	assert.Greater(t, client.AppPolicy.Query().CountX(ctx), 1, "应包含InitResourcePolicy创建的策略")

	// 验证 casbin_rules 表已创建
	casbinClient := casbinent.NewClient(casbinent.Driver(drv))
	count, err := casbinClient.CasbinRule.Query().Count(ctx)
	require.NoError(t, err, "casbin_rules 表应存在且可查询")
	assert.Greater(t, count, 0, "casbin_rules 表应包含 InitResourcePolicy 创建的策略")
}

func TestSetup_WithSqliteFileURI(t *testing.T) {
	// 测试使用 file: URI 格式的 DSN (setup wizard 使用的格式)
	dbFile := "setup_file_test.db"
	defer os.Remove(dbFile)
	dsn := "file:" + dbFile + "?cache=shared&_fk=1"

	assert.False(t, IsInitialized("sqlite3", dsn))

	cfg := Config{
		DriverName:         "sqlite3",
		DSN:                dsn,
		AdminPrincipalName: "admin",
		AdminDisplayName:   "Admin",
		AdminPassword:      "123456",
		AdminEmail:         "admin@localhost",
	}
	require.NotPanics(t, func() {
		err := Do(cfg)
		require.NoError(t, err)
	})

	assert.True(t, IsInitialized("sqlite3", dsn))

	// 验证 casbin_rules 表已创建
	drv, err := sql.Open("sqlite3", dsn)
	require.NoError(t, err)
	casbinClient := casbinent.NewClient(casbinent.Driver(drv))
	defer casbinClient.Close()
	count, err := casbinClient.CasbinRule.Query().Count(context.Background())
	require.NoError(t, err, "casbin_rules 表应存在且可查询")
	assert.Greater(t, count, 0, "casbin_rules 表应包含 InitResourcePolicy 创建的策略")
}

func TestSetup_WithSqliteCacheShared(t *testing.T) {
	// 测试 cache=shared 模式是否影响表创建
	dbFile := "setup_shared_test.db"
	defer os.Remove(dbFile)
	dsn := dbFile + "?cache=shared&_fk=1"

	cfg := Config{
		DriverName:         "sqlite3",
		DSN:                dsn,
		AdminPrincipalName: "admin",
		AdminDisplayName:   "Admin",
		AdminPassword:      "123456",
	}
	require.NoError(t, Do(cfg))

	// 用新连接验证表是否存在
	drv, err := sql.Open("sqlite3", dsn)
	require.NoError(t, err)
	defer drv.Close()

	// 检查 casbin_rules 表
	casbinClient := casbinent.NewClient(casbinent.Driver(drv))
	count, err := casbinClient.CasbinRule.Query().Count(context.Background())
	require.NoError(t, err, "casbin_rules 表应存在")
	assert.Greater(t, count, 0, "应包含策略数据")
}

func TestSetup_DefaultConfig(t *testing.T) {
	dbFile := "setup_default_test.db"
	defer os.Remove(dbFile)
	dsn := dbFile + "?_fk=1"

	cfg := DefaultConfig()
	cfg.DriverName = "sqlite3"
	cfg.DSN = dsn

	require.NotPanics(t, func() {
		err := Do(cfg)
		require.NoError(t, err)
	})

	assert.True(t, IsInitialized("sqlite3", dsn))
}
