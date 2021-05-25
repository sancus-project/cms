package view

import (
	"log"
	"net/http"

	"go.sancus.dev/web/errors"
)

func (v *View) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {

	path := v.config.GetRoutePath(r)
	log.Printf("%+n: %s", errors.Here(0), path)

	return errors.ErrNotFound
}
