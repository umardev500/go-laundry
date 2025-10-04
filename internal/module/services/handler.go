package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/module/services/dto"
	"github.com/umardev500/go-laundry/pkg/response"
	"github.com/umardev500/go-laundry/pkg/validator"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
	domain "github.com/umardev500/go-laundry/internal/domain/services"
)

type Handler struct {
	cfg       *config.Config
	service   domain.Service
	validator *validator.Validator
}

func NewHandler(cfg *config.Config, service domain.Service, v *validator.Validator) *Handler {
	return &Handler{
		cfg:       cfg,
		service:   service,
		validator: v,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/services")
	r.Use(middleware.CheckAuth(h.cfg), middleware.ScopedContextMiddleware())

	r.Get("/", h.list)
	r.Get("/:id", h.getByID)
	r.Post("/", h.create)
	r.Put("/:id", h.update)
	r.Delete("/:id", h.delete)
}

// create a new service
func (h *Handler) create(c *fiber.Ctx) error {
	var req dto.Create
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{Success: false, Error: "invalid body"})
	}

	// ✅ validation
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{Success: false, Error: err.Error()})
	}

	scopedCtx := appContext.GetScopedContext(c)
	if scopedCtx == nil {
		return nil
	}
	serviceDomain := req.ToDomain(scopedCtx.Scoped.TenantID)

	data, err := h.service.Create(scopedCtx, serviceDomain)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{Success: false, Error: err.Error()})
	}

	return c.JSON(response.APIResponse[*domain.Services]{
		Success: true,
		Message: "Service created",
		Data:    data,
	})
}

// list services with filter
func (h *Handler) list(c *fiber.Ctx) error {
	filter := &domain.Filter{
		Query:           c.Query("query"),
		Limit:           c.QueryInt("limit", 10),
		Offset:          c.QueryInt("offset", 0),
		OrderBy:         domain.OrderBy(c.Query("order_by", string(domain.OrderByNameAsc))),
		IncludeTenant:   c.QueryBool("include_tenant", false),
		IncludeCategory: c.QueryBool("include_category", false),
	}

	scopedCtx := appContext.GetScopedContext(c)
	if scopedCtx == nil {
		return nil
	}

	result, err := h.service.List(scopedCtx, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(response.APIResponse[any]{Success: false, Error: err.Error()})
	}

	return c.JSON(response.APIResponse[[]*domain.Services]{
		Success:    true,
		Message:    "Services fetched",
		Data:       result.Data,
		Pagination: result.Pagination,
	})
}

// get service by ID
func (h *Handler) getByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{Success: false, Error: "invalid service id"})
	}

	scopedCtx := appContext.GetScopedContext(c)
	if scopedCtx == nil {
		return nil
	}

	data, err := h.service.GetByID(scopedCtx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(response.APIResponse[any]{Success: false, Error: err.Error()})
	}

	return c.JSON(response.APIResponse[*domain.Services]{
		Success: true,
		Message: "Service retrieved",
		Data:    data,
	})
}

// update service by ID
func (h *Handler) update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{Success: false, Error: "invalid service id"})
	}

	var req dto.Update
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{Success: false, Error: "invalid body"})
	}

	// ✅ validation
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{Success: false, Error: err.Error()})
	}

	serviceDomain := req.ToDomain()

	scopedCtx := appContext.GetScopedContext(c)
	if scopedCtx == nil {
		return nil
	}

	data, err := h.service.Update(scopedCtx, id, serviceDomain)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{Success: false, Error: err.Error()})
	}

	return c.JSON(response.APIResponse[*domain.Services]{
		Success: true,
		Message: "Service updated",
		Data:    data,
	})
}

// delete service by ID
func (h *Handler) delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{Success: false, Error: "invalid service id"})
	}

	scopedCtx := appContext.GetScopedContext(c)
	if scopedCtx == nil {
		return nil
	}

	if err := h.service.Delete(scopedCtx, id); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{Success: false, Error: err.Error()})
	}

	return c.JSON(response.APIResponse[any]{
		Success: true,
		Message: "Service deleted",
	})
}
