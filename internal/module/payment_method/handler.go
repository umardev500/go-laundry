package paymentmethod

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	paymentmethod "github.com/umardev500/go-laundry/internal/domain/payment_method"
	"github.com/umardev500/go-laundry/internal/module/payment_method/dto"
	"github.com/umardev500/go-laundry/internal/utils/fiberutils"
	"github.com/umardev500/go-laundry/pkg/response"
)

type Handler struct {
	service paymentmethod.Service
	cfg     *config.Config
}

func NewHandler(service paymentmethod.Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/payment-methods")

	r.Use(middleware.CheckAuth(h.cfg))
	r.Get("/", h.List)
	r.Post("/", h.Create)
	r.Get("/:id", h.GetByID)
	r.Put("/:id", h.Update)
	r.Delete("/:id", h.Delete)
}

// List all payment methods (tenant scoped unless platform admin)
func (h *Handler) List(c *fiber.Ctx) error {
	tenantIDPtr := fiberutils.GetTenantIDfromCtx(c)

	filter, err := h.parseFilter(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	paymentMethods, err := h.service.List(c.Context(), tenantIDPtr, &filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[[]*paymentmethod.PaymentMethod]{
		Success: true,
		Message: "Payment methods fetched successfully",
		Data:    paymentMethods,
	})
}

// Create a new payment method
func (h *Handler) Create(c *fiber.Ctx) error {
	tenantIDPtr := fiberutils.GetTenantIDfromCtx(c)

	var req dto.Create
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	payload := req.ToPaymentMethodCreate(*tenantIDPtr)
	pm, err := h.service.Create(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response.APIResponse[*paymentmethod.PaymentMethod]{
		Success: true,
		Message: "Payment method created successfully",
		Data:    pm,
	})
}

// Get payment method by ID
func (h *Handler) GetByID(c *fiber.Ctx) error {
	tenantIDPtr := fiberutils.GetTenantIDfromCtx(c)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid payment method ID",
		})
	}

	filter, err := h.parseFilter(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	pm, err := h.service.GetByID(c.Context(), tenantIDPtr, id, &filter)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*paymentmethod.PaymentMethod]{
		Success: true,
		Message: "Payment method fetched successfully",
		Data:    pm,
	})
}

// Update payment method metadata
func (h *Handler) Update(c *fiber.Ctx) error {
	tenantIDPtr := fiberutils.GetTenantIDfromCtx(c)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid payment method ID",
		})
	}

	var req dto.Update
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid request body",
		})
	}

	payload := req.ToPaymentMethodCreate(*tenantIDPtr)
	pm, err := h.service.Update(c.Context(), tenantIDPtr, id, payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*paymentmethod.PaymentMethod]{
		Success: true,
		Message: "Payment method updated successfully",
		Data:    pm,
	})
}

// Delete a payment method by ID
func (h *Handler) Delete(c *fiber.Ctx) error {
	tenantIDPtr := fiberutils.GetTenantIDfromCtx(c)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid payment method ID",
		})
	}

	if err := h.service.Delete(c.Context(), tenantIDPtr, id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[any]{
		Success: true,
		Message: "Payment method deleted successfully",
	})
}

func (h *Handler) parseFilter(c *fiber.Ctx) (paymentmethod.Filter, error) {
	var filter paymentmethod.Filter
	if err := c.QueryParser(&filter); err != nil {
		return paymentmethod.Filter{}, err
	}
	return filter, nil
}
