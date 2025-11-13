package core

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/routerx"
)

func HandleError(c *routerx.Ctx, err error) error {
	if err == nil {
		return nil
	}

	var e *Error
	if AsError(err, &e) {
		return NewErrorResponse(c, e, e.statusCode)
	}

	// Fallback for unknown error
	log.Error().Err(err).Msg("unknown error")
	return c.Status(http.StatusInternalServerError).JSON(map[string]string{
		"error": "internal server error",
	})
}
