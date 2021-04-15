package server

import (
	"context"

	"go.sancus.dev/cms/os"
)

// CMS Server
type Server struct {
	Root  string
	Cache string

	root  os.Filesystem
	cache os.Filesystem
}

func (s *Server) Connect(ctx context.Context) error {
	if s.root != nil || s.cache != nil {
		return os.EBUSY
	}

	root, err := os.NewFilesystem(ctx, s.Root)
	if err != nil {
		return err
	}

	cache, err := os.NewFilesystem(ctx, s.Cache)
	if err != nil {
		defer root.Close()
		return err
	}

	s.root, s.Root = root, root.Root()
	s.cache, s.Cache = cache, cache.Root()

	return nil
}

func (s *Server) Close() error {
	root, cache := s.root, s.cache

	if cache != nil {
		defer cache.Close()
		s.cache = nil
	}
	if root != nil {
		defer root.Close()
		s.root = nil
	}

	return nil
}

// Server Options
func NewServer(root, cache string, options ...ServerOption) (*Server, error) {
	s := &Server{
		Root:  root,
		Cache: cache,
	}

	for _ = range options {
	}

	return s, nil
}
