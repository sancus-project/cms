package view

import (
	"log"
	"net/http"

	"go.sancus.dev/cms/errors"
)

func (v *View) Handler(w http.ResponseWriter, r *http.Request) error {
	log.Printf("%T.Handler: %s%s", v, r.Host, r.URL.Path)

	return errors.ErrNotFound
}
