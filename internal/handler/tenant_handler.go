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

type TenantHandler struct {
	service           *service.TenantService
	tenantUserService *service.TenantUserService
}

func NewTenantHandler(s *service.TenantService, tus *service.TenantUserService) *TenantHandler {
	return &TenantHandler{
		service:           s,
		tenantUserService: tus,
	}
}

// Register implements app.Route
func (h *TenantHandler) Register(app routerx.Router) {
	group := app.Group("/tenants")
	group.Post("/", h.Create)
	group.Delete("/{id}", h.Delete)
	group.Get("/", h.Find)
	group.Get("/{id}", h.FindByID)
	group.Put("/{id}", h.Update)

	// Tenant users
	group.Get("/{id}/users", h.FindUsers)
}

func (h *TenantHandler) Create(c *routerx.Ctx) error {
	var req dto.CreateTenant
	if err := c.BodyParser(&req); err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	// TODO: validate
	//

	cmd, err := req.ToCmd()
	if err != nil {
		return core.NewErrorResponse(c, err)
	}

	ctx := c.Locals(core.ContextKey).(*core.Context)
	result, err := h.service.Create(ctx, cmd)
	if err != nil {
		return core.HandleError(c, err)
	}

	return core.NewSuccessResponse(c, mapper.MapDomainTenantToDTO(result))
}

func (h *TenantHandler) Delete(c *routerx.Ctx) error {
	var idStr = c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	ctx := c.Locals(core.ContextKey).(*core.Context)
	err = h.service.Delete(ctx, id)
	if err != nil {
		return core.HandleError(c, err)
	}

	return c.SendStatus(http.StatusNoContent)
}

func (h *TenantHandler) Find(c *routerx.Ctx) error {
	var query dto.TenantFilter
	if err := c.QueryParser(&query); err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	filter, err := query.ToDomain()
	if err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	ctx := c.Locals(core.ContextKey).(*core.Context)
	tenants, count, err := h.service.Find(ctx, filter)
	if err != nil {
		return core.HandleError(c, err)
	}

	tenantDTOs := mapper.MapDomainTenantToDTOs(tenants)

	return core.NewPaginatedResponse(c, tenantDTOs, filter.Pagination, count)
}

func (h *TenantHandler) FindByID(c *routerx.Ctx) error {
	var idStr = c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	ctx := c.Locals(core.ContextKey).(*core.Context)
	result, err := h.service.FindByID(ctx, id)
	if err != nil {
		return core.HandleError(c, err)
	}

	return core.NewSuccessResponse(c, mapper.MapDomainTenantToDTO(result))
}

func (h *TenantHandler) Update(c *routerx.Ctx) error {
	var idStr = c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	var req dto.UpdateTenant
	if err := c.BodyParser(&req); err != nil {
		return core.NewErrorResponse(c, err)
	}

	// TODO: validate

	ctx := c.Locals(core.ContextKey).(*core.Context)
	cmd, err := req.ToCmd()
	if err != nil {
		return core.NewErrorResponse(c, err)
	}

	result, err := h.service.Update(ctx, id, cmd)
	if err != nil {
		return core.HandleError(c, err)
	}

	return core.NewSuccessResponse(c, mapper.MapDomainTenantToDTO(result))
}

// --- Tenant User Handlers ---
func (h *TenantHandler) FindUsers(c *routerx.Ctx) error {
	var query dto.TenantUserFilter
	if err := c.QueryParser(&query); err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	filter, err := query.ToDomain()
	if err != nil {
		return core.NewErrorResponse(c, err, http.StatusBadRequest)
	}

	ctx := c.Locals(core.ContextKey).(*core.Context)
	tenantUsers, count, err := h.tenantUserService.Find(ctx, filter)
	if err != nil {
		return core.HandleError(c, err)
	}

	mappedData := mapper.MapDomainTenantUserToDTOs(tenantUsers)

	return core.NewPaginatedResponse(
		c,
		mappedData,
		filter.Pagination,
		count,
	)
}
