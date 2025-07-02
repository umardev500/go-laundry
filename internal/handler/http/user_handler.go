package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/handler/http/middleware"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Setup(router fiber.Router) {
	router.Get("/users", middleware.AuthMiddleware, h.GetUsers)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	return c.JSON("hi")
}
