package audit

import "context"

type Service interface {
	// Create a new audit log
	Create(ctx context.Context, payload *Create) error
}
