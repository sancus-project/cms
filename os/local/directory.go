package local

type Directory struct {
	path string
}

func (d Directory) Path() string {
	return d.path
}
