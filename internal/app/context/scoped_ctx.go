package context

import (
	"context"
)

type ScopedContext struct {
	context.Context
	Scoped *Scoped
}
