package hook

import (
	"context"
	"encoding/json"
	"entgo.io/ent"
	"github.com/gin-gonic/gin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/service/quota"
)

// OrgUserPreferenceQuotaHook 用户偏好配额hook.
func OrgUserPreferenceQuotaHook() ent.Hook {
	return quota.EntHook(quota.ItemCodeClientPreference, quota.Resource{
		QuotaItemCode: string(quota.ItemCodeClientPreference),
		CalculateChange: func(ctx context.Context, m ent.Mutation, op ent.Op) (int64, error) {
			switch op {
			case ent.OpCreate, ent.OpUpdateOne:
				if val, ok := m.Field("client_preferences"); ok {
					if cp, ok := val.([]types.ClientPreference); ok {
						cpStr, err := json.Marshal(cp)
						if err != nil {
							return 0, err
						}
						ocpStr := make([]byte, 0)
						if op == ent.OpUpdateOne {
							ocp, err := m.OldField(ctx, "client_preferences")
							if err != nil {
								return 0, err
							}
							if ocp.([]types.ClientPreference) == nil {
								return int64(len(cpStr)), nil
							}
							ocpStr, _ = json.Marshal(ocp)
							if err != nil {
								return 0, err
							}
						}
						change := len(cpStr) - len(ocpStr)
						return int64(change), nil
					}
				}
				return 0, nil
			case ent.OpDeleteOne:
				// 删除置零处理
				oupm := m.(*gen.OrgUserPreferenceMutation)
				if id, ok := oupm.ID(); ok {
					oup, err := oupm.Client().OrgUserPreference.Get(quota.SkipQuota(ctx), id)
					if err != nil {
						return 0, err
					}
					ocpStr, err := json.Marshal(oup.ClientPreferences)
					if err != nil {
						return 0, err
					}
					return -int64(len(ocpStr)), nil
				}
				return 0, nil
			default:
				return 0, quota.ErrNoCurrentOp
			}
		},
		GetTarget: func(ctx context.Context, m ent.Mutation) (*quota.Target, error) {
			// 批量删除不支持
			if m.Op().Is(ent.OpDelete) && !m.Op().Is(ent.OpDeleteOne) {
				return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "operation not supported: %s", m.Op())
			}

			// 尝试从 mutation 中获取 UserID
			var userID, tenantID int
			if fd, ok := m.Field("user_id"); ok {
				if id, ok := fd.(int); ok {
					userID = id
				}
			}
			if fd, ok := m.Field("org_id"); ok {
				if id, ok := fd.(int); ok {
					tenantID = id
				}
			}
			if userID != 0 && tenantID != 0 {
				return &quota.Target{
					UserID:   userID,
					TenantID: tenantID,
				}, nil
			}
			// 如果是 DeleteOne 操作，尝试从数据库获取 UserID
			if m.Op().Is(ent.OpDeleteOne) {
				oupm := m.(*gen.OrgUserPreferenceMutation)
				if id, ok := oupm.ID(); ok {
					oup, err := oupm.Client().OrgUserPreference.Get(quota.SkipQuota(ctx), id)
					if err != nil {
						return nil, err
					}
					return &quota.Target{
						UserID:   oup.UserID,
						TenantID: oup.OrgID,
					}, nil
				}
			}
			// 如果是 UpdateOne 操作，尝试从现有记录中获取 UserID
			if m.Op().Is(ent.OpUpdateOne) {
				oUserID, err := m.OldField(ctx, "user_id")
				if err == nil {
					if id, ok := oUserID.(int); ok {
						userID = id
					}
				}
				oTenantID, err := m.OldField(ctx, "org_id")
				if err == nil {
					if id, ok := oTenantID.(int); ok {
						tenantID = id
					}
				}
				if userID != 0 && tenantID != 0 {
					return &quota.Target{
						UserID:   userID,
						TenantID: tenantID,
					}, nil
				}
			}

			return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "operation not supported: %s", m.Op())
		},
	})
}