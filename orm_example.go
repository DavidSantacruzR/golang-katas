package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang/orm/client"
	"golang/orm/connection"
	"golang/orm/handlers"
	"reflect"
	"strings"
)

type Tickets struct {
	Id    sql.NullString
	Plate sql.NullString
	Year  sql.NullString
	Value sql.NullFloat64
	City  sql.NullString
}

func (t *Tickets) GetFields() []string {
	modelFields := reflect.ValueOf(t).Elem()
	fields := make([]string, modelFields.NumField())
	for i := 0; i < modelFields.NumField(); i++ {
		fields[i] = strings.ToLower(modelFields.Type().Field(i).Name)
	}
	return fields
}

func main() {
	db := client.NewDbClient(connection.Connection{
		Host:     "localhost",
		Port:     5432,
		User:     "watchdog",
		Password: "watchdog",
		DbName:   "watchdog",
	})
	data, _ := handlers.Filter(&db, "transit_tickets", &Tickets{})
	fmt.Println("Record:", data)
}
