package subscription

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID        uuid.UUID  `json:"id"`
	PlanID    *uuid.UUID `json:"plan_id"`
	TenantID  *uuid.UUID `json:"tenant_id"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
