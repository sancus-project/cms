package cms

import (
	"context"
	"net/http"
)

type RequestHandler func(http.ResponseWriter, *http.Request) error

type View interface {
	http.Handler

	Middleware(next http.Handler) http.Handler
	Handler(w http.ResponseWriter, r *http.Request) error

	Config() *ViewConfig
}

type ViewConfig struct {
	GetUser      func(ctx context.Context) User
	GetRoutePath func(ctx context.Context) string
	SetResource  func(ctx context.Context, res Resource) (context.Context, error)
	GetResource  func(ctx context.Context) Resource

	Edit           string // per resource
	EditHandler    http.Handler
	Files          string // per directory
	FilesHandler   http.Handler
	Ping           string // optional per view
	PingHandler    http.Handler
	Sitemap        string // optional per view
	SitemapHandler http.Handler

	ResourceErrorHandler ResourceErrorHandler
	ErrorHandler         ErrorHandler
	PanicHandler         PanicHandler

	Index    string // default page
	ReadOnly bool   // storage can't be modified through this View
}

// Set View's File Editor
func (c *ViewConfig) SetEditHandler(path string, handler http.Handler) error {
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
func (c *ViewConfig) SetFilesHandler(path string, handler http.Handler) error {
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
func (c *ViewConfig) SetPingHandler(path string, handler http.Handler) error {
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
func (c *ViewConfig) SetSitemapHandler(path string, handler http.Handler) error {
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
