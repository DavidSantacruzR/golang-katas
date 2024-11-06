package redis

import (
	"fmt"
	"net"
	"strconv"
)

/*
* Create a new listener bind on port 3000
* Use a handle connection function as communication between the server, and the redis client.
* See where to initialise the in memory db client.
* Handle data sent, received using the protocols.
 */

type Ctx struct {
	key  string
	text string
	data string
}

type Server struct {
	server net.Listener
	conn   net.Conn
}

func (srv *Server) New(addr string, port int) *Server {
	listener, serverErr := net.Listen("tcp", addr+":"+strconv.Itoa(port))
	if serverErr != nil {
		fmt.Println("Could not start redis server")
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			srv.conn = connection
			fmt.Println("lola")
			go srv.handleConnection()
			fmt.Println("lala")
		}
	}
}

func (srv *Server) handleConnection() {
	fmt.Println("Connection established")
}

//func (srv *Server) readFromBuffer(text string) {
//	fmt.Println("readFromBuffer gets executed")
//	for {
//		buffer := make([]byte, 1024)
//		bf, err := srv.conn.Read(buffer)
//		if err != nil {
//			fmt.Println("Data", bf)
//		}
//	}
//}
//func (srv *Server) writeToBuffer(conn net.Conn) {
//	_, err := srv.conn.Write([]byte("Alive"))
//	if err != nil {
//		fmt.Println("Error writing to connection.")
//	}
//}
