package context

import "fmt"

var (
	ErrNotFound error = fmt.Errorf("not found")
)
