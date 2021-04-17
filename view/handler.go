package view

import (
	"log"
	"net/http"

	"go.sancus.dev/cms/errors"
)

func (v *View) Handler(w http.ResponseWriter, r *http.Request) error {
	var path string

	if v.config.GetRoutePath != nil {
		path = v.config.GetRoutePath(r.Context())
	} else {
		path = r.URL.Path
	}

	log.Printf("%T.Handler: %s", v, path)

	return errors.ErrNotFound
}
