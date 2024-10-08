package resource

import (
	"context"
	"fmt"
	"github.com/woocoos/entcache"
	"github.com/woocoos/knockout-go/api/msg"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout/api/graphql/model"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orgpolicy"
	"github.com/woocoos/knockout/ent/orgrole"
	"github.com/woocoos/knockout/ent/orgroleuser"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/orguserpreference"
	"github.com/woocoos/knockout/ent/permission"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/ent/userloginprofile"
	"github.com/woocoos/knockout/ent/userpassword"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// EnableOrganization 开启组织目录
func (s *Service) EnableOrganization(ctx context.Context, input model.EnableDirectoryInput) (*ent.Org, error) {
	client := ent.FromContext(ctx)
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	exist, err := client.Org.Query().Where(org.OwnerID(uid)).Exist(entcache.Skip(ctx))
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("directory service has enable")
	}
	orgd, err := client.Org.Create().SetOwnerID(uid).SetName(input.Name).SetDomain(input.Domain).
		SetKind(org.KindRoot).SetStatus(typex.SimpleStatusActive).Save(ctx)
	if err != nil {
		return nil, err
	}
	ur := client.User.GetX(ctx, uid)
	err = client.OrgUser.Create().SetOrgID(orgd.ID).SetUserID(uid).SetDisplayName(ur.DisplayName).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return orgd, err
}

// CreateRoot 创建组织root
func (s *Service) CreateRoot(ctx context.Context, input ent.CreateOrgInput) (*ent.Org, error) {
	client := ent.FromContext(ctx)
	o, err := client.Org.Create().SetInput(input).SetKind(org.KindRoot).Save(ctx)
	if err != nil {
		return nil, err
	}
	// 有管理用户则加入组织
	if input.OwnerID != nil {
		u, err := client.User.Query().Where(user.ID(*input.OwnerID)).Only(ctx)
		if err != nil {
			return nil, err
		}
		if u.UserType != user.UserTypeAccount {
			return nil, fmt.Errorf("owner must be account")
		}
		err = client.OrgUser.Create().SetOrgID(o.ID).SetUserID(*input.OwnerID).SetDisplayName(u.DisplayName).Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
}

// CreateOrganization 创建组织目录,基于根目录创建
func (s *Service) CreateOrganization(ctx context.Context, input ent.CreateOrgInput) (*ent.Org, error) {
	if input.ParentID == 0 {
		return nil, fmt.Errorf("parent id is required")
	}
	return s.Client.Org.Create().SetInput(input).SetKind(org.KindOrganization).Save(ctx)
}

// DeleteOrganization 删除组织目录
func (s *Service) DeleteOrganization(ctx context.Context, id int) error {
	client := ent.FromContext(ctx)
	count, err := client.Org.Query().Where(
		org.ParentID(id),
	).Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("organization has children")
	}
	count, err = client.Org.Query().Where(org.ID(id), org.HasUsers()).Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("organization has users")
	}
	return client.Org.DeleteOneID(id).Exec(ctx)
}

// CreateOrganizationAccount 创建组织目录账户,进入账户激活流程
//
// - 管理员账户才能创建下级组织目录的账户
func (s *Service) CreateOrganizationAccount(ctx context.Context, orgId int, input ent.CreateUserInput) (*ent.User, error) {
	return s.CreateOrganizationUser(ctx, orgId, input, user.UserTypeAccount)
}

// CreateOrganizationUser 创建组织目录用户
//
// TODO 新用户需要激活,如在国内,用户往往需要绑定手机或邮箱,然后通过邮件或短信激活.
func (s *Service) CreateOrganizationUser(ctx context.Context, orgId int, input ent.CreateUserInput, ut user.UserType) (*ent.User, error) {
	client := ent.FromContext(ctx)
	_, err := client.Org.Query().Where(org.ID(orgId), org.StatusEQ(typex.SimpleStatusActive)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("organization not exists or inactive")
	}

	us, err := client.User.Create().SetInput(input).
		SetUserType(ut).
		SetCreationType(user.CreationTypeManual).
		SetRegisterIP("").
		Save(ctx)
	if err != nil {
		return nil, err
	}

	_, err = client.UserIdentity.Create().SetUserID(us.ID).SetCode(input.PrincipalName).SetKind(useridentity.KindName).
		SetStatus(typex.SimpleStatusActive).Save(ctx)
	if err != nil {
		return nil, err
	}

	if ut != user.UserTypeAccount {
		_, err = client.OrgUser.Create().SetOrgID(orgId).SetUserID(us.ID).SetDisplayName(us.DisplayName).Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	// 自动生成密码,发送邮件
	ulp, err := us.QueryLoginProfile().Only(ctx)
	if err != nil {
		return nil, err
	}
	if ulp.SetKind == userloginprofile.SetKindAuto {
		err = s.generationAndSendUserPwd(ctx, us)
		if err != nil {
			return nil, err
		}
	}
	return us, nil
}

// generationAndSendUserPwd 自动生成密码并发邮件给用户
func (s *Service) generationAndSendUserPwd(ctx context.Context, usr *ent.User) error {
	if usr.Email == "" {
		return fmt.Errorf("email is nil")
	}
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	nPwd := RandomStr(6)
	shaPwd := SHA256(RandomStr(6))
	// 创建用户密码
	_, err = s.CreateUserPassword(ctx, &ent.CreateUserPasswordInput{
		Scene:    userpassword.SceneLogin,
		Password: &shaPwd,
		UserID:   &usr.ID,
	})
	if err != nil {
		return err
	}

	params := msg.PostableAlerts{
		{
			Annotations: map[string]string{
				"to":            usr.Email,
				"displayName":   usr.DisplayName,
				"principalName": usr.PrincipalName,
				"password":      nPwd,
			},
			Alert: &msg.Alert{
				Labels: map[string]string{
					"receiver":  "email",
					"alertname": "UserPasswordAndPrincipal",
					"tenant":    strconv.Itoa(tid),
					"timestamp": strconv.Itoa(int(time.Now().Unix())),
				},
			},
		},
	}
	return s.postAlerts(ctx, params)
}

func (s *Service) CreateUserPassword(ctx context.Context, input *ent.CreateUserPasswordInput) (pw *ent.UserPassword, err error) {
	if input == nil {
		return
	}
	salt := RandomStr(5)
	var hashPwd string
	if input.Password != nil || *input.Password != "" {
		hashPwd = SHA256(*input.Password + salt)
	} else {
		hashPwd = RandomStr(6)
		hashPwd = SHA256(hashPwd + salt)
	}
	pw, err = ent.FromContext(ctx).UserPassword.Create().
		SetInput(*input).
		SetPassword(hashPwd).
		SetSalt(salt).
		Save(ctx)

	return
}

// MoveOrganization 移动组织目录.
func (s *Service) MoveOrganization(ctx context.Context, src, tar int, action model.TreeAction) (err error) {
	client := ent.FromContext(ctx)
	tarOrg := client.Org.GetX(ctx, tar)
	builder := client.Org.UpdateOneID(src)
	var start int32 = 0
	var resort = true
	switch action {
	case model.TreeActionChild:
		var agg []struct {
			Max *int32
		}
		err = client.Org.Query().Where(org.ParentID(tarOrg.ID)).Aggregate(ent.Max(org.FieldDisplaySort)).Scan(ctx, &agg)
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
		err = client.Org.Update().Where(org.ParentID(tarOrg.ParentID), org.DisplaySortGTE(start)).AddDisplaySort(1).Exec(ctx)
		if err != nil {
			return
		}
	}

	return builder.Exec(ctx)
}

// AllotOrganizationUser 将用户加入组织目录
func (s *Service) AllotOrganizationUser(ctx context.Context, input ent.CreateOrgUserInput) error {
	client := ent.FromContext(ctx)
	has, err := client.OrgUser.Query().Where(orguser.OrgID(input.OrgID), orguser.UserID(input.UserID)).Exist(ctx)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("user already in organization")
	}
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	orgs, err := client.Org.Query().Where(org.IDIn(tid, input.OrgID)).Order(ent.Asc(org.FieldPath)).All(ctx)
	if err != nil {
		return err
	}
	if len(orgs) != 2 {
		return fmt.Errorf("invalid org id or root org id")
	}
	if !strings.HasPrefix(orgs[1].Path, orgs[0].Path) {
		return fmt.Errorf("org not match")
	}

	usr := client.User.GetX(ctx, input.UserID)
	if input.DisplayName == "" {
		input.DisplayName = usr.DisplayName
	}

	return client.OrgUser.Create().SetInput(input).Exec(ctx)
}

// RemoveOrganizationUser 将用户从组织目录中移除.
func (s *Service) RemoveOrganizationUser(ctx context.Context, orgID int, userID int) error {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	if orgID == tid {
		return fmt.Errorf("can not remove from root org")
	}
	i, err := client.OrgUser.Delete().Where(orguser.UserID(userID), orguser.OrgID(orgID)).Exec(ctx)
	if err != nil {
		return err
	}
	if i == 0 {
		return fmt.Errorf("user not in org")
	}

	return nil
}

// DeleteOrganizationUser 删除本域下的用户,在用户没有被引用时,允许删除
func (s *Service) DeleteOrganizationUser(ctx context.Context, userID int) error {
	client := ent.FromContext(ctx)

	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	code := strconv.FormatInt(int64(tid), 36)

	usr := client.User.GetX(ctx, userID)
	// 是否组织引用
	ins, err := client.OrgUser.Query().Where(orguser.UserID(userID),
		orguser.HasOrgWith(org.PathHasPrefix(code)),
	).Count(ctx)
	if err != nil {
		return err
	}
	if ins == 0 {
		return fmt.Errorf("user not in org")
	}
	if ins > 1 {
		return fmt.Errorf("user in more than one org")
	}
	// 根据授权判断是否被引用
	has, err := client.Permission.Query().Where(permission.UserID(userID), permission.HasOrgWith(org.PathHasPrefix(code))).Exist(ctx)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("user has been referenced")
	}
	// 根据角色判断是否被引用
	has, err = client.OrgRoleUser.Query().Where(orgroleuser.HasOrgUserWith(orguser.UserID(userID)), orgroleuser.HasOrgRoleWith(orgrole.HasOrgWith(org.ID(tid)))).Exist(ctx)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("user has been referenced")
	}

	_, err = client.OrgUser.Delete().Where(orguser.UserID(userID), orguser.OrgID(tid)).Exec(ctx)
	if err != nil {
		return err
	}
	if usr.CreationType != user.CreationTypeManual {
		// 非手动创建的用户,不允许删除,移除后返回
		return nil
	}
	_, err = client.UserIdentity.Delete().Where(useridentity.UserID(userID)).Exec(ctx)
	if err != nil {
		return err
	}
	_, err = client.UserPassword.Delete().Where(userpassword.UserID(userID)).Exec(ctx)
	if err != nil {
		return err
	}
	return client.User.Update().Where(user.ID(userID)).ClearIdentities().ClearPasswords().
		SetDeletedAt(time.Now()).SetStatus(typex.SimpleStatusInactive).Exec(ctx)
}

// UpdateUser 更新用户信息,允许更新用户的email,phone,但这些信息需要通过验证被引入UserIdentity中才能生效.
func (s *Service) UpdateUser(ctx context.Context, userID int, input ent.UpdateUserInput) (*ent.User, error) {
	if input.PrincipalName != nil {
		return nil, fmt.Errorf("principal name can not update")
	}
	client := ent.FromContext(ctx)
	return client.User.UpdateOneID(userID).SetInput(input).Save(ctx)
}

func (s *Service) ChangePassword(ctx context.Context, oldPwd, newPwd string) error {
	if oldPwd == newPwd {
		return fmt.Errorf("old password can not equal new password")
	}
	client := ent.FromContext(ctx)
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return err
	}
	usr := client.User.Query().Where(user.ID(uid)).
		WithPasswords(func(query *ent.UserPasswordQuery) {
			query.Where(userpassword.SceneEQ(userpassword.SceneLogin))
		}).OnlyX(ctx)

	o := SaltSecret(oldPwd, usr.Edges.Passwords[0].Salt)
	n := SaltSecret(newPwd, usr.Edges.Passwords[0].Salt)
	if o != usr.Edges.Passwords[0].Password {
		return fmt.Errorf("old password not match")
	}

	_, err = client.UserPassword.UpdateOneID(usr.Edges.Passwords[0].ID).
		SetPassword(n).Save(ctx)
	return err
}

func (s *Service) UpdateLoginProfile(ctx context.Context, userID int, input ent.UpdateUserLoginProfileInput) (*ent.UserLoginProfile, error) {
	client := ent.FromContext(ctx)
	lpID, err := client.UserLoginProfile.Query().Where(userloginprofile.UserID(userID)).Select(userloginprofile.FieldID).Int(ctx)
	if err != nil {
		return nil, err
	}
	return client.UserLoginProfile.UpdateOneID(lpID).SetInput(input).Save(ctx)
}

// CreateRole 创建角色或工作组
func (s *Service) CreateRole(ctx context.Context, input ent.CreateOrgRoleInput) (*ent.OrgRole, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return client.OrgRole.Create().SetInput(input).SetOrgID(tid).Save(ctx)
}

// UpdateRole 更新角色或工作组
func (s *Service) UpdateRole(ctx context.Context, roleID int, input ent.UpdateOrgRoleInput) (*ent.OrgRole, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return client.OrgRole.UpdateOneID(roleID).Where(orgrole.OrgID(tid)).SetInput(input).Save(ctx)
}

// DeleteRole 删除角色或工作组
func (s *Service) DeleteRole(ctx context.Context, roleID int) error {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	err = client.OrgRole.DeleteOneID(roleID).Where(orgrole.OrgID(tid)).Exec(ctx)
	return err
}

// CreateOrganizationPolicy 创建组织策略,该策略属于租户组织
func (s *Service) CreateOrganizationPolicy(ctx context.Context, input ent.CreateOrgPolicyInput) (*ent.OrgPolicy, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return client.OrgPolicy.Create().SetOrgID(tid).SetInput(input).Save(ctx)
}

func (s *Service) UpdateOrganizationPolicy(ctx context.Context, id int, input ent.UpdateOrgPolicyInput) (*ent.OrgPolicy, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	// rules不为空，则同步修改casbin授权信息
	if input.Rules != nil {
		err := updateOrgPolicyRules(ctx, id, input.Rules, tid)
		if err != nil {
			return nil, err
		}
	}

	data, err := client.OrgPolicy.UpdateOneID(id).Where(orgpolicy.OrgID(tid)).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Service) DeleteOrganizationPolicy(ctx context.Context, orgPolicyID int) error {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	// 存在引用，不能删除
	has, err := client.Permission.Query().Where(permission.OrgID(tid), permission.OrgPolicyID(orgPolicyID)).Exist(ctx)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("policy has be referenced，not allowed to delete")
	}
	return client.OrgPolicy.DeleteOneID(orgPolicyID).Where(orgpolicy.OrgID(tid)).Exec(ctx)
}

// GetRoleUserIds 获取组织用户组/角色用户ids
func (s *Service) GetRoleUserIds(ctx context.Context, roleID int) ([]int, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	exist, err := s.Client.OrgRole.Query().Where(orgrole.OrgID(tid), orgrole.ID(roleID)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("role not found")
	}
	ouIds, err := s.Client.OrgRoleUser.Query().Where(orgroleuser.OrgRoleID(roleID)).Select(orgroleuser.FieldOrgUserID).Ints(ctx)
	if err != nil {
		return nil, err
	}
	return s.Client.OrgUser.Query().Where(orguser.IDIn(ouIds...)).Select(orguser.FieldUserID).Ints(ctx)
}

// EnableMFA 启用用户的MFA验证
func (s *Service) EnableMFA(ctx context.Context, userID int) (*model.Mfa, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	usr, err := client.User.Query().Where(user.ID(userID), user.HasOrgUserWith(orguser.OrgID(tid))).Only(ctx)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, fmt.Errorf("user not found")
	}
	sec := GeneralMFASecret()
	err = client.UserLoginProfile.UpdateOneID(userID).SetMfaEnabled(true).SetMfaSecret(sec).SetMfaStatus(typex.SimpleStatusActive).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Mfa{
		Account: usr.PrincipalName,
		Secret:  sec,
	}, nil
}

func (s *Service) DisableMFA(ctx context.Context, userID int) error {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	usr, err := client.User.Query().Where(user.ID(userID), user.HasOrgUserWith(orguser.OrgID(tid))).Only(ctx)
	if err != nil {
		return err
	}
	if usr == nil {
		return fmt.Errorf("user not found")
	}
	return client.UserLoginProfile.Update().Where(userloginprofile.UserID(userID)).ClearMfaEnabled().ClearMfaSecret().ClearMfaStatus().Exec(ctx)
}

func (s *Service) GetUserMenus(ctx context.Context, appCode string) ([]*ent.AppMenu, error) {
	ams, err := s.Client.AppMenu.Query().Where(appmenu.HasAppWith(app.Code(appCode))).All(ctx)
	if err != nil {
		return nil, err
	}
	// 获取用户在当前app的所有权限
	ups, err := s.GetUserPermissions(ctx, &ent.AppActionWhereInput{
		HasAppWith: []*ent.AppWhereInput{
			{Code: &appCode},
		},
	})
	if err != nil {
		return nil, err
	}
	// 找出路由权限
	ras := make(map[int]*ent.AppAction)
	for _, a := range ups {
		if a.Kind == appaction.KindRoute && a.Method == appaction.MethodRead {
			ras[a.ID] = a
		}
	}
	// 找出用户授权的菜单
	ums := make([]*ent.AppMenu, 0)
	for _, m := range ams {
		if m.ActionID == nil {
			continue
		}
		if ras[*m.ActionID] != nil {
			ums = append(ums, m)
		}
	}
	// 找出菜单的父节点
	pms := make([]*ent.AppMenu, 0, len(ams))
	findMenuParents(ams, ums, &pms)
	// pms 去重
	temp := make(map[string]bool)
	result := make([]*ent.AppMenu, 0, len(pms))
	for _, v := range pms {
		key := strconv.Itoa(v.ID)
		if _, ok := temp[key]; !ok {
			temp[key] = true
			result = append(result, v)
		}
	}
	ums = append(ums, result...)
	return ums, nil
}

func findMenuParents(appMenus, userMenus []*ent.AppMenu, parentMenus *[]*ent.AppMenu) {
	for _, um := range userMenus {
		for _, am := range appMenus {
			if um.ParentID == am.ID {
				*parentMenus = append(*parentMenus, am)
				if um.ParentID != 0 {
					findMenuParents(appMenus, []*ent.AppMenu{am}, parentMenus)
				}
				continue
			}
		}
	}
}

// RecoverOrgUser 恢复删除用户
func (s *Service) RecoverOrgUser(ctx context.Context, userID int, userInput ent.UpdateUserInput, pwdKind userloginprofile.SetKind, pwdInput *ent.CreateUserPasswordInput) (*ent.User, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	has, err := client.Org.Query().Where(org.ID(tid), org.StatusEQ(typex.SimpleStatusActive)).Exist(ctx)
	if !has || err != nil {
		return nil, fmt.Errorf("organization not exists or inactive")
	}

	us, err := client.User.UpdateOneID(userID).SetInput(userInput).SetStatus(typex.SimpleStatusActive).ClearDeletedAt().Save(ctx)
	if err != nil {
		return nil, err
	}
	pn := us.PrincipalName
	if userInput.PrincipalName != nil {
		pn = *userInput.PrincipalName
	}
	_, err = client.UserIdentity.Create().SetUserID(us.ID).SetCode(pn).SetKind(useridentity.KindName).
		SetStatus(typex.SimpleStatusActive).Save(ctx)
	if err != nil {
		return nil, err
	}

	if pwdKind == userloginprofile.SetKindAuto {
		// 自动生成密码
		err = s.generationAndSendUserPwd(ctx, us)
		if err != nil {
			return nil, err
		}
	} else if pwdInput != nil {
		if pwdInput.UserID == nil {
			pwdInput.UserID = &userID
		}
		_, err = s.CreateUserPassword(ctx, pwdInput)
		if err != nil {
			return nil, err
		}
	}

	err = client.UserLoginProfile.Update().Where(userloginprofile.UserID(userID)).SetSetKind(pwdKind).Exec(ctx)
	if err != nil {
		return nil, err
	}

	_, err = client.OrgUser.Create().SetOrgID(tid).SetUserID(us.ID).SetDisplayName(us.DisplayName).Save(ctx)
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (s *Service) SendMFAToUserByEmail(ctx context.Context, userID int) error {
	usr, err := s.Client.User.Query().Where(user.ID(userID)).WithLoginProfile().Only(ctx)
	if err != nil {
		return err
	}
	if usr.Email == "" {
		return fmt.Errorf("email is null")
	}
	if !usr.Edges.LoginProfile.MfaEnabled {
		return fmt.Errorf("mfa is disabled")
	}
	if usr.Edges.LoginProfile.MfaSecret == "" {
		return fmt.Errorf("mfa secret is null")
	}

	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	params := msg.PostableAlerts{
		{
			Annotations: map[string]string{
				"to":          usr.Email,
				"displayName": usr.DisplayName,
				"mfaSecret":   usr.Edges.LoginProfile.MfaSecret,
			},
			Alert: &msg.Alert{
				Labels: map[string]string{
					"receiver":  "email",
					"alertname": "SendMFAToUser",
					"tenant":    strconv.Itoa(tid),
					"timestamp": strconv.Itoa(int(time.Now().Unix())),
				},
			},
		},
	}
	return s.postAlerts(ctx, params)
}

func (s *Service) ResetUserPasswordByEmail(ctx context.Context, userID int) error {
	client := ent.FromContext(ctx)
	usr, err := client.User.Query().Where(user.ID(userID)).WithIdentities().WithPasswords().Only(ctx)
	if err != nil {
		return err
	}
	if usr.Email == "" {
		return fmt.Errorf("email is null")
	}

	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}

	slat := ""
	for _, v := range usr.Edges.Passwords {
		if v.Scene == userpassword.SceneLogin {
			slat = v.Salt
			break
		}
	}
	newPwd := RandomStr(6)
	// 更新用户密码
	err = client.UserPassword.UpdateOneID(userID).Where(userpassword.SceneEQ(userpassword.SceneLogin)).SetPassword(SaltSecret(newPwd, slat)).Exec(ctx)
	if err != nil {
		return err
	}
	params := msg.PostableAlerts{
		{
			Annotations: map[string]string{
				"to":            usr.Email,
				"displayName":   usr.DisplayName,
				"principalName": usr.Edges.Identities[0].Code,
				"password":      newPwd,
			},
			Alert: &msg.Alert{
				Labels: map[string]string{
					"receiver":  "email",
					"alertname": "ResetUserPassword",
					"tenant":    strconv.Itoa(tid),
					"timestamp": strconv.Itoa(int(time.Now().Unix())),
				},
			},
		},
	}
	return s.postAlerts(ctx, params)
}

func (s *Service) postAlerts(ctx context.Context, params msg.PostableAlerts) error {
	resp, err := s.KOSDK.Msg().AlertAPI.PostAlerts(ctx, &msg.PostAlertsRequest{
		PostableAlerts: params,
	})
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf(resp.Status)
}

func (s *Service) SaveOrgUserPreference(ctx context.Context, input model.OrgUserPreferenceInput) (*ent.OrgUserPreference, error) {
	client := ent.FromContext(ctx)
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	uid, err := identity.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	oup, err := client.OrgUserPreference.Query().Where(orguserpreference.UserID(uid), orguserpreference.OrgID(tid)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// 未找到则新增
			create := client.OrgUserPreference.Create().SetUserID(uid).SetOrgID(tid)
			if input.MenuRecent != nil {
				create.SetMenuRecent(input.MenuRecent)
			}
			if input.MenuFavorite != nil {
				create.SetMenuFavorite(input.MenuFavorite)
			}
			return create.Save(ctx)
		}
		return nil, err
	}
	// 更新数据
	update := client.OrgUserPreference.UpdateOneID(oup.ID)
	if input.MenuRecent != nil {
		update.SetMenuRecent(input.MenuRecent)
	}
	if input.MenuFavorite != nil {
		update.SetMenuFavorite(input.MenuFavorite)
	}
	return update.Save(ctx)
}
