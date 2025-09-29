package paymentmethod

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	// Create registers a new payment method with metadata.
	// TenantID must not be nil (payment methods always belong to a tenant).
	Create(ctx context.Context, payload *Create) (*PaymentMethod, error)

	// GetByID retrieves a payment method by ID (scoped by tenantID if provided) and filter.
	GetByID(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID, filter *Filter) (*PaymentMethod, error)

	// List retrieves payment methods based on the filter (scoped by tenantID if provided).
	List(ctx context.Context, tenantID *uuid.UUID, filter *Filter) ([]*PaymentMethod, error)

	// Update modifies a payment method’s metadata (tenant scoped if tenantID != nil).
	Update(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID, payload *Update) (*PaymentMethod, error)

	// Delete removes a payment method permanently (tenant scoped if tenantID != nil).
	Delete(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID) error
}
