package domain

import (
	"github.com/umardev500/go-laundry/internal/ent/plan"
)

type CreatePlanRequest struct {
	Name         string            `json:"name" validate:"required"`
	Price        float64           `json:"price" validate:"required"`
	Description  string            `json:"description"`
	MaxBranch    int               `json:"max_branch"`
	MaxOrder     int               `json:"max_order"`
	MaxUsers     int               `json:"max_users"`
	MaxCustomers int               `json:"max_customers"`
	BillingCycle plan.BillingCycle `json:"billing_cycle"`
	DurationDays int               `json:"duration_days" validate:"min=1,max=365"`
	Enabled      bool              `json:"enabled"`
}
