package hook

import (
	"context"
	"entgo.io/ent"
	"github.com/gin-gonic/gin"
	"github.com/woocoos/knockout-go/pkg/fmterr"
	gen "github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/hook"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/user"
	"strconv"
)

// OrgPathHook 创建组织路径hook.
func OrgPathHook() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.OrgFunc(func(ctx context.Context, mutation *gen.OrgMutation) (gen.Value, error) {
				if _, ok := mutation.Path(); ok {
					return next.Mutate(ctx, mutation)
				}
				// 采用IntID，优先执行才有数据库id
				value, err := next.Mutate(ctx, mutation)
				if err != nil {
					return nil, err
				}
				ov := value.(*gen.Org)
				if pid, ok := mutation.ParentID(); ok {
					id, _ := mutation.ID()
					code := strconv.FormatInt(int64(id), 36)
					path := code
					if pid != 0 {
						parentPath := ""
						prow, err := mutation.Client().Org.Query().Where(org.ID(pid)).
							Select(org.FieldPath).Only(ctx)
						if err != nil {
							if !gen.IsNotFound(err) {
								return nil, err
							}
							parentPath = ""
						} else {
							parentPath = prow.Path + "/"
						}
						path = parentPath + path
					}
					ov.Path = path

					if ov.Code == "" {
						_, err = mutation.Client().ExecContext(ctx, "UPDATE "+org.Table+" SET path=?, code=? WHERE id=?", path, code, id)
						ov.Code = code
					} else {
						_, err = mutation.Client().ExecContext(ctx, "UPDATE "+org.Table+" SET path=? WHERE id=?", path, id)
					}
					if err != nil {
						return nil, err
					}
				}
				return value, nil
			})
		}, ent.OpCreate)
}

// OrgCheckDeleteHook 检查组织是否有子节点.
func OrgCheckDeleteHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.OrgFunc(func(ctx context.Context, mutation *gen.OrgMutation) (gen.Value, error) {
			if mutation.Op() == ent.OpDeleteOne {
				if id, ok := mutation.ID(); ok {
					count, err := mutation.Client().Org.Query().Where(
						org.ParentID(id),
					).Count(ctx)
					if err != nil {
						return nil, err
					}
					if count > 0 {
						return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "organization has children")
					}
				}
			}
			return next.Mutate(ctx, mutation)
		})
	}
}

// OrgOwnerCheckHook 检查owner是否为账户类型.
func OrgOwnerCheckHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.OrgFunc(func(ctx context.Context, m *gen.OrgMutation) (gen.Value, error) {
			if uid, ok := m.OwnerID(); ok {
				usr, err := m.Client().User.Get(ctx, uid)
				if err != nil {
					return nil, err
				}
				if usr.UserType != user.UserTypeAccount {
					return nil, fmterr.Newf(uint64(gin.ErrorTypePublic), "owner must be account: %s", usr.DisplayName)
				}
				m.SetKind(org.KindRoot)
				if m.Op().Is(ent.OpUpdateOne) {
					id, _ := m.ID()
					has, _ := m.Client().OrgUser.Query().Where(orguser.UserID(uid), orguser.OrgID(id)).Exist(ctx)
					if !has {
						err = m.Client().OrgUser.Create().SetOrgID(id).SetUserID(uid).SetDisplayName(usr.DisplayName).
							Exec(ctx)
						if err != nil {
							return nil, err
						}
					}
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}