package errors

import "fmt"

// wrappedError is a struct that wraps an error and a message
type wrappedError struct {
	err error
	msg string
}

// Error returns the error message formatted as a string
func (e *wrappedError) Error() string {
	if e.err == nil && e.msg == "" {
		return ""
	} else if e.err == nil && e.msg != "" {
		return e.msg
	} else if e.err != nil && e.msg == "" {
		return e.err.Error()
	} else {
		return fmt.Sprintf("%s : %v", e.msg, e.err)
	}
}

// New returns a new error with the given message
func New(msg string) error {
	return &wrappedError{msg: msg}
}

// Wrap returns a new error that wraps the given error with the given message
func Wrap(err error, msg string) error {
	return &wrappedError{
		err: err,
		msg: msg,
	}
}
