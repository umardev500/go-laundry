package payment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/payment"
	"github.com/umardev500/go-laundry/internal/module/orchestrator"
	"github.com/umardev500/go-laundry/internal/module/payment/dto"
	"github.com/umardev500/go-laundry/internal/utils/fiberutils"
	"github.com/umardev500/go-laundry/pkg/response"
	"github.com/umardev500/go-laundry/pkg/validator"
)

type Handler struct {
	service             payment.Service
	cfg                 *config.Config
	validator           *validator.Validator
	paymentOrchestrator *orchestrator.PaymentService
}

func NewHandler(service payment.Service, cfg *config.Config, validator *validator.Validator, paymentOrchestrator *orchestrator.PaymentService) *Handler {
	return &Handler{
		service:             service,
		cfg:                 cfg,
		validator:           validator,
		paymentOrchestrator: paymentOrchestrator,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/payments")

	r.Use(middleware.CheckAuth(h.cfg))
	r.Get("/", h.List)
	r.Get("/:id", h.GetByID)

	r.Patch("/:id/send-payment-proof", h.SendPaymentProof)
	r.Patch("/:id/process-payment", h.ProcessPayment)
}

func (h *Handler) ProcessPayment(c *fiber.Ctx) error {
	id, ok := fiberutils.GetUUIDParamOrAPIError(c, "id")
	if !ok {
		return nil
	}

	userID := c.Locals("user_id").(uuid.UUID)
	tenantID := fiberutils.GetTenantIDfromCtx(c)

	result, err := h.paymentOrchestrator.ProcessPayment(c.Context(), id, userID, tenantID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*payment.Payment]{
		Data:    result,
		Success: true,
		Message: "Payment processed successfully",
	})
}

func (h *Handler) SendPaymentProof(c *fiber.Ctx) error {
	id, ok := fiberutils.GetUUIDParamOrAPIError(c, "id")
	if !ok {
		return nil
	}

	var req dto.SendPaymentProof
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

	payload := &payment.PaymentUpdate{
		ProofURL: &req.ProofURL,
	}

	userID := c.Locals("user_id").(uuid.UUID)
	tenantIDPtr := fiberutils.GetTenantIDfromCtx(c)
	paymentData, err := h.service.Update(c.Context(), payload, id, userID, tenantIDPtr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*payment.Payment]{
		Success: true,
		Message: "Payment proof sent successfully",
		Data:    paymentData,
	})
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) List(c *fiber.Ctx) error {
	// Parse query params
	var filter payment.Filter
	if err := c.QueryParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	tenantIDPtr := fiberutils.GetTenantIDfromCtx(c)

	result, err := h.service.List(c.Context(), filter.WithDefaults(), tenantIDPtr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[[]*payment.Payment]{
		Success:    true,
		Message:    "Payments fetched successfully",
		Data:       result.Data,
		Pagination: result.Pagination,
	})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	return nil
}
