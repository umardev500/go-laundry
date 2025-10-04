package user

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

// UserService defined operation manageing users.
type Service interface {
	// Create registres a new user account.
	Create(ctx *appContext.ScopedContext, u *UserCreate) (*User, error)

	// CreateProfile creates a new profile for an existing user.
	// Typically used right after user creation.
	CreateProfile(ctx *appContext.ScopedContext, userID uuid.UUID, u *ProfileCreate) (*Profile, error)

	// Delete marks a user as inactive (soft delete)
	// Tenant admin can only delete their own users; platform admin can delete any user
	Delete(ctx *appContext.ScopedContext, id uuid.UUID) error

	// FindByEmail retrieves a user by email
	FindByEmail(ctx *appContext.ScopedContext, email string) (*User, error)

	// FindByToken retrieves a user by token
	FindByToken(ctx *appContext.ScopedContext, token string) (*User, error)

	// List retrieves users based on the provided filter.
	// The service can apply tenant-scoping and other business rules before calling the repository.
	List(ctx *appContext.ScopedContext, filter *Filter) (*types.PageResult[User], error)

	// Purge permanently removes a user from the system.
	// Tenant admin can purge only their own users; platform admin can purge any user.
	Purge(ctx *appContext.ScopedContext, id uuid.UUID) error

	// Update modifies a users's credentials (email and/or password).
	// Tenand admin can only update users within their own tenant.
	Update(ctx *appContext.ScopedContext, id uuid.UUID, payload *UserUpdate) (*User, error)

	// UpdateProfile udpates the profile information for a given user.
	// Fails if the user does not exist or in soft-deleted.
	UpdateProfile(ctx *appContext.ScopedContext, userID uuid.UUID, u *ProfileUpdate) (*Profile, error)
}
