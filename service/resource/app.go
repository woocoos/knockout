package resource

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"fmt"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/knockout/api/graphql/model"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/approle"
	"github.com/woocoos/knockout/ent/approlepolicy"
	"github.com/woocoos/knockout/ent/orgapp"
	"github.com/woocoos/knockout/ent/orgpolicy"
)

// CreateApp 创建应用,默认创建的应用都为公开的,不需要审核
//
// TODO 应用工作流
func (s *Service) CreateApp(ctx context.Context, input ent.CreateAppInput) (*ent.App, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	apl := client.App.Create().SetInput(input).SetOwnerOrgID(tid).SetPrivate(false).SaveX(ctx)
	return apl, nil
}

// CreateAppActions
func (s *Service) CreateAppActions(ctx context.Context, appID int, input []*ent.CreateAppActionInput) ([]*ent.AppAction, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	has := client.App.Query().Where(app.ID(appID), app.OwnerOrgID(tid)).ExistX(ctx)
	if !has {
		return nil, fmt.Errorf("app not exist")
	}
	builders := make([]*ent.AppActionCreate, len(input))
	names := make([]string, len(input))
	for i := range input {
		builders[i] = client.AppAction.Create().SetInput(*input[i]).SetAppID(appID)
		names[i] = input[i].Name
	}
	existNames, err := client.AppAction.Query().Where(appaction.NameIn(names...), appaction.AppID(appID)).Select(appaction.FieldName).Strings(ctx)
	if err != nil {
		return nil, err
	}
	if len(existNames) > 0 {
		return nil, fmt.Errorf("action %s is exist", existNames[0])
	}
	return client.AppAction.CreateBulk(builders...).Save(ctx)
}

// UpdateAppAction 更新action时，同步更新app_policy与org_policy引用的action
func (s *Service) UpdateAppAction(ctx context.Context, actionID int, input ent.UpdateAppActionInput) (*ent.AppAction, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	aa, err := client.AppAction.Query().WithApp(func(query *ent.AppQuery) {
		query.Where(app.OwnerOrgID(tid)).Select(app.FieldID, app.FieldCode)
	}).Where(appaction.ID(actionID)).Select(appaction.FieldName, appaction.FieldAppID).Only(ctx)
	if err != nil {
		return nil, err
	}
	if aa == nil {
		return nil, fmt.Errorf("action not exist")
	}
	// 判断应用下action name唯一
	exist, err := client.AppAction.Query().Where(appaction.Name(*input.Name), appaction.AppID(aa.AppID)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("action %s is exist", *input.Name)
	}
	//
	resaa, err := client.AppAction.UpdateOneID(actionID).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	// Name更新需同步更新policy中的引用
	if aa.Name != *input.Name {
		appid := aa.Edges.App.ID

		// 更新AppPolicy
		aps, err := client.AppPolicy.Query().Where(apppolicy.AppID(appid), func(selector *sql.Selector) {
			selector.Where(sqljson.StringContains(apppolicy.FieldRules, "\""+aa.Name+"\""))
		}).Select(apppolicy.FieldID, apppolicy.FieldRules).All(ctx)
		if err != nil {
			return nil, err
		}
		for _, policy := range aps {
			prs := policy.Rules
			if prs == nil {
				continue
			}
			for _, rule := range prs {
				if rule.Actions == nil {
					continue
				}
				rule.Actions = UpdateSliceElement[string](rule.Actions, *input.Name, aa.Name)
			}
			err = client.AppPolicy.UpdateOneID(policy.ID).SetRules(prs).Exec(ctx)
			if err != nil {
				return nil, err
			}
		}

		oac := aa.Edges.App.Code + ":" + aa.Name
		nac := aa.Edges.App.Code + ":" + *input.Name
		// 更新OrgPolicy
		ops, err := client.OrgPolicy.Query().Where(func(selector *sql.Selector) {
			selector.Where(sqljson.StringContains(orgpolicy.FieldRules, "\""+oac+"\""))
		}).Select(orgpolicy.FieldID, orgpolicy.FieldRules).All(ctx)
		if err != nil {
			return nil, err
		}
		for _, policy := range ops {
			prs := policy.Rules
			if prs == nil {
				continue
			}
			for _, rule := range prs {
				if rule.Actions == nil {
					continue
				}
				rule.Actions = UpdateSliceElement[string](rule.Actions, nac, oac)
			}

			// rules不为空，则同步修改casbin授权信息
			if prs != nil {
				err := updateOrgPolicyRules(ctx, policy.ID, prs, tid)
				if err != nil {
					return nil, err
				}
			}

			err = client.OrgPolicy.UpdateOneID(policy.ID).SetRules(prs).Exec(ctx)
			if err != nil {
				return nil, err
			}
		}
	}
	return resaa, nil
}

// DeleteAppAction 删除action时，同步删除app_policy与org_policy引用的action
func (s *Service) DeleteAppAction(ctx context.Context, actionID int) error {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	aa, err := client.AppAction.Query().WithApp(func(query *ent.AppQuery) {
		query.Where(app.OwnerOrgID(tid)).Select(app.FieldID, app.FieldCode)
	}).Where(appaction.ID(actionID)).Select(appaction.FieldName).Only(ctx)
	if err != nil {
		return err
	}
	if aa == nil {
		return fmt.Errorf("action not exist")
	}

	appid := aa.Edges.App.ID

	// 更新AppPolicy
	aps, err := client.AppPolicy.Query().Where(apppolicy.AppID(appid), func(selector *sql.Selector) {
		selector.Where(sqljson.StringContains(apppolicy.FieldRules, "\""+aa.Name+"\""))
	}).Select(apppolicy.FieldID, apppolicy.FieldRules).All(ctx)
	if err != nil {
		return err
	}
	for _, policy := range aps {
		prs := policy.Rules
		if prs == nil {
			continue
		}
		for _, rule := range prs {
			if rule.Actions == nil {
				continue
			}
			rule.Actions = RemoveSliceElement[string](rule.Actions, aa.Name)
		}
		err = client.AppPolicy.UpdateOneID(policy.ID).SetRules(prs).Exec(ctx)
		if err != nil {
			return err
		}
	}

	aac := aa.Edges.App.Code + ":" + aa.Name
	// 更新OrgPolicy
	ops, err := client.OrgPolicy.Query().Where(func(selector *sql.Selector) {
		selector.Where(sqljson.StringContains(orgpolicy.FieldRules, "\""+aac+"\""))
	}).Select(orgpolicy.FieldID, orgpolicy.FieldRules).All(ctx)
	if err != nil {
		return err
	}
	for _, policy := range ops {
		prs := policy.Rules
		if prs == nil {
			continue
		}
		for _, rule := range prs {
			if rule.Actions == nil {
				continue
			}
			rule.Actions = RemoveSliceElement[string](rule.Actions, aac)
		}

		// rules不为空，则同步修改casbin授权信息
		if prs != nil {
			err = updateOrgPolicyRules(ctx, policy.ID, prs, tid)
			if err != nil {
				return err
			}
		}

		err = client.OrgPolicy.UpdateOneID(policy.ID).SetRules(prs).Exec(ctx)
		if err != nil {
			return err
		}
	}
	client.AppAction.Delete().Where(appaction.ID(actionID)).ExecX(ctx)
	return nil
}

// CreateAppMenus 创建应用菜单，如果有route项，则相应创建action
func (s *Service) CreateAppMenus(ctx context.Context, appID int, input []*ent.CreateAppMenuInput) ([]*ent.AppMenu, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	has := client.App.Query().Where(app.ID(appID), app.OwnerOrgID(tid)).ExistX(ctx)
	if !has {
		return nil, fmt.Errorf("app not exist")
	}
	builders := make([]*ent.AppMenuCreate, len(input))
	for i, menu := range input {
		if menu.ActionID != nil {
			has := client.AppAction.Query().Where(appaction.ID(*menu.ActionID), appaction.AppID(appID)).ExistX(ctx)
			if !has {
				return nil, fmt.Errorf("app action not exist")
			}
		}
		builders[i] = client.AppMenu.Create().SetInput(*menu).SetAppID(appID)
		// 创建action
		if menu.Route != nil {
			aac := menu.Comments
			if aac == nil {
				aac = &menu.Name
			}
			aa, err := client.AppAction.Create().SetAppID(appID).SetName(*menu.Route).SetKind(appaction.KindRestful).SetMethod(appaction.MethodRead).SetComments(*aac).Save(ctx)
			if err != nil {
				return nil, err
			}
			builders[i].SetActionID(aa.ID)
		}
	}
	ams, err := client.AppMenu.CreateBulk(builders...).Save(ctx)
	return ams, err
}

// UpdateAppMenu 更新应用菜单，如果更新了route，则更新action
func (s *Service) UpdateAppMenu(ctx context.Context, menuID int, input ent.UpdateAppMenuInput) (*ent.AppMenu, error) {
	client := ent.FromContext(ctx)
	am, err := client.AppMenu.Query().Where(appmenu.ID(menuID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	// 更新route更新关联的action name
	var build *ent.AppMenuUpdateOne
	build = client.AppMenu.UpdateOneID(menuID).SetInput(input)
	if input.Route != nil {
		if am.ActionID != nil {
			if *input.Route == "" {
				err = s.DeleteAppAction(ctx, *am.ActionID)
				if err != nil {
					return nil, err
				}
			} else {
				_, err = s.UpdateAppAction(ctx, *am.ActionID, ent.UpdateAppActionInput{Name: input.Route})
				if err != nil {
					return nil, err
				}
			}
		} else {
			aac := &am.Comments
			if aac == nil {
				aac = &am.Name
			}
			aas, err := s.CreateAppActions(ctx, am.AppID, []*ent.CreateAppActionInput{
				{Name: *input.Route, Comments: aac, Kind: appaction.KindRestful, Method: appaction.MethodRead},
			})
			if err != nil {
				return nil, err
			}
			build.SetActionID(aas[0].ID)
		}
	}
	return build.Save(ctx)
}

// DeleteAppMenu 删除应用菜单，删除关联的action
func (s *Service) DeleteAppMenu(ctx context.Context, menuID int) error {
	client := ent.FromContext(ctx)
	am, err := client.AppMenu.Query().Where(appmenu.ID(menuID)).Only(ctx)
	if err != nil {
		return err
	}
	if am.ActionID != nil {
		// 删除action
		err = s.DeleteAppAction(ctx, *am.ActionID)
		if err != nil {
			return err
		}
	}
	return client.AppMenu.DeleteOneID(menuID).Exec(ctx)
}

// MoveAppMenu
func (s *Service) MoveAppMenu(ctx context.Context, src int, tar int, action model.TreeAction) (err error) {
	client := ent.FromContext(ctx)
	tarOrg := client.AppMenu.GetX(ctx, tar)
	builder := client.AppMenu.UpdateOneID(src)
	var start int32 = 0
	var resort = true
	switch action {
	case model.TreeActionChild:
		var agg []struct {
			Max *int32
		}
		err = client.AppMenu.Query().Where(appmenu.ParentID(tarOrg.ID)).Aggregate(ent.Max(appmenu.FieldDisplaySort)).Scan(ctx, &agg)
		if err != nil {
			return err
		}
		if agg[0].Max == nil {
			start = 1
		} else {
			start = *agg[0].Max + 1
		}
		builder.SetParentID(tarOrg.ID)
		resort = false
	case model.TreeActionUp:
		start = tarOrg.DisplaySort
		builder.SetParentID(tarOrg.ParentID).SetDisplaySort(start)
	case model.TreeActionDown:
		start = tarOrg.DisplaySort + 1
		builder.SetParentID(tarOrg.ParentID).SetDisplaySort(start)
	}
	if resort {
		err = client.AppMenu.Update().Where(appmenu.ParentID(tarOrg.ParentID), appmenu.DisplaySortGTE(start)).AddDisplaySort(1).Exec(ctx)
		if err != nil {
			return
		}
	}

	return builder.Exec(ctx)
}

// UpdateApp 更新应用
func (s *Service) UpdateApp(ctx context.Context, appID int, input ent.UpdateAppInput) (*ent.App, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	has, err := client.App.Query().Where(app.OwnerOrgID(tid), app.ID(appID)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("app not exist")
	}
	return client.App.UpdateOneID(appID).SetInput(input).Save(ctx)
}

// DeleteApp 删除应用
//
// 该应用没有被关联才可以删除: 1. 没有组织目录关联 2. 没有用户关联
func (s *Service) DeleteApp(ctx context.Context, appID int) error {
	client := ent.FromContext(ctx)
	apl := client.App.GetX(ctx, appID)
	if apl.Private != true {
		has, err := client.OrgApp.Query().Where(orgapp.HasAppWith(app.ID(appID))).Exist(ctx)
		if err != nil {
			return err
		}
		if has {
			return fmt.Errorf("app has been associated with org")
		}
	} else {
		client.OrgApp.Delete().Where(orgapp.AppID(appID)).ExecX(ctx)
	}
	client.AppAction.Delete().Where(appaction.AppID(appID)).ExecX(ctx)
	client.AppMenu.Delete().Where(appmenu.AppID(appID)).ExecX(ctx)
	client.AppPolicy.Delete().Where(apppolicy.AppID(appID)).ExecX(ctx)
	client.AppRole.Delete().Where(approle.AppID(appID)).ExecX(ctx)
	client.AppRolePolicy.Delete().Where(approlepolicy.AppID(appID)).ExecX(ctx)
	return nil
}

func (s *Service) UpdateAppRole(ctx context.Context, roleID int, input ent.UpdateAppRoleInput) (*ent.AppRole, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	has, err := client.AppRole.Query().Where(approle.ID(roleID), approle.HasAppWith(app.OwnerOrgID(tid))).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("role not exist")
	}
	return ent.FromContext(ctx).AppRole.UpdateOneID(roleID).SetInput(input).Save(ctx)
}

func (s *Service) DeleteAppRole(ctx context.Context, roleID int) error {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	has, err := client.AppRole.Query().Where(approle.ID(roleID), approle.HasAppWith(app.OwnerOrgID(tid))).Exist(ctx)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("role not exist")
	}
	return client.AppRole.DeleteOneID(roleID).Exec(ctx)
}

// AssignAppRolePolicy 角色添加权限
func (s *Service) AssignAppRolePolicy(ctx context.Context, appID int, roleID int, policyIDs []int) error {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	has, err := client.AppRole.Query().Where(approle.ID(roleID), approle.AppID(appID), approle.HasAppWith(app.OwnerOrgID(tid))).Exist(ctx)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("role not exist")
	}
	count, err := client.AppPolicy.Query().Where(apppolicy.IDIn(policyIDs...), apppolicy.AppID(appID)).Count(ctx)
	if err != nil {
		return err
	}
	if count != len(policyIDs) {
		return fmt.Errorf("invalid policy in policyIDs")
	}
	builders := make([]*ent.AppRolePolicyCreate, len(policyIDs))
	for i, v := range policyIDs {
		builders[i] = client.AppRolePolicy.Create().SetAppID(appID).SetRoleID(roleID).SetPolicyID(v)
	}
	return client.AppRolePolicy.CreateBulk(builders...).Exec(ctx)
}

// CreateAppPolicy 创建应用策略.
//
// 该方法会检查应用策略的规则中的action是否以应用代码开头.
func (s *Service) CreateAppPolicy(ctx context.Context, appID int, input ent.CreateAppPolicyInput) (*ent.AppPolicy, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	exist, err := client.App.Query().Where(app.ID(appID), app.OwnerOrgID(tid)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("app not exist")
	}

	return client.AppPolicy.Create().SetAppID(appID).SetInput(input).Save(ctx)
}

// UpdateAppPolicy 更新应用策略,该应用必须属于(创建者)该租户才可更新
// 当应用策略更新时,会被当前最新的策略模板,原有引用该策略的都保持不变
func (s *Service) UpdateAppPolicy(ctx context.Context, policyID int, input ent.UpdateAppPolicyInput) (*ent.AppPolicy, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	apl, err := client.AppPolicy.Query().WithApp(func(query *ent.AppQuery) {
		query.Where(app.OwnerOrgID(tid))
	}).Where(apppolicy.ID(policyID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	builder := client.AppPolicy.UpdateOneID(policyID).SetInput(input)
	builder.Mutation().SetAppID(apl.AppID)
	return builder.Save(ctx)
}

// DeleteAppPolicy 删除应用策略,该应用必须属于(创建者)该租户才可删除
// 当应用策略被删除时,原有引用该策略的都保持不变
func (s *Service) DeleteAppPolicy(ctx context.Context, policyID int) error {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	exists, err := client.AppPolicy.Query().WithApp(func(query *ent.AppQuery) {
		query.Where(app.OwnerOrgID(tid))
	}).Where(apppolicy.ID(policyID)).Exist(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("policy not exist")
	}
	return client.AppPolicy.DeleteOneID(policyID).Exec(ctx)
}
