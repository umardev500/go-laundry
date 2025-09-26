package plan

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/plan"
	"github.com/umardev500/go-laundry/internal/module/plan/dto"
	"github.com/umardev500/go-laundry/internal/utils/fiberutils"
	"github.com/umardev500/go-laundry/pkg/response"
	"github.com/umardev500/go-laundry/pkg/validator"
)

type Handler struct {
	service   plan.Service
	cfg       *config.Config
	validator *validator.Validator
}

func NewHandler(service plan.Service, cfg *config.Config, v *validator.Validator) *Handler {
	return &Handler{
		service:   service,
		cfg:       cfg,
		validator: v,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/plans")
	r.Use(middleware.CheckAuth(h.cfg))
	r.Get("/", h.List)
	r.Get("/:id", h.GetByID)

	r.Post("/:id/permissions", h.AddPermissions)

	r.Delete("/:id/permissions", h.RemovePermissions)
	r.Put("/:id/permissions", h.ReplacePermissions)
}

func (h *Handler) AddPermissions(c *fiber.Ctx) error {
	planID, ok := fiberutils.GetUUIDParamOrAPIError(c, "id")
	if !ok {
		return nil
	}

	var req dto.SetPermissionsRequest
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

	err := h.service.AddPermissions(c.Context(), planID, req.PermissionIDs)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[any]{
		Success: true,
		Message: "Permissions added successfully",
	})
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	planID, ok := fiberutils.GetUUIDParamOrAPIError(c, "id")
	if !ok {
		return nil
	}

	// Parse query params
	includeDeleted := c.QueryBool("include_deleted", false)
	IncludePermissions := c.QueryBool("include_permissions", false)

	filter := plan.PlanFilter{
		IncludeDeleted:     includeDeleted,
		IncludePermissions: IncludePermissions,
	}.WithDefaults()

	planData, err := h.service.GetByID(c.Context(), planID, &filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*plan.Plan]{
		Success: true,
		Message: "Plan fetched successfully",
		Data:    planData,
	})
}

func (h *Handler) List(c *fiber.Ctx) error {
	// Parse query params
	includeDeleted := c.QueryBool("include_deleted", false)
	IncludePermissions := c.QueryBool("include_permissions", false)

	filter := plan.PlanFilter{
		IncludeDeleted:     includeDeleted,
		IncludePermissions: IncludePermissions,
	}.WithDefaults()

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

func (h *Handler) RemovePermissions(c *fiber.Ctx) error {
	planID, ok := fiberutils.GetUUIDParamOrAPIError(c, "id")
	if !ok {
		return nil
	}

	var req dto.SetPermissionsRequest
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

	err := h.service.RemovePermissions(c.Context(), planID, req.PermissionIDs)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[any]{
		Success: true,
		Message: "Permissions removed successfully",
	})
}

func (h *Handler) ReplacePermissions(c *fiber.Ctx) error {
	planID, ok := fiberutils.GetUUIDParamOrAPIError(c, "id")
	if !ok {
		return nil
	}

	var req dto.SetPermissionsRequest
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

	err := h.service.ReplacePermissions(c.Context(), planID, req.PermissionIDs)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[any]{
		Success: true,
		Message: "Permissions replaced successfully",
	})
}
