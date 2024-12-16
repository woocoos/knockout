package quota

import (
	"context"
	"entgo.io/ent"
	"fmt"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/quota"
	"github.com/woocoos/knockout/ent/quotaitem"
)

type quotaKey struct{}

// SkipQuota returns a new context that skips the TenantRule interceptor/mutators.
func SkipQuota(parent context.Context) context.Context {
	return context.WithValue(parent, quotaKey{}, true)
}

// ifSkipQuota returns true if the TenantRule interceptor/mutators should be skipped.
func ifSkipQuota(ctx context.Context) bool {
	skip, _ := ctx.Value(quotaKey{}).(bool)
	return skip
}

// EntHook 创建一个通用的配额ent.Hook
//
// 当引入需要限额的资源的操作时,有一定的约束, 由于需要跟踪到对应资料所使用的目的用户,如果是批量类型的操作则无法支持.因为需要调用这类数据操作的写法.
// 如果需要跳过配额检查, 可以使用SkipQuota函数.
func EntHook(schemaType ItemCode, resource Resource) ent.Hook {
	RegisterQuotaResource(schemaType, resource)
	type client interface {
		Client() *gen.Client
	}
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if ifSkipQuota(ctx) {
				return next.Mutate(ctx, m)
			}
			// 获取 schema 类型, resourceRegistry应已经初始化,否则会出现竞态问题
			res, exists := resourceRegistry[schemaType]
			if !exists {
				return next.Mutate(ctx, m)
			}
			// 获取租户ID, 对于某些配额是以用户ID.
			tar, err := res.GetTarget(ctx, m)
			if err != nil {
				return nil, fmt.Errorf("failed to get org ID: %w", err)
			}

			var change int64
			op := m.Op()
			switch op {
			case ent.OpDelete:
				v, err := next.Mutate(ctx, m)
				if err != nil {
					return nil, err
				}
				change = int64(v.(int))
			default:
				// 计算资源变化量
				change, err = res.CalculateChange(ctx, m, op)
				if err != nil {
					return nil, fmt.Errorf("failed to calculate resource change: %w", err)
				}
			}
			if change == 0 {
				return next.Mutate(ctx, m)
			}
			mc, ok := m.(client)
			if !ok {
				return nil, fmt.Errorf("failed to get client from mutation")
			}
			// 查找相关的配额项
			quotaItem, err := mc.Client().QuotaItem.Query().
				Where(quotaitem.Code(res.QuotaItemCode)).
				Only(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to find quota item: %w", err)
			}

			q, err := mc.Client().Quota.Query().
				Where(
					quota.TenantID(tar.TenantID),
					quota.UserID(tar.UserID),
					quota.QuotaItemID(quotaItem.ID),
				).Only(ctx)
			if err != nil {
				if !gen.IsNotFound(err) {
					return nil, err
				}
				// 如果默认值为0, 说明不限制该配额.
				if quotaItem.DefaultLimit == 0 {
					return next.Mutate(ctx, m)
				}
				// 如果配额记录不存在，且是减少操作，则报错
				if change < 0 {
					return nil, fmt.Errorf("no quota record found for target %d:%d and item %s",
						tar.TenantID, tar.UserID, res.QuotaItemCode)
				}
				// 如果是增加操作，使用默认值创建新配额
				if quotaItem.DefaultLimit != 0 {
					q, err = mc.Client().Quota.Create().
						SetTenantID(tar.TenantID).
						SetUserID(tar.UserID).
						SetQuotaItemID(quotaItem.ID).
						SetLimit(quotaItem.DefaultLimit).
						SetUsed(change).
						Save(ctx)
				} else {
					return nil, fmt.Errorf("no quota limit defined for item %s", res.QuotaItemCode)
				}
			} else {
				// 检查是否超出配额
				newUsed := q.Used + change
				if newUsed > q.Limit {
					return nil, fmt.Errorf("quota exceeded for %s: limit %d, current %d, requested %d",
						res.QuotaItemCode, q.Limit, q.Used, change)
				}

				// 更新使用.
				q, err = mc.Client().Quota.UpdateOne(q).
					SetUsed(newUsed).SetUpdatedBy(0).
					Save(ctx)
			}
			if err != nil {
				return nil, err
			}

			// 执行原始操作
			return next.Mutate(ctx, m)
		})
	}
}
