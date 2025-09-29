package paymentmethodtype

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/config"
	paymentmethodtype "github.com/umardev500/go-laundry/internal/domain/payment_method_type"
	"github.com/umardev500/go-laundry/internal/module/payment_method_type/dto"
	"github.com/umardev500/go-laundry/pkg/response"
	"github.com/umardev500/go-laundry/pkg/validator"
)

type Handler struct {
	cfg       *config.Config
	validator *validator.Validator
	service   paymentmethodtype.Service
}

func NewHandler(cfg *config.Config, v *validator.Validator, service paymentmethodtype.Service) *Handler {
	return &Handler{
		cfg:       cfg,
		validator: v,
		service:   service,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/payment-method-types")

	r.Get("/", h.list)
	r.Post("/", h.create)
	r.Get("/:id", h.getByID)
	r.Put("/:id", h.update)
	r.Delete("/:id", h.delete)
}

func (h *Handler) create(c *fiber.Ctx) error {
	var req dto.Create
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		})
	}

	data, err := h.service.Create(c.Context(), req.ToPaymentMethodTypeCreate())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Failed to create payment method type",
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*paymentmethodtype.PaymentMethodType]{
		Success: true,
		Message: "Payment method type created successfully",
		Data:    data,
	})
}

func (h *Handler) update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Invalid ID",
			Error:   err.Error(),
		})
	}

	var req dto.Update
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Validation failed",
			Error:   err.Error(),
		})
	}

	data, err := h.service.Update(c.Context(), id, req.ToPaymentMethodTypeUpdate())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Failed to update payment method type",
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*paymentmethodtype.PaymentMethodType]{
		Success: true,
		Message: "Payment method type updated successfully",
		Data:    data,
	})
}

func (h *Handler) delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Invalid ID",
			Error:   err.Error(),
		})
	}

	if err := h.service.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Failed to delete payment method type",
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[any]{
		Success: true,
		Message: "Payment method type deleted successfully",
	})
}

func (h *Handler) getByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Invalid ID",
			Error:   err.Error(),
		})
	}

	data, err := h.service.GetByID(c.Context(), id, nil)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Payment method type not found",
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*paymentmethodtype.PaymentMethodType]{
		Success: true,
		Message: "Payment method type retrieved successfully",
		Data:    data,
	})
}

func (h *Handler) list(c *fiber.Ctx) error {
	var filter paymentmethodtype.Filter
	if status := c.Query("status"); status != "" {
		s := paymentmethodtype.Status(status)
		filter.Status = &s
	}

	data, err := h.service.List(c.Context(), &filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Message: "Failed to list payment method types",
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[[]*paymentmethodtype.PaymentMethodType]{
		Success: true,
		Message: "Payment method types retrieved successfully",
		Data:    data,
	})
}
