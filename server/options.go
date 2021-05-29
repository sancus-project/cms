package server

import (
	"context"
	"net/http"

	"go.sancus.dev/cms"
	"go.sancus.dev/web"
)

// Server Options
type ServerOption interface {
	ApplyOption(*Server) error
}

type ServerOptionFunc func(*Server) error

func (f ServerOptionFunc) ApplyOption(s *Server) error {
	return f(s)
}

// Set ViewConfig.GetUser
func SetGetUser(getUser func(context.Context) cms.User) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetGetUser(getUser)
	})
}

// Set ViewConfig.GetRoutePath
func SetGetRoutePath(getRoutePath func(r *http.Request) string) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetGetRoutePath(getRoutePath)
	})
}

// Set ViewConfig.SetResource
func SetSetResource(setResource func(context.Context, cms.Resource) context.Context) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetSetResource(setResource)
	})
}

// Set ViewConfig.GetResource
func SetGetResource(getResource func(context.Context) cms.Resource) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetGetResource(getResource)
	})
}

// Set Server's default File editor
func SetEditHandler(path string, handler web.HandlerFunc) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetEditHandler(path, handler)
	})
}

// Set Server's default File manager
func SetFilesHandler(path string, handler web.HandlerFunc) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetFilesHandler(path, handler)
	})
}

// Set Server's default Heartbeat handler
func SetPingHandler(path string, handler web.HandlerFunc) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetPingHandler(path, handler)
	})
}

// Set Server's default Sitemap handler
func SetSitemapHandler(path string, handler web.HandlerFunc) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetSitemapHandler(path, handler)
	})
}

// Set Server's default error Handler
func SetErrorHandler(handler web.ErrorHandlerFunc) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetErrorHandler(handler)
	})
}

// Set Server's default resource for a given directory
func SetIndexPage(index string) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetIndexPage(index)
	})
}
