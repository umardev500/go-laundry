package subscription

import "context"

// Repository defines the interface for interacting with Subscription entities.
type Repository interface {

	// List retrieves all subscriptions
	// Return a slice of subscriptions pointers and any error encountered
	List(ctx context.Context) ([]*Subscription, error)

	// Create inserts a new subscription from the given payload.
	Create(ctx context.Context, payload *SubscriptionCreate) (*Subscription, error)
}
