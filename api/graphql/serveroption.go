package graphql

import (
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout/ent"
)

type ServerOption func(srv *ServerOptions)

func WithPortalDB(db *ent.Client) ServerOption {
	return func(srv *ServerOptions) {
		srv.portalDB = db
	}
}

func WithCasbinDB(db *casbinent.Client) ServerOption {
	return func(srv *ServerOptions) {
		srv.casbinDB = db
	}
}

func WithKOSdk(kosdk *api.SDK) ServerOption {
	return func(srv *ServerOptions) {
		srv.kosdk = kosdk
	}
}
