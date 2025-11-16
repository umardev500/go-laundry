package handler

import (
	"net/http"

	"github.com/umardev500/laundry/internal/core"
	"github.com/umardev500/laundry/internal/dto"
	"github.com/umardev500/laundry/internal/mapper"
	"github.com/umardev500/laundry/internal/service"
	"github.com/umardev500/routerx"
)

type TenantHandler struct {
	service *service.TenantService
}

func NewTenantHandler(s *service.TenantService) *TenantHandler {
	return &TenantHandler{
		service: s,
	}
}

// Register implements app.Route
func (h *TenantHandler) Register(app routerx.Router) {
	group := app.Group("/tenants")
	group.Get("/", h.Find)
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

	ctx := core.NewCtx(c.Context())
	tenants, count, err := h.service.Find(ctx, filter)
	if err != nil {
		return core.HandleError(c, err)
	}

	tenantDTOs := mapper.MapDomainTenantToDTOs(tenants)

	return core.NewPaginatedResponse(c, tenantDTOs, filter.Pagination, count)
}
