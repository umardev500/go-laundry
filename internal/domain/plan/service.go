package plan

import "context"

type Service interface {
	// List retrieves all plans
	List(ctx context.Context, filter *PlanFilter) ([]*Plan, error)
}
