package dto

import (
	"github.com/google/uuid"
	domain "github.com/umardev500/go-laundry/internal/domain/category"
)

type Create struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

// ToDomain converts the DTO to the domain Create struct, setting tenant ID.
func (c Create) ToDomain(tenantID *uuid.UUID) *domain.Create {
	return &domain.Create{
		TenantID:    tenantID,
		Name:        c.Name,
		Description: c.Description,
	}
}
