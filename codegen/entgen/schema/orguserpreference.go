package schema

import (
	"context"
	"encoding/json"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout/codegen/entgen/types"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/quota"
)

// OrgUserPreference holds the schema definition for the OrgUserPreference entity.
type OrgUserPreference struct {
	ent.Schema
}

func (OrgUserPreference) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org_user_preference"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (OrgUserPreference) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the OrgUserPreference.
func (OrgUserPreference) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Comment("用户id"),
		field.Int("org_id").Immutable().Comment("租户ID").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.JSON("menu_favorite", []int{}).Optional().Comment("用户收藏菜单").Annotations(entgql.Type("[ID!]")),
		field.JSON("menu_recent", []int{}).Optional().Comment("用户最近访问菜单").Annotations(entgql.Type("[ID!]")),
		field.JSON("client_preferences", []types.ClientPreference{}).Optional().Comment("客户端偏好设置"),
	}
}

// Edges of the OrgUserPreference.
func (OrgUserPreference) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required().Field("user_id").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		edge.To("org", Org.Type).Immutable().Unique().Required().Field("org_id").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Hooks of the UserDevice.
func (p OrgUserPreference) Hooks() []ent.Hook {
	return []ent.Hook{
		p.quotaHook(),
	}
}

// quotaHook 对于Delete where操作,将不支持.
func (p OrgUserPreference) quotaHook() ent.Hook {
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
						ocp, err := m.OldField(ctx, "client_preferences")
						if err != nil {
							return 0, err
						}
						if ocp.([]types.ClientPreference) == nil {
							return int64(len(cpStr)), nil
						}
						ocpStr, err := json.Marshal(ocp)
						if err != nil {
							return 0, err
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
					oup, err := oupm.Client().OrgUserPreference.Get(ctx, id)
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
			// 对于删除操作，返回不支持
			if m.Op().Is(ent.OpDelete) {
				return nil, fmt.Errorf("operation not supported: %s", m.Op())
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
					oup, err := oupm.Client().OrgUserPreference.Get(ctx, id)
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

			return nil, fmt.Errorf("operation not supported: %s", m.Op())
		},
	})
}
