package subscription

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	// Pathcing
	r.Patch("/:id/activate", h.Activate)
}

func (h *Handler) Activate(c *fiber.Ctx) error {
	id, ok := fiberutils.GetUUIDParamOrAPIError(c, "id")
	if !ok {
		return nil
	}

	userID := c.Locals("user_id").(uuid.UUID)

	sub, err := h.service.Activate(c.Context(), id, userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*subscription.Subscription]{
		Success: true,
		Message: "Subscription activated successfully",
		Data:    sub,
	})
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
	userID := c.Locals("user_id").(uuid.UUID)

	sub, err := h.service.Create(c.Context(), userID, req.ToSubscriptionCreate(*tenantIDPtr))
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

	var filter subscription.SubscriptionFilter
	if err := c.QueryParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	subs, err := h.service.List(c.Context(), filter.WithDefaults())
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
