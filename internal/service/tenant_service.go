package service

import (
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/domain"
	"github.com/umardev500/laundry/internal/repository"
)

type TenantService struct {
	repo repository.TenantRepository
}

func NewTenantService(repo repository.TenantRepository) *TenantService {
	return &TenantService{
		repo: repo,
	}
}

func (s *TenantService) Find(ctx *core.Context, f *domain.TenantFilter) ([]*domain.Tenant, int, error) {
	return s.repo.Find(ctx, f)
}
