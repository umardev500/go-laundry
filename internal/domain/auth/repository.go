package auth

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	// GetActivePlanID retrieves the active plan for a tenant.
	GetActivePlanID(ctx context.Context, tenantID uuid.UUID) (uuid.UUID, error)
}
