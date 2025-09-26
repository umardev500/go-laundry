package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/subscription"
)

type CreateSubscriptionRequest struct {
	PlanID uuid.UUID `json:"plan_id" validate:"required"`
}

func (r *CreateSubscriptionRequest) ToSubscriptionCreate(tenantID uuid.UUID) *subscription.SubscriptionCreate {
	return &subscription.SubscriptionCreate{
		PlanID:    r.PlanID,
		TenantID:  tenantID,
		StartDate: nil,
		EndDate:   nil,
		Status:    nil,
	}
}
