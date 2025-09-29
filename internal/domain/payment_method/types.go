package paymentmethod

import (
	"time"

	"github.com/google/uuid"
	paymentmethodtype "github.com/umardev500/go-laundry/internal/domain/payment_method_type"
)

type QRISMetadata struct {
	ImageURL     string `json:"image_url"`
	MerchantName string `json:"merchant_name"`
}

type BankStransferMetadata struct {
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
	AccountHolder string `json:"account_holder"`
}

type CashMetadata struct {
	Note *string `json:"note"`
}

type PaymentMethod struct {
	ID        uuid.UUID                            `json:"id"`
	TenantID  *uuid.UUID                           `json:"tenant_id"`
	TypeID    uuid.UUID                            `json:"type_id"`
	Type      *paymentmethodtype.PaymentMethodType `json:"type"`
	Metadata  map[string]any                       `json:"metadata"`
	CreatedAt time.Time                            `json:"created_at"`
	UpdatedAt time.Time                            `json:"updated_at"`
}

type Create struct {
	TenantID uuid.UUID
	TypeID   uuid.UUID
	Metadata map[string]any
}

type Update struct {
	Metadata *map[string]any
}

// for now is empty
type Filter struct {
	IncludeType bool `query:"include_type"`
}

func (f *Filter) WithDefaults() *Filter {
	return f
}
