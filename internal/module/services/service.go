package services

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
	domain "github.com/umardev500/go-laundry/internal/domain/services"
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
func (s *serviceImpl) Create(ctx *appContext.ScopedContext, payload *domain.Create) (*domain.Services, error) {
	return s.repo.Create(ctx, payload)
}

// GetByID retrieves a service by its ID (optionally tenant scoped)
func (s *serviceImpl) GetByID(ctx *appContext.ScopedContext, id uuid.UUID) (*domain.Services, error) {
	return s.repo.GetByID(ctx, id)
}

// List retrieves services matching the filter (optionally tenant scoped)
func (s *serviceImpl) List(ctx *appContext.ScopedContext, filter *domain.Filter) (*types.PageResult[domain.Services], error) {
	filter = filter.WithDefaults()
	result, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	// build pagination response
	paginatedResult := utils.Paginate(result.Data, result.Total, filter.Offset, filter.Limit)
	return paginatedResult, nil
}

// Update modifies an existing service
func (s *serviceImpl) Update(ctx *appContext.ScopedContext, id uuid.UUID, payload *domain.Update) (*domain.Services, error) {
	return s.repo.Update(ctx, id, payload)
}

// Delete removes a service permanently
func (s *serviceImpl) Delete(ctx *appContext.ScopedContext, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
