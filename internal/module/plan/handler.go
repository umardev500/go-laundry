package plan

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/plan"
	"github.com/umardev500/go-laundry/pkg/response"
)

type Handler struct {
	service plan.Service
	cfg     *config.Config
}

func NewHandler(service plan.Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/plans")
	r.Use(middleware.CheckAuth(h.cfg))
	r.Get("/", h.List)
}

func (h *Handler) List(c *fiber.Ctx) error {
	// Parse query params
	includeDeleted := c.QueryBool("include_deleted", false)

	filter := plan.PlanFilter{
		IncludeDeleted: includeDeleted,
	}

	plans, err := h.service.List(c.Context(), &filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[[]*plan.Plan]{
		Success: true,
		Message: "Plans fetched successfully",
		Data:    plans,
	})
}
