package cms

import (
	"net/http"
)

// GET
type Resource interface {
	Methods() []string
	MimeTypes() []string
	Render(w http.ResponseWriter, r *http.Request) error
}

// HEAD
type PeekableResource interface {
	Head(w http.ResponseWriter, r *http.Request) error
}

// PUT
type CreatableResource interface {
	Resource
	Put(r *http.Request) error
}

// DELETE
type DeletableResource interface {
	Resource
	Delete() error
}

// POST
type SubmitableResource interface {
	Resource
	Post(r *http.Request) error
}
