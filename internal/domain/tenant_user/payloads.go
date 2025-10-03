package tenantuser

import "github.com/google/uuid"

// Create payload
type Create struct {
	TenantID uuid.UUID `json:"tenant_id"`
	UserID   uuid.UUID `json:"user_id"`
	Status   *Status   `json:"status,omitempty"`
}

// Update payload
type Update struct {
	Status *Status `json:"status,omitempty"`
}
