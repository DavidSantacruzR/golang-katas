package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang/orm/models"
	"log"
	"reflect"
	"strings"
)

type PSQLRepository struct {
	pool *sql.DB
}

func (p *PSQLRepository) Get(table string, model models.Model) Recordset {
	fields := strings.Join(model.GetFields(), ", ")
	response, _ := p.pool.Query(fmt.Sprintf("SELECT %s FROM %s", fields, table))
	defer func(response *sql.Rows) {
		closingErr := response.Close()
		if closingErr != nil {
			fmt.Println(closingErr)
		}
	}(response)
	result := make([]interface{}, len(model.GetFields()))
	for response.Next() {
		for i := range result {
			result[i] = new(interface{})
		}
		scanErr := response.Scan(result...)
		for i := range result {
			v := reflect.Indirect(reflect.ValueOf(result[i]))
			value := v.Interface()
			if byteSlice, ok := value.([]byte); ok {
				uuidString := string(byteSlice) // Convert byte slice to string
				uuidValue, err := uuid.Parse(uuidString)
				if err != nil {
					log.Printf("Failed to parse UUID: %v", err)
				} else {
					fmt.Println(uuidValue.String())
				}
			} else if nullString, ok := value.(*sql.NullString); ok {
				fmt.Println(nullString.String)
			} else {
				fmt.Println(value)
			}
		}
		if scanErr != nil {
			log.Printf("Failed to scan row: %v", scanErr)
		}
		fmt.Println()
	}
	return Recordset{}
}

func (p *PSQLRepository) Create(data map[string]any) OperationResult {
	//TODO implement me
	panic("implement me")
}

func (p *PSQLRepository) Update(id int, data map[string]any) OperationResult {
	//TODO implement me
	panic("implement me")
}

func (p *PSQLRepository) Delete(id int) OperationResult {
	//TODO implement me
	panic("implement me")
}

func NewPsqlRepository(connUrl string) Base {
	conn, err := sql.Open("postgres", connUrl)
	if err != nil {
		fmt.Println("Unexpected error", err)
	}
	return &PSQLRepository{conn}
}
