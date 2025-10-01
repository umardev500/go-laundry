package region

import (
	"context"

	"github.com/umardev500/go-laundry/internal/types"
)

type Repository interface {
	ListProvinces(ctx context.Context, f Filter) ([]*Province, error)
	GetProvinceByID(ctx context.Context, id string) (*Province, error)

	ListRegenciesByProvince(ctx context.Context, provinceID string, f Filter) (*types.PageData[Regency], error)
	GetRegencyByID(ctx context.Context, id string) (*Regency, error)

	ListDistrictsByRegency(ctx context.Context, regencyID string, f Filter) (*types.PageData[District], error)
	GetDistrictByID(ctx context.Context, id string) (*District, error)

	ListVillagesByDistrict(ctx context.Context, districtID string, f Filter) (*types.PageData[Village], error)
	GetVillageByID(ctx context.Context, id string) (*Village, error)
}
