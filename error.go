package cms

import (
	"net/http"
)

type Error interface {
	Error() string
	Status() int
	Unwrap() error
}

type ResourceErrorHandler func(http.ResponseWriter, *http.Request, Error)
type ErrorHandler func(http.ResponseWriter, *http.Request, error)
type PanicHandler func(http.ResponseWriter, *http.Request, interface{})
