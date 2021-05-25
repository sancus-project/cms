package view

import (
	"log"
	"net/http"

	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

func (v View) HandleError(w http.ResponseWriter, r *http.Request, err error) {

	if err == nil {
		return
	} else if h := v.config.ErrorHandler; h != nil {
		h(w, r, err)
		return
	}

	log.Printf("%+n: %s%s: %T: %s", errors.Here(0), r.Host, r.URL.Path, err, err.Error())

	// error knows how to render itself
	h, ok := err.(http.Handler)
	if !ok {
		var code int

		// but if it doesn't, wrap in HandlerError{}
		if e, ok := err.(web.Error); ok {
			code = e.Status()
		} else {
			code = http.StatusInternalServerError
		}

		h = &errors.HandlerError{Code: code, Err: err}
	}

	h.ServeHTTP(w, r)
}
