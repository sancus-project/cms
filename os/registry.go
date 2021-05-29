package os

import (
	"context"

	_ "go.sancus.dev/cms/os/local"

	"go.sancus.dev/cms/os/registry"
	"go.sancus.dev/cms/os/types"
)

func NewFilesystem(ctx context.Context, path string) (types.Filesystem, error) {
	return registry.NewFilesystem(ctx, path)
}
