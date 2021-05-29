package view

import (
	"net/http"

	"go.sancus.dev/cms"
	"go.sancus.dev/web/errors"
)

type View struct {
	config cms.ViewConfig
	server cms.Directory
}

func NewView(s cms.Directory, cfg cms.ViewConfig) (cms.View, error) {
	v := &View{
		config: cfg,
		server: s,
	}

	if err := v.SetDefaults(); err != nil {
		return nil, err
	}

	return v, nil
}

func (v *View) SetDefaults() error {
	return v.config.SetDefaults()
}

func (v *View) Middleware(prefix string) func(http.Handler) http.Handler {
	return v.NewMiddleware(prefix).Middleware
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.TryServeHTTP(w, r); err != nil {
		v.config.ErrorHandler(w, r, err)
	}
}

func (v *View) TryServeHTTP(w http.ResponseWriter, r *http.Request) error {

	if p, ok := v.pageInfo(r); ok {
		return p.TryServeHTTP(w, r)
	} else {
		return errors.ErrNotFound
	}
}

func (v *View) PageInfo(r *http.Request) (interface{}, bool) {
	return v.pageInfo(r)
}

func (v *View) Config() *cms.ViewConfig {
	return &v.config
}
