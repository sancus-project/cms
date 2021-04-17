package view

import (
	"log"
	"net/http"

	"go.sancus.dev/cms"
)

func (v View) HandleError(w http.ResponseWriter, r *http.Request, err error) {

	if h := v.config.ErrorHandler; h != nil {
		h(w, r, err)
		return
	}

	log.Printf("%T.HandleError: %s%s: %T: %s", v, r.Host, r.URL.Path, err, err.Error())

	// error knows how to render itself
	h, ok := err.(http.Handler)
	if !ok {
		var code int

		// but if it doesn't, wrap in HandlerError{}
		if e, ok := err.(cms.Error); ok {
			code = e.Status()
		} else {
			code = http.StatusInternalServerError
		}

		h = &cms.HandlerError{Code: code, Err: err}
	}

	h.ServeHTTP(w, r)
}

func (v View) HandlePanic(w http.ResponseWriter, r *http.Request, rvr interface{}) {

	if h := v.config.PanicHandler; h != nil {
		h(w, r, rvr)
		return
	}

	// TODO: backtrace
	log.Printf("%T.HandlePanic: %s%s: %T", v, r.Host, r.URL.Path, rvr)
	panic(rvr)
}
