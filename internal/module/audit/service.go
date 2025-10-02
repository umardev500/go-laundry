package audit

import (
	"context"

	"github.com/umardev500/go-laundry/internal/domain/audit"
)

type serviceImpl struct {
	repo audit.Repository
}

// Create implements audit.Service.
func (s *serviceImpl) Create(ctx context.Context, payload *audit.Create) error {
	return s.repo.Create(ctx, payload)
}

func NewService(repo audit.Repository) audit.Service {
	return &serviceImpl{
		repo: repo,
	}
}
