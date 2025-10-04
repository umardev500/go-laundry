package role

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

type Repository interface {
	Create(ctx *appContext.ScopedContext, role *RoleCreate) (*Role, error)
	List(ctx *appContext.ScopedContext, filter *Filter) (*types.PageData[Role], error)
	FindByName(ctx *appContext.ScopedContext, name string) (*Role, error)
	AssignRoleToUser(ctx *appContext.ScopedContext, userID, roleID uuid.UUID) error
}
