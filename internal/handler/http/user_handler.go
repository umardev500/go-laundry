package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/domain"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Setup(routers ...fiber.Router) {
	if len(routers) != 2 {
		log.Panic().Msg("expected exactly two routers: regular and protected")
	}

	regular := routers[0]
	// protected := routers[1]

	regular.Get("/users", h.GetUsers)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	return c.JSON("hi")
}
