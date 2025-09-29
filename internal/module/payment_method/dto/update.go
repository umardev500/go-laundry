package dto

import (
	"github.com/google/uuid"
	paymentmethod "github.com/umardev500/go-laundry/internal/domain/payment_method"
)

type Update struct {
	Metadata *map[string]any `json:"metadata"`
}

func (c Update) ToPaymentMethodCreate(tenantID uuid.UUID) *paymentmethod.Update {
	return &paymentmethod.Update{
		Metadata: c.Metadata,
	}
}
