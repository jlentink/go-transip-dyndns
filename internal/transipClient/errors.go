package transipClient

import (
	"errors"
)

var (
	NotFoundError   = NewNotFoundError(nil)
	NotChangedError = errors.New("not changed")
)

func NewNotFoundError(err error) error {
	return &errorNotFound{Err: err, S: "record not found"}
}

// errorString is a trivial implementation of error.
type errorNotFound struct {
	S   string
	Err error
}

func (e *errorNotFound) Error() string {
	return e.S
}
