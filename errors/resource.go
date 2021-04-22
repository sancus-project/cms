package errors

import (
	"go.sancus.dev/cms"
	"go.sancus.dev/web/errors"
)

type ResourceError struct {
	errors.HandlerError
	Resource cms.Resource
}
