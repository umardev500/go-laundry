package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/dto"
	"github.com/umardev500/laundry/internal/mapper"
	"github.com/umardev500/laundry/internal/service"
	"github.com/umardev500/routerx"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

// Register implements app.Route
func (u *UserHandler) Register(app routerx.Router) {
	group := app.Group("/users")

	group.Get("/", u.Find)
	group.Get("/{id}", u.FindByID)
}

func (u *UserHandler) Find(c *routerx.Ctx) error {
	var query dto.UserFilter
	if err := c.QueryParser(&query); err != nil {
		return err
	}

	filter, err := query.ToDomain()
	if err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	ctx := core.NewCtx(c.Context())
	users, count, err := u.service.Find(ctx, *filter)
	if err != nil {
		return core.HandleError(c, err)
	}

	userDTOs := mapper.MapDomainUsersToDTOs(users)

	return core.NewPaginatedResponse(c, userDTOs, filter.Pagination, count)
}

func (u *UserHandler) FindByID(c *routerx.Ctx) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	ctx := core.NewCtx(c.Context())
	user, err := u.service.FindByID(ctx, id)
	if err != nil {
		return core.HandleError(c, err)
	}

	userDTO := mapper.MapDomainUserToDTO(user)

	return core.NewSuccessResponse(c, userDTO)
}
