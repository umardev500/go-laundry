package subscription

import (
	"context"
	"time"

	"github.com/umardev500/go-laundry/internal/domain/plan"
	"github.com/umardev500/go-laundry/internal/domain/subscription"
)

type serviceImpl struct {
	repo        subscription.Repository
	planService plan.Service
}

// Create implements subscription.Service.
func (s *serviceImpl) Create(ctx context.Context, payload *subscription.SubscriptionCreate) (*subscription.Subscription, error) {
	planData, err := s.planService.GetByID(ctx, payload.PlanID)
	if err != nil {
		return nil, err
	}

	// If the selected plan has a price of 0 (treated as the "free" plan),
	// we immediately activate the subscription by setting:
	//	- Status	-> Active
	//	- StartDate	-> current time
	//	- EndDate	-> one day from now
	if *planData.Price == 0 {
		payload.Status = func() *subscription.SubscriptionStatus {
			status := subscription.SubscriptionStatusActive
			return &status
		}()
		payload.StartDate = func() *time.Time {
			now := time.Now()
			return &now
		}()
		payload.EndDate = func() *time.Time {
			now := time.Now().AddDate(0, 0, 1)
			return &now
		}()
	}

	return s.repo.Create(ctx, payload)
}

// List implements subscription.Service.
func (s *serviceImpl) List(ctx context.Context) ([]*subscription.Subscription, error) {
	return s.repo.List(ctx)
}

func NewService(repo subscription.Repository, planService plan.Service) subscription.Service {
	return &serviceImpl{
		repo:        repo,
		planService: planService,
	}
}
