package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/module/auth/dto"
	"github.com/umardev500/go-laundry/pkg/response"
	"github.com/umardev500/go-laundry/pkg/validator"
)

type Handler struct {
	service   Service
	validator *validator.Validator
}

func NewHandler(service Service, v *validator.Validator) *Handler {
	return &Handler{
		service:   service,
		validator: v,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/auth")

	r.Post("/login", h.Login)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, token, refreshToken, err := h.service.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Status(fiber.StatusOK).JSON(
		response.APIResponse[dto.LoginResponse]{
			Success: true,
			Message: "Login successful",
			Data: dto.LoginResponse{
				Token:        token,
				RefreshToken: refreshToken,
			},
		},
	)

	return nil
}

func (h *Handler) Register(c *fiber.Ctx) error {
	return nil
}
