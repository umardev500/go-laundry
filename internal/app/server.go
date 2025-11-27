package app

import (
	"github.com/umardev500/routerx"
)

func NewServer() *routerx.App {
	app := routerx.New()

	app.Use(func(c *routerx.Ctx) error {
		return c.Next()
	})

	return app
}
