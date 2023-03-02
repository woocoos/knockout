package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/graph/generated"
	"github.com/woocoos/knockout/service/resource"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type options struct {
	client   *ent.Client
	resource *resource.Service
}

type Option func(*options)

func RegisterEntClient(c *ent.Client) Option {
	return func(o *options) {
		o.client = c
	}
}

func RegistryService(svcs ...any) Option {
	return func(o *options) {
		for _, svc := range svcs {
			switch svc.(type) {
			case *resource.Service:
				o.resource = svc.(*resource.Service)
			}
		}
	}
}

type Resolver struct {
	client   *ent.Client
	resource *resource.Service
}

// NewSchema creates a graphql executable schema.
func NewSchema(opts ...Option) graphql.ExecutableSchema {
	opt := &options{}
	for _, o := range opts {
		o(opt)
	}
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			client:   opt.client,
			resource: opt.resource,
		},
	})
}
