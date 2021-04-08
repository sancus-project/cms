package errors

import (
	"fmt"
	"net/http"

	"go.sancus.dev/cms"
)

type ResourceError struct {
	Code     int
	Err      error
	Resource cms.Resource
}

func (e ResourceError) Status() int {
	var code int

	if e.Code == 0 {
		if e.Err == nil {
			code = http.StatusOK
		} else {
			code = http.StatusInternalServerError
		}
	} else {
		code = e.Code
	}

	return code
}

func (e ResourceError) Unwrap() error {
	return e.Err
}

func (e ResourceError) String() string {
	code := e.Status()
	text := http.StatusText(code)

	if len(text) == 0 {
		return fmt.Sprintf("Unknown Error %d", code)
	} else if code < 400 {
		return text
	} else {
		return fmt.Sprintf("%s (Error %d)", text, code)
	}
}

func (e ResourceError) Error() string {
	return e.String()
}
