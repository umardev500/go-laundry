package services

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/umardev500/go-laundry/internal/domain/services"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"
)

type serviceImpl struct {
	repo domain.Repository
}

// Ensure serviceImpl implements the domain Service interface
var _ domain.Service = (*serviceImpl)(nil)

// NewService creates a new services service
func NewService(repo domain.Repository) domain.Service {
	return &serviceImpl{repo: repo}
}

// Create a new service
func (s *serviceImpl) Create(ctx context.Context, payload *domain.Create, tenantID *uuid.UUID) (*domain.Services, error) {
	return s.repo.Create(ctx, payload, tenantID)
}

// GetByID retrieves a service by its ID (optionally tenant scoped)
func (s *serviceImpl) GetByID(ctx context.Context, id uuid.UUID, tenantID *uuid.UUID) (*domain.Services, error) {
	return s.repo.GetByID(ctx, id, tenantID)
}

// List retrieves services matching the filter (optionally tenant scoped)
func (s *serviceImpl) List(ctx context.Context, filter *domain.Filter, tenantID *uuid.UUID) (*types.PageResult[domain.Services], error) {
	filter = filter.WithDefaults()
	result, err := s.repo.List(ctx, filter, tenantID)
	if err != nil {
		return nil, err
	}

	// build pagination response
	paginatedResult := utils.Paginate(result.Data, result.Total, filter.Offset, filter.Limit)
	return paginatedResult, nil
}

// Update modifies an existing service
func (s *serviceImpl) Update(ctx context.Context, id uuid.UUID, payload *domain.Update, tenantID *uuid.UUID) (*domain.Services, error) {
	return s.repo.Update(ctx, id, payload, tenantID)
}

// Delete removes a service permanently
func (s *serviceImpl) Delete(ctx context.Context, id uuid.UUID, tenantID *uuid.UUID) error {
	return s.repo.Delete(ctx, id, tenantID)
}
