package dto

import (
	"github.com/google/uuid"
	paymentmethod "github.com/umardev500/go-laundry/internal/domain/payment_method"
)

type Create struct {
	TypeID   uuid.UUID      `json:"type_id" validate:"required"`
	Metadata map[string]any `json:"metadata" validate:"required"`
}

func (c Create) ToPaymentMethodCreate(tenantID uuid.UUID) *paymentmethod.Create {
	return &paymentmethod.Create{
		TenantID: tenantID,
		TypeID:   c.TypeID,
		Metadata: c.Metadata,
	}
}
