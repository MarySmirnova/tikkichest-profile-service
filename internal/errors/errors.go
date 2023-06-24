package errors

import "errors"

var (
	ErrInvalidPassword  = errors.New("invalid password")
	ErrInvalidInputData = errors.New("invalid input data")
	ErrInvalidReqBody   = errors.New("invalid request body")
	ErrForbidden        = errors.New("forbidden")
	ErrMissingFromDB    = errors.New("missing from the database")
)
