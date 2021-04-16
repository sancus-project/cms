package server

import (
	"go.sancus.dev/cms"
	"go.sancus.dev/cms/os/utils"
)

type Sandbox struct {
	root   string
	server *Server
}

// Spawn Sandbox from Server
func (s *Server) Chroot(path string) (cms.Server, error) {
	path, err := utils.ValidPath(path)
	if err != nil {
		return nil, err
	}

	v := &Sandbox{
		root:   path,
		server: s,
	}

	return v, nil
}

// Spawn Sandbox from Sandbox
func (s *Sandbox) Chroot(path string) (cms.Server, error) {
	path, err := utils.ValidPath(path)
	if err != nil {
		return nil, err
	}

	v := &Sandbox{
		root:   path,
		server: s.server,
	}

	return v, nil
}
