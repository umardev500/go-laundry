package paymentmethodtype

import (
	"context"

	"github.com/google/uuid"
	paymentmethodtype "github.com/umardev500/go-laundry/internal/domain/payment_method_type"
)

type serviceImpl struct {
	repo paymentmethodtype.Repository
}

// NewService constructs a new Service implementation.
func NewService(repo paymentmethodtype.Repository) paymentmethodtype.Service {
	return &serviceImpl{repo: repo}
}

// Create implements Service.
func (s *serviceImpl) Create(ctx context.Context, payload *paymentmethodtype.Create) (*paymentmethodtype.PaymentMethodType, error) {
	// Apply default status if not provided
	if payload.Status == nil {
		defaultStatus := paymentmethodtype.Active
		payload.Status = &defaultStatus
	}

	return s.repo.Create(ctx, payload)
}

// Update implements Service.
func (s *serviceImpl) Update(ctx context.Context, id uuid.UUID, payload *paymentmethodtype.Update) (*paymentmethodtype.PaymentMethodType, error) {
	return s.repo.Update(ctx, id, payload)
}

// Delete implements Service.
func (s *serviceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

// GetByID implements Service.
func (s *serviceImpl) GetByID(ctx context.Context, id uuid.UUID, filter *paymentmethodtype.Filter) (*paymentmethodtype.PaymentMethodType, error) {
	return s.repo.GetByID(ctx, id, filter)
}

// List implements Service.
func (s *serviceImpl) List(ctx context.Context, filter *paymentmethodtype.Filter) ([]*paymentmethodtype.PaymentMethodType, error) {
	return s.repo.List(ctx, filter)
}
