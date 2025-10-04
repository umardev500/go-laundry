package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/domain/auth"
	"github.com/umardev500/go-laundry/internal/module/auth/dto"
	"github.com/umardev500/go-laundry/internal/utils/fiberutils"
	"github.com/umardev500/go-laundry/pkg/response"
	"github.com/umardev500/go-laundry/pkg/validator"

	appContext "github.com/umardev500/go-laundry/internal/app/context"
)

type Handler struct {
	service   Service
	validator *validator.Validator
}

func NewHandler(service Service, v *validator.Validator) *Handler {
	return &Handler{
		service:   service,
		validator: v,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/auth")

	r.Post("/login", h.Login)
	r.Post("/reset-password", h.ResetPassword)
	r.Post("/request-password-reset", h.RequestPasswordReset)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	scoped := fiberutils.GetScopedFromLocals(c)
	scopedContext := &appContext.ScopedContext{
		Context: ctx,
		Scoped:  scoped,
	}

	_, token, refreshToken, reso, err := h.service.Login(scopedContext, req.Email, req.Password)
	if err != nil {
		if err == auth.ErrMultipleAccountTypes {
			return c.Status(fiber.StatusConflict).JSON(response.APIResponse[any]{
				Success: false,
				Error:   err.Error(),
				Data: dto.LoginResolution{
					PlatformUser: reso.PlatformUser,
					TenantUsers:  reso.TenantUsers,
				},
			})
		}

		if err == auth.ErrMultipleTenants {
			return c.Status(fiber.StatusConflict).JSON(response.APIResponse[any]{
				Success: false,
				Error:   err.Error(),
				Data: dto.LoginResolution{
					PlatformUser: reso.PlatformUser,
					TenantUsers:  reso.TenantUsers,
				},
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	c.Status(fiber.StatusOK).JSON(
		response.APIResponse[dto.LoginResponse]{
			Success: true,
			Message: "Login successful",
			Data: dto.LoginResponse{
				Token:        token,
				RefreshToken: refreshToken,
			},
		},
	)

	return nil
}

func (h *Handler) ResetPassword(c *fiber.Ctx) error {
	var req dto.PasswordResetRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	scopedCtx := appContext.GetScopedContext(c)
	if scopedCtx == nil {
		return nil
	}

	_, token, refreshToken, reso, err := h.service.ResetPassword(scopedCtx, req.Token, req.Password)
	if err != nil {
		if err == auth.ErrMultipleAccountTypes {
			return c.Status(fiber.StatusConflict).JSON(response.APIResponse[any]{
				Success: false,
				Error:   err.Error(),
				Data: dto.LoginResolution{
					PlatformUser: reso.PlatformUser,
					TenantUsers:  reso.TenantUsers,
				},
			})
		}

		if err == auth.ErrMultipleTenants {
			return c.Status(fiber.StatusConflict).JSON(response.APIResponse[any]{
				Success: false,
				Error:   err.Error(),
				Data: dto.LoginResolution{
					PlatformUser: reso.PlatformUser,
					TenantUsers:  reso.TenantUsers,
				},
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(
		response.APIResponse[dto.LoginResponse]{
			Success: true,
			Message: "Password reset successful",
			Data: dto.LoginResponse{
				Token:        token,
				RefreshToken: refreshToken,
			},
		},
	)
}

func (h *Handler) RequestPasswordReset(c *fiber.Ctx) error {
	var req dto.RequestPasswordResetRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	scopedCtx := appContext.GetScopedContext(c)
	if scopedCtx == nil {
		return nil
	}

	err := h.service.RequestPasswordReset(scopedCtx, req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.APIResponse[any]{
		Success: true,
		Message: "Password reset request sent successfully",
	})
}
