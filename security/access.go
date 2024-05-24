package security

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/tsingsun/woocoo/pkg/security"
	authz "github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/permission"
	"strconv"
)

// GrantPolicy 同步授权信息到cashbin.
//
//  1. 将授权信息同步到cashbin中.
//     如果授权主体为用户,则将用户作为特定角色.
//  2. 将授权信息同步到redis中
func GrantPolicy(rules []*types.PolicyRule, principal string, domain int, principalKind permission.PrincipalKind) error {
	tenant := strconv.Itoa(domain)
	authorizer := security.DefaultAuthorizer.(*authz.Authorizer)
	role := principal
	switch principalKind {
	case permission.PrincipalKindUser:
		role = "r_" + principal
		_, err := authorizer.Enforcer.AddRoleForUserInDomain(principal, role, tenant)
		if err != nil {
			return err
		}
	case permission.PrincipalKindRole:
	default:
		return errors.New("invalid principal kind")
	}
	pls := make([][]string, 0, len(rules))
	for _, rule := range rules {
		for _, action := range rule.Actions {
			p := []string{role, tenant, action, "read", rule.Effect.String()}
			if !authorizer.Enforcer.HasPolicy(p) {
				pls = append(pls, p)
			}
		}
		for _, resource := range rule.Resources {
			p := []string{role, tenant, resource, "read", rule.Effect.String()}
			if !authorizer.Enforcer.HasPolicy(p) {
				pls = append(pls, p)
			}
		}
	}
	if len(pls) > 0 {
		_, err := authorizer.Enforcer.AddPoliciesEx(pls)
		// 清除缓存
		_ = authorizer.Enforcer.(*casbin.CachedEnforcer).InvalidateCache()
		return err
	}
	return nil
}

func GrantByPermission(permission *ent.Permission, domain int) error {
	principal := strconv.Itoa(permission.RoleID)
	switch permission.PrincipalKind {
	case "user":
		principal = strconv.Itoa(permission.UserID)
	}
	return GrantPolicy(permission.Edges.OrgPolicy.Rules, principal, domain, permission.PrincipalKind)
}

// RevokePolicy 同步授权信息到cashbin.
//
//  1. 将授权信息同步到cashbin中.
//     如果授权主体为用户,则将用户作为特定角色.
//  2. 将授权信息同步到redis中
func RevokePolicy(rules []*types.PolicyRule, principal string, domain int, perm permission.PrincipalKind) error {
	tenant := strconv.Itoa(domain)
	authorizer := security.DefaultAuthorizer.(*authz.Authorizer)
	role := principal
	switch perm {
	case permission.PrincipalKindUser:
		// 需判断没有授权策略时才允许移除,暂时不移除
		role = "r_" + principal
		//_, err := authorizer.Enforcer.DeleteRoleForUserInDomain(principal, role, domain)
		//if err != nil {
		//	return err
		//}
	case permission.PrincipalKindRole:
	default:
		return errors.New("invalid principal kind")
	}
	pls := make([][]string, 0, len(rules))
	for _, rule := range rules {
		for _, action := range rule.Actions {
			pls = append(pls, []string{role, tenant, action, "read", rule.Effect.String()})
		}
		for _, resource := range rule.Resources {
			pls = append(pls, []string{role, tenant, resource, "read", rule.Effect.String()})
		}
	}

	_, err := authorizer.Enforcer.RemovePolicies(pls)
	if err != nil {
		return err
	}
	// 清除缓存
	err = authorizer.Enforcer.(*casbin.CachedEnforcer).InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}

func RevokeByPermission(perm *ent.Permission, domain int) error {
	principal := strconv.Itoa(perm.RoleID)
	switch perm.PrincipalKind {
	case permission.UserInverseTable:
		principal = strconv.Itoa(perm.UserID)
	}
	return RevokePolicy(perm.Edges.OrgPolicy.Rules, principal, domain, perm.PrincipalKind)
}

func GrantRoleForUser(userID, roleID int, domain int) error {
	authorizer := security.DefaultAuthorizer.(*authz.Authorizer)
	_, err := authorizer.Enforcer.AddRoleForUserInDomain(strconv.Itoa(userID), strconv.Itoa(roleID), strconv.Itoa(domain))
	// 清除缓存
	_ = authorizer.Enforcer.(*casbin.CachedEnforcer).InvalidateCache()
	return err
}

func RevokeGroupForUser(userID, roleID int, domain int) error {
	authorizer := security.DefaultAuthorizer.(*authz.Authorizer)
	_, err := authorizer.Enforcer.DeleteRoleForUserInDomain(strconv.Itoa(userID), strconv.Itoa(roleID), strconv.Itoa(domain))
	// 清除缓存
	_ = authorizer.Enforcer.(*casbin.CachedEnforcer).InvalidateCache()
	return err
}

func GetUserPermissions(userID int, domain int) [][]string {
	authorizer := security.DefaultAuthorizer.(*authz.Authorizer)
	return authorizer.Enforcer.GetPermissionsForUserInDomain(strconv.Itoa(userID), strconv.Itoa(domain))
}

func CheckUserPermission(rvals ...interface{}) (bool, error) {
	authorizer := security.DefaultAuthorizer.(*authz.Authorizer)
	return authorizer.Enforcer.Enforce(rvals...)
}
