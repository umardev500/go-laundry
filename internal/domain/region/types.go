package region

type OrderBy string

const (
	OrderByNameAsc  OrderBy = "name_asc"
	OrderByNameDesc OrderBy = "name_desc"
)

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Regency struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type District struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Village struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Filter struct {
	Query   string  `query:"query"`
	Offset  int     `query:"offset"`
	Limit   int     `query:"limit"`
	OrderBy OrderBy `query:"order_by"`
}

func (f Filter) WithDefaults() Filter {
	if f.Limit == 0 {
		f.Limit = 10 // default page size
	}
	if f.Offset == 0 {
		f.Offset = 0
	}
	if f.OrderBy == "" {
		f.OrderBy = OrderByNameAsc // default ordering
	}
	return f
}
