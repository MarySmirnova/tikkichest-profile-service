package xerrors

import "errors"

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidReqBody  = errors.New("invalid request body")
	ErrForbidden       = errors.New("forbidden")
	ErrMissingFromDB   = errors.New("missing from the database")
)
