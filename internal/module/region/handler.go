package region

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/go-laundry/internal/config"
	"github.com/umardev500/go-laundry/internal/domain/region"
	"github.com/umardev500/go-laundry/pkg/response"
)

type Handler struct {
	service region.Service
	cfg     *config.Config
}

func NewHandler(service region.Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
	}
}

func (h *Handler) SetupRoutes(router fiber.Router) {
	r := router.Group("/regions")

	// Provinces
	r.Get("/provinces", h.ListProvinces)
	r.Get("/provinces/:id", h.GetProvinceByID)

	// Regencies
	r.Get("/provinces/:provinceID/regencies", h.ListRegenciesByProvince)
	r.Get("/regencies/:id", h.GetRegencyByID)

	// Districts
	r.Get("/regencies/:regencyID/districts", h.ListDistrictsByRegency)
	r.Get("/districts/:id", h.GetDistrictByID)

	// Villages
	r.Get("/districts/:districtID/villages", h.ListVillagesByDistrict)
	r.Get("/villages/:id", h.GetVillageByID)
}

// Provinces
func (h *Handler) ListProvinces(c *fiber.Ctx) error {
	var f region.Filter
	if err := c.QueryParser(&f); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid query parameters",
		})
	}

	provs, err := h.service.ListProvinces(c.Context(), f)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[[]*region.Province]{
		Success: true,
		Message: "Provinces fetched successfully",
		Data:    provs,
	})
}

func (h *Handler) GetProvinceByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid province ID",
		})
	}

	prov, err := h.service.GetProvinceByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*region.Province]{
		Success: true,
		Message: "Province fetched successfully",
		Data:    prov,
	})
}

// Regencies
func (h *Handler) ListRegenciesByProvince(c *fiber.Ctx) error {
	provinceID := c.Params("provinceID")
	if provinceID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid province ID",
		})
	}

	var f region.Filter
	if err := c.QueryParser(&f); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid query parameters",
		})
	}

	result, err := h.service.ListRegenciesByProvince(c.Context(), provinceID, f)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[[]*region.Regency]{
		Success:    true,
		Message:    "Regencies fetched successfully",
		Data:       result.Data,
		Pagination: result.Pagination,
	})
}

func (h *Handler) GetRegencyByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid regency ID",
		})
	}

	reg, err := h.service.GetRegencyByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*region.Regency]{
		Success: true,
		Message: "Regency fetched successfully",
		Data:    reg,
	})
}

// Districts
func (h *Handler) ListDistrictsByRegency(c *fiber.Ctx) error {
	regencyID := c.Params("regencyID")
	if regencyID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid regency ID",
		})
	}

	var f region.Filter
	if err := c.QueryParser(&f); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid query parameters",
		})
	}

	result, err := h.service.ListDistrictsByRegency(c.Context(), regencyID, f)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[[]*region.District]{
		Success:    true,
		Message:    "Regencies fetched successfully",
		Data:       result.Data,
		Pagination: result.Pagination,
	})
}

func (h *Handler) GetDistrictByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid district ID",
		})
	}

	dist, err := h.service.GetDistrictByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*region.District]{
		Success: true,
		Message: "District fetched successfully",
		Data:    dist,
	})
}

// Villages
func (h *Handler) ListVillagesByDistrict(c *fiber.Ctx) error {
	districtID := c.Params("districtID")
	if districtID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid district ID",
		})
	}

	var f region.Filter
	if err := c.QueryParser(&f); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid query parameters",
		})
	}

	result, err := h.service.ListVillagesByDistrict(c.Context(), districtID, f)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[[]*region.Village]{
		Success:    true,
		Message:    "Regencies fetched successfully",
		Data:       result.Data,
		Pagination: result.Pagination,
	})
}

func (h *Handler) GetVillageByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.APIResponse[any]{
			Success: false,
			Error:   "Invalid village ID",
		})
	}

	vill, err := h.service.GetVillageByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.APIResponse[any]{
			Success: false,
			Error:   err.Error(),
		})
	}

	return c.JSON(response.APIResponse[*region.Village]{
		Success: true,
		Message: "Village fetched successfully",
		Data:    vill,
	})
}
