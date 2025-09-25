package subscription

import "context"

type Service interface {
	// List retrieves all subscriptions
	List(ctx context.Context) ([]*Subscription, error)
}
