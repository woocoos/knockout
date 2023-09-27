package graphql

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/knockout/api/graphql/generated"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/service/resource"
)

type Option func(*Resolver)

type Resolver struct {
	client   *ent.Client
	resource *resource.Service
}

func WithClient(client *ent.Client) Option {
	return func(r *Resolver) {
		r.client = client
	}
}

func WithResource(resource *resource.Service) Option {
	return func(r *Resolver) {
		r.resource = resource
	}
}

func NewResolver(opt ...Option) *Resolver {
	r := &Resolver{}
	for _, option := range opt {
		option(r)
	}
	return r
}

// NewSchema creates a graphql executable schema.
func NewSchema(opts ...Option) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: NewResolver(opts...),
	})
}
