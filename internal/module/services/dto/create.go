package dto

import (
	"github.com/google/uuid"
	domain "github.com/umardev500/go-laundry/internal/domain/services"
)

type Create struct {
	Name        string     `json:"name" validate:"required"`
	Description *string    `json:"description"`
	BasePrice   float64    `json:"base_price" validate:"required"`
	CategoryID  *uuid.UUID `json:"category_id"`
	UnitID      *uuid.UUID `json:"unit_id"`
}

// ToDomain converts the DTO to the domain Create struct.
func (c Create) ToDomain(tenantID *uuid.UUID) *domain.Create {
	return &domain.Create{
		TenantID:    tenantID,
		CategoryID:  c.CategoryID,
		Name:        c.Name,
		Description: c.Description,
		BasePrice:   c.BasePrice,
		UnitID:      c.UnitID,
	}
}
