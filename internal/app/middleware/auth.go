package middleware

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/pkg/httputil"
	"github.com/umardev500/go-laundry/pkg/response"
)

func CheckAuth(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		tokenStr, err := httputil.ExtractBearerToken(authHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Parse and validate JWT
		token, err := jwt.Parse([]byte(tokenStr), jwt.WithKey(jwa.HS256(), []byte(cfg.JWT.Secret)))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Attach user id to locals
		userIDStr, ok := token.Subject()
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Convert to uuid
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// Fetch tenant id
		var tenantIDStr string
		err = token.Get("tenant_id", &tenantIDStr)
		if err != nil {
			// Skip if claim does not exist
			// Only log or ignore
			log.Info().Msg("No tenant_id found in token")
		} else {
			tenantID, err := uuid.Parse(tenantIDStr)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			c.Locals("tenant_id", tenantID)
		}

		// Fetch plan id if exist
		var planIDStr string
		err = token.Get("plan_id", &planIDStr)
		if err != nil {
			// Skip if claim does not exist
			// Only log or ignore
			log.Info().Msg("No plan_id found in token")
		} else {
			planID, err := uuid.Parse(planIDStr)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			c.Locals("plan_id", planID)
		}

		// Set the user id
		var scope map[string]any
		if err = token.Get("user_scope", &scope); err != nil {
			log.Error().Err(err).Msg("No scope found in token")
			return c.Status(fiber.StatusUnauthorized).JSON(response.APIResponse[any]{
				Success: false,
				Error:   "Unauthorized",
			})
		} else {
			var claims types.Scoped
			data, _ := json.Marshal(scope)
			_ = json.Unmarshal(data, &claims)

			c.Locals("user_scope", claims)
		}

		c.Locals("user_id", userID)

		return c.Next()
	}
}
