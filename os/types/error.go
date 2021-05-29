package types

import (
	"fmt"
	"syscall"
)

const (
	EBUSY  = syscall.EBUSY
	EINVAL = syscall.EINVAL
)

type FileError struct {
	Name  string
	Errno syscall.Errno
}

func (err FileError) Error() string {
	return fmt.Sprintf("%s: %s", err.Name, err.Errno.Error())
}

func (err FileError) Unwrap() error {
	return err.Errno
}

// EINVAL
func ErrInvalid(pathname string) error {
	return &FileError{pathname, syscall.EINVAL}
}
