package subscription

import (
	"context"

	"github.com/google/uuid"
)

// Repository defines the interface for interacting with Subscription entities.
type Repository interface {
	// GetByID retrieves a subscription by its ID
	GetByID(ctx context.Context, id uuid.UUID, filter *SubscriptionFilter) (*Subscription, error)

	// List retrieves all subscriptions
	// Return a slice of subscriptions pointers and any error encountered
	List(ctx context.Context, filter *SubscriptionFilter) ([]*Subscription, error)

	// Create inserts a new subscription from the given payload.
	Create(ctx context.Context, payload *SubscriptionCreate) (*Subscription, error)

	// Update updates an existing subscription
	Update(ctx context.Context, payload *SubscriptionUpdate, id, userID uuid.UUID) (*Subscription, error)
}
