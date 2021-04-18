package os

import (
	"path/filepath"
)

func Join(d ...string) string {
	return filepath.Join(d...)
}
