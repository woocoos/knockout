// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppColumns holds the columns for the "app" table.
	AppColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 45},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 45},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"web", "native", "server"}},
		{Name: "redirect_uri", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "app_key", Type: field.TypeString, Nullable: true},
		{Name: "app_secret", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "scopes", Type: field.TypeString, Nullable: true, Size: 500},
		{Name: "token_validity", Type: field.TypeInt32, Nullable: true},
		{Name: "refresh_token_validity", Type: field.TypeInt32, Nullable: true},
		{Name: "logo", Type: field.TypeString, Nullable: true},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing"}, Default: "active"},
	}
	// AppTable holds the schema information for the "app" table.
	AppTable = &schema.Table{
		Name:       "app",
		Columns:    AppColumns,
		PrimaryKey: []*schema.Column{AppColumns[0]},
	}
	// AppActionColumns holds the columns for the "app_action" table.
	AppActionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"restful", "graphql", "rpc"}},
		{Name: "method", Type: field.TypeEnum, Enums: []string{"read", "write", "list"}},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "app_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// AppActionTable holds the schema information for the "app_action" table.
	AppActionTable = &schema.Table{
		Name:       "app_action",
		Columns:    AppActionColumns,
		PrimaryKey: []*schema.Column{AppActionColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_action_app_actions",
				Columns:    []*schema.Column{AppActionColumns[9]},
				RefColumns: []*schema.Column{AppColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AppMenuColumns holds the columns for the "app_menu" table.
	AppMenuColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "parent_id", Type: field.TypeInt},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"dir", "menu"}},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 100},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "display_sort", Type: field.TypeInt32, Nullable: true},
		{Name: "app_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "action_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// AppMenuTable holds the schema information for the "app_menu" table.
	AppMenuTable = &schema.Table{
		Name:       "app_menu",
		Columns:    AppMenuColumns,
		PrimaryKey: []*schema.Column{AppMenuColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_menu_app_menus",
				Columns:    []*schema.Column{AppMenuColumns[10]},
				RefColumns: []*schema.Column{AppColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "app_menu_app_action_menus",
				Columns:    []*schema.Column{AppMenuColumns[11]},
				RefColumns: []*schema.Column{AppActionColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AppPolicyColumns holds the columns for the "app_policy" table.
	AppPolicyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "comments", Type: field.TypeString},
		{Name: "rules", Type: field.TypeJSON},
		{Name: "version", Type: field.TypeString},
		{Name: "auto_grant", Type: field.TypeBool, Default: false},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing"}, Default: "active"},
		{Name: "app_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// AppPolicyTable holds the schema information for the "app_policy" table.
	AppPolicyTable = &schema.Table{
		Name:       "app_policy",
		Columns:    AppPolicyColumns,
		PrimaryKey: []*schema.Column{AppPolicyColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_policy_app_policies",
				Columns:    []*schema.Column{AppPolicyColumns[11]},
				RefColumns: []*schema.Column{AppColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AppResColumns holds the columns for the "app_res" table.
	AppResColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type_name", Type: field.TypeString},
		{Name: "arn_pattern", Type: field.TypeString},
		{Name: "app_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "app_action_resources", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// AppResTable holds the schema information for the "app_res" table.
	AppResTable = &schema.Table{
		Name:       "app_res",
		Columns:    AppResColumns,
		PrimaryKey: []*schema.Column{AppResColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_res_app_resources",
				Columns:    []*schema.Column{AppResColumns[8]},
				RefColumns: []*schema.Column{AppColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "app_res_app_action_resources",
				Columns:    []*schema.Column{AppResColumns[9]},
				RefColumns: []*schema.Column{AppActionColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AppRoleColumns holds the columns for the "app_role" table.
	AppRoleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "auto_grant", Type: field.TypeBool, Default: false},
		{Name: "editable", Type: field.TypeBool, Default: false},
		{Name: "app_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// AppRoleTable holds the schema information for the "app_role" table.
	AppRoleTable = &schema.Table{
		Name:       "app_role",
		Columns:    AppRoleColumns,
		PrimaryKey: []*schema.Column{AppRoleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_role_app_roles",
				Columns:    []*schema.Column{AppRoleColumns[9]},
				RefColumns: []*schema.Column{AppColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AppRolePolicyColumns holds the columns for the "app_role_policy" table.
	AppRolePolicyColumns = []*schema.Column{
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "role_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "policy_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// AppRolePolicyTable holds the schema information for the "app_role_policy" table.
	AppRolePolicyTable = &schema.Table{
		Name:       "app_role_policy",
		Columns:    AppRolePolicyColumns,
		PrimaryKey: []*schema.Column{AppRolePolicyColumns[4], AppRolePolicyColumns[5]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "app_role_policy_app_role_role",
				Columns:    []*schema.Column{AppRolePolicyColumns[4]},
				RefColumns: []*schema.Column{AppRoleColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "app_role_policy_app_policy_policy",
				Columns:    []*schema.Column{AppRolePolicyColumns[5]},
				RefColumns: []*schema.Column{AppPolicyColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrganizationColumns holds the columns for the "organization" table.
	OrganizationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "kind", Type: field.TypeEnum, Nullable: true, Enums: []string{"root", "org"}},
		{Name: "domain", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "code", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "name", Type: field.TypeString, Size: 100},
		{Name: "profile", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing"}},
		{Name: "path", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "display_sort", Type: field.TypeInt32, Nullable: true},
		{Name: "country_code", Type: field.TypeString, Nullable: true, Size: 10},
		{Name: "timezone", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true, Default: 0, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// OrganizationTable holds the schema information for the "organization" table.
	OrganizationTable = &schema.Table{
		Name:       "organization",
		Columns:    OrganizationColumns,
		PrimaryKey: []*schema.Column{OrganizationColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_organization_children",
				Columns:    []*schema.Column{OrganizationColumns[16]},
				RefColumns: []*schema.Column{OrganizationColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "organization_user_owner",
				Columns:    []*schema.Column{OrganizationColumns[17]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrganizationAppColumns holds the columns for the "organization_app" table.
	OrganizationAppColumns = []*schema.Column{
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "app_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "org_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// OrganizationAppTable holds the schema information for the "organization_app" table.
	OrganizationAppTable = &schema.Table{
		Name:       "organization_app",
		Columns:    OrganizationAppColumns,
		PrimaryKey: []*schema.Column{OrganizationAppColumns[5], OrganizationAppColumns[4]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_app_app_app",
				Columns:    []*schema.Column{OrganizationAppColumns[4]},
				RefColumns: []*schema.Column{AppColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "organization_app_organization_organization",
				Columns:    []*schema.Column{OrganizationAppColumns[5]},
				RefColumns: []*schema.Column{OrganizationColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrganizationRoleColumns holds the columns for the "organization_role" table.
	OrganizationRoleColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"group", "role"}},
		{Name: "name", Type: field.TypeString},
		{Name: "app_role_id", Type: field.TypeInt, Nullable: true},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "org_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// OrganizationRoleTable holds the schema information for the "organization_role" table.
	OrganizationRoleTable = &schema.Table{
		Name:       "organization_role",
		Columns:    OrganizationRoleColumns,
		PrimaryKey: []*schema.Column{OrganizationRoleColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_role_organization_rolesAndGroups",
				Columns:    []*schema.Column{OrganizationRoleColumns[9]},
				RefColumns: []*schema.Column{OrganizationColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// OrganizationUserColumns holds the columns for the "organization_user" table.
	OrganizationUserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "display_name", Type: field.TypeString},
		{Name: "org_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "user_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// OrganizationUserTable holds the schema information for the "organization_user" table.
	OrganizationUserTable = &schema.Table{
		Name:       "organization_user",
		Columns:    OrganizationUserColumns,
		PrimaryKey: []*schema.Column{OrganizationUserColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_user_organization_organization",
				Columns:    []*schema.Column{OrganizationUserColumns[6]},
				RefColumns: []*schema.Column{OrganizationColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "organization_user_user_user",
				Columns:    []*schema.Column{OrganizationUserColumns[7]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "organizationuser_org_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{OrganizationUserColumns[6], OrganizationUserColumns[7]},
			},
		},
	}
	// PermissionColumns holds the columns for the "permission" table.
	PermissionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "principal_kind", Type: field.TypeEnum, Enums: []string{"user", "role"}},
		{Name: "role_id", Type: field.TypeInt, Nullable: true},
		{Name: "org_policy_id", Type: field.TypeInt},
		{Name: "start_at", Type: field.TypeTime, Nullable: true},
		{Name: "end_at", Type: field.TypeTime, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing"}},
		{Name: "org_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "user_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// PermissionTable holds the schema information for the "permission" table.
	PermissionTable = &schema.Table{
		Name:       "permission",
		Columns:    PermissionColumns,
		PrimaryKey: []*schema.Column{PermissionColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "permission_organization_permissions",
				Columns:    []*schema.Column{PermissionColumns[11]},
				RefColumns: []*schema.Column{OrganizationColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "permission_user_user",
				Columns:    []*schema.Column{PermissionColumns[12]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrganizationPolicyColumns holds the columns for the "organization_policy" table.
	OrganizationPolicyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "app_id", Type: field.TypeInt, Nullable: true},
		{Name: "app_policy_id", Type: field.TypeInt, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "comments", Type: field.TypeString},
		{Name: "rules", Type: field.TypeJSON},
		{Name: "org_id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// OrganizationPolicyTable holds the schema information for the "organization_policy" table.
	OrganizationPolicyTable = &schema.Table{
		Name:       "organization_policy",
		Columns:    OrganizationPolicyColumns,
		PrimaryKey: []*schema.Column{OrganizationPolicyColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_policy_organization_policies",
				Columns:    []*schema.Column{OrganizationPolicyColumns[10]},
				RefColumns: []*schema.Column{OrganizationColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, SchemaType: map[string]string{"mysql": "bigint"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "principal_name", Type: field.TypeString, Unique: true},
		{Name: "display_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "mobile", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "user_type", Type: field.TypeEnum, Enums: []string{"account", "member"}},
		{Name: "creation_type", Type: field.TypeEnum, Enums: []string{"invitation", "register", "manual"}},
		{Name: "register_ip", Type: field.TypeString, Size: 45},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing"}},
		{Name: "comments", Type: field.TypeString, Nullable: true},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
	}
	// UserDeviceColumns holds the columns for the "user_device" table.
	UserDeviceColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "device_uid", Type: field.TypeString, Size: 64},
		{Name: "device_name", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "system_name", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "system_version", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "app_version", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "device_model", Type: field.TypeString, Nullable: true, Size: 45},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing"}},
		{Name: "comments", Type: field.TypeString, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// UserDeviceTable holds the schema information for the "user_device" table.
	UserDeviceTable = &schema.Table{
		Name:       "user_device",
		Columns:    UserDeviceColumns,
		PrimaryKey: []*schema.Column{UserDeviceColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_device_user_devices",
				Columns:    []*schema.Column{UserDeviceColumns[13]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserIdentityColumns holds the columns for the "user_identity" table.
	UserIdentityColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "kind", Type: field.TypeEnum, Enums: []string{"name", "email", "phone", "wechat", "qq"}},
		{Name: "code", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "code_extend", Type: field.TypeString, Nullable: true},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing"}},
		{Name: "user_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// UserIdentityTable holds the schema information for the "user_identity" table.
	UserIdentityTable = &schema.Table{
		Name:       "user_identity",
		Columns:    UserIdentityColumns,
		PrimaryKey: []*schema.Column{UserIdentityColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_identity_user_identities",
				Columns:    []*schema.Column{UserIdentityColumns[9]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserLoginProfileColumns holds the columns for the "user_login_profile" table.
	UserLoginProfileColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "last_login_ip", Type: field.TypeString, Nullable: true},
		{Name: "last_login_at", Type: field.TypeTime, Nullable: true},
		{Name: "can_login", Type: field.TypeBool, Nullable: true},
		{Name: "set_kind", Type: field.TypeEnum, Enums: []string{"keep", "customer", "auto"}},
		{Name: "password_reset", Type: field.TypeBool, Nullable: true},
		{Name: "verify_device", Type: field.TypeBool},
		{Name: "mfa_enabled", Type: field.TypeBool, Nullable: true},
		{Name: "mfa_secret", Type: field.TypeString, Nullable: true, Size: 100},
		{Name: "mfa_status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing"}},
		{Name: "user_id", Type: field.TypeInt, Unique: true, Nullable: true, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// UserLoginProfileTable holds the schema information for the "user_login_profile" table.
	UserLoginProfileTable = &schema.Table{
		Name:       "user_login_profile",
		Columns:    UserLoginProfileColumns,
		PrimaryKey: []*schema.Column{UserLoginProfileColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_login_profile_user_login_profile",
				Columns:    []*schema.Column{UserLoginProfileColumns[14]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserPasswordColumns holds the columns for the "user_password" table.
	UserPasswordColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"mysql": "int"}},
		{Name: "created_by", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeInt, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "scene", Type: field.TypeEnum, Nullable: true, Enums: []string{"login"}},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "salt", Type: field.TypeString, Size: 45},
		{Name: "status", Type: field.TypeEnum, Nullable: true, Enums: []string{"active", "inactive", "processing"}},
		{Name: "memo", Type: field.TypeString, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "bigint"}},
	}
	// UserPasswordTable holds the schema information for the "user_password" table.
	UserPasswordTable = &schema.Table{
		Name:       "user_password",
		Columns:    UserPasswordColumns,
		PrimaryKey: []*schema.Column{UserPasswordColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_password_user_passwords",
				Columns:    []*schema.Column{UserPasswordColumns[10]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppTable,
		AppActionTable,
		AppMenuTable,
		AppPolicyTable,
		AppResTable,
		AppRoleTable,
		AppRolePolicyTable,
		OrganizationTable,
		OrganizationAppTable,
		OrganizationRoleTable,
		OrganizationUserTable,
		PermissionTable,
		OrganizationPolicyTable,
		UserTable,
		UserDeviceTable,
		UserIdentityTable,
		UserLoginProfileTable,
		UserPasswordTable,
	}
)

func init() {
	AppTable.Annotation = &entsql.Annotation{
		Table: "app",
	}
	AppActionTable.ForeignKeys[0].RefTable = AppTable
	AppActionTable.Annotation = &entsql.Annotation{
		Table: "app_action",
	}
	AppMenuTable.ForeignKeys[0].RefTable = AppTable
	AppMenuTable.ForeignKeys[1].RefTable = AppActionTable
	AppMenuTable.Annotation = &entsql.Annotation{
		Table: "app_menu",
	}
	AppPolicyTable.ForeignKeys[0].RefTable = AppTable
	AppPolicyTable.Annotation = &entsql.Annotation{
		Table: "app_policy",
	}
	AppResTable.ForeignKeys[0].RefTable = AppTable
	AppResTable.ForeignKeys[1].RefTable = AppActionTable
	AppResTable.Annotation = &entsql.Annotation{
		Table: "app_res",
	}
	AppRoleTable.ForeignKeys[0].RefTable = AppTable
	AppRoleTable.Annotation = &entsql.Annotation{
		Table: "app_role",
	}
	AppRolePolicyTable.ForeignKeys[0].RefTable = AppRoleTable
	AppRolePolicyTable.ForeignKeys[1].RefTable = AppPolicyTable
	AppRolePolicyTable.Annotation = &entsql.Annotation{
		Table: "app_role_policy",
	}
	OrganizationTable.ForeignKeys[0].RefTable = OrganizationTable
	OrganizationTable.ForeignKeys[1].RefTable = UserTable
	OrganizationTable.Annotation = &entsql.Annotation{
		Table: "organization",
	}
	OrganizationAppTable.ForeignKeys[0].RefTable = AppTable
	OrganizationAppTable.ForeignKeys[1].RefTable = OrganizationTable
	OrganizationAppTable.Annotation = &entsql.Annotation{
		Table: "organization_app",
	}
	OrganizationRoleTable.ForeignKeys[0].RefTable = OrganizationTable
	OrganizationRoleTable.Annotation = &entsql.Annotation{
		Table: "organization_role",
	}
	OrganizationUserTable.ForeignKeys[0].RefTable = OrganizationTable
	OrganizationUserTable.ForeignKeys[1].RefTable = UserTable
	OrganizationUserTable.Annotation = &entsql.Annotation{
		Table: "organization_user",
	}
	PermissionTable.ForeignKeys[0].RefTable = OrganizationTable
	PermissionTable.ForeignKeys[1].RefTable = UserTable
	PermissionTable.Annotation = &entsql.Annotation{
		Table: "permission",
	}
	OrganizationPolicyTable.ForeignKeys[0].RefTable = OrganizationTable
	OrganizationPolicyTable.Annotation = &entsql.Annotation{
		Table: "organization_policy",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
	UserDeviceTable.ForeignKeys[0].RefTable = UserTable
	UserDeviceTable.Annotation = &entsql.Annotation{
		Table: "user_device",
	}
	UserIdentityTable.ForeignKeys[0].RefTable = UserTable
	UserIdentityTable.Annotation = &entsql.Annotation{
		Table: "user_identity",
	}
	UserLoginProfileTable.ForeignKeys[0].RefTable = UserTable
	UserLoginProfileTable.Annotation = &entsql.Annotation{
		Table: "user_login_profile",
	}
	UserPasswordTable.ForeignKeys[0].RefTable = UserTable
	UserPasswordTable.Annotation = &entsql.Annotation{
		Table: "user_password",
	}
}
