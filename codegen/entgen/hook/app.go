package hook

import (
	"context"
	"entgo.io/ent"
	"github.com/gin-gonic/gin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appmenu"
	"github.com/woocoos/knockout/ent/apppolicy"
	"github.com/woocoos/knockout/ent/approle"
	"github.com/woocoos/knockout/ent/approlepolicy"
	"github.com/woocoos/knockout/ent/hook"
	"github.com/woocoos/knockout/ent/orgapp"
)

// AppDeleteHook 删除应用前的清理hook.
func AppDeleteHook() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.AppFunc(func(ctx context.Context, m *gen.AppMutation) (gen.Value, error) {
			id, _ := m.ID()
			client := m.Client()
			apl, err := client.App.Get(ctx, id)
			if err != nil {
				return nil, err
			}
			if apl.OrgPrivate != true {
				has, err := client.OrgApp.Query().Where(orgapp.HasAppWith(app.ID(id))).Exist(ctx)
				if err != nil {
					return nil, err
				}
				if has {
					return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "app has been associated with org")
				}
			}
			if _, err = client.AppAction.Delete().Where(appaction.AppID(id)).Exec(ctx); err != nil {
				return nil, err
			}
			if _, err = client.AppMenu.Delete().Where(appmenu.AppID(id)).Exec(ctx); err != nil {
				return nil, err
			}
			if _, err = client.AppPolicy.Delete().Where(apppolicy.AppID(id)).Exec(ctx); err != nil {
				return nil, err
			}
			if _, err = client.AppRole.Delete().Where(approle.AppID(id)).Exec(ctx); err != nil {
				return nil, err
			}
			if _, err = client.AppRolePolicy.Delete().Where(approlepolicy.AppID(id)).Exec(ctx); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}, ent.OpDeleteOne)
}