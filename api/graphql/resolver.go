package graphql

import (
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/resource"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client   *ent.Client
	Resource *resource.Service
}
