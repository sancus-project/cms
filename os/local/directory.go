package local

import (
	"os"
	"path/filepath"

	"go.sancus.dev/cms/os/registry"
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

func (fs *Filesystem) MkdirAll(path string) (registry.Directory, error) {
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
