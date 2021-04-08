package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"go.sancus.dev/cms"
)

type ValidationError struct {
	Errors []error
}

func (e *ValidationError) AppendError(err error) {
	e.Errors = append(e.Errors, err)
}

func (e *ValidationError) AppendErrorString(s string) {
	e.AppendError(errors.New(s))
}

func (e *ValidationError) AppendErrorf(s string, args ...interface{}) {
	var err error

	if len(args) > 0 {
		err = fmt.Errorf(s, args...)
	} else {
		err = errors.New(s)
	}

	e.AppendError(err)
}

func (e *ValidationError) Ok() bool {
	return len(e.Errors) == 0
}

func (e *ValidationError) Code() int {
	if len(e.Errors) == 0 {
		return http.StatusOK
	} else {
		return http.StatusBadRequest
	}
}

func (e *ValidationError) String() string {
	var errors []string
	for _, err := range e.Errors {
		errors = append(errors, err.Error())
	}
	return strings.Join(errors, "\n")
}

func (e *ValidationError) Error() string {
	return e.String()
}

func (e *ValidationError) NewResourceError(r cms.Resource) ResourceError {
	return ResourceError{
		Code:     e.Code(),
		Resource: r,
		Wrapped:  e,
	}
}
