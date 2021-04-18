package view

import (
	"log"
	"net/http"

	"go.sancus.dev/cms"
)

func (v *View) Handler(w http.ResponseWriter, r *http.Request) error {

	path := v.config.GetRoutePath(r)
	log.Printf("%T.Handler: %s", v, path)

	return cms.ErrNotFound
}
