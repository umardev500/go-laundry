package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *UserCreate) (*User, error)
	CreateUserProfile(ctx context.Context, profile *ProfileCreate) (*Profile, error)
	UpdateUserProfile(ctx context.Context, userID uuid.UUID, u *ProfileUpdate) (*Profile, error)
}
