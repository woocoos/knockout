package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

func (Country) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "country"},
		entgql.RelayConnection(),
		entgql.QueryField().Description("国家查询"),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Country) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional().Comment("国家中文名称"),
		field.String("name_en").Optional().Comment("国家英文名称"),
		field.String("code").MaxLen(50).Unique().Comment("国家代码"),
		field.Int32("display_sort").Optional().
			Annotations(entgql.OrderField("displaySort"), entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusActive.String()).Optional().Comment("状态"),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("regions", Region.Type).Comment("地区信息"),
	}
}
