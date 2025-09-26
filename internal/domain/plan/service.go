package plan

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	// List retrieves all plans
	List(ctx context.Context, filter *PlanFilter) ([]*Plan, error)

	// GetByID retrieves a plan by its ID
	GetByID(ctx context.Context, id uuid.UUID) (*Plan, error)
}
