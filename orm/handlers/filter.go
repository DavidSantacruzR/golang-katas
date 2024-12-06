package handlers

import (
	"golang/orm/client"
	"golang/orm/models"
	"golang/orm/repository"
	"reflect"
)

func Filter(dbClient *client.DbClient, model any, fields map[string]any) (repository.Recordset, error) {
	m := reflect.TypeOf(model)
	field, found := m.FieldByName("Model")
	if !found {
		return repository.Recordset{}, &MissingModelTypeError{message: "No model characteristic found."}
	}
	if field.Type != reflect.TypeOf(models.Model{}) {
		return repository.Recordset{}, &WrongModelTypeError{message: "Missing model type, expected models.Model"}
	}
	return dbClient.Get(model, fields), nil
}
