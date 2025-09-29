package dto

import paymentmethodtype "github.com/umardev500/go-laundry/internal/domain/payment_method_type"

type Update struct {
	Name        *string                   `json:"name" validate:"required"`
	DisplayName *string                   `json:"display_name" validate:"required"`
	Status      *paymentmethodtype.Status `json:"status" validate:"required"`
}

func (u Update) ToPaymentMethodTypeUpdate() *paymentmethodtype.Update {
	return &paymentmethodtype.Update{
		Name:        u.Name,
		DisplayName: u.DisplayName,
		Status:      u.Status,
	}
}
