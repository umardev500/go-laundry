package fiberutils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/pkg/response"
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

// GetScopedFromCtx builds a Scoped object from Fiber context.
func GetScopedFromCtx(c *fiber.Ctx) *types.Scoped {
	var tenantID *uuid.UUID

	if v := c.Locals("tenant_id"); v != nil {
		if tid, ok := v.(uuid.UUID); ok {
			tenantID = &tid
		}
	}

	var scope types.Scope
	if v := c.Locals("user_scope"); v != nil {
		scope, _ = v.(types.Scope)
	}

	s := &types.Scoped{
		TenantID: tenantID,
		Scope:    scope,
	}

	// ✅ run validation immediately
	if err := s.Validate(); err != nil {
		_ = c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
		return nil
	}

	return s
}
