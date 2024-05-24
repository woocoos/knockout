package schema

import (
	"context"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"fmt"
	"github.com/woocoos/knockout-go/ent/schemax"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/identity"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/appdict"
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
		schemax.NotifyMixin{},
	}
}

// Fields of the AppDictItem.
func (AppDictItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("org_id").Optional().Immutable().Comment("组织ID,空为全局字典"),
		field.Int("dict_id").Optional().Immutable().Comment("所属字典"),
		field.String("ref_code").Comment("关联代码,由app_code和dict_code组成").Annotations(entgql.Skip(
			entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		),
		field.String("code").MinLen(3).MaxLen(20).Immutable().
			Comment("字典值唯一编码,生效后不可修改."),
		field.String("name").MaxLen(45).Comment("名称"),
		field.String("comments").Optional().Comment("备注").Annotations(entgql.Skip(entgql.SkipWhereInput)),
		field.Int32("display_sort").Optional().Annotations(entgql.OrderField("displaySort"),
			entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Enum("status").GoType(typex.SimpleStatus("")).Default(typex.SimpleStatusInactive.String()).
			Optional().Comment("状态"),
	}
}

// Edges of the AppDictItem.
func (AppDictItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dict", AppDict.Type).Ref("items").Unique().Immutable().Field("dict_id"),
		edge.To("org", Org.Type).Unique().Immutable().Field("org_id"),
	}
}

func (AppDictItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("ref_code"),
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
		}, ent.OpCreate|ent.OpUpdateOne),
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.AppDictItemFunc(func(ctx context.Context, m *gen.AppDictItemMutation) (gen.Value, error) {
				dict, _ := m.DictID()
				dr, err := m.Client().AppDict.Query().Where(appdict.ID(dict)).WithApp().Only(ctx)
				if err != nil {
					return nil, err
				}
				app, err := dr.App(ctx)
				if err != nil {
					return nil, err
				}
				m.SetRefCode(fmt.Sprintf("%s:%s", app.Code, dr.Code))
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.AppDictItemFunc(func(ctx context.Context, m *gen.AppDictItemMutation) (gen.Value, error) {
				tid, err := identity.TenantIDFromContext(ctx)
				if err != nil {
					return nil, err
				}
				if m.Op().Is(ent.OpCreate) {
					org, _ := m.OrgID()
					if org != 0 {
						m.SetOrgID(tid)
					} else {
						// TODO check pemission
					}
				}
				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate|ent.OpUpdateOne),
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.AppDictItemFunc(func(ctx context.Context, m *gen.AppDictItemMutation) (gen.Value, error) {
				id, _ := m.ID()
				status, ok := m.Status()
				if !ok {
					row, err := m.Client().AppDictItem.Get(ctx, id)
					if err != nil {
						return nil, err
					}
					status = row.Status
				}
				if status != typex.SimpleStatusInactive {
					return nil, fmt.Errorf("can't not delete not active stauts")
				}
				return next.Mutate(ctx, m)
			})
		}, ent.OpDeleteOne),
	}
}
