package repository

import "golang/orm/models"

type Recordset struct {
	IsSuccessful bool
	data         []map[string]any
}

func (r *Recordset) AppendRow(record map[string]any) {
	r.data = append(r.data, record)
}

type OperationResult struct {
	IsSuccessful bool
	Id           int
}

type Base interface {
	Get(table string, model models.Model) Recordset
	Create(data map[string]any) OperationResult
	Update(id int, data map[string]any) OperationResult
	Delete(id int) OperationResult
}
