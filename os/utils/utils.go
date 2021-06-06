package utils

import (
	"path/filepath"

	"go.sancus.dev/cms/os/types"
)

func Join(s ...string) (string, error) {
	path := filepath.Join(s...)

	if len(path) > 0 && path[0] == '/' {
		s := filepath.Clean(path)
		if len(s) > 0 && s[0] == '/' {
			return s, nil
		}
	}

	return "", types.ErrInvalid(path)
}
