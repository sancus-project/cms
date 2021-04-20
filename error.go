package cms

import (
	"net/http"

	"go.sancus.dev/mix/errors"
	"go.sancus.dev/mix/types"
)

var (
	ErrNotFound = errors.ErrNotFound
)

type (
	Error        = types.Error
	ErrorHandler = types.ErrorHandler
	HandlerError = errors.HandlerError
)

// Error Handlers
type PanicHandler func(http.ResponseWriter, *http.Request, interface{})
