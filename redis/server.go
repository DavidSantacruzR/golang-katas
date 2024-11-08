package redis

import (
	"fmt"
	"net"
	"strconv"
)

/*
TODO: Implement the redis protocol.
TODO: Handle memory storing, and retrieval of values (GET, PUT).
TODO: Error handling with exceptions and panics!
TODO: Manage byte length on messages to either write or read.
TODO: Distributed memory instances.
*/

type Ctx struct {
	key  string
	text string
	data string
}

type Server struct {
	server net.Listener
}

func (srv *Server) New(addr string, port int) *Server {
	fmt.Println("starting a new server")
	listener, serverErr := net.Listen("tcp", addr+":"+strconv.Itoa(port))
	if serverErr != nil {
		fmt.Println("Could not start redis server")
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("There was an error accepting connection")
		}
		go srv.handleConnection(connection)
	}
}

func (srv *Server) handleConnection(conn net.Conn) {
	defer srv.closeConnection(conn)
	fmt.Println("Handling connection data.")
	buffer := make([]byte, 1024)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			return
		}
		fmt.Printf("Received data: %s\n", string(buffer))
		conn.Write([]byte("moya"))
	}
}

func (srv *Server) closeConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Printf("Unable to close the current connection.")
		return
	}
}
