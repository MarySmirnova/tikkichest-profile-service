package xerrors

import "errors"

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrForbidden       = errors.New("forbidden")
	ErrMissingFromDB   = errors.New("missing from the database")
)
