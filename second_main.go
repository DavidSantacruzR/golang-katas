package main

import (
	"fmt"
	"golang/redis"
)

func main() {
	client := redis.Client{}
	client.New("127.0.0.1", 3000)
	var value = client.GetKey("juan")
	fmt.Println("Received from server", value)
}
