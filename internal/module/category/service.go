package category

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/umardev500/go-laundry/internal/domain/category"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"
	"github.com/umardev500/go-laundry/pkg/response"
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
func (s *serviceImpl) Create(ctx context.Context, payload *domain.Create) (*domain.Category, error) {
	return s.repo.Create(ctx, payload)
}

// GetByID retrieves a category by its ID (optionally tenant scoped)
func (s *serviceImpl) GetByID(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID) (*domain.Category, error) {
	return s.repo.GetByID(ctx, tenantID, id)
}

// List retrieves categories matching the filter (optionally tenant scoped)
func (s *serviceImpl) List(ctx context.Context, tenantID *uuid.UUID, f domain.Filter) (*types.PageResult[domain.Category], error) {
	f = f.WithDefaults()
	result, err := s.repo.List(ctx, tenantID, f)
	if err != nil {
		return nil, err
	}

	page := f.Offset + 1
	totalPages := utils.CalculateTotalPages(result.Total, f.Limit)

	return &types.PageResult[domain.Category]{
		Data: result.Data,
		Pagination: &response.Pagination{
			Page:       page,
			PageSize:   f.Limit,
			TotalItems: result.Total,
			TotalPages: totalPages,
			HasNext:    f.Offset+1 < totalPages,
			HasPrev:    f.Offset > 1,
		},
	}, nil
}

// Update modifies an existing category
func (s *serviceImpl) Update(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID, payload *domain.Update) (*domain.Category, error) {
	return s.repo.Update(ctx, tenantID, id, payload)
}

// Delete removes a category permanently
func (s *serviceImpl) Delete(ctx context.Context, tenantID *uuid.UUID, id uuid.UUID) error {
	return s.repo.Delete(ctx, tenantID, id)
}
