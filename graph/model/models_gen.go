// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/contrib/entgql"
)

type CreateOrganizationAccountInput struct {
	// 账号名称,组织内不可重复
	DisplayName string `json:"displayName"`
	// 账号登录名(含默认域名)
	PrincipalName string `json:"principalName"`
	// 邮箱
	Email string `json:"email"`
	// 所属组织ID
	OrgID int `json:"orgId"`
}

type EnableDirectoryInput struct {
	// 域名
	Domain string `json:"domain"`
	Name   string `json:"name"`
}

type GrantInput struct {
	Principal string `json:"principal"`
	OrgScope  int    `json:"orgScope"`
	PolicyID  int    `json:"policyID"`
}

// 消息协议
type Message struct {
	ID int `json:"id"`
	// 消息类型
	Topic string `json:"topic"`
	// 消息内容
	Body string `json:"body"`
	// 消息创建时间
	CreatedAt time.Time `json:"createdAt"`
	// 消息发送时间
	SentAt time.Time `json:"sentAt"`
}

// Ordering options for OrganizationRole connections
type OrganizationRoleOrder struct {
	// The ordering direction.
	Direction entgql.OrderDirection `json:"direction"`
	// The field by which to order OrganizationRoles.
	Field OrganizationRoleOrderField `json:"field"`
}

// Ordering options for OrganizationUser connections
type OrganizationUserOrder struct {
	// The ordering direction.
	Direction entgql.OrderDirection `json:"direction"`
	// The field by which to order OrganizationUsers.
	Field OrganizationUserOrderField `json:"field"`
}

// Properties by which OrganizationRole connections can be ordered.
type OrganizationRoleOrderField string

const (
	OrganizationRoleOrderFieldCreatedAt OrganizationRoleOrderField = "createdAt"
)

var AllOrganizationRoleOrderField = []OrganizationRoleOrderField{
	OrganizationRoleOrderFieldCreatedAt,
}

func (e OrganizationRoleOrderField) IsValid() bool {
	switch e {
	case OrganizationRoleOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e OrganizationRoleOrderField) String() string {
	return string(e)
}

func (e *OrganizationRoleOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrganizationRoleOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrganizationRoleOrderField", str)
	}
	return nil
}

func (e OrganizationRoleOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which OrganizationUser connections can be ordered.
type OrganizationUserOrderField string

const (
	OrganizationUserOrderFieldCreatedAt OrganizationUserOrderField = "createdAt"
)

var AllOrganizationUserOrderField = []OrganizationUserOrderField{
	OrganizationUserOrderFieldCreatedAt,
}

func (e OrganizationUserOrderField) IsValid() bool {
	switch e {
	case OrganizationUserOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e OrganizationUserOrderField) String() string {
	return string(e)
}

func (e *OrganizationUserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrganizationUserOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrganizationUserOrderField", str)
	}
	return nil
}

func (e OrganizationUserOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
