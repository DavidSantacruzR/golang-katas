package client

import (
	"golang/orm/connection"
	"golang/orm/repository"
)

type DbClient struct {
	repository repository.Base
}

func NewDbClient(connections int) DbClient {
	pool, _ := connection.StartConnectionPool(connections)
	if pool == nil {
		panic("connection pool is nil")
	}
	return DbClient{repository.NewPsqlRepository(pool)}
}

func (db *DbClient) Get(model any, fields map[string]any) repository.Recordset {
	return db.repository.Get(model, fields)
}
