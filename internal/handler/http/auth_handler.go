package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/ent"
	"github.com/umardev500/go-laundry/internal/handler/http/middleware"
	"github.com/umardev500/go-laundry/pkg/response"
)

type AuthHandler struct {
	aService domain.AuthService
}

func NewAuthHandler(aService domain.AuthService) *AuthHandler {
	return &AuthHandler{
		aService: aService,
	}
}

func (h *AuthHandler) Setup(router fiber.Router) {
	router.Get("/me", middleware.AuthMiddleware, h.Me)
	router.Post("/login", h.Login)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var payload domain.LoginRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	resp, err := h.aService.Login(c.UserContext(), &payload)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(response.APIResponse{
				Success: false,
				Message: "User not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(response.APIResponse{
		Success: true,
		Message: "Logged in successfully",
		Data:    resp,
	})
}

func (h *AuthHandler) Me(c *fiber.Ctx) error {
	u, err := h.aService.Me(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(response.APIResponse{
		Success: true,
		Message: "Get user successfully",
		Data:    u,
	})
}
