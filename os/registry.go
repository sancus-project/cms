package os

import (
	"context"

	_ "go.sancus.dev/cms/os/local"
	"go.sancus.dev/cms/os/registry"
)

type (
	Directory  = registry.Directory
	Filesystem = registry.Filesystem
)

func NewFilesystem(ctx context.Context, path string) (Filesystem, error) {
	return registry.NewFilesystem(ctx, path)
}
