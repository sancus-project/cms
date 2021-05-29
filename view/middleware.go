package view

import (
	"net/http"
	"strings"

	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

type Middleware struct {
	View
	prefix string

	// Original ViewConfig.GetRoutePath
	getRoutePath func(r *http.Request) string
}

// ViewConfig.GetRoutePath override
func (v *Middleware) GetRoutePath(r *http.Request) string {
	path := v.getRoutePath(r)
	path = strings.TrimPrefix(path, v.prefix)

	if len(path) > 0 {
		return path
	} else {
		return "/"
	}
}

func (v *Middleware) Middleware(next http.Handler) http.Handler {
	if next == nil {
		next = http.NotFoundHandler()
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.MiddlewareHandler(w, r, next)
	})
}

func (v *Middleware) MiddlewareHandler(w http.ResponseWriter, r *http.Request, next http.Handler) {
	path := v.GetRoutePath(r)
	if path[0] != '/' {
		// not for me
		next.ServeHTTP(w, r)
		return
	}

	defer func() {
		if err := errors.Recover(); err != nil && err != http.ErrAbortHandler {
			v.config.ErrorHandler(w, r, err)
		}
	}()

	if err := v.TryServeHTTP(w, r); err != nil {
		if e, ok := err.(web.Error); ok {
			if e.Status() == http.StatusNotFound {
				// 404
				next.ServeHTTP(w, r)
				return
			}
		}
		v.config.ErrorHandler(w, r, err)
	}
}

func (v *View) NewMiddleware(prefix string) *Middleware {
	if prefix == "/" {
		prefix = ""
	}

	m := &Middleware{
		View: View{
			config: v.config,
			server: v.server,
		},
		prefix:       prefix,
		getRoutePath: v.config.GetRoutePath,
	}

	// override ViewConfig.GetRoutePath
	m.config.GetRoutePath = m.GetRoutePath
	return m
}
