package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/woocoos/knockout-go/ent/schemax"
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
		field.Int64("default_limit").Optional().Comment("默认限制值"),
	}
}

func (QuotaItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("quota", Quota.Type).Annotations(entgql.RelayConnection()),
	}
}

// Quota 租户配额限制. 一个租户可以有多个不同配额限制,一类配额只能一种.
// 注意配额是一个高权限功能,需要注意对后台用户的授权
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
		field.Int("tenant_id").Comment("租户ID,来源于root的组织ID."),
		field.Int("user_id").Comment("来源于用户ID"),
		field.Int("quota_item_id").Comment("配额项ID"),
		field.Int64("limit").Min(0).Comment("限制值").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Int64("used").Default(0).Comment("已使用值").Annotations(
			entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipMutationCreateInput),
		),
		field.Time("start_at").Optional().Comment("生效时间"),
		field.Time("end_at").Optional().Comment("过期时间"),
	}
}

func (Quota) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("quota_item", QuotaItem.Type).Ref("quota").
			Field("quota_item_id").Unique().Required().Comment("配额定义"),
	}
}

// Indexes of the Quota.
func (Quota) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_id", "user_id", "quota_item_id").
			Unique(),
	}
}

func (Quota) Hooks() []ent.Hook {
	return nil
}
