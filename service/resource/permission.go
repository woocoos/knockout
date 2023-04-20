package resource

import (
	"context"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/api/graphql/model"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/approle"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgapp"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orgroleuser"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/security"
	"strconv"
	"time"
)

// AssignOrganizationApp 分配应用到根组织下. 如: 新账户创建时, 根账户分配已有应用给子账户(需要验证根用户是否该应用权限,可在外层验证).
func (s *Service) AssignOrganizationApp(ctx context.Context, orgID int, appID int) error {
	client := ent.FromContext(ctx)
	has, err := client.Org.Query().Where(org.ID(orgID), org.KindEQ(org.KindRoot),
		org.HasAppsWith(app.ID(appID))).Exist(ctx)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("org not found or already has app")
	}
	ap, err := client.App.Query().Where(app.ID(appID)).
		WithPolicies(func(query *ent.AppPolicyQuery) {
			query.Where(apppolicy.AutoGrant(true)).Unique(true).
				WithRoles(func(query *ent.AppRoleQuery) {
					query.Where(approle.AutoGrant(true))
				})
		}).Only(ctx)
	if err != nil {
		return err
	}
	err = client.OrgApp.Create().SetOrgID(orgID).SetAppID(ap.ID).Exec(ctx)
	if err != nil {
		return err
	}
	// 创建应用策略
	ps, err := ap.Policies(ctx)
	if err != nil {
		return err
	}
	// 角色
	rs, err := ap.Roles(ctx)
	if err != nil {
		return err
	}

	pbk := make([]*ent.OrgPolicyCreate, len(ps))
	for i, p := range ps {
		pbk[i] = client.OrgPolicy.Create().SetOrgID(orgID).SetAppID(ap.ID).SetAppPolicyID(p.ID).
			SetComments(p.Comments).SetRules(p.Rules).SetComments(p.Comments).SetName(p.Name)
	}
	err = client.OrgPolicy.CreateBulk(pbk...).Exec(ctx)
	if err != nil {
		return err
	}
	rbk := make([]*ent.OrgRoleCreate, len(rs))
	for i, r := range rs {
		rbk[i] = client.OrgRole.Create().SetOrgID(orgID).SetKind(orgrole.KindRole).SetAppRoleID(r.ID).
			SetComments(r.Comments).SetName(r.Name)
	}
	err = client.OrgRole.CreateBulk(rbk...).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// RevokeOrganizationApp 移除组织下的应用,同时物理删除授权信息与收回授权.
func (s *Service) RevokeOrganizationApp(ctx context.Context, orgID int, appID int) error {
	client := ent.FromContext(ctx)
	orgr, err := client.Org.Query().Where(org.ID(orgID), org.KindEQ(org.KindRoot)).Only(ctx)
	if err != nil {
		return err
	}
	ps, err := client.Permission.Query().Where(
		permission.HasOrgPolicyWith(orgpolicy.OrgID(orgID), orgpolicy.AppID(appID))).
		WithOrgPolicy().
		All(ctx)
	if err != nil {
		return err
	}

	domain, err := s.GetOrgDomain(ctx, orgID)
	if err != nil {
		return err
	}
	pids := make([]int, len(ps))
	for i, p := range ps {
		pids[i] = p.ID
		pid := getPrincipalID(p)
		err = security.RevokePolicy(p.Edges.OrgPolicy.Rules, strconv.Itoa(pid), domain, p.PrincipalKind)
		if err != nil {
			log.Error(err)
		}
	}
	_, err = client.Permission.Delete().Where(permission.IDIn(pids...)).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.OrgPolicy.Delete().Where(orgpolicy.OrgID(orgr.ID),
		orgpolicy.AppID(appID)).Exec(ctx)
	if err != nil {
		return err
	}
	rids, err := client.AppRole.Query().Where(approle.AppID(appID)).IDs(ctx)
	if err != nil {
		return err
	}
	if len(rids) > 0 {
		_, err = client.OrgRoleUser.Delete().Where(
			orgroleuser.HasOrgRoleWith(orgrole.OrgID(orgr.ID), orgrole.AppRoleIDIn(rids...))).Exec(ctx)
		if err != nil {
			return err
		}
		_, err = client.OrgRole.Delete().Where(orgrole.OrgID(orgr.ID), orgrole.AppRoleIDIn(rids...)).Exec(ctx)
		if err != nil {
			return err
		}
	}
	_, err = client.OrgApp.Delete().Where(orgapp.OrgID(orgr.ID), orgapp.AppID(appID)).Exec(ctx)
	return err
}

func (s *Service) AssignOrganizationAppPolicy(ctx context.Context, orgID int, appPolicyID int) error {
	client := ent.FromContext(ctx)
	ap, err := client.AppPolicy.Query().Where(apppolicy.ID(appPolicyID)).Only(ctx)
	if err != nil {
		return err
	}
	has, err := client.OrgApp.Query().Where(orgapp.OrgID(orgID), orgapp.AppID(ap.AppID)).Exist(ctx)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("org not found or not has app")
	}
	err = client.OrgPolicy.Create().SetOrgID(orgID).SetAppID(ap.ID).SetAppPolicyID(ap.ID).
		SetComments(ap.Comments).SetRules(ap.Rules).SetComments(ap.Comments).SetName(ap.Name).Exec(ctx)
	return err
}

// AssignRoleUser is the resolver for the assignRoleUser field.
func (s *Service) AssignRoleUser(ctx context.Context, input model.AssignRoleUserInput) error {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	ouid, err := client.OrgUser.Query().Where(orguser.OrgID(tid), orguser.UserID(input.UserID)).OnlyID(ctx)
	if err != nil {
		return err
	}

	orid, err := client.OrgRole.Query().Where(orgrole.ID(input.OrgRoleID), orgrole.OrgID(tid)).OnlyID(ctx)
	if err != nil {
		return err
	}
	has, err := client.OrgRoleUser.Query().Where(orgroleuser.OrgRoleID(orid), orgroleuser.OrgUserID(ouid)).Exist(ctx)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("user already in role")
	}
	err = client.OrgRoleUser.Create().SetOrgRoleID(input.OrgRoleID).SetOrgUserID(input.UserID).Exec(ctx)
	if err != nil {
		return err
	}
	domain, err := s.GetOrgDomain(ctx, tid)
	err = security.GrantRoleForUser(input.UserID, input.OrgRoleID, domain)
	if err != nil {
		log.Errorf("grant policy failed,policy is inactive: %w", err)
	}
	return err
}

// RevokeRoleUser is the resolver for the revokeRoleUser field.
func (s *Service) RevokeRoleUser(ctx context.Context, roleID int, userID int) error {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	has, err := client.OrgRoleUser.Query().Where(orgroleuser.HasOrgUserWith(orguser.OrgID(tid), orguser.UserID(userID)),
		orgroleuser.HasOrgRoleWith(orgrole.OrgID(tid), orgrole.ID(roleID))).Exist(ctx)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("role user not found")
	}
	domain, err := s.GetOrgDomain(ctx, tid)
	if err != nil {
		return err
	}
	err = security.RevokeGroupForUser(userID, roleID, domain)
	if err != nil {
		return err
	}
	_, err = client.OrgRoleUser.Delete().Where(orgroleuser.OrgRoleID(roleID), orgroleuser.OrgUserID(userID)).Exec(ctx)
	return err
}

func (s *Service) AssignOrganizationAppRole(ctx context.Context, orgID int, appRoleID int) error {
	panic("implement me")
}

func (s *Service) RevokeOrganizationAppPolicy(ctx context.Context, orgID int, appPolicyID int) error {
	client := ent.FromContext(ctx)

	domain, err := s.GetOrgDomain(ctx, orgID)
	if err != nil {
		return err
	}

	ps, err := client.Permission.Query().Where(
		permission.HasOrgPolicyWith(orgpolicy.OrgID(orgID), orgpolicy.AppPolicyID(appPolicyID))).
		WithOrgPolicy().
		All(ctx)
	if err != nil {
		return err
	}

	pids := make([]int, len(ps))
	for i, p := range ps {
		pids[i] = p.ID
		pid := getPrincipalID(p)
		err = security.RevokePolicy(p.Edges.OrgPolicy.Rules, strconv.Itoa(pid), domain, p.PrincipalKind)
		if err != nil {
			log.Error(err)
		}
	}
	_, err = client.Permission.Delete().Where(permission.IDIn(pids...)).Exec(ctx)
	return err
}

func getPrincipalID(perm *ent.Permission) int {
	switch perm.PrincipalKind {
	case permission.PrincipalKindUser:
		return perm.UserID
	case permission.PrincipalKindRole:
		return perm.RoleID
	}
	return 0
}

// Grant 给用户或角色授权.
//
// 此时先保证permission数据保存,如果cashbin操作失败,返回状态失败,再需要通过权限管理界面再次激活..
func (s *Service) Grant(ctx context.Context, input ent.CreatePermissionInput) (*ent.Permission, error) {
	return s.grantPolicy(ctx, input)
}

// grantPolicy 给用户或角色授权.
func (s *Service) grantPolicy(ctx context.Context, input ent.CreatePermissionInput) (*ent.Permission, error) {
	client := ent.FromContext(ctx)
	domain, err := s.GetOrgDomain(ctx, input.OrgID)
	if err != nil {
		return nil, err
	}
	pid := 0
	builder := client.Permission.Create().SetOrgID(input.OrgID).SetOrgPolicyID(input.OrgPolicyID).
		SetStatus(typex.SimpleStatusInactive)
	switch input.PrincipalKind {
	case permission.PrincipalKindUser:
		if input.UserID == nil {
			return nil, fmt.Errorf("user id is required")
		}
		pid = *input.UserID
		builder.SetUserID(pid)
		builder.SetPrincipalKind(permission.PrincipalKindUser)
	case permission.PrincipalKindRole:
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
	perm, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	policy, err := client.OrgPolicy.Query().Where(orgpolicy.ID(input.OrgPolicyID)).Select(orgpolicy.FieldRules).Only(ctx)
	if err != nil {
		return nil, err
	}
	perm.Edges.OrgPolicy = policy
	err = security.GrantByPermission(perm, domain)
	if err != nil {
		log.Errorf("grant policy failed,policy is inactive: %w", err)
		return perm, nil
	}
	perm.Status = typex.SimpleStatusActive
	return client.Permission.UpdateOneID(perm.ID).SetStatus(typex.SimpleStatusActive).Save(ctx)
}

// UpdatePermission 更新权限的
func (s *Service) UpdatePermission(ctx context.Context, permissionID int, input ent.UpdatePermissionInput) (*ent.Permission, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	domain, err := s.GetOrgDomain(ctx, tid)
	if err != nil {
		return nil, err
	}
	p, err := client.Permission.UpdateOneID(permissionID).Where(permission.ID(permissionID), permission.OrgID(tid)).SetInput(input).Save(ctx)
	if err != nil {
		return p, err
	}
	op, err := p.OrgPolicy(ctx)
	if err != nil {
		return nil, err
	}
	p.Edges.OrgPolicy = op
	// TODO 采用先清除再添加的方式
	err = security.RevokeByPermission(p, domain)
	if err != nil {
		return p, err
	}
	if p.Status == typex.SimpleStatusActive && p.EndAt.After(time.Now()) {
		err = security.GrantByPermission(p, domain)
	}
	return p, err
}

// Revoke 撤销用户或角色的权限.
func (s *Service) Revoke(ctx context.Context, orgID int, permissionID int) error {
	client := ent.FromContext(ctx)
	domain, err := s.GetOrgDomain(ctx, orgID)
	if err != nil {
		return err
	}
	p, err := client.Permission.Query().Where(permission.ID(permissionID), permission.OrgID(orgID)).Only(ctx)
	if err != nil {
		return err
	}

	pid := getPrincipalID(p)
	err = security.RevokePolicy(p.Edges.OrgPolicy.Rules, strconv.Itoa(pid), domain, p.PrincipalKind)
	if err != nil {
		log.Error(err)
	}
	_, err = client.Permission.Delete().Where(permission.ID(pid)).Exec(ctx)
	return err
}

// GetOrgDomain 获取组织域名.orgID为根组织.
func (s *Service) GetOrgDomain(ctx context.Context, orgID int) (string, error) {
	c := ent.FromContext(ctx)
	orgr := c.Org.Query().Where(org.ID(orgID)).Select(org.FieldDomain).OnlyX(ctx)
	if orgr.Domain == "" {
		return "", fmt.Errorf("organization %d domain is empty", orgID)
	}
	return orgr.Domain, nil
}

// GetRootOrgByUser 获取用户的最顶级的根组织.在组织中,一个账户可能存在多个根组织.需要从context获取租户ID
func (s *Service) GetRootOrgByUser(ctx context.Context, uid int) (*ent.Org, error) {
	c := ent.FromContext(ctx)
	return c.Org.Query().Where(org.HasUsersWith(user.ID(uid)), org.KindEQ(org.KindRoot)).
		Order(ent.Asc(org.FieldPath)).First(ctx)
}
