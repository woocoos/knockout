package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/entco/schemax/typex"
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
		field.String("device_name").MaxLen(45).Optional(),
		field.String("system_name").MaxLen(45).Optional(),
		field.String("system_version").MaxLen(45).Optional(),
		field.String("app_version").MaxLen(45).Optional(),
		field.String("device_model").MaxLen(45).Optional(),
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
