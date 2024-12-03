package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"fmt"
	"github.com/woocoos/knockout-go/ent/schemax"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/hook"
	"time"
)

// QuotaItem 配额项定义
type QuotaItem struct {
	ent.Schema
}

func (QuotaItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "quota_item"},
		entgql.RelayConnection(),
		entgql.QueryField().Description("配额定义"),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (QuotaItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

func (QuotaItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Unique().Comment("配额项代码,如: users,orgs"),
		field.String("name").Comment("配额项名称"),
		field.String("description").Optional().Comment("描述").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Enum("resource_type").
			NamedValues(
				"number", "number", // 数值类型
				"storage", "storage", // 存储容量
				"network", "network", // 网络带宽
			).Comment("资源类型"),
		field.String("unit").Optional().Comment("单位,如: 个,MB,GB").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Bool("active").Default(true).Comment("是否启用"),
	}
}

func (QuotaItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("quota", Quota.Type).Annotations(entgql.RelayConnection()),
	}
}

// Quota 租户配额限制. 一个租户可以有多个不同配额限制,一类配额只能一种.
type Quota struct {
	ent.Schema
}

func (Quota) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "quota"},
		entgql.RelayConnection(),
		entgql.QueryField("quotas").Description("配额管理"),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Quota) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the Quota.
//
// 目前暂时不添加审批流,只启用已使用值.对于生效期也暂时不限制.
func (Quota) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Comment("组织ID,为root型组织"),
		field.Int("quota_item_id").Comment("配额项ID"),
		field.Int64("limit").Comment("限制值").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Int64("used").Default(0).Comment("已使用值").Annotations(
			entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipMutationCreateInput),
		),
		field.Time("start_at").Optional().Comment("生效时间"),
		field.Time("end_at").Optional().Comment("过期时间"),
	}
}

func (Quota) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("org", Org.Type).
			Field("org_id").
			Unique().
			Required(),
		edge.From("quota_item", QuotaItem.Type).Ref("quota").
			Field("quota_item_id").Unique().Required().Comment("配额定义"),
	}
}

// Indexes of the Quota.
func (Quota) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("org_id", "quota_item_id").
			Unique(),
	}
}

func (Quota) Hooks() []ent.Hook {
	return []ent.Hook{
		quotaLimitHook(),
	}
}

// quotaLimitHook handles the business logic for quota limit changes
func quotaLimitHook() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.QuotaFunc(func(ctx context.Context, m *gen.QuotaMutation) (gen.Value, error) {
				now := time.Now()

				// 检查是否更新了limit
				if limit, exists := m.Limit(); exists {
					// 获取或设置生效时间
					var startAt time.Time
					if st, exists := m.StartAt(); exists {
						startAt = st
					} else if m.Op().Is(ent.OpCreate) {
						startAt = now
						m.SetStartAt(startAt)
					} else {
						// 对于更新操作，需要获取现有记录
						id, _ := m.ID()
						existing, err := m.Client().Quota.Get(ctx, id)
						if err != nil {
							return nil, fmt.Errorf("failed to get existing quota: %w", err)
						}
						startAt = existing.StartAt
					}

					// 获取结束时间（如果有）
					var endAt *time.Time
					if et, exists := m.EndAt(); exists {
						endAt = &et
					} else if m.Op().Is(ent.OpCreate) {
						// 创建时没有设置结束时间，保持为空
					} else {
						// 更新时获取现有的结束时间
						id, _ := m.ID()
						existing, err := m.Client().Quota.Get(ctx, id)
						if err != nil {
							return nil, fmt.Errorf("failed to get existing quota: %w", err)
						}
						endAt = &existing.EndAt
					}

					// 检查时间有效性
					if startAt.After(now) {
						// 生效时间在未来，设置used为0
						// 不更新, 保留原有值
					} else if endAt != nil && endAt.Before(now) {
						// 已过期, 提示错误.
						return nil, fmt.Errorf("record invalid,pleas set endAt")
					} else {
						// 当前在有效期内
						m.SetUsed(limit)
					}
				}

				return next.Mutate(ctx, m)
			})
		},
		ent.OpCreate|ent.OpUpdateOne,
	)
}
