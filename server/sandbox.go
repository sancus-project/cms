package server

import (
	"go.sancus.dev/cms"
)

// A Sandbox is a chroot-ed Directory
type Sandbox struct {
	cms.Directory

	server *Server
}

func (s *Server) newSandbox(root *Directory) cms.Directory {
	return &Sandbox{
		Directory: &Directory{
			path:     "/",
			fullpath: root.fullpath,
			data:     root.data,
			cache:    root.cache,
		},
		server: s,
	}
}

// Spawn Sandbox from Server
func (s *Server) Chroot(path string) (cms.Directory, error) {
	dir, err := s.chdir(path, false)
	if err != nil {
		return nil, err
	}

	v := s.newSandbox(dir)
	return v, nil
}

// Spawn Sandbox from Sandbox
func (s *Sandbox) Chroot(path string) (cms.Directory, error) {
	dir, err := s.chdir(path, false)
	if err != nil {
		return nil, err
	}

	v := s.server.newSandbox(dir)
	return v, nil
}
