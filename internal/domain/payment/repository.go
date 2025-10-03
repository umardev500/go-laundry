package payment

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"
)

type Repository interface {
	// Create inserts a new payment
	Create(ctx context.Context, payload *PaymentCreate) (*Payment, error)

	// GetByID retrieves a payment by its ID
	// filter is used to filter the payments
	// If tenantID is provided, only payments for that tenant will be returned
	// Return a payment pointer and any error encountered
	GetByID(ctx context.Context, id uuid.UUID, filter *Filter, tenantID *uuid.UUID) (*Payment, error)

	// List retrieves all payments based on the provider filter
	// If tenantID is provided, only payments for that tenant will be returned
	// Return a list of payments and any error encountered
	List(ctx context.Context, filter *Filter, tenantID *uuid.UUID) (*types.PageData[Payment], error)

	// Update updates a payment
	// If tenantID is provided, only payments for that tenant will be returned
	// Return a payment pointer and any error encountered
	Update(ctx context.Context, payload *PaymentUpdate, id, userID uuid.UUID, TenantID *uuid.UUID) (*Payment, error)
}
