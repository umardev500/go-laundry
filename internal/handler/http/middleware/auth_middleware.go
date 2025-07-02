package middleware

import (
	"context"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/domain"
	sharedctx "github.com/umardev500/go-laundry/pkg/context"
	sharedjwt "github.com/umardev500/go-laundry/pkg/jwt"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Check if the Authorization header is present
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Check if the Authorization header starts with "Bearer "
	parts := strings.Split(authHeader, "Bearer ")
	if len(parts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get JWT configuration from env
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var claims domain.Claims

	// Check if the token is valid and get the claims
	_, err := sharedjwt.Parse(parts[1], &claims, []byte(secret))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	ctx := context.WithValue(c.UserContext(), sharedctx.ClaimsContextKey, &claims)
	c.SetUserContext(ctx)

	return c.Next()
}
