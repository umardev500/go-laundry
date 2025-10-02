package role

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"
)

type Repository interface {
	Create(ctx context.Context, role *RoleCreate, tenantID *uuid.UUID) (*Role, error)
	List(ctx context.Context, filter *Filter, tenantID *uuid.UUID) (*types.PageData[Role], error)
	FindByName(ctx context.Context, name string, tenantID *uuid.UUID) (*Role, error)
	AssignRoleToUser(ctx context.Context, tenantID *uuid.UUID, userID, roleID uuid.UUID) error
}
