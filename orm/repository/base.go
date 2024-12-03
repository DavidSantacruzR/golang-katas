package repository

type Recordset struct {
	IsSuccessful bool
	data         map[string]any
}

type OperationResult struct {
	IsSuccessful bool
	Id           int
}

type Base interface {
	Get(model any, fields map[string]any) Recordset
	Create(data map[string]any) OperationResult
	Update(id int, data map[string]any) OperationResult
	Delete(id int) OperationResult
}
