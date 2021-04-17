package errors

import (
	"go.sancus.dev/cms"
)

type ResourceError struct {
	cms.HandlerError
	Resource cms.Resource
}
