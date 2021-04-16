package server

import (
	"go.sancus.dev/cms"
	"go.sancus.dev/cms/view"
)

// Spawn cms.View from Server
func (s *Server) View(path string) cms.View {
	return view.NewView(s, s.ViewConfig)
}
