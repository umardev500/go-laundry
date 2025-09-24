package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/user"
	"github.com/umardev500/go-laundry/internal/module/user/dto"
	"github.com/umardev500/go-laundry/pkg/response"
	"github.com/umardev500/go-laundry/pkg/validator"
)

type Handler struct {
	cfg       *config.Config
	validator *validator.Validator
	service   user.Service
}

func NewHandler(cfg *config.Config, v *validator.Validator, service user.Service) *Handler {
	return &Handler{
		cfg:       cfg,
		validator: v,
		service:   service,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/user")

	r.Use(middleware.CheckAuth(h.cfg))
	r.Put("/profile", h.updateProfile)
}

func (h *Handler) updateProfile(c *fiber.Ctx) error {
	var req dto.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID := c.Locals("user_id").(uuid.UUID)

	data, err := h.service.UpdateUserProfile(
		c.Context(),
		userID,
		req.ToUserProfileUpdate(),
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*user.Profile]{
		Success: true,
		Message: "Profile updated successfully",
		Data:    data,
	})
}
