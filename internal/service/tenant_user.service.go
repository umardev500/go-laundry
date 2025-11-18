package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/domain"
	"github.com/umardev500/laundry/internal/repository"
)

type TenantUserService struct {
	repo repository.TenantUserRepository
}

func NewTenantUseService(repo repository.TenantUserRepository) *TenantUserService {
	return &TenantUserService{
		repo: repo,
	}
}

func (s *TenantUserService) Find(ctx *core.Context, f *domain.TenantUserFilter) ([]*domain.TenantUser, int, error) {
	tenantID, _ := uuid.Parse("11111111-1111-1111-1111-111111111111")
	f.Criteria.TenantID = &tenantID

	return s.repo.Find(ctx, f)
}
