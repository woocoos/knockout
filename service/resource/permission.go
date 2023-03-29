package resource

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/approle"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/security"
	"strconv"
	"strings"
)

// CreateAppPolicies 创建应用策略.
//
// 该方法会检查应用策略的规则中的action是否以应用代码开头.
func (s *Service) CreateAppPolicies(ctx context.Context, input []*ent.CreateAppPolicyInput) error {
	c := ent.FromContext(ctx)
	builders := make([]*ent.AppPolicyCreate, len(input))
	for i, data := range input {
		apl, err := s.Client.App.Query().Where(app.ID(data.AppID)).Only(ctx)
		if err != nil {
			return err
		}
		for _, rule := range data.Rules {
			for _, action := range rule.Actions {
				if action == "*" {
					continue
				}
				if !strings.HasPrefix(action, apl.Code+":") {
					return fmt.Errorf("action %s must start with app code %s", action, apl.Code)
				}
			}
		}
		builders[i] = c.AppPolicy.Create().SetInput(*input[i])
	}
	err := c.AppPolicy.CreateBulk(builders...).Exec(ctx)
	return err
}

// AssignOrganizationApp 分配应用到根组织下. 如: 新账户创建时, 根账户分配已有应用给子账户(需要验证根用户是否该应用权限,可在外层验证).
func (s *Service) AssignOrganizationApp(ctx context.Context, orgID int, appID int) (bool, error) {
	c := ent.FromContext(ctx)
	org := c.Org.Query().Where(org.ID(orgID), org.KindEQ(org.KindRoot),
		org.Not(org.HasAppsWith(func(selector *sql.Selector) {
			selector.Where(sql.EQ(app.FieldID, appID))
		}))).OnlyX(ctx)
	ap := c.App.Query().Where(app.ID(appID)).
		WithPolicies(func(query *ent.AppPolicyQuery) {
			query.Where(apppolicy.AutoGrant(true)).Unique(true).
				WithRoles(func(query *ent.AppRoleQuery) {
					query.Where(approle.AutoGrant(true))
				})
		}).OnlyX(ctx)
	// 创建应用策略
	ps, err := ap.Policies(ctx)
	if err != nil {
		return false, err
	}
	// 角色
	rs, err := ap.Roles(ctx)
	if err != nil {
		return false, err
	}

	pbk := make([]*ent.OrgPolicyCreate, len(ps))
	for i, p := range ps {
		pbk[i] = c.OrgPolicy.Create().SetOrgID(org.ID).SetAppID(ap.ID).SetAppPolicyID(p.ID).
			SetComments(p.Comments).SetRules(p.Rules).SetComments(p.Comments).SetName(p.Name)
	}
	err = c.OrgPolicy.CreateBulk(pbk...).Exec(ctx)
	if err != nil {
		return false, err
	}
	rbk := make([]*ent.OrgRoleCreate, len(rs))
	for i, r := range rs {
		rbk[i] = c.OrgRole.Create().SetOrgID(org.ID).SetKind(orgrole.KindRole).SetAppRoleID(r.ID).
			SetComments(r.Comments).SetName(r.Name)
	}
	err = c.OrgRole.CreateBulk(rbk...).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// assignOrganizationPolicy 分配应用权限策略到组织.
func (s *Service) assignAppPolicyToOrganization(ctx context.Context, orgID int, policyID int) (*ent.OrgPolicy, error) {
	c := ent.FromContext(ctx)
	pm := c.AppPolicy.Query().Where(apppolicy.ID(policyID)).OnlyX(ctx)
	// 如应用存在,确定应用是否分配入组织
	// 对组织授权
	return c.OrgPolicy.Create().SetOrgID(orgID).SetAppID(pm.AppID).SetAppPolicyID(pm.ID).
		SetComments(pm.Comments).
		Save(ctx)
}

// Grant 给用户或角色授权.
//
// 此时先保证permission数据保存,如果cashbin操作失败,返回状态失败,再需要通过权限管理界面再次激活..
func (s *Service) Grant(ctx context.Context, input ent.CreatePermissionInput) (*ent.Permission, error) {
	c := ent.FromContext(ctx)
	domain, err := s.GetOrgDomain(ctx, input.OrgID)
	if err != nil {
		return nil, err
	}
	pid := 0
	builder := c.Permission.Create().SetOrgID(input.OrgID).SetOrgPolicyID(input.OrgPolicyID).
		SetStatus(typex.SimpleStatusInactive)
	switch input.PrincipalKind {
	case "user":
		if input.UserID == nil {
			return nil, fmt.Errorf("user id is required")
		}
		pid = *input.UserID
		builder.SetUserID(pid)
		builder.SetPrincipalKind(permission.PrincipalKindUser)
	case "role":
		if input.RoleID == nil {
			return nil, fmt.Errorf("role id is required")
		}
		pid = *input.RoleID
		builder.SetRoleID(pid)
		builder.SetPrincipalKind(permission.PrincipalKindRole)
	default:
		return nil, fmt.Errorf("grant type %s not support", input.PrincipalKind)
	}
	// save first
	row, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	policy := c.OrgPolicy.Query().Where(orgpolicy.ID(input.OrgPolicyID)).Select(orgpolicy.FieldRules).OnlyX(ctx)
	err = security.GrantPolicy(policy.Rules, strconv.Itoa(pid), input.PrincipalKind.String(), domain)
	if err != nil {
		log.Errorf("grant policy failed,policy is inactive: %w", err)
		return row, nil
	}
	row.Status = typex.SimpleStatusActive
	return c.Permission.UpdateOneID(row.ID).SetStatus(typex.SimpleStatusActive).Save(ctx)
}

// GetOrgDomain 获取组织域名.orgID为根组织.
func (s *Service) GetOrgDomain(ctx context.Context, orgID int) (string, error) {
	c := ent.FromContext(ctx)
	org := c.Org.Query().Where(org.ID(orgID)).Select(org.FieldDomain).OnlyX(ctx)
	if org.Domain == "" {
		return "", fmt.Errorf("organization %d domain is empty", orgID)
	}
	return org.Domain, nil
}

// GetRootOrgByUser 获取用户的最顶级的根组织.在组织中,一个账户可能存在多个根组织.需要从context获取租户ID
func (s *Service) GetRootOrgByUser(ctx context.Context, uid int) (*ent.Org, error) {
	c := ent.FromContext(ctx)
	return c.Org.Query().Where(org.HasUsersWith(user.ID(uid)), org.KindEQ(org.KindRoot)).
		Order(ent.Asc(org.FieldPath)).First(ctx)
}
