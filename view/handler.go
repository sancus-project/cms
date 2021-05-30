package view

import (
	"log"
	"net/http"

	"go.sancus.dev/cms"
	"go.sancus.dev/web/errors"
)

var ErrNotImplemented = &errors.HandlerError{
	Code: http.StatusServiceUnavailable,
}

type DirectoryHandler struct {
	d cms.Directory
	v *View
}

func (h DirectoryHandler) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	log.Println(errors.Here(0))
	return ErrNotImplemented
}

type EditHandler struct {
	r cms.Resource
	v *View
}

func (h EditHandler) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	log.Println(errors.Here(0))
	return ErrNotImplemented
}

type ResourceHandler struct {
	r cms.Resource
	v *View
}

func (h ResourceHandler) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	log.Println(errors.Here(0))
	return ErrNotImplemented
}
