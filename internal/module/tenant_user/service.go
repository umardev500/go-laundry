package tenantuser

import (
	"context"

	"github.com/google/uuid"
	domain "github.com/umardev500/go-laundry/internal/domain/tenant_user"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"
)

type serviceImpl struct {
	repo domain.Repository
}

var _ domain.Service = (*serviceImpl)(nil)

func NewService(repo domain.Repository) domain.Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) Create(ctx context.Context, payload *domain.Create) (*domain.TenantUser, error) {
	return s.repo.Create(ctx, payload)
}

func (s *serviceImpl) GetByID(ctx context.Context, id uuid.UUID) (*domain.TenantUser, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *serviceImpl) GetByUserID(ctx context.Context, userID uuid.UUID) (*domain.TenantUser, error) {
	return s.repo.GetByUserID(ctx, userID)
}

func (s *serviceImpl) List(ctx context.Context, f domain.Filter) (*types.PageResult[domain.TenantUser], error) {
	f = f.WithDefaults()
	result, err := s.repo.List(ctx, f)
	if err != nil {
		return nil, err
	}

	return utils.Paginate(result.Data, result.Total, f.Offset, f.Limit), nil
}

func (s *serviceImpl) Update(ctx context.Context, id uuid.UUID, payload *domain.Update) (*domain.TenantUser, error) {
	return s.repo.Update(ctx, id, payload)
}

func (s *serviceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
