package quota

import (
	"context"
	"entgo.io/ent"
	"errors"
	"sync"
)

var (
	ErrNoCurrentOp = errors.New("not correct ent operation")
)

type ItemCode string

const (
	// ItemCodeUserDevice 配额项代码: 用户设备
	ItemCodeUserDevice ItemCode = "user_device"
	// ItemCodeOrgUser 配额项代码: 组织用户
	ItemCodeOrgUser ItemCode = "org_user"
)

// Target 定义了配额目标,租户或者用户
type Target struct {
	TenantID int
	UserID   int
}

// Resource 定义了需要跟踪配额的资源
type Resource struct {
	// 配额项代码
	QuotaItemCode string
	// 获取配额目标.目标包括租户或者用户.根据资源类型不同,获取目标方式也不同.
	GetTarget func(context.Context, ent.Mutation) (*Target, error)
	// 计算资源使用量变化的函数 (正数表示增加，负数表示减少)
	CalculateChange func(context.Context, ent.Mutation, ent.Op) (int64, error)
}

var (
	resourceRegistry = make(map[ItemCode]Resource)
	registryMu       sync.RWMutex
)

// RegisterQuotaResource 注册需要跟踪配额的资源
func RegisterQuotaResource(schemaType ItemCode, resource Resource) {
	registryMu.Lock()
	defer registryMu.Unlock()
	resourceRegistry[schemaType] = resource
}
