package server

import (
	"log"

	"go.sancus.dev/cms"
	"go.sancus.dev/cms/view"
)

// Spawn cms.View from Server
func (s *Server) View(path string) cms.View {
	dir, err := s.Chroot(path)
	if err != nil {
		log.Fatal(err)
	}

	v, err := view.NewView(dir, s.ViewConfig)
	if err != nil {
		log.Fatal(err)
	}

	return v
}

// Spawn cms.View from Sandbox
func (s *Sandbox) View(path string) cms.View {
	dir, err := s.Chroot(path)
	if err != nil {
		log.Fatal(err)
	}

	v, err := view.NewView(dir, s.server.ViewConfig)
	if err != nil {
		log.Fatal(err)
	}

	return v
}
