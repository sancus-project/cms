package cms

import (
	"net/http"
)

// GET
type Resource interface {
	MimeTypes() []string
	Get(http.ResponseWriter, *http.Request) error
}

// HEAD
type PeekableResource interface {
	Resource
	Head(w http.ResponseWriter, r *http.Request) error
}

// PUT
type CreatableResource interface {
	Resource
	Put(http.ResponseWriter, *http.Request) error
}

// DELETE
type DeletableResource interface {
	Resource
	Delete(http.ResponseWriter, *http.Request) error
}

// POST
type SubmitableResource interface {
	Resource
	Post(http.ResponseWriter, *http.Request) error
}
