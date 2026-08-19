package hook

import (
	"context"

	"entgo.io/ent"
	"github.com/gin-gonic/gin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/quota"
)

// UserDeviceQuotaHook 用户设备配额hook.
func UserDeviceQuotaHook() ent.Hook {
	return quota.EntHook(quota.ItemCodeUserDevice, quota.Resource{
		QuotaItemCode: string(quota.ItemCodeUserDevice),
		CalculateChange: func(_ context.Context, m ent.Mutation, op ent.Op) (int64, error) {
			switch op {
			case ent.OpCreate:
				return 1, nil
			case ent.OpDeleteOne:
				return -1, nil
			case ent.OpUpdateOne:
				return 0, nil
			default:
				return 0, quota.ErrNoCurrentOp
			}
		},
		GetTarget: func(ctx context.Context, m ent.Mutation) (*quota.Target, error) {
			// 对于删除操作，返回不支持
			if m.Op().Is(ent.OpDelete) {
				return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "operation not supported: %s", m.Op())
			}

			// 尝试从 mutation 中获取 UserID
			if fd, ok := m.Field("user_id"); ok {
				if id, ok := fd.(int); ok {
					return &quota.Target{
						UserID: id,
					}, nil
				}
			}
			// 如果是 DeleteOne 操作，尝试从数据库获取 UserID
			if m.Op().Is(ent.OpDeleteOne) {
				udm := m.(*gen.UserDeviceMutation)
				if id, ok := udm.ID(); ok {
					ud, err := udm.Client().UserDevice.Get(ctx, id)
					if err != nil {
						return nil, err
					}
					return &quota.Target{
						UserID: ud.UserID,
					}, nil
				}
			}
			// 如果是 UpdateOne 操作，尝试从现有记录中获取 UserID
			if m.Op().Is(ent.OpUpdateOne) {
				oldValues, err := m.OldField(ctx, "user_id")
				if err == nil {
					if id, ok := oldValues.(int); ok {
						return &quota.Target{
							UserID: id,
						}, nil
					}
				}
			}

			return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "operation not supported: %s", m.Op())
		},
	})
}
