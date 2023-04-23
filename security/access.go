package security

import (
	"errors"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/authz"
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
func GrantPolicy(rules []*types.PolicyRule, principal, domain string, principalKind permission.PrincipalKind) error {
	authorizer := authz.DefaultAuthorization
	role := principal
	switch principalKind {
	case permission.PrincipalKindUser:
		role = "r_" + principal
		_, err := authorizer.Enforcer.AddRoleForUserInDomain(principal, role, domain)
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
			pls = append(pls, []string{role, domain, action, "read", rule.Effect.String()})
		}
	}
	if len(pls) > 0 {
		_, err := authorizer.Enforcer.AddPolicies(pls)
		if err != nil {
			return err
		}
		return authorizer.Enforcer.SavePolicy()
	}
	return fmt.Errorf("no policy to grant")
}

func GrantByPermission(permission *ent.Permission, domain string) error {
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
func RevokePolicy(rules []*types.PolicyRule, principal, domain string, perm permission.PrincipalKind) error {
	authorizer := authz.DefaultAuthorization
	role := principal
	switch perm {
	case permission.PrincipalKindUser:
		role = "r_" + principal
		_, err := authorizer.Enforcer.DeleteRoleForUserInDomain(principal, role, domain)
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
			pls = append(pls, []string{role, domain, action, "read", "allow"})
		}
	}
	_, err := authorizer.Enforcer.RemovePolicies(pls)
	if err != nil {
		return err
	}
	return authorizer.Enforcer.SavePolicy()
}

func RevokeByPermission(perm *ent.Permission, domain string) error {
	principal := strconv.Itoa(perm.RoleID)
	switch perm.PrincipalKind {
	case permission.UserInverseTable:
		principal = strconv.Itoa(perm.UserID)
	}
	return RevokePolicy(perm.Edges.OrgPolicy.Rules, principal, domain, perm.PrincipalKind)
}

func GrantRoleForUser(userID, roleID int, domain string) error {
	authorizer := authz.DefaultAuthorization
	_, err := authorizer.Enforcer.AddRoleForUserInDomain(strconv.Itoa(userID), strconv.Itoa(roleID), domain)
	if err != nil {
		return err
	}
	return authorizer.Enforcer.SavePolicy()
}

func RevokeGroupForUser(userID, roleID int, domain string) error {
	authorizer := authz.DefaultAuthorization
	_, err := authorizer.Enforcer.DeleteRoleForUserInDomain(strconv.Itoa(userID), strconv.Itoa(roleID), domain)
	if err != nil {
		return err
	}
	return authorizer.Enforcer.SavePolicy()
}
