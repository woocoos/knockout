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
		entgql.QueryField(),
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
		field.String("code").MaxLen(45).Unique().Comment("代码"),
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
		field.Int("created_org_id").Optional().Comment("创建租户").Annotations(entgql.Skip(entgql.SkipAll)),
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
		edge.From("organizations", Organization.Type).Ref("apps").Comment("使用该应用的组织").
			Through("organization_app", OrganizationApp.Type).
			Annotations(entgql.RelayConnection()),
	}
}
