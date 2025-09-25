package fiberutils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/pkg/response"
)

// GetUUIDParamOrAPIError parses a UUID from the URL path parameter.
// Returns the UUID, or writes an APIResponse with the proper HTTP status directly.
func GetUUIDParamOrAPIError(c *fiber.Ctx, key string) (uuid.UUID, bool) {
	val := c.Params(key)
	if val == "" {
		c.Status(fiber.StatusBadRequest).JSON(&response.APIResponse[any]{
			Success: false,
			Error:   "missing " + key,
		})
		return uuid.Nil, false
	}

	id, err := uuid.Parse(val)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(&response.APIResponse[any]{
			Success: false,
			Error:   "invalid " + key,
		})
		return uuid.Nil, false
	}

	return id, true
}
