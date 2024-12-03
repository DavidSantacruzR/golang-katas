package repository

import (
	"fmt"
	"golang/orm/connection"
)

type PSQLRepository struct {
	pool *connection.Pool
}

func (p *PSQLRepository) Get(model any, fields map[string]any) Recordset {
	conn := p.pool.Fetch()
	conn.ExecuteQuery(fmt.Sprintf("SELECT %V FROM %V", fields, model))
	panic("finish implementing me")
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

func NewPsqlRepository(conn *connection.Pool) Base {
	return &PSQLRepository{pool: conn}
}
