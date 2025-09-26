package plan

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	// AddPermissions attaches a list of permissions to a plan.
	AddPermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error

	// GetByID retrieves a plan by its ID.
	// Returns an error if not found or sofet-deleted
	GetByID(ctx context.Context, id uuid.UUID) (*Plan, error)

	// List retrieves all plans based on the provided filter.
	// Soft-deleted plans are excluded
	List(ctx context.Context, filter *PlanFilter) ([]*Plan, error)

	// RemovePermissions detaches a list of permission from a plan.
	RemovePermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error

	// ReplacePermissions replaces all permission attached to a plan.
	ReplacePermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error
}
