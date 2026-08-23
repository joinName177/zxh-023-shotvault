package vault

import "errors"

var (
	ErrWorkspaceNotFound = errors.New("workspace not found")
	ErrLayerNotFound     = errors.New("layer not found")
	ErrCancelled         = errors.New("operation cancelled")
	ErrInvalidLayer      = errors.New("invalid layer")
)
