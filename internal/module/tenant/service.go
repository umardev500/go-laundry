package tenant

import (
	"context"

	"github.com/umardev500/go-laundry/internal/domain/tenant"
)

type serviceImpl struct {
	repo tenant.Repository
}

// CreateTenant implements tenant.Service.
func (s *serviceImpl) CreateTenant(ctx context.Context, tenant *tenant.TenantCreate) (*tenant.Tenant, error) {
	return s.repo.CreateTenant(ctx, tenant)
}

func NewService(repo tenant.Repository) tenant.Service {
	return &serviceImpl{
		repo: repo,
	}
}
