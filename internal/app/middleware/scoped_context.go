package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/utils/fiberutils"
	"github.com/umardev500/go-laundry/pkg/response"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

func ScopedContextMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Build from locals
		scoped := fiberutils.GetScopedFromLocals(c)
		if scoped == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
				Success: false,
				Error:   "Unauthorized",
			})

		}

		scopedCtx := &appContext.ScopedContext{
			Context: c.UserContext(),
			Scoped:  scoped,
		}

		c.Locals(appContext.ScopedContextKey, scopedCtx)

		return c.Next()
	}
}
