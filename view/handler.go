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
	ctx := h.v.config.SetDirectory(r.Context(), h.d)
	r = r.WithContext(ctx)

	return h.v.config.FilesHandler(w, r)
}

type EditHandler struct {
	r cms.Resource
	v *View
}

func (h EditHandler) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	ctx := h.v.config.SetResource(r.Context(), h.r)
	r = r.WithContext(ctx)

	return h.v.config.EditHandler(w, r)
}

type ResourceHandler struct {
	r cms.Resource
	v *View
}

func (h ResourceHandler) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {
	log.Println(errors.Here())
	return ErrNotImplemented
}
