package local

import (
	"os"
	"path"

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

func (fs *Filesystem) MkdirAll(pathname string) (registry.Directory, error) {
	var err error
	var mode = NewDirectoryMode

	// validate path
	if pathname, err = CleanRoot(pathname); err != nil {
		return nil, err
	}

	fs.mu.Lock()
	defer fs.mu.Unlock()

	d := &Directory{
		path:   pathname,
		actual: path.Join(fs.root, pathname),
	}

	if err = os.MkdirAll(d.actual, mode); err != nil {
		return nil, err
	}

	return d, nil
}
