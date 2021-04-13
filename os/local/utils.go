package local

import (
	"fmt"
	"path/filepath"
	"strings"
)

func AbsoluteRoot(path string) (string, error) {
	if len(path) == 0 {
		return "", nil
	} else if s, err := filepath.Abs(path); err != nil {
		return "", err
	} else {
		return s, nil
	}
}

func CleanRoot(path string) (root string, err error) {

	if s := strings.TrimPrefix(path, "local:"); s != path {
		// local: can be relative
		root = filepath.Clean(s)
		root, err = AbsoluteRoot(root)
	} else if s := strings.TrimPrefix(path, "file://"); s != path {
		// but file:// needs to be absolute
		root = filepath.Clean(s)
	} else {
		// and if no prefix, accept relative paths too
		root = filepath.Clean(s)
		root, err = AbsoluteRoot(root)
	}

	// the final root needs to be non-empty, clean and absolute
	if err != nil {
		// error caused by Abs() when failing to Getwd()
		return "", err
	} else if len(root) == 0 || root[0] != '/' {
		err = fmt.Errorf("%q: invalid root path", path)
		return "", err
	} else {
		return root, nil
	}
}
