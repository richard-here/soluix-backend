package model

type GetQueries struct {
	Limit int    `validate:"omitempty,numeric,min=1,max=100"`
	Page  int    `validate:"omitempty,numeric,min=0"`
	Sort  string `validate:"omitempty,oneof=id+desc id+asc id%20desc id%20asc"`
}
