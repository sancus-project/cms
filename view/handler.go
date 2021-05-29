package view

import (
	"go.sancus.dev/cms"
	"go.sancus.dev/web"
)

func (v *View) pageFilesDirectory(d cms.Directory) (web.Handler, bool) {
	return nil, false
}

func (v *View) pageServeResource(r cms.Resource) (web.Handler, bool) {
	return nil, false
}

func (v *View) pageEditResource(r cms.Resource) (web.Handler, bool) {
	return nil, false
}
