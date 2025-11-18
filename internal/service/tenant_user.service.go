package service

import (
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
	return s.repo.Find(ctx, f)
}
