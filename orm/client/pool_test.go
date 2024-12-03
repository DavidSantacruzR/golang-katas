package client

import (
	"errors"
	"testing"
)

func TestPoolConnectionOnStart(t *testing.T) {
	connections := 1
	pool, connErr := StartConnectionPool(connections)
	if connErr != nil {
		t.Errorf("Unable to start connection pool: %v %d", connErr, connections)
	}
	if pool.liveConnections != 1 {
		t.Errorf("Pool initialised with the wrong connections number: %d", connections)
	}
}

func TestPoolConnectionMaxConnOnStart(t *testing.T) {
	connections := 11
	_, connErr := StartConnectionPool(connections)
	expectedError := &MaxConnectionLimitExceededError{connections: connections}
	if !errors.Is(expectedError, connErr) {
		t.Errorf("Expected error type: %v but got: %v", expectedError, connErr)
	}
}
