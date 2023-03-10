package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/entco/schemax/typex"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
		entproto.Message(),
		entproto.Service(
			entproto.Methods(entproto.MethodGet),
		),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("principal_name").Unique().Comment("登陆名称").
			Annotations(entproto.Field(7)),
		field.String("display_name").Comment("显示名").Annotations(entproto.Field(8)),
		field.String("email").MaxLen(45).Optional().Comment("邮箱").Annotations(entproto.Field(9)),
		field.String("mobile").MaxLen(45).Optional().Comment("手机").Annotations(entproto.Field(10)),
		field.Enum("user_type").NamedValues(
			"account", "account",
			"member", "member",
		).Comment("用户类型").
			Annotations(
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
				entproto.Field(11),
				entproto.Enum(map[string]int32{
					"account": 1,
					"member":  2,
				})),
		field.Enum("creation_type").NamedValues(
			"invitation", "invitation",
			"register", "register",
			"manual", "manual",
		).Comment("创建类型,邀请，注册,手工创建").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
				entproto.Field(12),
				entproto.Enum(map[string]int32{
					"invitation": 1,
					"register":   2,
					"manual":     3,
				})),
		field.String("register_ip").Comment("注册时IP").Nillable().MaxLen(45).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
				entproto.Skip()),
		field.Enum("status").GoType(typex.SimpleStatus("")).Optional().Comment("状态").
			Annotations(entgql.Skip(entgql.SkipMutationUpdateInput),
				entproto.Skip(),
			),
		field.String("comments").Optional().Comment("备注").Annotations(
			entgql.Skip(entgql.SkipWhereInput), entproto.Skip()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("identities", UserIdentity.Type).Comment("用户身份标识").Annotations(entproto.Skip()),
		edge.To("login_profile", UserLoginProfile.Type).Unique().Annotations(entproto.Skip()).Comment("登陆设置"),
		edge.To("passwords", UserPassword.Type).Annotations(entproto.Skip()).Comment("用户密码"),
		edge.To("devices", UserDevice.Type).Annotations(entproto.Skip()).Comment("用户设备"),
		edge.From("orgs", Org.Type).Ref("users").Comment("用户所属组织").
			Through("org_user", OrgUser.Type).
			Annotations(entgql.Skip(entgql.SkipAll), entproto.Skip()),
		edge.From("permissions", Permission.Type).Ref("user").Comment("用户权限").
			Annotations(entgql.RelayConnection(), entproto.Skip()),
	}
}
