package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/pkg/common"
)

type AuthHandler interface {
	common.Handler
}

type authHandler struct {
}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func (a *authHandler) Setup(router fiber.Router) {
	router.Get("/login", a.Login)
}

func (h *authHandler) Login(c *fiber.Ctx) error {

	return nil
}
