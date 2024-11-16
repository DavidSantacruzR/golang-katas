package main

import (
	"fmt"
	"golang/redis"
	"time"
)

func main() {
	client := redis.Client{}
	client.New("127.0.0.1", 6379)
	client.SetValue("juan", "moya", 4)
	client.SetValue("david", "santacruz", 2)
	fmt.Println("sleeping 3 seconds to have an expired key.")
	time.Sleep(3 * time.Second)
	var customer1 = client.GetKey("juan")
	var customer2 = client.GetKey("david")
	fmt.Println("Received from server:", customer1)
	fmt.Println("Received from server:", customer2)
}
