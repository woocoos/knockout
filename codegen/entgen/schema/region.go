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
	"github.com/woocoos/knockout/ent/region"
)

// Region holds the schema definition for the Region entity.
type Region struct {
	ent.Schema
}

func (Region) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "region"},
		entgql.RelayConnection(),
		entgql.QueryField().Description("地区查询"),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Region) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the Region.
func (Region) Fields() []ent.Field {
	return []ent.Field{
		field.Int("parent_id").Optional().Default(0).Comment("父id，0为顶级"),
		field.String("name").MaxLen(100).Optional().Comment("地区中文名称"),
		field.String("name_en").MaxLen(100).Optional().Comment("地区英文名称"),
		field.String("short_code").MaxLen(50).Optional().Comment("编码"),
		field.String("zip_code").MaxLen(10).Optional().Comment("邮政编码"),
		field.Int("country_id").Optional().Comment("国家id"),
		field.Int32("display_sort").Optional().
			Annotations(entgql.OrderField("displaySort"), entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusActive.String()).Optional().Comment("状态"),
	}
}

// Edges of the Region.
func (Region) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Region.Type).
			From("parent").Unique().Field("parent_id"),
		edge.From("country", Country.Type).Ref("regions").Field("country_id").Unique(),
	}
}

func (Region) Hooks() []ent.Hook {
	return []ent.Hook{
		InitDisplaySortHook(region.Table),
	}
}
