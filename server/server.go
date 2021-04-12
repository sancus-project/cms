package server

// CMS Server
type Server struct{}

func NewServer(options ...ServerOption) (*Server, error) {
	s := &Server{}

	for _ = range options {
	}

	return s, nil
}

// Options
type ServerOption interface{
	IsServerOption() ServerOption
}

type serverOption struct{}

func (s *serverOption) IsServerOption() ServerOption {
	return s
}
