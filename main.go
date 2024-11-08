package main

import (
	"golang/redis"
)

func main() {
	server := redis.Server{}
	server.New("127.0.0.1", 3000)
	//client := redis.Client{}
	//client.New("127.0.0.1", 3000)
	//client.ReadValue("Juan es un aspero")
}
