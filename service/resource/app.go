package resource

import (
	"context"
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
	"strings"
)

// CreateApp 创建应用,默认创建的应用都为公开的,不需要审核
//
// TODO 应用工作流
func (s *Service) CreateApp(ctx context.Context, input ent.CreateAppInput) (*ent.App, error) {
	client := ent.FromContext(ctx)
	tid := identity.TenantIDFromContext(ctx)
	apl := client.App.Create().SetInput(input).SetOrgID(tid).SetPrivate(false).SaveX(ctx)
	return apl, nil
}

// CreateAppActions
func (s *Service) CreateAppActions(ctx context.Context, appID int, input []*ent.CreateAppActionInput) ([]*ent.AppAction, error) {
	client := ent.FromContext(ctx)
	tid := identity.TenantIDFromContext(ctx)
	has := client.App.Query().Where(app.ID(appID), app.OrgID(tid)).ExistX(ctx)
	if !has {
		return nil, fmt.Errorf("app not exist")
	}
	builders := make([]*ent.AppActionCreate, len(input))
	for i := range input {
		builders[i] = client.AppAction.Create().SetInput(*input[i]).SetAppID(appID)
	}
	return client.AppAction.CreateBulk(builders...).Save(ctx)
}

// CreateAppMenus
func (s *Service) CreateAppMenus(ctx context.Context, appID int, input []*ent.CreateAppMenuInput) ([]*ent.AppMenu, error) {
	client := ent.FromContext(ctx)
	tid := identity.TenantIDFromContext(ctx)
	has := client.App.Query().Where(app.ID(appID), app.OrgID(tid)).ExistX(ctx)
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
	}
	ams, err := client.AppMenu.CreateBulk(builders...).Save(ctx)
	return ams, err
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

// CreateAppPolicies 创建应用策略.
//
// 该方法会检查应用策略的规则中的action是否以应用代码开头.
func (s *Service) CreateAppPolicies(ctx context.Context, appID int, input []*ent.CreateAppPolicyInput) ([]*ent.AppPolicy, error) {
	client := ent.FromContext(ctx)
	tid := identity.TenantIDFromContext(ctx)
	apl, err := client.App.Query().Where(app.ID(appID), app.OrgID(tid)).Only(ctx)
	if err != nil {
		return nil, err
	}

	builders := make([]*ent.AppPolicyCreate, len(input))
	for i, data := range input {
		for _, rule := range data.Rules {
			for _, action := range rule.Actions {
				if action == "*" {
					continue
				}
				if !strings.HasPrefix(action, apl.Code+":") {
					return nil, fmt.Errorf("action %s must start with app code %s", action, apl.Code)
				}
			}
		}
		builders[i] = client.AppPolicy.Create().SetInput(*input[i]).SetAppID(appID)
	}
	return client.AppPolicy.CreateBulk(builders...).Save(ctx)
}

func (s *Service) UpdateAppPolicy(ctx context.Context, policyID int, input ent.UpdateAppPolicyInput) (*ent.AppPolicy, error) {
	client := ent.FromContext(ctx)
	tid := identity.TenantIDFromContext(ctx)
	apl, err := client.AppPolicy.Query().WithApp(func(query *ent.AppQuery) {
		query.Where(app.OrgID(tid))
	}).Where(apppolicy.ID(policyID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	appcode := apl.Edges.App.Code
	for _, rule := range input.Rules {
		for _, action := range rule.Actions {
			if action == "*" {
				continue
			}
			if !strings.HasPrefix(action, appcode+":") {
				return nil, fmt.Errorf("action %s must start with app code %s", action, appcode)
			}
		}
	}

	return client.AppPolicy.UpdateOneID(policyID).SetInput(input).Save(ctx)
}

// UpdateApp 更新应用
func (s *Service) UpdateApp(ctx context.Context, appID int, input ent.UpdateAppInput) (*ent.App, error) {
	client := ent.FromContext(ctx)
	tid := identity.TenantIDFromContext(ctx)
	has, err := client.App.Query().Where(app.OrgID(tid), app.ID(appID)).Exist(ctx)
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
	tid := identity.TenantIDFromContext(ctx)
	has, err := client.AppRole.Query().Where(approle.ID(roleID), approle.HasAppWith(app.OrgID(tid))).Exist(ctx)
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
	tid := identity.TenantIDFromContext(ctx)
	has, err := client.AppRole.Query().Where(approle.ID(roleID), approle.HasAppWith(app.OrgID(tid))).Exist(ctx)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("role not exist")
	}
	return client.AppRole.DeleteOneID(roleID).Exec(ctx)
}
