package server

import (
	"go.sancus.dev/cms"
	"go.sancus.dev/cms/os"
	"go.sancus.dev/cms/os/types"
	"go.sancus.dev/cms/os/utils"
)

type Directory struct {
	path     string
	fullpath string
	data     types.Directory
	cache    types.Directory
	root     cms.Directory
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

func (d *Directory) Chdir(path string) (cms.Directory, error) {
	// Allow Relative paths, within the same root
	path, err := utils.ValidPath(os.Join(d.path, path))
	if err != nil {
		return nil, err
	}

	return d.root.Chdir(path)
}

func (s *Server) Path() string {
	return "/"
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
	return s.mkdirAll(path)
}

func (s *Server) chdir(path string) (*Directory, error) {
	path, err := utils.ValidPath(path)
	if err != nil {
		return nil, err
	}

	data, err := s.root.Chdir(path)
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

func (s *Server) Chdir(path string) (cms.Directory, error) {
	return s.chdir(path)
}

func (s *Sandbox) Path() string {
	return "/"
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

func (s *Sandbox) Chdir(path string) (cms.Directory, error) {
	path, err := utils.ValidPath(path)
	if err != nil {
		return nil, err
	}

	// ask Server for the directory
	d, err := s.server.chdir(os.Join(s.root.Path(), path))
	if err != nil {
		return nil, err
	}

	// but make its Path() relative to the Sandbox
	d.path = path
	d.root = s

	return d, nil
}
