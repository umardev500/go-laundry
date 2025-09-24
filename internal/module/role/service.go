package role

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/internal/domain/role"
)

type serviceImpl struct {
	repo role.Repository
}

// AssignRoleToUser implements role.Service.
func (s *serviceImpl) AssignRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error {
	return s.repo.AssignRoleToUser(ctx, userID, roleID)
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
func (s *serviceImpl) ListRoles(ctx context.Context, tenantID *uuid.UUID) ([]*role.Role, error) {
	return s.repo.List(ctx, tenantID)
}

func NewService(repo role.Repository) role.Service {
	return &serviceImpl{
		repo: repo,
	}
}
