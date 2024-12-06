package main

import (
	"fmt"
	"golang/orm/client"
	"golang/orm/handlers"
	"golang/orm/models"
)

type User struct {
	Model   models.Model
	Id      int
	Name    string
	Surname string
	Address string
	Mobile  string
}

func main() {
	fields := make(map[string]any)
	fields["Age"] = 30
	dbClient := client.NewDbClient(1)
	results, queryError := handlers.Filter(&dbClient, User{}, fields)
	fmt.Println(results)
	fmt.Println(queryError)
}
