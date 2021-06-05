package local

import (
	"io/fs"
	"os"

	"go.sancus.dev/cms/os/types"
	"go.sancus.dev/cms/os/utils"
)

const (
	NewDirectoryMode = os.FileMode(0755)
)

type Directory struct {
	fs.FileInfo

	path   string // relative to root
	actual string // relative to server
}

func (d Directory) Name() string {
	if d.path == "/" {
		return "/"
	} else {
		return d.FileInfo.Name()
	}
}

func (d Directory) Path() string {
	return d.path
}

func (d Directory) Size() int64 {
	return 0
}

func (d Directory) Type() fs.FileMode {
	return d.Mode()
}

func (d *Directory) Sys() interface{} {
	return d
}

func (d *Directory) Info() (fs.FileInfo, error) {
	return d, nil
}

func chdir(path string, actual string, mkdir bool) (types.Directory, error) {

	if mkdir {
		var mode = NewDirectoryMode

		if err := os.MkdirAll(actual, mode); err != nil {
			return nil, err
		}
	}

	if fi, err := os.Stat(actual); err != nil {
		return nil, err
	} else if !fi.IsDir() {
		return nil, os.ErrNotExist
	} else {
		next := &Directory{
			FileInfo: fi,
			path:     path,
			actual:   actual,
		}

		return next, nil
	}
}

func (d Directory) chdir(path string, mkdir bool) (types.Directory, error) {
	// Allow Relative paths, within the same root
	s0, err := utils.Join(d.path, path)
	if err != nil {
		return nil, err
	}

	s1, err := utils.Join(d.actual, path)
	if err != nil {
		return nil, err
	}

	return chdir(s0, s1, mkdir)
}

func (fs *Filesystem) chdir(path string, mkdir bool) (types.Directory, error) {
	// validate path
	s0, err := utils.Join(path)
	if err != nil {
		return nil, err
	}

	fs.mu.Lock()
	defer fs.mu.Unlock()

	s1, err := utils.Join(fs.root, s0)
	if err != nil {
		return nil, err
	}

	return chdir(s0, s1, mkdir)
}

func (d Directory) Chdir(path string) (types.Directory, error) {
	return d.chdir(path, false)
}

func (d Directory) MkdirAll(path string) (types.Directory, error) {
	return d.chdir(path, true)
}

func (fs *Filesystem) Chdir(path string) (types.Directory, error) {
	return fs.chdir(path, false)
}

func (fs *Filesystem) MkdirAll(path string) (types.Directory, error) {
	return fs.chdir(path, true)
}
