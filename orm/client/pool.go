package client

type ConnectionPool struct {
	connections     chan *Connection
	liveConnections int
}

func StartConnectionPool(connections int) (*ConnectionPool, error) {
	connectionsLimit := 10
	if connections > connectionsLimit {
		return nil, &MaxConnectionLimitExceededError{connections: connections}
	}
	channel := make(chan *Connection, connections)
	pool := &ConnectionPool{connections: channel, liveConnections: 0}
	for i := 0; i < connections; i++ {
		pool.connections <- &Connection{}
		pool.liveConnections++
	}
	return pool, nil
}

func (cp *ConnectionPool) Fetch() *Connection {
	var conn *Connection
	select {
	case conn = <-cp.connections:
		return conn
	default:
		conn = &Connection{}
		cp.connections <- conn
		return conn
	}
}

func (cp *ConnectionPool) Release(conn *Connection) {
	cp.connections <- conn
}
