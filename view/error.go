package view

import (
	"log"
	"net/http"

	"go.sancus.dev/cms"
)

func (v View) HandleResourceError(w http.ResponseWriter, r *http.Request, err cms.Error) {

	if h := v.config.ResourceErrorHandler; h != nil {
		h(w, r, err)
		return
	}

	log.Printf("%T.HandleResourceError: %s%s: %v: %s", v, r.Host, r.URL.Path, err.Status(), err.Error())

	w.WriteHeader(err.Status())
	w.Write([]byte(err.Error()))
}

func (v View) HandleError(w http.ResponseWriter, r *http.Request, err error) {

	if e, ok := err.(cms.Error); ok {
		v.HandleResourceError(w, r, e)
		return
	}

	if h := v.config.ErrorHandler; h != nil {
		h(w, r, err)
		return
	}

	log.Printf("%T.HandleError: %s%s: %T: %s", v, r.Host, r.URL.Path, err, err.Error())

	e := &cms.HandlerError{Code: http.StatusInternalServerError, Err: err}
	v.HandleResourceError(w, r, e)
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
