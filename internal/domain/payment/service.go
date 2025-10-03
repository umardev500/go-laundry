package payment

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"
)

type Service interface {
	// Create inserts a new payment
	Create(ctx context.Context, payload *PaymentCreate) (*Payment, error)

	// GetByID retrieves a payment by its ID
	GetByID(ctx context.Context, id uuid.UUID, filter *Filter) (*Payment, error)

	// List retrieves all payments based on the provider filter
	List(ctx context.Context, filter *Filter, tenantID *uuid.UUID) (*types.PageResult[Payment], error)

	// Update updates a payment
	Update(ctx context.Context, payload *PaymentUpdate, userID, id uuid.UUID, TenantID *uuid.UUID) (*Payment, error)
}
