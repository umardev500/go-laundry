package audit

import "context"

type Repository interface {
	Create(ctx context.Context, payload *Create) error
}
