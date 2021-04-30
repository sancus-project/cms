package view

import (
	"net/http"

	"go.sancus.dev/cms"
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
	cfg := &v.config

	// GetRoutePath
	if cfg.GetRoutePath == nil {
		cfg.GetRoutePath = func(r *http.Request) string {
			return r.URL.Path
		}
	}

	return nil
}

func (v *View) Middleware(prefix string) func(http.Handler) http.Handler {
	return v.NewMiddleware(prefix).Middleware
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.TryServeHTTP(w, r); err != nil {
		v.HandleError(w, r, err)
	}
}

func (v *View) Config() *cms.ViewConfig {
	return &v.config
}
