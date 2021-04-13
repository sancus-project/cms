package os

import (
	"context"

	_ "go.sancus.dev/cms/os/posix"
	"go.sancus.dev/cms/os/registry"
)

type Filesystem = registry.Filesystem

func NewFilesystem(ctx context.Context, path string) (Filesystem, error) {
	return registry.NewFilesystem(ctx, path)
}
