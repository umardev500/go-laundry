package subscription

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	// List retrieves all subscriptions
	List(ctx context.Context, filter *SubscriptionFilter) ([]*Subscription, error)

	// Create insertrs a new subscription
	Create(ctx context.Context, userID uuid.UUID, payload *SubscriptionCreate) (*Subscription, error)

	// Activate sets a subscription as active for the tenant
	Activate(ctx context.Context, id, userID uuid.UUID) (*Subscription, error)
}
