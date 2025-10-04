package role

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/internal/domain/role"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

type serviceImpl struct {
	repo role.Repository
}

// AssignRoleToUser implements role.Service.
func (s *serviceImpl) AssignRoleToUser(ctx *appContext.ScopedContext, userID uuid.UUID, roleID uuid.UUID) error {
	return s.repo.AssignRoleToUser(ctx, userID, roleID)
}

// CreateRole implements role.Service.
func (s *serviceImpl) CreateRole(ctx *appContext.ScopedContext, payload *role.RoleCreate) (*role.Role, error) {
	scoped := ctx.Scoped
	tenantID := scoped.TenantID

	// Check if role arelady exist for this tenant or globally
	existingRole, err := s.repo.FindByName(ctx, payload.Name)
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
	newRole, err := s.repo.Create(ctx, payload)
	if err != nil {
		return nil, err
	}

	return newRole, nil
}

// GetRoleByName implements role.Service.
func (s *serviceImpl) GetRoleByName(ctx *appContext.ScopedContext, name string) (*role.Role, error) {
	return s.repo.FindByName(ctx, name)
}

// ListRoles implements role.Service.
func (s *serviceImpl) ListRoles(ctx *appContext.ScopedContext, f *role.Filter) (*types.PageResult[role.Role], error) {
	f = f.WithDefaults()

	result, err := s.repo.List(ctx, f)
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
