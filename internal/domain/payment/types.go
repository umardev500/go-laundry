package payment

import (
	"time"

	"github.com/google/uuid"
	paymentmethod "github.com/umardev500/go-laundry/internal/domain/payment_method"
)

// ReferenceType represents the type of payment reference
type ReferenceType string

const (
	Subscription ReferenceType = "subscription"
)

// Currency represents the currency of the payment
type Currency string

const (
	IDR Currency = "IDR"
)

// Status represents the status of the payment
type Status string

const (
	Pending   Status = "pending"
	Completed Status = "completed"
	Failed    Status = "failed"
	Cancelled Status = "cancelled"
)

type Payment struct {
	ID            uuid.UUID                    `json:"id"`
	UserID        uuid.UUID                    `json:"user_id"`
	TenantID      *uuid.UUID                   `json:"tenant_id"`
	ReferenceID   uuid.UUID                    `json:"reference_id"`
	ReferenceType ReferenceType                `json:"reference_type"`
	Amount        float64                      `json:"amount"`
	Currency      Currency                     `json:"currency"`
	Status        Status                       `json:"status"`
	Method        *paymentmethod.PaymentMethod `json:"method"`
	ProofURL      *string                      `json:"proof_url"`
	PaidAt        *time.Time                   `json:"paid_at"`
	CreatedAt     time.Time                    `json:"created_at"`
	UpdatedAt     time.Time                    `json:"updated_at"`
}

type PaymentCreate struct {
	UserID          uuid.UUID
	TenantID        *uuid.UUID
	ReferenceID     uuid.UUID
	ReferenceType   ReferenceType
	PaymentMethodID uuid.UUID
	Amount          float64
	Currency        Currency
	Status          Status
	PaidAt          *time.Time
}

type PaymentUpdate struct {
	ProofURL *string
	AdminID  *uuid.UUID
	Amount   *float64
	Status   *Status
	PaidAt   *time.Time
}

type OrderBy string

const (
	OrderByCreatedAtAsc  OrderBy = "created_at_asc"
	OrderByCreatedAtDesc OrderBy = "created_at_desc"
)

type Filter struct {
	Query             string         `query:"query"`
	Limit             int            `query:"limit"`
	Offset            int            `query:"offset"`
	OrderBy           OrderBy        `query:"order_by"`
	Status            *Status        `query:"status"`
	Type              *ReferenceType `query:"type"`
	HasProof          bool           `query:"has_proof"`
	IncludeMethod     bool           `query:"include_method"`
	IncludeMethodType bool           `query:"include_method_type"`
	IncludeTenant     bool           `query:"include_tenant"`
}

func (f Filter) WithDefaults() *Filter {
	if f.Limit == 0 {
		f.Limit = 10
	}
	if f.Offset == 0 {
		f.Offset = 0
	}
	if f.OrderBy == "" {
		f.OrderBy = "created_at desc"
	}

	return &f
}
