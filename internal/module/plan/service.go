package plan

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/plan"
)

type serviceImpl struct {
	repo plan.Repository
}

// AddPermissions implements plan.Service.
func (s *serviceImpl) AddPermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error {
	return s.repo.AddPermissions(ctx, planID, permissionIDs)
}

// RemovePermissions implements plan.Service.
func (s *serviceImpl) RemovePermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error {
	return s.repo.RemovePermissions(ctx, planID, permissionIDs)
}

// ReplacePermissions implements plan.Service.
func (s *serviceImpl) ReplacePermissions(ctx context.Context, planID uuid.UUID, permissionIDs []uuid.UUID) error {
	return s.repo.ReplacePermissions(ctx, planID, permissionIDs)
}

// GetByID implements plan.Service.
func (s *serviceImpl) GetByID(ctx context.Context, id uuid.UUID, filter *plan.Filter) (*plan.Plan, error) {
	return s.repo.GetByID(ctx, id, filter)
}

// List implements plan.Service.
func (s *serviceImpl) List(ctx context.Context, filter *plan.Filter) ([]*plan.Plan, error) {
	return s.repo.List(ctx, filter)
}

func NewService(repo plan.Repository) plan.Service {
	return &serviceImpl{
		repo: repo,
	}
}
