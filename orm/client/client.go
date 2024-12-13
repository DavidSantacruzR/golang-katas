package client

import (
	"golang/orm/connection"
	"golang/orm/models"
	"golang/orm/repository"
)

type DbClient struct {
	repository repository.Base
}

func NewDbClient(connDetails connection.Connection) DbClient {
	return DbClient{repository.NewPsqlRepository(connDetails.ConnectionString())}
}

func (db *DbClient) Get(table string, model models.Model) repository.Recordset {
	return db.repository.Get(table, model)
}
