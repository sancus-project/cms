package utils

import (
	"path/filepath"

	"go.sancus.dev/cms/os/types"
)

func ValidPath(path string) (string, error) {

	var s string

	if path == "." {
		s = path
	} else if len(path) > 0 {

		s = filepath.Clean(path)
		if s == "." {
			s = ""
		}
	}

	if len(s) > 0 {
		return s, nil
	}

	return "", types.ErrInvalid(path)
}
