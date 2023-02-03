package model

import (
	localmodel "richard-here/soluix/product-api/model/local"
)

type AddBody struct {
	Name        string `json:"name" validate:"required"`
	Subcategory string `json:"subcategory" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	RetailPrice int    `json:"retail_price" validate:"required,numeric,min=1,max=50000000"`
	Status      bool   `json:"status" validate:"required"`
}

func (p *AddBody) ToLocalModel() *localmodel.Product {
	lp := new(localmodel.Product)
	lp.Name = p.Name
	lp.Subcategory = p.Subcategory
	lp.Brand = p.Brand
	lp.RetailPrice = p.RetailPrice
	lp.Status = p.Status
	return lp
}
