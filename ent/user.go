// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/ent/country"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/userloginprofile"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int `json:"created_by,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int `json:"updated_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// 登陆名称
	PrincipalName string `json:"principal_name,omitempty"`
	// 显示名
	DisplayName string `json:"display_name,omitempty"`
	// 用户类型
	UserType user.UserType `json:"user_type,omitempty"`
	// 创建类型,邀请，注册,手工创建
	CreationType user.CreationType `json:"creation_type,omitempty"`
	// 注册时IP
	RegisterIP *string `json:"register_ip,omitempty"`
	// 状态
	Status typex.SimpleStatus `json:"status,omitempty"`
	// 备注
	Comments string `json:"comments,omitempty"`
	// 头像地址
	Avatar string `json:"avatar,omitempty"`
	// 性别
	Gender user.Gender `json:"gender,omitempty"`
	// 国籍
	CitizenshipID *int `json:"citizenship_id,omitempty"`
	// 名字
	FirstName string `json:"first_name,omitempty"`
	// 中间名
	MiddleName string `json:"middle_name,omitempty"`
	// 姓氏
	LastName string `json:"last_name,omitempty"`
	// 语言
	Lang string `json:"lang,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// 用户身份标识
	Identities []*UserIdentity `json:"identities,omitempty"`
	// 登陆设置
	LoginProfile *UserLoginProfile `json:"login_profile,omitempty"`
	// 用户密码
	Passwords []*UserPassword `json:"passwords,omitempty"`
	// 用户设备
	Devices []*UserDevice `json:"devices,omitempty"`
	// 用户所属组织
	Orgs []*Org `json:"orgs,omitempty"`
	// 用户权限
	Permissions []*Permission `json:"permissions,omitempty"`
	// 用户AccessKey
	OauthClients []*OauthClient `json:"oauth_clients,omitempty"`
	// 用户联系信息
	Addresses []*UserAddr `json:"addresses,omitempty"`
	// 国籍信息
	Citizenship *Country `json:"citizenship,omitempty"`
	// OrgUser holds the value of the org_user edge.
	OrgUser []*OrgUser `json:"org_user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [10]bool
	// totalCount holds the count of the edges above.
	totalCount [7]map[string]int

	namedIdentities   map[string][]*UserIdentity
	namedPasswords    map[string][]*UserPassword
	namedDevices      map[string][]*UserDevice
	namedOrgs         map[string][]*Org
	namedPermissions  map[string][]*Permission
	namedOauthClients map[string][]*OauthClient
	namedAddresses    map[string][]*UserAddr
	namedOrgUser      map[string][]*OrgUser
}

// IdentitiesOrErr returns the Identities value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) IdentitiesOrErr() ([]*UserIdentity, error) {
	if e.loadedTypes[0] {
		return e.Identities, nil
	}
	return nil, &NotLoadedError{edge: "identities"}
}

// LoginProfileOrErr returns the LoginProfile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) LoginProfileOrErr() (*UserLoginProfile, error) {
	if e.LoginProfile != nil {
		return e.LoginProfile, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: userloginprofile.Label}
	}
	return nil, &NotLoadedError{edge: "login_profile"}
}

// PasswordsOrErr returns the Passwords value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PasswordsOrErr() ([]*UserPassword, error) {
	if e.loadedTypes[2] {
		return e.Passwords, nil
	}
	return nil, &NotLoadedError{edge: "passwords"}
}

// DevicesOrErr returns the Devices value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DevicesOrErr() ([]*UserDevice, error) {
	if e.loadedTypes[3] {
		return e.Devices, nil
	}
	return nil, &NotLoadedError{edge: "devices"}
}

// OrgsOrErr returns the Orgs value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OrgsOrErr() ([]*Org, error) {
	if e.loadedTypes[4] {
		return e.Orgs, nil
	}
	return nil, &NotLoadedError{edge: "orgs"}
}

// PermissionsOrErr returns the Permissions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PermissionsOrErr() ([]*Permission, error) {
	if e.loadedTypes[5] {
		return e.Permissions, nil
	}
	return nil, &NotLoadedError{edge: "permissions"}
}

// OauthClientsOrErr returns the OauthClients value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OauthClientsOrErr() ([]*OauthClient, error) {
	if e.loadedTypes[6] {
		return e.OauthClients, nil
	}
	return nil, &NotLoadedError{edge: "oauth_clients"}
}

// AddressesOrErr returns the Addresses value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AddressesOrErr() ([]*UserAddr, error) {
	if e.loadedTypes[7] {
		return e.Addresses, nil
	}
	return nil, &NotLoadedError{edge: "addresses"}
}

// CitizenshipOrErr returns the Citizenship value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CitizenshipOrErr() (*Country, error) {
	if e.Citizenship != nil {
		return e.Citizenship, nil
	} else if e.loadedTypes[8] {
		return nil, &NotFoundError{label: country.Label}
	}
	return nil, &NotLoadedError{edge: "citizenship"}
}

// OrgUserOrErr returns the OrgUser value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OrgUserOrErr() ([]*OrgUser, error) {
	if e.loadedTypes[9] {
		return e.OrgUser, nil
	}
	return nil, &NotLoadedError{edge: "org_user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldCreatedBy, user.FieldUpdatedBy, user.FieldCitizenshipID:
			values[i] = new(sql.NullInt64)
		case user.FieldPrincipalName, user.FieldDisplayName, user.FieldUserType, user.FieldCreationType, user.FieldRegisterIP, user.FieldStatus, user.FieldComments, user.FieldAvatar, user.FieldGender, user.FieldFirstName, user.FieldMiddleName, user.FieldLastName, user.FieldLang:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				u.CreatedBy = int(value.Int64)
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				u.UpdatedBy = int(value.Int64)
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = value.Time
			}
		case user.FieldPrincipalName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field principal_name", values[i])
			} else if value.Valid {
				u.PrincipalName = value.String
			}
		case user.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				u.DisplayName = value.String
			}
		case user.FieldUserType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_type", values[i])
			} else if value.Valid {
				u.UserType = user.UserType(value.String)
			}
		case user.FieldCreationType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field creation_type", values[i])
			} else if value.Valid {
				u.CreationType = user.CreationType(value.String)
			}
		case user.FieldRegisterIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field register_ip", values[i])
			} else if value.Valid {
				u.RegisterIP = new(string)
				*u.RegisterIP = value.String
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = typex.SimpleStatus(value.String)
			}
		case user.FieldComments:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comments", values[i])
			} else if value.Valid {
				u.Comments = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				u.Gender = user.Gender(value.String)
			}
		case user.FieldCitizenshipID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field citizenship_id", values[i])
			} else if value.Valid {
				u.CitizenshipID = new(int)
				*u.CitizenshipID = int(value.Int64)
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldMiddleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field middle_name", values[i])
			} else if value.Valid {
				u.MiddleName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldLang:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lang", values[i])
			} else if value.Valid {
				u.Lang = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryIdentities queries the "identities" edge of the User entity.
func (u *User) QueryIdentities() *UserIdentityQuery {
	return NewUserClient(u.config).QueryIdentities(u)
}

// QueryLoginProfile queries the "login_profile" edge of the User entity.
func (u *User) QueryLoginProfile() *UserLoginProfileQuery {
	return NewUserClient(u.config).QueryLoginProfile(u)
}

// QueryPasswords queries the "passwords" edge of the User entity.
func (u *User) QueryPasswords() *UserPasswordQuery {
	return NewUserClient(u.config).QueryPasswords(u)
}

// QueryDevices queries the "devices" edge of the User entity.
func (u *User) QueryDevices() *UserDeviceQuery {
	return NewUserClient(u.config).QueryDevices(u)
}

// QueryOrgs queries the "orgs" edge of the User entity.
func (u *User) QueryOrgs() *OrgQuery {
	return NewUserClient(u.config).QueryOrgs(u)
}

// QueryPermissions queries the "permissions" edge of the User entity.
func (u *User) QueryPermissions() *PermissionQuery {
	return NewUserClient(u.config).QueryPermissions(u)
}

// QueryOauthClients queries the "oauth_clients" edge of the User entity.
func (u *User) QueryOauthClients() *OauthClientQuery {
	return NewUserClient(u.config).QueryOauthClients(u)
}

// QueryAddresses queries the "addresses" edge of the User entity.
func (u *User) QueryAddresses() *UserAddrQuery {
	return NewUserClient(u.config).QueryAddresses(u)
}

// QueryCitizenship queries the "citizenship" edge of the User entity.
func (u *User) QueryCitizenship() *CountryQuery {
	return NewUserClient(u.config).QueryCitizenship(u)
}

// QueryOrgUser queries the "org_user" edge of the User entity.
func (u *User) QueryOrgUser() *OrgUserQuery {
	return NewUserClient(u.config).QueryOrgUser(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(u.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("principal_name=")
	builder.WriteString(u.PrincipalName)
	builder.WriteString(", ")
	builder.WriteString("display_name=")
	builder.WriteString(u.DisplayName)
	builder.WriteString(", ")
	builder.WriteString("user_type=")
	builder.WriteString(fmt.Sprintf("%v", u.UserType))
	builder.WriteString(", ")
	builder.WriteString("creation_type=")
	builder.WriteString(fmt.Sprintf("%v", u.CreationType))
	builder.WriteString(", ")
	if v := u.RegisterIP; v != nil {
		builder.WriteString("register_ip=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("comments=")
	builder.WriteString(u.Comments)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", u.Gender))
	builder.WriteString(", ")
	if v := u.CitizenshipID; v != nil {
		builder.WriteString("citizenship_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("middle_name=")
	builder.WriteString(u.MiddleName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("lang=")
	builder.WriteString(u.Lang)
	builder.WriteByte(')')
	return builder.String()
}

// NamedIdentities returns the Identities named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedIdentities(name string) ([]*UserIdentity, error) {
	if u.Edges.namedIdentities == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedIdentities[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedIdentities(name string, edges ...*UserIdentity) {
	if u.Edges.namedIdentities == nil {
		u.Edges.namedIdentities = make(map[string][]*UserIdentity)
	}
	if len(edges) == 0 {
		u.Edges.namedIdentities[name] = []*UserIdentity{}
	} else {
		u.Edges.namedIdentities[name] = append(u.Edges.namedIdentities[name], edges...)
	}
}

// NamedPasswords returns the Passwords named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedPasswords(name string) ([]*UserPassword, error) {
	if u.Edges.namedPasswords == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedPasswords[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedPasswords(name string, edges ...*UserPassword) {
	if u.Edges.namedPasswords == nil {
		u.Edges.namedPasswords = make(map[string][]*UserPassword)
	}
	if len(edges) == 0 {
		u.Edges.namedPasswords[name] = []*UserPassword{}
	} else {
		u.Edges.namedPasswords[name] = append(u.Edges.namedPasswords[name], edges...)
	}
}

// NamedDevices returns the Devices named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedDevices(name string) ([]*UserDevice, error) {
	if u.Edges.namedDevices == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedDevices[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedDevices(name string, edges ...*UserDevice) {
	if u.Edges.namedDevices == nil {
		u.Edges.namedDevices = make(map[string][]*UserDevice)
	}
	if len(edges) == 0 {
		u.Edges.namedDevices[name] = []*UserDevice{}
	} else {
		u.Edges.namedDevices[name] = append(u.Edges.namedDevices[name], edges...)
	}
}

// NamedOrgs returns the Orgs named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedOrgs(name string) ([]*Org, error) {
	if u.Edges.namedOrgs == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedOrgs[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedOrgs(name string, edges ...*Org) {
	if u.Edges.namedOrgs == nil {
		u.Edges.namedOrgs = make(map[string][]*Org)
	}
	if len(edges) == 0 {
		u.Edges.namedOrgs[name] = []*Org{}
	} else {
		u.Edges.namedOrgs[name] = append(u.Edges.namedOrgs[name], edges...)
	}
}

// NamedPermissions returns the Permissions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedPermissions(name string) ([]*Permission, error) {
	if u.Edges.namedPermissions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedPermissions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedPermissions(name string, edges ...*Permission) {
	if u.Edges.namedPermissions == nil {
		u.Edges.namedPermissions = make(map[string][]*Permission)
	}
	if len(edges) == 0 {
		u.Edges.namedPermissions[name] = []*Permission{}
	} else {
		u.Edges.namedPermissions[name] = append(u.Edges.namedPermissions[name], edges...)
	}
}

// NamedOauthClients returns the OauthClients named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedOauthClients(name string) ([]*OauthClient, error) {
	if u.Edges.namedOauthClients == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedOauthClients[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedOauthClients(name string, edges ...*OauthClient) {
	if u.Edges.namedOauthClients == nil {
		u.Edges.namedOauthClients = make(map[string][]*OauthClient)
	}
	if len(edges) == 0 {
		u.Edges.namedOauthClients[name] = []*OauthClient{}
	} else {
		u.Edges.namedOauthClients[name] = append(u.Edges.namedOauthClients[name], edges...)
	}
}

// NamedAddresses returns the Addresses named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedAddresses(name string) ([]*UserAddr, error) {
	if u.Edges.namedAddresses == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedAddresses[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedAddresses(name string, edges ...*UserAddr) {
	if u.Edges.namedAddresses == nil {
		u.Edges.namedAddresses = make(map[string][]*UserAddr)
	}
	if len(edges) == 0 {
		u.Edges.namedAddresses[name] = []*UserAddr{}
	} else {
		u.Edges.namedAddresses[name] = append(u.Edges.namedAddresses[name], edges...)
	}
}

// NamedOrgUser returns the OrgUser named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedOrgUser(name string) ([]*OrgUser, error) {
	if u.Edges.namedOrgUser == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedOrgUser[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedOrgUser(name string, edges ...*OrgUser) {
	if u.Edges.namedOrgUser == nil {
		u.Edges.namedOrgUser = make(map[string][]*OrgUser)
	}
	if len(edges) == 0 {
		u.Edges.namedOrgUser[name] = []*OrgUser{}
	} else {
		u.Edges.namedOrgUser[name] = append(u.Edges.namedOrgUser[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User
