package dto

import (
	"github.com/google/uuid"
	domain "github.com/umardev500/go-laundry/internal/domain/services"
)

type Update struct {
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	BasePrice   *float64   `json:"base_price"`
	UnitID      *uuid.UUID `json:"unit_id"`
	// NOTE: CategoryID is intentionally excluded from update (as per your rules).
}

// ToDomain converts the DTO to the domain Update struct.
func (u Update) ToDomain() *domain.Update {
	return &domain.Update{
		Name:        u.Name,
		Description: u.Description,
		BasePrice:   u.BasePrice,
		UnitID:      u.UnitID,
	}
}
