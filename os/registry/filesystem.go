package registry

import (
	"context"
	"fmt"
	"strings"
)

type Filesystem interface {
	Root() string
	Protocol() string

	Close() error
}

type FilesystemConstructor func(context.Context, string) (Filesystem, error)

var filesystems map[string]FilesystemConstructor

func RegisterFilesystem(prefix string, f FilesystemConstructor) error {
	if f == nil {
		// NOP
		err := fmt.Errorf("%s: no constructor for %q %s",
			"RegisterFilesystem", prefix, "given")
		return err
	}

	filesystems[prefix] = f
	return nil
}

func NewFilesystem(ctx context.Context, root string) (Filesystem, error) {
	for k, f := range filesystems {
		if len(k) > 0 {
			if strings.HasPrefix(root, k) {
				// hit
				return f(ctx, root)
			}
		}
	}

	if f, ok := filesystems[""]; ok {
		return f(ctx, root)
	}

	err := fmt.Errorf("%s: no constructor for %q %s",
		"NewFilesystem", root, "found")
	return nil, err
}

func init() {
	filesystems = make(map[string]FilesystemConstructor)
}
