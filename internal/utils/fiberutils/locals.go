package fiberutils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/pkg/response"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
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
	var tenantID *uuid.UUID
	if val := c.Locals("tenant_id"); val != nil {
		if id, ok := val.(uuid.UUID); ok && id != uuid.Nil {
			tenantID = func() *uuid.UUID {
				return &id
			}()
		}
	}
	return tenantID
}

// GetScopedFromLocals builds a Scoped object from Fiber context.
func GetScopedFromLocals(c *fiber.Ctx) *appContext.Scoped {
	var scope appContext.Scoped
	if v := c.Locals("user_scope"); v != nil {
		scope, _ = v.(appContext.Scoped)
	}

	// ✅ run validation immediately
	if err := scope.Validate(); err != nil {
		_ = c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
		return nil
	}

	return &scope
}
