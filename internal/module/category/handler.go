package category

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	domain "github.com/umardev500/go-laundry/internal/domain/category"
	"github.com/umardev500/go-laundry/internal/module/category/dto"
	"github.com/umardev500/go-laundry/internal/utils/fiberutils"
	"github.com/umardev500/go-laundry/pkg/response"
)

type Handler struct {
	cfg     *config.Config
	service domain.Service
}

func NewHandler(cfg *config.Config, service domain.Service) *Handler {
	return &Handler{
		cfg:     cfg,
		service: service,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/categories")
	r.Use(middleware.CheckAuth(h.cfg))

	r.Get("/", h.list)
	r.Post("/", h.create)
	r.Put("/:id", h.update)
	r.Delete("/:id", h.delete)
}

// create a new category
func (h *Handler) create(c *fiber.Ctx) error {
	var req dto.Create
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{
				Success: false,
				Error:   "invalid body",
			})
	}

	tenantID := fiberutils.GetTenantIDfromCtx(c)
	categoryDomain := req.ToDomain(tenantID)

	data, err := h.service.Create(c.Context(), categoryDomain)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{
				Success: false,
				Error:   err.Error(),
			})
	}

	return c.JSON(response.APIResponse[*domain.Category]{
		Success: true,
		Message: "Category created",
		Data:    data,
	})
}

// list categories with filter
func (h *Handler) list(c *fiber.Ctx) error {
	filter := domain.Filter{
		Query:   c.Query("query"),
		Limit:   c.QueryInt("limit", 10),
		Offset:  c.QueryInt("offset", 0),
		OrderBy: domain.OrderBy(c.Query("order_by", string(domain.OrderByNameAsc))),
	}.WithDefaults()

	tenantID := fiberutils.GetTenantIDfromCtx(c)

	result, err := h.service.List(c.Context(), tenantID, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(response.APIResponse[any]{
				Success: false,
				Error:   err.Error(),
			})
	}

	return c.JSON(response.APIResponse[[]*domain.Category]{
		Success:    true,
		Message:    "Categories fetched",
		Data:       result.Data,
		Pagination: result.Pagination,
	})
}

// update category by ID
func (h *Handler) update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{
				Success: false,
				Error:   "invalid category id",
			})
	}

	var req dto.Update
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{
				Success: false,
				Error:   "invalid body",
			})
	}

	tenantID := fiberutils.GetTenantIDfromCtx(c)
	categoryDomain := req.ToDomain()

	data, err := h.service.Update(c.Context(), tenantID, id, categoryDomain)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{
				Success: false,
				Error:   err.Error(),
			})
	}

	return c.JSON(response.APIResponse[*domain.Category]{
		Success: true,
		Message: "Category updated",
		Data:    data,
	})
}

// delete category by ID
func (h *Handler) delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{
				Success: false,
				Error:   "invalid category id",
			})
	}

	tenantID := fiberutils.GetTenantIDfromCtx(c)

	if err := h.service.Delete(c.Context(), tenantID, id); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(response.APIResponse[any]{
				Success: false,
				Error:   err.Error(),
			})
	}

	return c.JSON(response.APIResponse[any]{
		Success: true,
		Message: "Category deleted",
	})
}
