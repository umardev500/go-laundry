package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	// Create persists a new user record
	// The tenant ID should be injected into UserCreate at service layer (not from client request)
	// If tenantID is nil, the user is a platform user
	// Othwerwise, the user is a tenant user
	Create(ctx context.Context, user *UserCreate) (*User, error)

	// CreateProfile creates a profile associated with the given user.
	CreateProfile(ctx context.Context, userID uuid.UUID, profile *ProfileCreate) (*Profile, error)

	// Delete performs a soft delete by setting deleted_at on the user.
	// If tenantID is not nil, the deletion is scoped to that tenant (tenant users can
	// only delete users within their own tenant).
	// If tenantID is nil (platform user), deleletion is unrestricted.
	Delete(ctx context.Context, tenantID *uuid.UUID, userID uuid.UUID) error

	// FindByEmail retrieve a non-deleted user by email
	// Soft-deleted users are excluded
	FindByEmail(ctx context.Context, email string) (*User, error)

	// list retrieves users based on the filter criteria
	List(ctx context.Context, filter UserFilter) ([]*User, error)

	// PurgeUser performs a hard delete, physically removing the user record.
	//
	// If tenantID is not nil, the purge is restricted to users that belong to that tenant.
	// This allow tenant administators to permanently remove their own members.
	//
	// If tenantID is nil (platform user), the purge is unrestricted.
	// and can be applied across all tenants.
	PurgeUser(ctx context.Context, tenantID *uuid.UUID, userID uuid.UUID) error

	// UpdateProfile updated an existing profile for the give user.
	UpdateProfile(ctx context.Context, userID uuid.UUID, u *ProfileUpdate) (*Profile, error)
}
