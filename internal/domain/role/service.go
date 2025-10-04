package role

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

type Service interface {
	// AssignRoleToUser assigns a role to a user
	AssignRoleToUser(ctx *appContext.ScopedContext, userID, roleID uuid.UUID) error

	// CreateRole creates a tenant
	CreateRole(ctx *appContext.ScopedContext, payload *RoleCreate) (*Role, error)

	// GetRoleByName fetches a role by name for a tenant
	GetRoleByName(ctx *appContext.ScopedContext, name string) (*Role, error)

	// ListRoles fetches all roles for a tenant
	ListRoles(ctx *appContext.ScopedContext, filter *Filter) (*types.PageResult[Role], error)
}
