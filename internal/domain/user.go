package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/ent"
)

// Params

type CreateUserInput struct {
	Name       string
	Email      string
	Password   string
	MerchantID uuid.UUID
}

// DTO

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Interfaces

type UserRepository interface {
	Create(ctx context.Context, payload *CreateUserInput) (*ent.User, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
}

type UserService interface {
}
