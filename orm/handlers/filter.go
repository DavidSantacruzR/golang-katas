package handlers

import (
	"golang/orm/client"
	"golang/orm/repository"
)

func Filter(dbClient *client.DbClient, model any, fields map[string]any) (repository.Recordset, error) {
	/*TODO: handle logic validation that model fulfills certain criteria otherwise raise errors.*/
	return dbClient.Get(model, fields), nil
}
