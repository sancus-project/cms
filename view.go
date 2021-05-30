package cms

import (
	"context"
	"net/http"

	"go.sancus.dev/web"
	"go.sancus.dev/web/errors"
)

type View interface {
	http.Handler
	web.Handler
	web.RouterPageInfo

	Middleware(prefix string) func(next http.Handler) http.Handler

	Config() *ViewConfig
}

type ViewConfig struct {
	GetRoutePath func(r *http.Request) string
	GetUser      func(ctx context.Context) User
	SetResource  func(ctx context.Context, res Resource) context.Context
	GetResource  func(ctx context.Context) Resource
	SetDirectory func(ctx context.Context, res Directory) context.Context
	GetDirectory func(ctx context.Context) Directory

	Edit           string // per resource
	EditHandler    web.HandlerFunc
	Files          string // per directory
	FilesHandler   web.HandlerFunc
	Ping           string // optional per view
	PingHandler    web.HandlerFunc
	Sitemap        string // optional per view
	SitemapHandler web.HandlerFunc

	ErrorHandler web.ErrorHandlerFunc

	Index    string // default page
	ReadOnly bool   // storage can't be modified through this View
}

func (c *ViewConfig) SetGetRoutePath(f func(*http.Request) string) error {
	if f == nil {
		f = DefaultGetRoutePath
	}
	c.GetRoutePath = f
	return nil
}

func (c *ViewConfig) SetGetUser(f func(ctx context.Context) User) error {
	if f == nil {
		f = DefaultGetUser
	}
	c.GetUser = f
	return nil
}

// Put Resource into context.Context
func (c *ViewConfig) SetSetResource(f func(ctx context.Context, res Resource) context.Context) error {
	if f == nil {
		f = DefaultSetResource
	}
	c.SetResource = f
	return nil
}

// Get Resource from context.Context
func (c *ViewConfig) SetGetResource(f func(ctx context.Context) Resource) error {
	if f == nil {
		f = DefaultGetResource
	}
	c.GetResource = f
	return nil
}

// Put Directory into context.Context
func (c *ViewConfig) SetSetDirectory(f func(ctx context.Context, res Directory) context.Context) error {
	if f == nil {
		f = DefaultSetDirectory
	}
	c.SetDirectory = f
	return nil
}

// Get Directory from context.Context
func (c *ViewConfig) SetGetDirectory(f func(ctx context.Context) Directory) error {
	if f == nil {
		f = DefaultGetDirectory
	}
	c.GetDirectory = f
	return nil
}

// Set View's File Editor
func (c *ViewConfig) SetEditHandler(path string, handler web.HandlerFunc) error {
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
func (c *ViewConfig) SetFilesHandler(path string, handler web.HandlerFunc) error {
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
func (c *ViewConfig) SetPingHandler(path string, handler web.HandlerFunc) error {
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
func (c *ViewConfig) SetSitemapHandler(path string, handler web.HandlerFunc) error {
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

func (c *ViewConfig) SetErrorHandler(f web.ErrorHandlerFunc) error {
	if f == nil {
		f = errors.HandleError
	}
	c.ErrorHandler = f
	return nil
}

func (c *ViewConfig) SetIndexPage(index string) error {
	c.Index = index
	return nil
}

// Defaults
func (c *ViewConfig) SetDefaults() error {
	c.SetGetRoutePath(c.GetRoutePath)
	c.SetGetUser(c.GetUser)
	c.SetSetResource(c.SetResource)
	c.SetGetResource(c.GetResource)
	c.SetSetDirectory(c.SetDirectory)
	c.SetGetDirectory(c.GetDirectory)

	c.SetEditHandler(c.Edit, c.EditHandler)
	c.SetFilesHandler(c.Files, c.FilesHandler)
	c.SetPingHandler(c.Ping, c.PingHandler)
	c.SetSitemapHandler(c.Sitemap, c.SitemapHandler)

	c.SetErrorHandler(c.ErrorHandler)
	return nil
}
