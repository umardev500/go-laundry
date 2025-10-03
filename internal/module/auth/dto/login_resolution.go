package dto

import (
	platformuser "github.com/umardev500/go-laundry/internal/domain/platform_user"
	tenantuser "github.com/umardev500/go-laundry/internal/domain/tenant_user"
)

type LoginResolution struct {
	PlatformUser *platformuser.PlatformUser `json:"platform_user,omitempty"`
	TenantUsers  []*tenantuser.TenantUser   `json:"tenant_users,omitempty"`
}
