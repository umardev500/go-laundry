package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/domain"
	"github.com/umardev500/go-laundry/internal/handler/http/middleware"
	"github.com/umardev500/go-laundry/pkg/response"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Setup(router fiber.Router) {
	router.Get("/", middleware.AuthMiddleware, h.GetUsers)
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	params := &domain.GetUsersParams{}

	c.QueryParser(params)

	// Parse pagination info
	if params.Limit == 0 {
		params.Limit = 10
	}

	if params.Page > 0 {
		params.Offset = (params.Page - 1) * params.Limit
	}

	// Get users
	users, total, err := h.service.GetAll(c.UserContext(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	data := response.NewPaginatedData(users, params.Page, params.Limit, total)

	return c.JSON(response.APIResponse{
		Success: true,
		Message: "Get users successfully",
		Data:    data,
	})
}
