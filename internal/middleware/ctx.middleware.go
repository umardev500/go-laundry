package middleware

import (
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/routerx"
)

func ContextMiddleware() routerx.Handler {
	return func(c *routerx.Ctx) error {
		ctx := core.NewCtx(c.Context())

		c.Locals(core.ContextKey, ctx)

		return c.Next()
	}
}
