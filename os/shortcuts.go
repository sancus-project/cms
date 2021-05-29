package os

import (
	"os"
	"path/filepath"
)

var (
	ErrNotExist = os.ErrNotExist
)

func Join(d ...string) string {
	return filepath.Join(d...)
}
