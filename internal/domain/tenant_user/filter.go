package tenantuser

type OrderBy string

const (
	OrderByCreatedAtAsc  OrderBy = "created_at_asc"
	OrderByCreatedAtDesc OrderBy = "created_at_desc"
	OrderByUpdatedAtAsc  OrderBy = "updated_at_asc"
	OrderByUpdatedAtDesc OrderBy = "updated_at_desc"
)

type Filter struct {
	Limit   int
	Offset  int
	OrderBy OrderBy
}

func (f *Filter) WithDefaults() *Filter {
	if f == nil {
		return &Filter{}
	}

	if f.Limit <= 0 {
		f.Limit = 10
	}
	if f.Offset < 0 {
		f.Offset = 0
	}
	if f.OrderBy == "" {
		f.OrderBy = OrderByCreatedAtAsc
	}
	return f
}
