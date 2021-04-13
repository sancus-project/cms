package view

import (
	"net/http"

	"go.sancus.dev/cms"
	"go.sancus.dev/cms/errors"
)

type View struct {
	config cms.ViewConfig
}

func (v *View) Middleware(next http.Handler) http.Handler {

	if next == nil {
		next = http.NotFoundHandler()
	}

	fn := func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				v.HandlePanic(w, r, rvr)
			}
		}()

		if err := v.Handler(w, r); err != nil {
			if e, ok := err.(*errors.ResourceError); ok {
				if e.Code == http.StatusNotFound {
					// 404
					next.ServeHTTP(w, r)
					return
				}
				v.HandleResourceError(w, r, e)
			} else {
				v.HandleError(w, r, err)
			}
		}
	}

	return http.HandlerFunc(fn)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Handler(w, r); err != nil {
		panic(err)
	}
}

func (v *View) Config() *cms.ViewConfig {
	return &v.config
}