package types

type Filesystem interface {
	Root() string
	Protocol() string

	MkdirAll(string) (Directory, error)

	Close() error
}

type Directory interface {
	Path() string
}
