package plan

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	// AddPermissions attaches permissions to a plan
	AddPermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error

	// GetByID retrieves a plan by its ID
	GetByID(ctx context.Context, id uuid.UUID) (*Plan, error)

	// List retrieves all plans
	List(ctx context.Context, filter *PlanFilter) ([]*Plan, error)

	// RemovePermissions detaches permissions from a plan
	RemovePermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error

	// ReplacePermissions clears existing permissions and replaces them with the provided list
	ReplacePermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error
}
