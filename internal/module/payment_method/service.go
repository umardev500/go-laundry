package paymentmethod

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	domain "github.com/umardev500/go-laundry/internal/domain/payment_method"
)

type serviceImpl struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) domain.Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) Create(ctx context.Context, payload *domain.Create) (*domain.PaymentMethod, error) {
	if payload.TenantID == uuid.Nil {
		return nil, ErrTenantRequired
	}
	return s.repo.Create(ctx, payload)
}

func (s *serviceImpl) GetByID(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID, filter *domain.Filter) (*domain.PaymentMethod, error) {
	return s.repo.GetByID(ctx, tenantID, id, filter)
}

func (s *serviceImpl) List(ctx context.Context, tenantID *uuid.UUID, filter *domain.Filter) ([]*domain.PaymentMethod, error) {
	return s.repo.List(ctx, tenantID, filter)
}

func (s *serviceImpl) Update(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID, payload *domain.Update) (*domain.PaymentMethod, error) {
	return s.repo.Update(ctx, tenantID, id, payload)
}

func (s *serviceImpl) Delete(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID) error {
	return s.repo.Delete(ctx, tenantID, id)
}

// custom error
var ErrTenantRequired = fmt.Errorf("tenant_id is required for payment method")
