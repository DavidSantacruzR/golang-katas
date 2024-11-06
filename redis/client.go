package redis

import (
	"bufio"
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

func (c *Client) ReadValue(key string) {
	fmt.Println("ReadValue, connecting")
	conn, connectionErr := net.Dial("tcp", c.address)
	if connectionErr != nil {
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(status, key)
		}
	}
}

func (c *Client) WriteValue(key string, value string) {
	fmt.Println("writeValue gets executed")
}
