package server

import (
	"io/fs"
	"time"

	"go.sancus.dev/cms"
	"go.sancus.dev/cms/os"
	"go.sancus.dev/cms/os/types"
	"go.sancus.dev/cms/os/utils"
)

type Directory struct {
	name     string // relative to parent directory
	path     string // relative to root
	fullpath string // relative to server
	data     types.Directory
	cache    types.Directory
	root     cms.Directory
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Path() string {
	return d.path
}

func (d *Directory) IsDir() bool {
	return true
}

func (d *Directory) Size() int64 {
	return 0
}

func (d *Directory) Type() fs.FileMode {
	return fs.ModeDir
}

func (d *Directory) Mode() fs.FileMode {
	return fs.ModeDir
}

func (d *Directory) Info() (fs.FileInfo, error) {
	return d, nil
}

func (d *Directory) Sys() interface{} {
	return d
}

func (d *Directory) ModTime() time.Time {
	return d.data.ModTime()
}

func (d *Directory) MkdirAll(path string) (cms.Directory, error) {
	// Allow Relative paths, within the same root
	path, err := utils.Join(d.path, path)
	if err != nil {
		return nil, err
	}

	return d.root.MkdirAll(path)
}

func (d *Directory) Chdir(path string) (cms.Directory, error) {
	// Allow Relative paths, within the same root
	path, err := utils.Join(d.path, path)
	if err != nil {
		return nil, err
	}

	return d.root.Chdir(path)
}

func (d *Directory) Open(path string) (cms.Resource, error) {
	// Allow Relative paths, within the same root
	path, err := utils.Join(d.path, path)
	if err != nil {
		return nil, err
	}

	return d.root.Open(path)
}

func (s *Server) chdir(path string, mkdir bool) (*Directory, error) {
	// Validate directory
	path, err := utils.Join("/", path)
	if err != nil {
		return nil, err
	}

	// get data directory
	var data types.Directory

	if mkdir {
		data, err = s.root.MkdirAll(path)
	} else {
		data, err = s.root.Chdir(path)
	}

	if err != nil {
		return nil, err
	}

	// and cache directory
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
	return s.chdir(path, true)
}

func (s *Server) Chdir(path string) (cms.Directory, error) {
	return s.chdir(path, true)
}

func (s *Server) Open(path string) (cms.Resource, error) {
	path, err := utils.Join("/", path)
	if err != nil {
		return nil, err
	}

	return nil, os.ErrNotExist
}

func (s *Sandbox) chdir(path string, mkdir bool) (*Directory, error) {
	// Validate path
	path, err := utils.Join("/", path)
	if err != nil {
		return nil, err
	}

	// ask Server to create the directory
	fullpath := os.Join(s.Directory.Path(), path)

	d, err := s.server.chdir(fullpath, mkdir)
	if err != nil {
		return nil, err
	}

	// but make its Path() relative to the Sandbox
	d.path = path
	d.root = s

	return d, nil
}

func (s *Sandbox) MkdirAll(path string) (cms.Directory, error) {
	return s.chdir(path, true)
}

func (s *Sandbox) Chdir(path string) (cms.Directory, error) {
	return s.chdir(path, false)
}

func (s *Sandbox) Open(path string) (cms.Resource, error) {
	// Validate path
	path, err := utils.Join("/", path)
	if err != nil {
		return nil, err
	}

	// ask Server for the resource
	return s.server.Open(os.Join(s.Directory.Path(), path))
}
