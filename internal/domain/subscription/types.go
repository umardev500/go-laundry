package subscription

import (
	"time"

	"github.com/google/uuid"
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
	TenantID  *uuid.UUID         `json:"tenant_id"`
	StartDate *time.Time         `json:"start_date"`
	EndDate   *time.Time         `json:"end_date"`
	Status    SubscriptionStatus `json:"status"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type SubscriptionCreate struct {
	PlanID    uuid.UUID           `json:"plan_id"`
	TenantID  uuid.UUID           `json:"tenant_id"`
	StartDate *time.Time          `json:"start_date"`
	EndDate   *time.Time          `json:"end_date"`
	Status    *SubscriptionStatus `json:"status"`
}
