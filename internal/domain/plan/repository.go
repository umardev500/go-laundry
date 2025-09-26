package plan

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	// List retrieves all plans based on the provided filter.
	// Soft-deleted plans are excluded
	List(ctx context.Context, filter *PlanFilter) ([]*Plan, error)

	// GetByID retrieves a plan by its ID.
	// Returns an error if not found or sofet-deleted
	GetByID(ctx context.Context, id uuid.UUID) (*Plan, error)
}
