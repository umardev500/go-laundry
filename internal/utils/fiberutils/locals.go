package fiberutils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetTenantIDfromCtx extract the tenant_id from Fiber context locals.
//
// Returns a pointer to the UUID if present and valid, otherwise returns nil.
//
// Example usage:
//
//	tenantID := fiberutils.GetTenantIDFromCtx(c)
//	if tenantID != nil {
//	    // tenant-scoped operation
//	}
func GetTenantIDfromCtx(c *fiber.Ctx) *uuid.UUID {
	var tenantIDPtr *uuid.UUID
	if val := c.Locals("tenant_id"); val != nil {
		if id, ok := val.(uuid.UUID); ok && id != uuid.Nil {
			tenantIDPtr = func() *uuid.UUID {
				return &id
			}()
		}
	}
	return tenantIDPtr
}
