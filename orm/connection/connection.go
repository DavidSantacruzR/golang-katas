package connection

type Connection struct{}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) ExecuteQuery(input string) string {
	return ""
}
