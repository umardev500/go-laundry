package payment

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	// Create inserts a new payment
	Create(ctx context.Context, payload *PaymentCreate) (*Payment, error)

	// GetByID retrieves a payment by its ID
	GetByID(ctx context.Context, id uuid.UUID, filter *PaymentFilter) (*Payment, error)

	// List retrieves all payments based on the provider filter
	List(ctx context.Context, filter *PaymentFilter, tenantID *uuid.UUID) ([]*Payment, error)

	// Update updates a payment
	Update(ctx context.Context, payload *PaymentUpdate, userID, id uuid.UUID, TenantID *uuid.UUID) (*Payment, error)
}
