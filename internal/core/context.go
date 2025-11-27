package core

import (
	"context"

	"github.com/google/uuid"
)

var ContextKey = struct{}{}

type Context struct {
	context.Context
	RequestID string
	UserID    uuid.UUID
	TenantID  *uuid.UUID
}

func NewCtx(ctx context.Context) *Context {
	return &Context{
		Context:   ctx,
		RequestID: uuid.NewString(),
	}
}
