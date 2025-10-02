package payment

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/payment"
)

type serviceImpl struct {
	repo payment.Repository
}

// Create implements payment.Service.
func (s *serviceImpl) Create(ctx context.Context, payload *payment.PaymentCreate) (*payment.Payment, error) {
	return s.repo.Create(ctx, payload)
}

// GetByID implements payment.Service.
func (s *serviceImpl) GetByID(ctx context.Context, id uuid.UUID, filter *payment.PaymentFilter) (*payment.Payment, error) {
	return s.repo.GetByID(ctx, id, filter, nil)
}

// List implements payment.Service.
func (s *serviceImpl) List(ctx context.Context, filter *payment.PaymentFilter, tenantID *uuid.UUID) ([]*payment.Payment, error) {
	return s.repo.List(ctx, filter, tenantID)
}

// Update implements payment.Service.
func (s *serviceImpl) Update(ctx context.Context, payload *payment.PaymentUpdate, id, userID uuid.UUID, TenantID *uuid.UUID) (*payment.Payment, error) {
	return s.repo.Update(ctx, payload, id, userID, TenantID)
}

func NewService(repo payment.Repository) payment.Service {
	return &serviceImpl{
		repo: repo,
	}
}
