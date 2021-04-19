package cms

import (
	"net/http"

	"go.sancus.dev/mix/types"
)

var (
	ErrNotFound = types.ErrNotFound
)

type (
	Error        = types.Error
	ErrorHandler = types.ErrorHandler
	HandlerError = types.HandlerError
)

// Error Handlers
type PanicHandler func(http.ResponseWriter, *http.Request, interface{})
