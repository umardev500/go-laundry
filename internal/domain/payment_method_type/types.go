package paymentmethodtype

import (
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
)

type PaymentMethodType struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Create struct {
	Name        string
	DisplayName string
	Status      *Status
}

type Update struct {
	Name        *string
	DisplayName *string
	Status      *Status
}

type Filter struct {
	Status *Status
}

func (f *Filter) WithDefaults() *Filter {
	return f
}
