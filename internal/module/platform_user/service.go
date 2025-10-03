package platformuser

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/umardev500/go-laundry/internal/domain/platform_user"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"
)

type serviceImpl struct {
	repo domain.Repository
}

// Ensure serviceImpl implements the domain.Service interface
var _ domain.Service = (*serviceImpl)(nil)

// NewService creates a new PlatformUser service
func NewService(repo domain.Repository) domain.Service {
	return &serviceImpl{repo: repo}
}

// Create a new platform user
func (s *serviceImpl) Create(ctx context.Context, payload *domain.Create) (*domain.PlatformUser, error) {
	return s.repo.Create(ctx, payload)
}

// GetByID retrieves a platform user by ID
func (s *serviceImpl) GetByID(ctx context.Context, id uuid.UUID) (*domain.PlatformUser, error) {
	return s.repo.GetByID(ctx, id)
}

// GetByUserID retrieves a platform user by UserID
func (s *serviceImpl) GetByUserID(ctx context.Context, id uuid.UUID) (*domain.PlatformUser, error) {
	return s.repo.GetByUserID(ctx, id)
}

// List retrieves platform users with pagination
func (s *serviceImpl) List(ctx context.Context, f domain.Filter) (*types.PageResult[domain.PlatformUser], error) {
	f = f.WithDefaults()
	result, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return utils.Paginate(result.Data, result.Total, f.Offset, f.Limit), nil
}

// Update modifies an existing platform user
func (s *serviceImpl) Update(ctx context.Context, id uuid.UUID, payload *domain.Update) (*domain.PlatformUser, error) {
	return s.repo.Update(ctx, id, payload)
}

// Delete removes a platform user permanently
func (s *serviceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
