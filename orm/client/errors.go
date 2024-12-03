package client

import (
	"errors"
	"fmt"
)

type MaxConnectionLimitExceededError struct {
	connections int
}

func (e *MaxConnectionLimitExceededError) Error() string {
	return fmt.Sprintf("max connection limit exceeded: with %d", e.connections)
}

func (e *MaxConnectionLimitExceededError) Is(err error) bool {
	var other *MaxConnectionLimitExceededError
	ok := errors.As(err, &other)
	if !ok {
		return false
	}
	return e.connections == other.connections
}
