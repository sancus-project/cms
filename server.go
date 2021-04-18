package cms

type Directory interface {
	Path() string

	MkdirAll(string) (Directory, error)
}

type Server interface {
	MkdirAll(string) (Directory, error)
}
