package redis

import (
	"fmt"
	"net"
	"strconv"
)

type Client struct {
	server  Server
	address string
}

func (c *Client) New(addr string, port int) {
	c.address = addr + ":" + strconv.Itoa(port)
}

func (c *Client) GetKey(key string) string {
	command := fmt.Sprintf("*2\r\n$3\r\nGET\r\n$%d\r\n%s\r\n", len(key), key)
	var conn, _ = net.Dial("tcp", c.address)
	defer c.closeClientConnection(conn)
	_, err := conn.Write([]byte(command))
	if err != nil {
		fmt.Println("Unable to execute instruction", command, err)
	}
	buffer := make([]byte, 1024)
	value, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Could not get key", err)
		return ""
	}
	return string(buffer[:value])
}

func (c *Client) closeClientConnection(conn net.Conn) {
	/*TODO: Implement the redis parser set command.*/
	err := conn.Close()
	if err != nil {
		fmt.Printf("Unable to close the recently opened connection.")
	}
}
