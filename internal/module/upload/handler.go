package upload

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/domain/upload"
	"github.com/umardev500/go-laundry/pkg/response"
)

type Handler struct {
	service upload.Service
}

func NewHandler(service upload.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/uploads")
	r.Post("/", h.Upload)
}

func (h *Handler) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	dest, err := h.service.SaveFile(c.Context(), file, "payment_proofs")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	fileURL := h.service.GetFileURL(dest, c.Secure())

	return c.JSON(response.APIResponse[any]{
		Success: true,
		Message: "File uploaded successfully",
		Data:    fileURL,
	})
}
