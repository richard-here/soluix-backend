package model

type GetQueries struct {
	Limit       int    `validate:"omitempty,numeric,min=1,max=100"`
	Page        int    `validate:"omitempty,numeric,min=1"`
	Sort        string `validate:"omitempty,oneof=id+desc id+asc id%20desc id%20asc retail_price%20asc retail_price+asc retail_price%20desc retail_price+desc"`
	Name        string `validate:"omitempty"`
	Subcategory string `validate:"omitempty"`
	Brand       string `validate:"omitempty"`
	MinPrice    int    `validate:"omitempty,numeric,min=0,max=50000000"`
	MaxPrice    int    `validate:"omitempty,numeric,min=0,max=50000000,gtefield=MinPrice"`
	Status      string `validate:"omitempty,oneof=0 1"`
}
