package role

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	// SeedDefaultRoles inserts default roles for a tenant
	SeedDefaultRoles(ctx context.Context, tenantID uuid.UUID) error

	// CreateRole creates a tenant
	CreateRole(ctx context.Context, payload *RoleCreate, tenantID *uuid.UUID) (*Role, error)

	// GetRoleByName fetches a role by name for a tenant
	GetRoleByName(ctx context.Context, name string, tenantID *uuid.UUID) (*Role, error)

	// ListRoles fetches all roles for a tenant
	ListRoles(ctx context.Context, tenantID *uuid.UUID) ([]Role, error)
}
