package domain

import "errors"

var (
	ErrNotFound      = errors.New("Person not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrInternalError = errors.New("internal server error")
)

