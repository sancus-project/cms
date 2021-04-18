package cms

import (
	"context"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) error

type View interface {
	http.Handler

	Middleware(next http.Handler) http.Handler
	Handler(w http.ResponseWriter, r *http.Request) error

	Config() *ViewConfig
}

type ViewConfig struct {
	GetRoutePath func(r *http.Request) string
	GetUser      func(ctx context.Context) User
	SetResource  func(ctx context.Context, res Resource) (context.Context, error)
	GetResource  func(ctx context.Context) Resource

	Edit           string // per resource
	EditHandler    Handler
	Files          string // per directory
	FilesHandler   Handler
	Ping           string // optional per view
	PingHandler    Handler
	Sitemap        string // optional per view
	SitemapHandler Handler

	ErrorHandler ErrorHandler
	PanicHandler PanicHandler

	Index    string // default page
	ReadOnly bool   // storage can't be modified through this View
}

// Set View's File Editor
func (c *ViewConfig) SetEditHandler(path string, handler Handler) error {
	if path == "" || handler == nil {
		path = ""
		handler = nil
	} else if path[0] != '/' {
		path = "/" + path
	}

	c.Edit = path
	c.EditHandler = handler
	return nil
}

// Set View's File Manager
func (c *ViewConfig) SetFilesHandler(path string, handler Handler) error {
	if path == "" || handler == nil {
		path = ""
		handler = nil
	} else if path[0] != '/' {
		path = "/" + path
	}

	c.Files = path
	c.FilesHandler = handler
	return nil
}

// Set View's Ping Handler
func (c *ViewConfig) SetPingHandler(path string, handler Handler) error {
	if path == "" || handler == nil {
		path = ""
		handler = nil
	} else if path[0] != '/' {
		path = "/" + path
	}

	c.Ping = path
	c.PingHandler = handler
	return nil
}

// Set View's Sitemap Handler
func (c *ViewConfig) SetSitemapHandler(path string, handler Handler) error {
	if path == "" || handler == nil {
		path = ""
		handler = nil
	} else if path[0] != '/' {
		path = "/" + path
	}

	c.Sitemap = path
	c.SitemapHandler = handler
	return nil
}
