package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/quota"
)

// UserDevice holds the schema definition for the UserDevice entity.
type UserDevice struct {
	ent.Schema
}

// Annotations
//
// 用户信息暂时不需要通过直接的数据操作,因此未接入gql mutation
func (UserDevice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_device"},
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (UserDevice) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the UserDevice.
func (UserDevice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional().Immutable(),
		field.String("device_uid").MaxLen(64).Comment("设备唯一ID"),
		field.String("device_name").MaxLen(45).Optional().Comment("设备名称"),
		field.String("system_name").MaxLen(45).Optional().Comment("系统名称"),
		field.String("system_version").MaxLen(45).Optional().Comment("系统版本"),
		field.String("app_version").MaxLen(45).Optional().Comment("app版本"),
		field.String("device_model").MaxLen(45).Optional().Comment("设备型号"),
		field.Enum("status").GoType(typex.SimpleStatus("")).Optional().Comment("状态,可用或不可用及其他待确认状态"),
		field.String("comments").Optional().Comment("备注").Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the UserDevice.
func (UserDevice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("devices").Field("user_id").Unique().Immutable(),
	}
}

// Hooks of the UserDevice.
func (UserDevice) Hooks() []ent.Hook {
	return []ent.Hook{
		quotaHook(),
	}
}

// quotaHook 对于Delete where操作,将不支持.
func quotaHook() ent.Hook {
	return quota.EntHook(quota.ItemCodeUserDevice, quota.Resource{
		QuotaItemCode: string(quota.ItemCodeUserDevice),
		CalculateChange: func(_ context.Context, m ent.Mutation, op ent.Op) (int64, error) {
			switch op {
			case ent.OpCreate:
				return 1, nil
			case ent.OpDeleteOne:
				return -1, nil
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

			return nil, fmt.Errorf("operation not supported: %s", m.Op())
		},
	})
}
