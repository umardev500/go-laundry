package payment

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/payment"
	"github.com/umardev500/go-laundry/internal/utils/fiberutils"
	"github.com/umardev500/go-laundry/pkg/response"
)

type Handler struct {
	service payment.Service
	cfg     *config.Config
}

func NewHandler(service payment.Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/payments")

	r.Use(middleware.CheckAuth(h.cfg))
	r.Get("/", h.List)
	r.Get("/:id", h.GetByID)
	r.Post("/", h.Create)
	r.Put("/:id", h.Update)
	r.Delete("/:id", h.Delete)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	return nil
}

func (h *Handler) List(c *fiber.Ctx) error {
	// Parse query params
	status := c.Query("status")
	typ := c.Query("type")

	filter := payment.PaymentFilter{
		Status: (*payment.Status)(&status),
		Type:   (*payment.ReferenceType)(&typ),
	}.WithDefaults()

	tenantIDPtr := fiberutils.GetTenantIDfromCtx(c)

	payments, err := h.service.List(c.Context(), &filter, tenantIDPtr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[[]*payment.Payment]{
		Success: true,
		Message: "Payments fetched successfully",
		Data:    payments,
	})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	return nil
}
