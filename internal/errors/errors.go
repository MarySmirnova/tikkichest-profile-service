package errors

import "errors"

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrForbidden       = errors.New("forbidden")
)
