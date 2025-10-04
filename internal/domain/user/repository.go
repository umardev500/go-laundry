package user

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

type Repository interface {
	// Create persists a new user record
	// The tenant ID should be injected into UserCreate at service layer (not from client request)
	// If tenantID is nil, the user is a platform user
	// Othwerwise, the user is a tenant user
	Create(ctx *appContext.ScopedContext, user *UserCreate) (*User, error)

	// CreateProfile creates a profile associated with the given user.
	CreateProfile(ctx *appContext.ScopedContext, id uuid.UUID, profile *ProfileCreate) (*Profile, error)

	// Delete performs a soft delete by setting deleted_at on the user.
	// If tenantID is not nil, the deletion is scoped to that tenant (tenant users can
	// only delete users within their own tenant).
	// If tenantID is nil (platform user), deleletion is unrestricted.
	Delete(ctx *appContext.ScopedContext, id uuid.UUID) error

	// FindByEmail retrieve a non-deleted user by email
	// Soft-deleted users are excluded
	FindByEmail(ctx *appContext.ScopedContext, email string) (*User, error)

	// FindByToken retrieve a non-deleted user by token
	// Soft-deleted users are excluded
	FindByToken(ctx *appContext.ScopedContext, token string) (*User, error)

	// list retrieves users based on the filter criteria
	List(ctx *appContext.ScopedContext, filter *Filter) (*types.PageData[User], error)

	// PurgeUser performs a hard delete, physically removing the user record.
	//
	// If tenantID is not nil, the purge is restricted to users that belong to that tenant.
	// This allow tenant administators to permanently remove their own members.
	//
	// If tenantID is nil (platform user), the purge is unrestricted.
	// and can be applied across all tenants.
	PurgeUser(ctx *appContext.ScopedContext, id uuid.UUID) error

	// Update modifies a user's credentials (email and/or password).
	//
	// - payload: contains optional fields to update; only non-nil fields will be changed.
	// - userID: identifies the user to update.
	// - tenantID: if not nil, the update is scoped to the tenant. Tenant admin can only update users
	//   within their own tenant. Platform admins can pass nil to update any user.
	//
	//
	// Behavior:
	// - Only non-nil fields in payload are updated (partial update).
	// - If the user is soft-deleted, the update will fail and return an error.
	// - Returns the updated user on success
	Update(ctx *appContext.ScopedContext, payload *UserUpdate, id uuid.UUID) (*User, error)

	// UpdateProfile updated an existing profile for the give user.
	UpdateProfile(ctx *appContext.ScopedContext, id uuid.UUID, u *ProfileUpdate) (*Profile, error)
}
