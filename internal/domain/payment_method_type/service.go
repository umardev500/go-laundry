package paymentmethodtype

import (
	"context"

	"github.com/google/uuid"
)

// Service defines business logic for managing PaymentMethodType entities.
// The service layer applies validation, defaults, and business rules
// before delegating to the repository.
type Service interface {
	// Create registers a new payment method type.
	//
	// - If status is nil, defaults to "active".
	// - Ensures domain-level validations before persisting.
	Create(ctx context.Context, payload *Create) (*PaymentMethodType, error)

	// Update modifies an existing payment method type.
	//
	// - id: identifies the entity to update.
	// - payload: only non-nil fields will be updated.
	// - Returns updated entity on success.
	Update(ctx context.Context, id uuid.UUID, payload *Update) (*PaymentMethodType, error)

	// Delete permanently removes a payment method type.
	//
	// - id: identifies the entity to delete.
	// - This is a hard delete; after deletion, the record cannot be restored.
	Delete(ctx context.Context, id uuid.UUID) error

	// GetByID retrieves a payment method type by ID.
	//
	// - filter may be applied to narrow results (e.g., by status).
	// - Returns error if not found.
	GetByID(ctx context.Context, id uuid.UUID, filter *Filter) (*PaymentMethodType, error)

	// List returns payment method types matching the filter.
	//
	// - Supports filtering by status (active/inactive).
	// - Service may enforce additional business rules (e.g., hide inactive types for customers).
	List(ctx context.Context, filter *Filter) ([]*PaymentMethodType, error)
}
