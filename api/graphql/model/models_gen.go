// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/contrib/entgql"
)

// Ordering options for AppRolePolicy connections
type AppRolePolicyOrder struct {
	// The ordering direction.
	Direction entgql.OrderDirection `json:"direction"`
	// The field by which to order AppRolePolicies.
	Field AppRolePolicyOrderField `json:"field"`
}

type AssignRoleUserInput struct {
	// 授权类型为角色或用户组的ID
	OrgRoleID int `json:"orgRoleID"`
	UserID    int `json:"userID"`
	// 生效开始时间
	StartAt *time.Time `json:"startAt"`
	// 生效结束时间
	EndAt *time.Time `json:"endAt"`
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

type Mfa struct {
	Secret  string `json:"secret"`
	Account string `json:"account"`
}

// Ordering options for OrgRoleUser connections
type OrgRoleUserOrder struct {
	// The ordering direction.
	Direction entgql.OrderDirection `json:"direction"`
	// The field by which to order OrgRoleUsers.
	Field OrgRoleUserOrderField `json:"field"`
}

// Ordering options for OrgUser connections
type OrgUserOrder struct {
	// The ordering direction.
	Direction entgql.OrderDirection `json:"direction"`
	// The field by which to order OrgUsers.
	Field OrgUserOrderField `json:"field"`
}

// Properties by which AppRolePolicy connections can be ordered.
type AppRolePolicyOrderField string

const (
	AppRolePolicyOrderFieldCreatedAt AppRolePolicyOrderField = "createdAt"
)

var AllAppRolePolicyOrderField = []AppRolePolicyOrderField{
	AppRolePolicyOrderFieldCreatedAt,
}

func (e AppRolePolicyOrderField) IsValid() bool {
	switch e {
	case AppRolePolicyOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e AppRolePolicyOrderField) String() string {
	return string(e)
}

func (e *AppRolePolicyOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AppRolePolicyOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AppRolePolicyOrderField", str)
	}
	return nil
}

func (e AppRolePolicyOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which OrgRoleUser connections can be ordered.
type OrgRoleUserOrderField string

const (
	OrgRoleUserOrderFieldCreatedAt OrgRoleUserOrderField = "createdAt"
)

var AllOrgRoleUserOrderField = []OrgRoleUserOrderField{
	OrgRoleUserOrderFieldCreatedAt,
}

func (e OrgRoleUserOrderField) IsValid() bool {
	switch e {
	case OrgRoleUserOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e OrgRoleUserOrderField) String() string {
	return string(e)
}

func (e *OrgRoleUserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgRoleUserOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgRoleUserOrderField", str)
	}
	return nil
}

func (e OrgRoleUserOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which OrgUser connections can be ordered.
type OrgUserOrderField string

const (
	OrgUserOrderFieldCreatedAt OrgUserOrderField = "createdAt"
)

var AllOrgUserOrderField = []OrgUserOrderField{
	OrgUserOrderFieldCreatedAt,
}

func (e OrgUserOrderField) IsValid() bool {
	switch e {
	case OrgUserOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e OrgUserOrderField) String() string {
	return string(e)
}

func (e *OrgUserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrgUserOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrgUserOrderField", str)
	}
	return nil
}

func (e OrgUserOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// 树操作类型
type TreeAction string

const (
	// 作为子节点
	TreeActionChild TreeAction = "child"
	// 上移
	TreeActionUp TreeAction = "up"
	// 下移
	TreeActionDown TreeAction = "down"
)

var AllTreeAction = []TreeAction{
	TreeActionChild,
	TreeActionUp,
	TreeActionDown,
}

func (e TreeAction) IsValid() bool {
	switch e {
	case TreeActionChild, TreeActionUp, TreeActionDown:
		return true
	}
	return false
}

func (e TreeAction) String() string {
	return string(e)
}

func (e *TreeAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TreeAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TreeAction", str)
	}
	return nil
}

func (e TreeAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
