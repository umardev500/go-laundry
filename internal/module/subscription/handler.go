package subscription

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/subscription"
	"github.com/umardev500/go-laundry/pkg/response"
)

type Handler struct {
	service subscription.Service
	cfg     *config.Config
}

func NewHandler(service subscription.Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/subscriptions")

	r.Use(middleware.CheckAuth(h.cfg))
	r.Get("/", h.List)
}

func (h *Handler) List(c *fiber.Ctx) error {
	subs, err := h.service.List(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[[]*subscription.Subscription]{
		Success: true,
		Message: "Subscriptions fetched successfully",
		Data:    subs,
	})
}
