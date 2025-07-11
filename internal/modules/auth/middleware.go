package auth

import (
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware interface {
	WithAuth() fiber.Handler
}

type authMiddleware struct {
}

func NewAuthMiddleware() AuthMiddleware {
	return &authMiddleware{}
}

func (a *authMiddleware) WithAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// token := c.Get("Authorization")
		// if token == "" {
		// 	return fiber.ErrUnauthorized
		// }

		// tokenStr := strings.TrimPrefix(token, "Bearer ")
		// claims := &Claims{}

		// _, err := a.jwtManager.Verify(tokenStr, claims)
		// if err != nil {
		// 	return fiber.ErrUnauthorized
		// }

		// ctx := context.WithValue(c.UserContext(), pkgAuth.ClaimsKey, claims)
		// c.SetUserContext(ctx)

		return c.Next()
	}
}
