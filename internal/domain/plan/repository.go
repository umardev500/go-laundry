package plan

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	// AddPermissions attaches a list of permissions to a plan.
	AddPermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error

	// GetByID retrieves a plan by its ID and the provided filter.
	GetByID(ctx context.Context, id uuid.UUID, filter *PlanFilter) (*Plan, error)

	// List retrieves all plans based on the provided filter.
	List(ctx context.Context, filter *PlanFilter) ([]*Plan, error)

	// RemovePermissions detaches a list of permission from a plan.
	RemovePermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error

	// ReplacePermissions replaces all permission attached to a plan.
	ReplacePermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error
}
