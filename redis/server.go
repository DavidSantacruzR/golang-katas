package redis

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

/*
TODO: Implement the redis protocol.
TODO: Error handling with exceptions and panics!
TODO: Distributed memory instances.
*/

type Server struct {
	server net.Listener
	db     map[string]string
}

func (srv *Server) New(addr string, port int) *Server {
	fmt.Println("starting a new server")
	listener, serverErr := net.Listen("tcp", addr+":"+strconv.Itoa(port))
	if serverErr != nil {
		fmt.Println("Could not start redis server")
	}
	srv.db = make(map[string]string)
	fmt.Println("In memory db initialised.")
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
	buffer := make([]byte, 1024)
	for {
		_, err := conn.Read(buffer)
		if err != nil {
			return
		}
		instructions := strings.Split(string(buffer), "\r\n")
		if instructions[2] == "GET" {
			response := srv.fetchKeyFromDb(instructions[len(instructions)-2])
			_, _ = conn.Write([]byte(response))
			return
		}
		if instructions[2] == "SET" {
			/*TODO: Implement the complete SET method.*/
			response := srv.pushKeyToDb("test", "test")
			_, _ = conn.Write([]byte(response))
			return
		}
	}
}

func (srv *Server) closeConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Printf("Unable to close the current connection.")
		return
	}
}

func (srv *Server) fetchKeyFromDb(key string) string {
	fmt.Println(key)
	value, exists := srv.db[key]
	if exists {
		return value
	} else {
		return fmt.Sprintf("-WRONGKEY No value for key %s", key)
	}
}

func (srv *Server) pushKeyToDb(key string, value string) string {
	fmt.Println("Storing key, value pair.", key, value)
	return "done."
}
