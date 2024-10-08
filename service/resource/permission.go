package resource

import (
	"context"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/authz"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout/api/graphql/model"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/approle"
	"github.com/woocoos/knockout/ent/approlepolicy"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgapp"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orgroleuser"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/predicate"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/security"
	"strconv"
	"strings"
	"time"
)

const ArnSplit = authz.ArnSplit
const SplitPolicyEffect = "&&"

// AssignOrganizationApp 分配应用到根组织下. 如: 新账户创建时, 根账户分配已有应用给子账户(需要验证根用户是否该应用权限,可在外层验证).
func (s *Service) AssignOrganizationApp(ctx context.Context, orgID int, appID int) error {
	client := ent.FromContext(ctx)
	// 判断组织是否存在或者是否关联应用
	if has, err := client.Org.Query().Where(org.ID(orgID), org.KindEQ(org.KindRoot),
		org.HasAppsWith(app.ID(appID))).Exist(ctx); err != nil {
		return err
	} else if has {
		return fmt.Errorf("org not found or already has app")
	}
	// 判断组织是否关联根用户
	if o, err := client.Org.Query().Where(org.ID(orgID)).Only(ctx); err != nil {
		return err
	} else if o.OwnerID == nil {
		return fmt.Errorf("the organization owner was not found")
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
	// 关联app
	if err = client.OrgApp.Create().SetOrgID(orgID).SetAppID(ap.ID).Exec(ctx); err != nil {
		return err
	}

	// 应用策略
	ps, err := ap.Policies(ctx)
	if err != nil {
		return err
	}
	// 角色
	rs, err := ap.Roles(ctx)
	if err != nil {
		return err
	}

	rids := make([]int, len(rs))
	for i, r := range rs {
		rids[i] = r.ID
		err = s.AssignOrganizationAppRole(ctx, orgID, r.ID)
		if err != nil {
			return err
		}
	}
	// 角色关联的策略
	rpids, err := client.AppRolePolicy.Query().Where(approlepolicy.AppID(appID), approlepolicy.AppRoleIDIn(rids...)).Select(approlepolicy.FieldAppPolicyID).Ints(ctx)
	rps := make(map[int]bool)
	for _, pid := range rpids {
		rps[pid] = true
	}
	// 排除角色关联的策略
	nps := make([]*ent.AppPolicy, 0)
	for _, p := range ps {
		if rps[p.ID] {
			continue
		}
		nps = append(nps, p)
	}
	for _, v := range nps {
		err = s.AssignOrganizationAppPolicy(ctx, orgID, v.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// appPolicy to orgPolicy
func appPolicyToOrgPolicy(appCode string, rules []*types.PolicyRule, tenantID int) error {
	for _, rule := range rules {
		for j, action := range rule.Actions {
			rule.Actions[j] = appCode + ArnSplit + action
		}
		for j, resource := range rule.Resources {
			// 替换tenant_id
			resource = authz.ReplaceTenantID(resource, tenantID)
			// 补充appCode
			rule.Resources[j] = appCode + ArnSplit + authz.FormatResourceArn(resource)
		}
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

	isRoot, err := s.IsRootOrg(ctx, orgID)
	if err != nil {
		return err
	}
	if !isRoot {
		return fmt.Errorf("organization %d is not a root organization", orgID)
	}

	pids := make([]int, len(ps))
	for i, p := range ps {
		pids[i] = p.ID
		pid := getPrincipalID(p)
		err = security.RevokePolicy(p.Edges.OrgPolicy.Rules, strconv.Itoa(pid), orgID, p.PrincipalKind)
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
		//
		orus, err := client.OrgRoleUser.Query().Where(orgroleuser.HasOrgRoleWith(orgrole.OrgID(orgr.ID), orgrole.AppRoleIDIn(rids...))).WithOrgUser().All(ctx)
		if err != nil {
			return err
		}
		for _, oru := range orus {
			err = security.RevokeGroupForUser(oru.Edges.OrgUser.UserID, oru.OrgRoleID, orgID)
			if err != nil {
				return err
			}
		}
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

	isRoot, err := s.IsRootOrg(ctx, orgID)
	if err != nil {
		return err
	}
	if !isRoot {
		return fmt.Errorf("organization %d is not a root organization", orgID)
	}

	ap, err := client.AppPolicy.Query().Where(apppolicy.ID(appPolicyID)).WithApp().Only(ctx)
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
	has, err = client.OrgPolicy.Query().Where(orgpolicy.AppPolicyID(appPolicyID), orgpolicy.OrgID(orgID)).Exist(ctx)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("policy has assigned to org")
	}
	err = appPolicyToOrgPolicy(ap.Edges.App.Code, ap.Rules, orgID)
	if err != nil {
		return err
	}
	op, err := client.OrgPolicy.Create().SetOrgID(orgID).SetAppID(ap.AppID).SetAppPolicyID(ap.ID).
		SetComments(ap.Comments).SetRules(ap.Rules).SetComments(ap.Comments).SetName(ap.Name).Save(ctx)
	if err != nil {
		return err
	}
	// 授权应用策略给根组织owner
	err = s.assignAppPolicyToRootOwner(ctx, orgID, op.ID)
	return err
}

// assignAppPolicyToRootOwner 给根组织管理者授权策略
func (s *Service) assignAppPolicyToRootOwner(ctx context.Context, rootOrgID int, orgPolicyID int) error {
	client := ent.FromContext(ctx)
	rootOrg, err := client.Org.Query().Where(org.ID(rootOrgID)).Only(ctx)
	if err != nil {
		return err
	}
	// 给根用户授权策略
	_, err = s.Grant(ctx, ent.CreatePermissionInput{
		PrincipalKind: permission.PrincipalKindUser,
		OrgID:         rootOrgID,
		UserID:        rootOrg.OwnerID,
		OrgPolicyID:   orgPolicyID,
	})
	return err
}

func (s *Service) assignRoleUserByTid(ctx context.Context, input model.AssignRoleUserInput, tid int) error {
	client := ent.FromContext(ctx)
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
	err = client.OrgRoleUser.Create().SetOrgRoleID(input.OrgRoleID).SetOrgUserID(ouid).SetUserID(input.UserID).SetOrgID(tid).Exec(ctx)
	if err != nil {
		return err
	}
	err = security.GrantRoleForUser(input.UserID, input.OrgRoleID, tid)
	if err != nil {
		log.Errorf("grant policy failed,policy is inactive: %w", err)
	}
	return err
}

// AssignRoleUser is the resolver for the assignRoleUser field.
func (s *Service) AssignRoleUser(ctx context.Context, input model.AssignRoleUserInput) error {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	return s.assignRoleUserByTid(ctx, input, tid)
}

// RevokeRoleUser is the resolver for the revokeRoleUser field.
func (s *Service) RevokeRoleUser(ctx context.Context, roleID int, userID int) error {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	if isAllow, err := s.IsAllowRevokeOrgRole(ctx, userID, roleID); err != nil {
		return err
	} else if !isAllow {
		return fmt.Errorf("no allow to revoke")
	}
	has, err := client.OrgRoleUser.Query().Where(orgroleuser.HasOrgUserWith(orguser.OrgID(tid), orguser.UserID(userID)),
		orgroleuser.HasOrgRoleWith(orgrole.OrgID(tid), orgrole.ID(roleID))).Exist(ctx)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("role user not found")
	}
	err = security.RevokeGroupForUser(userID, roleID, tid)
	if err != nil {
		return err
	}
	ouId, err := client.OrgUser.Query().Where(orguser.OrgID(tid), orguser.UserID(userID)).Select(orguser.FieldID).Int(ctx)
	if err != nil {
		return err
	}
	_, err = client.OrgRoleUser.Delete().Where(orgroleuser.OrgRoleID(roleID), orgroleuser.OrgUserID(ouId)).Exec(ctx)
	return err
}

func (s *Service) IsAllowRevokeOrgRole(ctx context.Context, userID int, orgRoleID int) (bool, error) {
	has, err := s.Client.User.Query().Where(
		user.ID(userID),
		user.HasOrgsWith(
			org.OwnerID(userID),
			org.HasPermissionsWith(
				permission.HasRoleWith(
					orgrole.ID(orgRoleID),
					orgrole.AppRoleIDNotNil(),
				),
			),
		),
	).Exist(ctx)
	if err != nil {
		return false, err
	}
	if has {
		return false, nil
	}
	return true, nil
}

func (s *Service) AssignOrganizationAppRole(ctx context.Context, orgID int, appRoleID int) error {
	client := ent.FromContext(ctx)
	ar, err := client.AppRole.Query().Where(approle.ID(appRoleID)).WithApp().Only(ctx)
	if err != nil {
		return err
	}
	has, err := client.OrgApp.Query().Where(orgapp.OrgID(orgID), orgapp.AppID(ar.AppID)).Exist(ctx)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("org not found or not has app")
	}
	has, err = client.OrgRole.Query().Where(orgrole.AppRoleID(appRoleID), orgrole.OrgID(orgID)).Exist(ctx)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("role has assigned to org")
	}
	or, err := client.OrgRole.Create().SetOrgID(orgID).SetKind(orgrole.KindRole).SetAppRoleID(ar.ID).
		SetComments(ar.Comments).SetName(ar.Name).Save(ctx)
	// 查询角色关联的策略
	aps, err := client.AppPolicy.Query().Where(
		apppolicy.AppID(ar.AppID),
		apppolicy.HasAppRolePolicyWith(approlepolicy.AppID(ar.AppID), approlepolicy.AppRoleID(appRoleID)),
	).All(ctx)
	if err != nil {
		return err
	}
	// 查询aps已分配给组织的策略
	apids := make([]int, len(aps))
	for i, ap := range aps {
		apids[i] = ap.ID
	}
	hasOps, err := client.OrgPolicy.Query().Where(orgpolicy.OrgID(orgID), orgpolicy.AppID(ar.AppID), orgpolicy.AppPolicyIDIn(apids...)).All(ctx)
	if err != nil {
		return err
	}
	hasOpMap := make(map[int]bool)
	for _, op := range hasOps {
		hasOpMap[op.AppPolicyID] = true
	}
	// 分配策略给组织
	opbk := make([]*ent.OrgPolicyCreate, 0)
	for _, ap := range aps {
		if hasOpMap[ap.ID] {
			continue
		}
		if err = appPolicyToOrgPolicy(ar.Edges.App.Code, ap.Rules, orgID); err != nil {
			return err
		}
		opbk = append(opbk, client.OrgPolicy.Create().SetOrgID(orgID).SetAppID(ar.Edges.App.ID).SetAppPolicyID(ap.ID).
			SetComments(ap.Comments).SetRules(ap.Rules).SetComments(ap.Comments).SetName(ap.Name))
	}
	ops, err := client.OrgPolicy.CreateBulk(opbk...).Save(ctx)
	ops = append(ops, hasOps...)
	if err != nil {
		return err
	}
	// role关联权限
	pbk := make([]*ent.PermissionCreate, len(ops))
	for i, op := range ops {
		pbk[i] = client.Permission.Create().SetPrincipalKind(permission.PrincipalKindRole).SetRoleID(or.ID).
			SetOrgPolicyID(op.ID).SetOrgID(orgID).SetStatus(typex.SimpleStatusActive)
	}
	ps, err := client.Permission.CreateBulk(pbk...).Save(ctx)
	if err != nil {
		return err
	}
	pids := make([]int, len(ps))
	for i, p := range ps {
		pids[i] = p.ID
	}
	ps, err = client.Permission.Query().Where(permission.IDIn(pids...)).WithOrgPolicy().All(ctx)
	for _, p := range ps {
		err = security.GrantByPermission(p, orgID)
		if err != nil {
			return err
		}
	}
	// 授权role给根用户
	return s.assignAppRoleToRootOwner(ctx, orgID, or.ID)
}

// assignAppRoleToRootOwner 给根组织管理者授权角色
func (s *Service) assignAppRoleToRootOwner(ctx context.Context, rootOrgID int, orgRoleId int) error {
	client := ent.FromContext(ctx)
	rootOrg, err := client.Org.Query().Where(org.ID(rootOrgID)).Only(ctx)
	if err != nil {
		return err
	}
	// 给根用户授权策略
	return s.assignRoleUserByTid(ctx, model.AssignRoleUserInput{
		OrgRoleID: orgRoleId,
		UserID:    *rootOrg.OwnerID,
	}, rootOrgID)
}

func (s *Service) RevokeOrganizationAppRole(ctx context.Context, orgID int, appRoleID int) error {
	client := ent.FromContext(ctx)

	isRoot, err := s.IsRootOrg(ctx, orgID)
	if err != nil {
		return err
	}
	if !isRoot {
		return fmt.Errorf("organization %d is not a root organization", orgID)
	}

	ps, err := client.OrgRoleUser.Query().Where(
		orgroleuser.HasOrgUserWith(orguser.OrgID(orgID)),
		orgroleuser.HasOrgRoleWith(orgrole.AppRoleID(appRoleID), orgrole.OrgID(orgID))).
		WithOrgRole().WithOrgUser().All(ctx)
	if err != nil {
		return err
	}
	//
	pids := make([]int, len(ps))
	for i, p := range ps {
		pids[i] = p.ID
		// 解除角色授权
		err = security.RevokeGroupForUser(p.Edges.OrgUser.UserID, p.Edges.OrgRole.ID, orgID)
		if err != nil {
			log.Error(err)
		}
	}
	_, err = client.OrgRoleUser.Delete().Where(orgroleuser.IDIn(pids...)).Exec(ctx)
	if err != nil {
		return err
	}
	// 清理OrgRole授权
	orid, err := client.OrgRole.Query().Where(orgrole.AppRoleID(appRoleID), orgrole.OrgID(orgID)).Select(orgrole.FieldID).Int(ctx)
	if err != nil {
		return err
	}
	ops, err := client.OrgPolicy.Query().Where(
		orgpolicy.HasPermissionsWith(
			permission.RoleID(orid),
			permission.PrincipalKindEQ(permission.PrincipalKindRole),
			permission.OrgID(orgID),
		),
	).All(ctx)
	if err != nil {
		return err
	}
	rules := make([]*types.PolicyRule, 0)
	opids := make([]int, len(ops))
	for i, op := range ops {
		opids[i] = op.ID
		rs := op.Rules
		if rs != nil {
			rules = append(rules, rs...)
		}
	}
	err = security.RevokePolicy(rules, strconv.Itoa(orid), orgID, permission.PrincipalKindRole)
	if err != nil {
		return err
	}
	// 清理permission
	_, err = client.Permission.Delete().Where(permission.OrgID(orgID), permission.PrincipalKindEQ(permission.PrincipalKindRole), permission.RoleID(orid)).Exec(ctx)
	if err != nil {
		return err
	}
	// 删除orgPolicy
	_, err = client.OrgPolicy.Delete().Where(orgpolicy.IDIn(opids...)).Exec(ctx)
	if err != nil {
		return err
	}
	// 删除orgRole
	_, err = client.OrgRole.Delete().Where(orgrole.ID(orid)).Exec(ctx)
	return err
}

func (s *Service) RevokeOrganizationAppPolicy(ctx context.Context, orgID int, appPolicyID int) error {
	client := ent.FromContext(ctx)
	isRoot, err := s.IsRootOrg(ctx, orgID)
	if err != nil {
		return err
	}
	if !isRoot {
		return fmt.Errorf("organization %d is not a root organization", orgID)
	}

	if has, err := s.IsAllowRevokeAppPolicy(ctx, orgID, appPolicyID); err != nil {
		return err
	} else if !has {
		return fmt.Errorf("no allow to revoke")
	}

	// 查找对应授权的组织策略
	op, err := client.OrgPolicy.Query().Where(orgpolicy.OrgID(orgID), orgpolicy.AppPolicyID(appPolicyID)).Only(ctx)
	if err != nil {
		return err
	}
	// 清空授权
	err = updateOrgPolicyRules(ctx, op.ID, nil, orgID)
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
	}

	// 删除permission
	_, err = client.Permission.Delete().Where(permission.IDIn(pids...)).Exec(ctx)
	if err != nil {
		return err
	}
	// 删除orgPolicy
	_, err = client.OrgPolicy.Delete().Where(orgpolicy.ID(op.ID)).Exec(ctx)
	return err
}

// IsAllowRevokeAppPolicy 应用策略是否允许解除授权组织
// 如果授权组织的应用角色包含该策略，则该策略不允许单独解除授权
func (s *Service) IsAllowRevokeAppPolicy(ctx context.Context, orgID int, appPolicyID int) (bool, error) {
	has, err := s.Client.Permission.Query().Where(
		permission.PrincipalKindEQ(permission.PrincipalKindRole),
		permission.OrgID(orgID),
		permission.StatusEQ(typex.SimpleStatusActive),
		permission.HasOrgPolicyWith(orgpolicy.AppPolicyID(appPolicyID)),
		permission.HasRoleWith(orgrole.AppRoleIDNotNil()),
	).Exist(ctx)
	if err != nil {
		return false, err
	}
	if has {
		return false, nil
	}
	return true, nil
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
	// 检查是否已经存在该授权,不允许重复授权
	existsq := client.Permission.Query().Where(permission.OrgID(input.OrgID), permission.PrincipalKindEQ(input.PrincipalKind),
		permission.OrgPolicyID(input.OrgPolicyID))

	isRoot, err := s.IsRootOrg(ctx, input.OrgID)
	if err != nil {
		return nil, err
	}
	if !isRoot {
		return nil, fmt.Errorf("organization %d is not a root organization", input.OrgID)
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
		existsq.Where(permission.UserID(pid))
		builder.SetUserID(pid)
		builder.SetPrincipalKind(permission.PrincipalKindUser)
	case permission.PrincipalKindRole:
		if input.RoleID == nil {
			return nil, fmt.Errorf("role id is required")
		}
		pid = *input.RoleID
		existsq.Where(permission.RoleID(pid))
		builder.SetRoleID(pid)
		builder.SetPrincipalKind(permission.PrincipalKindRole)
	default:
		return nil, fmt.Errorf("grant type %s not support", input.PrincipalKind)
	}
	if has, _ := existsq.Exist(ctx); has {
		return nil, fmt.Errorf("permission already granted")
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
	err = security.GrantByPermission(perm, input.OrgID)
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
	err = security.RevokeByPermission(p, tid)
	if err != nil {
		return p, err
	}
	if p.Status == typex.SimpleStatusActive && p.EndAt.After(time.Now()) {
		err = security.GrantByPermission(p, tid)
	}
	return p, err
}

// Revoke 撤销用户或角色的权限.
func (s *Service) Revoke(ctx context.Context, orgID int, permissionID int) error {
	client := ent.FromContext(ctx)
	isRoot, err := s.IsRootOrg(ctx, orgID)
	if err != nil {
		return err
	}
	if !isRoot {
		return fmt.Errorf("organization %d is not a root organization", orgID)
	}
	p, err := client.Permission.Query().Where(permission.ID(permissionID), permission.OrgID(orgID)).WithOrgPolicy().Only(ctx)
	if err != nil {
		return err
	}
	if isAllow, err := s.IsAllowRevokePermission(ctx, p); err != nil {
		return err
	} else if !isAllow {
		return fmt.Errorf("no allow to revoke")
	}

	// 判断actions、resources是否存在主体其他授权的policy
	var ops []*ent.OrgPolicy
	wheres := []predicate.Permission{
		permission.PrincipalKindEQ(p.PrincipalKind),
		permission.OrgID(orgID),
	}
	if p.PrincipalKind == permission.PrincipalKindUser {
		// 用户
		wheres = append(wheres, permission.UserID(p.UserID))
	} else if p.PrincipalKind == permission.PrincipalKindRole {
		// 角色
		wheres = append(wheres, permission.RoleID(p.RoleID))
	} else {
		return fmt.Errorf("error PrincipalKind")
	}
	ps, err := client.Permission.Query().Where(wheres...).WithOrgPolicy().All(ctx)
	if err != nil {
		return err
	}
	for _, p := range ps {
		ops = append(ops, p.Edges.OrgPolicy)
	}
	delActions, delResources := joinPolicyRules(p.Edges.OrgPolicy.Rules)
	// 同步修改casbin授权信息
	err = updatePoliciesToCasbin(p.OrgPolicyID, map[string][]*ent.OrgPolicy{strconv.Itoa(getPrincipalID(p)): ops}, orgID, p.PrincipalKind, nil, nil, delActions, delResources)
	if err != nil {
		return err
	}

	_, err = client.Permission.Delete().Where(permission.ID(permissionID), permission.OrgID(orgID)).Exec(ctx)
	return err
}

func (s *Service) IsAllowRevokePermission(ctx context.Context, p *ent.Permission) (bool, error) {
	if p.PrincipalKind == permission.PrincipalKindUser {
		// 根用户不允许revoke
		has, err := s.Client.Permission.Query().Where(
			permission.ID(p.ID),
			permission.HasOrgWith(org.HasOwnerWith(user.ID(p.UserID))),
		).Exist(ctx)
		if err != nil {
			return false, err
		}
		if has {
			return false, err
		}
	} else if p.PrincipalKind == permission.PrincipalKindRole {
		// 系统角色不允许revoke
		has, err := s.Client.Permission.Query().Where(
			permission.ID(p.ID),
			permission.HasRoleWith(orgrole.ID(p.RoleID), orgrole.AppRoleIDNotNil()),
		).Exist(ctx)
		if err != nil {
			return false, err
		}
		if has {
			return false, err
		}
	}
	return true, nil
}

func (s *Service) GetUserPermissionsByUserID(ctx context.Context, userID int, where *ent.AppActionWhereInput) ([]*ent.AppAction, error) {
	client := s.Client
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// 根据appcode分组
	grantActions := make(map[string][]string)
	// 拥有全部权限的app
	fulGrantApps := make(map[string]*string)
	ups := security.GetUserPermissions(userID, tid)
	for _, p := range ups {
		parts := strings.Split(p[2], ArnSplit)
		if len(parts) != 2 {
			continue
		}
		ac := parts[0]
		// 标记app所有授权
		if parts[1] == "*" {
			fulGrantApps[ac] = &parts[1]
		}
		if grantActions[ac] == nil {
			grantActions[ac] = []string{
				parts[1],
			}
		} else {
			grantActions[ac] = append(grantActions[ac], parts[1])
		}
	}
	userActions := make([]*ent.AppAction, 0)
	for appcode, actions := range grantActions {
		if fulGrantApps[appcode] != nil {
			// 拥有应用全部权限
			query, err := where.Filter(client.AppAction.Query())
			if err != nil {
				return nil, err
			}
			aas, err := query.Where(appaction.HasAppWith(app.Code(appcode))).All(ctx)
			if err != nil {
				return nil, err
			}
			userActions = append(userActions, aas...)
		} else {
			query, err := where.Filter(client.AppAction.Query())
			if err != nil {
				return nil, err
			}
			// actions 去重
			newActions := RemoveDuplicateElement(actions)
			aas, err := query.Where(appaction.NameIn(newActions...), appaction.HasAppWith(app.Code(appcode))).All(ctx)
			if err != nil {
				return nil, err
			}
			userActions = append(userActions, aas...)
		}
	}
	return userActions, nil
}

// GetUserPermissions 获取用户的全部权限
func (s *Service) GetUserPermissions(ctx context.Context, where *ent.AppActionWhereInput) ([]*ent.AppAction, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return s.GetUserPermissionsByUserID(ctx, uid, where)
}

func (s *Service) GetUserApps(ctx context.Context) ([]*ent.App, error) {
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// 获取appcode
	acs := make([]string, 0)
	ups := security.GetUserPermissions(uid, tid)
	for _, p := range ups {
		parts := strings.Split(p[2], ArnSplit)
		if len(parts) > 2 {
			continue
		}
		acs = append(acs, parts[0])
	}
	acs = RemoveDuplicateElement(acs)
	return s.Client.App.Query().Where(app.CodeIn(acs...)).All(ctx)
}

func (s *Service) CheckPermission(ctx context.Context, permission string) (bool, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	// 检查permission有效
	parts := strings.SplitN(permission, ":", 2)
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid permission")
	}
	has, err := s.Client.AppAction.Query().Where(appaction.Name(parts[1]), appaction.HasAppWith(app.Code(parts[0]))).Exist(ctx)
	if err != nil {
		return false, err
	}
	if !has {
		return false, fmt.Errorf("invalid permission")
	}

	rule := []any{
		strconv.Itoa(uid),
		strconv.Itoa(tid),
		permission,
		"read",
	}
	has, err = security.CheckUserPermission(rule...)
	if err != nil {
		return false, err
	}
	if !has {
		return false, nil
	}
	return true, nil
}

// GetOrgDomain 获取组织域名.orgID为根组织.
func (s *Service) GetOrgDomain(ctx context.Context, orgID int) (string, error) {
	c := s.Client
	orgr := c.Org.Query().Where(org.ID(orgID)).Select(org.FieldDomain).OnlyX(ctx)
	if orgr.Domain == "" {
		return "", fmt.Errorf("organization %d domain is empty", orgID)
	}
	return orgr.Domain, nil
}

// IsRootOrg 判断组织是否root
func (s *Service) IsRootOrg(ctx context.Context, orgID int) (bool, error) {
	return s.Client.Org.Query().Where(org.ID(orgID)).Where(org.KindEQ(org.KindRoot)).Exist(ctx)
}

// GetRootOrgByUser 获取用户的最顶级的根组织.在组织中,一个账户可能存在多个根组织.需要从context获取租户ID
func (s *Service) GetRootOrgByUser(ctx context.Context, uid int) (*ent.Org, error) {
	c := s.Client
	return c.Org.Query().Where(org.HasUsersWith(user.ID(uid)), org.KindEQ(org.KindRoot)).
		Order(ent.Asc(org.FieldPath)).First(ctx)
}

// updateOrgPolicyRules 更新策略规则
// 策略rules更新，需根据策略的引用同步更新casbin
func updateOrgPolicyRules(ctx context.Context, orgPolicyId int, rules []*types.PolicyRule, domain int) error {
	client := ent.FromContext(ctx)
	// 更新的rules，拼接数据：action&&allow | action&&deny | resource&&allow | resource&&deny
	naas, nars := joinPolicyRules(rules)
	// 查询数据库policy
	op, err := client.OrgPolicy.Query().Where(orgpolicy.ID(orgPolicyId)).Only(ctx)
	if err != nil {
		return err
	}
	// 旧的actions、resources
	oaas, oars := joinPolicyRules(op.Rules)

	// 更新policy时,获取 addActions：新增actions、delActions：删除actions、addResources：新增resources、delResources：删除resources
	addActions, delActions := DiffArrays(naas, oaas)
	addResources, delResources := DiffArrays(nars, oars)

	// 查询policy授权的users、roles
	ps, err := client.Permission.Query().Where(permission.OrgPolicyID(orgPolicyId), permission.OrgID(domain)).All(ctx)
	if err != nil {
		return err
	}
	var uids, rids []int
	for _, p := range ps {
		if p.PrincipalKind == permission.PrincipalKindUser {
			uids = append(uids, p.UserID)
		} else if p.PrincipalKind == permission.PrincipalKindRole {
			rids = append(rids, p.RoleID)
		}
	}

	// 查询授权给user的policy
	ups, err := client.Permission.Query().Where(permission.PrincipalKindEQ(permission.PrincipalKindUser), permission.UserIDIn(uids...), permission.OrgID(domain)).WithOrgPolicy().All(ctx)
	if err != nil {
		return err
	}
	uops := make(map[string][]*ent.OrgPolicy)
	// 根据用户授权分组
	for _, p := range ups {
		if uops[strconv.Itoa(p.UserID)] == nil {
			uops[strconv.Itoa(p.UserID)] = []*ent.OrgPolicy{
				p.Edges.OrgPolicy,
			}
		} else {
			uops[strconv.Itoa(p.UserID)] = append(uops[strconv.Itoa(p.UserID)], p.Edges.OrgPolicy)
		}
	}
	// 更新权限
	err = updatePoliciesToCasbin(orgPolicyId, uops, domain, permission.PrincipalKindUser, addActions, addResources, delActions, delResources)
	if err != nil {
		return err
	}

	// 查询授权给role的policy
	rps, err := client.Permission.Query().Where(permission.PrincipalKindEQ(permission.PrincipalKindRole), permission.RoleIDIn(rids...), permission.OrgID(domain)).WithOrgPolicy().All(ctx)
	if err != nil {
		return err
	}
	rops := make(map[string][]*ent.OrgPolicy)
	// 根据用户授权分组
	for _, p := range rps {
		if rops[strconv.Itoa(p.RoleID)] == nil {
			rops[strconv.Itoa(p.RoleID)] = []*ent.OrgPolicy{
				p.Edges.OrgPolicy,
			}
		} else {
			rops[strconv.Itoa(p.RoleID)] = append(rops[strconv.Itoa(p.RoleID)], p.Edges.OrgPolicy)
		}
	}
	// 更新权限
	err = updatePoliciesToCasbin(orgPolicyId, rops, domain, permission.PrincipalKindRole, addActions, addResources, delActions, delResources)
	if err != nil {
		return err
	}
	return nil
}

// updatePoliciesToCasbin
func updatePoliciesToCasbin(exOrgPolicyId int, orgPolicies map[string][]*ent.OrgPolicy, domain int, principalKind permission.PrincipalKind, addActions []string, addResources []string, delActions []string, delResources []string) error {
	// 根据用户授权的policy来判断需要移除的actions及resources
	for principal, uop := range orgPolicies {
		var otherPolicyActions, otherPolicyResources []string
		for _, p := range uop {
			// 过滤当前policy id
			if p.ID == exOrgPolicyId {
				continue
			}
			for _, rule := range p.Rules {
				// actions
				for _, action := range rule.Actions {
					otherPolicyActions = append(otherPolicyActions, action+SplitPolicyEffect+rule.Effect.String())
				}
				// resource
				for _, res := range rule.Actions {
					otherPolicyResources = append(otherPolicyResources, res+SplitPolicyEffect+rule.Effect.String())
				}
			}
		}
		// 移除的actions与用户其他策略的actions对比，不存在其他策略则删除
		actuallyDelActions, _ := DiffArrays(delActions, otherPolicyActions)
		// 移除的resources与用户其他策略的resources对比，不存在其他策略则删除
		actuallyDelResources, _ := DiffArrays(delResources, otherPolicyResources)

		var addPolicyRule []*types.PolicyRule
		var delPolicyRule []*types.PolicyRule
		splitAddAllowActions, splitAddDenyActions := splitPolicyRules(addActions)
		splitAddAllowResources, splitAddDenyResources := splitPolicyRules(addResources)
		splitDelAllowActions, splitDelDenyActions := splitPolicyRules(actuallyDelActions)
		splitDelAllowResources, splitDelDenyResources := splitPolicyRules(actuallyDelResources)
		addAllowRule := types.PolicyRule{
			Effect:    types.PolicyEffectAllow,
			Actions:   splitAddAllowActions,
			Resources: splitAddAllowResources,
		}
		addDenyRule := types.PolicyRule{
			Effect:    types.PolicyEffectDeny,
			Actions:   splitAddDenyActions,
			Resources: splitAddDenyResources,
		}
		addPolicyRule = append(addPolicyRule, &addAllowRule)
		addPolicyRule = append(addPolicyRule, &addDenyRule)
		// 新增的授权
		err := security.GrantPolicy(addPolicyRule, principal, domain, principalKind)
		if err != nil {
			return err
		}
		delAllowRule := types.PolicyRule{
			Effect:    types.PolicyEffectAllow,
			Actions:   splitDelAllowActions,
			Resources: splitDelAllowResources,
		}
		delDenyRule := types.PolicyRule{
			Effect:    types.PolicyEffectDeny,
			Actions:   splitDelDenyActions,
			Resources: splitDelDenyResources,
		}
		delPolicyRule = append(delPolicyRule, &delAllowRule)
		delPolicyRule = append(delPolicyRule, &delDenyRule)
		// 删除的授权
		err = security.RevokePolicy(delPolicyRule, principal, domain, principalKind)
		if err != nil {
			return err
		}
	}
	return nil
}

// joinPolicyRules actions与resources拼接Effect
// 返回值：拼接后的actions、resources
func joinPolicyRules(rules []*types.PolicyRule) ([]string, []string) {
	var aas, ars []string
	for _, rule := range rules {
		for _, action := range rule.Actions {
			aas = append(aas, action+SplitPolicyEffect+rule.Effect.String())
		}
		for _, res := range rule.Resources {
			ars = append(ars, res+SplitPolicyEffect+rule.Effect.String())
		}
	}
	return aas, ars
}

// 截取actions或者resources
// 返回值：截取后返回allows、denies
func splitPolicyRules(data []string) ([]string, []string) {
	var allows, denies []string
	for _, res := range data {
		parts := strings.SplitN(res, SplitPolicyEffect, 2)
		if parts[1] == types.PolicyEffectAllow.String() {
			allows = append(allows, parts[0])
		} else if parts[1] == types.PolicyEffectDeny.String() {
			denies = append(denies, parts[0])
		}
	}
	return allows, denies
}
