package main

import (
	"fmt"
	"golang/orm/client"
	"golang/orm/handlers"
)

type User struct {
}

func main() {
	fields := make(map[string]any)
	fields["Id"] = 1
	dbClient := client.NewDbClient(1)
	results, _ := handlers.Filter(&dbClient, User{}, fields)
	fmt.Println(results)
}
