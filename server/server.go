package server

import (
	"context"
	"io/fs"
	"time"

	"go.sancus.dev/cms"
	"go.sancus.dev/cms/os"
	"go.sancus.dev/cms/os/types"
)

// CMS Server
type Server struct {
	cms.ViewConfig

	Root  string
	Cache string

	root  types.Filesystem
	cache types.Filesystem
}

func (s *Server) Info() (fs.FileInfo, error) {
	return s, nil
}

func (s *Server) Sys() interface{} {
	return s
}

func (s *Server) IsDir() bool {
	return true
}

func (s *Server) Name() string {
	return "/"
}

func (s *Server) Path() string {
	return "/"
}

func (s *Server) Size() int64 {
	return 0
}

func (s *Server) Mode() fs.FileMode {
	return fs.ModeDir
}

func (s *Server) Type() fs.FileMode {
	return fs.ModeDir
}
func (s *Server) ModTime() time.Time {
	return s.root.ModTime()
}

func (s *Server) Connect(ctx context.Context) error {
	if s.root != nil || s.cache != nil {
		return types.EBUSY
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

func NewServer(root, cache string, options ...ServerOption) (*Server, error) {
	s := &Server{
		Root:  root,
		Cache: cache,
	}

	for _, opt := range options {
		if err := opt.ApplyOption(s); err != nil {
			return nil, err
		}
	}

	if err := s.SetDefaults(); err != nil {
		return nil, err
	}

	return s, nil
}
