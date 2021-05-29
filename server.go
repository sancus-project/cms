package cms

type Directory interface {
	Path() string

	Chdir(string) (Directory, error)
	MkdirAll(string) (Directory, error)

	Open(string) (Resource, error)
}
