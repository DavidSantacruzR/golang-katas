package main

import (
	"fmt"
	"golang/redis"
)

func main() {
	client := redis.Client{}
	client.New("127.0.0.1", 6379)
	client.SetValue("juan", "moya")
	client.SetValue("david", "santacruz")
	var customer1 = client.GetKey("juan")
	var customer2 = client.GetKey("david")
	fmt.Println("Received from server:", customer1)
	fmt.Println("Received from server:", customer2)
}
