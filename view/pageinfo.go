package view

import (
	"log"
	"net/http"
	"strings"

	"go.sancus.dev/cms"
	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

func (v *View) pageInfo(r *http.Request) (web.Handler, bool) {
	path := v.config.GetRoutePath(r)
	log.Printf("%+n: %s", errors.Here(0), path)

	if p, ok := v.pageSitemap(path); ok {
		return p, true
	} else if p, ok := v.pagePing(path); ok {
		return p, true
	} else if p, ok := v.pageFiles(path); ok {
		return p, true
	} else if p, ok := v.pageEdit(path); ok {
		return p, true
	} else if p, ok := v.pageResource(path); ok {
		return p, true
	} else {
		return nil, false
	}
}

func (v *View) pageSitemap(path string) (web.Handler, bool) {
	if s0 := v.config.Sitemap; s0 != "" {
		if s1 := strings.TrimPrefix(path, s0); s1 != path {
			if s1 == "" || s1[0] == '/' {
				return v.config.SitemapHandler, true
			}
		}
	}

	return nil, false
}

func (v *View) pagePing(path string) (web.Handler, bool) {
	if s := v.config.Ping; s != "" && s == path {
		return v.config.PingHandler, true
	}
	return nil, false
}

func (v *View) pageFiles(path string) (web.Handler, bool) {
	if s0 := v.config.Files; s0 != "" {
		if s1 := strings.TrimSuffix(path, s0); s1 != path {
			if d := v.getDirectory(s1); d != nil {
				return v.pageFilesDirectory(d)
			}
		}
	}
	return nil, false
}

func (v *View) pageEdit(path string) (web.Handler, bool) {
	if s0 := v.config.Edit; s0 != "" {
		if s1 := strings.TrimSuffix(path, s0); s1 != path {
			if r := v.getResource(s1); r != nil {
				return v.pageEditResource(r)
			}
		}
	}
	return nil, false
}

func (v *View) pageResource(path string) (web.Handler, bool) {
	if r := v.getResource(path); r != nil {
		return v.pageServeResource(r)
	}
	return nil, false
}

func (v *View) getDirectory(path string) cms.Directory {
	return nil
}

func (v *View) getResource(path string) cms.Resource {
	return nil
}
