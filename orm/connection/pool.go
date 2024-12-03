package connection

type Pool struct {
	connections     chan *Connection
	liveConnections int
}

func StartConnectionPool(connections int) (*Pool, error) {
	connectionsLimit := 10
	if connections > connectionsLimit {
		return nil, &MaxConnectionLimitExceededError{connections: connections}
	}
	channel := make(chan *Connection, connections)
	pool := &Pool{connections: channel, liveConnections: 0}
	for i := 0; i < connections; i++ {
		pool.connections <- NewConnection()
		pool.liveConnections++
	}
	return pool, nil
}

func (cp *Pool) Fetch() *Connection {
	var conn *Connection
	select {
	case conn = <-cp.connections:
		return conn
	default:
		conn = NewConnection()
		cp.connections <- conn
		return conn
	}
}

func (cp *Pool) Release(conn *Connection) {
	cp.connections <- conn
}
