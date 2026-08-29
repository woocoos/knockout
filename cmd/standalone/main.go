package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/security"
	entadapter "github.com/woocoos/casbin-ent-adapter"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/ent/clientx"
	authzcasbin "github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	"github.com/woocoos/knockout-go/pkg/koapp"
	"github.com/woocoos/knockout/api/graphql"
	"github.com/woocoos/knockout/api/oas/auth"
	"github.com/woocoos/knockout/cmd/auth/casbin"
	setupapi "github.com/woocoos/knockout/cmd/standalone/setup"
	schemahook "github.com/woocoos/knockout/codegen/entgen/hook"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/job"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/woocoos/knockout-go/pkg/snowflake"
	_ "github.com/woocoos/knockout/ent/runtime"
)

// standalone 整合 adminx 与 auth 两个服务,用于开发环境一站式启动.
// 配置文件独立管理在 etc/ 目录下, woocoo 默认读取运行目录的 etc/.
//
// 启动流程:
// 1. 尝试加载配置, 检测数据库连接是否有效
// 2. 如果无效 → 启动内置 HTTP 安装向导 (不依赖配置文件)
// 3. 安装完成后自动重启, 使用新配置启动完整服务
func main() {
	// 尝试加载配置
	app := koapp.New(woocoo.WithAppConfiguration(
		conf.New(conf.WithGlobal(false)).Load(),
	))
	cnf := app.AppConfiguration()

	// 检测数据库配置是否有效
	driverName, dsn := getDBConfig(cnf)
	if !isValidDBConfig(driverName, dsn) {
		// 数据库配置无效, 启动内置安装向导 (不依赖 woocoo 配置)
		setupAddr := getSetupAddr(cnf)
		runSetupWizard(driverName, dsn, setupAddr)
		return
	}

	// 数据库配置有效, 启动完整服务
	ents := koapp.BuildEntComponents(cnf)
	drv := ents["portal"]

	portalClient := ent.NewClient(ent.Driver(drv))
	schemahook.RegisterAllHooks(portalClient)
	casbinClient := casbinent.NewClient(casbinent.Driver(drv))
	defer casbinClient.Close()
	if cnf.Development {
		portalClient = portalClient.Debug()
		casbinClient = casbinClient.Debug()
	}

	// 初始化全局 authorizer (casbin)
	var authorizer *authzcasbin.Authorizer
	if cnf.IsSet("authz") {
		adapter, err := entadapter.NewAdapterWithClient(casbinClient)
		if err != nil {
			panic(err)
		}
		authorizer, err = authzcasbin.NewAuthorizer(cnf.Sub("authz"), authzcasbin.WithAdapter(adapter))
		if err != nil {
			panic(err)
		}
		security.SetDefaultAuthorizer(authorizer)
	}

	if err := fmterr.InitErrorHandler(cnf.Sub("errors")); err != nil {
		panic(err)
	}
	kosdk, err := api.NewSDK(cnf.Sub("kosdk"))
	if err != nil {
		panic(err)
	}

	c, ok := graphql.DefaultCache(cnf)
	if !ok {
		panic("no cache configured")
	}

	graphqlSvr := graphql.NewServer(cnf,
		graphql.WithPortalDB(portalClient),
		graphql.WithKOSdk(kosdk),
		graphql.WithCache(c),
	)

	authParser, err := conf.NewParserFromFile("etc/auth.yaml")
	if err != nil {
		panic(err)
	}

	authCnf := &conf.AppConfiguration{
		Configuration: conf.NewFromParse(authParser, conf.WithGlobal(false)),
	}
	authSvr, err := auth.NewServer(authCnf, auth.WithAuthDB(drv), auth.WithCache(c))
	if err != nil {
		panic(err)
	}
	casbinSvr, err := casbin.NewServer(authCnf, casbin.WithAuthDB(drv), casbin.WithAuthorizer(authorizer))
	if err != nil {
		panic(err)
	}

	app.RegisterServer(graphqlSvr, authSvr, casbinSvr, clientx.ChangeSet)

	if cnf.IsSet("job") {
		jobSrv, err := job.NewServer(cnf.Sub("job"))
		if err != nil {
			panic(err)
		}
		if err := jobSrv.InitJobs(portalClient, kosdk); err != nil {
			panic(err)
		}
		app.RegisterServer(jobSrv)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}

// getDBConfig 从配置中获取数据库连接信息
func getDBConfig(cnf *conf.AppConfiguration) (string, string) {
	if cnf == nil {
		return "", ""
	}
	parser := cnf.Parser()
	driverName, _ := parser.Get("store.portal.driverName").(string)
	dsn, _ := parser.Get("store.portal.dsn").(string)
	return driverName, dsn
}

// isValidDBConfig 检测数据库配置是否有效
func isValidDBConfig(driverName, dsn string) bool {
	if driverName == "" || dsn == "" {
		return false
	}
	return setupapi.CanConnect(driverName, dsn)
}

// getSetupAddr 从配置中获取安装服务地址, 没有配置则使用默认值
func getSetupAddr(cnf *conf.AppConfiguration) string {
	if cnf == nil {
		return setupapi.DefaultAddr
	}
	addr, _ := cnf.Parser().Get("setup.addr").(string)
	if addr == "" {
		return setupapi.DefaultAddr
	}
	return addr
}

// runSetupWizard 启动内置安装向导
func runSetupWizard(driverName, dsn, addr string) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())

	setupHandler := setupapi.NewHandler(driverName, dsn)
	setupHandler.RegisterRoutes(engine.Group("/setup"))

	// 构建访问 URL
	displayAddr := addr
	if addr[0] == ':' {
		displayAddr = "localhost" + addr
	}
	println("系统未初始化, 请访问 http://" + displayAddr + "/setup/ 完成安装")

	// 创建 HTTP server 以支持 graceful shutdown
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	// 安装完成后: 写入 .env, 先响应客户端, 再关闭服务器, 最后重启
	setupHandler.SetPostInstallHook(func(cfg setupapi.Config) {
		go func() {
			time.Sleep(1 * time.Second)
			if err := writeDotEnv(cfg.DriverName, cfg.DSN, cfg.RedisAddr); err != nil {
				println("警告: 写入 .env 失败: " + err.Error())
			}
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			srv.Shutdown(ctx)
			time.Sleep(2 * time.Second)
			autoRestart()
		}()
	})

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

// writeDotEnv 将数据库和 Redis 配置写入 etc/.env 文件, woocoo 启动时自动加载
func writeDotEnv(driverName, dsn, redisAddr string) error {
	content := fmt.Sprintf(
		"# Auto-generated by setup wizard\nSTORE_PORTAL_DRIVERNAME=%s\nSTORE_PORTAL_DSN=%s\nREDIS_ADDR=%s\n",
		driverName, dsn, redisAddr,
	)
	return os.WriteFile("etc/.env", []byte(content), 0644)
}

// autoRestart 自动重启当前进程
func autoRestart() {
	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "windows" {
		// Windows: 启动新进程后退出
		cmd := exec.Command(executable, os.Args[1:]...)
		cmd.Dir = "."
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
		os.Exit(0)
	} else {
		// Unix: 用 syscall.Exec 替换当前进程
		//nolint:gosec
		syscall.Exec(executable, os.Args, os.Environ())
	}
}
