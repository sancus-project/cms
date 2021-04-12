package os

import (
	"context"

	"go.sancus.dev/cms/os/posix"
)

type Filesystem interface {
	Close() error
}

func NewFilesystem(ctx context.Context, path string) (Filesystem, error) {
	return posix.NewFilesystem(ctx, path)
}
