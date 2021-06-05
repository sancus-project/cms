package server

import (
	"go.sancus.dev/cms"
)

type Sandbox struct {
	root   cms.Directory
	server *Server
}

// Spawn Sandbox from Server
func (s *Server) Chroot(path string) (cms.Directory, error) {
	dir, err := s.Chdir(path)
	if err != nil {
		return nil, err
	}

	v := &Sandbox{
		root:   dir,
		server: s,
	}

	return v, nil
}

// Spawn Sandbox from Sandbox
func (s *Sandbox) Chroot(path string) (cms.Directory, error) {
	dir, err := s.Chdir(path)
	if err != nil {
		return nil, err
	}

	v := &Sandbox{
		root:   dir,
		server: s.server,
	}

	return v, nil
}
