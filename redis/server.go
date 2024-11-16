package redis

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	value string
	TTL   int64
}

func NewRecord(value string, TTL int64) Record {
	return Record{
		value: value,
		TTL:   TTL,
	}
}

type Server struct {
	server net.Listener
	db     map[string]Record
}

func (srv *Server) New(addr string, port int) *Server {
	fmt.Println("starting a new server")
	listener, serverErr := net.Listen("tcp", addr+":"+strconv.Itoa(port))
	if serverErr != nil {
		fmt.Println("Could not start redis server")
	}
	srv.db = make(map[string]Record)
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
			timeToLive, parseErr := strconv.ParseInt(instructions[10], 10, 64)
			if parseErr != nil {
				timeToLive = 0
			}
			response := srv.pushKeyToDb(instructions[4], instructions[6], timeToLive)
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
	record, exists := srv.db[key]
	if exists {
		expiredKey := srv.validateExpiration(record)
		if expiredKey {
			srv.deleteKeyFromDb(key)
			return fmt.Sprintf("-EXPIREDKEY Key already expired %s", key)
		}
		return record.value
	} else {
		return fmt.Sprintf("-WRONGKEY No value for key %s", key)
	}
}

func (srv *Server) pushKeyToDb(key string, value string, ttl int64) string {
	srv.db[key] = NewRecord(value, time.Now().Unix()+ttl)
	return "+OK\r\n"
}

func (srv *Server) validateExpiration(record Record) bool {
	if record.TTL == 0 {
		return false
	}
	if time.Now().Unix() > record.TTL {
		return true
	}
	return false
}

func (srv *Server) deleteKeyFromDb(key string) {
	delete(srv.db, key)
}
