package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/handler/http/middleware"
	"github.com/umardev500/go-laundry/pkg/response"
)

type MerchantHandler struct {
	usecase  domain.MerchantUsecase
	validate *validator.Validate
}

func NewMerchantHandler(merchantUsecase domain.MerchantUsecase, validate *validator.Validate) *MerchantHandler {
	return &MerchantHandler{
		usecase:  merchantUsecase,
		validate: validate,
	}
}

func (h *MerchantHandler) Setup(router fiber.Router) {
	router.Post("/", middleware.AuthMiddleware, h.Register)
}

func (h *MerchantHandler) Register(c *fiber.Ctx) error {
	var payload domain.CreateMerchantRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	// Validate
	if err := h.validate.Struct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse{
			Success: false,
			Message: "Validation error",
			Errors:  []string{err.Error()},
		})
	}

	err := h.usecase.Register(c.UserContext(), &payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse{
			Success: false,
			Message: "Failed to register merchant",
			Errors:  []string{err.Error()},
		})
	}

	return c.JSON(response.APIResponse{
		Success: true,
		Message: "Merchant registered successfully",
	})
}
