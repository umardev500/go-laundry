package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/ent"
)

// Params
type CreateMerchantInput struct {
	Name    string
	Email   string
	Phone   string
	Address string
}

// DTO
type CreateMerchantRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type MerchantRepository interface {
	Create(ctx context.Context, payload *CreateMerchantInput) (*ent.Merchant, error)
	ExistsByUserID(ctx context.Context, ownerID uuid.UUID) (bool, error)
}

type MerchantUsecase interface {
	Register(ctx context.Context, merchant *CreateMerchantRequest) error
}
