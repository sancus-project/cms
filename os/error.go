package os

import (
	"fmt"
	"syscall"
)

var (
	ErrBusy = &SysError{syscall.EBUSY}
)

type SysError struct {
	Errno syscall.Errno
}

func (err SysError) Error() string {
	return fmt.Sprintf("Errno %v", err.Errno)
}
