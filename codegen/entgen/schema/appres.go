package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
)

// AppRes 是应用下的资源信息
type AppRes struct {
	ent.Schema
}

func (AppRes) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_res"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (AppRes) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the AppRes.
func (AppRes) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Optional().Immutable().Comment("所属应用"),
		field.String("name").Comment("资源名称"),
		field.String("type_name").Comment("资源类型名称,如数据库表名"),
		field.String("arn_pattern").Comment("应用资源表达式"),
	}
}

// Edges of the AppRes.
func (AppRes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("resources").Unique().Immutable().Field("app_id"),
	}
}
