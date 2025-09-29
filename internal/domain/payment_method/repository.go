package paymentmethod

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	// Create persists a new payment method with arbitrary metadata.
	Create(ctx context.Context, payload *Create) (*PaymentMethod, error)

	// GetByID retrieves a payment method by ID (tenant scoped if tenantID != nil) and filter.
	GetByID(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID, filter *Filter) (*PaymentMethod, error)

	// List retrieves all payment methods matching the filter (tenant scoped if tenantID != nil).
	List(ctx context.Context, tenantID *uuid.UUID, filter *Filter) ([]*PaymentMethod, error)

	// Update modifies an existing payment method’s metadata.
	Update(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID, payload *Update) (*PaymentMethod, error)

	// Delete permanently removes a payment method by ID (tenant scoped if tenantID != nil).
	Delete(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID) error
}
