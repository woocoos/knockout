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
	"github.com/woocoos/knockout-go/ent/schemax/typex"
)

// AppDictItem holds the schema definition for the AppDictItem entity.
type AppDictItem struct {
	ent.Schema
}

func (AppDictItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_dict_item"},
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (AppDictItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the AppDictItem.
func (AppDictItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Optional().Immutable().Comment("租户ID,空为全局字典"),
		field.Int("dict_id").Optional().Immutable().Comment("所属字典"),
		field.String("ref_code").Comment("关联代码,由app_code和dict_code组成").Annotations(entgql.Skip(
			entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		field.String("code").MinLen(3).MaxLen(20).Immutable().
			Comment("字典值唯一编码,生效后不可修改."),
		field.String("name").MaxLen(45).Comment("名称"),
		field.String("comments").Optional().Comment("备注").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Int32("display_sort").Optional().Annotations(entgql.OrderField("displaySort"),
			entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusInactive.String()).
			Optional().Comment("状态"),
	}
}

// Edges of the AppDictItem.
func (AppDictItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dict", AppDict.Type).Ref("items").Unique().Immutable().Field("dict_id"),
		edge.To("org", Org.Type).Unique().Immutable().Field("org_id"),
	}
}

func (AppDictItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ref_code"),
	}
}
