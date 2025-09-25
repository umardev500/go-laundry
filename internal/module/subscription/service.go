package subscription

import (
	"context"

	"github.com/umardev500/go-laundry/internal/domain/subscription"
)

type serviceImpl struct {
	repo subscription.Repository
}

// List implements subscription.Service.
func (s *serviceImpl) List(ctx context.Context) ([]*subscription.Subscription, error) {
	return s.repo.List(ctx)
}

func NewService(repo subscription.Repository) subscription.Service {
	return &serviceImpl{
		repo: repo,
	}
}
