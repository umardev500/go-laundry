package payment

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/payment"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"
)

type serviceImpl struct {
	repo payment.Repository
}

// Create implements payment.Service.
func (s *serviceImpl) Create(ctx context.Context, payload *payment.PaymentCreate) (*payment.Payment, error) {
	return s.repo.Create(ctx, payload)
}

// GetByID implements payment.Service.
func (s *serviceImpl) GetByID(ctx context.Context, id uuid.UUID, filter *payment.Filter) (*payment.Payment, error) {
	return s.repo.GetByID(ctx, id, filter, nil)
}

// List implements payment.Service.
func (s *serviceImpl) List(ctx context.Context, f *payment.Filter, tenantID *uuid.UUID) (*types.PageResult[payment.Payment], error) {
	f = f.WithDefaults()

	result, err := s.repo.List(ctx, f, tenantID)
	if err != nil {
		return nil, err
	}

	paginateResult := utils.Paginate(result.Data, result.Total, f.Offset, f.Limit)

	return paginateResult, nil
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
