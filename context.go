package cms

import (
	"context"
	"net/http"

	sancus "go.sancus.dev/web/context"
)

func DefaultGetUser(ctx context.Context) User {
	// No user management
	return nil
}

func DefaultGetRoutePath(r *http.Request) string {
	if rctx := sancus.RouteContext(r.Context()); rctx != nil {
		return rctx.RoutePath
	}
	return r.URL.Path
}

func DefaultSetResource(ctx context.Context, res Resource) context.Context {
	return context.WithValue(ctx, ResourceCtxKey, res)
}

func DefaultGetResource(ctx context.Context) Resource {
	if res, ok := ctx.Value(ResourceCtxKey).(Resource); ok {
		return res
	}
	return nil
}

func DefaultSetDirectory(ctx context.Context, dir Directory) context.Context {
	return context.WithValue(ctx, DirectoryCtxKey, dir)
}

func DefaultGetDirectory(ctx context.Context) Directory {
	if dir, ok := ctx.Value(DirectoryCtxKey).(Directory); ok {
		return dir
	}
	return nil
}

var (
	ResourceCtxKey  = &contextKey{"Resource"}
	DirectoryCtxKey = &contextKey{"Directory"}
)

// contextKey is a value for use with context.WithValue
type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "cms context value " + k.name
}
