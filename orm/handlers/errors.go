package handlers

import (
	"errors"
	"fmt"
)

type WrongModelTypeError struct {
	message string
}

func (e *WrongModelTypeError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

func (e *WrongModelTypeError) Is(err error) bool {
	var other *WrongModelTypeError
	ok := errors.As(err, &other)
	if !ok {
		return false
	}
	return e.message == other.message
}

type MissingModelTypeError struct {
	message string
}

func (e *MissingModelTypeError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

func (e *MissingModelTypeError) Is(err error) bool {
	var other *MissingModelTypeError
	ok := errors.As(err, &other)
	if !ok {
		return false
	}
	return e.message == other.message
}
