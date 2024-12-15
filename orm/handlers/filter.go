package handlers

import (
	"golang/orm/client"
	"golang/orm/models"
	"golang/orm/repository"
)

func Filter(dbClient *client.DbClient, table string, model models.Model) (repository.Recordset, error) {
	return dbClient.Get(table, model), nil
}
