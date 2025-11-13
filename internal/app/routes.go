package app

import (
	"github.com/umardev500/laundry/internal/handler"
	"github.com/umardev500/routerx"
)

type Route interface {
	Register(app routerx.Router)
}

func NewRoutes(
	userHandler *handler.UserHandler,
) []Route {
	return []Route{
		userHandler,
	}
}
