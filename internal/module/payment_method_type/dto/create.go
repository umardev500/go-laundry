package dto

import paymentmethodtype "github.com/umardev500/go-laundry/internal/domain/payment_method_type"

type Create struct {
	Name        string                   `json:"name" validate:"required"`
	DisplayName string                   `json:"display_name" validate:"required"`
	Status      paymentmethodtype.Status `json:"status" validate:"required"`
}

func (c Create) ToPaymentMethodTypeCreate() *paymentmethodtype.Create {
	return &paymentmethodtype.Create{
		Name:        c.Name,
		DisplayName: c.DisplayName,
		Status:      &c.Status,
	}
}
