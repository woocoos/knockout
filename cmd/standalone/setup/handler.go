package setup

import (
	"context"
	"embed"
	"io/fs"
	"net/http"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
)

//go:embed all:static
var staticFiles embed.FS

// DefaultAddr 安装服务默认地址
const DefaultAddr = ":8081"

// Request 初始化请求
type Request struct {
	// 数据库配置
	DriverName string `json:"driverName"`
	DSN        string `json:"dsn"`
	// Redis 配置
	RedisAddr string `json:"redisAddr"`
	// 管理员配置
	AdminPrincipalName string `json:"adminPrincipalName"`
	AdminDisplayName   string `json:"adminDisplayName"`
	AdminPassword      string `json:"adminPassword"`
	AdminEmail         string `json:"adminEmail"`
}

// Handler 系统初始化 handler
type Handler struct {
	driverName    string
	dsn           string
	postInstallFn func(Config)
}

// NewHandler 创建初始化 handler
func NewHandler(driverName, dsn string) *Handler {
	return &Handler{driverName: driverName, dsn: dsn}
}

// SetPostInstallHook 设置安装完成后的回调函数
func (h *Handler) SetPostInstallHook(fn func(Config)) {
	h.postInstallFn = fn
}

// RegisterRoutes 注册初始化路由到 gin.RouterGroup
func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/status", h.status)
	rg.POST("/test-db", h.testDB)
	rg.GET("/", h.indexPage)
	rg.POST("/", h.setup)
}

// guard 检查系统是否已初始化, 已初始化则返回 403
func (h *Handler) guard(c *gin.Context) bool {
	if IsInitialized(h.driverName, h.dsn) {
		c.JSON(http.StatusForbidden, gin.H{"error": "system already initialized"})
		return true
	}
	return false
}

func (h *Handler) indexPage(c *gin.Context) {
	if h.guard(c) {
		return
	}
	content, err := fs.ReadFile(staticFiles, "static/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to load setup page")
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", content)
}

func (h *Handler) status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"initialized": IsInitialized(h.driverName, h.dsn),
	})
}

func (h *Handler) testDB(c *gin.Context) {
	var req struct {
		DriverName string `json:"driverName"`
		DSN        string `json:"dsn"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid request body"})
		return
	}
	if req.DriverName == "" || req.DSN == "" {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "driverName and dsn are required"})
		return
	}
	if err := TestConnection(req.DriverName, req.DSN); err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *Handler) setup(c *gin.Context) {
	if h.guard(c) {
		return
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// 验证数据库配置
	if req.DriverName == "" || req.DSN == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "database configuration is required"})
		return
	}

	// 验证管理员配置
	if req.AdminPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "admin password is required"})
		return
	}
	if len(req.AdminPassword) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 6 characters"})
		return
	}

	cfg := Config{
		DriverName:         req.DriverName,
		DSN:                req.DSN,
		RedisAddr:          req.RedisAddr,
		AdminPrincipalName: req.AdminPrincipalName,
		AdminDisplayName:   req.AdminDisplayName,
		AdminPassword:      req.AdminPassword,
		AdminEmail:         req.AdminEmail,
	}

	if cfg.AdminPrincipalName == "" {
		cfg.AdminPrincipalName = "admin"
	}
	if cfg.AdminDisplayName == "" {
		cfg.AdminDisplayName = cfg.AdminPrincipalName
	}
	if cfg.AdminEmail == "" {
		cfg.AdminEmail = "admin@localhost"
	}
	if cfg.RedisAddr == "" {
		cfg.RedisAddr = "127.0.0.1:6379"
	}

	// 执行安装
	if err := Do(cfg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 更新 handler 的数据库配置 (用于后续 guard 检测)
	h.driverName = req.DriverName
	h.dsn = req.DSN

	// 发送响应
	c.JSON(http.StatusOK, gin.H{"message": "setup completed successfully"})

	// 触发 post-install hook (由调用方控制重启时序)
	if h.postInstallFn != nil {
		h.postInstallFn(Config{
			DriverName: req.DriverName,
			DSN:        req.DSN,
			RedisAddr:  req.RedisAddr,
		})
	}
}

// CanConnect 检测数据库是否可以连接
func CanConnect(driverName, dsn string) bool {
	drv, err := sql.Open(driverName, dsn)
	if err != nil {
		return false
	}
	defer drv.Close()
	_, err = drv.ExecContext(context.Background(), "SELECT 1")
	return err == nil
}

// TestConnection 测试数据库连接, 返回详细错误信息
func TestConnection(driverName, dsn string) error {
	drv, err := sql.Open(driverName, dsn)
	if err != nil {
		return err
	}
	defer drv.Close()
	_, err = drv.ExecContext(context.Background(), "SELECT 1")
	return err
}
