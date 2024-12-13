package connection

import "fmt"

type Connection struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func (c *Connection) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.DbName)
}
