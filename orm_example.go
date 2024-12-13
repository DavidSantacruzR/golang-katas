package main

import (
	"database/sql"
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
	/*db, _ := sql.Open("postgres", "postgresql://watchdog:watchdog@localhost/watchdog?sslmode=disable")
	response, _ := db.Query("select id, plate from transit_tickets")
	var id sql.NullString
	var plate sql.NullString
	defer response.Close()
	for response.Next() {
		err := response.Scan(&id, &plate)
		fmt.Println(id.String, plate.String)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
		}
		fmt.Println()
	}*/
	db := client.NewDbClient(connection.Connection{
		Host:     "localhost",
		Port:     5432,
		User:     "watchdog",
		Password: "watchdog",
		DbName:   "watchdog",
	})
	_, _ = handlers.Filter(&db, "transit_tickets", &Tickets{})
}
