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
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/appdictitem"
	"github.com/woocoos/knockout/ent/hook"
)

// AppDictItem holds the schema definition for the AppDictItem entity.
type AppDictItem struct {
	ent.Schema
}

func (AppDictItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_dict_item"},
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (AppDictItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		schemax.IntID{},
		schemax.AuditMixin{},
	}
}

// Fields of the AppDictItem.
func (AppDictItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("app_id").Optional().Immutable().Comment("所属应用"),
		field.Int("org_id").Optional().Immutable().Comment("组织ID"),
		field.Int("dict_id").Optional().Immutable().Comment("所属字典"),
		field.String("code").MinLen(3).MaxLen(20).Immutable().
			Comment("用于标识应用资源的唯一代码,尽量简短"),
		field.String("name").MaxLen(45).Comment("名称"),
		field.String("comments").Optional().Comment("备注").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Int32("display_sort").Optional().Annotations(entgql.OrderField("displaySort"),
			entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the AppDictItem.
func (AppDictItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dict", AppDict.Type).Ref("items").Unique().Immutable().Field("dict_id"),
	}
}

func (AppDictItem) Hooks() []ent.Hook {
	return []ent.Hook{
		InitDisplaySortHookEx(appdictitem.Table, appdictitem.FieldDictID),
		// check code unique
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.AppDictItemFunc(func(ctx context.Context, m *gen.AppDictItemMutation) (gen.Value, error) {
				code, cok := m.Code()
				id, iok := m.ID()
				dictid, dok := m.DictID()
				if cok && iok && dok {
					return next.Mutate(ctx, m)
				}
				has, err := m.Client().AppDictItem.Query().Where(
					appdictitem.DictID(dictid), appdictitem.Code(code), appdictitem.IDNEQ(id)).Exist(ctx)
				if err != nil {
					return nil, err
				}
				if has {
					return nil, fmt.Errorf("code exists:%s", code)
				}
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate|ent.OpUpdate),
	}
}
