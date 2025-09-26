package subscription

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/subscription"
	"github.com/umardev500/go-laundry/internal/module/subscription/dto"
	"github.com/umardev500/go-laundry/internal/utils/fiberutils"
	"github.com/umardev500/go-laundry/pkg/response"
	"github.com/umardev500/go-laundry/pkg/validator"
)

type Handler struct {
	service   subscription.Service
	cfg       *config.Config
	validator *validator.Validator
}

func NewHandler(service subscription.Service, cfg *config.Config, v *validator.Validator) *Handler {
	return &Handler{
		service:   service,
		cfg:       cfg,
		validator: v,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/subscriptions")

	r.Use(middleware.CheckAuth(h.cfg))
	r.Get("/", h.List)
	r.Post("/", h.Create)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req dto.CreateSubscriptionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	tenantIDPtr := fiberutils.GetTenantIDfromCtx(c)

	sub, err := h.service.Create(c.Context(), req.ToSubscriptionCreate(*tenantIDPtr))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*subscription.Subscription]{
		Success: true,
		Message: "Subscription created successfully",
		Data:    sub,
	})
}

func (h *Handler) List(c *fiber.Ctx) error {
	// Parse query params
	includePlan := c.QueryBool("include_plan", false)
	includeTenant := c.QueryBool("include_tenant", false)

	filter := subscription.SubscriptionFilter{
		IncludePlan:   includePlan,
		IncludeTenant: includeTenant,
	}.WithDefaults()

	subs, err := h.service.List(c.Context(), &filter)
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
