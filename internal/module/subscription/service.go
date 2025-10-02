package subscription

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/payment"
	"github.com/umardev500/go-laundry/internal/domain/plan"
	"github.com/umardev500/go-laundry/internal/domain/subscription"
)

type serviceImpl struct {
	repo           subscription.Repository
	planService    plan.Service
	paymentService payment.Service
	client         *db.Client
}

// Activate implements subscription.Service.
func (s *serviceImpl) Activate(ctx context.Context, id, userID uuid.UUID) (*subscription.Subscription, error) {

	sub, err := s.repo.GetByID(ctx, id, &subscription.Filter{
		IncludePlan:    true,
		IncludePayment: true,
	})
	if err != nil {
		return nil, err
	}

	startDate := time.Now()
	duration := sub.Plan.Duration

	var payload *subscription.SubscriptionUpdate = &subscription.SubscriptionUpdate{
		Status: func() *subscription.SubscriptionStatus {
			status := subscription.SubscriptionStatusActive
			return &status
		}(),
		StartDate: func() *time.Time {
			return &startDate
		}(),
		EndDate: func() *time.Time {
			now := startDate.AddDate(0, 0, *duration)
			return &now
		}(),
	}

	// Set paid completed for payment
	paymentPayload := payment.PaymentUpdate{
		Status: func() *payment.Status {
			status := payment.Completed
			return &status
		}(),
		PaidAt: func() *time.Time {
			now := time.Now()
			return &now
		}(),
	}

	var subscriptonUpdated *subscription.Subscription

	err = s.client.WithTransaction(ctx, func(ctx context.Context) error {
		paymentInfo := sub.Payment

		_, err = s.paymentService.Update(ctx, &paymentPayload, paymentInfo.ID, userID, sub.TenantID)
		if err != nil {
			return err
		}

		subscriptonUpdated, err = s.repo.Update(ctx, payload, id, userID)

		return err
	})

	if err != nil {
		return nil, err
	}

	subscriptonUpdated, err = s.repo.GetByID(ctx, id, &subscription.Filter{
		IncludePlan:    true,
		IncludePayment: true,
		IncludeTenant:  true,
	})
	if err != nil {
		return nil, err
	}

	return subscriptonUpdated, nil
}

// Create implements subscription.Service.
func (s *serviceImpl) Create(
	ctx context.Context,
	userID uuid.UUID,
	payload *subscription.SubscriptionCreate,
) (*subscription.Subscription, error) {
	var sub *subscription.Subscription
	planData, err := s.planService.GetByID(ctx, payload.PlanID, &plan.Filter{
		IncludePermissions: false,
		IncludeDeleted:     false,
	})
	if err != nil {
		return nil, err
	}

	err = s.client.WithTransaction(ctx, func(ctx context.Context) error {

		var paymentStatus payment.Status = payment.Pending

		// If the selected plan has a price of 0 (treated as the "free" plan),
		// we immediately activate the subscription by setting:
		//	- Status	-> Active
		//	- StartDate	-> current time
		//	- EndDate	-> one day from now
		if *planData.Price == 0 {
			paymentStatus = payment.Completed

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

		sub, err = s.repo.Create(ctx, payload)
		if err != nil {
			return err
		}

		// Create payment
		pymnt, err := s.createPayment(
			ctx,
			userID,
			payload.TenantID,
			sub.ID,
			payload.PaymentMethodID,
			*planData.Price,
			paymentStatus,
		)
		if err != nil {
			return err
		}

		sub.Payment = &payment.Payment{
			ID:            pymnt.ID,
			UserID:        pymnt.UserID,
			TenantID:      pymnt.TenantID,
			ReferenceID:   pymnt.ReferenceID,
			ReferenceType: pymnt.ReferenceType,
			Amount:        pymnt.Amount,
			Currency:      pymnt.Currency,
			Status:        pymnt.Status,
			PaidAt:        pymnt.PaidAt,
			CreatedAt:     pymnt.CreatedAt,
			UpdatedAt:     pymnt.UpdatedAt,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return sub, nil

}

// List implements subscription.Service.
func (s *serviceImpl) List(ctx context.Context, filter *subscription.Filter) ([]*subscription.Subscription, error) {
	return s.repo.List(ctx, filter)
}

func (s *serviceImpl) createPayment(
	ctx context.Context,
	userID uuid.UUID,
	tenantID uuid.UUID,
	subID uuid.UUID,
	paymentMethodID uuid.UUID,
	amount float64,
	status payment.Status,
) (*payment.Payment, error) {
	var paymentCreate = payment.PaymentCreate{
		UserID: userID,
		TenantID: func() *uuid.UUID {
			if tenantID == uuid.Nil {
				return nil
			}

			return &tenantID
		}(),
		ReferenceID:     subID,
		ReferenceType:   payment.Subscription,
		PaymentMethodID: paymentMethodID,
		Amount:          amount,
		Currency:        payment.IDR,
		Status:          status,
		PaidAt: func() *time.Time {
			if status == payment.Completed {
				now := time.Now()
				return &now
			}

			return nil
		}(),
	}

	return s.paymentService.Create(ctx, &paymentCreate)
}

func NewService(
	repo subscription.Repository,
	planService plan.Service,
	paymentService payment.Service,
	client *db.Client,
) subscription.Service {
	return &serviceImpl{
		repo:           repo,
		planService:    planService,
		paymentService: paymentService,
		client:         client,
	}
}
