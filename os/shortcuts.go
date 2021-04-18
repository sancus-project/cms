package os

import (
	"path"
)

func Join(d ...string) string {
	return path.Join(d...)
}
