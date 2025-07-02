package context

import "fmt"

var (
	ErrNotFound error = fmt.Errorf("user context not found")
)
