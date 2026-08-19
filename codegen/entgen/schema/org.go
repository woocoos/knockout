package schema

import (
	"regexp"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/codegen/entgen/types"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/intercept"
)

// Org 组织目录定义,是企业目录的容器.
type Org struct {
	ent.Schema
}

func (Org) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "org"},
		entgql.RelayConnection(),
		entgql.QueryField("organizations"),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Org) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
		schemax.NewSoftDeleteMixin[intercept.Query, *gen.Client](intercept.NewQuery),
		schemax.NotifyMixin{},
	}
}

// Fields of the Org.
func (Org) Fields() []ent.Field {
	return []ent.Field{
		field.Int("owner_id").Optional().Nillable().Comment("管理账户ID,如果设置则该组织将升级为根组织"),
		field.Enum("kind").NamedValues(
			"root", "root",
			"organization", "org",
		).Default("org").Comment("分类: 根节点,组织节点").Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Int("parent_id").Default(0).Comment("父级ID,0为根组织."),
		field.String("domain").Optional().Unique().Comment("默认域名").
			Match(regexp.MustCompile(`^(?:(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]|)$`)),
		field.Strings("custom_domain").Optional().Comment("自定义域名"),
		field.String("code").MaxLen(45).Optional().Comment("系统代码").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("name").MaxLen(100).Comment("组织名称"),
		field.String("profile").Comment("简介").Optional().Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusActive.String()).Optional().Comment("状态"),
		field.Text("path").Optional().Comment("路径编码").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Int32("display_sort").Optional().
			Annotations(entgql.OrderField("displaySort"), entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("country_code").MaxLen(10).Optional().Comment("国家或地区2字码"),
		field.String("timezone").MaxLen(45).Optional().Comment("时区"),
		field.String("local_currency").MaxLen(10).Optional().Comment("组织本位币"),
		field.JSON("logo", &types.OrgLogo{}).Optional().Comment("组织图标"),
	}
}

// Edges of the Org.
func (Org) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Org.Type).
			From("parent").Unique().Required().Field("parent_id"),
		edge.To("owner", User.Type).Field("owner_id").Unique().Comment("管理账户"),
		edge.To("users", User.Type).Through("org_user", OrgUser.Type).
			Annotations(entgql.RelayConnection()).Comment("组织下用户"),
		edge.To("roles_and_groups", OrgRole.Type).
			Annotations(entgql.Skip(entgql.SkipType)).Comment("组织下角色及用户组."),
		edge.To("permissions", Permission.Type).Comment("组织授权信息").
			Annotations(entgql.RelayConnection()),
		edge.To("policies", OrgPolicy.Type).Comment("组织下权限策略").
			Annotations(entgql.RelayConnection()),
		edge.To("apps", App.Type).Comment("组织下应用").Through("org_app", OrgApp.Type).
			Annotations(entgql.RelayConnection()),
		edge.To("file_identities", FileIdentity.Type).Comment("组织下文件凭证"),
		edge.To("user_password_policy", UserPasswordPolicy.Type).Unique().Comment("组织下密码策略"),
		edge.To("org_quota", Quota.Type).Comment("组织下登录策略").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}
