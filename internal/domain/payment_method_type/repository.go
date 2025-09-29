package paymentmethodtype

import (
	"context"

	"github.com/google/uuid"
)

// Repository defines persistence operations for PaymentMethodType entities.
type Repository interface {
	// Create persists a new payment method type record.
	Create(ctx context.Context, payload *Create) (*PaymentMethodType, error)

	// Update modifies an existing payment method type.
	Update(ctx context.Context, id uuid.UUID, payload *Update) (*PaymentMethodType, error)

	// Delete removes a payment method type from the system (hard delete).
	Delete(ctx context.Context, id uuid.UUID) error

	// GetByID retrieves a single payment method type by ID and filter.
	GetByID(ctx context.Context, id uuid.UUID, filter *Filter) (*PaymentMethodType, error)

	// List retrieves payment method types based on filter criteria.
	List(ctx context.Context, filter *Filter) ([]*PaymentMethodType, error)
}
