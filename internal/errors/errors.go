package errors

import "errors"

var (
	ErrInvalidPassword        = errors.New("invalid password")
	ErrInvalidInputData       = errors.New("invalid input data")
	ErrInvalidReqBody         = errors.New("invalid request body")
	ErrInvalidParameterFormat = errors.New("invalid parameter format")
	ErrForbidden              = errors.New("forbidden")
	ErrMissingFromDB          = errors.New("missing from the database")
)
