package server

import (
	"context"

	"go.sancus.dev/cms"
)

// Server Options
type ServerOption interface {
	ApplyOption(*Server) error
}

type serverOption struct {
	apply func(*Server) error
}

func (opt serverOption) ApplyOption(s *Server) error {
	if opt.apply != nil {
		return opt.apply(s)
	}
	return nil
}

func ServerOptionFunc(apply func(*Server) error) ServerOption {
	return &serverOption{apply: apply}
}

// Set ViewConfig.GetUser
func (s *Server) SetGetUser(getUser func(context.Context) cms.User) error {
	if s != nil {
		s.ViewConfig.GetUser = getUser
	}
	return nil
}

func SetGetUser(getUser func(context.Context) cms.User) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetGetUser(getUser)
	})
}

// Set ViewConfig.GetRoutePath
func (s *Server) SetGetRoutePath(getRoutePath func(context.Context) string) error {
	if s != nil {
		s.ViewConfig.GetRoutePath = getRoutePath
	}
	return nil
}

func SetGetRoutePath(getRoutePath func(context.Context) string) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.SetGetRoutePath(getRoutePath)
	})
}

// Set ViewConfig.SetResource
func SetSetResource(setResource func(context.Context, cms.Resource) (context.Context, error)) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		s.ViewConfig.SetResource = setResource
		return nil
	})
}

// Set ViewConfig.GetResource
func SetGetResource(getResource func(context.Context) cms.Resource) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		s.ViewConfig.GetResource = getResource
		return nil
	})
}

// Set Server's default File editor
func SetEditHandler(path string, handler cms.Handler) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.ViewConfig.SetEditHandler(path, handler)
	})
}

// Set Server's default File manager
func SetFilesHandler(path string, handler cms.Handler) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.ViewConfig.SetFilesHandler(path, handler)
	})
}

// Set Server's default Heartbeat handler
func SetPingHandler(path string, handler cms.Handler) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.ViewConfig.SetPingHandler(path, handler)
	})
}

// Set Server's default Sitemap handler
func SetSitemapHandler(path string, handler cms.Handler) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		return s.ViewConfig.SetSitemapHandler(path, handler)
	})
}

// Set Server's default error Handler
func SetErrorHandler(handler cms.ErrorHandler) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		s.ViewConfig.ErrorHandler = handler
		return nil
	})
}

// Set Server's default recover Handler
func SetPanicHandler(handler cms.PanicHandler) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		s.ViewConfig.PanicHandler = handler
		return nil
	})
}

// Set Server's default resource for a given directory
func SetIndexPage(index string) ServerOption {
	return ServerOptionFunc(func(s *Server) error {
		s.ViewConfig.Index = index
		return nil
	})
}
