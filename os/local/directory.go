package local

import (
	"os"
	"path/filepath"

	"go.sancus.dev/cms/os/types"
	"go.sancus.dev/cms/os/utils"
)

const (
	NewDirectoryMode = os.FileMode(0755)
)

type Directory struct {
	path   string
	actual string
}

func (d Directory) Path() string {
	return d.path
}

func (d Directory) Chdir(path string) (types.Directory, error) {
	// Allow Relative paths
	s0, err := utils.ValidPath(filepath.Join(d.path, path))
	if err != nil {
		return nil, err
	}

	next := &Directory{
		path:   s0,
		actual: filepath.Join(d.actual, path),
	}

	if fi, err := os.Stat(next.actual); err != nil {
		return nil, err
	} else if fi.IsDir() {
		return d, nil
	} else {
		return nil, os.ErrNotExist
	}
}

func (fs *Filesystem) MkdirAll(path string) (types.Directory, error) {
	var err error
	var mode = NewDirectoryMode

	// validate path
	if path, err = CleanRoot(path); err != nil {
		return nil, err
	}

	fs.mu.Lock()
	defer fs.mu.Unlock()

	d := &Directory{
		path:   path,
		actual: filepath.Join(fs.root, path),
	}

	if err = os.MkdirAll(d.actual, mode); err != nil {
		return nil, err
	}

	return d, nil
}

func (fs *Filesystem) Chdir(path string) (types.Directory, error) {
	var err error

	// validate path
	if path, err = CleanRoot(path); err != nil {
		return nil, err
	}

	fs.mu.Lock()
	defer fs.mu.Unlock()

	d := &Directory{
		path:   path,
		actual: filepath.Join(fs.root, path),
	}

	if fi, err := os.Stat(d.actual); err != nil {
		return nil, err
	} else if fi.IsDir() {
		return d, nil
	} else {
		return nil, os.ErrNotExist
	}
}
