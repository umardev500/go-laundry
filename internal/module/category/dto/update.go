package dto

import (
	domain "github.com/umardev500/go-laundry/internal/domain/category"
)

type Update struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

// ToDomain converts the DTO to the domain Update struct.
func (u Update) ToDomain() *domain.Update {
	return &domain.Update{
		Name:        u.Name,
		Description: u.Description,
	}
}
