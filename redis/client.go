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
	var conn, _ = net.Dial("tcp", c.address)
	defer c.closeClientConnection(conn)
	fmt.Println("Byte representation", []byte(key))
	_, err := conn.Write([]byte(key))
	if err != nil {
		fmt.Println("Could not get key", err)
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
	err := conn.Close()
	if err != nil {
		fmt.Printf("Unable to close the recently opened connection.")
	}
}
