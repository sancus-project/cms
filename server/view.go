package server

import (
	"go.sancus.dev/cms"
	"go.sancus.dev/cms/view"
)

func (s *Server) View(path string) cms.View {
	v := &view.View{}
	return v
}
