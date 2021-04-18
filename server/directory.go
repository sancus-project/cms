package server

import (
	"go.sancus.dev/cms"
	"go.sancus.dev/cms/os"
	"go.sancus.dev/cms/os/utils"
)

type Directory struct {
	path     string
	fullpath string
	data     os.Directory
	cache    os.Directory
	root     cms.Server
}

func (d *Directory) Path() string {
	return d.path
}

func (d *Directory) MkdirAll(path string) (cms.Directory, error) {
	// Allow Relative paths, within the same root
	path, err := utils.ValidPath(os.Join(d.path, path))
	if err != nil {
		return nil, err
	}

	return d.root.MkdirAll(path)
}

func (s *Server) mkdirAll(path string) (*Directory, error) {
	path, err := utils.ValidPath(path)
	if err != nil {
		return nil, err
	}

	data, err := s.root.MkdirAll(path)
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
		data:     data,
		cache:    cache,
		root:     s,
	}

	return d, nil
}

func (s *Server) MkdirAll(path string) (cms.Directory, error) {
	d, err := s.mkdirAll(path)
	return d, err
}

func (s *Sandbox) MkdirAll(path string) (cms.Directory, error) {
	path, err := utils.ValidPath(path)
	if err != nil {
		return nil, err
	}

	// ask Server to create the directory
	d, err := s.server.mkdirAll(os.Join(s.root.Path(), path))
	if err != nil {
		return nil, err
	}

	// but make its Path() relative to the Sandbox
	d.path = path
	d.root = s

	return d, nil
}
