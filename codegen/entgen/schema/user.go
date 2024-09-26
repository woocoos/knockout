package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/fieldx"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
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
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NewSoftDeleteMixin[intercept.Query, *gen.Client](intercept.NewQuery),
		schemax.NotifyMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("principal_name").Unique().Comment("登陆名称").
			Annotations(entproto.Field(7)),
		field.String("display_name").Comment("显示名").Annotations(entproto.Field(8)),
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
		fieldx.File("avatar").MaxLen(255).Optional().Comment("头像地址").Annotations(
			entgql.Skip(entgql.SkipWhereInput), entproto.Skip()),
		field.Enum("gender").Default("privacy").NamedValues(
			"privacy", "privacy",
			"male", "male",
			"female", "female",
		).Comment("性别").Annotations(
			entproto.Field(13),
			entproto.Enum(map[string]int32{
				"privacy": 0,
				"male":    1,
				"female":  2,
			})),
		field.Int("citizenship_id").Optional().Nillable().Comment("国籍").Annotations(entproto.Field(14)),
		field.String("first_name").MaxLen(45).Optional().Comment("名字").Annotations(entproto.Field(15)),
		field.String("middle_name").MaxLen(45).Optional().Comment("中间名").Annotations(entproto.Field(16)),
		field.String("last_name").MaxLen(45).Optional().Comment("姓氏").Annotations(entproto.Field(17)),
		field.String("lang").Optional().Comment("语言").Annotations(entproto.Field(18)),
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
		edge.To("oauth_clients", OauthClient.Type).Comment("用户AccessKey").
			Annotations(entproto.Skip(), entgql.Skip(entgql.SkipMutationUpdateInput)),
		edge.To("addresses", UserAddr.Type).Comment("用户联系信息").
			Annotations(entproto.Skip(), entgql.Skip(entgql.SkipMutationUpdateInput)),
		edge.To("citizenship", Country.Type).Field("citizenship_id").Unique().Comment("国籍信息").
			Annotations(entproto.Skip()),
	}
}
