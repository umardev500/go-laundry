package plan

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/plan"
)

type serviceImpl struct {
	repo plan.Repository
}

// GetByID implements plan.Service.
func (s *serviceImpl) GetByID(ctx context.Context, id uuid.UUID) (*plan.Plan, error) {
	return s.repo.GetByID(ctx, id)
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
