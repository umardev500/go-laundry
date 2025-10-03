package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-laundry/internal/app/middleware"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/user"
	"github.com/umardev500/go-laundry/internal/module/user/dto"
	"github.com/umardev500/go-laundry/internal/utils/fiberutils"
	"github.com/umardev500/go-laundry/pkg/response"
	"github.com/umardev500/go-laundry/pkg/validator"
)

type Handler struct {
	cfg       *config.Config
	validator *validator.Validator
	service   user.Service
}

func NewHandler(cfg *config.Config, v *validator.Validator, service user.Service) *Handler {
	return &Handler{
		cfg:       cfg,
		validator: v,
		service:   service,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/users")

	r.Use(middleware.CheckAuth(h.cfg))
	r.Get("/", h.list)
	r.Post("/", h.createUser)
	r.Put("/profile", h.updateProfile)
	r.Delete("/:id", h.softDelete)
	r.Delete("/:id/purge", h.purge)
}

func (h *Handler) createUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var tenantIDPtr *uuid.UUID
	if val := c.Locals("tenant_id"); val != nil {
		if id, ok := val.(uuid.UUID); ok && id != uuid.Nil {
			tenantIDPtr = func() *uuid.UUID {
				return &id
			}()
		}
	}

	data, err := h.service.Create(c.Context(), req.ToUserCreate(tenantIDPtr))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*user.User]{
		Success: true,
		Message: "User created successfully",
		Data:    data,
	})
}

func (h *Handler) softDelete(c *fiber.Ctx) error {
	id, ok := fiberutils.GetUUIDParamOrAPIError(c, "id")
	if !ok {
		return nil // helper already wrote the response
	}

	scope := fiberutils.GetScopedFromCtx(c)
	if scope == nil {
		return nil
	}

	err := h.service.Delete(c.Context(), id, scope)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*user.User]{
		Success: true,
		Message: "User deleted successfully",
	})
}

func (h *Handler) list(c *fiber.Ctx) error {
	// Parse query params
	var filter user.Filter
	if err := c.QueryParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	scope := fiberutils.GetScopedFromCtx(c)
	if scope == nil {
		return nil
	}

	result, err := h.service.List(c.Context(), &filter, scope)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[[]*user.User]{
		Success:    true,
		Message:    "Users fetched successfully",
		Data:       result.Data,
		Pagination: result.Pagination,
	})
}

func (h *Handler) purge(c *fiber.Ctx) error {
	id, ok := fiberutils.GetUUIDParamOrAPIError(c, "id")
	if !ok {
		return nil // helper already wrote the response
	}

	scope := fiberutils.GetScopedFromCtx(c)
	if scope == nil {
		return nil
	}

	err := h.service.Purge(c.Context(), id, scope)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*user.User]{
		Success: true,
		Message: "User purged successfully",
	})
}

func (h *Handler) updateProfile(c *fiber.Ctx) error {
	var req dto.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userID := c.Locals("user_id").(uuid.UUID)

	data, err := h.service.UpdateProfile(
		c.Context(),
		userID,
		req.ToUserProfileUpdate(),
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*user.Profile]{
		Success: true,
		Message: "Profile updated successfully",
		Data:    data,
	})
}
