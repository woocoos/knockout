package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"time"

	"github.com/woocoos/knockout/graph/generated"
	"github.com/woocoos/knockout/graph/model"
)

// Message is the resolver for the message field.
func (r *subscriptionResolver) Message(ctx context.Context) (<-chan *model.Message, error) {
	ch := make(chan *model.Message)
	go func() {
		for {
			time.Sleep(10 * time.Second)
			select {
			case <-ctx.Done():
				return
			default:
				ch <- &model.Message{
					ID:        1,
					Body:      "Hello World",
					CreatedAt: time.Now(),
					SentAt:    time.Now(),
					Topic:     "msg",
				}
			}
		}
	}()
	return ch, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
