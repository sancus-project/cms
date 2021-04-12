package errors

import (
	"net/http"
)

var (
	ErrNotFound = &ResourceError{Code: http.StatusNotFound}
)
