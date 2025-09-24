package role

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	// AssignRoleToUser assigns a role to a user
	AssignRoleToUser(ctx context.Context, userID, roleID uuid.UUID) error

	// CreateRole creates a tenant
	CreateRole(ctx context.Context, payload *RoleCreate, tenantID *uuid.UUID) (*Role, error)

	// GetRoleByName fetches a role by name for a tenant
	GetRoleByName(ctx context.Context, name string, tenantID *uuid.UUID) (*Role, error)

	// ListRoles fetches all roles for a tenant
	ListRoles(ctx context.Context, tenantID *uuid.UUID) ([]Role, error)
}
