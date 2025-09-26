package redisutils

import (
	"fmt"

	"github.com/google/uuid"
)

// ActivePlan returns the key for the active plan for a tenant.
func ActivePlan(tenantID uuid.UUID) string {
	return fmt.Sprintf("tenant:%s:plan", tenantID.String())
}

// PlanPermissions returns the Redis key for a plan’s permissions set
func PlanPermissions(planID uuid.UUID) string {
	return "plan:" + planID.String() + ":permissions"
}
