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

func (c *ViewConfig) SetDefaults() error {
	// file editor
	if len(c.Edit) == 0 {
		c.EditHandler = nil
	} else if c.EditHandler == nil {
		c.Edit = ""
	} else if c.Edit[0] != '/' {
		c.Edit = "/" + c.Edit
	}

	// file manager
	if len(c.Files) == 0 {
		c.FilesHandler = nil
	} else if c.FilesHandler == nil {
		c.Files = ""
	} else if c.Files[0] != '/' {
		c.Files = "/" + c.Files
	}

	// heartbeat
	if len(c.Ping) == 0 {
		c.PingHandler = nil
	} else if c.PingHandler == nil {
		c.Ping = ""
	} else if c.Ping[0] != '/' {
		c.Ping = "/" + c.Ping
	}

	// sitemap
	if len(c.Sitemap) == 0 {
		c.SitemapHandler = nil
	} else if c.SitemapHandler == nil {
		c.Sitemap = ""
	} else if c.Sitemap[0] != '/' {
		c.Sitemap = "/" + c.Sitemap
	}

	return nil
}
