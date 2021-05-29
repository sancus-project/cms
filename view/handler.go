package view

import (
	"net/http"

	"go.sancus.dev/cms"
	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

func (v *View) pageNotImplemented(err error) (web.Handler, bool) {
	if err == nil {
		// 503 Temporarily Unavailable
		h := &errors.HandlerError{
			Code: http.StatusServiceUnavailable,
		}
		return h, true

	} else if e, ok := err.(web.Error); ok {

		if e.Status() == http.StatusNotFound {
			// Not for us
			return nil, false
		}

	}

	h := errors.NewFromError(err).(web.Handler)
	return h, true
}

func (v *View) pageFilesDirectory(d cms.Directory, err error) (web.Handler, bool) {
	return v.pageNotImplemented(err)
}

func (v *View) pageServeResource(r cms.Resource, err error) (web.Handler, bool) {
	return v.pageNotImplemented(err)
}

func (v *View) pageEditResource(r cms.Resource, err error) (web.Handler, bool) {
	return v.pageNotImplemented(err)
}
