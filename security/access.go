package security

import (
	"errors"
	"github.com/tsingsun/woocoo/pkg/authz"
	"github.com/woocoos/knockout/codegen/entgen/types"
)

// GrantPolicy 同步授权信息到cashbin.
//
//  1. 将授权信息同步到cashbin中.
//     如果授权主体为用户,则将用户作为特定角色.
//  2. 将授权信息同步到redis中
func GrantPolicy(rules []types.PolicyRule, principal, principalKind, domain string) error {
	authorizer := authz.DefaultAuthorization
	role := principal
	switch principalKind {
	case "user":
		role = "r_" + principal
		_, err := authorizer.Enforcer.AddRoleForUserInDomain(principal, role, domain)
		if err != nil {
			return err
		}
	case "role":
	default:
		return errors.New("invalid principal kind")
	}
	pls := make([][]string, 0, len(rules))
	for _, rule := range rules {
		for _, action := range rule.Actions {
			pls = append(pls, []string{role, domain, action, "read", "allow"})
		}
	}
	_, err := authorizer.Enforcer.AddPolicies(pls)
	if err != nil {
		return err
	}
	return authorizer.Enforcer.SavePolicy()
}
