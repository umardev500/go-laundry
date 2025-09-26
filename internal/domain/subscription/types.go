package subscription

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/plan"
	"github.com/umardev500/go-laundry/internal/domain/tenant"
)

type SubscriptionStatus string

const (
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusInactive  SubscriptionStatus = "inactive"
	SubscriptionStatusPending   SubscriptionStatus = "pending"
	SubscriptionStatusCancelled SubscriptionStatus = "cancelled"
	SubscriptionStatusSuspended SubscriptionStatus = "suspended"
)

type Subscription struct {
	ID        uuid.UUID          `json:"id"`
	PlanID    *uuid.UUID         `json:"plan_id"`
	Plan      *plan.Plan         `json:"plan"`
	TenantID  *uuid.UUID         `json:"tenant_id"`
	Tenant    *tenant.Tenant     `json:"tenant"`
	StartDate *time.Time         `json:"start_date"`
	EndDate   *time.Time         `json:"end_date"`
	Status    SubscriptionStatus `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type SubscriptionFilter struct {
	IncludePlan   bool
	IncludeTenant bool
}

func (f SubscriptionFilter) WithDefaults() SubscriptionFilter {
	return f
}

type SubscriptionCreate struct {
	PlanID    uuid.UUID           `json:"plan_id"`
	TenantID  uuid.UUID           `json:"tenant_id"`
	StartDate *time.Time          `json:"start_date"`
	EndDate   *time.Time          `json:"end_date"`
	Status    *SubscriptionStatus `json:"status"`
}

type SubscriptionUpdate struct {
	StartDate *time.Time          `json:"start_date"`
	EndDate   *time.Time          `json:"end_date"`
	Status    *SubscriptionStatus `json:"status"`
}
