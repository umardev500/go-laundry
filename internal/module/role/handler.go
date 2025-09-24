package role

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/role"
	"github.com/umardev500/go-laundry/internal/module/role/dto"
	"github.com/umardev500/go-laundry/pkg/response"
	"github.com/umardev500/go-laundry/pkg/validator"
)

type Handler struct {
	service   role.Service
	validator *validator.Validator
	cfg       *config.Config
}

func NewHandler(service role.Service, v *validator.Validator, cfg *config.Config) *Handler {
	return &Handler{
		service:   service,
		validator: v,
		cfg:       cfg,
	}
}

// SetupRoutes registers role-related routes
func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/roles")

	r.Use(middleware.CheckAuth(h.cfg))
	r.Post("/", h.CreateRole)
	r.Get("/", h.ListRoles)
}

func (h *Handler) CreateRole(c *fiber.Ctx) error {
	var req dto.CreateRoleRequest
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

	// userID := c.Locals("user_id").(uuid.UUID)
	var tenantIDPtr *uuid.UUID
	if val := c.Locals("tenant_id"); val != nil {
		if id, ok := val.(uuid.UUID); ok && id != uuid.Nil {
			tenantIDPtr = func() *uuid.UUID {
				return &id
			}()
		}
	}

	data, err := h.service.CreateRole(c.Context(), req.ToRoleCreate(), tenantIDPtr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*role.Role]{
		Success: true,
		Message: fmt.Sprintf("Role %s created successfully", req.Name),
		Data:    data,
	})
}

func (h *Handler) ListRoles(c *fiber.Ctx) error {
	return nil
}
