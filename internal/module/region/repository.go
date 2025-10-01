package region

import (
	"context"

	"github.com/umardev500/go-laundry/ent"
	"github.com/umardev500/go-laundry/internal/db"
	"github.com/umardev500/go-laundry/internal/domain/region"
	"github.com/umardev500/go-laundry/internal/types"

	districtEntity "github.com/umardev500/go-laundry/ent/district"
	provinceEntity "github.com/umardev500/go-laundry/ent/province"
	regencyEntity "github.com/umardev500/go-laundry/ent/regency"
	villageEntity "github.com/umardev500/go-laundry/ent/village"
)

type repositoryImpl struct {
	client *db.Client
}

// GetDistrictByID implements region.Repository.
func (r *repositoryImpl) GetDistrictByID(ctx context.Context, id string) (*region.District, error) {
	conn := r.client.GetConn(ctx)

	q := conn.District.
		Query().
		Where(districtEntity.IDEQ(id))

	entObj, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	reg := MapEntity([]*ent.District{entObj}, func(ent *ent.District) *region.District {
		return &region.District{
			ID:   ent.ID,
			Name: ent.Name,
		}
	})

	return reg[0], nil
}

// GetVillageByID implements region.Repository.
func (r *repositoryImpl) GetVillageByID(ctx context.Context, id string) (*region.Village, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Village.
		Query().
		Where(villageEntity.IDEQ(id))

	entObj, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	reg := MapEntity([]*ent.Village{entObj}, func(ent *ent.Village) *region.Village {
		return &region.Village{
			ID:   ent.ID,
			Name: ent.Name,
		}
	})

	return reg[0], nil
}

// ListDistrictsByRegency implements region.Repository.
func (r *repositoryImpl) ListDistrictsByRegency(ctx context.Context, regencyID string, f region.Filter) (*types.PageData[region.District], error) {
	conn := r.client.GetConn(ctx)

	q := conn.District.
		Query().
		Where(districtEntity.HasRegencyWith(regencyEntity.IDEQ(regencyID)))

	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	if f.Query != "" {
		q = q.Where(districtEntity.NameContainsFold(f.Query))
	}

	// Pagination
	q = q.Limit(f.Limit).Offset(f.Offset)

	districts, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	data := MapEntity(districts, func(ent *ent.District) *region.District {
		return &region.District{
			ID:   ent.ID,
			Name: ent.Name,
		}
	})

	return &types.PageData[region.District]{
		Data:  data,
		Total: total,
	}, nil
}

// ListVillagesByDistrict implements region.Repository.
func (r *repositoryImpl) ListVillagesByDistrict(ctx context.Context, districtID string, f region.Filter) (*types.PageData[region.Village], error) {
	conn := r.client.GetConn(ctx)

	q := conn.Village.
		Query().
		Where(villageEntity.HasDistrictWith(districtEntity.IDEQ(districtID)))

	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	if f.Query != "" {
		q = q.Where(villageEntity.NameContainsFold(f.Query))
	}

	// Pagination
	q = q.Limit(f.Limit).Offset(f.Offset)

	villages, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	data := MapEntity(villages, func(ent *ent.Village) *region.Village {
		return &region.Village{
			ID:   ent.ID,
			Name: ent.Name,
		}
	})

	return &types.PageData[region.Village]{
		Data:  data,
		Total: total,
	}, nil
}

// GetRegencyByID implements region.Repository.
func (r *repositoryImpl) GetRegencyByID(ctx context.Context, id string) (*region.Regency, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Regency.
		Query().
		Where(regencyEntity.IDEQ(id))

	entObj, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	reg := MapEntity([]*ent.Regency{entObj}, func(ent *ent.Regency) *region.Regency {
		return &region.Regency{
			ID:   ent.ID,
			Name: ent.Name,
		}
	})

	return reg[0], nil
}

// ListRegenciesByProvince implements region.Repository.
func (r *repositoryImpl) ListRegenciesByProvince(ctx context.Context, provinceID string, f region.Filter) (*types.PageData[region.Regency], error) {
	conn := r.client.GetConn(ctx)

	q := conn.Regency.
		Query().
		Where(regencyEntity.HasProvinceWith(provinceEntity.IDEQ(provinceID)))

	// Count total
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	if f.Query != "" {
		q = q.Where(regencyEntity.NameContainsFold(f.Query))
	}

	// Ordering
	switch f.OrderBy {
	case region.OrderByNameAsc:
		q = q.Order(ent.Asc(regencyEntity.FieldName))
	case region.OrderByNameDesc:
		q = q.Order(ent.Desc(regencyEntity.FieldName))
	}

	regencies, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	data := MapEntity(regencies, func(ent *ent.Regency) *region.Regency {
		return &region.Regency{
			ID:   ent.ID,
			Name: ent.Name,
		}
	})

	return &types.PageData[region.Regency]{
		Data:  data,
		Total: total,
	}, nil
}

// GetProvinceByID implements province.Repository.
func (r *repositoryImpl) GetProvinceByID(ctx context.Context, id string) (*region.Province, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Province.
		Query().
		Where(provinceEntity.IDEQ(id))

	entObj, err := q.Only(ctx)
	if err != nil {
		return nil, err
	}

	prov := MapEntity([]*ent.Province{entObj}, func(ent *ent.Province) *region.Province {
		return &region.Province{
			ID:   ent.ID,
			Name: ent.Name,
		}
	})

	return prov[0], nil
}

// ListProvinces implements province.Repository.
func (r *repositoryImpl) ListProvinces(ctx context.Context, f region.Filter) ([]*region.Province, error) {
	conn := r.client.GetConn(ctx)

	q := conn.Province.Query()

	if f.Query != "" {
		q = q.Where(provinceEntity.NameContains(f.Query))
	}

	// Ordering
	switch f.OrderBy {
	case region.OrderByNameAsc:
		q = q.Order(ent.Asc(provinceEntity.FieldName))
	case region.OrderByNameDesc:
		q = q.Order(ent.Desc(provinceEntity.FieldName))
	}

	// Pagination
	q = q.Offset(f.Offset).Limit(f.Limit)

	entObjs, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	provinces := MapEntity(entObjs, func(ent *ent.Province) *region.Province {
		return &region.Province{
			ID:   ent.ID,
			Name: ent.Name,
		}
	})

	return provinces, nil
}

func NewRepository(client *db.Client) region.Repository {
	return &repositoryImpl{
		client: client,
	}
}
