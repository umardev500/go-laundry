package category

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
	domain "github.com/umardev500/go-laundry/internal/domain/category"
)

type serviceImpl struct {
	repo domain.Repository
}

// Ensure serviceImpl implements the domain Service interface
var _ domain.Service = (*serviceImpl)(nil)

// NewService creates a new category service
func NewService(repo domain.Repository) domain.Service {
	return &serviceImpl{repo: repo}
}

// Create a new category
func (s *serviceImpl) Create(ctx *appContext.ScopedContext, payload *domain.Create) (*domain.Category, error) {
	return s.repo.Create(ctx, payload)
}

// GetByID retrieves a category by its ID (optionally tenant scoped)
func (s *serviceImpl) GetByID(ctx *appContext.ScopedContext, id uuid.UUID) (*domain.Category, error) {
	return s.repo.GetByID(ctx, id)
}

// List retrieves categories matching the filter (optionally tenant scoped)
func (s *serviceImpl) List(ctx *appContext.ScopedContext, f domain.Filter) (*types.PageResult[domain.Category], error) {
	f = f.WithDefaults()
	result, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	paginatedResult := utils.Paginate(result.Data, result.Total, f.Offset, f.Limit)
	return paginatedResult, nil
}

// Update modifies an existing category
func (s *serviceImpl) Update(ctx *appContext.ScopedContext, id uuid.UUID, payload *domain.Update) (*domain.Category, error) {
	return s.repo.Update(ctx, id, payload)
}

// Delete removes a category permanently
func (s *serviceImpl) Delete(ctx *appContext.ScopedContext, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
