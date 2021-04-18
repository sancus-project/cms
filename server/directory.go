package server

import (
	"go.sancus.dev/cms/os"
	"go.sancus.dev/cms/os/utils"
)

type Directory struct {
	path     string
	fullpath string
	data     os.Directory
	cache    os.Directory
	server   *Server
}

func (d *Directory) Path() string {
	return d.path
}

func (s *Server) MkdirAll(path string) (*Directory, error) {
	path, err := utils.ValidPath(path)
	if err != nil {
		return nil, err
	}

	dir, err := s.root.MkdirAll(path)
	if err != nil {
		return nil, err
	}

	cache, err := s.cache.MkdirAll(path)
	if err != nil {
		return nil, err
	}

	d := &Directory{
		path:     path,
		fullpath: path,
		data:     dir,
		cache:    cache,
		server:   s,
	}

	return d, nil
}

func (s *Sandbox) MkdirAll(path string) (*Directory, error) {
	path, err := utils.ValidPath(path)
	if err != nil {
		return nil, err
	}

	// ask Server to create the directory
	d, err := s.server.MkdirAll(os.Join(s.root.Path(), path))
	if err != nil {
		return nil, err
	}

	// but make its Path() relative to the Sandbox
	d.path = path

	return d, nil
}
