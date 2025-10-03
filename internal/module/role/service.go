package role

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/internal/domain/role"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"
)

type serviceImpl struct {
	repo role.Repository
}

// AssignRoleToUser implements role.Service.
func (s *serviceImpl) AssignRoleToUser(ctx context.Context, tenantID *uuid.UUID, userID uuid.UUID, roleID uuid.UUID) error {
	return s.repo.AssignRoleToUser(ctx, tenantID, userID, roleID)
}

// CreateRole implements role.Service.
func (s *serviceImpl) CreateRole(ctx context.Context, payload *role.RoleCreate, tenantID *uuid.UUID) (*role.Role, error) {
	// Check if role arelady exist for this tenant or globally
	existingRole, err := s.repo.FindByName(ctx, payload.Name, tenantID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if existingRole != nil {
		scope := "global"
		if tenantID != nil {
			scope = "tenant"
		}
		return nil, fmt.Errorf("role '%s' already exists in %s scope", payload.Name, scope)
	}

	// Create the role
	newRole, err := s.repo.Create(ctx, payload, tenantID)
	if err != nil {
		return nil, err
	}

	return newRole, nil
}

// GetRoleByName implements role.Service.
func (s *serviceImpl) GetRoleByName(ctx context.Context, name string, tenantID *uuid.UUID) (*role.Role, error) {
	return s.repo.FindByName(ctx, name, tenantID)
}

// ListRoles implements role.Service.
func (s *serviceImpl) ListRoles(ctx context.Context, f *role.Filter, tenantID *uuid.UUID) (*types.PageResult[role.Role], error) {
	f = f.WithDefaults()

	result, err := s.repo.List(ctx, f, tenantID)
	if err != nil {
		return nil, err
	}

	paginatedResult := utils.Paginate(result.Data, result.Total, f.Offset, f.Limit)
	return paginatedResult, nil
}

func NewService(repo role.Repository) role.Service {
	return &serviceImpl{
		repo: repo,
	}
}
