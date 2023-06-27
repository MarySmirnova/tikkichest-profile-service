package errors

import (
	"bytes"
	"fmt"
	"sync"
)

type validateError interface {
	Add(key, value string)
	Check() error
	Unwrap() error
	error
}

var ValidateError = NewValidError()

type validError struct {
	errs     map[string]string
	mu       sync.RWMutex
	makeOnce sync.Once
}

func NewValidError() validateError {
	return &validError{
		mu: sync.RWMutex{},
	}
}

func (e *validError) Add(key, value string) {
	e.makeOnce.Do(func() {
		if e.errs == nil {
			e.errs = make(map[string]string)
		}
	})

	e.mu.Lock()
	defer e.mu.Unlock()
	e.errs[key] = value
}

func (e *validError) Check() error {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if e.errs == nil || len(e.errs) == 0 {
		return nil
	}

	return validateError(e)
}

func (e *validError) Error() string {
	err := e.Check()
	if err == nil {
		return ""
	}

	e.mu.RLock()
	defer e.mu.RUnlock()

	var buf bytes.Buffer
	buf.WriteString("validation failed:\n")

	for f, v := range e.errs {
		buf.WriteString(fmt.Sprintf("%s: %s\n", f, v))
	}
	return buf.String()
}

func (e *validError) Unwrap() error {
	err := e.Check()
	if err != nil {
		return ValidateError
	}
	return nil
}
