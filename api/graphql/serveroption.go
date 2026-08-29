package graphql

import (
	"github.com/tsingsun/woocoo/pkg/cache"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout/ent"
)

type ServerOption func(srv *ServerOptions)

func WithPortalDB(db *ent.Client) ServerOption {
	return func(srv *ServerOptions) {
		srv.portalDB = db
	}
}

func WithKOSdk(kosdk *api.SDK) ServerOption {
	return func(srv *ServerOptions) {
		srv.kosdk = kosdk
	}
}

func WithCache(c cache.Cache) ServerOption {
	return func(srv *ServerOptions) {
		srv.cache = c
	}
}

// DefaultCache 从配置中获取默认 cache 实例.
// 优先读取 cache.default 指定的 driverName; 若未配置 default 但只有一个 cache, 则该 cache 为默认.
func DefaultCache(cnf *conf.AppConfiguration) (cache.Cache, bool) {
	cacheCnf := cnf.Sub("cache")
	// 优先使用 cache.default 指定的 driverName
	if driverName := cacheCnf.String("default"); driverName != "" {
		if c, err := cache.GetCache(driverName); err == nil {
			return c, true
		}
	}
	// 未配置 default, 若只有一个 cache 则使用它
	var (
		c     cache.Cache
		count int
	)
	cnf.Map("cache", func(root string, sub *conf.Configuration) {
		if root == "default" {
			return
		}
		count++
		if c == nil {
			driverName := sub.String("driverName")
			if driverName == "" {
				driverName = root
			}
			if inst, err := cache.GetCache(driverName); err == nil {
				c = inst
			}
		}
	})
	if count == 1 && c != nil {
		return c, true
	}
	return nil, false
}

// WithDefaultCache 从配置中获取默认 cache 实例
func WithDefaultCache(cnf *conf.AppConfiguration) ServerOption {
	if c, ok := DefaultCache(cnf); ok {
		return WithCache(c)
	}
	return func(srv *ServerOptions) {}
}
