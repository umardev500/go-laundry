package user

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	// Create registres a new user account.
	Create(ctx context.Context, u *UserCreate) (*User, error)

	// CreateProfile creates a new profile for an existing user.
	// Typically used right after user creation.
	CreateProfile(ctx context.Context, userID uuid.UUID, u *ProfileCreate) (*Profile, error)

	// Delete marks a user as inactive (soft delete)
	// Tenant admin can only delete their own users; platform admin can delete any user
	Delete(ctx context.Context, tenantID *uuid.UUID, userID uuid.UUID) error

	// FindByEmail retrieves a user by email
	FindByEmail(ctx context.Context, email string) (*User, error)

	// FindByToken retrieves a user by token
	FindByToken(ctx context.Context, token string) (*User, error)

	// List retrieves users based on the provided filter.
	// The service can apply tenant-scoping and other business rules before calling the repository.
	List(ctx context.Context, filter UserFilter) ([]*User, error)

	// Purge permanently removes a user from the system.
	// Tenant admin can purge only their own users; platform admin can purge any user.
	Purge(ctx context.Context, tenantID *uuid.UUID, userID uuid.UUID) error

	// Update modifies a users's credentials (email and/or password).
	// Tenand admin can only update users within their own tenant.
	Update(ctx context.Context, userID uuid.UUID, payload *UserUpdate, tenantID *uuid.UUID) (*User, error)

	// UpdateProfile udpates the profile information for a given user.
	// Fails if the user does not exist or in soft-deleted.
	UpdateProfile(ctx context.Context, userID uuid.UUID, u *ProfileUpdate) (*Profile, error)
}
