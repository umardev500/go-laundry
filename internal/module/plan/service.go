package plan

import (
	"context"

	"github.com/umardev500/go-laundry/internal/domain/plan"
)

type serviceImpl struct {
	repo plan.Repository
}

// List implements plan.Service.
func (s *serviceImpl) List(ctx context.Context, filter *plan.PlanFilter) ([]*plan.Plan, error) {
	return s.repo.List(ctx, filter)
}

func NewService(repo plan.Repository) plan.Service {
	return &serviceImpl{
		repo: repo,
	}
}
