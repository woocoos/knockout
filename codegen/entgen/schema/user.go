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
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/intercept"
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
		schemax.NewSoftDeleteMixin[intercept.Query, *gen.Client](intercept.NewQuery),
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
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
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
		field.Int("avatar_file_id").Optional().Comment("头像,存储路规则：/{appcode}/{tid}/xxx").Annotations(
			entgql.Skip(entgql.SkipWhereInput), entproto.Skip(), entgql.Type("ID")),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("identities", UserIdentity.Type).Comment("用户身份标识").
			Annotations(entproto.Skip(), entgql.Skip(entgql.SkipMutationUpdateInput)),
		edge.To("login_profile", UserLoginProfile.Type).Unique().Comment("登陆设置").
			Annotations(entproto.Skip(), entgql.Skip(entgql.SkipMutationUpdateInput)),
		edge.To("passwords", UserPassword.Type).Comment("用户密码").
			Annotations(entproto.Skip(), entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipType)),
		edge.To("devices", UserDevice.Type).Comment("用户设备").
			Annotations(entproto.Skip(), entgql.Skip(entgql.SkipMutationUpdateInput)),
		edge.From("orgs", Org.Type).Ref("users").Comment("用户所属组织").
			Through("org_user", OrgUser.Type).
			Annotations(entgql.Skip(entgql.SkipAll), entproto.Skip()),
		edge.From("permissions", Permission.Type).Ref("user").Comment("用户权限").
			Annotations(entgql.RelayConnection(), entproto.Skip(),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}
