package types

import (
	"io/fs"
)

type Filesystem interface {
	Root() string
	Protocol() string

	Chdir(string) (Directory, error)
	MkdirAll(string) (Directory, error)

	Close() error
}

type Directory interface {
	fs.FileInfo
	fs.DirEntry

	Path() string

	Chdir(string) (Directory, error)
}
