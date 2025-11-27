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

	group.Post("/", u.Create)
	group.Get("/", u.Find)
	group.Get("/{id}", u.FindByID)
	group.Put("/{id}", u.Update)
	group.Put("/{id}/profile", u.UpdateProfile)
}

func (u *UserHandler) Create(c *routerx.Ctx) error {
	var req dto.CreateUserDTO
	if err := c.BodyParser(&req); err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	cmd, err := req.ToCmd()
	if err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	ctx := c.Locals(core.ContextKey).(*core.Context)
	user, err := u.service.Create(ctx, cmd)
	if err != nil {
		return core.HandleError(c, err)
	}

	userDTO := mapper.MapDomainUserToDTO(user)

	return core.NewSuccessResponse(c, userDTO)
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

	ctx := c.Locals(core.ContextKey).(*core.Context)
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

	ctx := c.Locals(core.ContextKey).(*core.Context)
	user, err := u.service.FindByID(ctx, id)
	if err != nil {
		return core.HandleError(c, err)
	}

	userDTO := mapper.MapDomainUserToDTO(user)

	return core.NewSuccessResponse(c, userDTO)
}

// Update is update a user
func (u *UserHandler) Update(c *routerx.Ctx) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return core.NewErrorResponse(c, err)
	}

	var req dto.UpdateUserDTO
	if err := c.BodyParser(&req); err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	// TODO: validate

	cmd, err := req.ToCmd()
	if err != nil {
		return core.NewErrorResponse(c, err)
	}

	ctx := c.Locals(core.ContextKey).(*core.Context)
	user, err := u.service.Update(ctx, id, cmd)
	if err != nil {
		return core.HandleError(c, err)
	}

	userDTO := mapper.MapDomainUserToDTO(user)

	return core.NewSuccessResponse(c, userDTO)
}

func (u *UserHandler) UpdateProfile(c *routerx.Ctx) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return core.NewErrorResponse(c, err)
	}

	var req dto.UpdateProfileDTO
	if err := c.BodyParser(&req); err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	// TODO: validate

	cmd, err := req.ToCmd()
	if err != nil {
		return core.NewErrorResponse(c, err)
	}

	ctx := c.Locals(core.ContextKey).(*core.Context)
	user, err := u.service.UpdateProfile(ctx, id, cmd)
	if err != nil {
		return core.HandleError(c, err)
	}

	userDTO := mapper.MapDomainUserToDTO(user)

	return core.NewSuccessResponse(c, userDTO)
}
