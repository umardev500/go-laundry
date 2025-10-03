package tenantuser

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/domain/tenant"
)

type Status string

const (
	StatusActive    Status = "active"
	StatusSuspended Status = "suspended"
	StatusDeleted   Status = "deleted"
)

// TenantUser domain model
type TenantUser struct {
	ID        uuid.UUID      `json:"id"`
	TenantID  uuid.UUID      `json:"tenant_id"`
	Tenant    *tenant.Tenant `json:"tenant"`
	UserID    uuid.UUID      `json:"user_id"`
	Status    Status         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt *time.Time     `json:"deleted_at,omitempty"`
}
