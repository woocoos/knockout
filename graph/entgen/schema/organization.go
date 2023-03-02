package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"fmt"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/entco/schemax/typex"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/hook"
	"github.com/woocoos/knockout/ent/organization"
	"regexp"
	"strconv"
)

// Organization 组织目录定义,是企业目录的容器.
type Organization struct {
	ent.Schema
}

func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "organization"},
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}

func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
		SoftDeleteMixin{},
	}
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.Int("owner_id").Optional().Comment("管理账户ID"),
		field.Enum("kind").NamedValues(
			"root", "root",
			"organization", "org",
		).Optional().Comment("分类: 根节点,组织节点").Annotations(entgql.Skip(entgql.SkipAll)),
		field.Int("parent_id").Default(0).Comment("父级ID,0为根组织."),
		field.String("domain").Optional().Unique().Comment("默认域名").
			Match(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]\.[a-zA-Z]{2,}$`)),
		field.String("code").MaxLen(45).Optional().Comment("系统代码").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("name").MaxLen(100).Comment("组织名称"),
		field.String("profile").Comment("简介").Optional().Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Enum("status").GoType(typex.SimpleStatus("")).Optional().Comment("状态"),
		field.Text("path").Optional().Comment("路径编码").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Int32("display_sort").Optional().
			Annotations(entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("country_code").MaxLen(10).Optional().Comment("国家或地区2字码"),
		field.String("timezone").MaxLen(45).Optional().Comment("时区"),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Organization.Type).
			From("parent").Unique().Required().Field("parent_id"),
		edge.To("owner", User.Type).Field("owner_id").Unique().Comment("管理账户").
			Annotations(entgql.Skip(entgql.SkipType)),
		edge.To("users", User.Type).Through("organization_user", OrganizationUser.Type).
			Annotations(entgql.RelayConnection()).Comment("组织下用户"),
		edge.To("rolesAndGroups", OrganizationRole.Type).
			Annotations(entgql.Skip(entgql.SkipType)).Comment("组织下角色及用户组."),
		edge.To("permissions", Permission.Type).Comment("组织授权信息").
			Annotations(entgql.RelayConnection()),
		edge.To("policies", PermissionPolicy.Type).Comment("组织下权限策略").
			Annotations(entgql.RelayConnection()),
		edge.To("apps", App.Type).Comment("组织下应用").Through("organization_app", OrganizationApp.Type).
			Annotations(entgql.RelayConnection()),
	}
}

func (Organization) Hooks() []ent.Hook {
	return []ent.Hook{
		pathHook(),
		InitDisplaySortHook(organization.Table),
		hook.On(checkDeleteHook(), ent.OpDeleteOne),
	}
}

func pathHook() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.OrganizationFunc(func(ctx context.Context, mutation *gen.OrganizationMutation) (gen.Value, error) {
				if pid, ok := mutation.ParentID(); ok {
					id, _ := mutation.ID()
					code := strconv.FormatInt(int64(id), 36)
					if pid == 0 {
						mutation.SetPath(code)
					} else {
						prow, err := mutation.Client().Organization.Query().Where(organization.ID(pid)).
							Select(organization.FieldPath).First(ctx)
						if err != nil && !gen.IsNotFound(err) {
							return nil, err
						}
						mutation.SetPath(prow.Path + "/" + code)
					}
					if c, _ := mutation.Code(); c == "" {
						mutation.SetCode(code)
					}
				}
				return next.Mutate(ctx, mutation)
			})
		}, ent.OpCreate|ent.OpUpdateOne)
}

func checkDeleteHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.OrganizationFunc(func(ctx context.Context, mutation *gen.OrganizationMutation) (gen.Value, error) {
			if mutation.Op() == ent.OpDeleteOne {
				if id, ok := mutation.ID(); ok {
					count, err := mutation.Client().Organization.Query().Where(
						organization.ParentID(id),
					).Count(ctx)
					if err != nil {
						return nil, err
					}
					if count > 0 {
						return nil, fmt.Errorf("organization has children")
					}
				}
			}
			return next.Mutate(ctx, mutation)
		})
	}
}
