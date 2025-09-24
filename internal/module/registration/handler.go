package registration

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/module/registration/dto"
	"github.com/umardev500/go-laundry/pkg/response"
)

type Handler struct {
	service *service
}

func NewHandler(service *service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/registration")

	r.Post("/register-tenant", h.Register)
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	_, err := h.service.RegisterTenant(c.Context(), req.ToRegisterInput())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response.APIResponse[any]{
		Success: true,
		Message: "Register successful",
		Data:    req,
	})
}
