package cms

import (
	"net/http"

	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

var (
	ErrNotFound = errors.ErrNotFound
)

type (
	Error        = web.Error
	ErrorHandler = web.ErrorHandlerFunc
	HandlerError = errors.HandlerError
)

// Error Handlers
type PanicHandler func(http.ResponseWriter, *http.Request, interface{})
