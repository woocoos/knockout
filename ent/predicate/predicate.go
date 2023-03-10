// Code generated by ent, DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// App is the predicate function for app builders.
type App func(*sql.Selector)

// AppAction is the predicate function for appaction builders.
type AppAction func(*sql.Selector)

// AppMenu is the predicate function for appmenu builders.
type AppMenu func(*sql.Selector)

// AppPolicy is the predicate function for apppolicy builders.
type AppPolicy func(*sql.Selector)

// AppRes is the predicate function for appres builders.
type AppRes func(*sql.Selector)

// AppRole is the predicate function for approle builders.
type AppRole func(*sql.Selector)

// AppRolePolicy is the predicate function for approlepolicy builders.
type AppRolePolicy func(*sql.Selector)

// Org is the predicate function for org builders.
type Org func(*sql.Selector)

// OrgApp is the predicate function for orgapp builders.
type OrgApp func(*sql.Selector)

// OrgPolicy is the predicate function for orgpolicy builders.
type OrgPolicy func(*sql.Selector)

// OrgRole is the predicate function for orgrole builders.
type OrgRole func(*sql.Selector)

// OrgUser is the predicate function for orguser builders.
type OrgUser func(*sql.Selector)

// Permission is the predicate function for permission builders.
type Permission func(*sql.Selector)

// User is the predicate function for user builders.
type User func(*sql.Selector)

// UserDevice is the predicate function for userdevice builders.
type UserDevice func(*sql.Selector)

// UserIdentity is the predicate function for useridentity builders.
type UserIdentity func(*sql.Selector)

// UserLoginProfile is the predicate function for userloginprofile builders.
type UserLoginProfile func(*sql.Selector)

// UserPassword is the predicate function for userpassword builders.
type UserPassword func(*sql.Selector)
