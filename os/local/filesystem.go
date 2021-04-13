package local

import (
	"context"
	"os"

	"go.sancus.dev/cms/os/registry"
)

type Filesystem struct {
	root string
	ctx  context.Context
}

func NewFilesystem(ctx context.Context, root string) (registry.Filesystem, error) {

	// validate root
	if _, err := os.ReadDir(root); err != nil {
		return nil, err
	}

	// Context
	if ctx == nil {
		ctx = context.Background()
	}

	v := &Filesystem{
		root: root,
		ctx:  ctx,
	}

	return v, nil
}

func (v *Filesystem) Close() error {
	return nil
}

func init() {
	registry.RegisterFilesystem("", NewFilesystem)
}
