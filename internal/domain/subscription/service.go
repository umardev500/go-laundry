package subscription

import "context"

type Service interface {
	// List retrieves all subscriptions
	List(ctx context.Context) ([]*Subscription, error)

	// Create insertrs a new subscription
	Create(ctx context.Context, payload *SubscriptionCreate) (*Subscription, error)
}
