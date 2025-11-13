package core

import (
	"context"

	"github.com/google/uuid"
)

type Context struct {
	context.Context
	RequestID string
}

func NewCtx(ctx context.Context) *Context {
	return &Context{
		Context:   ctx,
		RequestID: uuid.NewString(),
	}
}
