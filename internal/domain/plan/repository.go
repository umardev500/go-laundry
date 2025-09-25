package plan

import (
	"context"
)

type Repository interface {
	// List retrieves all plans based on the provided filter.
	// Soft-deleted plans are excluded
	List(ctx context.Context, filter *PlanFilter) ([]*Plan, error)
}
