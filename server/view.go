package server

import (
	"log"

	"go.sancus.dev/cms"
	"go.sancus.dev/cms/view"
)

// Spawn cms.View from Server
func (s *Server) View(path string) cms.View {
	v, err := s.Chroot(path)
	if err != nil {
		log.Fatal(err)
	}

	return view.NewView(v, s.ViewConfig)
}

// Spawn cms.View from Sandbox
func (s *Sandbox) View(path string) cms.View {
	v, err := s.Chroot(path)
	if err != nil {
		log.Fatal(err)
	}

	return view.NewView(v, s.server.ViewConfig)
}
