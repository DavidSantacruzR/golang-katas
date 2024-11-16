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
	return c.executeCommand(command)
}

func (c *Client) SetValue(key string, value string) string {
	command := fmt.Sprintf("*2\r\n$3\r\nSET\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", len(key), key, len(value), value)
	return c.executeCommand(command)
}

func (c *Client) executeCommand(command string) string {
	var conn, _ = net.Dial("tcp", c.address)
	defer c.closeClientConnection(conn)
	_, err := conn.Write([]byte(command))
	if err != nil {
		fmt.Println("Unable to execute instruction", command, err)
	}
	buffer := make([]byte, 1024)
	response, valueErr := conn.Read(buffer)
	if valueErr != nil {
		fmt.Println("Unexpected error reading buffer", err)
		return ""
	}
	return string(buffer[:response])
}

func (c *Client) closeClientConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Printf("Unable to close the recently opened connection.")
	}
}
