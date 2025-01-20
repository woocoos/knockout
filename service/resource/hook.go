package resource

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/woocoos/knockout/ent/intercept"
	"github.com/woocoos/knockout/ent/org"
)

type EntHook struct {
	service *Service
}

func NewEntHook(service *Service) *EntHook {
	return &EntHook{
		service: service,
	}
}

// OrgTraverseFunc 只允许查询当前租户下的org
func (e *EntHook) OrgTraverseFunc() ent.Interceptor {
	var orgin = func(ids []int, w interface{ WhereP(...func(*sql.Selector)) }) {
		if len(ids) == 1 {
			w.WhereP(sql.FieldEQ(org.FieldID, ids[0]))
		} else {
			w.WhereP(sql.FieldIn(org.FieldID, ids...))
		}
	}
	return intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
		ids, err := e.service.GetRefTenants(ctx)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		orgin(ids, q)
		return nil
	})
}
