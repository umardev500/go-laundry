package permission

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	AssignPermissionsToRole(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error
	GetPermissionsByNames(ctx context.Context, names []string) ([]Permission, error)
}
