package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/commands"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/domain"
	"github.com/umardev500/laundry/internal/errors"
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

// Create creates a new tenant.
func (s *TenantService) Create(ctx *core.Context, cmd *commands.CreateTenantCmd) (*domain.Tenant, error) {
	newTenant, err := domain.NewTenant(cmd.Name)
	if err != nil {
		return nil, err
	}

	return s.repo.Create(ctx, newTenant)
}

// Delete deletes a tenant for the given id.
func (s *TenantService) Delete(ctx *core.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

// Find finds tenants by filter paramters.
func (s *TenantService) Find(ctx *core.Context, f *domain.TenantFilter) ([]*domain.Tenant, int, error) {
	return s.repo.Find(ctx, f)
}

// FindByID finds tenant for the given id.
func (s *TenantService) FindByID(ctx *core.Context, id uuid.UUID) (*domain.Tenant, error) {
	t, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.NewTenantNotFound(id)
		}

		return nil, err
	}

	return t, nil
}

// Update updates tenant for the given id.
func (s *TenantService) Update(ctx *core.Context, id uuid.UUID, cmd *commands.UpdateTenantCmd) (*domain.Tenant, error) {
	// Ensure tenant to update id exists.
	tExists, err := s.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	tExists.Update(cmd.Name)

	return s.repo.Update(ctx, tExists)
}
