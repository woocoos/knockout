package hook

import (
	"entgo.io/ent"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/appdictitem"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/apppolicyview"
	"github.com/woocoos/knockout/ent/hook"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/region"
)

// RegisterAllHooks 注册所有schema的hook到ent client.
// 在程序入口调用此函数来启用所有hook.
func RegisterAllHooks(client *gen.Client) {
	// App hooks
	client.App.Use(AppDeleteHook())

	// AppDict hooks
	client.AppDict.Use(AppDictDeleteHook())

	// AppDictItem hooks
	client.AppDictItem.Use(InitDisplaySortHookEx(appdictitem.Table, appdictitem.FieldDictID))
	client.AppDictItem.Use(AppDictItemCodeUniqueHook())
	client.AppDictItem.Use(AppDictItemRefCodeHook())
	client.AppDictItem.Use(AppDictItemOrgIDHook())
	client.AppDictItem.Use(AppDictItemDeleteStatusHook())

	// AppMenu hooks
	client.AppMenu.Use(InitDisplaySortHook(appmenu.Table))

	// AppPolicy hooks
	client.AppPolicy.Use(AppPolicyRulesHook())

	// AppPolicyView hooks
	client.AppPolicyView.Use(AppPolicyViewPathHook())
	client.AppPolicyView.Use(InitDisplaySortHook(apppolicyview.Table))

	// Org hooks
	client.Org.Use(OrgPathHook())
	client.Org.Use(InitDisplaySortHook(org.Table))
	client.Org.Use(hook.On(OrgCheckDeleteHook(), ent.OpDeleteOne))
	client.Org.Use(hook.On(OrgOwnerCheckHook(), ent.OpCreate|ent.OpUpdateOne|ent.OpUpdate))

	// OrgPolicy hooks
	client.OrgPolicy.Use(OrgPolicyRulesHook())

	// Quota hooks
	client.Quota.Use(QuotaTenantOrOrgIDHook())

	// Region hooks
	client.Region.Use(InitDisplaySortHook(region.Table))

	// UserAddr hooks
	client.UserAddr.Use(UserAddrContactUniqueHook())

	// UserDevice hooks
	client.UserDevice.Use(UserDeviceQuotaHook())

	// OrgUserPreference hooks
	client.OrgUserPreference.Use(OrgUserPreferenceQuotaHook())

	// UserIdentity hooks
	client.UserIdentity.Use(UserIdentityCodeUniqueHook())
}

