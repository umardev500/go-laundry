package role

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, role *RoleCreate, tenantID *uuid.UUID) (*Role, error)
	List(ctx context.Context, tenantID *uuid.UUID) ([]*Role, error)
	FindByName(ctx context.Context, name string, tenantID *uuid.UUID) (*Role, error)
	AssignRoleToUser(ctx context.Context, tenantID *uuid.UUID, userID, roleID uuid.UUID) error
}
