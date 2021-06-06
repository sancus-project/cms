package cms

import (
	"io/fs"
)

type Directory interface {
	fs.DirEntry
	fs.FileInfo

	Path() string

	Chdir(string) (Directory, error)
	MkdirAll(string) (Directory, error)

	Open(string) (Resource, error)
}

type Server interface {
	View(path string) View

	Chdir(string) (Directory, error)
	MkdirAll(string) (Directory, error)
	Chroot(string) (Directory, error)
}
