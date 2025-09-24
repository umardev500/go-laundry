package user

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	UpdateUserProfile(ctx context.Context, userID uuid.UUID, u *ProfileUpdate) (*Profile, error)
	CreateUser(ctx context.Context, u *UserCreate) (*User, error)
}
