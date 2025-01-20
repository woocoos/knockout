package resource

import (
	"context"
	"fmt"
	"github.com/woocoos/knockout-go/pkg/identity"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/userpasswordpolicy"
)

func (s *Service) UserPasswordPolicy(ctx context.Context) (*ent.UserPasswordPolicy, error) {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if tid == 0 {
		return s.defaultUserPasswordPolicy()
	}
	upp, err := s.Client.UserPasswordPolicy.Query().Where(userpasswordpolicy.TenantID(tid)).Only(ctx)
	if ent.IsNotFound(err) {
		return s.defaultUserPasswordPolicy()
	}
	if err != nil {
		return nil, err
	}
	return upp, nil
}

func (s *Service) defaultUserPasswordPolicy() (*ent.UserPasswordPolicy, error) {
	var upp = ent.UserPasswordPolicy{
		Length:               s.passwordPolicy.Length,
		IncludeElement:       s.passwordPolicy.IncludeElement,
		IncludeChar:          s.passwordPolicy.IncludeChar,
		AllowIncludeUserName: s.passwordPolicy.AllowIncludeUserName,
		InvalidDay:           s.passwordPolicy.InvalidDay,
		InvalidLoginLimit:    s.passwordPolicy.InvalidLoginLimit,
		Retry:                s.passwordPolicy.Retry,
		CaptchaTimes:         s.passwordPolicy.CaptchaTimes,
	}
	return &upp, nil
}

func (s *Service) CreateUserPasswordPolicy(ctx context.Context, orgID int, input ent.CreateUserPasswordPolicyInput) (*ent.UserPasswordPolicy, error) {
	client := s.Client
	// 查询租户是否已有密码策略
	has, err := client.UserPasswordPolicy.Query().Where(userpasswordpolicy.TenantID(orgID)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, fmt.Errorf("密码策略已经存在，不能再次创建")
	}
	input.OrgID = &orgID
	return client.UserPasswordPolicy.Create().SetInput(input).Save(ctx)
}

func (s *Service) UpdateUserPasswordPolicy(ctx context.Context, orgID int, input ent.UpdateUserPasswordPolicyInput) (*ent.UserPasswordPolicy, error) {
	client := s.Client
	// 查询出租户的密码策略
	upp, err := client.UserPasswordPolicy.Query().Where(userpasswordpolicy.TenantID(orgID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).UserPasswordPolicy.UpdateOne(upp).SetInput(input).Save(ctx)
}

func (s *Service) DeleteUserPasswordPolicy(ctx context.Context) (bool, error) {
	client := s.Client
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return false, err
	}
	// 查询出租户的密码策略
	upp, err := client.UserPasswordPolicy.Query().Where(userpasswordpolicy.TenantID(tid)).Only(ctx)
	if err != nil {
		return false, err
	}
	if upp == nil {
		return false, fmt.Errorf("密码策略不存在，删除失败")
	}
	err = client.UserPasswordPolicy.DeleteOne(upp).Exec(ctx)
	return err == nil, err
}
