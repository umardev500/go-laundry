package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/ent"
)

type AuthHandler struct {
	aService domain.AuthService
}

func NewAuthHandler(aService domain.AuthService) *AuthHandler {
	return &AuthHandler{
		aService: aService,
	}
}

func (h *AuthHandler) Setup(routers ...fiber.Router) {
	if len(routers) != 2 {
		log.Panic().Msg("expected exactly two routers: regular and protected")
	}

	regular := routers[0]
	// protected := routers[1]

	regular.Post("/auth/login", h.Login)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var payload domain.LoginRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.APIResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	resp, err := h.aService.Login(c.UserContext(), &payload)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(domain.APIResponse{
				Success: false,
				Message: "User not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(domain.APIResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(domain.APIResponse{
		Success: true,
		Message: "Logged in successfully",
		Data:    resp,
	})
}
