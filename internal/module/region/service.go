package region

import (
	"context"

	"github.com/umardev500/go-laundry/internal/domain/region"
	"github.com/umardev500/go-laundry/internal/types"
	"github.com/umardev500/go-laundry/internal/utils"
	"github.com/umardev500/go-laundry/pkg/response"
)

type serviceImpl struct {
	repo region.Repository
}

// Provinces
func (s *serviceImpl) ListProvinces(ctx context.Context, f region.Filter) ([]*region.Province, error) {
	return s.repo.ListProvinces(ctx, f)
}

func (s *serviceImpl) GetProvinceByID(ctx context.Context, id string) (*region.Province, error) {
	return s.repo.GetProvinceByID(ctx, id)
}

// Regencies
func (s *serviceImpl) ListRegenciesByProvince(ctx context.Context, provinceID string, f region.Filter) (*types.PageResult[region.Regency], error) {
	f = f.WithDefaults()

	pd, err := s.repo.ListRegenciesByProvince(ctx, provinceID, f)
	if err != nil {
		return nil, err
	}

	page := f.Offset + 1
	totalPages := utils.CalculateTotalPages(pd.Total, f.Limit)

	return &types.PageResult[region.Regency]{
		Data: pd.Data,
		Pagination: &response.Pagination{
			Page:       f.Offset + 1,
			PageSize:   f.Limit,
			TotalItems: pd.Total,
			TotalPages: utils.CalculateTotalPages(pd.Total, f.Limit),
			HasNext:    page < totalPages,
			HasPrev:    page > 1,
		},
	}, nil
}

func (s *serviceImpl) GetRegencyByID(ctx context.Context, id string) (*region.Regency, error) {
	return s.repo.GetRegencyByID(ctx, id)
}

// Districts
func (s *serviceImpl) ListDistrictsByRegency(ctx context.Context, regencyID string, f region.Filter) (*types.PageResult[region.District], error) {
	f = f.WithDefaults()

	pd, err := s.repo.ListDistrictsByRegency(ctx, regencyID, f)
	if err != nil {
		return nil, err
	}

	page := f.Offset + 1
	totalPages := utils.CalculateTotalPages(pd.Total, f.Limit)

	return &types.PageResult[region.District]{
		Data: pd.Data,
		Pagination: &response.Pagination{
			Page:       f.Offset + 1,
			PageSize:   f.Limit,
			TotalItems: pd.Total,
			TotalPages: utils.CalculateTotalPages(pd.Total, f.Limit),
			HasNext:    page < totalPages,
			HasPrev:    page > 1,
		},
	}, nil
}

func (s *serviceImpl) GetDistrictByID(ctx context.Context, id string) (*region.District, error) {
	return s.repo.GetDistrictByID(ctx, id)
}

// Villages
func (s *serviceImpl) ListVillagesByDistrict(ctx context.Context, districtID string, f region.Filter) (*types.PageResult[region.Village], error) {
	f = f.WithDefaults()

	pd, err := s.repo.ListVillagesByDistrict(ctx, districtID, f)
	if err != nil {
		return nil, err
	}

	page := f.Offset + 1
	totalPages := utils.CalculateTotalPages(pd.Total, f.Limit)

	return &types.PageResult[region.Village]{
		Data: pd.Data,
		Pagination: &response.Pagination{
			Page:       f.Offset + 1,
			PageSize:   f.Limit,
			TotalItems: pd.Total,
			TotalPages: utils.CalculateTotalPages(pd.Total, f.Limit),
			HasNext:    page < totalPages,
			HasPrev:    page > 1,
		},
	}, nil
}

func (s *serviceImpl) GetVillageByID(ctx context.Context, id string) (*region.Village, error) {
	return s.repo.GetVillageByID(ctx, id)
}

func NewService(repo region.Repository) region.Service {
	return &serviceImpl{
		repo: repo,
	}
}
