package errors

import "fmt"

type ErrHTTP struct {
	Err  error
	Code int
}

func NewErrHTTP(err error, code int) error {
	return &ErrHTTP{
		Err:  err,
		Code: code,
	}
}

func (e *ErrHTTP) Error() string {
	return fmt.Sprintf("http code %d: %v", e.Code, e.Err)
}

func (e *ErrHTTP) Unwrap() error {
	return e.Err
}
