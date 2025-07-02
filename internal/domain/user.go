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

type GetUsersParams struct {
	Page       int `query:"page"`
	Limit      int `query:"limit"`
	Offset     int `query:"-"`
	MerchantID uuid.UUID
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Interfaces

type UserRepository interface {
	Create(ctx context.Context, payload *CreateUserInput) (*ent.User, error)
	GetAll(ctx context.Context, params *GetUsersParams) ([]*ent.User, int, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
}

type UserService interface {
	Create(ctx context.Context, payload *CreateUserRequest) (*ent.User, error)
	GetAll(ctx context.Context, params *GetUsersParams) ([]*ent.User, int, error)
}
