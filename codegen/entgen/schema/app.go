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

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

func (App) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app"},
		entgql.RelayConnection(),
		entgql.QueryField().Description("公开应用查询"),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (App) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
	}
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(45).Unique().Comment("名称"),
		field.String("code").MinLen(3).MaxLen(10).Immutable().Unique().Comment("用于标识应用资源的唯一代码,尽量简短"),
		field.Enum("kind").NamedValues(
			"web", "web",
			"native", "native",
			"server", "server",
		).Comment("应用类型"),
		field.String("redirect_uri").MaxLen(255).Optional().Comment("回调地址"),
		field.String("app_key").Optional().Comment("应用ID"),
		field.String("app_secret").MaxLen(255).Optional().Comment("应用密钥"),
		field.String("scopes").MaxLen(500).Optional().Comment("权限范围"),
		field.Int32("token_validity").Optional().Comment("token有效期"),
		field.Int32("refresh_token_validity").Optional().Comment("refresh_token有效期"),
		field.String("logo").Optional().Comment("图标"),
		field.String("comments").Optional().Comment("备注"),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusActive.String()).Optional().Comment("状态"),
		field.Bool("private").Optional().Default(false).Comment("私有App,表示由组织创建").
			Annotations(entgql.Skip(entgql.SkipAll)),
		field.Int("org_id").Optional().Comment("创建的根组织ID").Annotations(entgql.Skip(entgql.SkipAll)),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("menus", AppMenu.Type).Comment("菜单").Annotations(entgql.RelayConnection()),
		edge.To("actions", AppAction.Type).Comment("权限").Annotations(entgql.RelayConnection()),
		edge.To("resources", AppRes.Type).Comment("资源").Annotations(entgql.RelayConnection()),
		edge.To("roles", AppRole.Type).Comment("角色"),
		edge.To("policies", AppPolicy.Type).Comment("策略"),
		edge.From("orgs", Org.Type).Ref("apps").Comment("使用该应用的组织").
			Through("org_app", OrgApp.Type).
			Annotations(entgql.RelayConnection(), entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}
