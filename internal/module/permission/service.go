package permission

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/permission"
)

type serviceImpl struct {
	repo permission.Repository
}

// GetPermissionsByNames implements permission.Service.
func (s *serviceImpl) GetPermissionsByNames(ctx context.Context, names []string) ([]permission.Permission, error) {
	return s.repo.GetPermissionsByNames(ctx, names)
}

// AssignPermissionsToRole implements permission.Service.
func (s *serviceImpl) AssignPermissionsToRole(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error {
	return s.repo.AssignPermissionsToRole(ctx, roleID, permissionIDs)
}

func NewService(repo permission.Repository) permission.Service {
	return &serviceImpl{
		repo: repo,
	}
}
