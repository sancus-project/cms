package server

// Options
type ServerOption interface {
	IsServerOption() ServerOption
}

type serverOption struct{}

func (s *serverOption) IsServerOption() ServerOption {
	return s
}
