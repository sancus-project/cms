package cms

type Directory interface {
	Path() string

	MkdirAll(string) (Directory, error)
}
