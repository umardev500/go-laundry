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
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type MerchantRepository interface {
	Create(ctx context.Context, payload *CreateMerchantInput) (*ent.Merchant, error)
	ExistsByUserID(ctx context.Context, ownerID uuid.UUID) (bool, error)
}

type MerchantService interface {
}
