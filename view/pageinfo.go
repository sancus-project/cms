package view

import (
	"log"
	"net/http"
	"strings"

	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

func (v *View) pageInfo(r *http.Request) (web.Handler, bool) {
	path := v.config.GetRoutePath(r)
	log.Printf("%+n: %s", errors.Here(), path)

	if p, ok := v.pageSitemap(path); ok {
		return p, true
	} else if p, ok := v.pagePing(path); ok {
		return p, true
	} else if p, ok := v.pageFiles(path); ok {
		return p, true
	} else if p, ok := v.pageEdit(path); ok {
		return p, true
	} else if p, ok := v.pageResource(path); ok {
		return p, true
	} else {
		return nil, false
	}
}

func (v *View) pageSitemap(path string) (web.Handler, bool) {
	if s0 := v.config.Sitemap; s0 != "" {
		if s1 := strings.TrimPrefix(path, s0); s1 != path {
			if s1 == "" || s1[0] == '/' {
				return v.config.SitemapHandler, true
			}
		}
	}

	return nil, false
}

func (v *View) pagePing(path string) (web.Handler, bool) {
	if s := v.config.Ping; s != "" && s == path {
		return v.config.PingHandler, true
	}
	return nil, false
}

func (v *View) pageFiles(path string) (web.Handler, bool) {
	if s0 := v.config.Files; s0 != "" {
		if s1 := strings.TrimSuffix(path, s0); s1 != path {

			if s1 == "" {
				s1 = "/"
			}

			if d, err := v.server.Chdir(s1); d != nil {
				if h, ok, done := v.pageCheckError(err); done {
					return h, ok
				}

				h := &DirectoryHandler{d, v}
				return h, true
			}
		}
	}
	return nil, false
}

func (v *View) pageEdit(path string) (web.Handler, bool) {
	if s0 := v.config.Edit; s0 != "" {
		if s1 := strings.TrimSuffix(path, s0); s1 != path {
			if r, err := v.server.Open(s1); r != nil {
				if h, ok, done := v.pageCheckError(err); done {
					return h, ok
				}

				h := &EditHandler{r, v}
				return h, true
			}
		}
	}
	return nil, false
}

func (v *View) pageResource(path string) (web.Handler, bool) {
	if r, err := v.server.Open(path); r != nil {
		if h, ok, done := v.pageCheckError(err); done {
			return h, ok
		}

		h := &ResourceHandler{r, v}
		return h, true
	}
	return nil, false
}

func (v *View) pageCheckError(err error) (web.Handler, bool, bool) {

	if err != nil {

		if e, ok := err.(web.Error); ok {

			if e.Status() == http.StatusNotFound {
				// PageInfo() -> nil, false
				return nil, false, true
			}

		}

		// PageInfo() -> err, true
		h := errors.NewFromError(err).(web.Handler)
		return h, true, true
	}

	// Carry on
	return nil, false, false
}
