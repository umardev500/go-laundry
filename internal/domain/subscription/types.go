package subscription

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/payment"
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
	Payment   *payment.Payment   `json:"payment"`
	StartDate *time.Time         `json:"start_date"`
	EndDate   *time.Time         `json:"end_date"`
	Status    SubscriptionStatus `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type OrderBy string

const (
	OrderByCreatedAtAsc  OrderBy = "created_at_asc"
	OrderByCreatedAtDesc OrderBy = "created_at_desc"
)

type Filter struct {
	Query          string  `query:"query"`
	Limit          int     `query:"limit"`
	Offset         int     `query:"offset"`
	OrderBy        OrderBy `query:"order_by"`
	IncludePlan    bool    `query:"include_plan"`
	IncludeTenant  bool    `query:"include_tenant"`
	IncludePayment bool    `query:"include_payment"`
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

type SubscriptionCreate struct {
	PlanID          uuid.UUID
	TenantID        uuid.UUID
	PaymentMethodID uuid.UUID
	StartDate       *time.Time
	EndDate         *time.Time
	Status          *SubscriptionStatus
}

type SubscriptionUpdate struct {
	StartDate *time.Time
	EndDate   *time.Time
	Status    *SubscriptionStatus
}
