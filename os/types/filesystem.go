package types

type Filesystem interface {
	Root() string
	Protocol() string

	Chdir(string) (Directory, error)
	MkdirAll(string) (Directory, error)

	Close() error
}

type Directory interface {
	Path() string

	Chdir(string) (Directory, error)
}
