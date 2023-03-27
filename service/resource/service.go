package resource

import (
	"ariga.io/entcache"
	"context"
	"fmt"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/knockout/api/graphql/model"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/useridentity"
	"github.com/woocoos/knockout/security"
)

// Service 企业目录服务管理
type Service struct {
	Client *ent.Client
}

// EnableOrganization 开启组织目录
func (s *Service) EnableOrganization(ctx context.Context, input model.EnableDirectoryInput) (*ent.Org, error) {
	client := ent.FromContext(ctx)
	uid := security.UserIDFromContext(ctx)

	exist, err := client.Org.Query().Where(org.OwnerID(uid)).Exist(entcache.Evict(ctx))
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("directory service has enable")
	}
	bulk := make([]*ent.OrgCreate, 1)
	bulk[0] = client.Org.Create().SetOwnerID(uid).SetName(input.Name).SetDomain(input.Domain).
		SetKind(org.KindRoot).SetStatus(typex.SimpleStatusActive)
	os, err := client.Org.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return os[0], err
}

// CreateOrganization 创建组织目录,基于根目录创建
func (s *Service) CreateOrganization(ctx context.Context, input ent.CreateOrgInput) (*ent.Org, error) {
	if input.ParentID == 0 {
		return nil, fmt.Errorf("parent id is required")
	}
	return s.Client.Org.Create().SetInput(input).SetKind(org.KindOrganization).Save(ctx)
}

// UpdateOrganization 更新组织目录
//
// - 如果更新的组织目录的管理账号，那么指向的用户必须是账户类型用户
func (s *Service) UpdateOrganization(ctx context.Context, id int, input ent.UpdateOrgInput) (*ent.Org, error) {
	client := ent.FromContext(ctx)
	if input.OwnerID != nil {
		u := client.User.Query().Where(user.ID(*input.OwnerID)).Select(user.FieldUserType).FirstX(ctx)
		if u.UserType != user.UserTypeAccount {
			return nil, fmt.Errorf("owner must be account")
		}
	}
	return client.Org.UpdateOneID(id).SetInput(input).Save(ctx)
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

	_, err = client.OrgUser.Create().SetOrgID(orgId).SetUserID(us.ID).SetDisplayName(us.DisplayName).Save(ctx)
	if err != nil {
		return nil, err
	}
	return us, nil
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

// DeleteApp 删除应用
//
// 该应用没有被关联才可以删除: 1. 没有组织目录关联 2. 没有用户关联
func (s *Service) DeleteApp() (bool, error) {
	panic("implement me")
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

// AddOrganizationUser 将用户加入组织目录
func (s *Service) AddOrganizationUser(ctx context.Context, orgID int, userID int, displayName *string) error {
	client := ent.FromContext(ctx)
	usr := ent.FromContext(ctx).User.GetX(ctx, userID)
	if usr.UserType == user.UserTypeAccount {
		// 邀请用户
	}
	if displayName == nil {
		displayName = &usr.DisplayName
	}
	return client.OrgUser.Create().SetOrgID(orgID).SetUserID(userID).SetDisplayName(*displayName).Exec(ctx)
}
