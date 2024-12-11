package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/woocoos/knockout-go/ent/schemax"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/apppolicyview"
	"github.com/woocoos/knockout/ent/hook"
	"strconv"
)

// AppPolicyView holds the schema definition for the AppPolicyView entity.
type AppPolicyView struct {
	ent.Schema
}

func (AppPolicyView) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_policy_view"},
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (AppPolicyView) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.SnowFlakeID{},
		schemax.AuditMixin{},
		schemax.NotifyMixin{},
	}
}

// Fields of the AppPolicyView.
func (AppPolicyView) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Optional().Immutable().Comment("所属应用"),
		field.Int("parent_id").Default(0).Comment("父级ID,0为顶级"),
		field.Enum("kind").Values("dir", "policy").Comment("分类：dir-目录、policy-权限策略"),
		field.String("name").Comment("名称"),
		field.String("comments").Optional().Comment("描述").
			Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Int("policy_id").Optional().Comment("关联的应用策略"),
		field.Text("path").Optional().Comment("路径编码").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Int32("display_sort").Optional().
			Annotations(entgql.OrderField("displaySort"), entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the AppPolicyView.
func (AppPolicyView) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("app", App.Type).Ref("policy_views").Unique().Immutable().Field("app_id"),
		edge.From("app_policy", AppPolicy.Type).Ref("policy_views").Unique().Field("policy_id"),
		edge.To("children", AppPolicyView.Type).
			From("parent").Unique().Required().Field("parent_id"),
	}
}

func (apv AppPolicyView) Hooks() []ent.Hook {
	return []ent.Hook{
		apv.pathHook(),
		InitDisplaySortHook(apppolicyview.Table),
	}
}

func (apv AppPolicyView) pathHook() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.AppPolicyViewFunc(func(ctx context.Context, mutation *gen.AppPolicyViewMutation) (gen.Value, error) {
				if _, ok := mutation.Path(); ok {
					return next.Mutate(ctx, mutation)
				}
				if pid, ok := mutation.ParentID(); ok {
					id, _ := mutation.ID()
					path := strconv.FormatInt(int64(id), 36)
					if pid == 0 {
						mutation.SetPath(path)
					} else {
						parentPath := ""
						prow, err := mutation.Client().AppPolicyView.Query().Where(apppolicyview.ID(pid)).
							Select(apppolicyview.FieldPath).Only(ctx)
						if err != nil {
							if !gen.IsNotFound(err) {
								return nil, err
							}
							parentPath = ""
						} else {
							parentPath = prow.Path + "/"
						}
						mutation.SetPath(parentPath + path)
					}
				}
				return next.Mutate(ctx, mutation)
			})
		}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}
