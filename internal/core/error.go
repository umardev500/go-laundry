package core

import "net/http"

type Error struct {
	err        error
	message    string
	statusCode int
}

func NewError(err error, message string, statusCode ...int) *Error {
	code := http.StatusBadRequest

	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	return &Error{
		err:        err,
		message:    message,
		statusCode: code,
	}
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Unwrap() error {
	return e.err
}

func AsError(err error, target **Error) bool {
	if e, ok := err.(*Error); ok {
		*target = e
		return true
	}
	return false
}
