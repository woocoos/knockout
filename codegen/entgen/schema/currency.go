package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
)

// Currency holds the schema definition for the Currency entity.
type Currency struct {
	ent.Schema
}

func (Currency) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "currency"},
		entgql.RelayConnection(),
		entgql.QueryField().Description("货币查询"),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Currency) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the Currency.
func (Currency) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").MaxLen(45).Comment("代码"),
		field.String("name").MaxLen(45).Comment("中文名称"),
		field.String("sign").MaxLen(45).Optional().Comment("货币符号"),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusActive.String()).Comment("状态"),
	}
}
