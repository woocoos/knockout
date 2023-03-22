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

// UserPassword 管理用户密码
type UserPassword struct {
	ent.Schema
}

func (UserPassword) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_password"},
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (UserPassword) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the UserPassword.
func (UserPassword) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Immutable().Optional(),
		field.Enum("scene").Values("login").Comment("场景: login 普通登陆"),
		field.String("password").Optional().Comment("密码").Sensitive(),
		field.String("salt").MaxLen(45).Comment("盐").Sensitive().Annotations(entgql.Skip(entgql.SkipAll)),
		field.Enum("status").GoType(typex.SimpleStatus("")).Optional(),
		field.String("memo").Optional().Annotations(entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Edges of the UserPassword.
func (UserPassword) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("passwords").Field("user_id").Immutable().Unique(),
	}
}
