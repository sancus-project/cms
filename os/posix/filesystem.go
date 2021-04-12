package posix

import (
	"context"
)

type Filesystem struct {
	root string
	ctx  context.Context
}

func NewFilesystem(ctx context.Context, root string) (*Filesystem, error) {
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
